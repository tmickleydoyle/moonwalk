package moonwalk

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// The WalkFunc function is passed inside Slide to get the information about the file
type WalkFunc func(path string, info os.FileInfo, depth int, err error) error

var lstat = os.Lstat

// moonWalk parses the path and moves back towards the root directory.
func moonWalk(path string, info os.FileInfo, walkFn WalkFunc, extension string, maxDepth int) error {
	currDepth := 0

	for path != string(os.PathSeparator) {
		if path == "." || path == "" || (maxDepth >= 0 && currDepth > maxDepth) {
			return nil
		}

		names, err := ioutil.ReadDir(path)

		if err != nil {
			return err
		}

		for _, name := range names {
			filename := filepath.Join(path, name.Name())
			fileInfo, err := lstat(filename)

			// Filter by extension if specified
			if extension != "" && !fileInfo.IsDir() {
				if filepath.Ext(filename) != extension {
					continue
				}
			}

			// Process both files and directories
			walkFn(filename, fileInfo, currDepth, err)
		}

		path = filepath.Dir(path)
		currDepth++
	}

	return nil
}

// Slide is the public func name that manages the returned files
// on the walk back towards the root directory.
func Slide(root string, walkFn WalkFunc, extension string, maxDepth int) error {
	info, err := os.Lstat(root)

	if err != nil {
		err = walkFn(root, nil, 0, err)
	} else {
		err = moonWalk(root, info, walkFn, extension, maxDepth)
	}

	return err
}
