package common

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HttpPostRequest - Common Utility Function for HTTP POST Requests
func HttpPostRequest(method string, path string, json string) (*http.Response, error) {

	var body io.ReadSeeker
	body = strings.NewReader(string(json))

	// Create API HTTP Request
	request, _ := http.NewRequest(method, path, body)
	//sessionToken, err := GetSessionToken()
	//HandleError(err)
	//bearerToken := string("Bearer " + sessionToken)

	bearerToken := "SomeBearerToken"
	request.Header.Add("Authorization", bearerToken)

	// Create HTTP Client and POST Request
	DebugMessage("[common] func HttpRequest() --> Submitting HTTP Request")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	DebugHttpRequest(request, "HttpPostRequest()")

	// Handle HTTP Response Status Codes
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		DebugMessage(fmt.Sprintf("[common] func HttpPostRequest() --> HTTP Status %v\n", response.StatusCode))
		return response, nil
	} else {
		DebugMessage(fmt.Sprintf("[common] func HttpPostRequest() --> HTTP Status %v\n", response.StatusCode))
		httpError := fmt.Errorf("ERROR: HTTP Status Code %v: %v", response.StatusCode, path)
		return response, httpError
	}

}
