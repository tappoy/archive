package conoha

import (
	"encoding/json"
	"fmt"
	"github.com/tappoy/cloud"
	"io"
	"net/http"
	"strconv"
	"time"
	//"strings"
)

const objectStorageUrl = "https://object-storage.c3j1.conoha.io/v1/AUTH_"

func (c ConohaClient) osUrl() string {
	return objectStorageUrl + c.TenantId
}

// CreateContainer creates a container.
//
// Reference:
//   - https://doc.conoha.jp/api-vps3/object-create_container-v3/
func (c ConohaClient) PutContainer(container string) error {
	apiUrl := c.osUrl() + "/" + container
	req, err := http.NewRequest(http.MethodPut, apiUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 && resp.StatusCode != 202 {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return nil
}

// GetContainer retrieves a objject list of a container.
//
// Reference:
//   - https://doc.conoha.jp/api-vps3/object-get_objects_list-v3/
func (c ConohaClient) GetContainer(container, query string) ([]cloud.Object, error) {
	apiUrl := c.osUrl() + "/" + container + "?" + query
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", c.Token)

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

	var objects []cloud.Object
	err = json.Unmarshal(body, &objects)
	if err != nil {
		return nil, err
	}

	return objects, nil
}

// PutObject uploads an object.
//
// References:
//   - https://doc.conoha.jp/api-vps3/object-upload_object-v3/
func (c ConohaClient) PutObject(container, object string, r io.Reader) error {
	apiUrl := c.osUrl() + "/" + container + "/" + object
	req, err := http.NewRequest(http.MethodPut, apiUrl, r)
	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", c.Token)

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

// DeleteObject deletes an object.
//
// References:
//   - https://doc.conoha.jp/api-vps3/object-delete_object-v3/
func (c ConohaClient) DeleteObject(container, object string) error {
	apiUrl := c.osUrl() + "/" + container + "/" + object
	req, err := http.NewRequest(http.MethodDelete, apiUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.Token)

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

// HeadObject retrieves an object metadata.
//
// References:
//   - (WRONG) https://doc.conoha.jp/api-vps3/object-get_objects_detail_specified-v3/
//   - https://docs.openstack.org/api-ref/object-store/#show-object-metadata
func (c ConohaClient) HeadObject(container, object string) (cloud.Object, error) {
	apiUrl := c.osUrl() + "/" + container + "/" + object
	req, err := http.NewRequest(http.MethodHead, apiUrl, nil)
	if err != nil {
		return cloud.Object{}, err
	}

	req.Header.Set("X-Auth-Token", c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return cloud.Object{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return cloud.Object{}, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	bytes, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return cloud.Object{}, err
	}

	lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))

	return cloud.Object{
		Name:         object,
		Hash:         resp.Header.Get("Etag"),
		Bytes:        bytes,
		ContentType:  resp.Header.Get("Content-Type"),
		LastModified: lastModified,
	}, nil
}
