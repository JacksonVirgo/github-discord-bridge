package controllers

import (
	"errors"
	"os"
)


type GithubContext struct {
	token string
	repo string
	author string
}

var githubContext *GithubContext

func LoadGithubContext() error {
	githubToken := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")
	author := os.Getenv("GITHUB_AUTHOR")

	if githubToken == "" || repo == "" || author == "" {
		return errors.New("missing environment variables")
	}

	*githubContext = GithubContext{
		token: githubToken,
		repo: repo,
		author: author,
	}

	return nil
}
