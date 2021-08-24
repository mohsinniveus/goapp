// models/book.go

package models

import (
	"time"
)

type Repo struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	RepoName   string    `json:"repoName"`
	RepoDetail string    `json:"repoDetail"`
	RepoURL    string    `json:"repoURL"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
