// Copyright (c) 2015 LunaNode Hosting Inc. All right reserved.
// Use of this source code is governed by the MIT License. See LICENSE file.

package linode

import "io/ioutil"
import "net/http"
import "net/http/httptest"
import "strings"
import "testing"

func setupTestServer(t *testing.T, expectedParams string, response string) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected method POST but got %s", r.Method)
			return
		}
		// validate params
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		contentsTrimmed := strings.TrimSpace(string(contents))
		if contentsTrimmed != expectedParams {
			t.Errorf("received parameters [%s], does not match expected [%s]", contentsTrimmed, expectedParams)
			return
		}
		// send response
		w.Write([]byte(response))
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}

func setupTestClient(url string) *Client {
	return &Client{
		APIKey: "testkey",
		HTTPClient: &http.Client{},
		apiURL: url,
	}
}

func TestRequestParams(t *testing.T) {
	type Response struct {
		Foo string `json:"foo"`
	}
	server := setupTestServer(t, "api_action=testaction&api_key=testkey&param1=value1&param2=value2", `{"ERRORARRAY":[],"ACTION":"testaction","DATA":{"foo":"bar"}}`)
	client := setupTestClient(server.URL)
	response := new(Response)
	err := client.request("testaction", map[string]string{"param1": "value1", "param2": "value2"}, response)
	if t.Failed() {
		return
	} else if err != nil {
		t.Fatal(err)
	} else if response.Foo != "bar" {
		t.Fatalf("response foo=%s but expected bar", response.Foo)
	}
}

func TestAPIError(t *testing.T) {
	server := setupTestServer(t, "api_action=testaction&api_key=testkey", `{"ERRORARRAY":[{"ERRORCODE":3,"ERRORMESSAGE":"The requested class does not exist"}],"DATA":{},"ACTION":"testaction"}`)
	client := setupTestClient(server.URL)
	err := client.request("testaction", nil, nil)
	if t.Failed() {
		return
	} else if err == nil {
		t.Fatal("expected API error but got nil")
	} else if !strings.Contains(err.Error(), "API error") {
		t.Fatalf("expected API error but got %s", err.Error())
	}
}

func TestMalformedResponse(t *testing.T) {
	server := setupTestServer(t, "api_action=testaction&api_key=testkey", `{"ERRORARRAY":[],"ACTION":"testaction","DATA":{`)
	client := setupTestClient(server.URL)
	err := client.request("testaction", nil, nil)
	if t.Failed() {
		return
	} else if err == nil {
		t.Fatal("expected json error but got nil")
	}
}
