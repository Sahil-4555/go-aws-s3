package configs

import (
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

const (
	META_SUCCESS = 1
	META_FAILED  = 0
)

const (
	STATUS_CREATED               = 201
	STATUS_DUPLICATE             = 409
	STATUS_OK                    = 200
	STATUS_FOUND                 = 302
	STATUS_BAD_REQUEST           = 400
	STATUS_UNAUTHORIZED          = 401
	STATUS_INTERNAL_SERVER_ERROR = 500
)

type UploadMediaResponse struct {
	FileName     string `bson:"file_name" json:"file_name" structs:"file_name"`
	PreSignedUrl string `bson:"pre_signed_url" json:"pre_signed_url" structs:"pre_signed_url"`
}

type Meta struct {
	Message string `json:"message" structs:"message"`
	Code    int    `json:"code" structs:"code"`
}

type BaseSuccessResponse struct {
	Data interface{} `json:"data" bson:"data"  structs:"data"`
	Meta Meta        `json:"meta" bson:"meta"  structs:"meta"`
}

func FinalResponse(data map[string]interface{}) map[string]interface{} {
	response := BaseSuccessResponse{
		Meta: Meta{
			Message: data["message"].(string),
			Code:    data["code"].(int),
		},
	}

	if rData, ok := data["data"]; ok {
		response.Data = rData

	} else {
		response.Data = nil
	}

	m := structs.Map(response)
	return m
}

func GetHTTPStatusCode(resCode interface{}) int {
	if resCode != nil {
		return resCode.(int)
	}
	return http.StatusOK
}

func Respond(c *gin.Context, status int, data map[string]interface{}) {
	d := FinalResponse(data)
	c.JSON(status, d)
}
