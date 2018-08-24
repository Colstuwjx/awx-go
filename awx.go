package awx

import (
	"fmt"
	"net/http"
)

// AWX represents awx api endpoints with services, and using
// client to communicate with awx server.
type AWX struct {
	client *Client

	PingService        *PingService
	InventoriesService *InventoriesService
	JobService         *JobService
	JobTemplateService *JobTemplateService
	ProjectService     *ProjectService
	UserService        *UserService
	GroupService       *GroupService
}

// Client implement http client.
type Client struct {
	BaseURL   string
	Requester *Requester
}

// CheckResponse do http response check, and return err if not in [200, 300).
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("responsed with %d, resp: %v", resp.StatusCode, resp)
}

// NewAWX news an awx handler with basic auth support, you could customize the http
// transport by passing custom client.
func NewAWX(baseURL, userName, passwd string, client *http.Client) *AWX {
	r := &Requester{Base: baseURL, BasicAuth: &BasicAuth{Username: userName, Password: passwd}, Client: client}
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	awxClient := &Client{
		BaseURL:   baseURL,
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
		ProjectService: &ProjectService{
			client: awxClient,
		},
		UserService: &UserService{
			client: awxClient,
		},
		GroupService: &GroupService{
			client: awxClient,
		},
	}
}
