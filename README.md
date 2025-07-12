# GopherLogs – Go CLI for Tailing & Filtering Logs
A practical Go CLI tool to tail, filter, and view log files efficiently.

🚀 Features
- Tail log files live with filters
- View last N lines of a log file
- Filter files for keywords
- Handle multiple log files concurrently
- Clean CLI experience with Cobra

## Usage
1. Tail live logs
```bash
./gopherlogs tail --file /path/to/logfile --keyword "ERROR" --lines 50
```

2. View last N lines: 
```bash
./gopherlogs last --file /path/to/logfile --lines 100
```

3. Filter file for a keyword

```bash
./gopherlogs filter --file /path/to/logfile --keyword "WARN"
```

4. Tail multiple files
```bash
./gopherlogs tail --file /var/log/syslog --file /var/log/auth.log --keyword "ssh"
```

## 🧩 Commands Overview
- `tail` : Tails log files live.
Options:
  - --file (multiple)   
  - --keyword (optional)
  - --lines (optional)

- `last` : Displays last N lines from a file.
Options:
  - --file (required)
  - --lines (required)

- `filter` : Filters an entire file for a keyword.
Options:
  - --file (required)
  - --keyword (required)

🚦 Next Step
1️⃣ Set up Cobra (cobra-cli init --pkg github.com/yourusername/gopherlogs).
2️⃣ Create subcommands: tail, last, filter.
3️⃣ Implement clean file handling, keyword filtering, and concurrency.
4️⃣ Add README and sample logs for testing.
5️⃣ Push to GitHub when done.

