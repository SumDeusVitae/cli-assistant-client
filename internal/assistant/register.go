package assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) Register(login, password, email string) (UserRespond, error) {
	url := baseUrl + "/register"
	// handling email
	emailStruct := struct {
		Email string `json:"email"`
		Valid bool   `json:"valid"`
	}{}
	if email != "" {
		emailStruct.Email = email
		emailStruct.Valid = true
	} else {
		emailStruct.Email = ""
		emailStruct.Valid = false
	}
	// Buiding our registration form
	register := RegistrationForm{
		Login:    login,
		Password: password,
		Email:    emailStruct,
	}
	// Marshaling into json
	jsonData, err := json.Marshal(register)
	if err != nil {
		log.Fatal(err)
		return UserRespond{}, err
	}
	// New Post request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return UserRespond{}, err
	}
	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return UserRespond{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		if res.StatusCode == http.StatusConflict {
			return UserRespond{}, fmt.Errorf("login already in use, please choose another")
		}
		return UserRespond{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	// Decode response into UserRespond struct
	// Doing unmarshaling cause response is only 2 strigns
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return UserRespond{}, err
	}
	if len(dat) == 0 {
		return UserRespond{}, fmt.Errorf("empty response body")
	}
	createdRespond := UserRespond{}
	err = json.Unmarshal(dat, &createdRespond)
	if err != nil {
		log.Printf("Failed to unmarshal response: %v\nResponse body: %s", err, string(dat))
		return UserRespond{}, err
	}
	return createdRespond, nil

}
