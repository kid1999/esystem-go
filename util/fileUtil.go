/**
@auther: kid1999
@date: 2021/7/15 14:01
@desc: fileUtil
**/

package util

import (
	"github.com/minio/minio-go/v6"
	"io"
	"log"
)

var endpoint  = "kid1999.top:9000"
var accessKeyID  = "root"
var secretAccessKey  = "pass=kid1999"
var useSSL  = false
var bucketName = "images"
var location = "us-east-1"


// 初使化minio client对象。
func  GetMinioClient()  *minio.Client {
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}


func  CreateBucket() {
	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)
}


func UploadFile(objectName,filePath,contentType string) (imgUrl string, err error) {
	n, err := GetMinioClient().FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	imgUrl = endpoint + "/" + bucketName + "/" + objectName
	log.Printf("Successfully uploaded %s of size %d\n", imgUrl, n)
	return
}

func UploadFileByStream(objectName string,fileStream io.Reader) (imgUrl string, err error) {
	ops := minio.PutObjectOptions{}
	n, err := GetMinioClient().PutObject(bucketName, objectName, fileStream, -1,ops)
	if err != nil {
		log.Fatalln(err)
	}

	imgUrl = endpoint + "/" + bucketName + "/" + objectName
	log.Printf("Successfully uploaded %s of size %d\n", imgUrl, n)
	return
}

