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

func (s *mockServer) ListInventoriesHandler(rw http.ResponseWriter, req *http.Request) {
	result := mockdata.MockedListInventoriesResponse
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
	mux.Handle("/api/v2/ping/", http.HandlerFunc(server.PingHandler))
	mux.Handle("/api/v2/inventories/", http.HandlerFunc(server.ListInventoriesHandler))
	server.server.Handler = mux
	server.server.Addr = ":8080"
}

// Close mock server
func Close() error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	return server.server.Shutdown(ctx)
}
