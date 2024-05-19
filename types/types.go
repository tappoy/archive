package types

import (
	"errors"
	"io"
	"time"
)

var ErrNotFound = errors.New("not found")

type Object struct {
	Name         string    `json:"name"`
	Hash         string    `json:"hash"`
	Bytes        int64     `json:"bytes"`
	LastModified time.Time `json:"last_modified"`
}

type Client interface {
	// String returns the client information.
	// This is used to logging or debugging.
	String() string

	// List retrieves a object list in the container.
	List(prefix string) ([]Object, error)

	// Put creates an object.
	Put(object string, body io.Reader) error

	// Delete deletes an object.
	Delete(object string) error

	// Head retrieves an object metadata.
	Head(object string) (Object, error)

	// Get retrieves an object.
	Get(object string) (Object, io.Reader, error)
}
