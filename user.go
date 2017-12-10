package jira

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string
	DisplayName string
}

func (u *User) Print() {
	fmt.Printf("%s (%s)\n", u.DisplayName, u.Name)
}

func GetUser(usernameOrKey string, usingUsername bool) (user User, err error) {
	endpoint := "/user/?"
	if usingUsername {
		endpoint = endpoint+"username="+usernameOrKey
	} else {
		endpoint = endpoint+"key="+usernameOrKey
	}

	status, res, err := callApi("GET", endpoint, nil)

	switch status {
	case 401:
		err = &JiraError{401, "You are not authenticated"}
		return
	case 404:
		err = &JiraError{404, "User not found"}
		return
	}

	err = json.Unmarshal(res, &user)

	return
}
