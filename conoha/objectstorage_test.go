package conoha

import (
	"os"
	"strings"
	"testing"
)

func TestNormal(t *testing.T) {
	userId := os.Getenv("OS_USER_ID")
	password := os.Getenv("OS_PASSWORD")
	tenantId := os.Getenv("OS_TENANT_ID")

	// Auth
	c, err := Auth(userId, password, tenantId)
	if err != nil {
		t.Fatal(err)
	}

	// PutContainer
	container := "test-container2"
	err = c.PutContainer(container)
	if err != nil {
		t.Error(err)
	}

	// PutObject
	object := "test-object"
	err = c.PutObject(container, object, strings.NewReader("test"))
	if err != nil {
		t.Error(err)
	}

	// PutObject
	object = "test-object2"
	err = c.PutObject(container, object, strings.NewReader("test"))
	if err != nil {
		t.Error(err)
	}

	// GetContainer
	objects, err := c.GetContainer(container, "")
	if err != nil {
		t.Error(err)
	}
	t.Log(objects)

	// HeadObject
	info, err := c.HeadObject(container, "test-object")
	if err != nil {
		t.Error(err)
	}
	t.Log(info)

	// DeleteObject
	err = c.DeleteObject(container, object)
	if err != nil {
		t.Error(err)
	}

	// DeleteObject
	err = c.DeleteObject(container, object)
	if err != nil {
		t.Error(err)
	}
}
