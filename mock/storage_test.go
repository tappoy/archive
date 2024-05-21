package mock

import (
	"github.com/tappoy/storage/v2/types"

	"io"
	"strings"
	"testing"
	"time"
)

func TestMockNormal(t *testing.T) {
	// 1s delay
	delay, err := time.ParseDuration("1ms")
	if err != nil {
		t.Fatal(err)
	}

	// NewClient
	c := NewClient(delay)
	if err != nil {
		t.Fatal(err)
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

	// Delete
	err = c.Delete("test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Log("Delete test.txt")

	// Delete
	err = c.Delete("test.txt")
	if err != types.ErrNotFound {
		t.Error(err)
	}

	// List
	ret, err = c.List("")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}

func TestMockNotFound(t *testing.T) {
	// 1s delay
	delay, err := time.ParseDuration("1ms")
	if err != nil {
		t.Fatal(err)
	}

	// NewClient
	c := NewClient(delay)
	if err != nil {
		t.Fatal(err)
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
