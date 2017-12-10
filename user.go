package jira

import (
	"encoding/json"
	"errors"
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
		err = errors.New("401")
		return
	case 404:
		err = errors.New("404")
		return
	}

	err = json.Unmarshal(res, &user)

	return
}
