package awx

//
//import (
//	"testing"
//	"time"
//)
//
//func TestListTeams(t *testing.T) {
//	var (
//		expectListTeamsResponse = []*Team{
//			{
//				ID:   1,
//				Type: "team",
//				URL:  "/api/v2/teams/1/",
//				Related: &Related{
//					CreatedBy:      "/api/v2/users/4/",
//					ModifiedBy:     "/api/v2/users/4/",
//					Users:          "/api/v2/teams/1/users/",
//					Roles:          "/api/v2/teams/1/roles/",
//					ObjectRoles:    "/api/v2/teams/1/object_roles/",
//					Credentials:    "/api/v2/teams/1/credentials/",
//					Projects:       "/api/v2/teams/1/projects/",
//					ActivityStream: "/api/v2/teams/1/activity_stream/",
//					AccessList:     "/api/v2/teams/1/access_list/",
//					Organization:   "/api/v2/organizations/1/",
//				},
//				SummaryFields: &Summary{
//					UserCapabilities: &UserCapabilities{
//						Edit:   true,
//						Delete: true,
//					},
//					Organization: &OrgnizationSummary{
//						ID:          1,
//						Name:        "Default",
//						Description: "",
//					},
//					ObjectRoles: &ObjectRoles{
//						AdminRole: &ApplyRole{
//							ID:          30,
//							Description: "Can manage all aspects of the team",
//							Name:        "Admin",
//						},
//
//						MemberRole: &ApplyRole{
//							ID:          29,
//							Description: "User is a member of the team",
//							Name:        "Member",
//						},
//
//						ReadRole: &ApplyRole{
//							ID:          31,
//							Description: "May view settings for the team",
//							Name:        "Read",
//						},
//					},
//					CreatedBy: &ByUserSummary{
//						ID:        4,
//						Username:  "admin",
//						FirstName: "",
//						LastName:  "",
//					},
//					ModifiedBy: &ByUserSummary{
//						ID:        4,
//						Username:  "admin",
//						FirstName: "",
//						LastName:  "",
//					},
//				},
//				Created: func() time.Time {
//					t, _ := time.Parse(time.RFC3339, "2018-12-01T13:33:30.692904Z")
//					return t
//				}(),
//				Modified: func() time.Time {
//					t, _ := time.Parse(time.RFC3339, "2018-12-01T13:33:30.692904Z")
//					return t
//				}(),
//				Name:         "test-team",
//				Organization: 1,
//				Description:  "",
//			},
//		}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, _, err := awx.TeamService.ListTeams(map[string]string{
//		"name": "test-team",
//	})
//
//	if err != nil {
//		t.Fatalf("ListTeams err: %s", err)
//	} else {
//		checkAPICallResult(t, expectListTeamsResponse, result)
//		t.Log("ListTeams passed!")
//	}
//}
//
//func TestCreateTeam(t *testing.T) {
//	var (
//		expectCreateTeamResponse = &Team{
//			ID:   1,
//			Type: "team",
//			URL:  "/api/v2/teams/1/",
//			Related: &Related{
//				CreatedBy:      "/api/v2/users/4/",
//				ModifiedBy:     "/api/v2/users/4/",
//				Users:          "/api/v2/teams/1/users/",
//				Roles:          "/api/v2/teams/1/roles/",
//				ObjectRoles:    "/api/v2/teams/1/object_roles/",
//				Credentials:    "/api/v2/teams/1/credentials/",
//				Projects:       "/api/v2/teams/1/projects/",
//				ActivityStream: "/api/v2/teams/1/activity_stream/",
//				AccessList:     "/api/v2/teams/1/access_list/",
//				Organization:   "/api/v2/organizations/1/",
//			},
//			SummaryFields: &Summary{
//				UserCapabilities: &UserCapabilities{
//					Edit:   true,
//					Delete: true,
//				},
//				Organization: &OrgnizationSummary{
//					ID:          1,
//					Name:        "Default",
//					Description: "",
//				},
//				ObjectRoles: &ObjectRoles{
//					AdminRole: &ApplyRole{
//						ID:          30,
//						Description: "Can manage all aspects of the team",
//						Name:        "Admin",
//					},
//
//					MemberRole: &ApplyRole{
//						ID:          29,
//						Description: "User is a member of the team",
//						Name:        "Member",
//					},
//
//					ReadRole: &ApplyRole{
//						ID:          31,
//						Description: "May view settings for the team",
//						Name:        "Read",
//					},
//				},
//				CreatedBy: &ByUserSummary{
//					ID:        4,
//					Username:  "admin",
//					FirstName: "",
//					LastName:  "",
//				},
//				ModifiedBy: &ByUserSummary{
//					ID:        4,
//					Username:  "admin",
//					FirstName: "",
//					LastName:  "",
//				},
//			},
//			Created: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-01T13:33:30.692904Z")
//				return t
//			}(),
//			Modified: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-01T13:33:30.692904Z")
//				return t
//			}(),
//			Name:         "test-team",
//			Organization: 1,
//			Description:  "",
//		}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, err := awx.TeamService.CreateTeam(map[string]interface{}{
//		"name":         "test-team",
//		"organization": 1,
//	}, map[string]string{})
//	if err != nil {
//		t.Fatalf("CreateTeam err: %s", err)
//	} else {
//		checkAPICallResult(t, expectCreateTeamResponse, result)
//		t.Log("CreateTeam passed!")
//	}
//}
//
//func TestUpdateTeam(t *testing.T) {
//	var (
//		expectUpdateTeamResponse = &Team{
//			ID:   1,
//			Type: "team",
//			URL:  "/api/v2/teams/1/",
//			Related: &Related{
//				CreatedBy:      "/api/v2/users/4/",
//				ModifiedBy:     "/api/v2/users/4/",
//				Users:          "/api/v2/teams/1/users/",
//				Roles:          "/api/v2/teams/1/roles/",
//				ObjectRoles:    "/api/v2/teams/1/object_roles/",
//				Credentials:    "/api/v2/teams/1/credentials/",
//				Projects:       "/api/v2/teams/1/projects/",
//				ActivityStream: "/api/v2/teams/1/activity_stream/",
//				AccessList:     "/api/v2/teams/1/access_list/",
//				Organization:   "/api/v2/organizations/1/",
//			},
//			SummaryFields: &Summary{
//				UserCapabilities: &UserCapabilities{
//					Edit:   true,
//					Delete: true,
//				},
//				Organization: &OrgnizationSummary{
//					ID:          1,
//					Name:        "Default",
//					Description: "",
//				},
//				ObjectRoles: &ObjectRoles{
//					AdminRole: &ApplyRole{
//						ID:          30,
//						Description: "Can manage all aspects of the team",
//						Name:        "Admin",
//					},
//
//					MemberRole: &ApplyRole{
//						ID:          29,
//						Description: "User is a member of the team",
//						Name:        "Member",
//					},
//
//					ReadRole: &ApplyRole{
//						ID:          31,
//						Description: "May view settings for the team",
//						Name:        "Read",
//					},
//				},
//				CreatedBy: &ByUserSummary{
//					ID:        4,
//					Username:  "admin",
//					FirstName: "",
//					LastName:  "",
//				},
//				ModifiedBy: &ByUserSummary{
//					ID:        4,
//					Username:  "admin",
//					FirstName: "",
//					LastName:  "",
//				},
//			},
//			Created: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-01T13:33:30.692904Z")
//				return t
//			}(),
//			Modified: func() time.Time {
//				t, _ := time.Parse(time.RFC3339, "2018-12-01T13:33:30.692904Z")
//				return t
//			}(),
//			Name:         "test-team",
//			Organization: 1,
//			Description:  "Update test-team",
//		}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, err := awx.TeamService.UpdateTeam(4, map[string]interface{}{
//		"name":         "test-team",
//		"organization": 1,
//		"description":  "Update test-team",
//	}, map[string]string{})
//	if err != nil {
//		t.Fatalf("CreateTeam err: %s", err)
//	} else {
//		checkAPICallResult(t, expectUpdateTeamResponse, result)
//		t.Log("CreateTeam passed!")
//	}
//}
//func TestDeleteTeam(t *testing.T) {
//	var (
//		expectDeleteTeamResponse = &Team{}
//	)
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	result, err := awx.TeamService.DeleteTeam(1)
//
//	if err != nil {
//		t.Fatalf("DeleteTeam err: %s", err)
//	} else {
//		checkAPICallResult(t, expectDeleteTeamResponse, result)
//		t.Log("DeleteTeam passed!")
//	}
//}
//
//func TestTeamGrantRole(t *testing.T) {
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	err := awx.TeamService.GrantRole(4, 170)
//
//	if err != nil {
//		t.Fatalf("TestTeamGrantRole err: %s", err)
//	} else {
//		checkAPICallResult(t, nil, nil)
//		t.Log("TestTeamGrantRole passed!")
//	}
//}
//
//func TestTeamRevokeRole(t *testing.T) {
//
//	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
//	err := awx.TeamService.RevokeRole(1, 170)
//
//	if err != nil {
//		t.Fatalf("TestTeamRevokeRole err: %s", err)
//	} else {
//		checkAPICallResult(t, nil, nil)
//		t.Log("TestTeamRevokeRole passed!")
//	}
//}
