package mockserver

import (
	"context"
	"net/http"
	"regexp"
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

func (s *mockServer) JobTemplatesHandler(rw http.ResponseWriter, req *http.Request) {
	var (
		// FIXME: try another framework or implement dictionary tree, rather than using raw net/http.
		singleJobTemplateLaunchRoute = "/api/v2/job_templates/[0-9]+/launch/"
		defaultListJobTemplatesRoute = "/api/v2/job_templates/"
	)

	if matched, _ := regexp.MatchString(singleJobTemplateLaunchRoute, req.URL.String()); matched {
		result := mockdata.MockedaunchJobTemplateResponse
		rw.Write(result)
		return
	}

	if matched, _ := regexp.MatchString(defaultListJobTemplatesRoute, req.URL.String()); matched {
		result := mockdata.MockedListJobTemplatesResponse
		rw.Write(result)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte("404 - router not found!"))
	return
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
	mux.Handle("/api/v2/job_templates/", http.HandlerFunc(server.JobTemplatesHandler))
	server.server.Handler = mux
	server.server.Addr = ":8080"
}

// Close mock server
func Close() error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	return server.server.Shutdown(ctx)
}
