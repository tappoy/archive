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

func makeObject(name string, mockObject mockObject) types.Object {
	return types.Object{
		Name:         name,
		Bytes:        int64(len(mockObject.body)),
		Hash:         crypto.Md5(string(mockObject.body)),
		LastModified: mockObject.lastModified,
	}
}

// List retrieves a object list in the container.
func (c MockClient) List(prefix string) ([]types.Object, error) {
	fmt.Printf("%v: List\n", c)
	fmt.Printf("prefix: %v\n", prefix)

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

	// show result
	for _, r := range ret {
		fmt.Printf("object: %v\n", r)
	}

	return ret, nil
}

// Put creates an object.
func (c MockClient) Put(object string, body io.Reader) error {
	fmt.Printf("%v: Put\n", c)
	fmt.Printf("object: %v\n", object)

	b, err := io.ReadAll(body)
	if err != nil {
		fmt.Printf("Put error: %v\n", err)
		return err
	}

	c.bucket[object] = mockObject{
		body:         b,
		lastModified: time.Now(),
	}

	// sleep
	time.Sleep(c.delay)

	// show result
	fmt.Printf("body: %x\n", b)

	return nil
}

// Get retrieves an object.
//
// Errors:
//   - ErrNotFound: if the object is not found.
func (c MockClient) Get(object string) (types.Object, io.Reader, error) {
	fmt.Printf("%v: Get\n", c)
	fmt.Printf("object: %v\n", object)

	if _, ok := c.bucket[object]; !ok {
		fmt.Printf("Get error: %v\n", types.ErrNotFound)
		return types.Object{}, nil, types.ErrNotFound
	}

	// make Object
	obj := makeObject(object, c.bucket[object])

	// sleep
	time.Sleep(c.delay)

	// show result
	fmt.Printf("body: %x\n", c.bucket[object].body)

	return obj, bytes.NewReader(c.bucket[object].body), nil
}

// Delete deletes an object.
//
// Errors:
//   - ErrNotFound: if the object is not found.
func (c MockClient) Delete(object string) error {
	fmt.Printf("%v: Delete\n", c)
	fmt.Printf("object: %v\n", object)

	if _, ok := c.bucket[object]; !ok {
		fmt.Printf("Delete error: %v\n", types.ErrNotFound)
		return types.ErrNotFound
	}

	delete(c.bucket, object)

	// sleep
	time.Sleep(c.delay)

	// show result
	fmt.Printf("Deleted\n")

	return nil
}

// Head retrieves an object metadata.
//
// Errors:
//   - ErrNotFound: if the object is not found.
func (c MockClient) Head(object string) (types.Object, error) {
	fmt.Printf("%v: Head\n", c)
	fmt.Printf("object: %v\n", object)

	if _, ok := c.bucket[object]; !ok {
		fmt.Printf("Head error: %v\n", types.ErrNotFound)
		return types.Object{}, types.ErrNotFound
	}

	// sleep
	time.Sleep(c.delay)

	ret := makeObject(object, c.bucket[object])

	// show result
	fmt.Printf("head: %v\n", ret)

	return ret, nil
}
