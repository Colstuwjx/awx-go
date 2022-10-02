package awx

import (
	"time"
)

// Common types definition here
// For common usage, we made `Related` and `Summary` as two common field,
// it maybe happened that some structs don't have some fields in `Related` or `Summary`.

// Pagination represents the awx api pagination params.
type Pagination struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
}

// ProjectUpdateCancel represents the awx project update cancel api response.
type ProjectUpdateCancel struct {
	CanCancel bool `json:"can_cancel"`
}

// Related represents the awx api related field.
type Related struct {
	NamedURL                     string `json:"named_url"`
	CreatedBy                    string `json:"created_by"`
	ModifiedBy                   string `json:"modified_by"`
	JobTemplates                 string `json:"job_templates"`
	VariableData                 string `json:"variable_data"`
	RootGroups                   string `json:"root_groups"`
	ObjectRoles                  string `json:"object_roles"`
	AdHocCommands                string `json:"ad_hoc_commands"`
	Script                       string `json:"script"`
	Tree                         string `json:"tree"`
	AccessList                   string `json:"access_list"`
	ActivityStream               string `json:"activity_stream"`
	InstanceGroups               string `json:"instance_groups"`
	Hosts                        string `json:"hosts"`
	Job                          string `json:"job"`
	Host                         string `json:"host"`
	Groups                       string `json:"groups"`
	Copy                         string `json:"copy"`
	UpdateInventorySources       string `json:"update_inventory_sources"`
	InventorySources             string `json:"inventory_sources"`
	InventorySource              string `json:"inventory_source"`
	FactVersions                 string `json:"fact_versions"`
	SmartInventories             string `json:"smart_inventories"`
	Insights                     string `json:"insights"`
	Organization                 string `json:"organization"`
	Labels                       string `json:"labels"`
	Inventory                    string `json:"inventory"`
	Project                      string `json:"project"`
	Credential                   string `json:"credential"`
	ExtraCredentials             string `json:"extra_credentials"`
	Credentials                  string `json:"credentials"`
	NotificationTemplatesError   string `json:"notification_templates_error"`
	NotificationTemplatesSuccess string `json:"notification_templates_success"`
	Jobs                         string `json:"jobs"`
	NotificationTemplatesAny     string `json:"notification_templates_any"`
	Launch                       string `json:"launch"`
	Schedules                    string `json:"schedules"`
	SurveySpec                   string `json:"survey_spec"`
	UnifiedJobTemplate           string `json:"unified_job_template"`
	Stdout                       string `json:"stdout"`
	Notifications                string `json:"notifications"`
	JobHostSummaries             string `json:"job_host_summaries"`
	JobEvents                    string `json:"job_events"`
	JobTemplate                  string `json:"job_template"`
	Cancel                       string `json:"cancel"`
	ProjectUpdate                string `json:"project_update"`
	CreateSchedule               string `json:"create_schedule"`
	Relaunch                     string `json:"relaunch"`
	AdminOfOrganizations         string `json:"admin_of_organizations"`
	Organizations                string `json:"organizations"`
	Roles                        string `json:"roles"`
	Teams                        string `json:"teams"`
	Projects                     string `json:"projects"`
	PotentialChildren            string `json:"potential_children"`
	AllHosts                     string `json:"all_hosts"`
	AllGroups                    string `json:"all_groups"`
	AdHocCommandEvents           string `json:"ad_hoc_command_events"`
	Children                     string `json:"children"`
	AnsibleFacts                 string `json:"ansible_facts"`
	Events                       string `json:"events"`
}

