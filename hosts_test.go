package awx

import (
	"testing"
	"time"
)

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
					AnsibleFacts:       "/api/v2/hosts/1/ansible_facts/",
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
		t.Log("ListHosts passed!")
	}
}

func TestCreateHost(t *testing.T) {
	var (
		expectCreateHostResponse = &Host{
			ID:   3,
			Type: "host",
			URL:  "/api/v2/hosts/3/",
			Related: &Related{
				NamedURL:           "/api/v2/hosts/test++Demo Inventory++Default/",
				CreatedBy:          "/api/v2/users/2/",
				ModifiedBy:         "/api/v2/users/2/",
				JobHostSummaries:   "/api/v2/hosts/3/job_host_summaries/",
				VariableData:       "/api/v2/hosts/3/variable_data/",
				JobEvents:          "/api/v2/hosts/3/job_events/",
				AdHocCommands:      "/api/v2/hosts/3/ad_hoc_commands/",
				InventorySources:   "/api/v2/hosts/3/inventory_sources/",
				FactVersions:       "/api/v2/hosts/3/fact_versions/",
				SmartInventories:   "/api/v2/hosts/3/smart_inventories/",
				Groups:             "/api/v2/hosts/3/groups/",
				AllGroups:          "/api/v2/hosts/3/all_groups/",
				ActivityStream:     "/api/v2/hosts/3/activity_stream/",
				AdHocCommandEvents: "/api/v2/hosts/3/ad_hoc_command_events/",
				Insights:           "/api/v2/hosts/3/insights/",
				Inventory:          "/api/v2/inventories/1/",
				AnsibleFacts:       "/api/v2/hosts/3/ansible_facts/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           1,
					Name:                         "Demo Inventory",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   1,
					HostsWithActiveFailures:      0,
					TotalGroups:                  3,
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
					Count:   0,
					Results: []Result{},
				},
				RecentJobs: []interface{}{},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456501Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456512Z")
				return t
			}(),
			Name:                 "test",
			Description:          "test create host",
			Inventory:            1,
			Enabled:              true,
			InstanceID:           "",
			Variables:            "ansible_host: localhost",
			HasActiveFailures:    false,
			HasInventorySources:  false,
			LastJob:              nil,
			LastJobHostSummary:   nil,
			InsightsSystemID:     nil,
			AnsibleFactsModified: nil,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.HostService.CreateHost(map[string]interface{}{
		"name":        "test",
		"inventory":   1,
		"description": "test create host",
		"enabled":     true,
		"instance_id": "",
		"variables":   "ansible_host: localhost",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("CreateHost err: %s", err)
	} else {
		checkAPICallResult(t, expectCreateHostResponse, result)
		t.Log("CreateHost passed!")
	}
}

func TestUpdateHost(t *testing.T) {
	var (
		expectUpdateHostResponse = &Host{
			ID:   3,
			Type: "host",
			URL:  "/api/v2/hosts/3/",
			Related: &Related{
				NamedURL:           "/api/v2/hosts/testUpdate++Demo Inventory++Default/",
				CreatedBy:          "/api/v2/users/2/",
				ModifiedBy:         "/api/v2/users/2/",
				JobHostSummaries:   "/api/v2/hosts/3/job_host_summaries/",
				VariableData:       "/api/v2/hosts/3/variable_data/",
				JobEvents:          "/api/v2/hosts/3/job_events/",
				AdHocCommands:      "/api/v2/hosts/3/ad_hoc_commands/",
				InventorySources:   "/api/v2/hosts/3/inventory_sources/",
				FactVersions:       "/api/v2/hosts/3/fact_versions/",
				SmartInventories:   "/api/v2/hosts/3/smart_inventories/",
				Groups:             "/api/v2/hosts/3/groups/",
				AllGroups:          "/api/v2/hosts/3/all_groups/",
				ActivityStream:     "/api/v2/hosts/3/activity_stream/",
				AdHocCommandEvents: "/api/v2/hosts/3/ad_hoc_command_events/",
				Insights:           "/api/v2/hosts/3/insights/",
				Inventory:          "/api/v2/inventories/1/",
				AnsibleFacts:       "/api/v2/hosts/3/ansible_facts/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           1,
					Name:                         "Demo Inventory",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   1,
					HostsWithActiveFailures:      0,
					TotalGroups:                  3,
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
					Count:   0,
					Results: []Result{},
				},
				RecentJobs: []interface{}{},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456501Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456512Z")
				return t
			}(),
			Name:                 "testUpdate",
			Description:          "test create host",
			Inventory:            1,
			Enabled:              true,
			InstanceID:           "",
			Variables:            "ansible_host: localhost",
			HasActiveFailures:    false,
			HasInventorySources:  false,
			LastJob:              nil,
			LastJobHostSummary:   nil,
			InsightsSystemID:     nil,
			AnsibleFactsModified: nil,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.HostService.UpdateHost(3, map[string]interface{}{
		"name":      "testUpdate",
		"inventory": 1,
	}, map[string]string{})

	if err != nil {
		t.Fatalf("UpdateHost err: %s", err)
	} else {
		checkAPICallResult(t, expectUpdateHostResponse, result)
		t.Log("UpdateHost passed!")
	}
}

func TestAssociateHostGroup(t *testing.T) {
	var (
		expectAssociateHostGroupResponse = &Host{
			ID:   3,
			Type: "host",
			URL:  "/api/v2/hosts/3/",
			Related: &Related{
				NamedURL:           "/api/v2/hosts/testUpdate++Demo Inventory++Default/",
				CreatedBy:          "/api/v2/users/2/",
				ModifiedBy:         "/api/v2/users/2/",
				JobHostSummaries:   "/api/v2/hosts/3/job_host_summaries/",
				VariableData:       "/api/v2/hosts/3/variable_data/",
				JobEvents:          "/api/v2/hosts/3/job_events/",
				AdHocCommands:      "/api/v2/hosts/3/ad_hoc_commands/",
				InventorySources:   "/api/v2/hosts/3/inventory_sources/",
				FactVersions:       "/api/v2/hosts/3/fact_versions/",
				SmartInventories:   "/api/v2/hosts/3/smart_inventories/",
				Groups:             "/api/v2/hosts/3/groups/",
				AllGroups:          "/api/v2/hosts/3/all_groups/",
				ActivityStream:     "/api/v2/hosts/3/activity_stream/",
				AdHocCommandEvents: "/api/v2/hosts/3/ad_hoc_command_events/",
				Insights:           "/api/v2/hosts/3/insights/",
				Inventory:          "/api/v2/inventories/1/",
				AnsibleFacts:       "/api/v2/hosts/3/ansible_facts/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           1,
					Name:                         "Demo Inventory",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   1,
					HostsWithActiveFailures:      0,
					TotalGroups:                  3,
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
					Count: 1,
					Results: []Result{
						{ID: 10,
							Name: "testGroup"},
					},
				},
				RecentJobs: []interface{}{},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456501Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456512Z")
				return t
			}(),
			Name:                 "testUpdate",
			Description:          "test create host",
			Inventory:            1,
			Enabled:              true,
			InstanceID:           "",
			Variables:            "ansible_host: localhost",
			HasActiveFailures:    false,
			HasInventorySources:  false,
			LastJob:              nil,
			LastJobHostSummary:   nil,
			InsightsSystemID:     nil,
			AnsibleFactsModified: nil,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.HostService.AssociateGroup(3, map[string]interface{}{
		"id": 10,
	}, map[string]string{})

	if err != nil {
		t.Fatalf("AssociateGroup err: %s", err)
	} else {
		checkAPICallResult(t, expectAssociateHostGroupResponse, result)
		t.Log("AssociateGroup passed!")
	}
}

func TestDisAssociateHostGroup(t *testing.T) {
	var (
		expectDisAssociateHostGroupResponse = &Host{
			ID:   3,
			Type: "host",
			URL:  "/api/v2/hosts/3/",
			Related: &Related{
				NamedURL:           "/api/v2/hosts/testUpdate++Demo Inventory++Default/",
				CreatedBy:          "/api/v2/users/2/",
				ModifiedBy:         "/api/v2/users/2/",
				JobHostSummaries:   "/api/v2/hosts/3/job_host_summaries/",
				VariableData:       "/api/v2/hosts/3/variable_data/",
				JobEvents:          "/api/v2/hosts/3/job_events/",
				AdHocCommands:      "/api/v2/hosts/3/ad_hoc_commands/",
				InventorySources:   "/api/v2/hosts/3/inventory_sources/",
				FactVersions:       "/api/v2/hosts/3/fact_versions/",
				SmartInventories:   "/api/v2/hosts/3/smart_inventories/",
				Groups:             "/api/v2/hosts/3/groups/",
				AllGroups:          "/api/v2/hosts/3/all_groups/",
				ActivityStream:     "/api/v2/hosts/3/activity_stream/",
				AdHocCommandEvents: "/api/v2/hosts/3/ad_hoc_command_events/",
				Insights:           "/api/v2/hosts/3/insights/",
				Inventory:          "/api/v2/inventories/1/",
				AnsibleFacts:       "/api/v2/hosts/3/ansible_facts/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           1,
					Name:                         "Demo Inventory",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   1,
					HostsWithActiveFailures:      0,
					TotalGroups:                  3,
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
					Count:   0,
					Results: []Result{},
				},
				RecentJobs: []interface{}{},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456501Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-09-01T11:18:16.456512Z")
				return t
			}(),
			Name:                 "testUpdate",
			Description:          "test create host",
			Inventory:            1,
			Enabled:              true,
			InstanceID:           "",
			Variables:            "ansible_host: localhost",
			HasActiveFailures:    false,
			HasInventorySources:  false,
			LastJob:              nil,
			LastJobHostSummary:   nil,
			InsightsSystemID:     nil,
			AnsibleFactsModified: nil,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.HostService.DisAssociateGroup(3, map[string]interface{}{
		"id": 10,
	}, map[string]string{})

	if err != nil {
		t.Fatalf("DisAssociateGroup err: %s", err)
	} else {
		checkAPICallResult(t, expectDisAssociateHostGroupResponse, result)
		t.Log("DisAssociateGroup passed!")
	}
}
func TestDeleteHost(t *testing.T) {
	var (
		expectDeleteHostResponse = &Host{}
	)
	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.HostService.DeleteHost(3)

	if err != nil {
		t.Fatalf("DeleteHost err: %s", err)
	} else {
		checkAPICallResult(t, expectDeleteHostResponse, result)
		t.Log("DeleteHost passed!")
	}
}
