// This package is a mock implementation of the client package.
// It is used for testing purposes.
package mock

import (
	"github.com/tappoy/storage/types"

	"fmt"
	"time"
)

type mockObject struct {
	body         []byte
	lastModified time.Time
}

type MockClient struct {
	bucket map[string]mockObject
	delay  time.Duration
}

// String returns the client information.
// This is used to logging or debugging.
func (c MockClient) String() string {
	return fmt.Sprintf("MockClient{delay: %s}", c.delay)
}

// NewClient is a factory method for MockClient.
// delay is the delay in seconds for each operation.
func NewClient(delay time.Duration) types.Client {
	return MockClient{bucket: make(map[string]mockObject), delay: delay}
}

// NewClientFromConfig is a factory method for MockClient.
func NewClientFromConfig(config map[string]string) (types.Client, error) {
	// check required fields
	if _, ok := config["MOCK_DELAY"]; !ok {
		return nil, fmt.Errorf("missing MOCK_DELAY")
	}

	delay, err := time.ParseDuration(config["MOCK_DELAY"])
	if err != nil {
		return nil, err
	}

	return NewClient(delay), nil
}
