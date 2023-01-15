package request

import (
	"core/dto"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type PutRequestContract interface {
	Main(url string, body []byte) dto.Response
}

type PutRequest struct{}

func ProviderPutRequest() PutRequest {
	return PutRequest{}
}

func (p *PutRequest) Main(url string, param_body io.Reader) dto.Response {

	var result dto.Result
	var client = &http.Client{}

	request, err := http.NewRequest(http.MethodPut, url, param_body)
	if err != nil {
		return dto.Response{ResponseCode: http.StatusBadGateway, ErrorMessage: err, Result: result}
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(request)
	if err != nil {
		return dto.Response{ResponseCode: http.StatusBadGateway, ErrorMessage: err, Result: result}
	}
	defer resp.Body.Close()

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
