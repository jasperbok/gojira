package jira

import "fmt"

type Comment struct {
	Id     string
	Author User
	Body   string
}

func (c *Comment) AuthorDisplayName() string {
	return c.Author.DisplayName
}

func (c *Comment) Print() {
	fmt.Printf("%s:\n%s", c.Author.DisplayName, c.Body)
}
