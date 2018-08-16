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

func (s *mockServer) InventoriesHandler(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == "POST":
		result := mockdata.MockedCreateInventoryResponse
		rw.Write(result)
		return
	case req.Method == "PATCH", req.Method == "PUT":
		result := mockdata.MockedUpdateInventoryResponse
		rw.Write(result)
		return
	case req.Method == "DELETE":
		result := mockdata.MockedDeleteInventoryResponse
		rw.Write(result)
		return
	case req.RequestURI == "/api/v2/inventories/1/":
		result := mockdata.MockedGetInventoryByIDResponse
		rw.Write(result)
		return
	default:
		result := mockdata.MockedListInventoriesResponse
		rw.Write(result)
	}
}

func (s *mockServer) JobTemplatesHandler(rw http.ResponseWriter, req *http.Request) {
	var (
		// FIXME: try another framework or implement dictionary tree, rather than using raw net/http.
		singleJobTemplateLaunchRoute = "/api/v2/job_templates/[0-9]+/launch/"
		defaultListJobTemplatesRoute = "/api/v2/job_templates/"
	)

	if matched, _ := regexp.MatchString(singleJobTemplateLaunchRoute, req.URL.String()); matched {
		result := mockdata.MockedLaunchJobTemplateResponse
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
}

func (s *mockServer) JobsHandler(rw http.ResponseWriter, req *http.Request) {
	var (
		// FIXME: try another framework or implement dictionary tree, rather than using raw net/http.
		singleJobGetJobEventsRoute     = "/api/v2/jobs/[0-9]+/job_events/"
		singleJobGetHostSummariesRoute = "/api/v2/jobs/[0-9]+/job_host_summaries/"
		singleJobCancelRoute           = "/api/v2/jobs/[0-9]+/cancel/"
		singleJobRelaunchRoute         = "/api/v2/jobs/[0-9]+/relaunch/"
		singleJobGetRoute              = "/api/v2/jobs/[0-9]+/"
	)

	if matched, _ := regexp.MatchString(singleJobGetJobEventsRoute, req.URL.String()); matched {
		result := mockdata.MockedJobEventsResponse
		rw.Write(result)
		return
	}

	if matched, _ := regexp.MatchString(singleJobGetHostSummariesRoute, req.URL.String()); matched {
		result := mockdata.MockedHostSummariesResponse
		rw.Write(result)
		return
	}

	if matched, _ := regexp.MatchString(singleJobCancelRoute, req.URL.String()); matched {
		result := mockdata.MockedCancelJobResponse
		rw.Write(result)
		return
	}

	if matched, _ := regexp.MatchString(singleJobRelaunchRoute, req.URL.String()); matched {
		result := mockdata.MockedLaunchJobTemplateResponse
		rw.Write(result)
		return
	}

	if matched, _ := regexp.MatchString(singleJobGetRoute, req.URL.String()); matched {
		result := mockdata.MockedGetJobResponse
		rw.Write(result)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte("404 - router not found!"))
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
	mux.Handle("/api/v2/inventories/", http.HandlerFunc(server.InventoriesHandler))
	mux.Handle("/api/v2/job_templates/", http.HandlerFunc(server.JobTemplatesHandler))
	mux.Handle("/api/v2/jobs/", http.HandlerFunc(server.JobsHandler))
	server.server.Handler = mux
	server.server.Addr = ":8080"
}

// Close mock server
func Close() error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	return server.server.Shutdown(ctx)
}
