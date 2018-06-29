package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	JobStatusNew        = "new"
	JobStatusPending    = "pending"
	JobStatusWaiting    = "waiting"
	JobStatusRunning    = "running"
	JobStatusSuccessful = "successful"
	JobStatusFailed     = "failed"
	JobStatusError      = "error"
	JobStatusCanceled   = "canceled"
)

type JobService struct {
	client *Client
}

type HostSummariesResponse struct {
	Pagination
	Results []HostSummary `json:"results"`
}

type JobEventsResponse struct {
	Pagination
	Results []JobEvent `json:"results"`
}

type CancelJobResponse struct {
	Detail string `json:"detail"`
}

func (this *JobService) GetJob(id int, params map[string]string) (*Job, *http.Response, error) {
	result := new(Job)
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/", id)
	resp, err := this.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

func (this *JobService) CancelJob(id int, data map[string]interface{}, params map[string]string) (*CancelJobResponse, *http.Response, error) {
	result := new(CancelJobResponse)
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/cancel/", id)
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

func (this *JobService) RelaunchJob(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, *http.Response, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/relaunch/", id)
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

func (this *JobService) GetHostSummaries(id int, params map[string]string) ([]HostSummary, *http.Response, error) {
	result := new(HostSummariesResponse)
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/job_host_summaries/", id)
	resp, err := this.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, resp, err
	}

	return result.Results, resp, nil
}

func (this *JobService) GetJobEvents(id int, params map[string]string) ([]JobEvent, *http.Response, error) {
	result := new(JobEventsResponse)
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/job_events/", id)
	resp, err := this.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, resp, err
	}

	return result.Results, resp, nil
}
