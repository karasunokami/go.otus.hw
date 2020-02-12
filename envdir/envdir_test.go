package envdir

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "test_")
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	assert.Equal(t, nil, err)

	tmp, err := os.Create(dir + "/_env")
	defer func() {
		_ = tmp.Close()
	}()

	assert.Equal(t, nil, err)

	_, err = tmp.Write([]byte("value"))
	assert.Equal(t, nil, err)

	result, err := ReadDir(dir)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "value", result["_env"])
}

func TestReadDirEmpty(t *testing.T) {
	dir, err := ioutil.TempDir("", "test_")
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	assert.Equal(t, nil, err)

	tmp, err := os.Create(dir + "/_env")
	defer func() {
		_ = tmp.Close()
	}()
	assert.Equal(t, nil, err)

	result, err := ReadDir(dir)
	assert.Equal(t, nil, err)
	assert.False(t, len(result) > 0)
}

func TestReadDirEmptyFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "test_")
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	assert.Equal(t, nil, err)

	result, err := ReadDir(dir)
	assert.Equal(t, nil, err)
	assert.False(t, len(result) > 0)
}

func TestReadDirWrong(t *testing.T) {
	_, err := ReadDir("none")
	assert.NotEqual(t, nil, err)
}
