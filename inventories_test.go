package awx

import (
	"reflect"
	"testing"
	"time"
)

func TestListInventories(t *testing.T) {
	var (
		expectListInventoriesResponse = []*Inventory{
			{
				ID:   1,
				Type: "inventory",
				URL:  "/api/v2/inventories/1/",
				Related: &Related{
					CreatedBy:      "/api/v2/users/1/",
					ModifiedBy:     "/api/v2/users/1/",
					JobTemplates:   "/api/v2/inventories/1/job_templates/",
					VariableData:   "/api/v2/inventories/1/variable_data/",
					RootGroups:     "/api/v2/inventories/1/root_groups/",
					ObjectRoles:    "/api/v2/inventories/1/object_roles/",
					AdHocCommands:  "/api/v2/inventories/1/ad_hoc_commands/",
					Script:         "/api/v2/inventories/1/script/",
					Tree:           "/api/v2/inventories/1/tree/",
					AccessList:     "/api/v2/inventories/1/access_list/",
					ActivityStream: "/api/v2/inventories/1/activity_stream/",
					InstanceGroups: "/api/v2/inventories/1/instance_groups/",
					Hosts:          "/api/v2/inventories/1/hosts/",
					Groups:         "/api/v2/inventories/1/groups/",
					Copy:           "/api/v2/inventories/1/copy/",
					UpdateInventorySources: "/api/v2/inventories/1/update_inventory_sources/",
					InventorySources:       "/api/v2/inventories/1/inventory_sources/",
					Organization:           "/api/v2/organizations/1/",
				},
				SummaryFields: &Summary{
					Organization: &OrgnizationSummary{
						ID:          1,
						Name:        "Default",
						Description: "",
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
						UseRole: &ApplyRole{
							ID:          23,
							Description: "Can use the inventory in a job template",
							Name:        "Use",
						},

						AdminRole: &ApplyRole{
							ID:          21,
							Description: "Can manage all aspects of the inventory",
							Name:        "Admin",
						},

						AdhocRole: &ApplyRole{
							ID:          20,
							Description: "May run ad hoc commands on an inventory",
							Name:        "Ad Hoc",
						},

						UpdateRole: &ApplyRole{
							ID:          24,
							Description: "May update project or inventory or group using the configured source update system",
							Name:        "Update",
						},

						ReadRole: &ApplyRole{
							ID:          22,
							Description: "May view settings for the inventory",
							Name:        "Read",
						},
					},

					UserCapabilities: &UserCapabilities{
						Edit:   true,
						Copy:   true,
						Adhoc:  true,
						Delete: true,
					},
				},

				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.657185Z")
					return t
				}(),

				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-05-30T09:42:22.412749Z")
					return t
				}(),

				Name:                         "Demo Inventory",
				Description:                  "",
				Organization:                 1,
				Kind:                         "",
				HostFilter:                   nil,
				Variables:                    "",
				HasActiveFailures:            false,
				TotalHosts:                   2,
				HostsWithActiveFailures:      0,
				TotalGroups:                  0,
				GroupsWithActiveFailures:     0,
				HasInventorySources:          false,
				TotalInventorySources:        0,
				InventorySourcesWithFailures: 0,
				InsightsCredential:           nil,
				PendingDeletion:              false,
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.InventoriesService.ListInventories(map[string]string{
		"name": "Demo Inventory",
	})

	if err != nil {
		t.Fatalf("ListInventories err: %s", err)
	} else if !reflect.DeepEqual(result, expectListInventoriesResponse) {
		t.Logf("expected: %v", expectListInventoriesResponse[0])
		t.Logf("result: %v", result[0])
		t.Fatal("ListInventories resp not as expected")
	} else {
		t.Log("ListInventories passed!")
	}
}

func TestCreateInventory(t *testing.T) {
	var (
		expectCreateInventoryResponse = &Inventory{
			ID:   6,
			Type: "inventory",
			URL:  "/api/v2/inventories/6/",
			Related: &Related{
				NamedURL:       "/api/v2/inventories/TestInventory++Default/",
				CreatedBy:      "/api/v2/users/1/",
				ModifiedBy:     "/api/v2/users/1/",
				JobTemplates:   "/api/v2/inventories/6/job_templates/",
				VariableData:   "/api/v2/inventories/6/variable_data/",
				RootGroups:     "/api/v2/inventories/6/root_groups/",
				ObjectRoles:    "/api/v2/inventories/6/object_roles/",
				AdHocCommands:  "/api/v2/inventories/6/ad_hoc_commands/",
				Script:         "/api/v2/inventories/6/script/",
				Tree:           "/api/v2/inventories/6/tree/",
				AccessList:     "/api/v2/inventories/6/access_list/",
				ActivityStream: "/api/v2/inventories/6/activity_stream/",
				InstanceGroups: "/api/v2/inventories/6/instance_groups/",
				Hosts:          "/api/v2/inventories/6/hosts/",
				Groups:         "/api/v2/inventories/6/groups/",
				Copy:           "/api/v2/inventories/6/copy/",
				UpdateInventorySources: "/api/v2/inventories/6/update_inventory_sources/",
				InventorySources:       "/api/v2/inventories/6/inventory_sources/",
				Organization:           "/api/v2/organizations/1/",
			},
			SummaryFields: &Summary{
				Organization: &OrgnizationSummary{
					ID:          1,
					Name:        "Default",
					Description: "",
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
					UseRole: &ApplyRole{
						ID:          80,
						Description: "Can use the inventory in a job template",
						Name:        "Use",
					},

					AdminRole: &ApplyRole{
						ID:          78,
						Description: "Can manage all aspects of the inventory",
						Name:        "Admin",
					},

					AdhocRole: &ApplyRole{
						ID:          77,
						Description: "May run ad hoc commands on an inventory",
						Name:        "Ad Hoc",
					},

					UpdateRole: &ApplyRole{
						ID:          81,
						Description: "May update project or inventory or group using the configured source update system",
						Name:        "Update",
					},

					ReadRole: &ApplyRole{
						ID:          79,
						Description: "May view settings for the inventory",
						Name:        "Read",
					},
				},

				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Copy:   true,
					Adhoc:  true,
					Delete: true,
				},
			},

			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-08-13T01:59:47.160127Z")
				return t
			}(),

			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-08-13T01:59:47.160140Z")
				return t
			}(),

			Name:                         "TestInventory",
			Description:                  "for testing CreateInventory api",
			Organization:                 1,
			Kind:                         "",
			HostFilter:                   nil,
			Variables:                    "",
			HasActiveFailures:            false,
			TotalHosts:                   0,
			HostsWithActiveFailures:      0,
			TotalGroups:                  0,
			GroupsWithActiveFailures:     0,
			HasInventorySources:          false,
			TotalInventorySources:        0,
			InventorySourcesWithFailures: 0,
			InsightsCredential:           nil,
			PendingDeletion:              false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.InventoriesService.CreateInventory(map[string]interface{}{
		"name":         "TestInventory",
		"description":  "for testing CreateInventory api",
		"organization": 1,
		"kind":         "",
		"host_filter":  "",
		"variables":    "",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("CreateInventory err: %s", err)
	} else if !reflect.DeepEqual(result, expectCreateInventoryResponse) {
		t.Logf("expected: %v", expectCreateInventoryResponse)
		t.Logf("result: %v", result)
		t.Fatal("CreateInventory response not as expected")
	} else {
		t.Log("CreateInventory passed!")
	}
}

func TestUpdateInventory(t *testing.T) {
	var (
		expectUpdateInventoryResponse = &Inventory{
			ID:   6,
			Type: "inventory",
			URL:  "/api/v2/inventories/6/",
			Related: &Related{
				NamedURL:       "/api/v2/inventories/TestInventory-update1++Default/",
				CreatedBy:      "/api/v2/users/1/",
				ModifiedBy:     "/api/v2/users/1/",
				JobTemplates:   "/api/v2/inventories/6/job_templates/",
				VariableData:   "/api/v2/inventories/6/variable_data/",
				RootGroups:     "/api/v2/inventories/6/root_groups/",
				ObjectRoles:    "/api/v2/inventories/6/object_roles/",
				AdHocCommands:  "/api/v2/inventories/6/ad_hoc_commands/",
				Script:         "/api/v2/inventories/6/script/",
				Tree:           "/api/v2/inventories/6/tree/",
				AccessList:     "/api/v2/inventories/6/access_list/",
				ActivityStream: "/api/v2/inventories/6/activity_stream/",
				InstanceGroups: "/api/v2/inventories/6/instance_groups/",
				Hosts:          "/api/v2/inventories/6/hosts/",
				Groups:         "/api/v2/inventories/6/groups/",
				Copy:           "/api/v2/inventories/6/copy/",
				UpdateInventorySources: "/api/v2/inventories/6/update_inventory_sources/",
				InventorySources:       "/api/v2/inventories/6/inventory_sources/",
				Organization:           "/api/v2/organizations/1/",
			},
			SummaryFields: &Summary{
				Organization: &OrgnizationSummary{
					ID:          1,
					Name:        "Default",
					Description: "",
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
					UseRole: &ApplyRole{
						ID:          80,
						Description: "Can use the inventory in a job template",
						Name:        "Use",
					},

					AdminRole: &ApplyRole{
						ID:          78,
						Description: "Can manage all aspects of the inventory",
						Name:        "Admin",
					},

					AdhocRole: &ApplyRole{
						ID:          77,
						Description: "May run ad hoc commands on an inventory",
						Name:        "Ad Hoc",
					},

					UpdateRole: &ApplyRole{
						ID:          81,
						Description: "May update project or inventory or group using the configured source update system",
						Name:        "Update",
					},

					ReadRole: &ApplyRole{
						ID:          79,
						Description: "May view settings for the inventory",
						Name:        "Read",
					},
				},

				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Copy:   true,
					Adhoc:  true,
					Delete: true,
				},
			},

			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-08-13T01:59:47.160127Z")
				return t
			}(),

			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-08-13T01:59:47.160140Z")
				return t
			}(),

			Name:                         "TestInventory-update1",
			Description:                  "for testing UpdateInventory api",
			Organization:                 1,
			Kind:                         "",
			HostFilter:                   nil,
			Variables:                    "",
			HasActiveFailures:            false,
			TotalHosts:                   0,
			HostsWithActiveFailures:      0,
			TotalGroups:                  0,
			GroupsWithActiveFailures:     0,
			HasInventorySources:          false,
			TotalInventorySources:        0,
			InventorySourcesWithFailures: 0,
			InsightsCredential:           nil,
			PendingDeletion:              false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.InventoriesService.UpdateInventory(6, map[string]interface{}{
		"name":         "TestInventory-update1",
		"description":  "for testing UpdateInventory api",
		"organization": 1,
		"kind":         "",
		"host_filter":  "",
		"variables":    "",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("UpdateInventory err: %s", err)
	} else if !reflect.DeepEqual(result, expectUpdateInventoryResponse) {
		t.Logf("expected: %v", expectUpdateInventoryResponse)
		t.Logf("result: %v", result)
		t.Fatal("UpdateInventory resp not as expected")
	} else {
		t.Log("UpdateInventory passed!")
	}
}

func TestGetInventory(t *testing.T) {
	var (
		expectGetInventoryResponse = &Inventory{
			ID:   1,
			Type: "inventory",
			URL:  "/api/v2/inventories/1/",
			Related: &Related{
				CreatedBy:      "/api/v2/users/1/",
				ModifiedBy:     "/api/v2/users/1/",
				JobTemplates:   "/api/v2/inventories/1/job_templates/",
				VariableData:   "/api/v2/inventories/1/variable_data/",
				RootGroups:     "/api/v2/inventories/1/root_groups/",
				ObjectRoles:    "/api/v2/inventories/1/object_roles/",
				AdHocCommands:  "/api/v2/inventories/1/ad_hoc_commands/",
				Script:         "/api/v2/inventories/1/script/",
				Tree:           "/api/v2/inventories/1/tree/",
				AccessList:     "/api/v2/inventories/1/access_list/",
				ActivityStream: "/api/v2/inventories/1/activity_stream/",
				InstanceGroups: "/api/v2/inventories/1/instance_groups/",
				Hosts:          "/api/v2/inventories/1/hosts/",
				Groups:         "/api/v2/inventories/1/groups/",
				Copy:           "/api/v2/inventories/1/copy/",
				UpdateInventorySources: "/api/v2/inventories/1/update_inventory_sources/",
				InventorySources:       "/api/v2/inventories/1/inventory_sources/",
				Organization:           "/api/v2/organizations/1/",
			},
			SummaryFields: &Summary{
				Organization: &OrgnizationSummary{
					ID:          1,
					Name:        "Default",
					Description: "",
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
					UseRole: &ApplyRole{
						ID:          80,
						Description: "Can use the inventory in a job template",
						Name:        "Use",
					},

					AdminRole: &ApplyRole{
						ID:          78,
						Description: "Can manage all aspects of the inventory",
						Name:        "Admin",
					},

					AdhocRole: &ApplyRole{
						ID:          77,
						Description: "May run ad hoc commands on an inventory",
						Name:        "Ad Hoc",
					},

					UpdateRole: &ApplyRole{
						ID:          81,
						Description: "May update project or inventory or group using the configured source update system",
						Name:        "Update",
					},

					ReadRole: &ApplyRole{
						ID:          79,
						Description: "May view settings for the inventory",
						Name:        "Read",
					},
				},

				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Copy:   true,
					Adhoc:  true,
					Delete: true,
				},
			},

			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-05-21T01:34:35.657185Z")
				return t
			}(),

			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-05-30T09:42:22.412749Z")
				return t
			}(),

			Name:                         "Demo Inventory",
			Description:                  "",
			Organization:                 1,
			Kind:                         "",
			HostFilter:                   nil,
			Variables:                    "",
			HasActiveFailures:            false,
			TotalHosts:                   2,
			HostsWithActiveFailures:      0,
			TotalGroups:                  0,
			GroupsWithActiveFailures:     0,
			HasInventorySources:          false,
			TotalInventorySources:        0,
			InventorySourcesWithFailures: 0,
			InsightsCredential:           nil,
			PendingDeletion:              false,
		}
	)
	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.InventoriesService.GetInventory(1, map[string]string{})

	if err != nil {
		t.Fatalf("GetInventory by ID err: %s", err)
	} else if !reflect.DeepEqual(result, expectGetInventoryResponse) {
		t.Logf("expected: %v", expectGetInventoryResponse)
		t.Logf("result: %v", result)
		t.Fatal("GetInventory by ID resp not as expected")
	} else {
		t.Log("GetInventory by ID passed!")
	}
}

func TestDeleteInventory(t *testing.T) {
	var (
		expectDeleteInventoryResponse = &Inventory{}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.InventoriesService.DeleteInventory(1)
	if err != nil {
		t.Fatalf("DeleteInventory by ID err: %s", err)
	} else if !reflect.DeepEqual(result, expectDeleteInventoryResponse) {
		t.Logf("expected: %v", expectDeleteInventoryResponse)
		t.Logf("result: %v", result)
		t.Fatal("DeleteInventory resp not as expected")
	} else {
		t.Log("DeleteInventory passed!")
	}
}
