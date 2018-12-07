package mockdata

var (
	// MockedListCredentialResponse mocked `ListCredentials` endpoint response
	MockedListCredentialResponse = []byte(`
        {
            "count": 1,
            "next": null,
            "previous": null,
            "results": [
                {
                    "id": 1,
                    "type": "credential",
                    "url": "/api/v2/credentials/1/",
                    "related": {
                        "modified_by": "/api/v2/users/4/",
                        "organization": "/api/v2/organizations/1/",
                        "owner_teams": "/api/v2/credentials/1/owner_teams/",
                        "owner_users": "/api/v2/credentials/1/owner_users/",
                        "activity_stream": "/api/v2/credentials/1/activity_stream/",
                        "access_list": "/api/v2/credentials/1/access_list/",
                        "object_roles": "/api/v2/credentials/1/object_roles/",
                        "copy": "/api/v2/credentials/1/copy/",
                        "credential_type": "/api/v2/credential_types/1/"
                    },
                    "summary_fields": {
                        "host": {},
                        "project": {},
                        "organization": {
                            "id": 1,
                            "name": "Default",
                            "description": ""
                        },
                        "modified_by": {
                            "id": 4,
                            "username": "admin",
                            "first_name": "",
                            "last_name": ""
                        },
                        "object_roles": {
                            "admin_role": {
                                "id": 18,
                                "description": "Can manage all aspects of the credential",
                                "name": "Admin"
                            },
                            "use_role": {
                                "id": 20,
                                "description": "Can use the credential in a job template",
                                "name": "Use"
                            },
                            "read_role": {
                                "id": 19,
                                "description": "May view settings for the credential",
                                "name": "Read"
                            }
                        },
                        "user_capabilities": {
                            "edit": true,
                            "copy": true,
                            "delete": true
                        },
                        "owners": [
                            {
                                "url": "/api/v2/organizations/1/",
                                "description": "",
                                "type": "organization",
                                "id": 1,
                                "name": "Default"
                            }
                        ]
                    },
                    "created": "2018-12-01T12:10:00.496424Z",
                    "modified": "2018-12-07T16:17:48.131210Z",
                    "name": "Demo Credential",
                    "description": "",
                    "organization": 1,
                    "credential_type": 1,
                    "inputs": {
                        "username": "admin"
                    }
                }
            ]
    }`)
	// MockedCreateCredentialResponse mocked `CreateCredential` endpoint response
	MockedCreateCredentialResponse = []byte(`
                {
                    "id": 1,
                    "type": "credential",
                    "url": "/api/v2/credentials/1/",
                    "related": {
                        "modified_by": "/api/v2/users/4/",
                        "organization": "/api/v2/organizations/1/",
                        "owner_teams": "/api/v2/credentials/1/owner_teams/",
                        "owner_users": "/api/v2/credentials/1/owner_users/",
                        "activity_stream": "/api/v2/credentials/1/activity_stream/",
                        "access_list": "/api/v2/credentials/1/access_list/",
                        "object_roles": "/api/v2/credentials/1/object_roles/",
                        "copy": "/api/v2/credentials/1/copy/",
                        "credential_type": "/api/v2/credential_types/1/"
                    },
                    "summary_fields": {
                        "host": {},
                        "project": {},
                        "organization": {
                            "id": 1,
                            "name": "Default",
                            "description": ""
                        },
                        "modified_by": {
                            "id": 4,
                            "username": "admin",
                            "first_name": "",
                            "last_name": ""
                        },
                        "object_roles": {
                            "admin_role": {
                                "id": 18,
                                "description": "Can manage all aspects of the credential",
                                "name": "Admin"
                            },
                            "use_role": {
                                "id": 20,
                                "description": "Can use the credential in a job template",
                                "name": "Use"
                            },
                            "read_role": {
                                "id": 19,
                                "description": "May view settings for the credential",
                                "name": "Read"
                            }
                        },
                        "user_capabilities": {
                            "edit": true,
                            "copy": true,
                            "delete": true
                        },
                        "owners": [
                            {
                                "url": "/api/v2/organizations/1/",
                                "description": "",
                                "type": "organization",
                                "id": 1,
                                "name": "Default"
                            }
                        ]
                    },
                    "created": "2018-12-01T12:10:00.496424Z",
                    "modified": "2018-12-07T16:17:48.131210Z",
                    "name": "Demo Credential",
                    "description": "",
                    "organization": 1,
                    "credential_type": 1,
                    "inputs": {
                        "username": "admin"
                    }
    }`)
	// MockedUpdateCredentialResponse mocked `UpdateCredential` endpoint response
	MockedUpdateCredentialResponse = []byte(`
        {
            "id": 1,
            "type": "credential",
            "url": "/api/v2/credentials/1/",
            "related": {
                "modified_by": "/api/v2/users/4/",
                "organization": "/api/v2/organizations/1/",
                "owner_teams": "/api/v2/credentials/1/owner_teams/",
                "owner_users": "/api/v2/credentials/1/owner_users/",
                "activity_stream": "/api/v2/credentials/1/activity_stream/",
                "access_list": "/api/v2/credentials/1/access_list/",
                "object_roles": "/api/v2/credentials/1/object_roles/",
                "copy": "/api/v2/credentials/1/copy/",
                "credential_type": "/api/v2/credential_types/1/"
            },
            "summary_fields": {
                "host": {},
                "project": {},
                "organization": {
                    "id": 1,
                    "name": "Default",
                    "description": ""
                },
                "modified_by": {
                    "id": 4,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "object_roles": {
                    "admin_role": {
                        "id": 18,
                        "description": "Can manage all aspects of the credential",
                        "name": "Admin"
                    },
                    "use_role": {
                        "id": 20,
                        "description": "Can use the credential in a job template",
                        "name": "Use"
                    },
                    "read_role": {
                        "id": 19,
                        "description": "May view settings for the credential",
                        "name": "Read"
                    }
                },
                "user_capabilities": {
                    "edit": true,
                    "copy": true,
                    "delete": true
                },
                "owners": [
                    {
                        "url": "/api/v2/organizations/1/",
                        "description": "",
                        "type": "organization",
                        "id": 1,
                        "name": "Default"
                    }
                ]
            },
            "created": "2018-12-01T12:10:00.496424Z",
            "modified": "2018-12-07T16:17:48.131210Z",
            "name": "Demo Credential",
            "description": "Demo Credential",
            "organization": 1,
            "credential_type": 1,
            "inputs": {
                "username": "admin"
            }
}`)
	// MockedDeleteCredentialResponse mocked `DeleteCredential` endpoint response
	MockedDeleteCredentialResponse = []byte(`{}`)
)
