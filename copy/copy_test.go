package copy

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func createFiles() (string, string) {
	test := "123456790"
	b := []byte(test)
	path := "/tmp/test1"
	pathTo := "/tmp/test2"
	_ = ioutil.WriteFile(path, b, 0644)

	return path, pathTo
}

func TestCopyFiles(t *testing.T) {
	path, pathTo := createFiles()

	_ = Copy(path, pathTo, 0, 0)
	b, _ := ioutil.ReadFile(pathTo)
	_ = os.Remove(path)
	_ = os.Remove(pathTo)

	assert.Equal(t, "123456790", string(b))
}

func TestCopyFilesWithOffset(t *testing.T) {
	path, pathTo := createFiles()

	_ = Copy(path, pathTo, 0, 3)
	b, _ := ioutil.ReadFile(pathTo)
	_ = os.Remove(path)
	_ = os.Remove(pathTo)

	assert.Equal(t, "456790", string(b))
}

func TestCopyFilesWithOffsetAndLimit(t *testing.T) {
	path, pathTo := createFiles()

	_ = Copy(path, pathTo, 2, 3)
	b, _ := ioutil.ReadFile(pathTo)
	_ = os.Remove(path)
	_ = os.Remove(pathTo)

	assert.Equal(t, "45", string(b))
}

func TestCopyFilesWithBadLimit(t *testing.T) {
	path, pathTo := createFiles()

	err := Copy(path, pathTo, 22, 3)
	_, _ = ioutil.ReadFile(pathTo)
	_ = os.Remove(path)
	_ = os.Remove(pathTo)

	assert.Equal(t, errors.New("limit cant be more then file size or less then 0"), err)
}

func TestCopyFilesWithBadOffset(t *testing.T) {
	path, pathTo := createFiles()

	err := Copy(path, pathTo, 3, 22)
	_, _ = ioutil.ReadFile(pathTo)
	_ = os.Remove(path)
	_ = os.Remove(pathTo)

	assert.Equal(t, errors.New("offset cant be more then file size or less then 0"), err)
}

func TestCopyFilesWithBadOffsetAndLimit(t *testing.T) {
	path, pathTo := createFiles()

	err := Copy(path, pathTo, 9, 9)
	_, _ = ioutil.ReadFile(pathTo)
	_ = os.Remove(path)
	_ = os.Remove(pathTo)

	assert.Equal(t, errors.New("failed to copy file - offset + limit more then file size"), err)
}

func TestCopyFilesWithWrongPath(t *testing.T) {
	err := Copy("none", "", 9, 9)

	assert.Contains(t, fmt.Sprint(err), "no such file or directory")
}
