package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AWX struct {
	baseUrl   string
	requester *Requester
}

func NewAWX(baseUrl, userName, passwd string, client *http.Client) *AWX {
	r := &Requester{Base: baseUrl, BasicAuth: &BasicAuth{Username: userName, Password: passwd}, Client: client}
	if r.Client == nil {
		r.Client = http.DefaultClient
	}

	return &AWX{
		baseUrl, r,
	}
}

func (a *AWX) Ping() (*PingResponse, *http.Response, error) {
	var result *PingResponse
	endpoint := "/api/v1/ping/"
	resp, err := a.requester.GetJSON(endpoint, &result, map[string]string{})
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d, resp: %v", resp.StatusCode, resp)
	}

	return result, resp, nil
}

func (a *AWX) GetInventories(params map[string]string) (*InventoryResponse, *http.Response, error) {
	var result *InventoryResponse
	endpoint := "/api/v2/inventories/"
	resp, err := a.requester.GetJSON(endpoint, &result, params)
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d", resp.StatusCode)
	}

	return result, resp, nil
}

func (a *AWX) RunJobTemplate(jobTpltId int, data map[string]interface{}, params map[string]string) (*RunJobTemplateResponse, *http.Response, error) {
	var result = new(RunJobTemplateResponse)
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d/launch/", jobTpltId)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	resp, err := a.requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d", resp.StatusCode)
	}

	return result, resp, nil
}

func (a *AWX) RelaunchJob(jobId int, data map[string]interface{}, params map[string]string) (*RunJobTemplateResponse, *http.Response, error) {
	var result = new(RunJobTemplateResponse)
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/relaunch/", jobId)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	resp, err := a.requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d", resp.StatusCode)
	}

	return result, resp, nil
}

func (a *AWX) CancelJob(jobId int, data map[string]interface{}, params map[string]string) (*CancelJobResponse, *http.Response, error) {
	var result = new(CancelJobResponse)
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/cancel/", jobId)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	resp, err := a.requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d", resp.StatusCode)
	}

	return result, resp, nil
}

func (a *AWX) GetJob(jobId int, params map[string]string) (*JobResponse, *http.Response, error) {
	var result *JobResponse
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/", jobId)
	resp, err := a.requester.GetJSON(endpoint, &result, params)
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d", resp.StatusCode)
	}

	return result, resp, nil
}

func (a *AWX) GetJobHostSummaries(jobId int, params map[string]string) (*HostSummaryResponse, *http.Response, error) {
	var result *HostSummaryResponse
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/job_host_summaries/", jobId)
	resp, err := a.requester.GetJSON(endpoint, &result, params)
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d", resp.StatusCode)
	}

	return result, resp, nil
}

func (a *AWX) GetJobEvents(jobId int, params map[string]string) (*JobEventResponse, *http.Response, error) {
	var result *JobEventResponse
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/job_events/", jobId)
	resp, err := a.requester.GetJSON(endpoint, &result, params)
	if err != nil {
		return result, resp, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return result, resp, fmt.Errorf("responsed with %d", resp.StatusCode)
	}

	return result, resp, nil
}
