package awx

import (
	"fmt"
	"net/http"
)

type AWX struct {
	client *Client

	PingService        *PingService
	InventoriesService *InventoriesService
	JobService         *JobService
	JobTemplateService *JobTemplateService
}

type Client struct {
	BaseUrl   string
	Requester *Requester
}

type Pagination struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
}

func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	} else {
		return fmt.Errorf("responsed with %d, resp: %v", resp.StatusCode, resp)
	}
}

func NewAWX(baseUrl, userName, passwd string, client *http.Client) *AWX {
	r := &Requester{Base: baseUrl, BasicAuth: &BasicAuth{Username: userName, Password: passwd}, Client: client}
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	awxClient := &Client{
		BaseUrl:   baseUrl,
		Requester: r,
	}

	return &AWX{
		client: awxClient,
		PingService: &PingService{
			client: awxClient,
		},
		InventoriesService: &InventoriesService{
			client: awxClient,
		},
		JobService: &JobService{
			client: awxClient,
		},
		JobTemplateService: &JobTemplateService{
			client: awxClient,
		},
	}
}
