package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CredentialService implements awx Credentials apis.
type CredentialService struct {
	client *Client
}

// ListCredentialsResponse represents `ListCredentials` endpoint response.
type ListCredentialsResponse struct {
	Pagination
	Results []*Credential `json:"results"`
}

// ListCredentials shows list of awx Credentials.
func (t *CredentialService) ListCredentials(params map[string]string) ([]*Credential, *ListCredentialsResponse, error) {
	result := new(ListCredentialsResponse)
	endpoint := "/api/v2/credentials/"
	resp, err := t.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateCredential creates an awx Credential.
func (t *CredentialService) CreateCredential(data map[string]interface{}, params map[string]string) (*Credential, error) {
	mandatoryFields = []string{"name", "credential_type"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Credential)
	endpoint := "/api/v2/credentials/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if Credential exists and return proper error

	resp, err := t.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateCredential update an awx user.
func (t *CredentialService) UpdateCredential(id int, data map[string]interface{}, params map[string]string) (*Credential, error) {
	result := new(Credential)
	endpoint := fmt.Sprintf("/api/v2/credentials/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Requester.PutJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteCredential delete an awx Credential.
func (t *CredentialService) DeleteCredential(id int) (*Credential, error) {
	result := new(Credential)
	endpoint := fmt.Sprintf("/api/v2/credentials/%d", id)

	resp, err := t.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
