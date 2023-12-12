package folders

import (
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func GetFolderHandler(context *gin.Context) {
	path := context.Query("path")
	bucket := context.Query("bucket")

	fullPath := filepath.Join(helpers.BasePath, bucket, path)

	zipBuffer, err := helpers.GetZipFiles(fullPath)
	if err != nil {
		context.JSON(500, response.Response{
			Message: "Could not get folder: " + err.Error(),
			Code:    500,
		})
		return
	}
	context.Header("Content-Disposition", "attachment; filename="+bucket+".zip")
	context.Data(200, "application/zip", zipBuffer.Bytes())
}
