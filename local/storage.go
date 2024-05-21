package local

import (
	"github.com/tappoy/crypto"
	"github.com/tappoy/storage/v2/types"

	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func makeObject(path, name string) (types.Object, bool) {
	stat, err := os.Stat(path)
	if err != nil {
		return types.Object{}, false
	}

	body, err := os.ReadFile(path)
	if err != nil {
		return types.Object{}, false
	}

	return types.Object{
		Name:         name,
		Bytes:        int64(stat.Size()),
		Hash:         crypto.Md5(string(body)),
		LastModified: stat.ModTime(),
	}, true
}

// List retrieves a object list in the container.
func (c LocalClient) List(prefix string) ([]types.Object, error) {
	lists := []types.Object{}

	err := filepath.Walk(c.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		obj, ok := makeObject(path, strings.TrimPrefix(path, c.dir+"/"))
		if !ok {
			return fmt.Errorf("failed to make object")
		}

		lists = append(lists, obj)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return lists, nil
}

// Put creates an object.
func (c LocalClient) Put(object string, body io.Reader) error {
	path := filepath.Join(c.dir, object)

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	// dir
	dir := filepath.Dir(path)

	// mkdir
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// write
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Get retrieves an object.
//
// Errors:
//   - ErrNotFound: if the object is not found.
func (c LocalClient) Get(object string) (types.Object, io.Reader, error) {
	// path
	path := filepath.Join(c.dir, object)

	// stat
	_, err := os.Stat(path)
	if err != nil {
		return types.Object{}, nil, types.ErrNotFound
	}

	// make Object
	obj, ok := makeObject(path, object)
	if !ok {
		return types.Object{}, nil, fmt.Errorf("failed to make object")
	}

	// read all
	b, err := os.ReadFile(path)
	if err != nil {
		return types.Object{}, nil, err
	}

	return obj, bytes.NewReader(b), nil
}

// Delete deletes an object.
//
// Errors:
//   - ErrNotFound: if the object is not found.
func (c LocalClient) Delete(object string) error {
	// path
	path := filepath.Join(c.dir, object)

	// stat
	_, err := os.Stat(path)
	if err != nil {
		return types.ErrNotFound
	}

	// remove
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// Head retrieves an object metadata.
//
// Errors:
//   - ErrNotFound: if the object is not found.
func (c LocalClient) Head(object string) (types.Object, error) {
	// path
	path := filepath.Join(c.dir, object)

	// stat
	_, err := os.Stat(path)
	if err != nil {
		return types.Object{}, types.ErrNotFound
	}

	obj, ok := makeObject(path, object)
	if !ok {
		return types.Object{}, fmt.Errorf("failed to make object")
	}

	return obj, nil
}
