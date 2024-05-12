package main

import (
	"go-aws-s3/configs"
	"go-aws-s3/configs/middleware"
	"go-aws-s3/controllers"
	"go-aws-s3/log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Init()
	port := configs.Port()
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.GinMiddleware())

	r.POST("/upload-file", controllers.UploadMedia)
	r.GET("/object-list", controllers.GetListOfObject)
	r.DELETE("/delete-object", controllers.DeleteObject)
	r.GET("/get-signed-url", controllers.GenerateSignedUrl)
	r.POST("/create-bucket/:bucket", controllers.CreateBucket)

	r.Run(":" + port)
}
