# forcepath

**forcepath** is an automatic directory discovery tool written in Go that works without relying on static wordlists.

Unlike traditional tools that depend on large, predefined wordlists, **forcepath** takes a brute-force approach: it dynamically generates and tests subdirectory names in real time. This avoids the need to load massive lists into memory, making it faster, more memory-efficient, and suitable for adaptive directory enumeration.

Perfect for reconnaissance, Capture The Flag (CTF) challenges, and general web enumeration tasks.

---

## 📦 Installation

To install `forcepath`, you need to have [Go](https://go.dev/dl/) installed (version 1.16 or higher recommended).

Run the following command to install the tool:

```bash
go install github.com/redseverity/forcepath@latest
```

---

## 🚀 Usage

Basic syntax
```bash
forcepath -url <target_url> [-charset <charset>] [flags]
```

Scan a website with default settings (charset: abc123)
```bash
forcepath -url "https://example.com"
```

Specify a custom charset for brute forcing
```bash
forcepath -url "https://example.com" -charset "abc123+-.=?"
```

Show help
```bash
forcepath -h
```
