package items

import (
	"net/http"
	"strings"

	"api/app/models"

	"github.com/gin-gonic/gin"
)

// GetItem ...
func GetItem(c *gin.Context) {
	itemID := strings.TrimSpace(c.Param("id"))
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
		return
	}

	item, err := Is.Item(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}
	c.JSON(200, item)
	return
}

// GetItems ...
func GetItems(c *gin.Context) {
	items, err := Is.Items()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}
	c.JSON(200, items)
	return
}

// PostItem ...
func PostItem(c *gin.Context) {
	i := &models.Item{}
	if err := c.BindJSON(i); c.Request.ContentLength == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bind_error", "description": err.Error()})
		return
	}
	err := Is.CreateItem(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "save_error", "description": err.Error()})
		return
	}
	c.JSON(201, i)
}

// DeleteItem
func DeleteItem(c *gin.Context) {
	itemID := strings.TrimSpace(c.Param("id"))
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
		return
	}

	err := Is.DeleteItem(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id":itemID})
	return
}

// GetNullItem
func GetNullItem(c *gin.Context) {
	c.JSON(200, gin.H{"id":0, "name": "", "description": ""})
	return
}
