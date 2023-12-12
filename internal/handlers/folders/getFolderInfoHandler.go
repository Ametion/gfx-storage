package folders

import (
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

func GetFolderInfoHandler(c *gin.Context) {
	bucket := c.Query("bucket")
	folder := c.Query("folder")

	if bucket == "" {
		c.JSON(400, response.Response{
			Message: "Bucket name is required",
			Code:    400,
		})
		return
	}

	bucket = filepath.Clean(bucket)
	folder = filepath.Clean(folder)

	if filepath.IsAbs(bucket) || filepath.IsAbs(folder) || strings.Contains(bucket, "..") || strings.Contains(folder, "..") {
		c.JSON(400, response.Response{
			Message: "Invalid bucket or folder name",
			Code:    400,
		})
		return
	}

	fullPath := filepath.Join(helpers.BasePath, bucket, folder)

	if !strings.HasPrefix(fullPath, filepath.Clean(helpers.BasePath)) {
		c.JSON(400, response.Response{
			Message: "Invalid bucket or folder name",
			Code:    400,
		})
		return
	}

	info, err := os.Stat(fullPath)

	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(404, response.Response{
				Message: "Folder not found",
				Code:    404,
			})
		} else {
			c.JSON(500, response.Response{
				Message: "Unable to read folder info",
				Code:    500,
			})
		}
		return
	}

	if !info.IsDir() {
		c.JSON(400, response.Response{
			Message: "Invalid folder name",
			Code:    400,
		})
		return
	}

	entries, err := os.ReadDir(fullPath)

	if err != nil {
		c.JSON(500, response.Response{
			Message: "Unable to read folder info",
			Code:    500,
		})
		return
	}

	var contents []string

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			name += "/"
		}
		contents = append(contents, name)
	}

	c.JSON(200, contents)
}
