package request

import (
	"customer/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HTTPRequestContract interface {
	Get(url string) entity.Response
}

type HTTPRequest struct{}

func ProviderRequest() HTTPRequest {
	return HTTPRequest{}
}

func (r *HTTPRequest) Get(url string) entity.Response {

	var result entity.Result

	resp, err := http.Get(url)
	if err != nil {
		return entity.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
	}

	json.Unmarshal([]byte(body), &result)

	return entity.Response{ResponseCode: resp.StatusCode, ErrorMessage: err, Result: result}
}