// OrgnizationSummary represents the awx api orgnization summary fields.
type OrgnizationSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ByUserSummary represents the awx api user summary fields.
type ByUserSummary struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// JobTemplateSummary represents the awx api job template summary fields.
type JobTemplateSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// InstanceGroupSummary represents the awx api instance group summary fields.
type InstanceGroupSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ApplyRole represents the awx api apply role.
type ApplyRole struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ObjectRoles represents the awx api object roles.
type ObjectRoles struct {
	UseRole     *ApplyRole `json:"use_role"`
	AdminRole   *ApplyRole `json:"admin_role"`
	AdhocRole   *ApplyRole `json:"adhoc_role"`
	UpdateRole  *ApplyRole `json:"update_role"`
	ReadRole    *ApplyRole `json:"read_role"`
	ExecuteRole *ApplyRole `json:"execute_role"`
}

// UserCapabilities represents the awx api user capabilities.
type UserCapabilities struct {
	Edit     bool `json:"edit"`
	Start    bool `json:"start"`
	Schedule bool `json:"schedule"`
	Copy     bool `json:"copy"`
	Adhoc    bool `json:"adhoc"`
	Delete   bool `json:"delete"`
}

// Labels represents the awx api labels.
type Labels struct {
	Count   int           `json:"count"`
	Results []interface{} `json:"results"`
}

// Summary represents the awx api summary fields.
type Summary struct {
	InstanceGroup      *InstanceGroupSummary  `json:"instance_group"`
	Organization       *OrgnizationSummary    `json:"organization"`
	CreatedBy          *ByUserSummary         `json:"created_by"`
	ModifiedBy         *ByUserSummary         `json:"modified_by"`
	ObjectRoles        *ObjectRoles           `json:"object_roles"`
	UserCapabilities   *UserCapabilities      `json:"user_capabilities"`
	Project            *Project               `json:"project"`
	LastJob            map[string]interface{} `json:"last_job"`
	CurrentJob         map[string]interface{} `json:"current_job"`
	LastUpdate         map[string]interface{} `json:"last_update"`
	Inventory          *Inventory             `json:"inventory"`
	RecentJobs         []interface{}          `json:"recent_jobs"`
	Groups             *Groups                `json:"groups"`
	Credentials        []Credential           `json:"credentials"`
	Credential         *Credential            `json:"credential"`
	Labels             *Labels                `json:"labels"`
	JobTemplate        *JobTemplateSummary    `json:"job_template"`
	UnifiedJobTemplate *UnifiedJobTemplate    `json:"unified_job_template"`
	ExtraCredentials   []interface{}          `json:"extra_credentials"`
	ProjectUpdate      *ProjectUpdate         `json:"project_update"`
	InventorySource    *InventorySource       `json:"inventory_source"`
}

type InventorySource struct {
	Source      string    `json:"source"`
	LastUpdated time.Time `json:"last_updated"`
	Status      string    `json:"status"`
}

// ProjectUpdate represents the awx api project update.
type ProjectUpdate struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Failed      bool   `json:"failed"`
}

