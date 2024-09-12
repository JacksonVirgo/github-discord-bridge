package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Issue struct {
	Title string `json:"title"`
}

func GetIssues() ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", GithubContext.author, GithubContext.repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	req.Header.Set("Authorization", "token "+GithubContext.token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return nil, err
	}

	var issues []Issue
	err = json.Unmarshal(body, &issues)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		return nil, err
	}

	var titles []string
	for _, issue := range issues {
		titles = append(titles, issue.Title)
	}

	return titles, nil
}
