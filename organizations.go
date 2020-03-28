package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// OrganizationService implements awx Organizations apis.
type OrganizationService struct {
	client *Client
}

// ListOrganizationsResponse represents `ListOrganizations` endpoint response.
type ListOrganizationsResponse struct {
	Pagination
	Results []*Organization `json:"results"`
}

// ListOrganizations shows list of awx Organizations.
func (t *OrganizationService) ListOrganizations(params map[string]string) ([]*Organization, *ListOrganizationsResponse, error) {
	result := new(ListOrganizationsResponse)
	endpoint := "/api/v2/organizations/"
	resp, err := t.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateOrganization creates an awx Organization.
func (t *OrganizationService) CreateOrganization(data map[string]interface{}, params map[string]string) (*Organization, error) {
	mandatoryFields = []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Organization)
	endpoint := "/api/v2/organizations/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if Organization exists and return proper error

	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateOrganization update an awx user.
func (t *OrganizationService) UpdateOrganization(id int, data map[string]interface{}, params map[string]string) (*Organization, error) {
	result := new(Organization)
	endpoint := fmt.Sprintf("/api/v2/organizations/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteOrganization delete an awx Organization.
func (t *OrganizationService) DeleteOrganization(id int) (*Organization, error) {
	result := new(Organization)
	endpoint := fmt.Sprintf("/api/v2/organizations/%d", id)

	resp, err := t.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
