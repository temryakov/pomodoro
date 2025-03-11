# Pomodoro CLI

Pomodoro CLI is a command-line utility for time management using the Pomodoro Technique. It allows you to run 25-minute work sessions followed by short breaks to enhance productivity.

▎Key Features:

• Start, pause, and end Pomodoro sessions.

• Customize session durations.

• Manage breaks.

A simple and convenient tool to boost your work efficiency!

```bash
pomodoro start
pomodoro start -d, --duration int # Time duration of pomodoro (default 25)
pomodoro start -s, --set int # Default set time of pomodoro (default 25)

pomodoro break
pomodoro break  -d, --duration int # Time duration of break (default 5)
pomodoro break -s, --set int # Default set time of break (default 5)

pomodoro --help
```
