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
	mandatoryFields = []string{"name", "organization"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

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
func (i *InventoriesService) GetInventory(id int, params map[string]string) (*Inventory, error) {
	endpoint := fmt.Sprintf("/api/v2/inventories/%d", id)
	result := new(Inventory)
	resp, err := i.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil

}

// DeleteInventory delete an inventory from AWX
func (i *InventoriesService) DeleteInventory(id int) (*Inventory, error) {
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

// SyncInventorySourcesByInventoryID sends a sync request for a single inventory.
// Accepts only ID, which is the inventory id, not the inventory source id
func (i *InventoriesService) SyncInventorySourcesByInventoryID(id int) ([]*InventoryUpdate, error) {
	result := new([]*InventoryUpdate)

	endpoint := fmt.Sprintf("/api/v2/inventories/%d/update_inventory_sources/", id)
	resp, err := i.client.Requester.PostJSON(endpoint, nil, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return *result, nil
}

func (i *InventoriesService) CreateInventorySource(id int, data map[string]interface{}, params map[string]string) (*InventorySource, error) {
	mandatoryFields = []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(InventorySource)
	endpoint := fmt.Sprintf("/api/v2/inventories/%d/inventory_sources/", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if inventory_source exists and return proper error

	resp, err := i.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
