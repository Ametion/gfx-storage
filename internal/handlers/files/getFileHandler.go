package files

import (
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func GetFileHandler(context *gin.Context) {
	bucket := context.Query("bucket")
	path := context.Query("path")
	fileName := context.Query("fileName")

	fullPath := filepath.Join(helpers.BasePath, bucket, path)

	if fileName == "" {
		context.JSON(400, response.Response{
			Message: "Filename is required",
			Code:    400,
		})
		return
	}

	filePath, err := helpers.GetSingleFile(fullPath, fileName)

	if err != nil {
		context.JSON(500, response.Response{
			Message: "Could not get file: " + err.Error(),
			Code:    500,
		})
		return
	}

	context.File(filePath)
}
