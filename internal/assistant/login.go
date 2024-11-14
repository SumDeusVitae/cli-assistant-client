package assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) Login(login, password string) (UserRespond, error) {
	url := baseUrl + "/login"
	// Building login form
	loginForm := LoginForm{
		Login:    login,
		Password: password,
	}
	// Marshal into json
	jsonData, err := json.Marshal(loginForm)
	if err != nil {
		log.Fatal(err)
		return UserRespond{}, err
	}
	// New POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return UserRespond{}, err
	}
	// Set header
	req.Header.Set("Content-Type", "application/json")

	// Send
	res, err := c.httpClient.Do(req)
	if err != nil {
		return UserRespond{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return UserRespond{}, fmt.Errorf("wrong login or password")
		}
		return UserRespond{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	// Decode
	var createRespond UserRespond
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&createRespond)
	if err != nil {
		return UserRespond{}, err
	}
	return createRespond, nil

}
