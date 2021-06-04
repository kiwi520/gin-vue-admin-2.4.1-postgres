package response

import (
	"gin-vue-admin/model/postgres"
)

type ExaFileResponse struct {
	File postgres.ExaFileUploadAndDownload `json:"file"`
}
