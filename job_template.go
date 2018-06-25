package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JobTemplateService struct {
	client *Client
}

type JobTemplate struct {
	ID                    int         `json:"id"`
	Type                  string      `json:"type"`
	URL                   string      `json:"url"`
	Related               *Related    `json:"related"`
	SummaryFields         *Summary    `json:"summary_fields"`
	Created               time.Time   `json:"created"`
	Modified              time.Time   `json:"modified"`
	Name                  string      `json:"name"`
	Description           string      `json:"description"`
	JobType               string      `json:"job_type"`
	Inventory             int         `json:"inventory"`
	Project               int         `json:"project"`
	Playbook              string      `json:"playbook"`
	Forks                 int         `json:"forks"`
	Limit                 string      `json:"limit"`
	Verbosity             int         `json:"verbosity"`
	ExtraVars             string      `json:"extra_vars"`
	JobTags               string      `json:"job_tags"`
	ForceHandlers         bool        `json:"force_handlers"`
	SkipTags              string      `json:"skip_tags"`
	StartAtTask           string      `json:"start_at_task"`
	Timeout               int         `json:"timeout"`
	UseFactCache          bool        `json:"use_fact_cache"`
	LastJobRun            interface{} `json:"last_job_run"`
	LastJobFailed         bool        `json:"last_job_failed"`
	NextJobRun            interface{} `json:"next_job_run"`
	Status                string      `json:"status"`
	HostConfigKey         string      `json:"host_config_key"`
	AskDiffModeOnLaunch   bool        `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch  bool        `json:"ask_variables_on_launch"`
	AskLimitOnLaunch      bool        `json:"ask_limit_on_launch"`
	AskTagsOnLaunch       bool        `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch   bool        `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch    bool        `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch  bool        `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch  bool        `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch bool        `json:"ask_credential_on_launch"`
	SurveyEnabled         bool        `json:"survey_enabled"`
	BecomeEnabled         bool        `json:"become_enabled"`
	DiffMode              bool        `json:"diff_mode"`
	AllowSimultaneous     bool        `json:"allow_simultaneous"`
	CustomVirtualenv      interface{} `json:"custom_virtualenv"`
	Credential            int         `json:"credential"`
	VaultCredential       interface{} `json:"vault_credential"`
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

func (this *JobTemplateService) Launch(id int, data map[string]interface{}, params map[string]string) (*Job, *http.Response, error) {
	result := new(Job)
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
