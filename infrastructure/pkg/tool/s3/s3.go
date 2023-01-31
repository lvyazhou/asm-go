package s3

import "asm_platform/infrastructure/config"

type S3Client struct {
	S3Bucket    string //Bucket名称，从hulk可以获取
	S3AccessKey string //AccessKey名称，从hulk可以获取
	S3SecretKey string //SecretKey名称，从hulk可以获取
	S3Region    string
	S3EndPoint  string //S3的域名，从hulk可以获取
}

var Client *S3Client

// NewS3Client s3 client
func NewS3Client(cfg *config.S3Config) *S3Client {
	client := &S3Client{
		S3Bucket:    cfg.S3Bucket,
		S3AccessKey: cfg.S3AccessKey,
		S3SecretKey: cfg.S3SecretKey,
		S3Region:    cfg.S3Region,
		S3EndPoint:  cfg.S3EndPoint,
	}
	return client
}
