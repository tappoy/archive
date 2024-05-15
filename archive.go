// This package provides an interface for archiving to cloud services.
//
// Supported protocols:
//   - OpenStack  https://docs.openstack.org/2024.1/api/
//   - AWS        https://docs.aws.amazon.com/AmazonS3/latest/API/Welcome.html
//
// Tested services:
//   - ConoHa     https://doc.conoha.jp/api-vps3/
//   - Sakura     https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html
//
// May be supported:
//   - CloudFlare https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/
//
// Others:
//   - Google     https://cloud.google.com/storage/docs/json_api/v1
package archive

import (
	"errors"
	"github.com/tappoy/archive/openstack"
	"github.com/tappoy/archive/types"
)

var (
	// ErrUnsupportedProtocol is returned when the protocol is not supported.
	ErrUnsupportedProtocol = errors.New("unsupported protocol")
)

func NewClientFromConfig(config map[string]string) (types.Client, error) {
	switch config["ARCHIVE_PROTOCOL"] {
	case "openstack":
		return openstack.NewClientFromConfig(config)
	// case "aws":
	// 	return aws.NewClientFromConfig(config)
	default:
		return nil, ErrUnsupportedProtocol
	}
}
