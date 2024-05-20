package aws

import (
	"github.com/tappoy/archive/types"

	"io"
	"os"
	"strings"
	"testing"
)

func TestAWSNormal(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("AWS_ENDPOINT")
	bucket := os.Getenv("AWS_BUCKET")

	// check env
	if region == "" || accessKey == "" || secretKey == "" || endpoint == "" || bucket == "" {
		t.Skip("AWS_REGION, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_ENDPOINT, AWS_BUCKET are required")
	}

	// NewClient
	c, err := NewClient(region, accessKey, secretKey, endpoint, bucket)
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
	// AWS should not return error when deleting nonexistent file
	err = c.Delete("test.txt")
	if err != nil {
		t.Error(err)
	}

	// List
	ret, err = c.List("")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)

}

func TestAWSNotFound(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("AWS_ENDPOINT")
	bucket := os.Getenv("AWS_BUCKET")

	// check env
	if region == "" || accessKey == "" || secretKey == "" || endpoint == "" || bucket == "" {
		t.Skip("AWS_REGION, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_ENDPOINT, AWS_BUCKET are required")
	}

	// NewClient
	c, err := NewClient(region, accessKey, secretKey, endpoint, bucket)
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
