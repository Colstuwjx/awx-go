package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type JobTemplateService struct {
	client *Client
}

type ListJobTemplatesResponse struct {
	Pagination
	Results []*JobTemplate `json:"results"`
}

func (this *JobTemplateService) ListJobTemplates(params map[string]string) ([]*JobTemplate, *http.Response, error) {
	result := new(ListJobTemplatesResponse)
	endpoint := "/api/v2/job_templates/"
	resp, err := this.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, resp, err
	}

	return result.Results, resp, nil
}

func (this *JobTemplateService) Launch(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, *http.Response, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d/launch/", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	resp, err := this.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
