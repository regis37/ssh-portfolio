package ui

type Section struct {
	Title   string
	Content string
}

var sections = []Section{
	{
		Title: "Home",
		Content: `Regis Safack
────────────────────────────────────────
Electrical & Information Engineering · Software

Student in Elektrotechnik und Informationstechnik (B.Eng.)
4th Semester · TH Nürnberg

Werkstudent System Engineering
@ Siemens Healthineers · Forchheim

Based in Nuremberg, Germany

Open to opportunities and collaborations.`,
	},
	{
		Title: "Stack",
		Content: `Languages
────────────────────────────────────────
  C++          ████████░░  Advanced
  TypeScript   ███████░░░  Proficient
  Java         ██████░░░░  Proficient
  Python       ██████░░░░  Proficient
  MATLAB       █████░░░░░  Intermediate

Frameworks & Tools
────────────────────────────────────────
  Angular      Frontend framework (TypeScript)
  Spring Boot  Java backend framework
  Git          Version control
  CMake        C++ build system
  Linux        Daily driver OS`,
	},
	{
		Title: "Experience",
		Content: `Werkstudent — System Engineering
Siemens Healthineers · Forchheim
────────────────────────────────────────
  · Signal analysis for medical imaging systems
  · Control systems design & simulation
  · MATLAB/Simulink: PID tuning, system identification
  · Documentation of embedded system behavior

Math & EE Tutor
────────────────────────────────────────
  · Private tutoring in mathematics and electrical
    engineering fundamentals
  · Topics: circuit analysis, signals & systems,
    linear algebra, differential equations`,
	},
	{
		Title: "Projects",
		Content: `tcp-chat-cpp
────────────────────────────────────────
  Multi-client TCP chat server in C++
  · Winsock2 networking, multithreaded architecture
  · Custom protocol, room management

measurement-data-analyzer
────────────────────────────────────────
  Python data analysis pipeline
  · pandas + matplotlib for CSV measurement data
  · Automated report generation

cpp-inventory-management-system
────────────────────────────────────────
  C++ OOP inventory system
  · SQLite persistence layer
  · CLI interface with CRUD operations

file-sharing-app
────────────────────────────────────────
  AirDrop-style file sharing
  · Java Spring Boot backend
  · Angular frontend
  · WebSocket real-time relay`,
	},
	{
		Title: "Focus",
		Content: `Signal Processing
────────────────────────────────────────
  · Bode plots, frequency response analysis
  · Filter design (low-pass, high-pass, band-pass)
  · Fourier, Laplace, and Z transforms

Control Systems
────────────────────────────────────────
  · PT1/PT2 systems, step response analysis
  · PID controller design & tuning
  · System identification techniques

Network Programming
────────────────────────────────────────
  · TCP/UDP socket programming
  · Multi-threaded server architecture
  · Protocol design

Backend Development
────────────────────────────────────────
  · RESTful APIs with Spring Boot
  · WebSocket real-time communication
  · Database integration (SQLite, PostgreSQL)`,
	},
	{
		Title: "Education",
		Content: `B.Eng. Elektrotechnik und Informationstechnik
TH Nürnberg · Georg Simon Ohm
────────────────────────────────────────
  Semester:   4th (ongoing)
  Location:   Nuremberg, Germany
  Focus:      Signals & Systems, Embedded Systems,
              Control Theory, Digital Electronics

Key Courses
────────────────────────────────────────
  · Signals and Systems (Signale und Systeme)
  · Control Engineering (Regelungstechnik)
  · Embedded Systems Programming
  · Digital Signal Processing
  · Software Engineering`,
	},
	{
		Title: "Contact",
		Content: `Regis Safack
────────────────────────────────────────

  GitHub    github.com/regis37

  Location  Nuremberg, Germany

  Status    Open to opportunities
            and collaborations

────────────────────────────────────────
Feel free to reach out via GitHub.`,
	},
}
