package moonwalk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMoonWalk(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "moonwalktest")

	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}

	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	err = Slide(tmpDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			t.Fatal(err)
		}

		if !strings.HasPrefix(filepath.Dir(tmpDir), filepath.Dir(path)) {
			t.Errorf("expected moonwalk path to be in %s: %s", path, tmpDir)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestMoonWalkNoDir(t *testing.T) {
	_, err := ioutil.TempDir("", "")

	tmpDir := "."

	err = Slide(tmpDir, func(path string, info os.FileInfo, err error) error {

		fmt.Println(path)

		if err != nil {
			t.Fatal(err)
		}

		if !strings.HasPrefix(filepath.Dir(tmpDir), filepath.Dir(path)) {
			t.Errorf("expected moonwalk path to be in %s: %s", path, tmpDir)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
