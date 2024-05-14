package aws

import (
	"os"
	//"strings"
	"testing"
)

func TestAWSNormal(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("AWS_ENDPOINT")
	bucket := os.Getenv("AWS_BUCKET")

	// NewClient
	c, err := NewClient(region, accessKey, secretKey, endpoint, bucket)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := c.List("")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
