package common

import (
	"fmt"
	"net/http"
)

// HttpGetRequest - Common Utility Function for SUPPLY API HTTP GET Requests
func HttpGetRequest(path string) (*http.Response, error) {

	// Create API HTTP GET Request
	request, _ := http.NewRequest("GET", path, nil)
	//sessionToken, err := GetSessionToken()
	//HandleError(err)
	//bearerToken := string("Bearer " + sessionToken)

	bearerToken := "SomeBearerToken"
	request.Header.Add("Authorization", bearerToken)

	request.Header.Add("Authorization", bearerToken)
	DebugMessage("[common] func HttpGetRequest() --> Adding Bearer Token to Authorization Header")
	DebugMessage(bearerToken)

	// Create HTTP Client and Do Request
	DebugMessage("[common] func HttpGetRequest() --> Submitting HTTP Request to API")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	DebugHttpRequest(request, "HttpGetRequest() API Request")

	// Handle HTTP Response Status Codes
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		DebugMessage(fmt.Sprintf("[common] func HttpGetRequest() --> HTTP Status %v\n", response.StatusCode))
		return response, nil
	} else {
		DebugMessage(fmt.Sprintf("[common] func HttpGetRequest() --> HTTP Status %v\n", response.StatusCode))
		httpError := fmt.Errorf("ERROR: HTTP Status Code %v: %v", response.StatusCode, path)
		return response, httpError
	}

}
