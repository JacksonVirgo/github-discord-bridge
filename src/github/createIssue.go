package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CreateIssueRequest struct {
	Owner   string   `json:"owner"`
	Repo    string   `json:"repo"`
	Title   string   `json:"title"`
	Body    string   `json:"body"`
	Labels  []string `json:"labels"`
	Headers Headers  `json:"headers"`
}

type Headers struct {
	XGitHubApiVersion string `json:"X-GitHub-Api-Version"`
}

type CreateIssueResponse struct {
	Number  int    `json:"number"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

func CreateIssue(newIssue CreateIssueRequest) (CreateIssueResponse, error) {
	newIssue.Owner = GithubContext.author
	newIssue.Repo = GithubContext.repo
	newIssue.Headers.XGitHubApiVersion = "2022-11-28"

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", newIssue.Owner, newIssue.Repo)
	reqBody, err := json.Marshal(newIssue)
	if err != nil {
		log.Printf("Error marshalling request body: %s", err)
		return CreateIssueResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return CreateIssueResponse{}, err
	}

	req.Header.Set("Authorization", "token "+GithubContext.token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return CreateIssueResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return CreateIssueResponse{}, err
	}

	var response CreateIssueResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		return CreateIssueResponse{}, err
	}

	return response, err
}
