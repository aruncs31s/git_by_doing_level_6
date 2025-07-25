package remote

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
	Author struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
	} `json:"author"`
}

// the following function does this request to github
/*
	curl -L \
	-H "Accept: application/vnd.github+json" \
	-H "Authorization: Bearer <YOUR-TOKEN>" \
	-H "X-GitHub-Api-Version: 2022-11-28" \
	https://api.github.com/repos/OWNER/REPO/activity

*/

// NOTE: This can be run in parrallel if they wait longer. or the selected option is play
func GetGithubCommits() int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/aruncs31s/git_by_doing_level_6/commits", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// fmt.Printf("API Error: %d %s\n", resp.StatusCode, resp.Status)
		return 0
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var commits []Commit
	err = json.Unmarshal(body, &commits)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// for i, commit := range commits {
	// 	// if i >= 5 { // Show only first 5 commits
	// 	// 	break
	// 	// }

	// 	fmt.Printf("🔹 Commit #%d\n", i+1)
	// 	fmt.Printf("   SHA: %s\n", commit.SHA[:8]) // Show first 8 characters
	// 	fmt.Printf("   Author: %s (%s)\n", commit.Commit.Author.Name, commit.Author.Login)
	// 	fmt.Printf("   Date: %s\n", commit.Commit.Author.Date)
	// 	fmt.Printf("   Message: %s\n", commit.Commit.Message)
	// }

	// for _, commit := range commits {
	// 	fmt.Println(commit.SHA[:7])
	// }
	// fmt.Print("📊 Total Num Commits ")
	// fmt.Println(len(commits))
	// return commit_lists
	return len(commits)
}
