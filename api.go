// Package jira provides functions to work with the Jira REST API.
package jira

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

const API_URL = "https://fluxility.atlassian.net/rest/api/latest"

func jiraGetAuth() (username, password string) {
	return os.Getenv("JIRA_USER"), os.Getenv("JIRA_PASS")
}

func callApi(method string, path string, data []byte) (statusCode int, body []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, API_URL+path, bytes.NewBuffer(data))

	if err != nil {
		return
	}

	req.SetBasicAuth(jiraGetAuth())
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	res.Body.Close()
	statusCode = res.StatusCode
	return
}
