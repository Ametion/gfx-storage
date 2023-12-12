package files

import (
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, response.Response{
			Message: "File is required: " + err.Error(),
			Code:    400,
		})
		return
	}

	path := c.DefaultPostForm("path", "")
	bucket := c.DefaultPostForm("bucket", "")
	path = filepath.Clean(path)

	filename := filepath.Clean(file.Filename)

	targetPath := filepath.Join(helpers.BasePath, bucket, path, filepath.Base(filename))

	baseDir := filepath.Clean(helpers.BasePath) + string(os.PathSeparator)
	targetDir := filepath.Clean(targetPath)

	if !strings.HasPrefix(targetDir, baseDir) {
		c.JSON(400, response.Response{
			Message: "Invalid path provided for upload file",
			Code:    400,
		})
		return
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
		c.JSON(500, response.Response{
			Message: "Could not create directory: " + err.Error(),
			Code:    500,
		})
		return
	}

	if err := c.SaveUploadedFile(file, targetPath); err != nil {
		c.JSON(500, response.Response{
			Message: "Could not save file: " + err.Error(),
			Code:    500,
		})
		return
	}

	c.JSON(200, response.Response{
		Message: "File uploaded successfully at: " + targetPath,
		Code:    200,
	})
}
