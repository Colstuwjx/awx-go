package awx

import (
	"net/http"
	"time"
)

type PingService struct {
	client *Client
}

type InstanceGroup struct {
	Instances []string `json:"instances"`
	Capacity  int      `json:"capacity"`
	Name      string   `json:"name"`
}

type Instance struct {
	Node      string    `json:"node"`
	Heartbeat time.Time `json:"heartbeat"`
	Version   string    `json:"version"`
	Capacity  int       `json:"capacity"`
}

type Ping struct {
	Instances      []Instance      `json:"instances"`
	InstanceGroups []InstanceGroup `json:"instance_groups"`
	Ha             bool            `json:"ha"`
	Version        string          `json:"version"`
	ActiveNode     string          `json:"active_node"`
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
