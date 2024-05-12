package controllers

import (
	"go-aws-s3/configs"
	"go-aws-s3/log"
	service "go-aws-s3/services"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadMedia(c *gin.Context) {
	log.GetLog().Info("INFO : ", "S3 Controller Called(UploadMedia).")

	var req configs.UploadMediaRequest

	if c.BindJSON(&req) != nil {
		data := map[string]interface{}{
			"message":  "Failed to read the body.",
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_OK,
		}
		statusCode := configs.GetHTTPStatusCode(data["res_code"])
		configs.Respond(c, statusCode, data)
		return
	}

	basePath := "Sahil"
	data, err := service.CreateMediaRequest(req.FileSize, req.FileName, "test", basePath)

	var filedata configs.UploadMediaData
	if err == nil {
		filedata.Key = data["key"].(string)
		filedata.FileName = data["name"].(string)
	} else {
		data := map[string]interface{}{
			"message":  "Please provide video file as per the specification : Allowed size: 10MB & Allowed type:[.mkv, .mp4, .flv, .pdf, .png, .jpg, .jpeg]",
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_OK,
		}
		statusCode := configs.GetHTTPStatusCode(data["res_code"])
		configs.Respond(c, statusCode, data)
		return
	}

	resp := service.UploadMedia(filedata)
	statusCode := configs.GetHTTPStatusCode(resp["res_code"])
	configs.Respond(c, statusCode, resp)
	log.GetLog().Info("INFO : ", "Media Uploded successfully...")

}

func GetListOfObject(c *gin.Context) {
	log.GetLog().Info("INFO : ", "S3 Controller Called(GetListOfObject).")

	basePath := "Sahil/"
	resp := service.GetListOfObject(basePath)
	statusCode := configs.GetHTTPStatusCode(resp["res_code"])
	configs.Respond(c, statusCode, resp)
	log.GetLog().Info("INFO : ", "Object Listed successfully...")
}

func DeleteObject(c *gin.Context) {
	log.GetLog().Info("INFO : ", "S3 Controller Called(DeleteObject).")

	var req configs.DeleteObjectReq
	if c.BindJSON(&req) != nil {
		data := map[string]interface{}{
			"message":  "Failed to read the body.",
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_OK,
		}
		statusCode := configs.GetHTTPStatusCode(data["res_code"])
		configs.Respond(c, statusCode, data)
		return
	}

	basePath := "Sahil/"
	key := basePath + req.FileName
	resp := service.DeleteObject(key)
	statusCode := configs.GetHTTPStatusCode(resp["res_code"])
	configs.Respond(c, statusCode, resp)
	log.GetLog().Info("INFO : ", "Object Deleted successfully...")
}

func GenerateSignedUrl(c *gin.Context) {
	log.GetLog().Info("INFO : ", "S3 Controller Called(DeleteObject).")

	var req configs.DeleteObjectReq
	if c.BindJSON(&req) != nil {
		data := map[string]interface{}{
			"message":  "Failed to read the body.",
			"code":     configs.META_FAILED,
			"res_code": configs.STATUS_OK,
		}
		statusCode := configs.GetHTTPStatusCode(data["res_code"])
		configs.Respond(c, statusCode, data)
		return
	}

	basePath := "Sahil/"
	key := basePath + req.FileName
	resp := service.SignedURL(key, time.Duration(3))
	statusCode := configs.GetHTTPStatusCode(resp["res_code"])
	configs.Respond(c, statusCode, resp)
	log.GetLog().Info("INFO : ", "Object Deleted successfully...")
}

func CreateBucket(c *gin.Context) {
	log.GetLog().Info("INFO : ", "S3 Controller Called(CreateBucket).")

	bucket := c.Param("bucket")

	resp := service.CreateBucket(bucket)
	statusCode := configs.GetHTTPStatusCode(resp["res_code"])
	configs.Respond(c, statusCode, resp)
	log.GetLog().Info("INFO : ", "Successfully created the bucket.")
}
