# Rinaco Elevate (CLI-G)
# Rinaco : Elevate Your Habit.
# Version : v1.5.19-B3 [FIB] #7a2b9f1

A terminal-based recovery tracker with real-time countdown.

## Features

- Track recovery progress
- Dynamic real-time countdown
- Session history
- Standard recovery plan guide
- Trigger note tracking
- Works everywhere: Windows, Linux, Android Termux

---

## Installation & Usage

### Platform 1: Android (Termux)

**First Time Setup (once)**

1. Download Termux
   - From F-Droid (recommended): https://f-droid.org/en/packages/com.termux/
   - Or Google Play (older version)

2. Open Termux and run:

```
pkg update
pkg install golang git
```

Wait for installation. If asked, type `y` and press Enter.

3. Create the project folder and set it up (only once):

```
mkdir ~/rinaco
cd ~/rinaco
go mod init rinaco-elevate
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
```

**Run the App** (every time you want to use it)

```
cd ~/rinaco
go run .
```

Or build it once and run the binary:

```
cd ~/rinaco
go build
./rinaco-elevate
```

After building, next time just run:

```
cd ~/rinaco
./rinaco-elevate
```

---

### Platform 2: Linux (Ubuntu / Debian / Fedora / Arch)

**First Time Setup (once)**

1. Install Go (if not installed)

```
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go git

# Fedora/RedHat
sudo dnf install golang git

# Arch
sudo pacman -S go git
```

2. Create the project folder and set it up (only once):

```
mkdir ~/rinaco
cd ~/rinaco
go mod init rinaco-elevate
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
```

**Run the App** (every time you want to use it)

```
cd ~/rinaco
go run .
```

Or build it once and run the binary:

```
cd ~/rinaco
go build
./rinaco-elevate
```

After building, next time just run:

```
cd ~/rinaco
./rinaco-elevate
```

---

### Platform 3: Windows

**First Time Setup (once)**

1. Download and Install Go
   - Go to: https://golang.org/dl/
   - Download Windows installer (.msi file)
   - Run installer, click Next then Finish
   - Restart your computer

2. Download and Install Git
   - Go to: https://git-scm.com/download/win
   - Download Windows installer
   - Run installer, click Next then Finish

3. Open PowerShell or Command Prompt
   - Press `Win + R`
   - Type `cmd` or `powershell`
   - Press Enter

4. Create the project folder and set it up (only once):

```
mkdir C:\Users\YOUR-USERNAME\rinaco
cd C:\Users\YOUR-USERNAME\rinaco
go mod init rinaco-elevate
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
```

**Run the App** (every time you want to use it)

```
cd C:\Users\YOUR-USERNAME\rinaco
go run .
```

Or build it once and run the binary:

```
cd C:\Users\YOUR-USERNAME\rinaco
go build
rinaco-elevate.exe
```

After building, next time just run:

```
cd C:\Users\YOUR-USERNAME\rinaco
rinaco-elevate.exe
```

---

## File Structure

```
rinaco-elevate/
├── main.go           Main entry point
├── models.go         Data structures
├── styles.go         UI colors & design
├── helpers.go        Utility functions
├── validation.go     Input validation
├── storage.go        Load data
├── plan_data.go      Recovery plan content
├── tea_setup.go      App initialization
├── update.go         Handle user input
├── views.go          Render UI screens
├── go.mod            Dependencies list
├── go.sum            Dependency checksums
└── README.md         Current file
```

---

## Controls

| Key | Action           |
|-----|------------------|
| s   | Set start time   |
| r   | Mark as relapsed |
| h   | View history     |
| p   | View plan        |
| t   | View statistics  |
| j   | View triggers    |
| e   | Edit profile     |
| q   | Quit             |

---

## Data Storage

Your progress is saved in: `rinaco-elevate_data.json`

Keep this file in the same folder as the app.

---

## Troubleshooting

### Problem: "command not found: go"

Fix:
- Make sure Go is installed
- Restart terminal or PowerShell after installation
- On Linux, you may need to add Go to PATH:

```
export PATH=$PATH:/usr/local/go/bin
```

---

### Problem: "no required module provides package github.com/charmbracelet/bubbletea"

Fix:

```
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go run .
```

---

### Problem: Can't find the app after building

Android / Linux:

```
ls -la
./rinaco-elevate
```

Windows:

```
dir
rinaco-elevate.exe
```

---

## Development

To modify and test:

1. Edit any `.go` file
2. Run: `go run .`
3. Changes take effect immediately

---

## Building for Distribution

Android / Linux:

```
go build -o rinaco-elevate
```

Creates: `rinaco-elevate` (executable)

Windows:

```
go build -o rinaco-elevate.exe
```

Creates: `rinaco-elevate.exe` (executable)

---

## Contributing

Found a bug or have an idea?

1. Fork the repository
2. Create a branch: `git checkout -b feature/your-feature`
3. Make your changes
4. Commit: `git commit -m "description"`
5. Push: `git push origin feature/your-feature`
6. Create a Pull Request

---

## License

Free to use, modify, and share.

Made with 🤩.
