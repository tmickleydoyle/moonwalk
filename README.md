# MoonWalk: Walk to the root directory

`moonwalk` recursively walks the working directory back to the root returning all files in the walk back. It now includes filtering, search, multiple output formats, and more detailed statistics about files found in the walk.

## Features

- Walk from current directory or a specified path back to root
- Filter files by extension
- Search for text within files
- Multiple output formats (text, JSON, CSV)
- File size reporting
- Statistics and summary information
- Control depth of directory traversal

## Examples

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

### Combine multiple options

```bash
$ moonwalk -path /Users/tmickleydoyle/Projects -ext .go -search "func" -size -format json
```

## Example Output (Text Format)

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

## Example Output (Summary)

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

## Installation

Run the following command in the terminal.

```bash
$ go get github.com/tmickleydoyle/moonwalk
```

## Usage

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
```
