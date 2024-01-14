package files

import (
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func DeleteFileHandler(context *gin.Context) {
	bucket := context.Query("bucket")
	path := context.Query("path")
	filename := context.Query("filename")

	fullPath := filepath.Join(helpers.BasePath, bucket, path)

	if err := helpers.DeleteFile(fullPath, filename); err != nil {
		context.JSON(400, response.Response{
			Message: "Error while try to delete file",
			Code:    400,
		})
		return
	}

	context.JSON(200, response.Response{
		Message: "Successfully Deleted File by Path: " + fullPath + "/" + filename,
		Code:    200,
	})
}
