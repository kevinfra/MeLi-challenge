package documents

import (
	"net/http"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Get /search-in-doc/:id ?word=
func SearchInDoc(c *gin.Context) {
	fileID := strings.TrimSpace(c.Param("id"))
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
		return
	}
	word := strings.TrimSpace(c.Query("word"))
	if word == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty word"})
		return
	}

	driveFile, err := Di.Document(fileID)
	if err != nil {
		errormsg := "An error occurred: " + err.Error()
		c.JSON(404, gin.H{"error": errormsg})
		return
	} else {
		c.JSON(200, gin.H{"description": driveFile.Description})
	}
	return

}

func AuthForDrive(c *gin.Context) {
	if Di.Authorized() {
		c.JSON(400, gin.H{"error":"Already authorized"})
	} else {
		token := c.Query("token")
		if token == "" {
			err := Di.LoadFromDB()
			if err != nil {
				c.JSON(200, gin.H{"success":Di.LoadURLForTokenAuth()})
				return
			}
		} else {
			err := Di.LoadFromToken(token)
			if err != nil {
				fmt.Printf(err.Error())
				c.JSON(500, gin.H{"error":err.Error()})
				return
			}
		}
		c.JSON(200, gin.H{"success":"authorized"})
	}
	return
}
