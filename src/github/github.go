package github

import (
	"errors"
	"os"
)

type GithubContextStructure struct {
	token  string
	repo   string
	author string
}

var GithubContext = &GithubContextStructure{}

func LoadGithubContext() error {
	githubToken := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")
	author := os.Getenv("GITHUB_AUTHOR")

	if githubToken == "" || repo == "" || author == "" {
		return errors.New("missing environment variables")
	}

	*GithubContext = GithubContextStructure{
		token:  githubToken,
		repo:   repo,
		author: author,
	}

	return nil
}

func GetRepo() string {
	return GithubContext.repo
}

func GetAuthor() string {
	return GithubContext.author
}
