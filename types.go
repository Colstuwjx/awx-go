package awx

import (
	"time"
)

const (
	JobStatusNew        = "new"
	JobStatusPending    = "pending"
	JobStatusWaiting    = "waiting"
	JobStatusRunning    = "running"
	JobStatusSuccessful = "successful"
	JobStatusFailed     = "failed"
	JobStatusError      = "error"
	JobStatusCanceled   = "canceled"
)

type PingResponse struct {
	Instances struct {
		Primary     string        `json:"primary"`
		Secondaries []interface{} `json:"secondaries"`
	} `json:"instances"`
	Ha      bool   `json:"ha"`
	Role    string `json:"role"`
	Version string `json:"version"`
}

type InventoryResponse struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		URL     string `json:"url"`
		Related struct {
			CreatedBy              string `json:"created_by"`
			ModifiedBy             string `json:"modified_by"`
			JobTemplates           string `json:"job_templates"`
			VariableData           string `json:"variable_data"`
			RootGroups             string `json:"root_groups"`
			ObjectRoles            string `json:"object_roles"`
			AdHocCommands          string `json:"ad_hoc_commands"`
			Script                 string `json:"script"`
			Tree                   string `json:"tree"`
			AccessList             string `json:"access_list"`
			ActivityStream         string `json:"activity_stream"`
			InstanceGroups         string `json:"instance_groups"`
			Hosts                  string `json:"hosts"`
			Groups                 string `json:"groups"`
			Copy                   string `json:"copy"`
			UpdateInventorySources string `json:"update_inventory_sources"`
			InventorySources       string `json:"inventory_sources"`
			Organization           string `json:"organization"`
		} `json:"related"`
		SummaryFields struct {
			Organization struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"organization"`
			CreatedBy struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			} `json:"created_by"`
			ModifiedBy struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			} `json:"modified_by"`
			ObjectRoles struct {
				UseRole struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"use_role"`
				AdminRole struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"admin_role"`
				AdhocRole struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"adhoc_role"`
				UpdateRole struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"update_role"`
				ReadRole struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"read_role"`
			} `json:"object_roles"`
			UserCapabilities struct {
				Edit   bool `json:"edit"`
				Copy   bool `json:"copy"`
				Adhoc  bool `json:"adhoc"`
				Delete bool `json:"delete"`
			} `json:"user_capabilities"`
		} `json:"summary_fields"`
		Created                      time.Time   `json:"created"`
		Modified                     time.Time   `json:"modified"`
		Name                         string      `json:"name"`
		Description                  string      `json:"description"`
		Organization                 int         `json:"organization"`
		Kind                         string      `json:"kind"`
		HostFilter                   interface{} `json:"host_filter"`
		Variables                    string      `json:"variables"`
		HasActiveFailures            bool        `json:"has_active_failures"`
		TotalHosts                   int         `json:"total_hosts"`
		HostsWithActiveFailures      int         `json:"hosts_with_active_failures"`
		TotalGroups                  int         `json:"total_groups"`
		GroupsWithActiveFailures     int         `json:"groups_with_active_failures"`
		HasInventorySources          bool        `json:"has_inventory_sources"`
		TotalInventorySources        int         `json:"total_inventory_sources"`
		InventorySourcesWithFailures int         `json:"inventory_sources_with_failures"`
		InsightsCredential           interface{} `json:"insights_credential"`
		PendingDeletion              bool        `json:"pending_deletion"`
	} `json:"results"`
}

