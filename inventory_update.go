package awx

import "fmt"

// InventoryUpdateService implements awx inventory_update apis.
type InventoryUpdatesService struct {
	client *Client
}

// GetInventoryUpdate retrives the inventory_update information from its ID or Name
func (i *InventoryUpdatesService) GetInventoryUpdate(id int, params map[string]string) (*InventoryUpdate, error) {
	endpoint := fmt.Sprintf("/api/v2/inventory_updates/%d/", id)
	result := new(InventoryUpdate)

	resp, err := i.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
