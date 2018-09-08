package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	mandatoryFields = []string{"name", "organization", "scm_type"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

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

// UpdateProject update an awx Project.
func (p *ProjectService) UpdateProject(id int, data map[string]interface{}, params map[string]string) (*Project, error) {
	result := new(Project)
	endpoint := fmt.Sprintf("/api/v2/projects/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteProject delete an awx Project.
func (p *ProjectService) DeleteProject(id int) (*Project, error) {
	result := new(Project)
	endpoint := fmt.Sprintf("/api/v2/projects/%d", id)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CancelUpdate cancel of awx projects update.
func (p *ProjectService) CancelUpdate(id int) (*ProjectUpdateCancel, error) {
	result := new(ProjectUpdateCancel)
	endpoint := fmt.Sprintf("/api/v2/project_updates/%d/cancel", id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUpdate get of awx projects update.
func (p *ProjectService) GetUpdate(id int) (*Job, error) {
	result := new(Job)
	endpoint := fmt.Sprintf("/api/v2/project_updates/%d", id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}
