package gw_s3_handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// S3Upload is to deal with S3 Upload with appropriate directory and year-month
// directory structure
// TODO: update to deal with multiple files https://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#Uploader.UploadWithIterator
func S3Upload(appName string, file graphql.Upload) (*s3manager.UploadOutput, error) {
	bucket := "gw-everyday"

	id, err := uuid.NewUUID()
	if err != nil {
		// handle error
	}
	newFilename := fmt.Sprintf("%s%s", id, filepath.Ext(file.Filename))

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewSharedCredentials("", "sipp11-account"),
	})

	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	basePath := fmt.Sprintf("/%s/%s/%s", appName, now.Format("2006"), now.Format("01"))

	// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
	// for more information on configuring part size, and concurrency.
	//
	// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
	uploader := s3manager.NewUploader(sess)
	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fmt.Sprintf("%s/%s", basePath, newFilename)),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: bytes.NewReader(content),
	})
	if err != nil {
		// Print the error and exit.
		// exitErrorf("Unable to upload %q to %q, %v", newFilename, bucket, err)
		return nil, err
	}
	return output, nil
}
