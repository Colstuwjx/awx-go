package awx

import (
	"testing"
	"time"
)

func TestGetJob(t *testing.T) {
	var (
		expectGetJobResponse = &Job{
			ID:   301,
			Type: "job",
			URL:  "/api/v2/jobs/301/",
			Related: &Related{
				CreatedBy:          "/api/v2/users/1/",
				Labels:             "/api/v2/jobs/301/labels/",
				Inventory:          "/api/v2/inventories/1/",
				Project:            "/api/v2/projects/6/",
				ExtraCredentials:   "/api/v2/jobs/301/extra_credentials/",
				Credentials:        "/api/v2/jobs/301/credentials/",
				UnifiedJobTemplate: "/api/v2/job_templates/8/",
				Stdout:             "/api/v2/jobs/301/stdout/",
				Notifications:      "/api/v2/jobs/301/notifications/",
				JobHostSummaries:   "/api/v2/jobs/301/job_host_summaries/",
				JobEvents:          "/api/v2/jobs/301/job_events/",
				ActivityStream:     "/api/v2/jobs/301/activity_stream/",
				JobTemplate:        "/api/v2/job_templates/8/",
				Cancel:             "/api/v2/jobs/301/cancel/",
				ProjectUpdate:      "/api/v2/project_updates/303/",
				CreateSchedule:     "/api/v2/jobs/301/create_schedule/",
				Relaunch:           "/api/v2/jobs/301/relaunch/",
			},
			SummaryFields: &Summary{
				InstanceGroup: &InstanceGroupSummary{
					ID:   1,
					Name: "tower",
				},

				JobTemplate: &JobTemplateSummary{
					ID:          8,
					Name:        "Hello-world",
					Description: "",
				},

				Inventory: &Inventory{
					ID:                           1,
					Name:                         "Demo Inventory",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   2,
					HostsWithActiveFailures:      0,
					TotalGroups:                  0,
					GroupsWithActiveFailures:     0,
					HasInventorySources:          false,
					TotalInventorySources:        0,
					InventorySourcesWithFailures: 0,
					OrganizationID:               1,
					Kind:                         "",
				},

				ProjectUpdate: &ProjectUpdate{
					ID:          303,
					Name:        "DeployBook",
					Description: "",
					Status:      "successful",
					Failed:      false,
				},

				UnifiedJobTemplate: &UnifiedJobTemplate{
					ID:             8,
					Description:    "",
					Name:           "Hello-world",
					UnifiedJobType: "job",
				},

				Project: &Project{
					ID:          6,
					Name:        "DeployBook",
					Description: "",
					Status:      "successful",
					ScmType:     "git",
				},

				CreatedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},

				UserCapabilities: &UserCapabilities{
					Start:  true,
					Delete: true,
				},

				Labels: &Labels{
					Count:   0,
					Results: []interface{}{},
				},

				ExtraCredentials: []interface{}{},
				Credentials:      []Credential{},
			},

			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:43.046077Z")
				return t
			}(),

			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:49.132905Z")
				return t
			}(),

			Name:               "Hello-world",
			Description:        "",
			JobType:            "run",
			Inventory:          1,
			Project:            6,
			Playbook:           "hello-world.yml",
			Forks:              0,
			Limit:              "localhost",
			Verbosity:          0,
			ExtraVars:          "{}",
			JobTags:            "",
			ForceHandlers:      false,
			SkipTags:           "",
			StartAtTask:        "",
			Timeout:            0,
			UseFactCache:       false,
			UnifiedJobTemplate: 8,
			LaunchType:         "manual",
			Status:             "canceled",
			Failed:             true,
			Started: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:47.060276Z")
				return t
			}(),
			Finished: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:54.740342Z")
				return t
			}(),
			Elapsed:                 7.68,
			JobArgs:                 "[\"ansible-playbook\", \"-i\", \"/tmp/awx_301_3kuP6Z/tmpupf88B\", \"-u\", \"root\", \"-l\", \"localhost\", \"-e\", \"@/tmp/awx_301_3kuP6Z/tmpjlw9Dp\", \"hello-world.yml\"]",
			JobCwd:                  "/var/lib/awx/projects/_6__deploybook",
			JobEnv:                  map[string]string{},
			JobExplanation:          "",
			ExecutionNode:           "awx",
			ResultTraceback:         "",
			EventProcessingFinished: true,
			JobTemplate:             8,
			PasswordsNeededToStart:  []interface{}{},
			AskDiffModeOnLaunch:     false,
			AskVariablesOnLaunch:    true,
			AskLimitOnLaunch:        true,
			AskTagsOnLaunch:         false,
			AskSkipTagsOnLaunch:     false,
			AskJobTypeOnLaunch:      false,
			AskVerbosityOnLaunch:    false,
			AskInventoryOnLaunch:    true,
			AskCredentialOnLaunch:   false,
			AllowSimultaneous:       false,
			Artifacts:               map[string]string{},
			ScmRevision:             "ae9cc6b0fee9aea55e10a4a325c9ae3d97c42f1a",
			InstanceGroup:           1,
			DiffMode:                false,
			Credential:              nil,
			VaultCredential:         nil,
		}

		testJobID = 301
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.JobService.GetJob(testJobID, map[string]string{})

	if err != nil {
		t.Fatalf("GetJob err: %s", err)
	} else {
		checkAPICallResult(t, *expectGetJobResponse, *result)
		t.Log("GetJob passed!")
	}
}

