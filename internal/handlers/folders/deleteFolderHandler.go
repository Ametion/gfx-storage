package folders

import (
	"fmt"
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func DeleteFolderHandler(context *gin.Context) {
	bucket := context.Query("bucket")
	path := context.Query("path")

	if path == "" || path == "/" {
		context.JSON(400, response.Response{
			Message: "Can not delete bucket while delete folder",
			Code:    400,
		})
		return
	}

	fullPath := filepath.Join(helpers.BasePath, bucket, path)

	err := helpers.DeleteFolder(fullPath)

	if err != nil {
		context.JSON(500, response.Response{
			Message: "Error while delete directory",
			Code:    500,
		})

		fmt.Println(err.Error())
		return
	}

	context.JSON(200, response.Response{
		Message: "Successfully delete folder by path: " + fullPath,
		Code:    200,
	})
}
