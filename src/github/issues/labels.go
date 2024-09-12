package issues

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/JacksonVirgo/github-discord-bridge/src/github"
)

type Label struct {
	Name string `json:"name"`
}

func GetIssueLabels() ([]Label, error) {
	token := github.GetToken()
	repo := github.GetRepo()
	author := github.GetAuthor()

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/labels", author, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
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

	var labels []Label
	err = json.Unmarshal(body, &labels)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		return nil, err
	}

	return labels, nil
}
