package mockdata

var (
	// MockedPingResponse mocked `Ping` endpoint response
	MockedPingResponse = []byte(`
{
    "instance_groups": [
        {
            "instances": [
                "awx"
            ],
            "capacity": 138,
            "name": "tower"
        }
    ],
    "instances": [
        {
            "node": "awx",
            "heartbeat": "2018-06-25T03:14:34.688447Z",
            "version": "1.0.6.5",
            "capacity": 138
        }
    ],
    "ha": false,
    "version": "1.0.6.5",
    "active_node": "awx"
}`)
)