func TestCancelJob(t *testing.T) {
	var (
		expectCancelJobResponse = &CancelJobResponse{
			Detail: "Method \"POST\" not allowed.",
		}

		testJobID = 301
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.JobService.CancelJob(testJobID, map[string]interface{}{}, map[string]string{})

	if err != nil {
		t.Fatalf("CancelJob err: %s", err)
	} else {
		checkAPICallResult(t, *expectCancelJobResponse, *result)
		t.Log("CancelJob passed!")
	}
}

func TestRelaunchJob(t *testing.T) {
	var (
		expectLaunchJobResponse = &JobLaunch{
			Job:           499,
			IgnoredFields: map[string]string{},
			ID:            499,
			Type:          "job",
			URL:           "/api/v2/jobs/499/",
			Related: &Related{
				CreatedBy:          "/api/v2/users/1/",
				ModifiedBy:         "/api/v2/users/1/",
				Labels:             "/api/v2/jobs/499/labels/",
				Inventory:          "/api/v2/inventories/1/",
				Project:            "/api/v2/projects/4/",
				Credential:         "/api/v2/credentials/1/",
				ExtraCredentials:   "/api/v2/jobs/499/extra_credentials/",
				Credentials:        "/api/v2/jobs/499/credentials/",
				UnifiedJobTemplate: "/api/v2/job_templates/5/",
				Stdout:             "/api/v2/jobs/499/stdout/",
				Notifications:      "/api/v2/jobs/499/notifications/",
				JobHostSummaries:   "/api/v2/jobs/499/job_host_summaries/",
				JobEvents:          "/api/v2/jobs/499/job_events/",
				ActivityStream:     "/api/v2/jobs/499/activity_stream/",
				JobTemplate:        "/api/v2/job_templates/5/",
				Cancel:             "/api/v2/jobs/499/cancel/",
				CreateSchedule:     "/api/v2/jobs/499/create_schedule/",
				Relaunch:           "/api/v2/jobs/499/relaunch/",
			},

			SummaryFields: &Summary{
				JobTemplate: &JobTemplateSummary{
					ID:          5,
					Name:        "Demo Job Template",
					Description: "",
				},

				Inventory: &Inventory{
					ID:                           1,
					Name:                         "Demo Inventory",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   2,
					HostsWithActiveFailures:      0,
					TotalGroups:                  0,
					GroupsWithActiveFailures:     0,
					HasInventorySources:          false,
					TotalInventorySources:        0,
					InventorySourcesWithFailures: 0,
					OrganizationID:               1,
					Kind:                         "",
				},

				Credential: &Credential{
					Description:      "",
					CredentialTypeID: 1,
					ID:               1,
					Kind:             "ssh",
					Name:             "Demo Credential",
				},

				UnifiedJobTemplate: &UnifiedJobTemplate{
					ID:             5,
					Name:           "Demo Job Template",
					Description:    "",
					UnifiedJobType: "job",
				},

				Project: &Project{
					ID:          4,
					Name:        "Demo Project",
					Description: "",
					Status:      "successful",
					ScmType:     "git",
				},

				CreatedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},

				ModifiedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},

				UserCapabilities: &UserCapabilities{
					Start:  true,
					Delete: true,
				},

				Labels: &Labels{
					Count:   0,
					Results: []interface{}{},
				},

				ExtraCredentials: []interface{}{},
				Credentials: []Credential{
					{
						Description:      "",
						CredentialTypeID: 1,
						ID:               1,
						Kind:             "ssh",
						Name:             "Demo Credential",
					},
				},
			},

			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-25T04:25:11.312072Z")
				return t
			}(),

			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-25T04:25:11.362046Z")
				return t
			}(),

			Name:                    "Demo Job Template",
			Description:             "",
			JobType:                 "run",
			Inventory:               1,
			Project:                 4,
			Playbook:                "hello-world.yml",
			Forks:                   0,
			Limit:                   "",
			Verbosity:               0,
			ExtraVars:               "{}",
			JobTags:                 "",
			ForceHandlers:           false,
			SkipTags:                "",
			StartAtTask:             "",
			Timeout:                 0,
			UseFactCache:            false,
			UnifiedJobTemplate:      5,
			LaunchType:              "manual",
			Status:                  "pending",
			Failed:                  false,
			Started:                 nil,
			Finished:                nil,
			Elapsed:                 0,
			JobArgs:                 "",
			JobCwd:                  "",
			JobEnv:                  map[string]string{},
			JobExplanation:          "",
			ExecutionNode:           "",
			ResultTraceback:         "",
			EventProcessingFinished: false,
			JobTemplate:             5,
			PasswordsNeededToStart:  []interface{}{},
			AskDiffModeOnLaunch:     false,
			AskVariablesOnLaunch:    false,
			AskLimitOnLaunch:        false,
			AskTagsOnLaunch:         false,
			AskSkipTagsOnLaunch:     false,
			AskJobTypeOnLaunch:      false,
			AskVerbosityOnLaunch:    false,
			AskInventoryOnLaunch:    false,
			AskCredentialOnLaunch:   false,
			AllowSimultaneous:       false,
			Artifacts:               map[string]string{},
			ScmRevision:             "",
			InstanceGroup:           nil,
			DiffMode:                false,
			Credential:              1,
			VaultCredential:         nil,
		}

		testJobID = 508
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.JobService.RelaunchJob(testJobID, map[string]interface{}{"hosts": "all"}, map[string]string{})

	if err != nil {
		t.Fatalf("RelaunchJob err: %s", err)
	} else {
		checkAPICallResult(t, *expectLaunchJobResponse, *result)
		t.Log("RelaunchJob passed!")
	}
}

