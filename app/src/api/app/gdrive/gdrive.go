package gdrive

import (
	//"strings"
	//"net/http"
	"github.com/gin-gonic/gin"
    "google.golang.org/api/drive/v3"
)

var (
	driveService *drive.Service
)

func Configure(r *gin.Engine, gdriveService *drive.Service) {
	driveService = gdriveService
	r.GET("/search-in-doc/:id", SearchInDoc)
}
