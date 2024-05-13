package aws

import (
	"os"
	//"strings"
	"testing"
)

func TestNormal(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("AWS_ENDPOINT")

	// NewClient
	c, err := NewClient(region, accessKey, secretKey, endpoint)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := c.GetContainer("tappoydev", "")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
