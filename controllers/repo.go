// controllers/books.go

package controllers

import (
	// "bookc/models"
	"goapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateRepoInput struct {
	RepoName   string `json:"repoName"`
	RepoDetail string `json:"repoDetail"`
	RepoURL    string `json:"repoURL`
}

// GET /tasks
// Get all tasks
func FindRepos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var repos []models.Repo
	db.Find(&repos)

	c.JSON(http.StatusOK, gin.H{"data": repos})
}
