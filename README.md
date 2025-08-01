# ⚡ Fany CLI – Bootstrap NineVerse Projects

**Fany** is a **tiny, fast, cross-platform CLI** to scaffold [NineVerse](https://github.com/Appsventory/NineVerse) – a lightweight PHP-MVC starter template, without the overhead of large frameworks.

---

## 📥 Installation (v1.0)

### 1. Download the latest release

Visit the [Releases](https://github.com/Appsventory/fany-cli/releases) page and download the appropriate binary for your system:

| OS      | File       |
| ------- | ---------- |
| Windows | `fany.exe` |
| macOS   | `fany`     |
| Linux   | `fany`     |

---

### 2. Add to your system `PATH`

#### **Windows (PowerShell)**

```powershell
cd path\to\download
.\fany.exe install
```

#### **macOS / Linux (bash)**

```bash
cd /path/to/download
./fany install
```

> Restart your terminal after installation.

---

### 3. Manual Build (Optional)

```bash
git clone https://github.com/Appsventory/fany-cli.git
cd fany-cli

# Windows
go build -o fany.exe ./cmd/fany
.\fany.exe install

# macOS / Linux
go build -o fany ./cmd/fany
./fany install
```

---

## 🚀 Usage

| Task               | Command                          |
| ------------------ | -------------------------------- |
| Create new project | `fany new Blog --dir ~/Projects` |
| Force online pull  | `fany get-new Blog`              |
| Clone repository   | `fany git-clone Appsventory/NineVerse` |
| Initialize cache   | `fany cache-init ./NineVerse`    |
| Update cache       | `fany cache-update`              |
| Install Git        | `fany git-install`               |
| Install Composer   | `fany cp-install`                |
| Upgrade Composer   | `fany cp-upgrade`                |
| Show version       | `fany --v`                       |
| Help               | `fany --help`                    |

---

## 🔧 Inside a NineVerse Project

```bash
fany get-new MyApp
cd MyApp

# Create a new controller named User
./fany make:controller User

# Start local dev server on port 8080
./fany server --p 8080
```

---

## 🧹 Uninstall

```bash
fany uninstall        # Remove installed binary

# Or manually:
rm $(which fany)
```

---

## 🐞 Issues & Contributions

Found a bug or want to contribute? Open an issue or pull request at:
👉 [https://github.com/Appsventory/fany-cli](https://github.com/Appsventory/fany-cli)

---

MIT License © ICK Network

---

Let me know if you'd like a minimal version for publishing on package managers (like Homebrew, Scoop, etc.) or if you want a website/docs layout.
