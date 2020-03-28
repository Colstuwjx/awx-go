package awx

//
//import (
//	"testing"
//	"time"
//)
//
//func TestListCredentials(t *testing.T) {
//	var (
//		expectListCredentialsResponse = []*Credential{
//			{
//				ID:   1,
//				Type: "credential",
//				URL:  "/api/v2/credentials/1/",
//				Related: &Related{
//					ObjectRoles:    "/api/v2/credentials/1/object_roles/",
//					AccessList:     "/api/v2/credentials/1/access_list/",
//					CredentialType: "/api/v2/credential_types/1/",
//					ModifiedBy:     "/api/v2/users/4/",
//					OwnerUsers:     "/api/v2/credentials/1/owner_users/",
//					OwnerTeams:     "/api/v2/credentials/1/owner_teams/",
//					Organization:   "/api/v2/organizations/1/",
//					Copy:           "/api/v2/credentials/1/copy/",
//					ActivityStream: "/api/v2/credentials/1/activity_stream/",
//				},
//				SummaryFields: &Summary{
//					Organization: &OrgnizationSummary{
//						ID:          1,
//						Name:        "Default",
//						Description: "",
//					},
//					Project: &Project{},
//					ModifiedBy: &ByUserSummary{
//						ID:        4,
//						Username:  "admin",
//						FirstName: "",
//						LastName:  "",
//					},
//					ObjectRoles: &ObjectRoles{
//						AdminRole: &ApplyRole{
//							ID:          18,
//							Description: "Can manage all aspects of the credential",
//							Name:        "Admin",
//						},
//						UseRole: &ApplyRole{
//							ID:          20,
//							Description: "Can use the credential in a job template",
//							Name:        "Use",
//						},
//						ReadRole: &ApplyRole{
//							ID:          19,
//							Description: "May view settings for the credential",
//							Name:        "Read",
//						},
//					},
//					UserCapabilities: &UserCapabilities{
//						Edit:   true,
//						Delete: true,
//						Copy:   true,
//					},
//				},
//				Created: func() time.Time {
//					t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.496424Z")
//					return t
//				}(),
//				Modified: func() time.Time {
//					t, _ := time.Parse(time.RFC3339, "2018-12-07T16:17:48.131210Z")
//					return t
//				}(),
//				Name:           "Demo Credential",
//				Description:    "",
//				Organization:   1,
//				CredentialType: 1,
//				Inputs: map[string]interface{}{
//					"username": "admin",
//				},
//			},
//		}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, _, err := awx.CredentialService.ListCredentials(map[string]string{
//		"name": "Default",
//	})
//
//	if err != nil {
//		t.Fatalf("ListCredentials err: %s", err)
//	} else {
//		checkAPICallResult(t, expectListCredentialsResponse, result)
//		t.Log("ListCredentials passed!")
//	}
//}
//
//func TestCreateCredentials(t *testing.T) {
//	var (
//		expectCreateCredentialsResponse = &Credential{
//			ID:   1,
//			Type: "credential",
//			URL:  "/api/v2/credentials/1/",
//			Related: &Related{
//				ObjectRoles:    "/api/v2/credentials/1/object_roles/",
//				AccessList:     "/api/v2/credentials/1/access_list/",
//				CredentialType: "/api/v2/credential_types/1/",
//				ModifiedBy:     "/api/v2/users/4/",
//				OwnerUsers:     "/api/v2/credentials/1/owner_users/",
//				OwnerTeams:     "/api/v2/credentials/1/owner_teams/",
//				Organization:   "/api/v2/organizations/1/",
//				Copy:           "/api/v2/credentials/1/copy/",
//				ActivityStream: "/api/v2/credentials/1/activity_stream/",
//			},
//			SummaryFields: &Summary{
//				Organization: &OrgnizationSummary{
//					ID:          1,
//					Name:        "Default",
//					Description: "",
//				},
//				Project: &Project{},
//				ModifiedBy: &ByUserSummary{
//					ID:        4,
//					Username:  "admin",
//					FirstName: "",
//					LastName:  "",
//				},
//				ObjectRoles: &ObjectRoles{
//					AdminRole: &ApplyRole{
//						ID:          18,
//						Description: "Can manage all aspects of the credential",
//						Name:        "Admin",
//					},
//					UseRole: &ApplyRole{
//						ID:          20,
//						Description: "Can use the credential in a job template",
//						Name:        "Use",
//					},
//					ReadRole: &ApplyRole{
//						ID:          19,
//						Description: "May view settings for the credential",
//						Name:        "Read",
//					},
//				},
//				UserCapabilities: &UserCapabilities{
//					Edit:   true,
//					Delete: true,
//					Copy:   true,
//				},
//			},
//			Created: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.496424Z")
//				return t
//			}(),
//			Modified: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-07T16:17:48.131210Z")
//				return t
//			}(),
//			Name:           "Demo Credential",
//			Description:    "",
//			Organization:   1,
//			CredentialType: 1,
//			Inputs: map[string]interface{}{
//				"username": "admin",
//			},
//		}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, err := awx.CredentialService.CreateCredential(map[string]interface{}{
//		"name":            "Demo Credential",
//		"organization":    1,
//		"credential_type": 1,
//	}, map[string]string{})
//	if err != nil {
//		t.Fatalf("CreateCredentials err: %s", err)
//	} else {
//		checkAPICallResult(t, expectCreateCredentialsResponse, result)
//		t.Log("CreateCredentials passed!")
//	}
//}
//
//func TestUpdateCredentials(t *testing.T) {
//	var (
//		expectUpdateCredentialsResponse = &Credential{
//			ID:   1,
//			Type: "credential",
//			URL:  "/api/v2/credentials/1/",
//			Related: &Related{
//				ObjectRoles:    "/api/v2/credentials/1/object_roles/",
//				AccessList:     "/api/v2/credentials/1/access_list/",
//				CredentialType: "/api/v2/credential_types/1/",
//				ModifiedBy:     "/api/v2/users/4/",
//				OwnerUsers:     "/api/v2/credentials/1/owner_users/",
//				OwnerTeams:     "/api/v2/credentials/1/owner_teams/",
//				Organization:   "/api/v2/organizations/1/",
//				Copy:           "/api/v2/credentials/1/copy/",
//				ActivityStream: "/api/v2/credentials/1/activity_stream/",
//			},
//			SummaryFields: &Summary{
//				Organization: &OrgnizationSummary{
//					ID:          1,
//					Name:        "Default",
//					Description: "",
//				},
//				Project: &Project{},
//				ModifiedBy: &ByUserSummary{
//					ID:        4,
//					Username:  "admin",
//					FirstName: "",
//					LastName:  "",
//				},
//				ObjectRoles: &ObjectRoles{
//					AdminRole: &ApplyRole{
//						ID:          18,
//						Description: "Can manage all aspects of the credential",
//						Name:        "Admin",
//					},
//					UseRole: &ApplyRole{
//						ID:          20,
//						Description: "Can use the credential in a job template",
//						Name:        "Use",
//					},
//					ReadRole: &ApplyRole{
//						ID:          19,
//						Description: "May view settings for the credential",
//						Name:        "Read",
//					},
//				},
//				UserCapabilities: &UserCapabilities{
//					Edit:   true,
//					Delete: true,
//					Copy:   true,
//				},
//			},
//			Created: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-01T12:10:00.496424Z")
//				return t
//			}(),
//			Modified: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-07T16:17:48.131210Z")
//				return t
//			}(),
//			Name:           "Demo Credential",
//			Description:    "Demo Credential",
//			Organization:   1,
//			CredentialType: 1,
//			Inputs: map[string]interface{}{
//				"username": "admin",
//			},
//		}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, err := awx.CredentialService.UpdateCredential(1, map[string]interface{}{
//		"name":            "Demo Credential",
//		"description":     "Demo credential",
//		"organization":    1,
//		"credential_type": 1,
//		"inputs": map[string]interface{}{
//			"username": "admin",
//		},
//	}, map[string]string{})
//	if err != nil {
//		t.Fatalf("CreateCredentials err: %s", err)
//	} else {
//		checkAPICallResult(t, expectUpdateCredentialsResponse, result)
//		t.Log("CreateCredentials passed!")
//	}
//}
//
//func TestDeleteCredentials(t *testing.T) {
//	var (
//		expectDeleteCredentialsResponse = &Credential{}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, err := awx.CredentialService.DeleteCredential(1)
//
//	if err != nil {
//		t.Fatalf("DeleteCredentials err: %s", err)
//	} else {
//		checkAPICallResult(t, expectDeleteCredentialsResponse, result)
//		t.Log("DeleteCredentials passed!")
//	}
//}
