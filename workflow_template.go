package awx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// WorkflowJobTemplateService implements awx job template apis.
type WorkflowJobTemplateService struct {
	client *Client
}

// ListWorkflowJobTemplatesResponse represents `ListWorkflowJobTemplates` endpoint response.
type ListWorkflowJobTemplatesResponse struct {
	Pagination
	Results []*WorkflowJobTemplate `json:"results"`
}

// ListWorkflowJobTemplates shows a list of job templates.
func (jt *WorkflowJobTemplateService) ListWorkflowJobTemplates(params map[string]string) ([]*WorkflowJobTemplate, *ListWorkflowJobTemplatesResponse, error) {
	result := new(ListWorkflowJobTemplatesResponse)
	endpoint := "/api/v2/workflow_job_templates/"
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
func (jt *WorkflowJobTemplateService) Launch(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("/api/v2/workflow_job_templates/%d/launch/", id)
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

	// in case invalid job id return
	if result.Job == 0 {
		return nil, errors.New("invalid job id 0")
	}

	return result, nil
}

// CreateWorkflowJobTemplate creates a job template
func (jt *WorkflowJobTemplateService) CreateWorkflowJobTemplate(data map[string]interface{}, params map[string]string) (*WorkflowJobTemplate, error) {
	result := new(WorkflowJobTemplate)
	mandatoryFields = []string{"name", "job_type", "inventory", "project"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	endpoint := "/api/v2/workflow_job_templates/"
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

// UpdateWorkflowJobTemplate updates a job template
func (jt *WorkflowJobTemplateService) UpdateWorkflowJobTemplate(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplate, error) {
	result := new(WorkflowJobTemplate)
	endpoint := fmt.Sprintf("/api/v2/workflow_job_templates/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteWorkflowJobTemplate deletes a job template
func (jt *WorkflowJobTemplateService) DeleteWorkflowJobTemplate(id int) (*WorkflowJobTemplate, error) {
	result := new(WorkflowJobTemplate)
	endpoint := fmt.Sprintf("/api/v2/workflow_job_templates/%d", id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
