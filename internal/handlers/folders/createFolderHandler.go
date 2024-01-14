package folders

import (
	"gfx-storage/internal/models/request"
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

func CreateFolderHandler(c *gin.Context) {
	var body request.CreateFolderBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, response.Response{
			Message: "Invalid request body",
			Code:    400,
		})
		return
	}

	body.Folder = filepath.Clean(body.Folder)

	if filepath.IsAbs(body.Folder) || strings.Contains(body.Folder, "..") {
		c.JSON(400, response.Response{
			Message: "Invalid folder path",
			Code:    400,
		})
		return
	}

	newFolderPath := filepath.Join(helpers.BasePath, body.Bucket, body.Folder)

	if _, err := os.Stat(newFolderPath); err == nil {
		c.JSON(400, response.Response{
			Message: "Folder already exists",
			Code:    400,
		})
		return
	} else if !os.IsNotExist(err) {
		c.JSON(500, response.Response{
			Message: "Error checking folder existence",
			Code:    500,
		})
		return
	}

	err := os.MkdirAll(newFolderPath, os.ModePerm)
	if err != nil {
		c.JSON(500, response.Response{
			Message: "Could not create folder",
			Code:    500,
		})
		return
	}

	c.JSON(200, response.Response{
		Message: "Folder created successfully at " + newFolderPath,
		Code:    200,
	})
}
