package awx

import (
	"bytes"
	"encoding/json"
)

// InventoriesService implements awx inventories apis.
type InventoriesService struct {
	client *Client
}

// ListInventoriesResponse represents `ListInventories` endpoint response.
type ListInventoriesResponse struct {
	Pagination
	Results []*Inventory `json:"results"`
}

// ListInventories shows list of awx inventories.
func (i *InventoriesService) ListInventories(params map[string]string) ([]*Inventory, *ListInventoriesResponse, error) {
	result := new(ListInventoriesResponse)
	endpoint := "/api/v2/inventories/"
	resp, err := i.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateInventory creates an awx inventory.
func (i *InventoriesService) CreateInventory(data map[string]interface{}, params map[string]string) (*Inventory, error) {
	result := new(Inventory)
	endpoint := "/api/v2/inventories/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := i.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateInventory update an awx inventory
func (this *InventoriesService) UpdateInventory(data map[string]interface{}, params map[string]string, id string) (*Inventory, error) {
        result := new(Inventory)
        endpoint := fmt.Sprintf("/api/v2/inventories/%s", id)
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
