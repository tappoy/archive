package conoha

import (
	"fmt"
	"github.com/tappoy/cloud"
	"net/http"
	"strings"
)

const authUrl = "https://identity.c3j1.conoha.io/v3/auth/tokens"

// ConohaClient is a client for ConoHa API.
type ConohaClient struct {
	Token    string
	TenantId string
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

// NewClient is a factory method for ConohaClient.
//
// Errors:
//   - http.NewRequest
//   - http.DefaultClient.Do
//   - "status code: %d" if response status code is not 201
func NewClient(userId, password, tenantId string) (cloud.Client, error) {
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
	return ConohaClient{Token: token, TenantId: tenantId}, nil
}
