#!/usr/bin/env bash
# Reproducible deployment script for ssh-portfolio
# Usage: ./deploy/deploy.sh
set -euo pipefail

REMOTE_USER="ubuntu"
REMOTE_HOST="152.70.21.255"
REMOTE_DIR="/opt/portfolio"
SSH_KEY="${SSH_KEY:-$HOME/.ssh/oracle.key}"
BINARY="portfolio"

echo "==> Cross-compiling static Linux/amd64 binary..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o "$BINARY" ./cmd/portfolio

echo "==> Stopping service (releases binary file lock)..."
ssh -i "$SSH_KEY" -o StrictHostKeyChecking=no "${REMOTE_USER}@${REMOTE_HOST}" \
  "sudo systemctl stop portfolio 2>/dev/null || true"

echo "==> Copying binary to ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/"
ssh -i "$SSH_KEY" -o StrictHostKeyChecking=no "${REMOTE_USER}@${REMOTE_HOST}" \
  "sudo mkdir -p ${REMOTE_DIR} && sudo chown ${REMOTE_USER}:${REMOTE_USER} ${REMOTE_DIR}"
scp -i "$SSH_KEY" -o StrictHostKeyChecking=no "$BINARY" \
  "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/${BINARY}"
ssh -i "$SSH_KEY" -o StrictHostKeyChecking=no "${REMOTE_USER}@${REMOTE_HOST}" \
  "chmod +x ${REMOTE_DIR}/${BINARY}"

echo "==> Installing systemd service..."
scp -i "$SSH_KEY" -o StrictHostKeyChecking=no deploy/portfolio.service \
  "${REMOTE_USER}@${REMOTE_HOST}:/tmp/portfolio.service"
ssh -i "$SSH_KEY" -o StrictHostKeyChecking=no "${REMOTE_USER}@${REMOTE_HOST}" \
  "sudo mv /tmp/portfolio.service /etc/systemd/system/portfolio.service && \
   sudo systemctl daemon-reload && \
   sudo systemctl enable portfolio && \
   sudo systemctl restart portfolio"

echo "==> Ensuring iptables rule for TCP 23234..."
ssh -i "$SSH_KEY" -o StrictHostKeyChecking=no "${REMOTE_USER}@${REMOTE_HOST}" "
  if ! sudo iptables -C INPUT -p tcp --dport 23234 -m state --state NEW -j ACCEPT 2>/dev/null; then
    REJECT_LINE=\$(sudo iptables -L INPUT --line-numbers -n | awk '/REJECT/{print \$1; exit}')
    sudo iptables -I INPUT \"\${REJECT_LINE:-5}\" -p tcp --dport 23234 -m state --state NEW -j ACCEPT
    sudo netfilter-persistent save
  else
    echo 'iptables rule already present'
  fi
"

echo "==> Service status:"
ssh -i "$SSH_KEY" -o StrictHostKeyChecking=no "${REMOTE_USER}@${REMOTE_HOST}" \
  "sudo systemctl status portfolio --no-pager && sudo ss -tlnp | grep 23234"

echo ""
echo "==> Done. Connect with:"
echo "    ssh -p 23234 registsafack.duckdns.org"
