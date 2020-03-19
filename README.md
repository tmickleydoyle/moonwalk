[![Latest Release](https://img.shields.io/github/release/tmickleydoyle/moonwalk.svg)](https://github.com/tmickleydoyle/moonwalk/releases)
[![Build Status](https://github.com/tmickleydoyle/moonwalk/workflows/build/badge.svg)](https://github.com/tmickleydoyle/moonwalk/actions)
[![Coverage Status](https://coveralls.io/repos/github/tmickleydoyle/moonwalk/badge.svg?branch=master)](https://coveralls.io/github/tmickleydoyle/moonwalk?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/charmbracelet/glow)](http://goreportcard.com/report/tmickleydoyle/moonwalk)

# MoonWalk: Walk to the Root Directory

MoonWalk travels the file tree from a child directory back to the root directory calling walkFn for each file in the path back. The walkFn function will handle the errors from the visited files and return the name of the file with the path to the file. MoonWalk does not return paths with only directory names.

Example:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tmickleydoyle/moonwalk"
)

func main() {
	err := moonwalk.Slide("/Users/tmickleydoyle/Desktop", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

```

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

[Go Playground Example](https://play.golang.org/p/iwX6I3cMc3k)