// Project represents the awx api project.
type Project struct {
	ID                    int       `json:"id"`
	Type                  string    `json:"type"`
	URL                   string    `json:"url"`
	Related               *Related  `json:"related"`
	SummaryFields         *Summary  `json:"summary_fields"`
	Created               time.Time `json:"created"`
	Modified              time.Time `json:"modified"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	LocalPath             string    `json:"local_path"`
	ScmType               string    `json:"scm_type"`
	ScmURL                string    `json:"scm_url"`
	ScmBranch             string    `json:"scm_branch"`
	ScmClean              bool      `json:"scm_clean"`
	ScmDeleteOnUpdate     bool      `json:"scm_delete_on_update"`
	Credential            string    `json:"credential"`
	Timeout               int       `json:"timeout"`
	LastJobRun            time.Time `json:"last_job_run"`
	LastJobFailed         bool      `json:"last_job_failed"`
	NextJobRun            time.Time `json:"next_job_run"`
	Status                string    `json:"status"`
	Organization          int       `json:"organization"`
	ScmDeleteOnNextUpdate bool      `json:"scm_delete_on_next_update"`
	ScmUpdateOnLaunch     bool      `json:"scm_update_on_launch"`
	ScmUpdateCacheTimeout int       `json:"scm_update_cache_timeout"`
	ScmRevision           string    `json:"scm_revision"`
	LastUpdateFailed      bool      `json:"last_update_failed"`
	LastUpdated           time.Time `json:"last_updated"`
}

// Inventory represents the awx api inventory.
type Inventory struct {
	ID                           int         `json:"id"`
	Type                         string      `json:"type"`
	URL                          string      `json:"url"`
	Related                      *Related    `json:"related"`
	SummaryFields                *Summary    `json:"summary_fields"`
	Created                      time.Time   `json:"created"`
	Modified                     time.Time   `json:"modified"`
	Name                         string      `json:"name"`
	Description                  string      `json:"description"`
	Organization                 int         `json:"organization"`
	OrganizationID               int         `json:"organization_id"`
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
}

// InventoryUpdate represents the awx api inventory update.
type InventoryUpdate struct {
	InventorySource         int         `json:"inventory_source"`
	Status                  string      `json:"status"`
	ID                      int         `json:"id"`
	Type                    string      `json:"type"`
	URL                     string      `json:"url"`
	Related                 *Related    `json:"related"`
	SummaryFields           *Summary    `json:"summary_fields"`
	Created                 time.Time   `json:"created"`
	Modified                time.Time   `json:"modified"`
	Name                    string      `json:"name"`
	Description             string      `json:"description"`
	Source                  string      `json:"source"`
	SourcePath              string      `json:"source_path"`
	SourceScript            interface{} `json:"source_script"`
	SourceVars              string      `json:"source_vars"`
	Credential              interface{} `json:"credential"`
	EnabledVar              string      `json:"enabled_var"`
	EnabledValue            string      `json:"enabled_value"`
	HostFilter              string      `json:"host_filter"`
	Overwrite               bool        `json:"overwrite"`
	OverwriteVars           bool        `json:"overwrite_vars"`
	CustomVirtualenv        interface{} `json:"custom_virtualenv"`
	Timeout                 int         `json:"timeout"`
	Verbosity               int         `json:"verbosity"`
	UnifiedJobTemplate      int         `json:"unified_job_template"`
	LaunchType              string      `json:"launch_type"`
	Failed                  bool        `json:"failed"`
	Started                 interface{} `json:"started"`
	Finished                interface{} `json:"finished"`
	CanceledOn              interface{} `json:"canceled_on"`
	Elapsed                 float64     `json:"elapsed"`
	JobArgs                 string      `json:"job_args"`
	JobCwd                  string      `json:"job_cwd"`
	JobEnv                  interface{} `json:"job_env"`
	JobExplanation          string      `json:"job_explanation"`
	ExecutionNode           string      `json:"execution_node"`
	ResultTraceback         string      `json:"result_traceback"`
	EventProcessingFinished bool        `json:"event_processing_finished"`
	Inventory               int         `json:"inventory"`
	LicenseError            bool        `json:"license_error"`
	OrgHostLimitError       bool        `json:"org_host_limit_error"`
	SourceProjectUpdate     interface{} `json:"source_project_update"`
	SourceProject           interface{} `json:"source_project"`
	InventoryUpdate         int         `json:"inventory_update"`
}

// Credential represents the awx api credential.
type Credential struct {
	Description      string `json:"description"`
	CredentialTypeID int    `json:"credential_type_id"`
	ID               int    `json:"id"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
}

// UnifiedJobTemplate represents the awx api unified job template.
type UnifiedJobTemplate struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	UnifiedJobType string `json:"unified_job_type"`
}

// InstanceGroup represents the awx api instance group.
type InstanceGroup struct {
	Instances []string `json:"instances"`
	Capacity  int      `json:"capacity"`
	Name      string   `json:"name"`
}

