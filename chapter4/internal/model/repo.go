package model

import "time"

// Repository はGitHubリポジトリの情報を表す構造体。
type Repository struct {
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Stars       int       `json:"stargazers_count"`
	Forks       int       `json:"forks_count"`
	HTMLURL     string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
