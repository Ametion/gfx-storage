package buckets

import (
	"gfx-storage/pkg/helpers"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func GetBucketHandler(context *gin.Context) {
	bucket := context.Param("bucket")

	files, err := os.ReadDir(filepath.Join(helpers.BasePath, bucket))

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var folders []string

	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}

	context.JSON(200, folders)
}