func TestGetHostSummaries(t *testing.T) {
	var (
		expectHostSummariesResponse = []HostSummary{
			{
				ID:   65,
				Type: "job_host_summary",
				URL:  "/api/v2/job_host_summaries/65/",
				Related: &Related{
					Job:  "/api/v2/jobs/301/",
					Host: "/api/v2/hosts/1/",
				},
				SummaryFields: &HostSummaryFields{
					Job: &HostSummaryJob{
						ID:              301,
						Name:            "Hello-world",
						Description:     "",
						Status:          "canceled",
						Failed:          true,
						Elapsed:         7.68,
						JobTemplateID:   8,
						JobTemplateName: "Hello-world",
					},
					Host: &HostSummaryHost{
						ID:                  1,
						Name:                "localhost",
						Description:         "",
						HasActiveFailures:   false,
						HasInventorySources: false,
					},
				},

				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:51.619497Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:51.619511Z")
					return t
				}(),

				Job:       301,
				Host:      1,
				HostName:  "localhost",
				Changed:   1,
				Dark:      0,
				Failures:  0,
				Ok:        1,
				Processed: 1,
				Skipped:   0,
				Failed:    false,
			},
		}

		testJobID = 301
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.JobService.GetHostSummaries(testJobID, map[string]string{})

	if err != nil {
		t.Fatalf("GetHostSummaries err: %s", err)
	} else {
		checkAPICallResult(t, expectHostSummariesResponse, result)
		t.Log("GetHostSummaries passed!")
	}
}

func TestGetJobEvents(t *testing.T) {
	var (
		expectJobEventsResponse = []JobEvent{
			{
				ID:   682,
				Type: "job_event",
				URL:  "/api/v2/job_events/682/",
				Related: &Related{
					Job: "/api/v2/jobs/301/",
				},
				SummaryFields: &HostSummaryFields{
					Role: map[string]string{},
					Job: &HostSummaryJob{
						ID:              301,
						Name:            "Hello-world",
						Description:     "",
						Status:          "canceled",
						Failed:          true,
						Elapsed:         7.68,
						JobTemplateID:   8,
						JobTemplateName: "Hello-world",
					},
				},

				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:51.249805Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-06-05T09:13:51.249818Z")
					return t
				}(),

				Job:          301,
				Event:        "verbose",
				Counter:      1,
				EventDisplay: "Verbose",
				EventData:    &EventData{},
				EventLevel:   0,
				Failed:       false,
				Changed:      false,
				UUID:         "",
				ParentUUID:   "",
				Host:         nil,
				HostName:     "",
				Parent:       nil,
				Playbook:     "",
				Play:         "",
				Task:         "",
				Role:         "",
				Stdout:       "\u001b[0;31m [ERROR]:\u001b[0m",
				StartLine:    0,
				EndLine:      1,
				Verbosity:    0,
			},
		}
		testJobID = 301
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.JobService.GetJobEvents(testJobID, map[string]string{
		"order_by":  "start_line",
		"page_size": "1000000",
	})

	if err != nil {
		t.Fatalf("GetJobEvents err: %s", err)
	} else {
		checkAPICallResult(t, expectJobEventsResponse, result)
		t.Log("GetJobEvents passed!")
	}
}
