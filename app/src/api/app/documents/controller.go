package documents

import (
	// "strings"
	// "os"
	"github.com/gin-gonic/gin"
)

// Get /search-in-doc/:id
func SearchInDoc(c *gin.Context) {
	// fileID := strings.TrimSpace(c.Param("id"))
	// //word := strings.TrimSpace(c.Param("word"))

	// driveFile, err := driveService.Files.Get(fileID).Do()
	// if err != nil {
	// 	c.JSON(404, gin.H{"error": "File not found"})
	// 	return
	// }

	// f, err := os.Create("/tmp/" + driveFile.Name)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "An error ocurred while searching word in file"})
	// 	return
	// }
	// defer f.Close()

	// f.WriteString(driveFile.Description)
	// f.Sync()
	return

}