package buckets

import (
	"gfx-storage/internal/models/response"
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"os"
)

func GetBucketsHandler(context *gin.Context) {
	dir, err := os.ReadDir(helpers.BasePath)

	if err != nil {
		context.JSON(500, response.Response{
			Message: "Could not read buckets",
			Code:    500,
		})
		return
	}

	var buckets []string

	for _, file := range dir {
		if file.IsDir() {
			buckets = append(buckets, file.Name())
		}
	}

	context.JSON(200, buckets)
}
