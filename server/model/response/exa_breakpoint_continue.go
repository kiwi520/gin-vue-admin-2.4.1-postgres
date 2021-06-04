package response

import (
	"gin-vue-admin/model/postgres"
)

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File postgres.ExaFile `json:"file"`
}
