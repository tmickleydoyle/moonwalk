# MoonWalk: Walk to the root directory 🌙

`moonwalk` recursively walks the working directory back to the root returning all files in the walk back. It now includes filtering, search, multiple output formats, an **interactive TUI mode**, and detailed statistics about files found in the walk.

![MoonWalk Demo](https://img.shields.io/badge/TUI-Interactive-blue?style=for-the-badge)
![Go Version](https://img.shields.io/badge/Go-1.15+-00ADD8?style=for-the-badge&logo=go)

## ✨ Features

- 🚀 Walk from current directory or a specified path back to root
- 🔍 Filter files by extension
- 📝 Search for text within files
- 📊 Multiple output formats (text, JSON, CSV)
- 🖥️  **Interactive TUI mode with real-time navigation**
- 📏 File size reporting
- 📈 Statistics and summary information
- 🎯 Control depth of directory traversal

## Examples

## 🎮 Interactive TUI Mode

Launch the beautiful terminal user interface:

```bash
$ moonwalk -tui
```

### TUI Features:
- 📁 **Visual file browser** with file-type specific icons (🐹 Go, 🐍 Python, 🟨 JavaScript, etc.)
- ⌨️  **Keyboard navigation** with arrow keys or vim-style (j/k)
- 🔍 **Live search** - press `/` to filter files by name
- 📊 **Toggle statistics** - press `d` to show/hide detailed info
- 📁 **Path display** - press `p` to toggle current directory path
- ❓ **Interactive help** - press `h` or `?` for command reference
- 🔄 **Real-time refresh** - press `r` to rescan directory
- 🎨 **Beautiful styling** with syntax highlighting
- 📱 **Responsive design** that adapts to terminal size

### TUI Keyboard Shortcuts:
| Key | Action |
|-----|--------|
| `↑/↓` or `j/k` | Navigate up/down |
| `g` / `G` | Go to top/bottom |
| `/` | Enter search mode |
| `d` | Toggle detailed statistics |
| `p` | Toggle path display |
| `h` or `?` | Toggle help |
| `r` | Refresh/rescan files |
| `q` or `Ctrl+C` | Quit |

### Combine TUI with other options:

```bash
# TUI mode with Go files only
$ moonwalk -tui -ext .go

# TUI mode with content search
$ moonwalk -tui -search "TODO"

# TUI mode from specific directory
$ moonwalk -tui -path /home/user/projects
```

## 📚 CLI Examples

### Basic usage from the current directory

```bash
$ moonwalk
```

### Moonwalk from a specific directory

```bash
$ moonwalk -path /Users/tmickleydoyle/Desktop
```

_The path needs to be the absolute path to the working directory._

### Filter by file extension

```bash
$ moonwalk -ext .go
```

### Search for text within files

```bash
$ moonwalk -search "import"
```

### Limit directory depth

```bash
$ moonwalk -depth 2
```

### Show file sizes

```bash
$ moonwalk -size
```

### Generate summary statistics

```bash
$ moonwalk -summary
```

### Output in different formats

```bash
$ moonwalk -format json
$ moonwalk -format csv
```

### 🎯 Advanced combinations

```bash
# Search for Go functions with file sizes in JSON format
$ moonwalk -path /Users/tmickleydoyle/Projects -ext .go -search "func" -size -format json

# Get summary of Python files up to 3 levels deep
$ moonwalk -ext .py -depth 3 -summary

# Interactive TUI for JavaScript files with content search
$ moonwalk -tui -ext .js -search "import"
```

## 📋 Example Output (Text Format)

```text
[Directory] "/Users/tmickleydoyle/Desktop"
[File] "/Users/tmickleydoyle/Desktop/data.csv" (2.3 KB)
[File] "/Users/tmickleydoyle/Desktop/data.json" (4.1 KB)
[File] "/Users/tmickleydoyle/Desktop/filename.txt" (512 B)
[File] "/Users/tmickleydoyle/Desktop/download.png" (2.5 MB)
[Directory] "/Users/tmickleydoyle"
[File] "/Users/tmickleydoyle/cohorts.json" (1.2 KB)
[File] "/Users/tmickleydoyle/house_query.sql" (3.4 KB)
[Directory] "/Users"
[File] "/Users/setup.txt" (128 B)
```

## 📊 Example Output (Summary)

```text
Summary:
Total Files: 8
Total Directories: 3
Total Size: 2.5 MB
Average File Size: 320.0 KB
Largest File: /Users/tmickleydoyle/Desktop/download.png (2.5 MB)

File Extensions:
.csv: 1 files
.json: 2 files
.txt: 2 files
.png: 1 files
.sql: 1 files
.html: 1 files
```

## 🚀 Installation

### Using Go Install
```bash
go install github.com/tmickleydoyle/moonwalk@latest
```

### From Source
```bash
git clone https://github.com/tmickleydoyle/moonwalk.git
cd moonwalk
go build -o moonwalk .
```

### Dependencies
- Go 1.15 or higher
- Terminal with color support (for best TUI experience)

## 📖 Usage

```
Usage:
  moonwalk [flags]

Flags:
  -path string       starting point directory
  -ext string        filter files by extension (e.g., .go, .txt)
  -search string     search for text within files
  -format string     output format (text, json, csv) (default "text")
  -size              show file sizes
  -summary           show summary statistics
  -depth int         maximum directory depth (-1 for unlimited) (default -1)
  -tui               run in interactive TUI mode
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## 🙏 Acknowledgments

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI interface
- Styled with [Lip Gloss](https://github.com/charmbracelet/lipgloss) for beautiful terminal output
- Inspired by the need for better directory exploration tools
