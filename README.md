# 🗓️ goatimtable

[![Go Reference](https://pkg.go.dev/badge/github.com/benzjeremy/goatimtable.svg)](https://pkg.go.dev/github.com/benzjeremy/goatimtable)
[![Go Report Card](https://goreportcard.com/badge/github.com/benzjeremy/goatimtable.svg)](https://goreportcard.com/report/github.com/benzjeremy/goatimtable)
[![Release](https://img.shields.io/github/v/release/benzjeremy/goatimtable)](https://github.com/benzjeremy/goatimtable/releases)
[![Status: Pre-Release](https://img.shields.io/badge/Status-Pre--Release%20%2F%20WIP-orange.svg)](https://github.com/benzjeremy/goatimtable)
[![License: GPL-3.0](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20Windows-lightgrey)](#installation)

> 🌐 **Official Website:** [https://goatimtable.benzjeremy.pp.ua/](https://goatimtable.benzjeremy.pp.ua/)
> 
> 📖 **Official Wiki & Documentation:** [https://goatimtable.benzjeremy.pp.ua/wiki/](https://goatimtable.benzjeremy.pp.ua/wiki/)

> [!IMPORTANT]
> ### 🚧 Pre-Release / Active Development Notice
> **This software is not yet finished and is actively being worked on.**  
> All versions (including `v2.3`) are **Pre-Releases** (Work in Progress), even if not originally announced as such. Active development, architectural refinements, and feature updates are ongoing.  
> If you encounter **errors, display bugs, or unexpected behavior**, please open an issue directly under [**GitHub Issues**](https://github.com/benzjeremy/goatimtable/issues)! Feedback and pull requests are warmly welcome.

---

## 🎯 What is goatimtable?

A fast, native, and modern **WebUntis PC desktop client for students and teachers** – written in Go with a sleek desktop shell (inspired by clean sidebars of apps like Discord and Spotify Desktop).

Forget slow web interfaces, cluttered layouts, or resource-heavy web wrappers. **goatimtable** brings your timetable, substitution schedules, homework, announcements, and absences lightning-fast and cross-platform to your desktop – **without Electron bloat!**

---

## ✨ Features

- 🏠 **Overview (Dashboard)**:
  - Personal greeting with name and school.
  - Direct display of next upcoming lesson and daily timeline.
  - Quick overview of open homework assignments and unread messages.
- 📅 **My Timetable**:
  - Personal student or teacher timetable.
  - Day and week view with **red LIVE-TIMELINE**.
  - Click on any lesson opens detailed information (teachers, substitutions, rooms, subject, homework).
- 👥 **Other Timetables**:
  - Quick switching between **all classes** (with live search over hundreds of classes), **teachers**, and **classrooms**.
- 📝 **Homework Management**:
  - Overview of all homework from WebUntis plus manually added tasks.
  - Tick off tasks, filter, and add new homework via a modal dialog.
- 🩺 **Absences**:
  - Absence list with status (Excused / Unexcused), period, and reason.
  - Enter new absences / sick notes.
- 💬 **Messages (Message Center)**:
  - Full-featured inbox for all official school messages, parent letters, and teacher announcements with full-text display.
- ⚙️ **Schools & Profile Management**:
  - Manage any number of schools and user profiles in parallel.
  - **Delete schools**: Each profile can be removed with a single click.
  - **Live school search**: Find your school worldwide via the official WebUntis school search.
- 🛡️ **100% Local-First & Privacy Compliant (GDPR / DSGVO)**:
  - All data stays strictly local on your machine in an encrypted SQLite database (`~/.local/share/untis-go/untis.db`).
  - Zero telemetry, zero tracking, zero external third-party cloud dependencies.
  - Full compliance with European GDPR (DSGVO) and State Data Protection Authority (LDI NRW) standards.
- 📅 **RFC-5545 iCalendar (.ics) Export**:
  - 1-click timetable export for Google Calendar, Apple Calendar & Mozilla Thunderbird.
  - Generates standard-compliant VEVENT blocks with room, teachers, notes, and cancellation status.
- 🔔 **Background Sync & Desktop Notifications**:
  - Automatic background monitoring daemon detecting timetable mutations (room changes, cancellations, substitutions).
  - Native OS desktop notifications (via Linux `notify-send` and Windows toast).
- 🖥️ **Native System Tray Integration**:
  - Minimize window to system tray / taskbar instead of quitting.
  - Quick menu for one-click schedule access, force refresh, status check, and graceful exit.

---

## 🔒 Security & Privacy (Zero-Telemetry)

- **100% Local-First & Privacy (GDPR / LDI NRW)**: Zero telemetry, zero tracking. All communication occurs directly with your school's WebUntis instance.
- **AES-256-GCM Encryption**: Credentials and cached secrets are never stored in plain text.
- **SQLite Cache**: Timetables load in under 1 ms directly from local storage.
- **Random Port & Crypto Session Token**: Protection against unauthorized local access (Strict Anti-DNS-Rebinding & Anti-CSRF).

---

## 📦 Installation & Download

### 1. Download Precompiled Packages (Recommended)

Download the matching file for your operating system from the [**Releases**](https://github.com/benzjeremy/goatimtable/releases) page:

- **Linux (x86_64)**:
  ```bash
  tar -xzf goatimtable-v2.5-linux-amd64.tar.gz
  sudo cp goatimtable /usr/local/bin/
  goatimtable
  ```
- **Windows (x86_64)**:
  - Unzip `goatimtable-v2.5-windows-amd64.zip` and run `goatimtable.exe`.

### 2. Installation via Go (`go install`)

If you have Go (version 1.21 or newer) installed:

```bash
go install github.com/benzjeremy/goatimtable@latest
```

The binary will be compiled automatically into your `$GOPATH/bin` (or `~/go/bin`) and can be called directly as `goatimtable` in your terminal.

### 3. Compile from Source

#### Prerequisites (Linux):
- Go 1.21+
- GTK 3 & WebKitGTK development libraries:
  - **Arch Linux / CachyOS**: `sudo pacman -S webkit2gtk-4.1 gtk3 gcc`
  - **Ubuntu / Debian**: `sudo apt install libwebkit2gtk-4.1-dev libgtk-3-dev build-essential`
  - **Fedora**: `sudo dnf install webkit2gtk4.1-devel gtk3-devel gcc`

#### Building:
```bash
# 1. Clone the repository
git clone https://github.com/benzjeremy/goatimtable.git
cd goatimtable

# 2. Build binary
go build -o goatimtable .

# 3. Launch
./goatimtable
```

---

## ⌨️ Keyboard Shortcuts

| Key | Action |
|---|---|
| `←` / `→` | Previous / Next day (or week) |
| `T` | Jump to **Today** |
| `D` | Activate **Day view** |
| `W` | Activate **Week view** |
| `Esc` | Close open dialogs, info-sheets, and menus |

---

## 🐛 Bug Reports & Contributing

Bug reports and contributions are welcome:
1. Open the [**Issues**](https://github.com/benzjeremy/goatimtable/issues) tab.
2. Click **New Issue**.
3. Briefly describe:
   - Operating system and desktop environment.
   - Steps to reproduce.
   - Expected vs actual behavior (including console log output).

---

## 📚 Wiki & Documentation

Detailed documentation and step-by-step guides are available in our official web wiki:  
👉 **[goatimtable Wiki: https://goatimtable.benzjeremy.pp.ua/wiki/](https://goatimtable.benzjeremy.pp.ua/wiki/)**

- **Getting Started & Onboarding**: [Getting Started Guide](https://goatimtable.benzjeremy.pp.ua/wiki/#quickstart)
- **Feature Deep Dive**: [Timetables, Homework & Absences](https://goatimtable.benzjeremy.pp.ua/wiki/#features)
- **Installation**: [Linux & Windows Setup](https://goatimtable.benzjeremy.pp.ua/wiki/#installation)
- **Security Architecture**: [AES-256-GCM & PBKDF2](https://goatimtable.benzjeremy.pp.ua/wiki/#security)
- **FAQ & Troubleshooting**: [Frequently Asked Questions](https://goatimtable.benzjeremy.pp.ua/wiki/#faq)

---

## 📄 License & Author

- **Developer:** Jeremy Benz ([@benzjeremy](https://github.com/benzjeremy))
- **License:** [GNU General Public License v3.0 (GPL-3.0)](LICENSE)

## Project rename

untis-go is now **goatimtable**. This project remains in development (pre-release). Existing local data continues to use the legacy storage directory for compatibility.
Users upgrading from v2.4.1 or earlier must download this release manually: the old updater only accepts the former repository path and executable name.

## A personal note from Jeremy Benz

> I decided we needed to rename untis-go to goatimtable to avoid potential intellectual-property and trademark issues. The previous name directly referenced Untis GmbH's brand. I wanted to make this change early, before it could lead to legal trouble. Same project, new name — thank you for sticking with it.
>
> — Jeremy Benz, project creator

Windows icon resource: regenerate `app_windows_amd64.syso` after icon changes with `x86_64-w64-mingw32-windres app.rc -O coff -o app_windows_amd64.syso`.
