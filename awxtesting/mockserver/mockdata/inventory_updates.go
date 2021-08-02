package mockdata

var (
	// MockedInventoryUpdatesResponse mocked `SyncInventorySourcesByInventoryID` endpoint response
	MockedInventoryUpdatesResponse = []byte(`
    [
        {
            "inventory_source": 10,
            "status": "started",
            "id": 305,
            "type": "inventory_update",
            "url": "/api/v2/inventory_updates/305/",
            "related": {
                "created_by": "/api/v2/users/5/",
                "modified_by": "/api/v2/users/5/",
                "unified_job_template": "/api/v2/inventory_sources/10/",
                "stdout": "/api/v2/inventory_updates/305/stdout/",
                "inventory_source": "/api/v2/inventory_sources/10/",
                "cancel": "/api/v2/inventory_updates/305/cancel/",
                "notifications": "/api/v2/inventory_updates/305/notifications/",
                "events": "/api/v2/inventory_updates/305/events/",
                "inventory": "/api/v2/inventories/1/",
                "credentials": "/api/v2/inventory_updates/305/credentials/"
            },
            "summary_fields": {
                "organization": {
                    "id": 1,
                    "name": "Default",
                    "description": ""
                },
                "inventory": {
                    "id": 1,
                    "name": "Default",
                    "description": "",
                    "has_active_failures": true,
                    "total_hosts": 7,
                    "hosts_with_active_failures": 6,
                    "total_groups": 14,
                    "has_inventory_sources": true,
                    "total_inventory_sources": 1,
                    "inventory_sources_with_failures": 0,
                    "organization_id": 1,
                    "kind": ""
                },
                "unified_job_template": {
                    "id": 10,
                    "name": "Default",
                    "description": "",
                    "unified_job_type": "inventory_update"
                },
                "inventory_source": {
                    "source": "scm",
                    "last_updated": "2021-07-30T10:12:44.553099Z",
                    "status": "pending"
                },
                "created_by": {
                    "id": 5,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 5,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "user_capabilities": {
                    "delete": true,
                    "start": true
                },
                "credentials": []
            },
            "created": "2021-08-02T01:45:22.144755Z",
            "modified": "2021-08-02T01:45:22.157220Z",
            "name": "Default",
            "description": "",
            "source": "scm",
            "source_path": "",
            "source_script": null,
            "source_vars": "",
            "credential": null,
            "enabled_var": "",
            "enabled_value": "",
            "host_filter": "",
            "overwrite": true,
            "overwrite_vars": true,
            "custom_virtualenv": null,
            "timeout": 0,
            "verbosity": 2,
            "unified_job_template": 10,
            "launch_type": "manual",
            "failed": false,
            "started": null,
            "finished": null,
            "canceled_on": null,
            "elapsed": 0.0,
            "job_args": "",
            "job_cwd": "",
            "job_env": null,
            "job_explanation": "",
            "execution_node": "",
            "result_traceback": "",
            "event_processing_finished": false,
            "inventory": 1,
            "license_error": false,
            "org_host_limit_error": false,
            "source_project_update": null,
            "source_project": null,
            "inventory_update": 305
        }
    ]`)

	// MockedInventoryUpdateResponse mocked `GetInventoryUpdate` endpoint response
	MockedInventoryUpdateResponse = []byte(`
        {
            "status": "successful",
            "id": 305,
            "type": "inventory_update",
            "url": "/api/v2/inventory_updates/305/",
            "related": {
                "created_by": "/api/v2/users/5/",
                "modified_by": "/api/v2/users/5/",
                "unified_job_template": "/api/v2/inventory_sources/10/",
                "stdout": "/api/v2/inventory_updates/305/stdout/",
                "inventory_source": "/api/v2/inventory_sources/10/",
                "cancel": "/api/v2/inventory_updates/305/cancel/",
                "notifications": "/api/v2/inventory_updates/305/notifications/",
                "events": "/api/v2/inventory_updates/305/events/",
                "inventory": "/api/v2/inventories/1/",
                "credentials": "/api/v2/inventory_updates/305/credentials/"
            },
            "summary_fields": {
                "organization": {
                    "id": 1,
                    "name": "Default",
                    "description": ""
                },
                "inventory": {
                    "id": 1,
                    "name": "Default",
                    "description": "",
                    "has_active_failures": true,
                    "total_hosts": 7,
                    "hosts_with_active_failures": 6,
                    "total_groups": 14,
                    "has_inventory_sources": true,
                    "total_inventory_sources": 1,
                    "inventory_sources_with_failures": 0,
                    "organization_id": 1,
                    "kind": ""
                },
                "unified_job_template": {
                    "id": 10,
                    "name": "Default",
                    "description": "",
                    "unified_job_type": "inventory_update"
                },
                "inventory_source": {
                    "source": "scm",
                    "last_updated": "2021-07-30T10:12:44.553099Z",
                    "status": "pending"
                },
                "created_by": {
                    "id": 5,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "modified_by": {
                    "id": 5,
                    "username": "admin",
                    "first_name": "",
                    "last_name": ""
                },
                "user_capabilities": {
                    "delete": true,
                    "start": true
                },
                "credentials": []
            },
            "created": "2021-08-02T01:45:22.144755Z",
            "modified": "2021-08-02T01:45:22.157220Z",
            "name": "Default",
            "description": "",
            "source": "scm",
            "source_path": "",
            "source_script": null,
            "source_vars": "",
            "credential": null,
            "enabled_var": "",
            "enabled_value": "",
            "host_filter": "",
            "overwrite": true,
            "overwrite_vars": true,
            "custom_virtualenv": null,
            "timeout": 0,
            "verbosity": 2,
            "unified_job_template": 10,
            "launch_type": "manual",
            "failed": false,
            "started": null,
            "finished": null,
            "canceled_on": null,
            "elapsed": 0.0,
            "job_args": "",
            "job_cwd": "",
            "job_env": null,
            "job_explanation": "",
            "execution_node": "",
            "result_traceback": "",
            "event_processing_finished": false,
            "inventory": 1,
            "license_error": false,
            "org_host_limit_error": false,
            "source_project_update": null,
            "source_project": null,
            "inventory_update": 305
		}`)
)
