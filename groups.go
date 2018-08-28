package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GroupService implements awx Groups apis.
type GroupService struct {
	client *Client
}

// ListGroupsResponse represents `ListGroups` endpoint response.
type ListGroupsResponse struct {
	Pagination
	Results []*Group `json:"results"`
}

// ListGroups shows list of awx Groups.
func (g *GroupService) ListGroups(params map[string]string) ([]*Group, *ListGroupsResponse, error) {
	result := new(ListGroupsResponse)
	endpoint := "/api/v2/groups/"
	resp, err := g.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateGroup creates an awx Group.
func (g *GroupService) CreateGroup(data map[string]interface{}, params map[string]string) (*Group, error) {
	mandatoryFields = []string{"name", "inventory"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Group)
	endpoint := "/api/v2/groups/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if Group exists and return proper error

	resp, err := g.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
