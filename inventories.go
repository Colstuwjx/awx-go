package awx

import (
	"bytes"
	"encoding/json"
)

type InventoriesService struct {
	client *Client
}

type ListInventoriesResponse struct {
	Pagination
	Results []*Inventory `json:"results"`
}

func (this *InventoriesService) ListInventories(params map[string]string) ([]*Inventory, *ListInventoriesResponse, error) {
	result := new(ListInventoriesResponse)
	endpoint := "/api/v2/inventories/"
	resp, err := this.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (this *InventoriesService) CreateInventory(data map[string]interface{}, params map[string]string) (*Inventory, error) {
	result := new(Inventory)
	endpoint := "/api/v2/inventories/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := this.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
