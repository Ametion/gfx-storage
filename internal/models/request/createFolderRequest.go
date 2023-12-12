package request

type CreateFolderBody struct {
	Folder string `json:"folder"`
	Bucket string `json:"bucket"`
}
