package mockserver

import (
	"context"
	"net/http"
	"time"

	"github.com/Colstuwjx/awx-go/internal/mockserver/mockdata"
)

type mockServer struct {
	// TODO: add auth mock

	server http.Server
}

func (s *mockServer) PingHandler(rw http.ResponseWriter, req *http.Request) {
	result := mockdata.MockedPingResponse
	rw.Write(result)
}

var server *mockServer

// Run mock server
func Run() error {
	initServer()
	return server.server.ListenAndServe()
}

func initServer() {
	server = &mockServer{}
	mux := http.NewServeMux()
	mux.Handle("/api/v1/ping/", http.HandlerFunc(server.PingHandler))
	server.server.Handler = mux
	server.server.Addr = ":8080"
}

// Close mock server
func Close() error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	return server.server.Shutdown(ctx)
}
