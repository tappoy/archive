// This package provides an interface for archiving to cloud services.
//
// References:
//   - ConoHa     https://doc.conoha.jp/api-vps3/
//   - OpenStack  https://docs.openstack.org/2024.1/api/
//   - Sakura     https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html
//   - Google     https://cloud.google.com/storage/docs/json_api/v1
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
	// List retrieves a object list in the container.
	List(prefix string) ([]Object, error)

	// Put creates an object.
	Put(object string, body io.Reader) error

	// Delete deletes an object.
	Delete(object string) error

	// Head retrieves an object metadata.
	Head(object string) (Object, error)
}
