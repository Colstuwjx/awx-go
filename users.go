package awx

import (
	"bytes"
	"encoding/json"
)

// UserService implements awx Users apis.
type UserService struct {
	client *Client
}

// ListUsersResponse represents `ListUsers` endpoint response.
type ListUsersResponse struct {
	Pagination
	Results []*User `json:"results"`
}

// ListUsers shows list of awx Users.
func (i *UserService) ListUsers(params map[string]string) ([]*User, *ListUsersResponse, error) {
	result := new(ListUsersResponse)
	endpoint := "/api/v2/users/"
	resp, err := i.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateUser creates an awx User.
func (i *UserService) CreateUser(data map[string]interface{}, params map[string]string) (*CreateUser, error) {
	result := new(CreateUser)
	endpoint := "/api/v2/users/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if User exists and return proper error

	resp, err := i.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
