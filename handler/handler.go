package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harry21711/go-url-shortener/shortener"
	"github.com/harry21711/go-url-shortener/store"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {

	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}
