// This package is a local storage implementation of the client package.
package local

import (
	"github.com/tappoy/archive/types"

	"fmt"
	"os"
)

// LocalClient is a local storage implementation of the client package.
type LocalClient struct {
	dir string
}

// String returns the client information.
// This is used to logging or debugging.
func (c LocalClient) String() string {
	return fmt.Sprintf("LocalClient{dir: %s}", c.dir)
}

// NewClient is a factory method for LocalClient.
// dir is the directory path to store the files.
func NewClient(dir string) (types.Client, error) {
	// make dir
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	return LocalClient{dir: dir}, nil
}

// NewClientFromConfig is a factory method for MockClient.
func NewClientFromConfig(config map[string]string) (types.Client, error) {
	// check required fields
	if _, ok := config["LOCAL_DIR"]; !ok {
		return nil, fmt.Errorf("missing LOCAL_DIR")
	}

	ret, err := NewClient(config["LOCAL_DIR"])
	if err != nil {
		return nil, err
	}

	return ret, nil
}
