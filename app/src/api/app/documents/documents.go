package documents

import (
	//"strings"
	//"net/http"
	"api/app/models"
	"database/sql"
	"github.com/gin-gonic/gin"
    "google.golang.org/api/drive/v3"
)

var (
	Di models.DocumentServiceInterface
)

func Configure(r *gin.Engine, gdriveService *drive.Service, db *sql.DB) {
	Di = &DocumentService{DB: db, GDrive: gdriveService}
	r.GET("/search-in-doc/:id", SearchInDoc)
}
