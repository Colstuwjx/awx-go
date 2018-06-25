package awx

import (
	"time"
)

// common types here

type Pagination struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
}

type Related struct {
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
	Groups                       string `json:"groups"`
	Copy                         string `json:"copy"`
	UpdateInventorySources       string `json:"update_inventory_sources"`
	InventorySources             string `json:"inventory_sources"`
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
}

type OrgnizationSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ByUserSummary struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ApplyRole struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ObjectRoles struct {
	UseRole     *ApplyRole `json:"use_role"`
	AdminRole   *ApplyRole `json:"admin_role"`
	AdhocRole   *ApplyRole `json:"adhoc_role"`
	UpdateRole  *ApplyRole `json:"update_role"`
	ReadRole    *ApplyRole `json:"read_role"`
	ExecuteRole *ApplyRole `json:"execute_role"`
}

type UserCapabilities struct {
	Edit     bool `json:"edit"`
	Start    bool `json:"start"`
	Schedule bool `json:"schedule"`
	Copy     bool `json:"copy"`
	Adhoc    bool `json:"adhoc"`
	Delete   bool `json:"delete"`
}

type Labels struct {
	Count   int           `json:"count"`
	Results []interface{} `json:"results"`
}

type Summary struct {
	Organization     *OrgnizationSummary `json:"organization"`
	CreatedBy        *ByUserSummary      `json:"created_by"`
	ModifiedBy       *ByUserSummary      `json:"modified_by"`
	ObjectRoles      *ObjectRoles        `json:"object_roles"`
	UserCapabilities *UserCapabilities   `json:"user_capabilities"`
	Project          *Project            `json:"project"`
	Inventory        *Inventory          `json:"inventory"`
	RecentJobs       []interface{}       `json:"recent_jobs"`
	Credentials      []Credential        `json:"credentials"`
	Labels           *Labels             `json:"labels"`
}

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ScmType     string `json:"scm_type"`
}

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

type Credential struct {
	Description      string `json:"description"`
	CredentialTypeID int    `json:"credential_type_id"`
	ID               int    `json:"id"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
}

type InstanceGroup struct {
	Instances []string `json:"instances"`
	Capacity  int      `json:"capacity"`
	Name      string   `json:"name"`
}

type Instance struct {
	Node      string    `json:"node"`
	Heartbeat time.Time `json:"heartbeat"`
	Version   string    `json:"version"`
	Capacity  int       `json:"capacity"`
}

type Ping struct {
	Instances      []Instance      `json:"instances"`
	InstanceGroups []InstanceGroup `json:"instance_groups"`
	Ha             bool            `json:"ha"`
	Version        string          `json:"version"`
	ActiveNode     string          `json:"active_node"`
}

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
