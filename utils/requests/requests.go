package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const PATCH = "PATCH"
const DELETE = "DELETE"

type Headers map[string]string
type Payload map[string]interface{}

type Response struct {
	Body []byte
	*http.Response
}

func (resp *Response) BindJSON(destination interface{}) error {
	return json.Unmarshal(resp.Body, destination)
}

type Request struct {
	Method  string
	Payload interface{}
	Headers Headers
}

func (req Request) Send(url string, destination interface{}) (*Response, error) {
	var payload io.Reader = nil
	if req.Payload != nil {
		jsonBytes, err := json.Marshal(req.Payload)
		if err != nil {
			return nil, err
		} else {
			payload = bytes.NewReader(jsonBytes)
		}
	}

	method := http.MethodGet
	if req.Method != "" {
		method = req.Method
	}

	httpReq, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if req.Headers != nil {
		for key, value := range req.Headers {
			httpReq.Header.Set(key, value)
		}
	}
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		Body:     body,
		Response: httpResp,
	}

	if destination != nil {
		err = resp.BindJSON(destination)
	}

	return resp, err
}
