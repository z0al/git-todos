// Copyright © 2017 Ahmed T. Ali <ah.tajelsir@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package todos

import (
	// Native
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	// Packages
	"github.com/ttacon/chalk"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/log"
)

// Provider interface
// ------------------
type Provider interface {
	// Name returns a friendly <Provider> name
	Name() string
	// Search for issues on <Provider>
	Search(term string) []Todo
}

// GitHubProvider for importing issues from GitHub
// -----------------------------------------------
type GitHubProvider struct {
	URL string
}

// Name returns a friendly GitHub name
func (gh GitHubProvider) Name() string {
	return "GitHub"
}

// Search for issues on GitHub
func (gh GitHubProvider) Search(term string) []Todo {
	todosList := []Todo{}
	pageSize := "15"

	info, err := url.Parse(gh.URL)
	if err != nil {
		log.Error("An error occured while paring the remote URL")
		os.Exit(1)
	}

	var apiURL string
	if info.Host == "github.com" {
		apiURL = "https://api.github.com/search/issues"
	} else {
		// Assuming GitHub Enterprise
		apiURL = info.Scheme + "://" + info.Host + "/api/v3/search/issues"
	}

	// Search parameters
	query := url.Values{}
	q := term + "+repo:" + getRepoSlug(info.Path)
	query.Add("q", q)
	query.Add("sort", "created")
	query.Add("order", "desc")
	query.Add("per_page", pageSize)

	// Get issues
	res, err := http.Get(apiURL + "?" + query.Encode())
	if err != nil {
		log.Error(
			fmt.Sprintf(
				"An error occured while fetching issues from %s",
				chalk.Yellow.Color(apiURL),
			),
		)
		os.Exit(1)
	} else if res.StatusCode != 200 {
		log.Error(
			fmt.Sprintf(
				"An error occured while fetching issues from %s (code %d)",
				chalk.Yellow.Color(apiURL),
				res.StatusCode,
			),
		)
		fmt.Println(query)
		os.Exit(1)
	}
	defer res.Body.Close()

	// Parse result
	data := GitHubResponse{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Error("Something went wrong. Please try again :/")
		os.Exit(1)
	}

	if data.Total > 0 {
		// Page size hard coded :p
		if data.Total > 15 {
			log.Warn(
				fmt.Sprintf(
					"Showing up to %s out of %d matches found", pageSize, data.Total,
				),
			)
		}
		// Copy issues
		for _, i := range data.Items {
			todosList = append(
				todosList, Todo{Title: i.Title, Description: i.Body, ID: i.ID},
			)
		}
	} else {
		log.Warn("No items found matching the search query")
	}

	return todosList
}

// Helpers
// -------
func getRepoSlug(path string) string {
	str := []rune(path)
	str = str[1:]          // Removes '/' prefix
	str = str[:len(str)-4] // Removes '.git' suffix
	return string(str)
}

// GitHubIssue type
type GitHubIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	ID    int    `json:"number"`
}

// GitHubResponse type
type GitHubResponse struct {
	Total int           `json:"total_count"`
	Items []GitHubIssue `json:"items"`
}
