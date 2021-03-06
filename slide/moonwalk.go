package moonwalk

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// The WalkFunc function is passed inside Slide to get the information about the tile
type WalkFunc func(path string, info os.FileInfo, err error) error

var lstat = os.Lstat

// moonWalk parses the path and moves back towards the root directory.
func moonWalk(path string, info os.FileInfo, walkFn WalkFunc) error {
	for path != string(os.PathSeparator) {

		if path == "." || path == "" {
			return nil
		}

		names, err := ioutil.ReadDir(path)

		if err != nil {
			return err
		}

		for _, name := range names {
			filename := filepath.Join(path, name.Name())
			fileInfo, err := lstat(filename)

			if !fileInfo.IsDir() {
				walkFn(filename, fileInfo, err)
			}
		}

		path = filepath.Dir(path)
	}

	return nil
}

// Slide is the public func name that manages the returned files
// on the walk back towards the root directory.
func Slide(root string, walkFn WalkFunc) error {
	info, err := os.Lstat(root)

	if err != nil {
		err = walkFn(root, nil, err)
	} else {
		err = moonWalk(root, info, walkFn)
	}

	return err
}