type RunJobTemplateResponse struct {
	Job           int `json:"job"`
	IgnoredFields struct {
	} `json:"ignored_fields"`
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		CreatedBy          string `json:"created_by"`
		ModifiedBy         string `json:"modified_by"`
		Labels             string `json:"labels"`
		Inventory          string `json:"inventory"`
		Project            string `json:"project"`
		Credential         string `json:"credential"`
		ExtraCredentials   string `json:"extra_credentials"`
		Credentials        string `json:"credentials"`
		UnifiedJobTemplate string `json:"unified_job_template"`
		Stdout             string `json:"stdout"`
		Notifications      string `json:"notifications"`
		JobHostSummaries   string `json:"job_host_summaries"`
		JobEvents          string `json:"job_events"`
		ActivityStream     string `json:"activity_stream"`
		JobTemplate        string `json:"job_template"`
		Cancel             string `json:"cancel"`
		CreateSchedule     string `json:"create_schedule"`
		Relaunch           string `json:"relaunch"`
	} `json:"related"`
	SummaryFields struct {
		JobTemplate struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"job_template"`
		Inventory struct {
			ID                           int    `json:"id"`
			Name                         string `json:"name"`
			Description                  string `json:"description"`
			HasActiveFailures            bool   `json:"has_active_failures"`
			TotalHosts                   int    `json:"total_hosts"`
			HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
			TotalGroups                  int    `json:"total_groups"`
			GroupsWithActiveFailures     int    `json:"groups_with_active_failures"`
			HasInventorySources          bool   `json:"has_inventory_sources"`
			TotalInventorySources        int    `json:"total_inventory_sources"`
			InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
			OrganizationID               int    `json:"organization_id"`
			Kind                         string `json:"kind"`
		} `json:"inventory"`
		Credential struct {
			Description      string `json:"description"`
			CredentialTypeID int    `json:"credential_type_id"`
			ID               int    `json:"id"`
			Kind             string `json:"kind"`
			Name             string `json:"name"`
		} `json:"credential"`
		UnifiedJobTemplate struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Description    string `json:"description"`
			UnifiedJobType string `json:"unified_job_type"`
		} `json:"unified_job_template"`
		Project struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Status      string `json:"status"`
			ScmType     string `json:"scm_type"`
		} `json:"project"`
		CreatedBy struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"created_by"`
		ModifiedBy struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"modified_by"`
		UserCapabilities struct {
			Start  bool `json:"start"`
			Delete bool `json:"delete"`
		} `json:"user_capabilities"`
		Labels struct {
			Count   int           `json:"count"`
			Results []interface{} `json:"results"`
		} `json:"labels"`
		ExtraCredentials []interface{} `json:"extra_credentials"`
		Credentials      []struct {
			Description      string `json:"description"`
			CredentialTypeID int    `json:"credential_type_id"`
			ID               int    `json:"id"`
			Kind             string `json:"kind"`
			Name             string `json:"name"`
		} `json:"credentials"`
	} `json:"summary_fields"`
	Created            time.Time   `json:"created"`
	Modified           time.Time   `json:"modified"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	JobType            string      `json:"job_type"`
	Inventory          int         `json:"inventory"`
	Project            int         `json:"project"`
	Playbook           string      `json:"playbook"`
	Forks              int         `json:"forks"`
	Limit              string      `json:"limit"`
	Verbosity          int         `json:"verbosity"`
	ExtraVars          string      `json:"extra_vars"`
	JobTags            string      `json:"job_tags"`
	ForceHandlers      bool        `json:"force_handlers"`
	SkipTags           string      `json:"skip_tags"`
	StartAtTask        string      `json:"start_at_task"`
	Timeout            int         `json:"timeout"`
	UseFactCache       bool        `json:"use_fact_cache"`
	UnifiedJobTemplate int         `json:"unified_job_template"`
	LaunchType         string      `json:"launch_type"`
	Status             string      `json:"status"`
	Failed             bool        `json:"failed"`
	Started            interface{} `json:"started"`
	Finished           interface{} `json:"finished"`
	Elapsed            int         `json:"elapsed"`
	JobArgs            string      `json:"job_args"`
	JobCwd             string      `json:"job_cwd"`
	JobEnv             struct {
	} `json:"job_env"`
	JobExplanation          string        `json:"job_explanation"`
	ExecutionNode           string        `json:"execution_node"`
	ResultTraceback         string        `json:"result_traceback"`
	EventProcessingFinished bool          `json:"event_processing_finished"`
	JobTemplate             int           `json:"job_template"`
	PasswordsNeededToStart  []interface{} `json:"passwords_needed_to_start"`
	AskDiffModeOnLaunch     bool          `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch    bool          `json:"ask_variables_on_launch"`
	AskLimitOnLaunch        bool          `json:"ask_limit_on_launch"`
	AskTagsOnLaunch         bool          `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch     bool          `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch      bool          `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch    bool          `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch    bool          `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch   bool          `json:"ask_credential_on_launch"`
	AllowSimultaneous       bool          `json:"allow_simultaneous"`
	Artifacts               struct {
	} `json:"artifacts"`
	ScmRevision     string      `json:"scm_revision"`
	InstanceGroup   interface{} `json:"instance_group"`
	DiffMode        bool        `json:"diff_mode"`
	Credential      int         `json:"credential"`
	VaultCredential interface{} `json:"vault_credential"`
}

