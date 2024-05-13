// This package provides the way to use Cloud API like ConoHa, Sakura, etc.
//
// References:
//   - ConoHa     https://doc.conoha.jp/api-vps3/
//   - OpenStack  https://docs.openstack.org/2024.1/api/
//   - Google     https://cloud.google.com/storage/docs/json_api/v1
//   - Sakura     https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html
//   - CloudFlare https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/
package archive

import (
	"io"
	"time"
)

type Object struct {
	Name         string    `json:"name"`
	Hash         string    `json:"hash"`
	Bytes        int64     `json:"bytes"`
	ContentType  string    `json:"content_type"`
	LastModified time.Time `json:"last_modified"`
}

type Client interface {
	// PutContainer creates a container.
	PutContainer(container string) error

	// GetContainer retrieves a object list in the container.
	GetContainer(container, query string) ([]Object, error)

	// PutObject creates an object.
	PutObject(container, object string, body io.Reader) error

	// DeleteObject deletes an object.
	DeleteObject(container, object string) error

	// HeadObject retrieves an object metadata.
	HeadObject(container, object string) (Object, error)
}
