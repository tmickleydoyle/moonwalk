package moonwalk

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestMoonWalk(t *testing.T) {
	tmpDir := os.TempDir()

    tmpFile, err := ioutil.TempFile(tmpDir, "prefix-")
    if err != nil {
        t.Errorf("Cannot create temporary file: %s", err)
    }

    defer os.Remove(tmpFile.Name())

    fmt.Println("Created File: " + tmpFile.Name())

    text := []byte("This is a golangcode.com example!")
    if _, err = tmpFile.Write(text); err != nil {
        t.Errorf("Failed to write to temporary file: %s", err)
    }

    if err := tmpFile.Close(); err != nil {
        t.Errorf("Failed to close file: %s", err)
	}
	
	tmpDir, err = ioutil.TempDir(tmpDir, "moonwalktest")

	err = Slide(tmpDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			t.Errorf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Errorf("%s", err)
	}
}
