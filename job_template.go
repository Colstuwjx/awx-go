package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JobTemplateService implements awx job template apis.
type JobTemplateService struct {
	client *Client
}

// ListJobTemplatesResponse represents `ListJobTemplates` endpoint response.
type ListJobTemplatesResponse struct {
	Pagination
	Results []*JobTemplate `json:"results"`
}

// ListJobTemplates shows a list of job templates.
func (jt *JobTemplateService) ListJobTemplates(params map[string]string) ([]*JobTemplate, *ListJobTemplatesResponse, error) {
	result := new(ListJobTemplatesResponse)
	endpoint := "/api/v2/job_templates/"
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// Launch lauchs a job with the job template.
func (jt *JobTemplateService) Launch(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d/launch/", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
