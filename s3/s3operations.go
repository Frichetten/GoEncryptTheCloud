package s3

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// WriteS3 will upload binary data to S3
func WriteS3(bucketname string, filename string, ciphertext []byte) {
	sess := session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {})

	reader := bytes.NewReader(ciphertext)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("test-bucket"),
		Key:    aws.String(filename),
		Body:   reader,
	})
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}

// DownloadS3 will retrieve a file from S3
func DownloadS3(filename string, bucketname string) []byte {
	sess := session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	downloader := s3manager.NewDownloader(sess, func(u *s3manager.Downloader) {})

	buff := &aws.WriteAtBuffer{}

	_, err := downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String("test-bucket"),
		Key:    aws.String(filename),
	})
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}
