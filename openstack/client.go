package openstack

import (
	"fmt"
	"github.com/tappoy/storage/types"
	"net/http"
	"strings"
)

const authUrl = "https://identity.c3j1.conoha.io/v3/auth/tokens"

// OpenstackClient is a client for Openstack Object Storage.
type OpenstackClient struct {
	token    string
	tenantId string
	endpoint string
	bucket   string
}

const authFormat = `{
	"auth": {
		"identity": {
			"methods": ["password"],
			"password": {
				"user": {
					"id": "%s",
					"password": "%s"
				}
			}
		},
		"scope": {
			"project": {
				"id": "%s"
			}
		}
	}
}`

// String returns the client information.
// This is used to logging or debugging.
func (c OpenstackClient) String() string {
	return fmt.Sprintf("OpenstackClient{bucket: %s}", c.bucket)
}

// NewClient is a factory method for OpenstackClient.
//
// Errors:
//   - http.NewRequest
//   - http.DefaultClient.Do
//   - "status code: %d" if response status code is not 201
func NewClient(userId, password, tenantId, endpoint, bucket string) (types.Client, error) {
	body := fmt.Sprintf(authFormat, userId, password, tenantId)
	req, err := http.NewRequest(http.MethodPost, authUrl, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	token := resp.Header.Get("X-Subject-Token")
	return OpenstackClient{token: token, tenantId: tenantId, endpoint: endpoint, bucket: bucket}, nil
}

// NewClientFromConfig is a factory method for OpenstackClient.
func NewClientFromConfig(config map[string]string) (types.Client, error) {
	// check required fields
	if _, ok := config["OS_USER_ID"]; !ok {
		return nil, fmt.Errorf("missing OS_USER_ID")
	}

	if _, ok := config["OS_PASSWORD"]; !ok {
		return nil, fmt.Errorf("missing OS_PASSWORD")
	}

	if _, ok := config["OS_TENANT_ID"]; !ok {
		return nil, fmt.Errorf("missing OS_TENANT_ID")
	}

	if _, ok := config["OS_ENDPOINT"]; !ok {
		return nil, fmt.Errorf("missing OS_ENDPOINT")
	}

	if _, ok := config["OS_BUCKET"]; !ok {
		return nil, fmt.Errorf("missing OS_BUCKET")
	}

	return NewClient(config["OS_USER_ID"], config["OS_PASSWORD"], config["OS_TENANT_ID"], config["OS_ENDPOINT"], config["OS_BUCKET"])
}
