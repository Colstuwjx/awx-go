package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// HostService implements awx Hosts apis.
type HostService struct {
	client *Client
}

// ListHostsResponse represents `ListHosts` endpoint response.
type ListHostsResponse struct {
	Pagination
	Results []*Host `json:"results"`
}

// ListHosts shows list of awx Hosts.
func (h *HostService) ListHosts(params map[string]string) ([]*Host, *ListHostsResponse, error) {
	result := new(ListHostsResponse)
	endpoint := "/api/v2/hosts/"
	resp, err := h.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateHost creates an awx Host.
func (h *HostService) CreateHost(data map[string]interface{}, params map[string]string) (*Host, error) {
	mandatoryFields = []string{"name", "inventory"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Host)
	endpoint := "/api/v2/hosts/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if Host exists and return proper error

	resp, err := h.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateHost update an awx Host
func (h *HostService) UpdateHost(id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("/api/v2/hosts/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateGroup update an awx Host
func (h *HostService) AssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("/api/v2/hosts/%d/groups/", id)
	data["associate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DisAssociateGroup update an awx Host
func (h *HostService) DisAssociateGroup(id int, data map[string]interface{}, params map[string]string) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("/api/v2/hosts/%d/groups/", id)
	data["disassociate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteHost delete an awx Host.
func (h *HostService) DeleteHost(id int) (*Host, error) {
	result := new(Host)
	endpoint := fmt.Sprintf("/api/v2/hosts/%d", id)

	resp, err := h.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
