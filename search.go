package jira

import (
	"encoding/json"
	"fmt"
	"strings"
)

const IN_PROGRESS_STATUSES string = "In Progress"
const TODO_PROGRESS_STATUSES string = `"To Do","Open","Selected for Development"`
const CLOSED_PROGRESS_STATUSES string = "Closed,Ready For Client,Done"
const BACKLOG_STATUSES string = "Backlog"

type SearchResult struct {
	MaxResults int
	Total      int
	Issues     []Issue
}

func (s *SearchResult) Print() {
	longestKey := 0
	longestSummary := 0

	for _, i := range s.Issues {
		keyLen := len(i.Key)
		summaryLen := len(i.Fields.Summary)

		if keyLen > longestKey {
			longestKey = keyLen
		}
		if summaryLen > longestSummary {
			longestSummary = summaryLen
		}
	}

	fmtStr := fmt.Sprintf("%%-%ds  %%-%ds\n", longestKey, longestSummary)

	for _, i := range s.Issues {
		fmt.Printf(fmtStr, i.Key, i.Fields.Summary)
	}
}

// Search executes the given JQL query and returns a SearchResult.
func Search(query string) (res SearchResult, err error) {
	query = strings.Replace(query, `"`, `\"`, -1)
	q := []byte(fmt.Sprintf(`{"jql": "%s"}`, query))
	status, result, err := callApi("POST", "/search/", q)

	if err != nil {
		return
	}

	if status == 400 {
		return res, &JiraError{400, "There was a problem with the search query"}
	}

	err = json.Unmarshal(result, &res)

	return
}
