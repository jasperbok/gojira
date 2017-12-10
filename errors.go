package jira

import "fmt"

type JiraError struct {
	Code int
	Err  string
}

func (e *JiraError) Error() string {
	return fmt.Sprintf("jira: %d: %s", e.Code, e.Err)
}
