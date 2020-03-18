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

	text := []byte("This is a golangcode.com example!")
	if _, err = tmpFile.Write(text); err != nil {
		t.Errorf("Failed to write to temporary file: %s", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Errorf("Failed to close file: %s", err)
	}

	tmpDir, err = ioutil.TempDir(tmpDir, "moonwalktest")

	if err != nil {
		t.Errorf("could not append new directory: %v\n", err)
		return err
	}

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

func TestMoonWalkNoDir(t *testing.T) {
	_, err := ioutil.TempDir("", "tempFolder")

	if err != nil {
		t.Errorf("could not create TempDir: %v\n", err)
		return err
	}

	tmpDir := "."

	err = Slide(tmpDir, func(path string, info os.FileInfo, err error) error {

		fmt.Println(path)

		if err != nil {
			t.Fatal(err)
		}

		if path != "" {
			t.Errorf("expected moonwalk to skip: %s", path)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
