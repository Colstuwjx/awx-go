package awx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/twinj/uuid"
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

// Launch launches a job with the job template.
func (jt *JobTemplateService) Launch(id int, data *JobLaunchOpts, params map[string]string) (*JobLaunch, error) {
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

// CreateJobTemplateCallBack executes a PATCH HTTP Request to create the callback url and the generated host_config_key
func (jt *JobTemplateService) CreateJobTemplateCallBack(template *JobTemplate) (*JobTemplate, error) {
	if template.ID == 0 {
		return nil, fmt.Errorf("Job template ID must be passed")
	}

	endpoint := "/api/v2/job_templates/" + strconv.Itoa(template.ID)
	template.AllowCallbacks = true
	template.HostConfigKey = uuid.NewV4().String()

	jsonPayload, err := json.Marshal(template)

	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PatchJSON(endpoint, bytes.NewReader(jsonPayload), template, map[string]string{})

	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return template, nil
}

// CreateJobTemplate creates a job template
func (jt *JobTemplateService) CreateJobTemplate(data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)
	mandatoryFields = []string{"name", "job_type", "inventory", "project", "playbook"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	endpoint := "/api/v2/job_templates/"
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

	if result.AllowCallbacks {
		return jt.CreateJobTemplateCallBack(result)
	}

	return result, nil
}

// UpdateJobTemplate updates a job template
func (jt *JobTemplateService) UpdateJobTemplate(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d", id)
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

// DeleteJobTemplate deletes a job template
func (jt *JobTemplateService) DeleteJobTemplate(id int) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d", id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// GetJobTemplate gets a job template
func (jt *JobTemplateService) GetJobTemplate(id int) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d", id)

	resp, err := jt.client.Requester.Get(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func (jt *JobTemplateService) AddJobTemplateCredential(jobTemplateID int, credID int) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d/credentials/", jobTemplateID)

	payload := map[string]int{
		"id": credID,
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(jsonPayload), result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func (jt *JobTemplateService) GetSurveySpec(jobTemplate *JobTemplate) ([]byte, error) {
	endpoint := jobTemplate.Related.SurveySpec
	spec := make(map[string]interface{})
	resp, err := jt.client.Requester.Get(endpoint, spec, map[string]string{})

	if err != nil {
		return nil, err
	}

	if err = CheckResponse(resp); err != nil {
		return nil, err
	}

	return json.Marshal(spec)
}
