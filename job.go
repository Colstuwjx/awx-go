package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JobService struct {
	client *Client
}

type HostSummaryHost struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	HasActiveFailures   bool   `json:"has_active_failures"`
	HasInventorySources bool   `json:"has_inventory_sources"`
}

type HostSummaryJob struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Failed      bool    `json:"failed"`
	Elapsed     float64 `json:"elapsed"`
}

type HostSummaryFields struct {
	Host HostSummaryHost `json:"host"`
	Job  HostSummaryJob  `json:"job"`
}

type HostSummaryResult struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		Job  string `json:"job"`
		Host string `json:"host"`
	} `json:"related"`
	SummaryFields HostSummaryFields `json:"summary_fields"`
	Created       time.Time         `json:"created"`
	Modified      time.Time         `json:"modified"`
	Job           int               `json:"job"`
	Host          int               `json:"host"`
	HostName      string            `json:"host_name"`
	Changed       int               `json:"changed"`
	Dark          int               `json:"dark"`
	Failures      int               `json:"failures"`
	Ok            int               `json:"ok"`
	Processed     int               `json:"processed"`
	Skipped       int               `json:"skipped"`
	Failed        bool              `json:"failed"`
}

type HostSummariesResponse struct {
	Pagination
	Results []*HostSummaryResult `json:"results"`
}

type JobEvent struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		Job  string `json:"job"`
		Host string `json:"host"`
	} `json:"related"`
	SummaryFields struct {
		Role struct {
		} `json:"role"`
		Job struct {
			ID              int     `json:"id"`
			Name            string  `json:"name"`
			Description     string  `json:"description"`
			Status          string  `json:"status"`
			Failed          bool    `json:"failed"`
			Elapsed         float64 `json:"elapsed"`
			JobTemplateID   int     `json:"job_template_id"`
			JobTemplateName string  `json:"job_template_name"`
		} `json:"job"`
		Host struct {
			ID                  int    `json:"id"`
			Name                string `json:"name"`
			Description         string `json:"description"`
			HasActiveFailures   bool   `json:"has_active_failures"`
			HasInventorySources bool   `json:"has_inventory_sources"`
		} `json:"host"`
	} `json:"summary_fields"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Job          int       `json:"job"`
	Event        string    `json:"event"`
	Counter      int       `json:"counter"`
	EventDisplay string    `json:"event_display"`
	EventData    struct {
		PlayPattern string      `json:"play_pattern"`
		Play        string      `json:"play"`
		EventLoop   interface{} `json:"event_loop"`
		TaskArgs    string      `json:"task_args"`
		RemoteAddr  string      `json:"remote_addr"`
		Res         struct {
			AnsibleParsed bool     `json:"_ansible_parsed"`
			StderrLines   []string `json:"stderr_lines"`
			Changed       bool     `json:"changed"`
			End           string   `json:"end"`
			AnsibleNoLog  bool     `json:"_ansible_no_log"`
			Stdout        string   `json:"stdout"`
			Cmd           string   `json:"cmd"`
			Start         string   `json:"start"`
			Delta         string   `json:"delta"`
			Stderr        string   `json:"stderr"`
			Rc            int      `json:"rc"`
			Invocation    struct {
				ModuleArgs struct {
					Creates    interface{} `json:"creates"`
					Executable interface{} `json:"executable"`
					UsesShell  bool        `json:"_uses_shell"`
					RawParams  string      `json:"_raw_params"`
					Removes    interface{} `json:"removes"`
					Warn       bool        `json:"warn"`
					Chdir      string      `json:"chdir"`
					Stdin      interface{} `json:"stdin"`
				} `json:"module_args"`
			} `json:"invocation"`
			StdoutLines []string `json:"stdout_lines"`
			Warnings    []string `json:"warnings"`
		} `json:"res"`
		Pid          int    `json:"pid"`
		PlayUUID     string `json:"play_uuid"`
		TaskUUID     string `json:"task_uuid"`
		Task         string `json:"task"`
		PlaybookUUID string `json:"playbook_uuid"`
		Playbook     string `json:"playbook"`
		TaskAction   string `json:"task_action"`
		Host         string `json:"host"`
		Role         string `json:"role"`
		TaskPath     string `json:"task_path"`
	} `json:"event_data"`
	EventLevel int         `json:"event_level"`
	Failed     bool        `json:"failed"`
	Changed    bool        `json:"changed"`
	UUID       string      `json:"uuid"`
	ParentUUID string      `json:"parent_uuid"`
	Host       int         `json:"host"`
	HostName   string      `json:"host_name"`
	Parent     interface{} `json:"parent"`
	Playbook   string      `json:"playbook"`
	Play       string      `json:"play"`
	Task       string      `json:"task"`
	Role       string      `json:"role"`
	Stdout     string      `json:"stdout"`
	StartLine  int         `json:"start_line"`
	EndLine    int         `json:"end_line"`
	Verbosity  int         `json:"verbosity"`
}

type JobEventsResponse struct {
	Pagination
	Results []*JobEvent `json:"results"`
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

func (this *JobService) RelaunchJob(id int, data map[string]interface{}, params map[string]string) (*Job, *http.Response, error) {
	result := new(Job)
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

func (this *JobService) GetHostSummaries(id int, params map[string]string) ([]*HostSummaryResult, *http.Response, error) {
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

func (this *JobService) GetJobEvents(id int, params map[string]string) ([]*JobEvent, *http.Response, error) {
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
