# MoonWalk: Walk to the root directory

`moonwalk` recursively walks the working directory back to the root returning all files in the walk back. `moonwalk` does not return paths with only directory names.

## Examples

### Moonwalk from the current directory

```bash
$ moonwalk
```

### Moonwalk from a specific directory

```bash
$ moonwalk -path /Users/tmickleydoyle/Desktop
```

_The path needs to be the absolute path to working directory._

Example Output

```text
/Users/tmickleydoyle/Desktop/data.csv
/Users/tmickleydoyle/Desktop/data.json
/Users/tmickleydoyle/Desktop/filename.txt
/Users/tmickleydoyle/Desktop/download.png
/Users/tmickleydoyle/Desktop/network.json
/Users/tmickleydoyle/cohorts.json
/Users/tmickleydoyle/house_query.sql
/Users/tmickleydoyle/website.html
/Users/tmickleydoyle/blog.txt
/Users/setup.txt
```

## Installation

Run the following command in the terminal.

```bash
$ go get github.com/tmickleydoyle/moonwalk
```
