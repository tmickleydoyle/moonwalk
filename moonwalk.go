package moonwalk

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type WalkFunc func(path string, info os.FileInfo, err error) error

var lstat = os.Lstat
var pathSep = string(os.PathSeparator)

func moonWalk(path string, info os.FileInfo, walkFn WalkFunc) error {

	backDir := true

	for backDir == true {

		if path == "." || path == "" {
			backDir = false
		}

		names, err := ioutil.ReadDir(path)

		err1 := walkFn(path, info, err)

		if err != nil || err1 != nil {
			return err1
		}

		if err != nil {
			return err
		}

		for _, name := range names {
			filename := filepath.Join(path, name.Name())
			fileInfo, err := lstat(filename)

			if err != nil {
				if err := walkFn(filename, fileInfo, err); err != nil {
					return err
				}
			} else {
				walkFn(filename, fileInfo, err)
			}
		}

		path = filepath.Dir(path)

		if path == string(os.PathSeparator) {
			path = "."
		}
	}

	return nil
}

func Slide(root string, walkFn WalkFunc) error {
	info, err := os.Lstat(root)
	if err != nil {
		err = walkFn(root, nil, err)
	} else {
		err = moonWalk(root, info, walkFn)
	}

	return err
}
