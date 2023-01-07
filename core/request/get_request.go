package request

import (
	"core/dto"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetRequestContract interface {
	Main(url string) dto.Response
}

type GetRequest struct{}

func ProviderGetRequest() GetRequest {
	return GetRequest{}
}

func (g *GetRequest) Main(url string) dto.Response {

	var result dto.Result

	resp, err := http.Get(url)
	if err != nil {
		return dto.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
	}

	json.Unmarshal([]byte(body), &result)

	return dto.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
}