// Result data type
type Result struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Groups represents the awx api hosts group list
type Groups struct {
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

// Instance represents the awx api instance.
type Instance struct {
	Node      string    `json:"node"`
	Heartbeat time.Time `json:"heartbeat"`
	Version   string    `json:"version"`
	Capacity  int       `json:"capacity"`
}

// Ping represents the awx api ping.
type Ping struct {
	Instances      []Instance      `json:"instances"`
	InstanceGroups []InstanceGroup `json:"instance_groups"`
	Ha             bool            `json:"ha"`
	Version        string          `json:"version"`
	ActiveNode     string          `json:"active_node"`
}

// JobTemplate represents the awx api job template.
type JobTemplate struct {
	ID                    int         `json:"id"`
	Type                  string      `json:"type"`
	URL                   string      `json:"url"`
	Related               *Related    `json:"related"`
	SummaryFields         *Summary    `json:"summary_fields"`
	Created               time.Time   `json:"created"`
	Modified              time.Time   `json:"modified"`
	Name                  string      `json:"name"`
	Description           string      `json:"description"`
	JobType               string      `json:"job_type"`
	Inventory             int         `json:"inventory"`
	Project               int         `json:"project"`
	Playbook              string      `json:"playbook"`
	Forks                 int         `json:"forks"`
	Limit                 string      `json:"limit"`
	Verbosity             int         `json:"verbosity"`
	ExtraVars             string      `json:"extra_vars"`
	JobTags               string      `json:"job_tags"`
	ForceHandlers         bool        `json:"force_handlers"`
	SkipTags              string      `json:"skip_tags"`
	StartAtTask           string      `json:"start_at_task"`
	Timeout               int         `json:"timeout"`
	UseFactCache          bool        `json:"use_fact_cache"`
	LastJobRun            interface{} `json:"last_job_run"`
	LastJobFailed         bool        `json:"last_job_failed"`
	NextJobRun            interface{} `json:"next_job_run"`
	Status                string      `json:"status"`
	HostConfigKey         string      `json:"host_config_key"`
	AskDiffModeOnLaunch   bool        `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch  bool        `json:"ask_variables_on_launch"`
	AskLimitOnLaunch      bool        `json:"ask_limit_on_launch"`
	AskTagsOnLaunch       bool        `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch   bool        `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch    bool        `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch  bool        `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch  bool        `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch bool        `json:"ask_credential_on_launch"`
	SurveyEnabled         bool        `json:"survey_enabled"`
	BecomeEnabled         bool        `json:"become_enabled"`
	DiffMode              bool        `json:"diff_mode"`
	AllowSimultaneous     bool        `json:"allow_simultaneous"`
	CustomVirtualenv      interface{} `json:"custom_virtualenv"`
	Credential            int         `json:"credential"`
	VaultCredential       interface{} `json:"vault_credential"`
}

// JobLaunch represents the awx api job launch.
type JobLaunch struct {
	Job                     int               `json:"job"`
	IgnoredFields           map[string]string `json:"ignored_fields"`
	ID                      int               `json:"id"`
	Type                    string            `json:"type"`
	URL                     string            `json:"url"`
	Related                 *Related          `json:"related"`
	SummaryFields           *Summary          `json:"summary_fields"`
	Created                 time.Time         `json:"created"`
	Modified                time.Time         `json:"modified"`
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	JobType                 string            `json:"job_type"`
	Inventory               int               `json:"inventory"`
	Project                 int               `json:"project"`
	Playbook                string            `json:"playbook"`
	Forks                   int               `json:"forks"`
	Limit                   string            `json:"limit"`
	Verbosity               int               `json:"verbosity"`
	ExtraVars               string            `json:"extra_vars"`
	JobTags                 string            `json:"job_tags"`
	ForceHandlers           bool              `json:"force_handlers"`
	SkipTags                string            `json:"skip_tags"`
	StartAtTask             string            `json:"start_at_task"`
	Timeout                 int               `json:"timeout"`
	UseFactCache            bool              `json:"use_fact_cache"`
	UnifiedJobTemplate      int               `json:"unified_job_template"`
	LaunchType              string            `json:"launch_type"`
	Status                  string            `json:"status"`
	Failed                  bool              `json:"failed"`
	Started                 interface{}       `json:"started"`
	Finished                interface{}       `json:"finished"`
	Elapsed                 int               `json:"elapsed"`
	JobArgs                 string            `json:"job_args"`
	JobCwd                  string            `json:"job_cwd"`
	JobEnv                  map[string]string `json:"job_env"`
	JobExplanation          string            `json:"job_explanation"`
	ExecutionNode           string            `json:"execution_node"`
	ResultTraceback         string            `json:"result_traceback"`
	EventProcessingFinished bool              `json:"event_processing_finished"`
	JobTemplate             int               `json:"job_template"`
	PasswordsNeededToStart  []interface{}     `json:"passwords_needed_to_start"`
	AskDiffModeOnLaunch     bool              `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch    bool              `json:"ask_variables_on_launch"`
	AskLimitOnLaunch        bool              `json:"ask_limit_on_launch"`
	AskTagsOnLaunch         bool              `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch     bool              `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch      bool              `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch    bool              `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch    bool              `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch   bool              `json:"ask_credential_on_launch"`
	AllowSimultaneous       bool              `json:"allow_simultaneous"`
	Artifacts               map[string]string `json:"artifacts"`
	ScmRevision             string            `json:"scm_revision"`
	InstanceGroup           interface{}       `json:"instance_group"`
	DiffMode                bool              `json:"diff_mode"`
	Credential              int               `json:"credential"`
	VaultCredential         interface{}       `json:"vault_credential"`
}

// Job represents the awx api job.
type Job struct {
	ID                      int               `json:"id"`
	Type                    string            `json:"type"`
	URL                     string            `json:"url"`
	Related                 *Related          `json:"related"`
	SummaryFields           *Summary          `json:"summary_fields"`
	Created                 time.Time         `json:"created"`
	Modified                time.Time         `json:"modified"`
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	JobType                 string            `json:"job_type"`
	Inventory               int               `json:"inventory"`
	Project                 int               `json:"project"`
	Playbook                string            `json:"playbook"`
	Forks                   int               `json:"forks"`
	Limit                   string            `json:"limit"`
	Verbosity               int               `json:"verbosity"`
	ExtraVars               string            `json:"extra_vars"`
	JobTags                 string            `json:"job_tags"`
	ForceHandlers           bool              `json:"force_handlers"`
	SkipTags                string            `json:"skip_tags"`
	StartAtTask             string            `json:"start_at_task"`
	Timeout                 int               `json:"timeout"`
	UseFactCache            bool              `json:"use_fact_cache"`
	UnifiedJobTemplate      int               `json:"unified_job_template"`
	LaunchType              string            `json:"launch_type"`
	Status                  string            `json:"status"`
	Failed                  bool              `json:"failed"`
	Started                 time.Time         `json:"started"`
	Finished                time.Time         `json:"finished"`
	Elapsed                 float64           `json:"elapsed"`
	JobArgs                 string            `json:"job_args"`
	JobCwd                  string            `json:"job_cwd"`
	JobEnv                  map[string]string `json:"job_env"`
	JobExplanation          string            `json:"job_explanation"`
	ExecutionNode           string            `json:"execution_node"`
	ResultTraceback         string            `json:"result_traceback"`
	EventProcessingFinished bool              `json:"event_processing_finished"`
	JobTemplate             int               `json:"job_template"`
	PasswordsNeededToStart  []interface{}     `json:"passwords_needed_to_start"`
	AskDiffModeOnLaunch     bool              `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch    bool              `json:"ask_variables_on_launch"`
	AskLimitOnLaunch        bool              `json:"ask_limit_on_launch"`
	AskTagsOnLaunch         bool              `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch     bool              `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch      bool              `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch    bool              `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch    bool              `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch   bool              `json:"ask_credential_on_launch"`
	AllowSimultaneous       bool              `json:"allow_simultaneous"`
	Artifacts               map[string]string `json:"artifacts"`
	ScmRevision             string            `json:"scm_revision"`
	InstanceGroup           int               `json:"instance_group"`
	DiffMode                bool              `json:"diff_mode"`
	Credential              *Credential       `json:"credential"`
	VaultCredential         interface{}       `json:"vault_credential"`
}

// HostSummaryHost represents the awx api host summary host fields.
type HostSummaryHost struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	HasActiveFailures   bool   `json:"has_active_failures"`
	HasInventorySources bool   `json:"has_inventory_sources"`
}

