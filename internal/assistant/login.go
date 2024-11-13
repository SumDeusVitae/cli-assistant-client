package assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) Login(login, password string) (LoginRespond, error) {
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
		return LoginRespond{}, err
	}
	// New POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return LoginRespond{}, err
	}
	// Set header
	req.Header.Set("Content-Type", "application/json")

	// Send
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LoginRespond{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return LoginRespond{}, fmt.Errorf("wrong login or password")
		}
		return LoginRespond{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	// Decode
	var createRespond LoginRespond
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&createRespond)
	if err != nil {
		return LoginRespond{}, err
	}
	return createRespond, nil

}
