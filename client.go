// Copyright (c) 2015 LunaNode Hosting Inc. All right reserved.
// Use of this source code is governed by the MIT License. See LICENSE file.

package linode

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"

const API_URL = "https://api.linode.com/"

// Linode API client.
// Note that Linode API encodes booleans as JSON integers, and returned objects will correspondingly contain integer fields.
type Client struct {
	// The Linode API key.
	APIKey string

	// An HTTP client to perform API requests.
	HTTPClient *http.Client

	apiURL string
}

type apiError struct {
	Code int `json:"ERRORCODE"`
	Message string `json:"ERRORMESSAGE"`
}

type genericResponse struct {
	Errors []apiError `json:"ERRORARRAY"`
	Action string `json:"ACTION"`
	Data interface{} `json:"DATA"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{},
		apiURL: API_URL,
	}
}

func (client *Client) request(action string, params map[string]string, dataTarget interface{}) error {
	// setup post parameters
	postParams := make(url.Values)
	postParams.Set("api_key", client.APIKey)
	postParams.Set("api_action", action)
	if params != nil {
		for key, value := range params {
			postParams.Set(key, value)
		}
	}

	// determine URL to use
	apiURL := API_URL
	if client.apiURL != "" {
		apiURL = client.apiURL
	}

	// do request
	response, err := client.HTTPClient.PostForm(apiURL, postParams)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// unmarshal json
	responseTarget := new(genericResponse)
	responseTarget.Data = dataTarget
	err = json.Unmarshal(contents, responseTarget)
	if err != nil {
		return err
	} else if responseTarget.Action != action {
		return fmt.Errorf("expected %s for API response action, but got %s", action, responseTarget.Action)
	}

	// check for non-0 errors (0 is "ok" error)
	for _, apiErr := range responseTarget.Errors {
		if apiErr.Code != 0 {
			return fmt.Errorf("API error (%s) %d: %s", action, apiErr.Code, apiErr.Message)
		}
	}
	return nil
}
