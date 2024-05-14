package openstack

import (
	"os"
	"testing"
)

func TestOSAuth(t *testing.T) {
	userId := os.Getenv("OS_USER_ID")
	password := os.Getenv("OS_PASSWORD")
	tenantId := os.Getenv("OS_TENANT_ID")
	endpoint := os.Getenv("OS_ENDPOINT")
	bucket := os.Getenv("OS_BUCKET")

	_, err := NewClient(userId, password, tenantId, endpoint, bucket)
	if err != nil {
		t.Fatal(err)
	}
}
