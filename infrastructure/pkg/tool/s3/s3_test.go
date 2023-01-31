package s3

import (
	"asm_platform/infrastructure/config"
	"fmt"
	"testing"
)

var s3Client = config.S3Config{
	S3Bucket:    "mss-company-push-test",
	S3AccessKey: "nDmdXAr5jyOOnx5QBYN5",
	S3SecretKey: "X8aT6sAUod80oJzvTwQvfimAbEx3exYBrb0uZXcn",
	S3Region:    "beijing",
	S3EndPoint:  "http://beijing.xstore.qihoo.net",
}

//func TestUpload(t *testing.T) {
//	key := "test1/test1.png"
//	file := "C:\\Users\\lvyazhou\\Downloads\\lv.jpg"
//	client := NewS3Client()
//	client.S3Bucket = S3Bucket
//	client.S3AccessKey = S3AccessKey
//	client.S3SecretKey = S3SecretKey
//	client.S3Region = S3Region
//	client.S3EndPoint = S3EndPoint
//
//	fmt.Println("Test File")
//	path, _ := client.UploadFile(key, file)
//	fmt.Println(path)
//	//newfile := "E:\\test.png"
//	//client.DownloadFile(key, newfile)
//	//
//	//fmt.Println("Test Object")
//	//f, _ := ioutil.ReadFile(file)
//	//newkey := "test2"
//	//path1, _ := client.PutObject(newkey, f)
//	//fmt.Println(path1)
//	//client.GetObject(newkey)
//	//client.DeleteObject(newkey)
//}

//func TestMain(m *testing.M) {
//	setup()
//	code := m.Run()
//	// update
//	key := "mss/exposed_asset/update.txt"
//	file := "E:\\MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\update.txt"
//	fmt.Println("update file txt")
//	client := NewS3Client(config.Conf.S3Config)
//	path, _ := client.UploadFile(key, file)
//	fmt.Println(path)
//
//	// update md5
//	key2 := "mss/exposed_asset/update.txt.md5"
//	file2 := "E:\\MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\update.txt.md5"
//	fmt.Println("update md5 file txt")
//	client2 := NewS3Client(config.Conf.S3Config)
//	path2, _ := client2.UploadFile(key2, file2)
//	fmt.Println(path2)
//
//	// asset
//	key3 := "mss/exposed_asset/c1/asset.txt"
//	file3 := "E:\\MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\asset.txt"
//	fmt.Println("asset file txt")
//	client3 := NewS3Client(config.Conf.S3Config)
//	path3, _ := client3.UploadFile(key3, file3)
//	fmt.Println(path3)
//
//	// asset md5
//	key4 := "mss/exposed_asset/c1/asset.txt.md5"
//	file4 := "E:\\MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\asset.txt.md5"
//	fmt.Println("asset md5 file txt")
//	client4 := NewS3Client(config.Conf.S3Config)
//	path4, _ := client4.UploadFile(key4, file4)
//	fmt.Println(path4)
//	teardown()
//	os.Exit(code)
//}

func TestFaker(t *testing.T) {
	// update
	key := "mss/fake/update.txt"
	file := "E:\\01-MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\update.txt"
	fmt.Println("update file txt")
	client := NewS3Client(&s3Client)
	path, _ := client.UploadFile(key, file)
	fmt.Println(path)

	// update md5
	key2 := "mss/fake/update.txt.md5"
	file2 := "E:\\01-MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\update.txt.md5"
	fmt.Println("update md5 file txt")
	client2 := NewS3Client(&s3Client)
	path2, _ := client2.UploadFile(key2, file2)
	fmt.Println(path2)

	// asset
	key3 := "mss/fake/162644137102174208/asset.txt"
	file3 := "E:\\01-MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\fake\\asset.txt"
	fmt.Println("asset file txt")
	client3 := NewS3Client(&s3Client)
	path3, _ := client3.UploadFile(key3, file3)
	fmt.Println(path3)

	// asset md5
	key4 := "mss/fake/162644137102174208/asset.txt.md5"
	file4 := "E:\\01-MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\fake\\asset.txt.md5"
	fmt.Println("asset md5 file txt")
	client4 := NewS3Client(&s3Client)
	path4, _ := client4.UploadFile(key4, file4)
	fmt.Println(path4)
}

//func TestMain(m *testing.M) {
//	setup()
//	code := m.Run()
//	// update
//	key := "mss/dns/update.txt"
//	file := "E:\\MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\update.txt"
//	fmt.Println("update file txt")
//	client := NewS3Client(config.Conf.S3Config)
//	path, _ := client.UploadFile(key, file)
//	fmt.Println(path)
//
//	// update md5
//	key2 := "mss/dns/update.txt.md5"
//	file2 := "E:\\MSS\\03 系统设计\\数据模型文档\\mss企业订阅服务样例数据\\update.txt.md5"
//	fmt.Println("update md5 file txt")
//	client2 := NewS3Client(config.Conf.S3Config)
//	path2, _ := client2.UploadFile(key2, file2)
//	fmt.Println(path2)
//
//	teardown()
//	os.Exit(code)
//}
