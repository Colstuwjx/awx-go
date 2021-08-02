package awx

import (
	"testing"
	"time"
)

func TestGetInventoryUpdate(t *testing.T) {
	var (
		expectInventoryUpdateResponse = &InventoryUpdate{
			Status: "successful",
			ID:     305,
			Type:   "inventory_update",
			URL:    "/api/v2/inventory_updates/305/",
			Related: &Related{
				CreatedBy:          "/api/v2/users/5/",
				ModifiedBy:         "/api/v2/users/5/",
				UnifiedJobTemplate: "/api/v2/inventory_sources/10/",
				Stdout:             "/api/v2/inventory_updates/305/stdout/",
				InventorySource:    "/api/v2/inventory_sources/10/",
				Cancel:             "/api/v2/inventory_updates/305/cancel/",
				Notifications:      "/api/v2/inventory_updates/305/notifications/",
				Events:             "/api/v2/inventory_updates/305/events/",
				Inventory:          "/api/v2/inventories/1/",
				Credentials:        "/api/v2/inventory_updates/305/credentials/",
			},
			SummaryFields: &Summary{
				Organization: &OrgnizationSummary{
					ID:          1,
					Name:        "Default",
					Description: "",
				},
				Inventory: &Inventory{
					ID:                           1,
					Name:                         "Default",
					Description:                  "",
					HasActiveFailures:            true,
					TotalHosts:                   7,
					HostsWithActiveFailures:      6,
					TotalGroups:                  14,
					HasInventorySources:          true,
					TotalInventorySources:        1,
					InventorySourcesWithFailures: 0,
					OrganizationID:               1,
					Kind:                         "",
				},
				UnifiedJobTemplate: &UnifiedJobTemplate{
					ID:             10,
					Name:           "Default",
					Description:    "",
					UnifiedJobType: "inventory_update",
				},
				InventorySource: &InventorySource{
					Source: "scm",
					LastUpdated: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2021-07-30T10:12:44.553099Z")
						return t
					}(),
					Status: "pending",
				},
				CreatedBy: &ByUserSummary{
					ID:        5,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},
				ModifiedBy: &ByUserSummary{
					ID:        5,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},
				UserCapabilities: &UserCapabilities{
					Delete: true,
					Start:  true,
				},
				Credentials: []Credential{},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2021-08-02T01:45:22.144755Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2021-08-02T01:45:22.157220Z")
				return t
			}(),
			Name:                    "Default",
			Description:             "",
			Source:                  "scm",
			SourcePath:              "",
			SourceScript:            nil,
			SourceVars:              "",
			Credential:              nil,
			EnabledVar:              "",
			EnabledValue:            "",
			HostFilter:              "",
			Overwrite:               true,
			OverwriteVars:           true,
			CustomVirtualenv:        nil,
			Timeout:                 0,
			Verbosity:               2,
			UnifiedJobTemplate:      10,
			LaunchType:              "manual",
			Failed:                  false,
			Started:                 nil,
			Finished:                nil,
			CanceledOn:              nil,
			Elapsed:                 0.0,
			JobArgs:                 "",
			JobCwd:                  "",
			JobEnv:                  nil,
			JobExplanation:          "",
			ExecutionNode:           "",
			ResultTraceback:         "",
			EventProcessingFinished: false,
			Inventory:               1,
			LicenseError:            false,
			OrgHostLimitError:       false,
			SourceProjectUpdate:     nil,
			SourceProject:           nil,
			InventoryUpdate:         305,
		}
	)
	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.InventoryUpdatesService.GetInventoryUpdate(305, map[string]string{})

	if err != nil {
		t.Fatalf("GetInventoryUpdate err: %s", err)
	} else {
		checkAPICallResult(t, expectInventoryUpdateResponse, result)
		t.Log("GetInventoryUpdate passed!")
	}
}
