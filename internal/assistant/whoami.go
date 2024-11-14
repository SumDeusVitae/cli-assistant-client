package assistant

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) Whoami(api string) (UserRespond, error) {
	url := baseUrl + "/whoami"
	apiKey := "ApiKey " + api

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return UserRespond{}, err
	}
	req.Header.Set("Authorization", apiKey)
	// Sending
	res, err := c.httpClient.Do(req)
	if err != nil {
		return UserRespond{}, err
	}
	defer res.Body.Close()

	// Check status code.
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return UserRespond{}, errors.New("unauthorized")
		}
		return UserRespond{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// Decode
	createRespond := UserRespond{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&createRespond)
	if err != nil {
		return UserRespond{}, err
	}
	return createRespond, nil

}
