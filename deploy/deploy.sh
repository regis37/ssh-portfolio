#!/usr/bin/env bash
# Reproducible deployment script for ssh-portfolio
# Usage: SSH_KEY=~/.ssh/oracle.key ./deploy/deploy.sh
set -euo pipefail

REMOTE_USER="ubuntu"
REMOTE_HOST="152.70.21.255"
REMOTE_DIR="/opt/portfolio"
LOG_DIR="/var/log/portfolio"
SSH_KEY="${SSH_KEY:-$HOME/.ssh/oracle.key}"
SSH="ssh -i $SSH_KEY -o StrictHostKeyChecking=no ${REMOTE_USER}@${REMOTE_HOST}"
SCP="scp -i $SSH_KEY -o StrictHostKeyChecking=no"

echo "==> Cross-compiling static Linux/amd64 binaries..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o portfolio  ./cmd/portfolio
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o stats      ./cmd/stats

echo "==> Stopping service (releases binary file lock)..."
$SSH "sudo systemctl stop portfolio 2>/dev/null || true"

echo "==> Ensuring remote directories..."
$SSH "sudo mkdir -p ${REMOTE_DIR} ${LOG_DIR} && \
      sudo chown ${REMOTE_USER}:${REMOTE_USER} ${REMOTE_DIR} ${LOG_DIR} && \
      sudo chmod 750 ${LOG_DIR}"

echo "==> Copying binaries..."
$SCP portfolio stats "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/"
$SSH "chmod +x ${REMOTE_DIR}/portfolio ${REMOTE_DIR}/stats"

echo "==> Installing systemd service..."
$SCP deploy/portfolio.service "${REMOTE_USER}@${REMOTE_HOST}:/tmp/portfolio.service"
$SSH "sudo mv /tmp/portfolio.service /etc/systemd/system/portfolio.service && \
      sudo systemctl daemon-reload && \
      sudo systemctl enable portfolio"

echo "==> Generating persistent log salt (once — skipped if /opt/portfolio/.env exists)..."
$SSH "
  if [ ! -f ${REMOTE_DIR}/.env ]; then
    SALT=\$(openssl rand -hex 32)
    printf 'PORTFOLIO_LOG_SALT=%s\n' \"\$SALT\" > ${REMOTE_DIR}/.env
    chmod 600 ${REMOTE_DIR}/.env
    echo 'Salt generated and saved to ${REMOTE_DIR}/.env'
  else
    echo 'Salt already exists — keeping it (hashes remain consistent across deploys)'
  fi
"

echo "==> Installing logrotate config..."
$SCP deploy/logrotate-portfolio "${REMOTE_USER}@${REMOTE_HOST}:/tmp/logrotate-portfolio"
$SSH "sudo mv /tmp/logrotate-portfolio /etc/logrotate.d/portfolio && \
      sudo chmod 644 /etc/logrotate.d/portfolio"

echo "==> Ensuring iptables rule for TCP 22 (portfolio)..."
$SSH "
  if ! sudo iptables -C INPUT -p tcp --dport 22 -m state --state NEW -j ACCEPT 2>/dev/null; then
    REJECT_LINE=\$(sudo iptables -L INPUT --line-numbers -n | awk '/REJECT/{print \$1; exit}')
    sudo iptables -I INPUT \"\${REJECT_LINE:-5}\" -p tcp --dport 22 -m state --state NEW -j ACCEPT
    sudo netfilter-persistent save
    echo 'iptables ACCEPT rule added for port 22'
  else
    echo 'iptables rule for port 22 already present'
  fi
"

echo "==> Starting service..."
$SSH "sudo systemctl restart portfolio"
sleep 2

echo "==> Service status:"
$SSH "sudo systemctl status portfolio --no-pager -l && sudo ss -tlnp | grep ':22 '"

echo ""
echo "==> Verify sha256 (local vs remote):"
LOCAL_SUM=$(sha256sum portfolio | awk '{print $1}')
REMOTE_SUM=$($SSH "sha256sum ${REMOTE_DIR}/portfolio | awk '{print \$1}'")
echo "  local : $LOCAL_SUM"
echo "  remote: $REMOTE_SUM"
if [ "$LOCAL_SUM" = "$REMOTE_SUM" ]; then
  echo "  MATCH ✓"
else
  echo "  MISMATCH ✗ — deploy may have failed"
  exit 1
fi

echo ""
echo "==> Stats binary on server:"
$SSH "${REMOTE_DIR}/stats 2>/dev/null || echo '(no visits yet)'"

echo ""
echo "==> Done. Connect with:"
echo "    ssh registsafack.duckdns.org"
echo ""
echo "==> Read stats on server:"
echo "    ssh ubuntu@152.70.21.255 /opt/portfolio/stats"
