package awx

import (
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func checkAPICallResult(t *testing.T, expected interface{}, got interface{}) {
	if diff := pretty.Compare(expected, got); diff != "" {
		t.Fatalf("Diff: (-got +want)\n%s", diff)
	}
}

func TestListHosts(t *testing.T) {
	var (
		expectListHostsResponse = []*Host{
			{
				ID:   1,
				Type: "host",
				URL:  "/api/v2/hosts/1/",
				Related: &Related{
					CreatedBy:          "/api/v2/users/2/",
					ModifiedBy:         "/api/v2/users/2/",
					JobHostSummaries:   "/api/v2/hosts/1/job_host_summaries/",
					VariableData:       "/api/v2/hosts/1/variable_data/",
					JobEvents:          "/api/v2/hosts/1/job_events/",
					AdHocCommands:      "/api/v2/hosts/1/ad_hoc_commands/",
					InventorySources:   "/api/v2/hosts/1/inventory_sources/",
					FactVersions:       "/api/v2/hosts/1/fact_versions/",
					SmartInventories:   "/api/v2/hosts/1/smart_inventories/",
					Groups:             "/api/v2/hosts/1/groups/",
					ActivityStream:     "/api/v2/hosts/1/activity_stream/",
					AllGroups:          "/api/v2/hosts/1/all_groups/",
					AdHocCommandEvents: "/api/v2/hosts/1/ad_hoc_command_events/",
					Insights:           "/api/v2/hosts/1/insights/",
					Inventory:          "/api/v2/inventories/1/",

					// FIXME: why this line could NOT be unmarshaled in json？
					// AnsibleFacts: "/api/v2/hosts/1/ansible_facts/",
				},
				SummaryFields: &Summary{
					Inventory: &Inventory{
						ID:                           1,
						Name:                         "Demo Inventory",
						Description:                  "",
						HasActiveFailures:            false,
						TotalHosts:                   1,
						HostsWithActiveFailures:      0,
						TotalGroups:                  2,
						GroupsWithActiveFailures:     0,
						HasInventorySources:          false,
						TotalInventorySources:        0,
						InventorySourcesWithFailures: 0,
						OrganizationID:               1,
						Kind:                         "",
					},
					CreatedBy: &ByUserSummary{
						ID:        2,
						Username:  "admin",
						FirstName: "",
						LastName:  "",
					},
					ModifiedBy: &ByUserSummary{
						ID:        2,
						Username:  "admin",
						FirstName: "",
						LastName:  "",
					},
					UserCapabilities: &UserCapabilities{
						Edit:   true,
						Delete: true,
					},
					Groups: &Groups{
						Count: 2,
						Results: []Result{
							{ID: 19,
								Name: "ciao"},
							{ID: 21,
								Name: "test"},
						},
					},
					RecentJobs: []interface{}{},
				},
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-08-27T13:47:11.145028Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-08-27T13:47:11.145042Z")
					return t
				}(),
				Name:                 "localhost",
				Description:          "",
				Inventory:            1,
				Enabled:              true,
				InstanceID:           "",
				Variables:            "ansible_connection: local",
				HasActiveFailures:    false,
				HasInventorySources:  false,
				LastJob:              nil,
				LastJobHostSummary:   nil,
				InsightsSystemID:     nil,
				AnsibleFactsModified: nil,
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.HostService.ListHosts(map[string]string{
		"name": "localhost",
	})
	if err != nil {
		t.Fatalf("ListHosts err: %s", err)
	} else {
		checkAPICallResult(t, expectListHostsResponse, result)
		t.Log("ListHost passed!")
	}
}

/*
func TestCreateGroup(t *testing.T) {
	var (
		expectCreateGroupResponse = &Group{
			ID:   21,
			Type: 0,
			URL:  "/api/v2/groups/21/",
			Related: &Related{
				CreatedBy:         "/api/v2/users/11/",
				JobHostSummaries:  "/api/v2/groups/21/job_host_summaries/",
				VariableData:      "/api/v2/groups/21/variable_data/",
				JobEvents:         "/api/v2/groups/21/job_events/",
				PotentialChildren: "/api/v2/groups/21/potential_children/",
				AdHocCommands:     "/api/v2/groups/21/ad_hoc_commands/",
				AllHosts:          "/api/v2/groups/21/all_hosts/",
				ActivityStream:    "/api/v2/groups/21/activity_stream/",
				Hosts:             "/api/v2/groups/21/hosts/",
				Children:          "/api/v2/groups/21/children/",
				InventorySources:  "/api/v2/groups/21/inventory_sources/",
				Inventory:         "/api/v2/inventories/9/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           9,
					Name:                         "test",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   8,
					HostsWithActiveFailures:      0,
					TotalGroups:                  4,
					GroupsWithActiveFailures:     0,
					HasInventorySources:          false,
					TotalInventorySources:        0,
					InventorySourcesWithFailures: 0,
					OrganizationID:               16,
					Kind:                         "",
				},
				CreatedBy: &ByUserSummary{
					ID:        11,
					Username:  "demouser",
					FirstName: "Demo",
					LastName:  "User",
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Copy:   true,
					Delete: true,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:27:46.686176Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:28:07.127040Z")
				return t
			}(),
			Name:                     "Demo Group",
			Description:              "",
			Inventory:                9,
			Variables:                "",
			HasActiveFailures:        false,
			TotalHosts:               3,
			HostsWithActiveFailures:  0,
			TotalGroups:              0,
			GroupsWithActiveFailures: 0,
			HasInventorySources:      false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.GroupService.CreateGroup(map[string]interface{}{
		"name":        "Test Group",
		"description": "for testing CreateGroup api",
		"inventory":   1,
		"variables":   "",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("CreateGroup err: %s", err)
	} else if !reflect.DeepEqual(result, expectCreateGroupResponse) {
		t.Logf("expected: %v", expectCreateGroupResponse)
		t.Logf("result: %v", result)
		t.Fatal("CreateGroup response not as expected")
	} else {
		t.Log("CreateGroup passed!")
	}
}

func TestUpdateGroup(t *testing.T) {
	var (
		expectUpdateGroupResponse = &Group{
			ID:   21,
			Type: 0,
			URL:  "/api/v2/groups/21/",
			Related: &Related{
				CreatedBy:         "/api/v2/users/11/",
				JobHostSummaries:  "/api/v2/groups/21/job_host_summaries/",
				VariableData:      "/api/v2/groups/21/variable_data/",
				JobEvents:         "/api/v2/groups/21/job_events/",
				PotentialChildren: "/api/v2/groups/21/potential_children/",
				AdHocCommands:     "/api/v2/groups/21/ad_hoc_commands/",
				AllHosts:          "/api/v2/groups/21/all_hosts/",
				ActivityStream:    "/api/v2/groups/21/activity_stream/",
				Hosts:             "/api/v2/groups/21/hosts/",
				Children:          "/api/v2/groups/21/children/",
				InventorySources:  "/api/v2/groups/21/inventory_sources/",
				Inventory:         "/api/v2/inventories/9/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           9,
					Name:                         "test",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   8,
					HostsWithActiveFailures:      0,
					TotalGroups:                  4,
					GroupsWithActiveFailures:     0,
					HasInventorySources:          false,
					TotalInventorySources:        0,
					InventorySourcesWithFailures: 0,
					OrganizationID:               16,
					Kind:                         "",
				},
				CreatedBy: &ByUserSummary{
					ID:        11,
					Username:  "demouser",
					FirstName: "Demo",
					LastName:  "User",
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Copy:   true,
					Delete: true,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:27:46.686176Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:28:07.127040Z")
				return t
			}(),
			Name:                     "Demo Group",
			Description:              "Add description",
			Inventory:                9,
			Variables:                "",
			HasActiveFailures:        false,
			TotalHosts:               3,
			HostsWithActiveFailures:  0,
			TotalGroups:              0,
			GroupsWithActiveFailures: 0,
			HasInventorySources:      false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.GroupService.UpdateGroup(21, map[string]interface{}{
		"description": "Add description",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("UpdateGroup err: %s", err)
	} else if !reflect.DeepEqual(result, expectUpdateGroupResponse) {
		t.Logf("expected: %v", expectUpdateGroupResponse)
		t.Logf("result: %v", result)
		t.Fatal("UpdateGroup response not as expected")
	} else {
		t.Log("UpdateGroup passed!")
	}
}
func TestDeleteGroup(t *testing.T) {
	var (
		expectDeleteGroupResponse = &Group{}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.GroupService.DeleteGroup(21)

	if err != nil {
		t.Fatalf("DeleteGroup by ID err: %s", err)
	}

	if !reflect.DeepEqual(result, expectDeleteGroupResponse) {
		t.Fatalf("DeleteGroup resp not as expected, expected: %v, resp result: %v", expectDeleteGroupResponse, result)
	}

	t.Log("DeleteGroup passed!")
}
*/
