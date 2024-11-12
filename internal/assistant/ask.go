package assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) Ask(model, API, question string) (QuestionRespond, error) {
	url := baseUrl + "/ask"

	questionForm := QuestionForm{
		Model:   model,
		Request: question,
	}
	jsonData, err := json.Marshal(questionForm)
	if err != nil {
		log.Fatal(err)
		return QuestionRespond{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return QuestionRespond{}, err
	}
	// SET HEADERS
	req.Header.Set("Content-Type", "application/json")
	apiKey := "ApiKey " + API
	req.Header.Set("Authorization", apiKey)

	// send
	res, err := c.httpClient.Do(req)
	if err != nil {
		return QuestionRespond{}, err
	}
	defer res.Body.Close()

	// Check status code.
	if res.StatusCode != http.StatusCreated {
		return QuestionRespond{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// Decode
	createRespond := QuestionRespond{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&createRespond)
	if err != nil {
		return QuestionRespond{}, err
	}
	return createRespond, nil

}
