package controllers

import (
	"net/http"

	"github.com/adrianomonteiroweb/library-golang-api/database"
	"github.com/adrianomonteiroweb/library-golang-api/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
    var books []models.Book
    database.DB.Find(&books)
    c.JSON(http.StatusOK, gin.H{"data": books})
}
