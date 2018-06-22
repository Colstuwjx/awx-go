package awx

import (
	"testing"
)

var (
	apiUrl        = "http://your_awx_host"
	userName      = "admin"
	passwd        = "password"
	jobTemplateId = 1
	jobId         = 1
)

func testAwxAPIRelaunchJob(t *testing.T, jobId int) int {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	relaunchResult, _, err := awx.RelaunchJob(jobId, map[string]interface{}{"hosts": "all"}, map[string]string{}) // hosts could be `all` or `failed`
	if err == nil {
		t.Log("relaunched job: ", relaunchResult.Job)
	} else {
		t.Fatalf("relaunch job err: %v", err)
	}

	return relaunchResult.Job
}

func testAwxAPICancelJob(t *testing.T, jobId int) bool {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	_, _, err := awx.CancelJob(jobId, map[string]interface{}{}, map[string]string{})
	if err == nil {
		t.Log("canceled job: ", jobId)
	} else {
		t.Log("cancel job err: ", err)
		return false
	}

	return true
}

func testAwxAPIHostSummaries(t *testing.T, jobId int) {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	hostSummaries, _, err := awx.GetJobHostSummaries(jobId, map[string]string{})
	if err == nil {
		t.Log("get job summary: ", hostSummaries)
	} else {
		t.Fatalf("get job %d host summary failed, err: %v", jobId, err)
	}
}

func TestAwxAPIInventories(t *testing.T) {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	inventoriesResp, _, err := awx.GetInventories(map[string]string{
		"name": "Demo Inventory",
	})

	if err != nil {
		t.Fatalf("get inventories err: %v", err)
	}

	t.Log(inventoriesResp)
	t.Log("Test passed!")
}

func TestAwxAPIPing(t *testing.T) {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	_, _, err := awx.Ping()
	if err != nil {
		t.Fatalf("ping awx err: %v", err)
	}

	t.Log("Test passed!")
}

func TestAwxAPIRunJobTemplate(t *testing.T) {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	result, _, err := awx.RunJobTemplate(jobTemplateId, map[string]interface{}{
		"inventory":  1,
		"limit":      "localhost",
		"extra_vars": "{\"version\": \"1234\"}",
	}, map[string]string{})
	if err != nil {
		t.Fatal(err)
	}

	if err == nil {
		t.Log("fired job: ", result.Job)

		testAwxAPICancelJob(t, result.Job)
		relaunchJobId := testAwxAPIRelaunchJob(t, result.Job)
		testAwxAPIHostSummaries(t, relaunchJobId)
	} else {
		t.Fatalf("run job template err: %s", err)
	}
}

func TestAwxAPIGetJob(t *testing.T) {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	result, _, err := awx.GetJob(31, map[string]string{})
	if err != nil {
		t.Fatal(err)
	}

	if err == nil {
		t.Log("Job ", result.Name, " status: ", result.Status)
		t.Log("Test Passed!")
	} else {
		t.Fatalf("response err: %s", err)
	}
}

func TestAwxAPIGetJobEventsResult(t *testing.T) {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	result, _, err := awx.GetJobEvents(31, map[string]string{
		"order_by":  "start_line",
		"page_size": "1000000",
	})
	if err != nil {
		t.Fatal(err)
	}

	if err == nil {
		hostStatuses := make(map[string][]*JobEventResponseResult)
		output := ""

		for _, r := range result.Results {
			output = output + r.Stdout
			if r.HostName != "" {
				hostStatuses[r.HostName] = append(hostStatuses[r.HostName], r)
			}
		}

		for host, hs := range hostStatuses {
			t.Log("Host: ", host)
			t.Log("Status: ", hs)
		}

		t.Log(output)
		t.Log("Test Passed!")
	} else {
		t.Fatalf("response err: %s", err)
	}
}

func TestAwxAPIGetJobEventsStat(t *testing.T) {
	awx := NewAWX(apiUrl, userName, passwd, nil)
	result, _, err := awx.GetJobEvents(jobId, map[string]string{
		"event": "playbook_on_stats",
	})
	if err != nil {
		t.Fatal(err)
	}

	if err == nil {
		output := ""
		for _, r := range result.Results {
			output = output + r.Stdout
		}

		t.Log(output)
		t.Log("Test Passed!")
	} else {
		t.Fatalf("response err: %s", err)
	}
}
