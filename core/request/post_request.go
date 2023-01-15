package request

import (
	"core/dto"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type PostRequestContract interface {
	Main(url string, body []byte) dto.Response
}

type PostRequest struct{}

func ProviderPostRequest() PostRequest {
	return PostRequest{}
}

func (p *PostRequest) Main(url string, param_body io.Reader) dto.Response {

	var result dto.Result
	app_type := "application/json"

	resp, err := http.Post(url, app_type, param_body)
	if err != nil {
		return dto.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
	}

	json.Unmarshal([]byte(body), &result)

	return dto.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
}
