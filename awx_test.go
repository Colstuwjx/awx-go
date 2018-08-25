package awx

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/Colstuwjx/awx-go/awxtesting/mockserver"
)

var (
	testAwxHost     = "http://127.0.0.1:8080"
	testAwxUserName = "admin"
	testAwxPasswd   = "password"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()

	// FIXME: fix codecov issue, then reopen it.
	// teardown()

	os.Exit(code)
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
