package mock

import (
	"github.com/tappoy/archive/types"
	"github.com/tappoy/crypto"

	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

func makeObject(name string, mockObject MockObject) types.Object {
	return types.Object{
		Name:         name,
		Bytes:        int64(len(mockObject.body)),
		Hash:         crypto.Md5(string(mockObject.body)),
		LastModified: mockObject.lastModified,
	}
}

// List retrieves a object list in the container.
func (c MockClient) List(prefix string) ([]types.Object, error) {
	keys := []string{}
	for k := range c.bucket {
		// check if the key has the prefix
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
	}

	ret := []types.Object{}
	for _, k := range keys {
		ret = append(ret, makeObject(k, c.bucket[k]))
	}

	// sleep
	time.Sleep(c.delay)

	return ret, nil
}

// Put creates an object.
func (c MockClient) Put(object string, body io.Reader) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	c.bucket[object] = MockObject{
		body:         b,
		lastModified: time.Now(),
	}

	// sleep
	time.Sleep(c.delay)

	return nil
}

// Get retrieves an object.
func (c MockClient) Get(object string) (types.Object, io.Reader, error) {
	if _, ok := c.bucket[object]; !ok {
		return types.Object{}, nil, fmt.Errorf("object not found")
	}

	// make Object
	obj := makeObject(object, c.bucket[object])

	// sleep
	time.Sleep(c.delay)

	return obj, bytes.NewReader(c.bucket[object].body), nil
}

// Delete deletes an object.
func (c MockClient) Delete(object string) error {
	if _, ok := c.bucket[object]; !ok {
		return fmt.Errorf("object not found")
	}

	delete(c.bucket, object)

	// sleep
	time.Sleep(c.delay)

	return nil
}

// Head retrieves an object metadata.
func (c MockClient) Head(object string) (types.Object, error) {
	if _, ok := c.bucket[object]; !ok {
		return types.Object{}, fmt.Errorf("object not found")
	}

	// sleep
	time.Sleep(c.delay)

	return makeObject(object, c.bucket[object]), nil
}