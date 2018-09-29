package awx

import (
	"fmt"
)

// ProjectUpdatesService implements awx projects apis.
type ProjectUpdatesService struct {
	client *Client
}

// ProjectUpdateCancel cancel of awx projects update.
func (p *ProjectUpdatesService) ProjectUpdateCancel(id int) (*ProjectUpdateCancel, error) {
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

// ProjectUpdateGet get of awx projects update.
func (p *ProjectUpdatesService) ProjectUpdateGet(id int) (*Job, error) {
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
