package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// RequestCardtokens sends a HTTP request to a specified endpoint with a JSON payload.
// It returns an error and a map containing the response.
//
// Parameters:
// method: The HTTP method to use (e.g., "GET", "POST").
// endpoint: The URL of the endpoint to send the request to.
// mPayload: A map containing the payload to send in the request body.
// apikey: The API key to use for authentication.
func RequestCardtokens(method string, endpoint string, mPayload map[string]string, apikey string) (err error, mResponse map[string]interface{}) {
	//
	// Initialize error to nil and response to an empty string
	//
	err = nil
	var response string = ""
	mResponse = make(map[string]interface{}, 0)

	//
	// Convert the payload map to JSON
	//
	jsonpayload, err := json.Marshal(mPayload)
	if err != nil {
		fmt.Println("Unable to generate json from realtime valudation map: " + err.Error())
		return err, mResponse
	}

	//
	// Create a new HTTP client and request
	//
	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, strings.NewReader(string(jsonpayload)))
	if err != nil {
		fmt.Println("Create http request: " + err.Error())
		return err, mResponse
	}

	//
	// Set the request headers
	//
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("x-api-key", apikey)
	req.Header.Set("User-Agent", "Cardtokens/1")

	//
	// Send the request and get the response
	//
	resp, err := client.Do(req)

	//
	// If there is no response, return an error
	//
	if resp == nil {
		return errors.New("Unable to get response from client"), mResponse
	}

	//
	// If there is an error in the response, return it
	//
	if err != nil {
		fmt.Println("Error communicating to mc ", err)
		return err, mResponse
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, mResponse
	}
	response = string(bodyBytes)

	//
	// Unmarshal the response body into the response map
	//
	err = json.Unmarshal([]byte(response), &mResponse)
	if err != nil {
		err = errors.New("Unable to process response " + err.Error() + " - " + response)
		return err, mResponse
	}

	//
	// If the response status code is 200 or 201 and the response body is not empty, return the response map
	//
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		return nil, mResponse
	} else {
		return errors.New("Invalid response received"), mResponse
	}
}
