# gosubfinder

**gosubfinder** is an automatic directory discovery tool written in Go that works without relying on static wordlists.

Unlike traditional tools that depend on large, predefined wordlists, **gosubfinder** takes a brute-force approach: it dynamically generates and tests subdirectory names in real time. This avoids the need to load massive lists into memory, making it faster, more memory-efficient, and suitable for adaptive directory enumeration.

Perfect for reconnaissance, Capture The Flag (CTF) challenges, and general web enumeration tasks.

---

## 📦 Installation

To install `gosubfinder`, you need to have [Go](https://go.dev/dl/) installed (version 1.16 or higher recommended).

Run the following command to install the tool:

```bash
go install github.com/redseverity/gosubfinder@latest
