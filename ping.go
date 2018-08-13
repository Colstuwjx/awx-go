package awx

// PingService implements awx ping apis.
type PingService struct {
	client *Client
}

// Ping do ping with awx servers.
func (p *PingService) Ping() (*Ping, error) {
	result := new(Ping)
	endpoint := "/api/v2/ping/"
	resp, err := p.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