// HostSummaryJob represents the awx api host summary job fields.
type HostSummaryJob struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Status          string  `json:"status"`
	Failed          bool    `json:"failed"`
	Elapsed         float64 `json:"elapsed"`
	JobTemplateID   int     `json:"job_template_id"`
	JobTemplateName string  `json:"job_template_name"`
}

// HostSummaryFields represents the awx api host summary fields.
type HostSummaryFields struct {
	Role map[string]string `json:"role"`
	Host *HostSummaryHost  `json:"host"`
	Job  *HostSummaryJob   `json:"job"`
}

// HostSummary represents the awx api host summary.
type HostSummary struct {
	ID            int                `json:"id"`
	Type          string             `json:"type"`
	URL           string             `json:"url"`
	Related       *Related           `json:"related"`
	SummaryFields *HostSummaryFields `json:"summary_fields"`
	Created       time.Time          `json:"created"`
	Modified      time.Time          `json:"modified"`
	Job           int                `json:"job"`
	Host          int                `json:"host"`
	HostName      string             `json:"host_name"`
	Changed       int                `json:"changed"`
	Dark          int                `json:"dark"`
	Failures      int                `json:"failures"`
	Ok            int                `json:"ok"`
	Processed     int                `json:"processed"`
	Skipped       int                `json:"skipped"`
	Failed        bool               `json:"failed"`
}

