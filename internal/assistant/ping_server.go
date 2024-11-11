package assistant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) HealthCheck() (string, error) {
	url := baseUrl + "/healthz"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "Failed to create request: ", err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "Failed to send request: ", err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Failing to read response body", err
	}
	var result map[string]interface{}
	err = json.Unmarshal(dat, &result)
	if err != nil {
		return "Couldn't unmarshal response", err
	}
	fmt.Printf("!!!! :%v\n", result)
	return "should be good", nil

}
