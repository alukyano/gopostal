package postal

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeDirectRequest(requestURL string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("Could not create request: %s\n", err)
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		return "", err
	}

	//fmt.Printf("client: got response!\n")
	//fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s\n", err)
		return "", err
	}
	//fmt.Printf("Response body: %s\n", resBody)
	return string(resBody), nil
}
