package awx

import (
	"net/http"
)

type PingService struct {
	client *Client
}

type Instances struct {
	Primary     string   `json:"primary"`
	Secondaries []string `json:"secondaries"`
}

type Ping struct {
	Instances Instances `json:"instances"`
	Ha        bool      `json:"ha"`
	Role      string    `json:"role"`
	Version   string    `json:"version"`
}

func (this *PingService) Ping() (*Ping, *http.Response, error) {
	result := new(Ping)
	endpoint := "/api/v1/ping/"
	resp, err := this.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return result, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return result, resp, err
	}

	return result, resp, nil
}