// EventModuleArgs represents the awx api event module args.
type EventModuleArgs struct {
	Creates    interface{} `json:"creates"`
	Executable interface{} `json:"executable"`
	UsesShell  bool        `json:"_uses_shell"`
	RawParams  string      `json:"_raw_params"`
	Removes    interface{} `json:"removes"`
	Warn       bool        `json:"warn"`
	Chdir      string      `json:"chdir"`
	Stdin      interface{} `json:"stdin"`
}

// EventInvocation represents the awx api event invocation.
type EventInvocation struct {
	ModuleArgs *EventModuleArgs `json:"module_args"`
}

// EventRes represents the awx api event response.
type EventRes struct {
	AnsibleParsed bool             `json:"_ansible_parsed"`
	AnsibleFacts  interface{}      `json:"ansible_facts"`
	StderrLines   []string         `json:"stderr_lines"`
	Changed       bool             `json:"changed"`
	End           string           `json:"end"`
	AnsibleNoLog  bool             `json:"_ansible_no_log"`
	Stdout        string           `json:"stdout"`
	Cmd           string           `json:"cmd"`
	Start         string           `json:"start"`
	Delta         string           `json:"delta"`
	Stderr        string           `json:"stderr"`
	Rc            int              `json:"rc"`
	Invocation    *EventInvocation `json:"invocation"`
	StdoutLines   []string         `json:"stdout_lines"`
	Warnings      []string         `json:"warnings"`
}

// EventData represents the awx api event data.
type EventData struct {
	PlayPattern  string      `json:"play_pattern"`
	Play         string      `json:"play"`
	EventLoop    interface{} `json:"event_loop"`
	TaskArgs     string      `json:"task_args"`
	RemoteAddr   string      `json:"remote_addr"`
	Res          *EventRes   `json:"res"`
	Pid          int         `json:"pid"`
	PlayUUID     string      `json:"play_uuid"`
	TaskUUID     string      `json:"task_uuid"`
	Task         string      `json:"task"`
	PlaybookUUID string      `json:"playbook_uuid"`
	Playbook     string      `json:"playbook"`
	TaskAction   string      `json:"task_action"`
	Host         string      `json:"host"`
	Role         string      `json:"role"`
	TaskPath     string      `json:"task_path"`
}

