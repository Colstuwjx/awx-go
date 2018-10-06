package mockdata

var (
	// MockedListProjectsResponse mocked `ListProjects` endpoint response
	MockedListProjectsResponse = []byte(`
    {
        "count": 1,
        "next": null,
        "previous": null,
        "results": [{
            "id": 4,
            "type": "project",
            "url": "/api/v2/projects/4/",
            "related": {
                "created_by": "/api/v2/users/1/",
                "modified_by": "/api/v2/users/1/",
                "object_roles": "/api/v2/projects/4/object_roles/",
                "copy": "/api/v2/projects/4/copy/",
                "access_list": "/api/v2/projects/4/access_list/",
                "schedules": "/api/v2/projects/4/schedules/",
                "organization": "/api/v2/organizations/1/"
            },
            "summary_fields": {
                "organization": {
                    "id": 1,
                    "name": "Default",
                    "description": ""
                },
                "created_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "admin",
                    "last_name": "admin"
                },
                "modified_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "admin",
                    "last_name": "admin"
                },
                "object_roles": {
                    "admin_role": {
                        "id": 14,
                        "description": "Can manage all aspects of the project",
                        "name": "Admin"
                    },
                    "use_role": {
                        "id": 16,
                        "description": "Can manage all aspects of the project",
                        "name": "Use"
                    },
                    "update_role": {
                        "id": 17,
                        "description": "May update project or inventory or group using the configured source update system",
                        "name": "Update"
                    },
                    "read_role": {
                        "id": 15,
                        "description": "May view settings for the project",
                        "name": "Read"
                    }
                },
                "user_capabilities": {
                    "edit": true,
                    "start": true,
                    "copy": true,
                    "schedule": true,
                    "delete": true
                }
            },
            "created": "2018-06-28T16:31:15.809617Z",
            "modified": "2018-06-28T16:31:15.923732Z",
            "name": "Demo Project",
            "description": "",
            "local_path": "",
            "scm_type": "git",
            "scm_url": "https://github.com/ansible/ansible-tower-samples",
            "scm_clean": false,
            "scm_delete_on_update": false,
            "status": "never updated",
            "scm_delete_on_next_update": false,
            "scm_update_on_launch": true,
            "scm_update_cache_timeout": 0,
            "last_update_failed": false
        }]
    }`)

	// MockedCreateProjectResponse mocked `CreateProject` endpoint response
	MockedCreateProjectResponse = []byte(`
    {
        "id": 4,
        "type": "project",
        "url": "/api/v2/projects/4/",
        "related": {
            "created_by": "/api/v2/users/1/",
            "modified_by": "/api/v2/users/1/",
            "object_roles": "/api/v2/projects/4/object_roles/",
            "copy": "/api/v2/projects/4/copy/",
            "access_list": "/api/v2/projects/4/access_list/",
            "schedules": "/api/v2/projects/4/schedules/",
            "organization": "/api/v2/organizations/1/"
        },
        "summary_fields": {
            "organization": {
                "id": 1,
                "name": "Default",
                "description": ""
            },
            "created_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "modified_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "object_roles": {
                "admin_role": {
                    "id": 14,
                    "description": "Can manage all aspects of the project",
                    "name": "Admin"
                },
                "use_role": {
                    "id": 16,
                    "description": "Can manage all aspects of the project",
                    "name": "Use"
                },
                "update_role": {
                    "id": 17,
                    "description": "May update project or inventory or group using the configured source update system",
                    "name": "Update"
                },
                "read_role": {
                    "id": 15,
                    "description": "May view settings for the project",
                    "name": "Read"
                }
            },
            "user_capabilities": {
                "edit": true,
                "start": true,
                "copy": true,
                "schedule": true,
                "delete": true
            }
        },
        "created": "2018-06-28T16:31:15.809617Z",
        "modified": "2018-06-28T16:31:15.923732Z",
        "name": "Demo Project",
        "description": "",
        "local_path": "",
        "scm_type": "git",
        "scm_url": "https://github.com/ansible/ansible-tower-samples",
        "scm_clean": false,
        "scm_delete_on_update": false,
        "status": "never updated",
        "scm_delete_on_next_update": false,
        "scm_update_on_launch": true,
        "scm_update_cache_timeout": 0,
        "last_update_failed": false
    }`)
	// MockedUpdateProjectResponse mocked `UpdateProject` endpoint response
	MockedUpdateProjectResponse = []byte(`
    {
        "id": 4,
        "type": "project",
        "url": "/api/v2/projects/4/",
        "related": {
            "created_by": "/api/v2/users/1/",
            "modified_by": "/api/v2/users/1/",
            "object_roles": "/api/v2/projects/4/object_roles/",
            "copy": "/api/v2/projects/4/copy/",
            "access_list": "/api/v2/projects/4/access_list/",
            "schedules": "/api/v2/projects/4/schedules/",
            "organization": "/api/v2/organizations/1/"
        },
        "summary_fields": {
            "organization": {
                "id": 1,
                "name": "Default",
                "description": ""
            },
            "created_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "modified_by": {
                "id": 1,
                "username": "admin",
                "first_name": "admin",
                "last_name": "admin"
            },
            "object_roles": {
                "admin_role": {
                    "id": 14,
                    "description": "Can manage all aspects of the project",
                    "name": "Admin"
                },
                "use_role": {
                    "id": 16,
                    "description": "Can manage all aspects of the project",
                    "name": "Use"
                },
                "update_role": {
                    "id": 17,
                    "description": "May update project or inventory or group using the configured source update system",
                    "name": "Update"
                },
                "read_role": {
                    "id": 15,
                    "description": "May view settings for the project",
                    "name": "Read"
                }
            },
            "user_capabilities": {
                "edit": true,
                "start": true,
                "copy": true,
                "schedule": true,
                "delete": true
            }
        },
        "created": "2018-06-28T16:31:15.809617Z",
        "modified": "2018-06-28T16:31:15.923732Z",
        "name": "Demo Project",
        "description": "Test Update",
        "local_path": "",
        "scm_type": "git",
        "scm_url": "https://github.com/ansible/ansible-tower-samples",
        "scm_clean": false,
        "scm_delete_on_update": false,
        "status": "never updated",
        "scm_delete_on_next_update": false,
        "scm_update_on_launch": true,
        "scm_update_cache_timeout": 0,
        "last_update_failed": false
    }`)

	// MockedDeleteProjectResponse mocked `DeleteProject` endpoint response
	MockedDeleteProjectResponse = []byte(`{}`)

	// MockedCancelUpdateProjectResponse mocked `CancelUpdate` endpoint response
	MockedCancelUpdateProjectResponse = []byte(`{
        "can_cancel": true
    }`)

	// MockedGetUpdateProjectResponse mocked `GetUpdate` endpoint response
	MockedGetUpdateProjectResponse = []byte(`
        {
            "id": 301,
            "type": "job",
            "url": "/api/v2/jobs/301/",
            "related": {
                "created_by": "/api/v2/users/1/",
                "labels": "/api/v2/jobs/301/labels/",
                "inventory": "/api/v2/inventories/1/",
                "project": "/api/v2/projects/6/",
                "extra_credentials": "/api/v2/jobs/301/extra_credentials/",
                "credentials": "/api/v2/jobs/301/credentials/",
                "unified_job_template": "/api/v2/job_templates/8/",
                "stdout": "/api/v2/jobs/301/stdout/",
                "notifications": "/api/v2/jobs/301/notifications/",
                "job_host_summaries": "/api/v2/jobs/301/job_host_summaries/",
                "job_events": "/api/v2/jobs/301/job_events/",
                "activity_stream": "/api/v2/jobs/301/activity_stream/",
                "job_template": "/api/v2/job_templates/8/",
                "cancel": "/api/v2/jobs/301/cancel/",
                "project_update": "/api/v2/project_updates/303/",
                "create_schedule": "/api/v2/jobs/301/create_schedule/",
                "relaunch": "/api/v2/jobs/301/relaunch/"
            },
            "summary_fields": {
                "instance_group": {
                    "id": 1,
                    "name": "tower"
                },
                "job_template": {
                    "id": 8,
                    "name": "Hello-world",
                    "description": ""
                },
                "inventory": {
                    "id": 1,
                    "name": "Demo Inventory",
                    "description": "",
                    "has_active_failures": false,
                    "total_hosts": 2,
                    "hosts_with_active_failures": 0,
                    "total_groups": 0,
                    "groups_with_active_failures": 0,
                    "has_inventory_sources": false,
                    "total_inventory_sources": 0,
                    "inventory_sources_with_failures": 0,
                    "organization_id": 1,
                    "kind": ""
                },
                "project_update": {
                    "id": 303,
                    "name": "DeployBook",
                    "description": "",
                    "status": "successful",
                    "failed": false
                },
                "unified_job_template": {
                    "id": 8,
                    "name": "Hello-world",
                    "description": "",
                    "unified_job_type": "job"
                },
                "project": {
                    "id": 6,
                    "name": "DeployBook",
                    "description": "",
                    "status": "successful",
                    "scm_type": "git"
                },
                "created_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "user_capabilities": {
                    "start": true,
                    "delete": true
                },
                "labels": {
                    "count": 0,
                    "results": []
                },
                "extra_credentials": [],
                "credentials": []
            },
            "created": "2018-06-05T09:13:43.046077Z",
            "modified": "2018-06-05T09:13:49.132905Z",
            "name": "Hello-world",
            "description": "",
            "job_type": "run",
            "inventory": 1,
            "project": 6,
            "playbook": "hello-world.yml",
            "forks": 0,
            "limit": "localhost",
            "verbosity": 0,
            "extra_vars": "{}",
            "job_tags": "",
            "force_handlers": false,
            "skip_tags": "",
            "start_at_task": "",
            "timeout": 0,
            "use_fact_cache": false,
            "unified_job_template": 8,
            "launch_type": "manual",
            "status": "canceled",
            "failed": true,
            "started": "2018-06-05T09:13:47.060276Z",
            "finished": "2018-06-05T09:13:54.740342Z",
            "elapsed": 7.68,
            "job_args": "[\"ansible-playbook\", \"-i\", \"/tmp/awx_301_3kuP6Z/tmpupf88B\", \"-u\", \"root\", \"-l\", \"localhost\", \"-e\", \"@/tmp/awx_301_3kuP6Z/tmpjlw9Dp\", \"hello-world.yml\"]",
            "job_cwd": "/var/lib/awx/projects/_6__deploybook",
            "job_env": {},
            "job_explanation": "",
            "execution_node": "awx",
            "result_traceback": "",
            "event_processing_finished": true,
            "job_template": 8,
            "passwords_needed_to_start": [],
            "ask_diff_mode_on_launch": false,
            "ask_variables_on_launch": true,
            "ask_limit_on_launch": true,
            "ask_tags_on_launch": false,
            "ask_skip_tags_on_launch": false,
            "ask_job_type_on_launch": false,
            "ask_verbosity_on_launch": false,
            "ask_inventory_on_launch": true,
            "ask_credential_on_launch": false,
            "allow_simultaneous": false,
            "artifacts": {},
            "scm_revision": "ae9cc6b0fee9aea55e10a4a325c9ae3d97c42f1a",
            "instance_group": 1,
            "diff_mode": false,
            "credential": null,
            "vault_credential": null
        }`)
)
