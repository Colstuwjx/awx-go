package awx

import (
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	var (
		expectPingResponse = &Ping{
			Instances: []Instance{
				{
					Node: "awx",
					Heartbeat: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2018-06-25T03:14:34.688447Z")
						return t
					}(),
					Version:  "1.0.6.5",
					Capacity: 138,
				},
			},
			InstanceGroups: []InstanceGroup{
				{
					Instances: []string{"awx"},
					Capacity:  138,
					Name:      "tower",
				},
			},
			Ha:         false,
			Version:    "1.0.6.5",
			ActiveNode: "awx",
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.PingService.Ping()

	if err != nil {
		t.Fatalf("Ping err: %s", err)
	} else {
		checkAPICallResult(t, *expectPingResponse, *result)
		t.Log("Ping passed!")
	}
}
