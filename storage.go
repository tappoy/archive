// This package provides an interface for storage services.
//
// Supported protocols:
//   - OpenStack:  https://docs.openstack.org/2024.1/api/
//   - AWS:        https://docs.aws.amazon.com/AmazonS3/latest/API/Welcome.html
//   - Local:      Local file system
//   - Mock:       In memeory mock
//
// Tested clouds:
//   - ConoHa     https://doc.conoha.jp/api-vps3/
//   - Sakura     https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html
//
// May be supported:
//   - CloudFlare https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/
//
// Others:
//   - Google     https://cloud.google.com/storage/docs/json_api/v1
package storage

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/tappoy/storage/v2/aws"
	"github.com/tappoy/storage/v2/local"
	"github.com/tappoy/storage/v2/mock"
	"github.com/tappoy/storage/v2/openstack"
	"github.com/tappoy/storage/v2/types"
)

var (
	// ErrUnsupportedProtocol is returned when the protocol is not supported.
	ErrUnsupportedProtocol = errors.New("unsupported protocol")

	// ErrMissingProtocol is returned when the protocol is missing.
	ErrMissingProtocol = errors.New("missing protocol")
)

// NewClientFromConfig creates a new client from a configuration map.
//
// Supported protocols:
//   - openstack
//   - aws
//   - local
//   - mock
//
// Errors:
//   - ErrMissingProtocol: when the "ARCHIVE_PROTOCOL" key is missing.
//   - ErrUnsupportedProtocol: when the protocol is not supported.
//   - Any error returned by the client constructor.
func NewClientFromConfig(config map[string]string) (types.Client, error) {
	protocol, ok := config["ARCHIVE_PROTOCOL"]
	if !ok {
		return nil, ErrMissingProtocol
	}
	switch protocol {
	case "openstack":
		return openstack.NewClientFromConfig(config)
	case "aws":
		return aws.NewClientFromConfig(config)
	case "local":
		return local.NewClientFromConfig(config)
	case "mock":
		return mock.NewClientFromConfig(config)
	default:
		return nil, ErrUnsupportedProtocol
	}
}

// ParseError is an error type for parsing configuration.
type ParseError struct {
	// LineNo is the line number where the error occurred.
	LineNo int
	// Line is the line where the error occurred.
	Line string
	// Message is the error message.
	Message string
}

// Error returns the error message.
func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d: %s. %s", e.LineNo, e.Message, e.Line)
}

// NewClientFromString creates a new client from a configuration string.
//
// Errors:
//   - ParseError: when the configuration string is invalid.
//   - ErrUnsupportedProtocol: when the protocol is not supported.
//   - ErrMissingProtocol: when the protocol is missing.
//   - Any error returned by the client constructor.
func NewClientFromString(src string) (types.Client, error) {
	reader := bufio.NewReaderSize(strings.NewReader(src), 4096)

	// read line by line
	config := make(map[string]string)

	lineNo := 0
	for {
		lineNo++
		bytes, _, err := reader.ReadLine()
		// check EOF
		if err != nil {
			break
		}
		line := string(bytes)

		// check empty line
		if len(line) == 0 {
			continue
		}

		// check comment
		if line[0] == '#' {
			continue
		}

		// parse line
		fields := strings.Fields(line)

		// check key-value pair
		if len(fields) != 2 {
			return nil, &ParseError{lineNo, line, "invalid key-value pair"}
		}

		// set key-value pair
		config[fields[0]] = fields[1]
	}

	// create client
	return NewClientFromConfig(config)
}
