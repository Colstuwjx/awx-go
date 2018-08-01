package awx

type PingService struct {
	client *Client
}

func (this *PingService) Ping() (*Ping, error) {
	result := new(Ping)
	endpoint := "/api/v2/ping/"
	resp, err := this.client.Requester.GetJSON(endpoint, result, map[string]string{})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
