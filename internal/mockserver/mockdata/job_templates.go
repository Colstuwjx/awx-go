package mockdata

var (
	MockedListJobTemplatesResponse = []byte(`
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
        {
            "id": 5,
            "type": "job_template",
            "url": "/api/v2/job_templates/5/",
            "related": {
                "created_by": "/api/v2/users/1/",
                "modified_by": "/api/v2/users/1/",
                "labels": "/api/v2/job_templates/5/labels/",
                "inventory": "/api/v2/inventories/1/",
                "project": "/api/v2/projects/4/",
                "credential": "/api/v2/credentials/1/",
                "extra_credentials": "/api/v2/job_templates/5/extra_credentials/",
                "credentials": "/api/v2/job_templates/5/credentials/",
                "notification_templates_error": "/api/v2/job_templates/5/notification_templates_error/",
                "notification_templates_success": "/api/v2/job_templates/5/notification_templates_success/",
                "jobs": "/api/v2/job_templates/5/jobs/",
                "object_roles": "/api/v2/job_templates/5/object_roles/",
                "notification_templates_any": "/api/v2/job_templates/5/notification_templates_any/",
                "access_list": "/api/v2/job_templates/5/access_list/",
                "launch": "/api/v2/job_templates/5/launch/",
                "instance_groups": "/api/v2/job_templates/5/instance_groups/",
                "schedules": "/api/v2/job_templates/5/schedules/",
                "copy": "/api/v2/job_templates/5/copy/",
                "activity_stream": "/api/v2/job_templates/5/activity_stream/",
                "survey_spec": "/api/v2/job_templates/5/survey_spec/"
            },
            "summary_fields": {
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
                "project": {
                    "id": 4,
                    "name": "Demo Project",
                    "description": "",
                    "status": "never updated",
                    "scm_type": "git"
                },
                "created_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 1,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "object_roles": {
                    "admin_role": {
                        "id": 27,
                        "description": "Can manage all aspects of the job template",
                        "name": "Admin"
                    },
                    "execute_role": {
                        "id": 26,
                        "description": "May run the job template",
                        "name": "Execute"
                    },
                    "read_role": {
                        "id": 25,
                        "description": "May view settings for the job template",
                        "name": "Read"
                    }
                },
                "user_capabilities": {
                    "edit": true,
                    "start": true,
                    "copy": true,
                    "schedule": true,
                    "delete": true
                },
                "labels": {
                    "count": 0,
                    "results": []
                },
                "recent_jobs": [],
                "credentials": [
                    {
                        "description": "",
                        "credential_type_id": 1,
                        "id": 1,
                        "kind": "ssh",
                        "name": "Demo Credential"
                    }
                ]
            },
            "created": "2018-05-21T01:34:35.773593Z",
            "modified": "2018-05-21T01:34:35.773605Z",
            "name": "Demo Job Template",
            "description": "",
            "job_type": "run",
            "inventory": 1,
            "project": 4,
            "playbook": "hello_world.yml",
            "forks": 0,
            "limit": "",
            "verbosity": 0,
            "extra_vars": "",
            "job_tags": "",
            "force_handlers": false,
            "skip_tags": "",
            "start_at_task": "",
            "timeout": 0,
            "use_fact_cache": false,
            "last_job_run": null,
            "last_job_failed": false,
            "next_job_run": null,
            "status": "never updated",
            "host_config_key": "",
            "ask_diff_mode_on_launch": false,
            "ask_variables_on_launch": false,
            "ask_limit_on_launch": false,
            "ask_tags_on_launch": false,
            "ask_skip_tags_on_launch": false,
            "ask_job_type_on_launch": false,
            "ask_verbosity_on_launch": false,
            "ask_inventory_on_launch": false,
            "ask_credential_on_launch": false,
            "survey_enabled": false,
            "become_enabled": false,
            "diff_mode": false,
            "allow_simultaneous": false,
            "custom_virtualenv": null,
            "credential": 1,
            "vault_credential": null
        }
    ]
}`)
)
