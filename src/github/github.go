package github

import (
	"errors"
	"os"
)

type GithubContextStructure struct {
	token         string
	repo          string
	author        string
	webhookSecret string
}

var GithubContext = &GithubContextStructure{}

func LoadGithubContext() error {
	githubToken := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")
	author := os.Getenv("GITHUB_AUTHOR")
	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")

	if githubToken == "" || repo == "" || author == "" {
		return errors.New("missing environment variables")
	}

	*GithubContext = GithubContextStructure{
		token:         githubToken,
		repo:          repo,
		author:        author,
		webhookSecret: secret,
	}

	return nil
}

func GetRepo() string {
	return GithubContext.repo
}

func GetAuthor() string {
	return GithubContext.author
}

func GetToken() string {
	return GithubContext.token
}

func GetWebhookSecret() string {
	return GithubContext.webhookSecret
}
