package request

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type PostRequestContract interface {
	Main(url string, body []byte) (interface{}, error)
}

type PostRequest struct{}

func ProviderPostRequest() PostRequest {
	return PostRequest{}
}

func (p *PostRequest) Main(url string, param_body io.Reader) (interface{}, error) {

	var result interface{}
	app_type := "application/json"

	resp, err := http.Post(url, app_type, param_body)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	json.Unmarshal([]byte(body), &result)
	return body, nil
}
