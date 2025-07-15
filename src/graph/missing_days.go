package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const githubAPI = "https://api.github.com/graphql"

type GraphQLRequest struct {
	Query string `json:"query"`
}

type ContributionDay struct {
	Date              string `json:"date"`
	ContributionCount int    `json:"contributionCount"`
}

type GraphQLResponse struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					Weeks []struct {
						ContributionDays []ContributionDay `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}

func Get_days() {
	username := os.Getenv("GITHUB_USER")
	token := os.Getenv("GITHUB_TOKEN")
	todayDate := time.Now().Format("2006-01-02")
	query := fmt.Sprintf(`
query {
  user(login: "%s") {
    contributionsCollection(from: "2025-01-01T00:00:00Z", to: "%sT23:59:59Z") {
      contributionCalendar {
        weeks {
          contributionDays {
            date
            contributionCount
          }
        }
      }
    }
  }
}
`, username, todayDate)

	payload := GraphQLRequest{Query: query}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", githubAPI, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result GraphQLResponse
	json.Unmarshal(body, &result)

	missingDays := []string{}

	for _, week := range result.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			if day.ContributionCount == 0 {
				missingDays = append(missingDays, day.Date)
			}
		}
	}

	fmt.Println("Days with zero contributions:")
	for _, day := range missingDays {
		fmt.Println(day)
	}
}
