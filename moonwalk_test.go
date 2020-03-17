package moonwalk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

func TestMoonWalkNoEndSep(t *testing.T) {
	tmpDir, err := prepareTestDirTree("prepare/test/directory/to/walk")
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
