package request

import (
	"core/dto"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type PatchRequestContract interface {
	Main(url string, body []byte) dto.Response
}

type PatchRequest struct{}

func ProviderPatchRequest() PatchRequest {
	return PatchRequest{}
}

func (p *PatchRequest) Main(url string) dto.Response {

	var result dto.Result
	var client = &http.Client{}

	var io io.Reader

	request, err := http.NewRequest(http.MethodPatch, url, io)
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
