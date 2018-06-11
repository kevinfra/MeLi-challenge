package documents

import (
	//"strings"
	//"net/http"
	"api/app/models"
	"database/sql"
	"github.com/gin-gonic/gin"
)

var (
	Di models.DocumentServiceInterface
)

func Configure(r *gin.Engine, db *sql.DB) {
	Di = &DocumentService{DB: db, GDrive: nil, ClientConfig: nil}
	Di.InitializeConfig()
	r.GET("/search-in-doc/:id", SearchInDoc)
	r.POST("/file", CreateFile)
	r.POST("/auth-for-drive", AuthForDrive)
}
