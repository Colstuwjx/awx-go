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
