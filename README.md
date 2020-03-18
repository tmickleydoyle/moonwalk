![Build Status](https://github.com/tmickleydoyle/moonwalk/workflows/build/badge.svg)[![Coverage Status](https://coveralls.io/repos/github/tmickleydoyle/moonwalk/badge.svg?branch=master)](https://coveralls.io/github/tmickleydoyle/moonwalk?branch=master)[![Go ReportCard](http://goreportcard.com/badge/charmbracelet/glow)](http://goreportcard.com/report/tmickleydoyle/moonwalk)

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
	err := moonwalk.Slide("/directory/to/walk", func(path string, info os.FileInfo, err error) error {

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

[Go Playground Example](https://play.golang.org/p/Av0CtSmoimR)
