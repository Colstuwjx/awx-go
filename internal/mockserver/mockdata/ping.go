package mockdata

var (
	MockedPingResponse = []byte(`
{
    "instances": {
        "primary": "localhost",
        "secondaries": []
    },
    "ha": false,
    "role": "primary",
    "version": "2.2.2"
}`)
)
