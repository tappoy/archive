package conoha

import (
	"os"
	"testing"
)

func TestAuth(t *testing.T) {
	userId := os.Getenv("OS_USER_ID")
	password := os.Getenv("OS_PASSWORD")
	tenantId := os.Getenv("OS_TENANT_ID")

	_, err := Auth(userId, password, tenantId)
	if err != nil {
		t.Fatal(err)
	}
}