// JobEvent represents the awx api job event.
type JobEvent struct {
	ID            int                `json:"id"`
	Type          string             `json:"type"`
	URL           string             `json:"url"`
	Related       *Related           `json:"related"`
	SummaryFields *HostSummaryFields `json:"summary_fields"`
	Created       time.Time          `json:"created"`
	Modified      time.Time          `json:"modified"`
	Job           int                `json:"job"`
	Event         string             `json:"event"`
	Counter       int                `json:"counter"`
	EventDisplay  string             `json:"event_display"`
	EventData     *EventData         `json:"event_data"`
	EventLevel    int                `json:"event_level"`
	Failed        bool               `json:"failed"`
	Changed       bool               `json:"changed"`
	UUID          string             `json:"uuid"`
	ParentUUID    string             `json:"parent_uuid"`

	// FIXME: inconsistent value type from tower API, int, null
	Host interface{} `json:"host"`

	HostName  string      `json:"host_name"`
	Parent    interface{} `json:"parent"`
	Playbook  string      `json:"playbook"`
	Play      string      `json:"play"`
	Task      string      `json:"task"`
	Role      string      `json:"role"`
	Stdout    string      `json:"stdout"`
	StartLine int         `json:"start_line"`
	EndLine   int         `json:"end_line"`
	Verbosity int         `json:"verbosity"`
}

// User represents an user
type User struct {
	ID              int         `json:"id"`
	Type            int         `json:"type"`
	URL             string      `json:"url"`
	Related         *Related    `json:"related"`
	SummaryFields   *Summary    `json:"summary_fields"`
	Created         time.Time   `json:"created"`
	Username        string      `json:"username"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Email           string      `json:"email"`
	IsSuperUser     bool        `json:"is_superuser"`
	IsSystemAuditor bool        `json:"is_system_auditor"`
	Password        string      `json:"password"`
	LdapDn          string      `json:"ldap_dn"`
	ExternalAccount interface{} `json:"external_account"`
}

// Group represents a group
type Group struct {
	ID                       int       `json:"id"`
	Type                     int       `json:"type"`
	URL                      string    `json:"url"`
	Related                  *Related  `json:"related"`
	SummaryFields            *Summary  `json:"summary_fields"`
	Created                  time.Time `json:"created"`
	Modified                 time.Time `json:"modified"`
	Name                     string    `json:"name"`
	Description              string    `json:"description"`
	Inventory                int       `json:"inventory"`
	Variables                string    `json:"variables"`
	HasActiveFailures        bool      `json:"has_active_failures"`
	TotalHosts               int       `json:"total_hosts"`
	HostsWithActiveFailures  int       `json:"hosts_with_active_failures"`
	TotalGroups              int       `json:"total_groups"`
	GroupsWithActiveFailures int       `json:"groups_with_active_failures"`
	HasInventorySources      bool      `json:"has_inventory_sources"`
}

// Host represents a host
type Host struct {
	ID                   int          `json:"id"`
	Type                 string       `json:"type"`
	URL                  string       `json:"url"`
	Related              *Related     `json:"related"`
	SummaryFields        *Summary     `json:"summary_fields"`
	Created              time.Time    `json:"created"`
	Modified             time.Time    `json:"modified"`
	Name                 string       `json:"name"`
	Description          string       `json:"description"`
	Inventory            int          `json:"inventory"`
	Enabled              bool         `json:"enabled"`
	InstanceID           string       `json:"instance_id"`
	Variables            string       `json:"variables"`
	HasActiveFailures    bool         `json:"has_active_failures"`
	HasInventorySources  bool         `json:"has_inventory_sources"`
	LastJob              *Job         `json:"last_job"`
	LastJobHostSummary   *HostSummary `json:"last_job_host_summary"`
	InsightsSystemID     interface{}  `json:"insights_system_id"`
	AnsibleFactsModified interface{}  `json:"ansible_facts_modified"`
}
