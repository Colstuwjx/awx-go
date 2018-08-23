package awx

import (
	"bytes"
	"encoding/json"
)

// ProjectService implements awx projects apis.
type ProjectService struct {
	client *Client
}

// ListProjectsResponse represents `ListProjects` endpoint response.
type ListProjectsResponse struct {
	Pagination
	Results []*Project `json:"results"`
}

// ListProjects shows list of awx projects.
func (p *ProjectService) ListProjects(params map[string]string) ([]*Project, *ListProjectsResponse, error) {
	result := new(ListProjectsResponse)
	endpoint := "/api/v2/projects/"
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateProject creates an awx project.
func (p *ProjectService) CreateProject(data map[string]interface{}, params map[string]string) (*Project, error) {
	result := new(Project)
	endpoint := "/api/v2/projects/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if project exists and return proper error

	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
