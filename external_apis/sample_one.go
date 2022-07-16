package external_apis

import (
	"bytes"
	"encoding/json"
	"github.com/web-app-with-golang/client"
	env "github.com/web-app-with-golang/project_env"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func SendSampleRequest() []byte {
	endpoint := "https://httpbin.org/post"
	payload := map[string]string{"name": "faisal"}
	requestParams, err := json.Marshal(payload)
	myClient := client.MyHttpClient()

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(requestParams))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := myClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	if env.GoDotEnvVariable("DEBUG") == "1" {
		log.Println("Status code: ", response.StatusCode)
	}

	// Close the connection to reuse it
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	if env.GoDotEnvVariable("DEBUG") == "1" {
		log.Println(string(responseBody))
	}
	return responseBody
}
