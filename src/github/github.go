package github

import (
	"errors"
	"os"
)

type GithubContext struct {
	token  string
	repo   string
	author string
}

var GH_Context = &GithubContext{}

func LoadGithubContext() error {
	githubToken := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")
	author := os.Getenv("GITHUB_AUTHOR")

	if githubToken == "" || repo == "" || author == "" {
		return errors.New("missing environment variables")
	}

	*GH_Context = GithubContext{
		token:  githubToken,
		repo:   repo,
		author: author,
	}

	return nil
}

func GetRepo() string {
	return GH_Context.repo
}

func GetAuthor() string {
	return GH_Context.author
}
