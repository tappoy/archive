package openstack

import (
	"encoding/json"
	"fmt"
	"github.com/tappoy/archive/types"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (c OpenstackClient) osUrl() string {
	return c.endpoint + c.tenantId
}

// List retrieves a objject list of a container.
//
// Reference:
//   - https://doc.conoha.jp/api-vps3/object-get_objects_list-v3/
func (c OpenstackClient) List(prefix string) ([]types.Object, error) {
	apiUrl := c.osUrl() + "/" + c.bucket
	if len(prefix) > 0 {
		apiUrl += "?prefix=" + prefix
	}
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var objects []types.Object
	err = json.Unmarshal(body, &objects)
	if err != nil {
		return nil, err
	}

	return objects, nil
}

// Put uploads an object.
//
// References:
//   - https://doc.conoha.jp/api-vps3/object-upload_object-v3/
func (c OpenstackClient) Put(object string, r io.Reader) error {
	apiUrl := c.osUrl() + "/" + c.bucket + "/" + object
	req, err := http.NewRequest(http.MethodPut, apiUrl, r)
	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return nil
}

// Delete deletes an object.
//
// References:
//   - https://doc.conoha.jp/api-vps3/object-delete_object-v3/
func (c OpenstackClient) Delete(object string) error {
	apiUrl := c.osUrl() + "/" + c.bucket + "/" + object
	req, err := http.NewRequest(http.MethodDelete, apiUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 && resp.StatusCode != 404 {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return nil
}

// Head retrieves an object metadata.
//
// Errors:
//   - ErrNotFound: if the object is not found.
//
// References:
//   - https://doc.conoha.jp/api-vps3/object-get_objects_detail_specified-v3/
//     2024-05-15: It's wrong. It says 'GET', but it's actually 'HEAD'.
//   - https://docs.openstack.org/api-ref/object-store/#show-object-metadata
func (c OpenstackClient) Head(object string) (types.Object, error) {
	apiUrl := c.osUrl() + "/" + c.bucket + "/" + object
	req, err := http.NewRequest(http.MethodHead, apiUrl, nil)
	if err != nil {
		return types.Object{}, err
	}

	req.Header.Set("X-Auth-Token", c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return types.Object{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return types.Object{}, types.ErrNotFound
		}
		return types.Object{}, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	bytes, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return types.Object{}, err
	}

	lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))

	return types.Object{
		Name:         object,
		Hash:         resp.Header.Get("Etag"),
		Bytes:        bytes,
		LastModified: lastModified,
	}, nil
}

// Get retrieves an object.
//
// Errors:
//   - ErrNotFound: if the object is not found.
//
// References:
//   - https://doc.conoha.jp/api-vps3/object-download_object-v3/
func (c OpenstackClient) Get(object string) (types.Object, io.Reader, error) {
	apiUrl := c.osUrl() + "/" + c.bucket + "/" + object
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return types.Object{}, nil, err
	}

	req.Header.Set("X-Auth-Token", c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return types.Object{}, nil, err
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return types.Object{}, nil, types.ErrNotFound
		}
		return types.Object{}, nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	bytes, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return types.Object{}, nil, err
	}

	lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
	if err != nil {
		return types.Object{}, nil, err
	}

	return types.Object{
		Name:         object,
		Hash:         resp.Header.Get("Etag"),
		Bytes:        bytes,
		LastModified: lastModified,
	}, resp.Body, nil
}