type JobEventResponseResult struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		Job  string `json:"job"`
		Host string `json:"host"`
	} `json:"related"`
	SummaryFields struct {
		Role struct {
		} `json:"role"`
		Job struct {
			ID              int     `json:"id"`
			Name            string  `json:"name"`
			Description     string  `json:"description"`
			Status          string  `json:"status"`
			Failed          bool    `json:"failed"`
			Elapsed         float64 `json:"elapsed"`
			JobTemplateID   int     `json:"job_template_id"`
			JobTemplateName string  `json:"job_template_name"`
		} `json:"job"`
		Host struct {
			ID                  int    `json:"id"`
			Name                string `json:"name"`
			Description         string `json:"description"`
			HasActiveFailures   bool   `json:"has_active_failures"`
			HasInventorySources bool   `json:"has_inventory_sources"`
		} `json:"host"`
	} `json:"summary_fields"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Job          int       `json:"job"`
	Event        string    `json:"event"`
	Counter      int       `json:"counter"`
	EventDisplay string    `json:"event_display"`
	EventData    struct {
		PlayPattern string      `json:"play_pattern"`
		Play        string      `json:"play"`
		EventLoop   interface{} `json:"event_loop"`
		TaskArgs    string      `json:"task_args"`
		RemoteAddr  string      `json:"remote_addr"`
		Res         struct {
			AnsibleParsed bool     `json:"_ansible_parsed"`
			StderrLines   []string `json:"stderr_lines"`
			Changed       bool     `json:"changed"`
			End           string   `json:"end"`
			AnsibleNoLog  bool     `json:"_ansible_no_log"`
			Stdout        string   `json:"stdout"`
			Cmd           string   `json:"cmd"`
			Start         string   `json:"start"`
			Delta         string   `json:"delta"`
			Stderr        string   `json:"stderr"`
			Rc            int      `json:"rc"`
			Invocation    struct {
				ModuleArgs struct {
					Creates    interface{} `json:"creates"`
					Executable interface{} `json:"executable"`
					UsesShell  bool        `json:"_uses_shell"`
					RawParams  string      `json:"_raw_params"`
					Removes    interface{} `json:"removes"`
					Warn       bool        `json:"warn"`
					Chdir      string      `json:"chdir"`
					Stdin      interface{} `json:"stdin"`
				} `json:"module_args"`
			} `json:"invocation"`
			StdoutLines []string `json:"stdout_lines"`
			Warnings    []string `json:"warnings"`
		} `json:"res"`
		Pid          int    `json:"pid"`
		PlayUUID     string `json:"play_uuid"`
		TaskUUID     string `json:"task_uuid"`
		Task         string `json:"task"`
		PlaybookUUID string `json:"playbook_uuid"`
		Playbook     string `json:"playbook"`
		TaskAction   string `json:"task_action"`
		Host         string `json:"host"`
		Role         string `json:"role"`
		TaskPath     string `json:"task_path"`
	} `json:"event_data"`
	EventLevel int         `json:"event_level"`
	Failed     bool        `json:"failed"`
	Changed    bool        `json:"changed"`
	UUID       string      `json:"uuid"`
	ParentUUID string      `json:"parent_uuid"`
	Host       int         `json:"host"`
	HostName   string      `json:"host_name"`
	Parent     interface{} `json:"parent"`
	Playbook   string      `json:"playbook"`
	Play       string      `json:"play"`
	Task       string      `json:"task"`
	Role       string      `json:"role"`
	Stdout     string      `json:"stdout"`
	StartLine  int         `json:"start_line"`
	EndLine    int         `json:"end_line"`
	Verbosity  int         `json:"verbosity"`
}

type JobEventResponse struct {
	Count    int                       `json:"count"`
	Next     interface{}               `json:"next"`
	Previous interface{}               `json:"previous"`
	Results  []*JobEventResponseResult `json:"results"`
}

type CancelJobResponse struct {
	Detail string `json:"detail"`
}

type JobResponse struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		CreatedBy          string `json:"created_by"`
		Labels             string `json:"labels"`
		Inventory          string `json:"inventory"`
		Project            string `json:"project"`
		Credential         string `json:"credential"`
		ExtraCredentials   string `json:"extra_credentials"`
		Credentials        string `json:"credentials"`
		UnifiedJobTemplate string `json:"unified_job_template"`
		Stdout             string `json:"stdout"`
		Notifications      string `json:"notifications"`
		JobHostSummaries   string `json:"job_host_summaries"`
		JobEvents          string `json:"job_events"`
		ActivityStream     string `json:"activity_stream"`
		JobTemplate        string `json:"job_template"`
		Cancel             string `json:"cancel"`
		ProjectUpdate      string `json:"project_update"`
		CreateSchedule     string `json:"create_schedule"`
		Relaunch           string `json:"relaunch"`
	} `json:"related"`
	SummaryFields struct {
		InstanceGroup struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"instance_group"`
		JobTemplate struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"job_template"`
		Inventory struct {
			ID                           int    `json:"id"`
			Name                         string `json:"name"`
			Description                  string `json:"description"`
			HasActiveFailures            bool   `json:"has_active_failures"`
			TotalHosts                   int    `json:"total_hosts"`
			HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
			TotalGroups                  int    `json:"total_groups"`
			GroupsWithActiveFailures     int    `json:"groups_with_active_failures"`
			HasInventorySources          bool   `json:"has_inventory_sources"`
			TotalInventorySources        int    `json:"total_inventory_sources"`
			InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
			OrganizationID               int    `json:"organization_id"`
			Kind                         string `json:"kind"`
		} `json:"inventory"`
		Credential struct {
			Description      string `json:"description"`
			CredentialTypeID int    `json:"credential_type_id"`
			ID               int    `json:"id"`
			Kind             string `json:"kind"`
			Name             string `json:"name"`
		} `json:"credential"`
		ProjectUpdate struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Status      string `json:"status"`
			Failed      bool   `json:"failed"`
		} `json:"project_update"`
		UnifiedJobTemplate struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Description    string `json:"description"`
			UnifiedJobType string `json:"unified_job_type"`
		} `json:"unified_job_template"`
		Project struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Status      string `json:"status"`
			ScmType     string `json:"scm_type"`
		} `json:"project"`
		CreatedBy struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"created_by"`
		UserCapabilities struct {
			Start  bool `json:"start"`
			Delete bool `json:"delete"`
		} `json:"user_capabilities"`
		Labels struct {
			Count   int           `json:"count"`
			Results []interface{} `json:"results"`
		} `json:"labels"`
		ExtraCredentials []interface{} `json:"extra_credentials"`
		Credentials      []struct {
			Description      string `json:"description"`
			CredentialTypeID int    `json:"credential_type_id"`
			ID               int    `json:"id"`
			Kind             string `json:"kind"`
			Name             string `json:"name"`
		} `json:"credentials"`
	} `json:"summary_fields"`
	Created            time.Time `json:"created"`
	Modified           time.Time `json:"modified"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	JobType            string    `json:"job_type"`
	Inventory          int       `json:"inventory"`
	Project            int       `json:"project"`
	Playbook           string    `json:"playbook"`
	Forks              int       `json:"forks"`
	Limit              string    `json:"limit"`
	Verbosity          int       `json:"verbosity"`
	ExtraVars          string    `json:"extra_vars"`
	JobTags            string    `json:"job_tags"`
	ForceHandlers      bool      `json:"force_handlers"`
	SkipTags           string    `json:"skip_tags"`
	StartAtTask        string    `json:"start_at_task"`
	Timeout            int       `json:"timeout"`
	UseFactCache       bool      `json:"use_fact_cache"`
	UnifiedJobTemplate int       `json:"unified_job_template"`
	LaunchType         string    `json:"launch_type"`
	Status             string    `json:"status"`
	Failed             bool      `json:"failed"`
	Started            time.Time `json:"started"`
	Finished           time.Time `json:"finished"`
	Elapsed            float64   `json:"elapsed"`
	JobArgs            string    `json:"job_args"`
	JobCwd             string    `json:"job_cwd"`
	JobEnv             struct {
		MEMCACHEDPORT11211TCP           string `json:"MEMCACHED_PORT_11211_TCP"`
		ANSIBLERETRYFILESENABLED        string `json:"ANSIBLE_RETRY_FILES_ENABLED"`
		AWXWEBENVAWXADMINUSER           string `json:"AWXWEB_ENV_AWX_ADMIN_USER"`
		HTTPProxy                       string `json:"http_proxy"`
		POSTGRESENVPOSTGRESPASSWORD     string `json:"POSTGRES_ENV_POSTGRES_PASSWORD"`
		RABBITMQPORT4369TCPADDR         string `json:"RABBITMQ_PORT_4369_TCP_ADDR"`
		AWXWEBENVDATABASEUSER           string `json:"AWXWEB_ENV_DATABASE_USER"`
		AWXWEBPORT8052TCPPROTO          string `json:"AWXWEB_PORT_8052_TCP_PROTO"`
		AWXWEBPORT8052TCPADDR           string `json:"AWXWEB_PORT_8052_TCP_ADDR"`
		RABBITMQENVRABBITMQLOGS         string `json:"RABBITMQ_ENV_RABBITMQ_LOGS"`
		VIRTUALENV                      string `json:"VIRTUAL_ENV"`
		RABBITMQPORT15671TCPPORT        string `json:"RABBITMQ_PORT_15671_TCP_PORT"`
		MEMCACHEDENVMEMCACHEDVERSION    string `json:"MEMCACHED_ENV_MEMCACHED_VERSION"`
		SUPERVISORENABLED               string `json:"SUPERVISOR_ENABLED"`
		RABBITMQPORT15672TCPPORT        string `json:"RABBITMQ_PORT_15672_TCP_PORT"`
		RABBITMQPORT5672TCPPROTO        string `json:"RABBITMQ_PORT_5672_TCP_PROTO"`
		DJANGOPROJECTDIR                string `json:"DJANGO_PROJECT_DIR"`
		RABBITMQPORT                    string `json:"RABBITMQ_PORT"`
		JOBID                           string `json:"JOB_ID"`
		RABBITMQHOST                    string `json:"RABBITMQ_HOST"`
		PYTHONPATH                      string `json:"PYTHONPATH"`
		POSTGRESENVPOSTGRESUSER         string `json:"POSTGRES_ENV_POSTGRES_USER"`
		POSTGRESENVLANG                 string `json:"POSTGRES_ENV_LANG"`
		LCCTYPE                         string `json:"LC_CTYPE"`
		MEMCACHEDPORT11211TCPPROTO      string `json:"MEMCACHED_PORT_11211_TCP_PROTO"`
		MPFORKLOGFILE                   string `json:"_MP_FORK_LOGFILE_"`
		RABBITMQPORT4369TCP             string `json:"RABBITMQ_PORT_4369_TCP"`
		MEMCACHEDPORT11211TCPPORT       string `json:"MEMCACHED_PORT_11211_TCP_PORT"`
		HOSTNAME                        string `json:"HOSTNAME"`
		AWXWEBENVDATABASENAME           string `json:"AWXWEB_ENV_DATABASE_NAME"`
		ANSIBLEINVENTORYENABLED         string `json:"ANSIBLE_INVENTORY_ENABLED"`
		ANSIBLEVENVPATH                 string `json:"ANSIBLE_VENV_PATH"`
		MEMCACHEDHOST                   string `json:"MEMCACHED_HOST"`
		AWXWEBENVRABBITMQHOST           string `json:"AWXWEB_ENV_RABBITMQ_HOST"`
		RABBITMQPORT25672TCPPROTO       string `json:"RABBITMQ_PORT_25672_TCP_PROTO"`
		MEMCACHEDNAME                   string `json:"MEMCACHED_NAME"`
		CELERYLOGREDIRECTLEVEL          string `json:"CELERY_LOG_REDIRECT_LEVEL"`
		SUPERVISORGROUPNAME             string `json:"SUPERVISOR_GROUP_NAME"`
		SECRETKEY                       string `json:"SECRET_KEY"`
		POSTGRESENVPGDATA               string `json:"POSTGRES_ENV_PGDATA"`
		RABBITMQENVRABBITMQERLANGCOOKIE string `json:"RABBITMQ_ENV_RABBITMQ_ERLANG_COOKIE"`
		RABBITMQPORT5672TCPADDR         string `json:"RABBITMQ_PORT_5672_TCP_ADDR"`
		SUPERVISORSERVERURL             string `json:"SUPERVISOR_SERVER_URL"`
		LANGUAGE                        string `json:"LANGUAGE"`
		SHLVL                           string `json:"SHLVL"`
		AWXADMINPASSWORD                string `json:"AWX_ADMIN_PASSWORD"`
		CELERYLOGFILE                   string `json:"CELERY_LOG_FILE"`
		RABBITMQPORT5672TCPPORT         string `json:"RABBITMQ_PORT_5672_TCP_PORT"`
		ANSIBLEHOSTKEYCHECKING          string `json:"ANSIBLE_HOST_KEY_CHECKING"`
		POSTGRESPORT5432TCP             string `json:"POSTGRES_PORT_5432_TCP"`
		LANG                            string `json:"LANG"`
		ANSIBLESTDOUTCALLBACK           string `json:"ANSIBLE_STDOUT_CALLBACK"`
		AWXWEBENVMEMCACHEDPORT          string `json:"AWXWEB_ENV_MEMCACHED_PORT"`
		SUPERVISORPROCESSNAME           string `json:"SUPERVISOR_PROCESS_NAME"`
		RABBITMQPORT5671TCPADDR         string `json:"RABBITMQ_PORT_5671_TCP_ADDR"`
		RABBITMQPORT4369TCPPORT         string `json:"RABBITMQ_PORT_4369_TCP_PORT"`
		POSTGRESENVPOSTGRESDB           string `json:"POSTGRES_ENV_POSTGRES_DB"`
		POSTGRESPORT5432TCPPORT         string `json:"POSTGRES_PORT_5432_TCP_PORT"`
		NAMING_FAILED                   string `json:"_"`
		RABBITMQVHOST                   string `json:"RABBITMQ_VHOST"`
		AWXWEBENVRABBITMQPASSWORD       string `json:"AWXWEB_ENV_RABBITMQ_PASSWORD"`
		DATABASENAME                    string `json:"DATABASE_NAME"`
		RABBITMQENVRABBITMQVERSION      string `json:"RABBITMQ_ENV_RABBITMQ_VERSION"`
		POSTGRESPORT5432TCPADDR         string `json:"POSTGRES_PORT_5432_TCP_ADDR"`
		RABBITMQENVRABBITMQGITHUBTAG    string `json:"RABBITMQ_ENV_RABBITMQ_GITHUB_TAG"`
		INVENTORYID                     string `json:"INVENTORY_ID"`
		DJANGOSETTINGSMODULE            string `json:"DJANGO_SETTINGS_MODULE"`
		AWXWEBENVNoProxy                string `json:"AWXWEB_ENV_no_proxy"`
		RABBITMQPORT15671TCPPROTO       string `json:"RABBITMQ_PORT_15671_TCP_PROTO"`
		ANSIBLEPARAMIKORECORDHOSTKEYS   string `json:"ANSIBLE_PARAMIKO_RECORD_HOST_KEYS"`
		AWXPRIVATEDATADIR               string `json:"AWX_PRIVATE_DATA_DIR"`
		RABBITMQPORT5671TCPPORT         string `json:"RABBITMQ_PORT_5671_TCP_PORT"`
		AWXWEBENVMEMCACHEDHOST          string `json:"AWXWEB_ENV_MEMCACHED_HOST"`
		LCALL                           string `json:"LC_ALL"`
		RABBITMQPORT5671TCP             string `json:"RABBITMQ_PORT_5671_TCP"`
		HOME                            string `json:"HOME"`
		RABBITMQPORT15672TCPPROTO       string `json:"RABBITMQ_PORT_15672_TCP_PROTO"`
		POSTGRESENVGOSUVERSION          string `json:"POSTGRES_ENV_GOSU_VERSION"`
		MPFORKLOGFORMAT                 string `json:"_MP_FORK_LOGFORMAT_"`
		AWXHOST                         string `json:"AWX_HOST"`
		RABBITMQPORT25672TCPPORT        string `json:"RABBITMQ_PORT_25672_TCP_PORT"`
		RABBITMQPORT15672TCPADDR        string `json:"RABBITMQ_PORT_15672_TCP_ADDR"`
		HTTPSProxy                      string `json:"https_proxy"`
		RABBITMQPORT4369TCPPROTO        string `json:"RABBITMQ_PORT_4369_TCP_PROTO"`
		RABBITMQENVRABBITMQHOME         string `json:"RABBITMQ_ENV_RABBITMQ_HOME"`
		CACHE                           string `json:"CACHE"`
		AWXWEBENVDATABASEPORT           string `json:"AWXWEB_ENV_DATABASE_PORT"`
		RABBITMQENVRABBITMQGPGKEY       string `json:"RABBITMQ_ENV_RABBITMQ_GPG_KEY"`
		AWXWEBPORT8052TCP               string `json:"AWXWEB_PORT_8052_TCP"`
		MPFORKLOGLEVEL                  string `json:"_MP_FORK_LOGLEVEL_"`
		AWXWEBENVRABBITMQUSER           string `json:"AWXWEB_ENV_RABBITMQ_USER"`
		AWXWEBENVRABBITMQPORT           string `json:"AWXWEB_ENV_RABBITMQ_PORT"`
		CELERYLOGLEVEL                  string `json:"CELERY_LOG_LEVEL"`
		DATABASEHOST                    string `json:"DATABASE_HOST"`
		AWXWEBPORT                      string `json:"AWXWEB_PORT"`
		POSTGRESNAME                    string `json:"POSTGRES_NAME"`
		POSTGRESENVPGVERSION            string `json:"POSTGRES_ENV_PG_VERSION"`
		AWXWEBENVRABBITMQVHOST          string `json:"AWXWEB_ENV_RABBITMQ_VHOST"`
		POSTGRESENVPGMAJOR              string `json:"POSTGRES_ENV_PG_MAJOR"`
		RABBITMQPORT15671TCP            string `json:"RABBITMQ_PORT_15671_TCP"`
		RABBITMQENVRABBITMQDEFAULTVHOST string `json:"RABBITMQ_ENV_RABBITMQ_DEFAULT_VHOST"`
		DJANGOLIVETESTSERVERADDRESS     string `json:"DJANGO_LIVE_TEST_SERVER_ADDRESS"`
		PROJECTREVISION                 string `json:"PROJECT_REVISION"`
		CELERYLOGREDIRECT               string `json:"CELERY_LOG_REDIRECT"`
		AWXWEBENVAWXADMINPASSWORD       string `json:"AWXWEB_ENV_AWX_ADMIN_PASSWORD"`
		RABBITMQPORT25672TCPADDR        string `json:"RABBITMQ_PORT_25672_TCP_ADDR"`
		AWXWEBENVDATABASEPASSWORD       string `json:"AWXWEB_ENV_DATABASE_PASSWORD"`
		ANSIBLEINVENTORYUNPARSEDFAILED  string `json:"ANSIBLE_INVENTORY_UNPARSED_FAILED"`
		PATH                            string `json:"PATH"`
		RABBITMQPORT5672TCP             string `json:"RABBITMQ_PORT_5672_TCP"`
		ANSIBLESSHCONTROLPATHDIR        string `json:"ANSIBLE_SSH_CONTROL_PATH_DIR"`
		DATABASEUSER                    string `json:"DATABASE_USER"`
		ANSIBLECALLBACKPLUGINS          string `json:"ANSIBLE_CALLBACK_PLUGINS"`
		RABBITMQPORT15671TCPADDR        string `json:"RABBITMQ_PORT_15671_TCP_ADDR"`
		TZ                              string `json:"TZ"`
		POSTGRESPORT5432TCPPROTO        string `json:"POSTGRES_PORT_5432_TCP_PROTO"`
		NoProxy                         string `json:"no_proxy"`
		MAXEVENTRES                     string `json:"MAX_EVENT_RES"`
		RABBITMQENVRABBITMQSASLLOGS     string `json:"RABBITMQ_ENV_RABBITMQ_SASL_LOGS"`
		AWXWEBPORT8052TCPPORT           string `json:"AWXWEB_PORT_8052_TCP_PORT"`
		AWXWEBNAME                      string `json:"AWXWEB_NAME"`
		AWXWEBENVSECRETKEY              string `json:"AWXWEB_ENV_SECRET_KEY"`
		AWXWEBENVDATABASEHOST           string `json:"AWXWEB_ENV_DATABASE_HOST"`
		MEMCACHEDPORT                   string `json:"MEMCACHED_PORT"`
		MEMCACHEDENVMEMCACHEDSHA1       string `json:"MEMCACHED_ENV_MEMCACHED_SHA1"`
		RABBITMQNAME                    string `json:"RABBITMQ_NAME"`
		DATABASEPORT                    string `json:"DATABASE_PORT"`
		CELERYLOADER                    string `json:"CELERY_LOADER"`
		RABBITMQPORT5671TCPPROTO        string `json:"RABBITMQ_PORT_5671_TCP_PROTO"`
		AWXADMINUSER                    string `json:"AWX_ADMIN_USER"`
		MEMCACHEDPORT11211TCPADDR       string `json:"MEMCACHED_PORT_11211_TCP_ADDR"`
		INVENTORYHOSTVARS               string `json:"INVENTORY_HOSTVARS"`
		ANSIBLEFORCECOLOR               string `json:"ANSIBLE_FORCE_COLOR"`
		RABBITMQUSER                    string `json:"RABBITMQ_USER"`
		POSTGRESPORT                    string `json:"POSTGRES_PORT"`
		AWXWEBENVHTTPProxy              string `json:"AWXWEB_ENV_http_proxy"`
		PWD                             string `json:"PWD"`
		AWXWEBENVHTTPSProxy             string `json:"AWXWEB_ENV_https_proxy"`
		RABBITMQPORT15672TCP            string `json:"RABBITMQ_PORT_15672_TCP"`
		DATABASEPASSWORD                string `json:"DATABASE_PASSWORD"`
		RABBITMQPASSWORD                string `json:"RABBITMQ_PASSWORD"`
		RABBITMQPORT25672TCP            string `json:"RABBITMQ_PORT_25672_TCP"`
	} `json:"job_env"`
	JobExplanation          string        `json:"job_explanation"`
	ExecutionNode           string        `json:"execution_node"`
	ResultTraceback         string        `json:"result_traceback"`
	EventProcessingFinished bool          `json:"event_processing_finished"`
	JobTemplate             int           `json:"job_template"`
	PasswordsNeededToStart  []interface{} `json:"passwords_needed_to_start"`
	AskDiffModeOnLaunch     bool          `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch    bool          `json:"ask_variables_on_launch"`
	AskLimitOnLaunch        bool          `json:"ask_limit_on_launch"`
	AskTagsOnLaunch         bool          `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch     bool          `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch      bool          `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch    bool          `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch    bool          `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch   bool          `json:"ask_credential_on_launch"`
	AllowSimultaneous       bool          `json:"allow_simultaneous"`
	Artifacts               struct {
	} `json:"artifacts"`
	ScmRevision     string      `json:"scm_revision"`
	InstanceGroup   int         `json:"instance_group"`
	DiffMode        bool        `json:"diff_mode"`
	Credential      int         `json:"credential"`
	VaultCredential interface{} `json:"vault_credential"`
}

type HostSummaryHost struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	HasActiveFailures   bool   `json:"has_active_failures"`
	HasInventorySources bool   `json:"has_inventory_sources"`
}

type HostSummaryJob struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Failed      bool    `json:"failed"`
	Elapsed     float64 `json:"elapsed"`
}

type HostSummaryFields struct {
	Host HostSummaryHost `json:"host"`
	Job  HostSummaryJob  `json:"job"`
}

type HostSummaryResult struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		Job  string `json:"job"`
		Host string `json:"host"`
	} `json:"related"`
	SummaryFields HostSummaryFields `json:"summary_fields"`
	Created       time.Time         `json:"created"`
	Modified      time.Time         `json:"modified"`
	Job           int               `json:"job"`
	Host          int               `json:"host"`
	HostName      string            `json:"host_name"`
	Changed       int               `json:"changed"`
	Dark          int               `json:"dark"`
	Failures      int               `json:"failures"`
	Ok            int               `json:"ok"`
	Processed     int               `json:"processed"`
	Skipped       int               `json:"skipped"`
	Failed        bool              `json:"failed"`
}

type HostSummaryResponse struct {
	Count    int                  `json:"count"`
	Next     interface{}          `json:"next"`
	Previous interface{}          `json:"previous"`
	Results  []*HostSummaryResult `json:"results"`
}
