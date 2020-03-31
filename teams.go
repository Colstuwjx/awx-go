package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TeamService implements awx Teams apis.
type TeamService struct {
	client *Client
}

// ListTeamsResponse represents `ListTeams` endpoint response.
type ListTeamsResponse struct {
	Pagination
	Results []*Team `json:"results"`
}

// ListTeams shows list of awx Teams.
func (t *TeamService) ListTeams(params map[string]string) ([]*Team, *ListTeamsResponse, error) {
	result := new(ListTeamsResponse)
	endpoint := "/api/v2/teams/"
	resp, err := t.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateTeam creates an awx Team.
func (t *TeamService) CreateTeam(data map[string]interface{}, params map[string]string) (*Team, error) {
	mandatoryFields = []string{"name", "organization"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Team)
	endpoint := "/api/v2/teams/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if Team exists and return proper error

	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTeam update an awx user.
func (t *TeamService) UpdateTeam(id int, data map[string]interface{}, params map[string]string) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("/api/v2/teams/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteTeam delete an awx Team.
func (t *TeamService) DeleteTeam(id int) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("/api/v2/teams/%d", id)

	resp, err := t.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// GrantRole grant the provided role to the AWX Team
func (t *TeamService) GrantRole(id int, roleID int) error {
	result := new(Team)
	endpoint := fmt.Sprintf("/api/v2/teams/%d/roles/", id)
	data := map[string]interface{}{
		"id": roleID,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return err
	}

	if err := CheckResponse(resp); err != nil {
		return err
	}

	return nil
}

// RevokeRole revoke the provided role to the AWX Team
func (t *TeamService) RevokeRole(id int, roleID int) error {
	result := new(Team)
	endpoint := fmt.Sprintf("/api/v2/teams/%d/roles/", id)
	data := map[string]interface{}{
		"id":           roleID,
		"disassociate": "true",
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return err
	}

	if err := CheckResponse(resp); err != nil {
		return err
	}

	return nil
}
