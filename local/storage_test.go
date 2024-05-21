package local

import (
	"github.com/tappoy/archive/types"

	"io"
	"strings"
	"testing"
)

const testDir = "/tmp/tappoy/archive/local"

func TestLocalNormal(t *testing.T) {
	// NewClient
	c, err := NewClient(testDir)
	if err != nil {
		t.Error(err)
	}

	// List
	ret, err := c.List("")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)

	// Put
	err = c.Put("test.txt", strings.NewReader("test"))
	if err != nil {
		t.Error(err)
	}
	t.Log("Put test.txt")

	// Head
	head, err := c.Head("test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Log(head)

	// List
	ret, err = c.List("")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)

	// Get
	head, reader, err := c.Get("test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Log(head)

	// Show content
	content, err := io.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(content))

	// Put
	err = c.Put("a/b/c/test.txt", strings.NewReader("test"))
	if err != nil {
		t.Error(err)
	}
	t.Log("Put a/b/c/test.txt")

	// // Delete
	// err = c.Delete("test.txt")
	// if err != nil {
	// 	t.Error(err)
	// }
	// t.Log("Delete test.txt")

	// // Delete
	// err = c.Delete("test.txt")
	// if err != types.ErrNotFound {
	// 	t.Error(err)
	// }

	// List
	ret, err = c.List("")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestLocalNotFound(t *testing.T) {
	// NewClient
	c, err := NewClient(testDir)
	if err != nil {
		t.Error(err)
	}

	// Head
	_, err = c.Head("nonexistent.txt")
	if err != types.ErrNotFound {
		t.Error(err)
	}

	// Get
	_, _, err = c.Get("nonexistent.txt")
	if err != types.ErrNotFound {
		t.Error(err)
	}
}
