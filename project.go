package jira

import (
	"encoding/json"
	"fmt"
)

type Project struct {
	Id          string
	Key         string
	Description string
	Name        string
	Lead        User
}

func (p *Project) LeadDisplayName() string {
	return p.Lead.DisplayName
}

func (p *Project) Print() {
	fmt.Printf("%s (%s)", p.Name, p.Key)
}

// GetProject gets a project's information and returns it as a Project.
func GetProject(keyOrId string) (p Project, err error) {
	status, res, err := callApi("GET", "/project/"+keyOrId, nil)

	if err != nil {
		return
	}

	if status == 404 {
		return p, &JiraError{404, "Project not found"}
	}

	err = json.Unmarshal(res, &p)

	return
}
