package awx

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/Colstuwjx/awx-go/internal/mockserver"
)

func TestMain(m *testing.M) {
	setup()
	defer teardown()

	m.Run()
}

func setup() {
	go func() {
		if err := mockserver.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for mock server to run
	time.Sleep(time.Millisecond * 10)
}

func teardown() {
	mockserver.Close()
}

func TestPing(t *testing.T) {
	var (
		expectPingResponse = &Ping{
			Instances: Instances{
				Primary:     "localhost",
				Secondaries: []string{},
			},
			Ha:      false,
			Role:    "primary",
			Version: "2.2.2",
		}
	)

	awx := NewAWX("http://127.0.0.1:8080", "", "", nil)
	result, _, err := awx.PingService.Ping()
	if err != nil {
		log.Fatalf("Ping err: %s", err)
	}

	if !reflect.DeepEqual(*result, *expectPingResponse) {
		log.Fatalf("Ping resp not as expected, expected: %v, resp result: %v", *expectPingResponse, *result)
	}

	log.Println("Ping passed!")
}
