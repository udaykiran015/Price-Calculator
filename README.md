# Price-Calculator

A small Go utility that reads a `prices.txt` file, converts values, and computes totals.

## 🔧 Prerequisites
- Go 1.20+ installed
- Git for source control

## 🚀 Quick start
Clone and run locally:

```bash
git clone <repo-url>
cd Price-Calculator
go run .
```

Build a binary:

```bash
go build -o price-calculator
./price-calculator
```

## 📄 `prices.txt` format (example)
```
USD 12.50
EUR 9.99
JPY 1500
```

## 🧪 Tests
Run unit tests:

```bash
go test ./...
```

## 📦 Contributing
- Create a branch: `git checkout -b feat/describe-change`
- Stage & commit: `git add . && git commit -m "Add feature"
- Push: `git push -u origin feat/describe-change`

## 💡 Notes
- See `conversion/`, `filemanager/`, and `prices/` packages for implementation details.

---

Happy coding! ✨