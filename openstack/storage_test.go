package openstack

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestOSNormal(t *testing.T) {
	userId := os.Getenv("OS_USER_ID")
	password := os.Getenv("OS_PASSWORD")
	tenantId := os.Getenv("OS_TENANT_ID")
	endpoint := os.Getenv("OS_ENDPOINT")
	bucket := os.Getenv("OS_BUCKET")

	// check env
	if userId == "" || password == "" || tenantId == "" || endpoint == "" || bucket == "" {
		t.Skip("OS_USER_ID, OS_PASSWORD, OS_TENANT_ID, OS_ENDPOINT, OS_BUCKET are required")
	}

	// NewClient
	c, err := NewClient(userId, password, tenantId, endpoint, bucket)
	if err != nil {
		t.Fatal(err)
	}

	// PutObject
	object := "test-object"
	err = c.Put(object, strings.NewReader("test"))
	if err != nil {
		t.Error(err)
	}

	// PutObject
	object = "test-object2"
	err = c.Put(object, strings.NewReader("test"))
	if err != nil {
		t.Error(err)
	}

	// List
	objects, err := c.List("")
	if err != nil {
		t.Error(err)
	}
	t.Log(objects)

	// Head
	info, err := c.Head("test-object")
	if err != nil {
		t.Error(err)
	}
	t.Log(info)

	// Get
	info, reader, err := c.Get("test-object")
	if err != nil {
		t.Error(err)
	}
	t.Log(info)

	// Show content
	content, err := io.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(content))

	// Delete
	err = c.Delete(object)
	if err != nil {
		t.Error(err)
	}

	// Delete
	err = c.Delete(object)
	if err != nil {
		t.Error(err)
	}
}
