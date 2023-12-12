package buckets

import (
	"gfx-storage/internal/models/request"
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"regexp"
)

func CreateBucketHandler(context *gin.Context) {
	var body request.CreateBucketBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, response.Response{
			Message: "Invalid request body",
			Code:    400,
		})
		return
	}

	validBucketName := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

	if !validBucketName(body.BucketName) {
		context.JSON(400, response.Response{
			Message: "Invalid bucket name",
			Code:    400,
		})
		return
	}

	newBucketPath := filepath.Join(helpers.BasePath, body.BucketName)

	if _, err := os.Stat(newBucketPath); err == nil {
		context.JSON(409, response.Response{
			Message: "Bucket already exists",
			Code:    409,
		})
		return
	} else if !os.IsNotExist(err) {
		context.JSON(500, response.Response{
			Message: "Unable to create bucket",
			Code:    500,
		})
		return
	}

	err := os.MkdirAll(newBucketPath, os.ModePerm)
	if err != nil {
		context.JSON(500, response.Response{
			Message: "Unable to create bucket",
			Code:    500,
		})
		return
	}

	context.JSON(200, response.Response{
		Message: "Bucket created successfully",
		Code:    200,
	})
}
