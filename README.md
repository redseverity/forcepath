# forcepath

**forcepath** is an automatic directory discovery tool written in Go that works without relying on static wordlists.

Unlike traditional tools that depend on large, predefined wordlists, **forcepath** takes a brute-force approach: it dynamically generates and tests subdirectory names in real time. This avoids the need to load massive lists into memory, making it faster, more memory-efficient, and suitable for adaptive directory enumeration.

Ideal for reconnaissance, Capture The Flag (CTF) challenges, and general web enumeration tasks.

---

## 📦 Installation

To install `forcepath`, you need to have [Go](https://go.dev/dl/) installed (version 1.16 or higher recommended).

Run the following command to install the tool:

```bash
go install github.com/redseverity/forcepath@latest
```

---

## ⚙️ Flags Guide

### `-url`

Defines the target URL to scan.

**Accepted formats:**

```
https://example.com/
http://example.com
example.com
```

If the URL doesn't include a scheme, the tool will **automatically prepend `https://`**.

---

### `-charset`

Specifies the set of characters used for generating directory names.

**Allowed characters:**

```
a-z  A-Z  0-9  -  _  .  ~
```

* Duplicate characters are **automatically removed**.
* If omitted, the default charset is: `"abc123"`

**Example:**

```bash
-charset "abbcc1232"  → becomes "abc123"
```

---

### `-min`

Sets the **minimum length** of generated directory names.
Default: `1`

### `-max`

Sets the **maximum length** of generated directory names.
Default: `3`

> The generator will start with the smallest combination (`min` length) and go up to the largest (`max` length).

---

### `-timeout`

Sets the timeout (in seconds) for each HTTP request.
Default: `3`

---

### `-delay`

Sets the delay (in milliseconds) between each HTTP request.  
Default: `0`

---

### `-help`

Displays the help message with usage examples and exits.

---

## ⚙️ Behavior

* Automatically adds `https://` to URLs missing a scheme.
* Removes duplicate characters in the `-charset` input.
* Generates all combinations of characters from `min` to `max` length.
* Attempts to connect to the **host** before starting.
* Tests each generated path individually and checks for **valid HTTP responses**.
* Applies a timeout to each request as defined by the `-timeout` flag.
* If the `-delay` flag is set (> 0), waits the specified milliseconds between
each request to avoid rate limiting.
* Validates required flags and exits with helpful error messages if any are missing.

---

## 🚀 Usage Examples

Basic syntax:

```bash
forcepath -url <target_url> -charset <charset> [flags]
```

Run with defaults:

```bash
forcepath -url https://example.com
```

Custom charset with specific length range:

```bash
forcepath -url https://example.com -charset "abc123-_.~" -min 2 -max 2 -timeout 5 -delay 1
```

Show help:

```bash
forcepath -help
```

---

## 📚 Learn More

For more information, source code, updates, or to contribute, visit the repository:
👉 [https://github.com/redseverity/forcepath](https://github.com/redseverity/forcepath)

