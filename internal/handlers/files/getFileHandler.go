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
	filename := context.Query("filename")

	fullPath := filepath.Join(helpers.BasePath, bucket, path)

	if filename == "" {
		context.JSON(400, response.Response{
			Message: "Filename is required",
			Code:    400,
		})
		return
	}

	filePath, err := helpers.GetSingleFile(fullPath, filename)

	if err != nil {
		context.JSON(500, response.Response{
			Message: "Could not get file: " + err.Error(),
			Code:    500,
		})
		return
	}

	context.File(filePath)
}
