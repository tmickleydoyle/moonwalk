package moonwalk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMoonWalkNoEndSep(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testDir")

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

		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
