package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

/**
*    上传本地文件到S3
*    params :
*        key S3上保存文件的路径（名字）
*        filename 本地文件名字
 */
func (sc *S3Client) UploadFile(key, filename string) (string, error) {
	creds := credentials.NewStaticCredentials(sc.S3AccessKey, sc.S3SecretKey, "")
	config := &aws.Config{
		Region:           aws.String(sc.S3Region),
		Endpoint:         aws.String(sc.S3EndPoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
	}
	sess := session.Must(session.NewSession(config))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("failed to open file %q, %v", filename, err)
		return "", err
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(sc.S3Bucket),
		Key:    aws.String(key),
		Body:   f,
	})
	if err != nil {
		log.Printf("failed to upload file, %v", err)
		return "", err
	}
	path := result.Location
	return path, nil
}

/**
*    从S3下载文件到本地并保存
*    params :
*        key S3上需要下载的文件的路径（名字）
*        filename 本地文件名字
 */
func (sc *S3Client) DownloadFile(key, filename string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(sc.S3Region),
		Endpoint:         aws.String(sc.S3EndPoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(sc.S3AccessKey, sc.S3SecretKey, ""),
	}))

	log.Printf("sc.S3EndPoint= %v", sc.S3EndPoint)
	// Create an uploader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("failed to create file %q, %v", filename, err)
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	// Write the contents of S3 Object to the file
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(sc.S3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Printf("failed to download file, %v", err)
		return err
	}

	return nil
}
