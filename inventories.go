package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	// Add check if inventory exists and return proper error

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
func (i *InventoriesService) UpdateInventory(id int, data map[string]interface{}, params map[string]string) (*Inventory, error) {
	result := new(Inventory)
	endpoint := fmt.Sprintf("/api/v2/inventories/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := i.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// GetInventory retrives the inventory information from its ID or Name
func (i *InventoriesService) GetInventory(params map[string]string) (*Inventory, error) {
	endpoint := "/api/v2/inventories/"
	if _, ok := params["id"]; ok {
		endpoint = fmt.Sprintf("%s/%s", endpoint, params["id"])
		result := new(Inventory)
		resp, err := i.client.Requester.GetJSON(endpoint, result, map[string]string{})
		if err != nil {
			return nil, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, err
		}

		return result, nil

	} else if _, ok := params["name"]; ok {
		result := new(ListInventoriesResponse)
		resp, err := i.client.Requester.GetJSON(endpoint, result, params)
		if err != nil {
			return nil, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, err
		}

		if len(result.Results) >= 1 {
			return result.Results[0], nil
		}
		return nil, fmt.Errorf("Inventory not found")
	} else {
		return nil, fmt.Errorf("params must be name or id")
	}

}

// Delete an inventory from AWX
func (i *InventoriesService) Delete(id int) (*Inventory, error) {
	result := new(Inventory)
	endpoint := fmt.Sprintf("/api/v2/inventories/%d", id)

	resp, err := i.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
