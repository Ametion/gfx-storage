package main

import (
	"gfx-storage/internal/handlers/buckets"
	"gfx-storage/internal/handlers/files"
	"gfx-storage/internal/handlers/folders"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.GET("/file", files.GetFileHandler)
	engine.POST("/file", files.UploadFileHandler)
	engine.DELETE("/file", files.DeleteFileHandler)

	engine.GET("/folder", folders.GetFolderHandler)
	engine.GET("/folderInfo", folders.GetFolderInfoHandler)
	engine.POST("/folder", folders.CreateFolderHandler)
	engine.DELETE("/folder", folders.DeleteFolderHandler)

	engine.GET("/bucket/:bucket", buckets.GetBucketHandler)
	engine.GET("/buckets", buckets.GetBucketsHandler)
	engine.POST("/bucket", buckets.CreateBucketHandler)

	runErr := engine.Run(":5783")
	if runErr != nil {
		log.Fatalf("Could not start server: %s\n", runErr.Error())
	}
}
