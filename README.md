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

## 🚀 Guide

#### Flag -url

Allowed formats:

    https://example.com/

    http://example.com

    example.com

If the URL does not include a scheme (like http:// or https://), the tool will automatically add the https:// scheme.

#### Flag -charset

Allowed Characters for -charset:

    a-z  A-Z  0-9  -  _  .  ~

Duplicate characters in the -charset value will be automatically removed.
For example, if you provide -charset "abbcc1232", it will be treated as "abc123".

---

## 🚀 Usage

Basic syntax
```bash
forcepath -url <target_url> -charset <charset>
```

Scan a website with default settings (charset: abc123)
```bash
forcepath -url "https://example.com"
```

Specify a custom charset for brute forcing
```bash
forcepath -url "https://example.com" -charset "abc123-_.~"
```

Show help
```bash
forcepath -h
```
