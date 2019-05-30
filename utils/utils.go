package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var HTTPTypes = [4]string{"GET", "POST", "PATCH", "DELETE"}
var client = &http.Client{Timeout: 10 * time.Second}

func isHttpType(value string) bool {
	if value == "" {
		return false
	}
	for _, httpType := range HTTPTypes {
		if strings.ToLower(value) == strings.ToLower(httpType) {
			return true
		}
	}
	return false
}

func MakeAPIRequest(requestURL, httpType string, output interface{}) error {
	// Validate HTTP type
	if httpTypeValid := isHttpType(httpType); httpTypeValid != true {
		return errors.New("invalid http type: " + httpType)
	}

	// Build request
	request, err := http.NewRequest(httpType, requestURL, nil)
	if err != nil {
		return err
	}

	// Add authorisation token to header
	request.SetBasicAuth(os.Getenv("JENKINS_USERNAME"), os.Getenv("JENKINS_API_TOKEN"))

	// Execute request
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if !strings.HasPrefix(strconv.Itoa(response.StatusCode), "2") {
		return fmt.Errorf("%s %d %s", "HTTP Status Code", response.StatusCode, "returned")
	}

	var data []byte
	response.Body.Read(data)
	if len(data) != 0 {
		return json.Unmarshal(data, output)
	}
	return nil
}
