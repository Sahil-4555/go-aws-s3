package service

import (
	"errors"
	"fmt"
	"go-aws-s3/configs"
	uploadtoaws "go-aws-s3/configs/aws"
	"time"

	"go-aws-s3/log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	maxFileSize = 10 * 1024 * 1024 // 10 MB in bytes
)

func CreateMediaRequest(filesize int64, filename string, uuId string, basePath string) (map[string]interface{}, error) {
	log.GetLog().Info("INFO : ", "S3 Service Called(createMediaRequest).")
	if filesize <= maxFileSize {
		extension := filename
		extArr := strings.Split(extension, ".")
		fileName := uuId + "." + extArr[1]
		key := basePath + "/" + fileName
		response := map[string]interface{}{
			"key":  key,
			"name": fileName,
		}
		return response, nil
	} else {
		return map[string]interface{}{}, errors.New("Invalid file type or size")
	}
}

func UploadMedia(req configs.UploadMediaData) map[string]interface{} {
	log.GetLog().Info("INFO : ", "S3 Service Called(UploadMedia).")
	var respData configs.UploadMediaResponse

	url, err := uploadtoaws.UploadToS3(req.Key)
	if err != nil {
		log.GetLog().Error("ERROR : ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_OK,
		}
	}

	fileName := strings.Split(req.Key, "/")
	respData.FileName = fileName[len(fileName)-1]
	respData.PreSignedUrl = url

	return map[string]interface{}{
		"message": "File uploaded successfully.",
		"code":    configs.STATUS_OK,
		"data":    respData,
	}
}

func GetListOfObject(basePath string) map[string]interface{} {
	log.GetLog().Info("INFO : ", "S3 Service Called(GetListOfObject).")
	accessKey := configs.AccessKey()
	secret := configs.SecretKey()
	region := configs.Region()
	bucket := configs.Bucket()

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
		Region:      aws.String(region),
	})
	if err != nil {
		log.GetLog().Error("ERROR: ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_BAD_REQUEST,
		}
	}

	svc := s3.New(sess)
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(basePath),
	}

	resp, err := svc.ListObjects(params)
	if err != nil {
		log.GetLog().Error("ERROR: ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_BAD_REQUEST,
		}
	}

	var objects []string
	for _, key := range resp.Contents {
		objects = append(objects, *key.Key)
	}

	return map[string]interface{}{
		"message": "Object listed successfully.",
		"code":    configs.META_SUCCESS,
		"data":    objects,
	}
}

func DeleteObject(path string) map[string]interface{} {

	log.GetLog().Info("INFO : ", "S3 Service Called(GetListOfObject).")
	accessKey := configs.AccessKey()
	secret := configs.SecretKey()
	region := configs.Region()
	bucket := configs.Bucket()

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
		Region:      aws.String(region),
	})
	if err != nil {
		log.GetLog().Error("ERROR: ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_BAD_REQUEST,
		}
	}

	svc := s3.New(sess)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}

	d, err := svc.DeleteObject(input)
	if err != nil {
		log.GetLog().Error("ERROR: ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_BAD_REQUEST,
		}
	}

	return map[string]interface{}{
		"message": "Deleted successfully.",
		"code":    configs.META_SUCCESS,
		"data":    d,
	}
}

func SignedURL(path string, minutes time.Duration) map[string]interface{} {

	log.GetLog().Info("INFO : ", "S3 Service Called(SignedURL).")
	accessKey := configs.AccessKey()
	secret := configs.SecretKey()
	region := configs.Region()
	bucket := configs.Bucket()

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
		Region:      aws.String(region),
	})
	if err != nil {
		log.GetLog().Error("ERROR: ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_BAD_REQUEST,
		}
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})
	urlStr, err := req.Presign(minutes * time.Hour)
	if err != nil {
		log.GetLog().Error("ERROR: ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_BAD_REQUEST,
		}
	}

	return map[string]interface{}{
		"message": "Generated SignedURL successfully.",
		"code":    configs.META_SUCCESS,
		"data":    urlStr,
	}
}

func CreateBucket(bucket string) map[string]interface{} {
	accessKey := configs.AccessKey()
	secret := configs.SecretKey()
	region := configs.Region()

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
		Region:      aws.String(region),
	}))

	svc := s3.New(sess)

	data, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		log.GetLog().Error("ERROR: ", err.Error())
		return map[string]interface{}{
			"message":  err.Error(),
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_BAD_REQUEST,
		}
	}

	return map[string]interface{}{
		"message": fmt.Sprintf("Successfully Created the %s", bucket),
		"code":    configs.META_SUCCESS,
		"data":    data.String(),
	}
}
