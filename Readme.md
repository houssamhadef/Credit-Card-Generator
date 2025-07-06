# 💳 CCGen - Credit Card Generator CLI

![Go Version](https://img.shields.io/badge/Go-1.22+-brightgreen?logo=go)
![License](https://img.shields.io/github/license/yourusername/ccgen)
![Stars](https://img.shields.io/github/stars/yourusername/ccgen?style=social)

> A blazing-fast, Luhn-valid credit card generator built with Go. Supports BIN targeting, expiration dates, CVV generation, and CLI control.

---

## 🚀 Features

- ✅ Generate Luhn-valid credit card numbers
- 🏦 Support for BIN prefix (e.g. `471300`)
- 📆 Random expiration date generator (up to 5 years ahead)
- 🔒 CVV (3-digit) generator
- 📁 Saves to `logs/` with timestamped folders
- ⚡ Fast and memory-efficient
- 🧠 Learn how credit card checksums work via Luhn algorithm
- 🖥️ Fully functional CLI using [Cobra](https://github.com/spf13/cobra) (Soon..)

---

## 📦 Installation

### Clone & Build:

```bash
git clone https://github.com/yourusername/ccgen.git
cd ccgen
go build -o ccgen
