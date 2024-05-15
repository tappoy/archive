package aws

import (
	"context"
	"fmt"
	"github.com/tappoy/archive/types"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// List retrieves a object list in the container.
func (c S3Client) List(prefix string) ([]types.Object, error) {
	// Set the parameters based on the CLI flag inputs.
	params := &s3.ListObjectsV2Input{
		Bucket: &c.bucket,
	}
	if len(prefix) != 0 {
		params.Prefix = &prefix
	}

	// Create the Paginator for the ListObjectsV2 operation.
	p := s3.NewListObjectsV2Paginator(c.client, params, func(o *s3.ListObjectsV2PaginatorOptions) {})

	// Iterate through the S3 object pages.
	var i int
	result := []types.Object{}

	for p.HasMorePages() {
		i++

		// Next Page takes a new context for each page retrieval. This is where
		// you could add timeouts or deadlines.
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("failed to get page %v, %v", i, err)
		}

		// Append the objects found
		for _, obj := range page.Contents {
			o := types.Object{
				Name:         *obj.Key,
				Hash:         *obj.ETag,
				Bytes:        *obj.Size,
				LastModified: *obj.LastModified,
			}
			result = append(result, o)
		}
	}
	return result, nil
}

// Put creates an object.
func (c S3Client) Put(object string, body io.Reader) error {
	params := &s3.PutObjectInput{
		Bucket: &c.bucket,
		Key:    &object,
		Body:   body,
	}
	_, err := c.client.PutObject(context.TODO(), params)
	if err != nil {
		return fmt.Errorf("failed to put object %v, %v", object, err)
	}
	return nil
}

// Delete deletes an object.
func (c S3Client) Delete(object string) error {
	params := &s3.DeleteObjectInput{
		Bucket: &c.bucket,
		Key:    &object,
	}
	_, err := c.client.DeleteObject(context.TODO(), params)
	if err != nil {
		return fmt.Errorf("failed to delete object %v, %v", object, err)
	}
	return nil
}

// Head retrieves an object metadata.
func (c S3Client) Head(object string) (types.Object, error) {
	ret := types.Object{}
	params := &s3.HeadObjectInput{
		Bucket: &c.bucket,
		Key:    &object,
	}
	obj, err := c.client.HeadObject(context.TODO(), params)
	if err != nil {
		return ret, fmt.Errorf("failed to head object %v, %v", object, err)
	}
	ret.Name = object
	ret.Hash = *obj.ETag
	ret.Bytes = *obj.ContentLength
	ret.LastModified = *obj.LastModified
	return ret, nil
}
