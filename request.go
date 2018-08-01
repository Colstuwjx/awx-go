package awx

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type APIRequest struct {
	Method   string
	Endpoint string
	Payload  io.Reader
	Headers  http.Header
	Suffix   string
}

func (ar *APIRequest) SetHeader(key string, value string) *APIRequest {
	ar.Headers.Set(key, value)
	return ar
}

func NewAPIRequest(method string, endpoint string, payload io.Reader) *APIRequest {
	var headers = http.Header{}
	var suffix string
	ar := &APIRequest{method, endpoint, payload, headers, suffix}
	return ar
}

type BasicAuth struct {
	Username string
	Password string
}

type Requester struct {
	Base      string
	BasicAuth *BasicAuth
	Client    *http.Client
}

func (r *Requester) Do(ar *APIRequest, responseStruct interface{}, options ...interface{}) (*http.Response, error) {
	if !strings.HasSuffix(ar.Endpoint, "/") && ar.Method != "POST" {
		ar.Endpoint += "/"
	}

	URL, err := url.Parse(r.Base + ar.Endpoint + ar.Suffix)
	if err != nil {
		return nil, err
	}

	for _, o := range options {
		switch v := o.(type) {
		case map[string]string:
			querystring := make(url.Values)
			for key, val := range v {
				querystring.Set(key, val)
			}

			URL.RawQuery = querystring.Encode()
		}
	}

	var req *http.Request
	req, err = http.NewRequest(ar.Method, URL.String(), ar.Payload)
	if err != nil {
		return nil, err
	}

	if r.BasicAuth != nil {
		req.SetBasicAuth(r.BasicAuth.Username, r.BasicAuth.Password)
	}

	for k := range ar.Headers {
		req.Header.Add(k, ar.Headers.Get(k))
	}

	if response, err := r.Client.Do(req); err != nil {
		return nil, err
	} else {
		switch responseStruct.(type) {
		case *string:
			return r.ReadRawResponse(response, responseStruct)
		default:
			return r.ReadJSONResponse(response, responseStruct)
		}
	}
}

func (r *Requester) ReadRawResponse(response *http.Response, responseStruct interface{}) (*http.Response, error) {
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if str, ok := responseStruct.(*string); ok {
		*str = string(content)
	} else {
		return nil, fmt.Errorf("Could not cast responseStruct to *string")
	}

	return response, nil
}

func (r *Requester) ReadJSONResponse(response *http.Response, responseStruct interface{}) (*http.Response, error) {
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(responseStruct)
	return response, nil
}

func (r *Requester) Get(endpoint string, responseStruct interface{}, querystring map[string]string) (*http.Response, error) {
	ar := NewAPIRequest("GET", endpoint, nil)
	ar.Suffix = ""
	return r.Do(ar, responseStruct, querystring)
}

func (r *Requester) GetJSON(endpoint string, responseStruct interface{}, query map[string]string) (*http.Response, error) {
	ar := NewAPIRequest("GET", endpoint, nil)
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ar, &responseStruct, query)
}

func (r *Requester) Post(endpoint string, payload io.Reader, responseStruct interface{}, querystring map[string]string) (*http.Response, error) {
	ar := NewAPIRequest("POST", endpoint, payload)
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ar, &responseStruct, querystring)
}

func (r *Requester) Put(endpoint string, payload io.Reader, responseStruct interface{}, querystring map[string]string) (*http.Response, error) {
	ar := NewAPIRequest("PUT", endpoint, payload)
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ar, &responseStruct, querystring)
}

func (r *Requester) PostJSON(endpoint string, payload io.Reader, responseStruct interface{}, querystring map[string]string) (*http.Response, error) {
	ar := NewAPIRequest("POST", endpoint, payload)
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ar, &responseStruct, querystring)
}

func (r *Requester) PatchJSON(endpoint string, payload io.Reader, responseStruct interface{}, querystring map[string]string) (*http.Response, error) {
	ar := NewAPIRequest("PATCH", endpoint, payload)
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ar, &responseStruct, querystring)
}
