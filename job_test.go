package awx

import (
	"log"
	"reflect"
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

		testJobId = 301
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.JobService.GetJob(testJobId, map[string]string{})
	if err != nil {
		log.Fatalf("GetJob err: %s", err)
	}

	if !reflect.DeepEqual(*result, *expectGetJobResponse) {
		log.Fatalf("GetJob resp not as expected, expected: %v, resp result: %v", *expectGetJobResponse, *result)
	}

	log.Println("GetJob passed!")
}
