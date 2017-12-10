package jira

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Issue represents a Jira issue.
type Issue struct {
	Id     string
	Key    string
	Fields Fields
}

type IssueType struct {
	Id          string
	Name        string
	Description string
	Subtask     bool
}

type Fields struct {
	Summary     string
	Description string
	Issuetype   IssueType
	Creator     User
	Reporter    User
	Assignee    User
	Project	    Project
	Comments    []Comment
}
func (i *Issue) CreatorDisplayName() string {
	if i.Fields.Creator == (User{}) {
		return "Unassigned"
	}
	return i.Fields.Creator.DisplayName
}

func (i *Issue) ReporterDisplayName() string {
	if i.Fields.Reporter == (User{}) {
		return "Unassigned"
	}
	return i.Fields.Reporter.DisplayName
}

func (i *Issue) AssigneeDisplayName() string {
	if i.Fields.Assignee == (User{}) {
		return "Unassigned"
	}
	return i.Fields.Assignee.DisplayName
}

// Print formats an issue and prints it to stdout.
func (i *Issue) Print() {
	fmt.Printf("%s - %s\n", i.Key, i.Fields.Summary)
	fmt.Printf("Reporter: %s   Assignee: %s\n", i.ReporterDisplayName(), i.AssigneeDisplayName())

	if i.Fields.Description != "" {
		fmt.Printf("\n%s\n", i.Fields.Description)
	}

	if len(i.Fields.Comments) > 0 {
		fmt.Printf("Comments:\n")
		for _, c := range i.Fields.Comments {
			fmt.Printf("\n%s:\n%s\n", c.Author.DisplayName, c.Body)
		}
	}
}

// GetIssue calls the API's GET Issue endpoint.
func GetIssue(key string) (i Issue, err error) {
	status, res, err := callApi("GET", "/issue/"+key, nil)

	if err != nil {
		return
	}

	if status == 404 {
		return i, &JiraError{404, "Issue not found"}
	}

	err = json.Unmarshal(res, &i)

	return
}
