package aws

import (
	"context"
	"fmt"
	"github.com/tappoy/archive"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// PutContainer creates a container.
func (c S3Client) PutContainer(container string) error {
	return nil
}

// GetContainer retrieves a object list in the container.
func (c S3Client) GetContainer(container, query string) ([]archive.Object, error) {
	// Set the parameters based on the CLI flag inputs.
	params := &s3.ListObjectsV2Input{
		Bucket: &container,
	}
	if len(query) != 0 {
		params.Prefix = &query
	}

	// Create the Paginator for the ListObjectsV2 operation.
	p := s3.NewListObjectsV2Paginator(c.client, params, func(o *s3.ListObjectsV2PaginatorOptions) {
	})

	// Iterate through the S3 object pages, printing each object returned.
	var i int
	fmt.Println("Objects:")
	result := []archive.Object{}

	for p.HasMorePages() {
		i++

		// Next Page takes a new context for each page retrieval. This is where
		// you could add timeouts or deadlines.
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("failed to get page %v, %v", i, err)
		}

		// Log the objects found
		for _, obj := range page.Contents {
			o := archive.Object{
				Name:         *obj.Key,
				Hash:         *obj.ETag,
				Bytes:        *obj.Size,
				LastModified: *obj.LastModified,
			}
			result = append(result, o)
			fmt.Println("Object:", *obj.Key)
			fmt.Println("ETag:", *obj.ETag)
			fmt.Println("Size:", *obj.Size)
			fmt.Println("LastModified:", *obj.LastModified)
		}
	}
	return result, nil
}

// PutObject creates an object.
func (c S3Client) PutObject(container, object string, body io.Reader) error {
	return nil
}

// DeleteObject deletes an object.
func (c S3Client) DeleteObject(container, object string) error {
	return nil
}

// HeadObject retrieves an object metadata.
func (c S3Client) HeadObject(container, object string) (archive.Object, error) {
	return archive.Object{}, nil
}
