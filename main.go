package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	slide "github.com/tmickleydoyle/moonwalk/slide"
)

var (
	dir  string
	path string
)

func main() {
	flag.StringVar(&dir, "path", "", "starting point")
	flag.Parse()

	if dir != "" {
		path = dir
	} else {
		path, _ = os.Getwd()
	}

	err := slide.Slide(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		fmt.Printf("%q\n", path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}
