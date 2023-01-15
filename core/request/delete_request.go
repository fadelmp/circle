package request

import (
	"core/dto"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type DeleteRequestContract interface {
	Main(url string, body []byte) dto.Response
}

type DeleteRequest struct{}

func ProviderDeleteRequest() DeleteRequest {
	return DeleteRequest{}
}

func (d *DeleteRequest) Main(url string) dto.Response {

	var result dto.Result
	var client = &http.Client{}
	var io io.Reader

	request, err := http.NewRequest(http.MethodDelete, url, io)
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
