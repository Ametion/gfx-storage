package buckets

import (
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
)

func DeleteBucketHandler(context *gin.Context) {
	bucket := context.Param("bucket")

	if strings.Contains(bucket, "/") {
		context.JSON(400, response.Response{
			Message: "You can not delete folder while try to delete bucket",
			Code:    400,
		})
		return
	}

	fullPath := filepath.Join(helpers.BasePath, bucket)

	err := helpers.DeleteFolder(fullPath)

	if err != nil {
		context.JSON(500, response.Response{
			Message: "Eror while delete bucket",
			Code:    500,
		})
		return
	}

	context.JSON(200, response.Response{
		Message: "Bucket deleted successfully",
		Code:    200,
	})
}
