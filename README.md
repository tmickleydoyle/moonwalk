[![Coverage Status](https://coveralls.io/repos/github/tmickleydoyle/moonwalk/badge.svg?branch=master)](https://coveralls.io/github/tmickleydoyle/moonwalk?branch=master)

# MoonWalk: Walk to the Root Directory

MoonWalk travels the file tree from a child directory back to the root directory calling walkFn for each file or direcory in the path back. The walkFn function will handle the errors from the visited files and directories.

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
