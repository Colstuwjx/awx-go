package awx

import (
	"testing"
	"time"
)

func TestListWorkflowJobTemplates(t *testing.T) {
	var (
		expectListJobTempaltesResponse = []*WorkflowJobTemplate{
			{
				ID:   5,
				Type: "job_template",
				URL:  "/api/v2/workflow_job_templates/5/",
				Related: &Related{
					CreatedBy:                    "/api/v2/users/1/",
					ModifiedBy:                   "/api/v2/users/1/",
					Labels:                       "/api/v2/workflow_job_templates/5/labels/",
					Inventory:                    "/api/v2/inventories/1/",
					Project:                      "/api/v2/projects/4/",
					Credential:                   "/api/v2/credentials/1/",
					ExtraCredentials:             "/api/v2/workflow_job_templates/5/extra_credentials/",
					Credentials:                  "/api/v2/workflow_job_templates/5/credentials/",
					NotificationTemplatesError:   "/api/v2/workflow_job_templates/5/notification_templates_error/",
					NotificationTemplatesSuccess: "/api/v2/workflow_job_templates/5/notification_templates_success/",
					NotificationTemplatesAny:     "/api/v2/workflow_job_templates/5/notification_templates_any/",
					Jobs:                         "/api/v2/workflow_job_templates/5/jobs/",
					ObjectRoles:                  "/api/v2/workflow_job_templates/5/object_roles/",
					AccessList:                   "/api/v2/workflow_job_templates/5/access_list/",
					Launch:                       "/api/v2/workflow_job_templates/5/launch/",
					InstanceGroups:               "/api/v2/workflow_job_templates/5/instance_groups/",
					Schedules:                    "/api/v2/workflow_job_templates/5/schedules/",
					Copy:                         "/api/v2/workflow_job_templates/5/copy/",
					ActivityStream:               "/api/v2/workflow_job_templates/5/activity_stream/",
					SurveySpec:                   "/api/v2/workflow_job_templates/5/survey_spec/",
				},
				SummaryFields: &Summary{
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

					Project: &Project{
						ID:          4,
						Name:        "Demo Project",
						Description: "",
						Status:      "never updated",
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

					ObjectRoles: &ObjectRoles{
						AdminRole: &ApplyRole{
							ID:          27,
							Description: "Can manage all aspects of the job template",
							Name:        "Admin",
						},

						ExecuteRole: &ApplyRole{
							ID:          26,
							Description: "May run the job template",
							Name:        "Execute",
						},

						ReadRole: &ApplyRole{
							ID:          25,
							Description: "May view settings for the job template",
							Name:        "Read",
						},
					},

					UserCapabilities: &UserCapabilities{
						Edit:     true,
						Start:    true,
						Copy:     true,
						Schedule: true,
						Delete:   true,
					},

					Labels: &Labels{
						Count:   0,
						Results: []interface{}{},
					},

					RecentJobs: []interface{}{},

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
					t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.773593Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.773605Z")
					return t
				}(),

				Name:                 "Demo Job Template",
				Description:          "",
				Inventory:            1,
				Limit:                "",
				ExtraVars:            "",
				LastJobRun:           nil,
				LastJobFailed:        false,
				NextJobRun:           nil,
				Status:               "never updated",
				AskVariablesOnLaunch: false,
				AskLimitOnLaunch:     false,
				AskInventoryOnLaunch: false,
				SurveyEnabled:        false,
				AllowSimultaneous:    false,
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.WorkflowJobTemplateService.ListWorkflowJobTemplates(map[string]string{})

	if err != nil {
		t.Logf("CreateWorkflowJobTemplate err: %s", err)
	} else {
		checkAPICallResult(t, expectListJobTempaltesResponse, result)
	}
}

func TestLauchWorkflowJob(t *testing.T) {
	var (
		testWorkflowJobTemplateID = 5
		testInventoryID           = 1

		expectLaunchWorkflowJobTemplateResponse = &WorkflowJobLaunch{
			Job:           499,
			IgnoredFields: map[string]string{},
			ID:            499,
			Type:          "job",
			URL:           "/api/v2/jobs/499/",
			Related: &Related{
				CreatedBy:                  "/api/v2/users/1/",
				ModifiedBy:                 "/api/v2/users/1/",
				Labels:                     "/api/v2/jobs/499/labels/",
				Inventory:                  "/api/v2/inventories/1/",
				Project:                    "/api/v2/projects/4/",
				Credential:                 "/api/v2/credentials/1/",
				ExtraCredentials:           "/api/v2/jobs/499/extra_credentials/",
				Credentials:                "/api/v2/jobs/499/credentials/",
				Stdout:                     "/api/v2/jobs/499/stdout/",
				Notifications:              "/api/v2/jobs/499/notifications/",
				JobHostSummaries:           "/api/v2/jobs/499/job_host_summaries/",
				JobEvents:                  "/api/v2/jobs/499/job_events/",
				ActivityStream:             "/api/v2/jobs/499/activity_stream/",
				Cancel:                     "/api/v2/jobs/499/cancel/",
				CreateSchedule:             "/api/v2/jobs/499/create_schedule/",
				Relaunch:                   "/api/v2/jobs/499/relaunch/",
			},

			// SummaryFields: &Summary{
			// 	WorkflowJobTemplate: &WorkflowJobTemplateSummary{
			// 		ID:          5,
			// 		Name:        "Demo Job Template",
			// 		Description: "",
			// 	},

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

				UnifiedWorkflowJobTemplate: &UnifiedWorkflowJobTemplate{
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

			Name:                       "Demo Job Template",
			Description:                "",
			JobType:                    "run",
			Inventory:                  1,
			Project:                    4,
			Playbook:                   "hello-world.yml",
			Forks:                      0,
			Limit:                      "",
			Verbosity:                  0,
			ExtraVars:                  "{}",
			JobTags:                    "",
			ForceHandlers:              false,
			SkipTags:                   "",
			StartAtTask:                "",
			Timeout:                    0,
			UseFactCache:               false,
			UnifiedWorkflowJobTemplate: 5,
			LaunchType:                 "manual",
			Status:                     "pending",
			Failed:                     false,
			Started:                    nil,
			Finished:                   nil,
			Elapsed:                    0,
			JobArgs:                    "",
			JobCwd:                     "",
			JobEnv:                     map[string]string{},
			JobExplanation:             "",
			ExecutionNode:              "",
			ResultTraceback:            "",
			EventProcessingFinished:    false,
			WorkflowJobTemplate:        5,
			PasswordsNeededToStart:     []interface{}{},
			AskDiffModeOnLaunch:        false,
			AskVariablesOnLaunch:       false,
			AskLimitOnLaunch:           false,
			AskTagsOnLaunch:            false,
			AskSkipTagsOnLaunch:        false,
			AskJobTypeOnLaunch:         false,
			AskVerbosityOnLaunch:       false,
			AskInventoryOnLaunch:       false,
			AskCredentialOnLaunch:      false,
			AllowSimultaneous:          false,
			Artifacts:                  map[string]string{},
			ScmRevision:                "",
			InstanceGroup:              nil,
			DiffMode:                   false,
			Credential:                 1,
			VaultCredential:            nil,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.WorkflowJobTemplateService.Launch(testWorkflowJobTemplateID, map[string]interface{}{
		"inventory": testInventoryID,
	}, map[string]string{})

	if err != nil {
		t.Logf("CreateWorkflowJobTemplate err: %s", err)
	} else {
		checkAPICallResult(t, expectLaunchWorkflowJobTemplateResponse, result)
	}
}

func TestCreateWorkflowJobTemplate(t *testing.T) {
	var (
		expectCreateJobTempalteResponse = &WorkflowJobTemplate{
			ID:   5,
			Type: "job_template",
			URL:  "/api/v2/workflow_job_templates/5/",
			Related: &Related{
				NamedURL:                     "/api/v2/workflow_job_templates/Demo Job Template/",
				CreatedBy:                    "/api/v2/users/1/",
				ModifiedBy:                   "/api/v2/users/1/",
				Labels:                       "/api/v2/workflow_job_templates/5/labels/",
				Inventory:                    "/api/v2/inventories/1/",
				Project:                      "/api/v2/projects/4/",
				Credential:                   "/api/v2/credentials/1/",
				ExtraCredentials:             "/api/v2/workflow_job_templates/5/extra_credentials/",
				Credentials:                  "/api/v2/workflow_job_templates/5/credentials/",
				NotificationTemplatesError:   "/api/v2/workflow_job_templates/5/notification_templates_error/",
				NotificationTemplatesSuccess: "/api/v2/workflow_job_templates/5/notification_templates_success/",
				NotificationTemplatesAny:     "/api/v2/workflow_job_templates/5/notification_templates_any/",
				Jobs:                         "/api/v2/workflow_job_templates/5/jobs/",
				ObjectRoles:                  "/api/v2/workflow_job_templates/5/object_roles/",
				AccessList:                   "/api/v2/workflow_job_templates/5/access_list/",
				Launch:                       "/api/v2/workflow_job_templates/5/launch/",
				InstanceGroups:               "/api/v2/workflow_job_templates/5/instance_groups/",
				Schedules:                    "/api/v2/workflow_job_templates/5/schedules/",
				Copy:                         "/api/v2/workflow_job_templates/5/copy/",
				ActivityStream:               "/api/v2/workflow_job_templates/5/activity_stream/",
				SurveySpec:                   "/api/v2/workflow_job_templates/5/survey_spec/",
			},
			SummaryFields: &Summary{
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
				Project: &Project{
					ID:          4,
					Name:        "Demo Project",
					Description: "",
					Status:      "never updated",
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
				ObjectRoles: &ObjectRoles{
					AdminRole: &ApplyRole{
						ID:          28,
						Description: "Can manage all aspects of the job template",
						Name:        "Admin",
					},
					ExecuteRole: &ApplyRole{
						ID:          27,
						Description: "May run the job template",
						Name:        "Execute",
					},
					ReadRole: &ApplyRole{
						ID:          26,
						Description: "May view settings for the job template",
						Name:        "Read",
					},
				},
				UserCapabilities: &UserCapabilities{
					Edit:     true,
					Start:    true,
					Copy:     true,
					Schedule: true,
					Delete:   true,
				},
				Labels: &Labels{
					Count:   0,
					Results: []interface{}{},
				},
				RecentJobs: []interface{}{},
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
				t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:16.238510Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:16.238524Z")
				return t
			}(),
			Name:                 "Demo Job Template",
			Description:          "",
			Inventory:            1,
			Limit:                "",
			ExtraVars:            "",
			LastJobRun:           nil,
			LastJobFailed:        false,
			NextJobRun:           nil,
			Status:               "never updated",
			AskVariablesOnLaunch: false,
			AskLimitOnLaunch:     false,
			AskInventoryOnLaunch: false,
			SurveyEnabled:        false,
			AllowSimultaneous:    false,
		}
	)
	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.WorkflowJobTemplateService.CreateWorkflowJobTemplate(map[string]interface{}{
		"name":        "TestProject",
		"description": "Test project",
	}, map[string]string{})
	if err != nil {
		t.Logf("CreateWorkflowJobTemplate err: %s", err)
	} else {
		checkAPICallResult(t, expectCreateJobTempalteResponse, result)
	}
	t.Log("CreateWorkflowJobTemplate passed!")
}

func TestUpdateWorkflowJobTemplate(t *testing.T) {
	var (
		expectUpdateJobTempalteResponse = &WorkflowJobTemplate{
			ID:   5,
			Type: "job_template",
			URL:  "/api/v2/workflow_job_templates/5/",
			Related: &Related{
				NamedURL:                     "/api/v2/workflow_job_templates/Demo Job Template/",
				CreatedBy:                    "/api/v2/users/1/",
				ModifiedBy:                   "/api/v2/users/1/",
				Labels:                       "/api/v2/workflow_job_templates/5/labels/",
				Inventory:                    "/api/v2/inventories/1/",
				Project:                      "/api/v2/projects/4/",
				Credential:                   "/api/v2/credentials/1/",
				ExtraCredentials:             "/api/v2/workflow_job_templates/5/extra_credentials/",
				Credentials:                  "/api/v2/workflow_job_templates/5/credentials/",
				NotificationTemplatesError:   "/api/v2/workflow_job_templates/5/notification_templates_error/",
				NotificationTemplatesSuccess: "/api/v2/workflow_job_templates/5/notification_templates_success/",
				NotificationTemplatesAny:     "/api/v2/workflow_job_templates/5/notification_templates_any/",
				Jobs:                         "/api/v2/workflow_job_templates/5/jobs/",
				ObjectRoles:                  "/api/v2/workflow_job_templates/5/object_roles/",
				AccessList:                   "/api/v2/workflow_job_templates/5/access_list/",
				Launch:                       "/api/v2/workflow_job_templates/5/launch/",
				InstanceGroups:               "/api/v2/workflow_job_templates/5/instance_groups/",
				Schedules:                    "/api/v2/workflow_job_templates/5/schedules/",
				Copy:                         "/api/v2/workflow_job_templates/5/copy/",
				ActivityStream:               "/api/v2/workflow_job_templates/5/activity_stream/",
				SurveySpec:                   "/api/v2/workflow_job_templates/5/survey_spec/",
			},
			SummaryFields: &Summary{
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
				Project: &Project{
					ID:          4,
					Name:        "Demo Project",
					Description: "",
					Status:      "never updated",
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
				ObjectRoles: &ObjectRoles{
					AdminRole: &ApplyRole{
						ID:          28,
						Description: "Can manage all aspects of the job template",
						Name:        "Admin",
					},
					ExecuteRole: &ApplyRole{
						ID:          27,
						Description: "May run the job template",
						Name:        "Execute",
					},
					ReadRole: &ApplyRole{
						ID:          26,
						Description: "May view settings for the job template",
						Name:        "Read",
					},
				},
				UserCapabilities: &UserCapabilities{
					Edit:     true,
					Start:    true,
					Copy:     true,
					Schedule: true,
					Delete:   true,
				},
				Labels: &Labels{
					Count:   0,
					Results: []interface{}{},
				},
				RecentJobs: []interface{}{},
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
				t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:16.238510Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-06-28T16:31:16.238524Z")
				return t
			}(),
			Name:                 "Demo Job Template",
			Description:          "Test Update",
			Inventory:            1,
			Limit:                "",
			ExtraVars:            "",
			LastJobRun:           nil,
			LastJobFailed:        false,
			NextJobRun:           nil,
			Status:               "never updated",
			AskVariablesOnLaunch: false,
			AskLimitOnLaunch:     false,
			AskInventoryOnLaunch: false,
			SurveyEnabled:        false,
			AllowSimultaneous:    false,
		}
	)
	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.WorkflowJobTemplateService.UpdateWorkflowJobTemplate(5, map[string]interface{}{
		"description": "Test Update",
	}, map[string]string{})
	if err != nil {
		t.Logf("UpdateWorkflowJobTemplate err: %s", err)
	} else {
		checkAPICallResult(t, expectUpdateJobTempalteResponse, result)
	}
	t.Log("UpdateWorkflowJobTemplate passed!")
}
func TestDeleteWorkflowJobTemplate(t *testing.T) {
	var (
		expectDeleteJobTempleteResponse = &WorkflowJobTemplate{}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.WorkflowJobTemplateService.DeleteWorkflowJobTemplate(5)

	if err != nil {
		t.Fatalf("DeleteWorkflowJobTemplate err: %s", err)
	} else {
		checkAPICallResult(t, expectDeleteJobTempleteResponse, result)
		t.Log("DeleteWorkflowJobTemplate passed!")
	}
}
