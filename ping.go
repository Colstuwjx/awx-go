package awx

import (
	"net/http"
)

type PingService struct {
	client *Client
}

func (this *PingService) Ping() (*Ping, *http.Response, error) {
	result := new(Ping)
	endpoint := "/api/v2/ping/"
	resp, err := this.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
