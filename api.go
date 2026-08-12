// github.com/Infrawrench/infrawrench-go v1.14.0 | MIT | Copyright (c) 2026 Infrawrench LLC
// https://github.com/Infrawrench/Infrawrench
//
// Generated from the Infrawrench API OpenAPI 3.1 spec (API version 1.14.0).
//
// DO NOT EDIT. Regenerate with:
//   pnpm --filter @infrawrench/web generate:sdk
//
// Internal routes are absent by construction: the generator consumes the same
// published spec that /openapi.json serves, which drops every operation
// marked x-internal.

package infrawrench

import (
	"context"
	"io"
	"net/http"
)

// WithOrgID sets the orgId every scoped call falls back to, so it can be left
// out of the parameters. A call that passes its own wins; a call that has
// neither returns ErrMissingPathParam.
func WithOrgID(orgID string) ClientOption {
	return withScope(orgID)
}

// APIV1Client is a client for the Infrawrench API.
//
// Calls hang off the namespace fields below, which mirror the URL structure —
// `client.Accounts.Credentials.Get(ctx, …)`. The zero value is not usable; build
// one with NewAPIV1Client.
type APIV1Client struct {
	t *transport

	// AccessRequests: `client.accessRequests`.
	AccessRequests *AccessRequestsNamespace
	// Accounts: `client.accounts`.
	Accounts *AccountsNamespace
	// Agents: `client.agents`.
	Agents *AgentsNamespace
	// AlertRules: `client.alertRules`.
	AlertRules *AlertRulesNamespace
	// APIKeys: `client.apiKeys`.
	APIKeys *APIKeysNamespace
	// Artifacts: `client.artifacts`.
	Artifacts *ArtifactsNamespace
	// Associations: `client.associations`.
	Associations *AssociationsNamespace
	// AuditLogs: `client.auditLogs`.
	AuditLogs *AuditLogsNamespace
	// Auth: `client.auth`.
	Auth *AuthNamespace
	// Bastions: `client.bastions`.
	Bastions *BastionsNamespace
	// Billing: `client.billing`.
	Billing *BillingNamespace
	// BillingRules: `client.billingRules`.
	BillingRules *BillingRulesNamespace
	// BlastRadius: `client.blastRadius`.
	BlastRadius *BlastRadiusNamespace
	// Budgets: `client.budgets`.
	Budgets *BudgetsNamespace
	// BusinessMetrics: `client.businessMetrics`.
	BusinessMetrics *BusinessMetricsNamespace
	// ChangeFreezes: `client.changeFreezes`.
	ChangeFreezes *ChangeFreezesNamespace
	// Changes: `client.changes`.
	Changes *ChangesNamespace
	// Commitments: `client.commitments`.
	Commitments *CommitmentsNamespace
	// Config: `client.config`.
	Config *ConfigNamespace
	// Connect: `client.connect`.
	Connect *ConnectNamespace
	// CostAlerts: `client.costAlerts`.
	CostAlerts *CostAlertsNamespace
	// CostAnnotations: `client.costAnnotations`.
	CostAnnotations *CostAnnotationsNamespace
	// CostCentres: `client.costCentres`.
	CostCentres *CostCentresNamespace
	// CostExports: `client.costExports`.
	CostExports *CostExportsNamespace
	// CostReportFolders: `client.costReportFolders`.
	CostReportFolders *CostReportFoldersNamespace
	// CostReportNotifications: `client.costReportNotifications`.
	CostReportNotifications *CostReportNotificationsNamespace
	// CostReports: `client.costReports`.
	CostReports *CostReportsNamespace
	// CostScenarios: `client.costScenarios`.
	CostScenarios *CostScenariosNamespace
	// Costs: `client.costs`.
	Costs *CostsNamespace
	// CredentialHygiene: `client.credentialHygiene`.
	CredentialHygiene *CredentialHygieneNamespace
	// Credits: `client.credits`.
	Credits *CreditsNamespace
	// Currency: `client.currency`.
	Currency *CurrencyNamespace
	// CustomGraphs: `client.customGraphs`.
	CustomGraphs *CustomGraphsNamespace
	// Dashboards: `client.dashboards`.
	Dashboards *DashboardsNamespace
	// DependencyGraph: `client.dependencyGraph`.
	DependencyGraph *DependencyGraphNamespace
	// Deployments: `client.deployments`.
	Deployments *DeploymentsNamespace
	// Digest: `client.digest`.
	Digest *DigestNamespace
	// DNS: `client.dns`.
	DNS *DNSNamespace
	// Docker: `client.docker`.
	Docker *DockerNamespace
	// EnvironmentDiff: `client.environmentDiff`.
	EnvironmentDiff *EnvironmentDiffNamespace
	// Expiring: `client.expiring`.
	Expiring *ExpiringNamespace
	// Invitations: `client.invitations`.
	Invitations *InvitationsNamespace
	// Invoices: `client.invoices`.
	Invoices *InvoicesNamespace
	// Jira: `client.jira`.
	Jira *JiraNamespace
	// KV: `client.kv`.
	KV *KVNamespace
	// Leases: `client.leases`.
	Leases *LeasesNamespace
	// Linear: `client.linear`.
	Linear *LinearNamespace
	// LogWorkspaces: `client.logWorkspaces`.
	LogWorkspaces *LogWorkspacesNamespace
	// ManagedAccounts: `client.managedAccounts`.
	ManagedAccounts *ManagedAccountsNamespace
	// MetricAlerts: `client.metricAlerts`.
	MetricAlerts *MetricAlertsNamespace
	// Moment: `client.moment`.
	Moment *MomentNamespace
	// Msteams: `client.msteams`.
	Msteams *MsteamsNamespace
	// NetworkFlows: `client.networkFlows`.
	NetworkFlows *NetworkFlowsNamespace
	// Orgs: `client.orgs`.
	Orgs *OrgsNamespace
	// Orphans: `client.orphans`.
	Orphans *OrphansNamespace
	// Ownership: `client.ownership`.
	Ownership *OwnershipNamespace
	// Pages: `client.pages`.
	Pages *PagesNamespace
	// Posture: `client.posture`.
	Posture *PostureNamespace
	// Probes: `client.probes`.
	Probes *ProbesNamespace
	// Profile: `client.profile`.
	Profile *ProfileNamespace
	// Quotas: `client.quotas`.
	Quotas *QuotasNamespace
	// Resources: `client.resources`.
	Resources *ResourcesNamespace
	// Rightsizing: `client.rightsizing`.
	Rightsizing *RightsizingNamespace
	// SavedCostFilters: `client.savedCostFilters`.
	SavedCostFilters *SavedCostFiltersNamespace
	// Schedules: `client.schedules`.
	Schedules *SchedulesNamespace
	// Search: `client.search`.
	Search *SearchNamespace
	// SessionRecordings: `client.sessionRecordings`.
	SessionRecordings *SessionRecordingsNamespace
	// SFTP: `client.sftp`.
	SFTP *SFTPNamespace
	// SharedConsoles: `client.sharedConsoles`.
	SharedConsoles *SharedConsolesNamespace
	// Slack: `client.slack`.
	Slack *SlackNamespace
	// SQL: `client.sql`.
	SQL *SQLNamespace
	// SSHFanout: `client.sshFanout`.
	SSHFanout *SSHFanoutNamespace
	// SSHKeys: `client.sshKeys`.
	SSHKeys *SSHKeysNamespace
	// SSHTunnels: `client.sshTunnels`.
	SSHTunnels *SSHTunnelsNamespace
	// Status: `client.status`.
	Status *StatusNamespace
	// StatusIncidents: `client.statusIncidents`.
	StatusIncidents *StatusIncidentsNamespace
	// StatusPages: `client.statusPages`.
	StatusPages *StatusPagesNamespace
	// Storage: `client.storage`.
	Storage *StorageNamespace
	// TagPolicy: `client.tagPolicy`.
	TagPolicy *TagPolicyNamespace
	// Team: `client.team`.
	Team *TeamNamespace
	// WorkflowApprovals: `client.workflowApprovals`.
	WorkflowApprovals *WorkflowApprovalsNamespace
	// Workflows: `client.workflows`.
	Workflows *WorkflowsNamespace
}

// NewAPIV1Client builds a client. With no options it talks to
// https://app.infrawrench.com anonymously, which is rarely what you want: pass
// WithAPIKey, and WithOrgID if you would rather not repeat the organization id
// on every call.
func NewAPIV1Client(opts ...ClientOption) *APIV1Client {
	t := newTransport(opts)
	c := &APIV1Client{t: t}
	c.AccessRequests = newAccessRequestsNamespace(t)
	c.Accounts = newAccountsNamespace(t)
	c.Agents = newAgentsNamespace(t)
	c.AlertRules = newAlertRulesNamespace(t)
	c.APIKeys = newAPIKeysNamespace(t)
	c.Artifacts = newArtifactsNamespace(t)
	c.Associations = newAssociationsNamespace(t)
	c.AuditLogs = newAuditLogsNamespace(t)
	c.Auth = newAuthNamespace(t)
	c.Bastions = newBastionsNamespace(t)
	c.Billing = newBillingNamespace(t)
	c.BillingRules = newBillingRulesNamespace(t)
	c.BlastRadius = newBlastRadiusNamespace(t)
	c.Budgets = newBudgetsNamespace(t)
	c.BusinessMetrics = newBusinessMetricsNamespace(t)
	c.ChangeFreezes = newChangeFreezesNamespace(t)
	c.Changes = newChangesNamespace(t)
	c.Commitments = newCommitmentsNamespace(t)
	c.Config = newConfigNamespace(t)
	c.Connect = newConnectNamespace(t)
	c.CostAlerts = newCostAlertsNamespace(t)
	c.CostAnnotations = newCostAnnotationsNamespace(t)
	c.CostCentres = newCostCentresNamespace(t)
	c.CostExports = newCostExportsNamespace(t)
	c.CostReportFolders = newCostReportFoldersNamespace(t)
	c.CostReportNotifications = newCostReportNotificationsNamespace(t)
	c.CostReports = newCostReportsNamespace(t)
	c.CostScenarios = newCostScenariosNamespace(t)
	c.Costs = newCostsNamespace(t)
	c.CredentialHygiene = newCredentialHygieneNamespace(t)
	c.Credits = newCreditsNamespace(t)
	c.Currency = newCurrencyNamespace(t)
	c.CustomGraphs = newCustomGraphsNamespace(t)
	c.Dashboards = newDashboardsNamespace(t)
	c.DependencyGraph = newDependencyGraphNamespace(t)
	c.Deployments = newDeploymentsNamespace(t)
	c.Digest = newDigestNamespace(t)
	c.DNS = newDNSNamespace(t)
	c.Docker = newDockerNamespace(t)
	c.EnvironmentDiff = newEnvironmentDiffNamespace(t)
	c.Expiring = newExpiringNamespace(t)
	c.Invitations = newInvitationsNamespace(t)
	c.Invoices = newInvoicesNamespace(t)
	c.Jira = newJiraNamespace(t)
	c.KV = newKVNamespace(t)
	c.Leases = newLeasesNamespace(t)
	c.Linear = newLinearNamespace(t)
	c.LogWorkspaces = newLogWorkspacesNamespace(t)
	c.ManagedAccounts = newManagedAccountsNamespace(t)
	c.MetricAlerts = newMetricAlertsNamespace(t)
	c.Moment = newMomentNamespace(t)
	c.Msteams = newMsteamsNamespace(t)
	c.NetworkFlows = newNetworkFlowsNamespace(t)
	c.Orgs = newOrgsNamespace(t)
	c.Orphans = newOrphansNamespace(t)
	c.Ownership = newOwnershipNamespace(t)
	c.Pages = newPagesNamespace(t)
	c.Posture = newPostureNamespace(t)
	c.Probes = newProbesNamespace(t)
	c.Profile = newProfileNamespace(t)
	c.Quotas = newQuotasNamespace(t)
	c.Resources = newResourcesNamespace(t)
	c.Rightsizing = newRightsizingNamespace(t)
	c.SavedCostFilters = newSavedCostFiltersNamespace(t)
	c.Schedules = newSchedulesNamespace(t)
	c.Search = newSearchNamespace(t)
	c.SessionRecordings = newSessionRecordingsNamespace(t)
	c.SFTP = newSFTPNamespace(t)
	c.SharedConsoles = newSharedConsolesNamespace(t)
	c.Slack = newSlackNamespace(t)
	c.SQL = newSQLNamespace(t)
	c.SSHFanout = newSSHFanoutNamespace(t)
	c.SSHKeys = newSSHKeysNamespace(t)
	c.SSHTunnels = newSSHTunnelsNamespace(t)
	c.Status = newStatusNamespace(t)
	c.StatusIncidents = newStatusIncidentsNamespace(t)
	c.StatusPages = newStatusPagesNamespace(t)
	c.Storage = newStorageNamespace(t)
	c.TagPolicy = newTagPolicyNamespace(t)
	c.Team = newTeamNamespace(t)
	c.WorkflowApprovals = newWorkflowApprovalsNamespace(t)
	c.Workflows = newWorkflowsNamespace(t)
	return c
}

// BaseURL reports the normalized base URL every call is sent to.
func (c *APIV1Client) BaseURL() string {
	return c.t.baseURL()
}

// AccessRequestsNamespace is `client.accessRequests`.
type AccessRequestsNamespace struct {
	t *transport
}

func newAccessRequestsNamespace(t *transport) *AccessRequestsNamespace {
	n := &AccessRequestsNamespace{t: t}
	return n
}

// AccessRequestsApproveParams holds the parameters for
// `client.accessRequests.approve`.
type AccessRequestsApproveParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	RequestID string
	// Body: the JSON request body.
	Body *AccessDecision
}

// Approve: Approve an access request
//
// Opens the elevation window: the requester holds the requested permissions from
// now until `grantExpiresAt`, on every surface at once (HTTP, the WebSocket
// gateway, chat, MCP tools). Two rules are enforced here and cannot be bypassed:
// you cannot decide your own request (403 `self_approval`), and you cannot grant
// a permission you do not hold yourself (403 `exceeds_approver`) — denying
// something aimed higher than you is allowed. Deciding a request that has
// already been decided or has timed out is a 409. Audit-logged.
//
// _Requires permission: `access:approve`._
//
// POST /api/org/{orgId}/access-requests/{requestId}/approve
//
// Raises on 400: Bad request
//
// Raises on 403: Self-approval, or granting beyond the approver's own
// permissions
//
// Raises on 404: Not found
//
// Raises on 409: Already decided, or the request timed out
func (n *AccessRequestsNamespace) Approve(ctx context.Context, params AccessRequestsApproveParams, opts ...RequestOption) (*AccessRequest, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/access-requests/{requestId}/approve")
	r.setPath("orgId", params.OrgID)
	r.setPath("requestId", params.RequestID)
	r.setJSONBody(params.Body)
	var out *AccessRequest
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccessRequestsCatalogParams holds the parameters for
// `client.accessRequests.catalog`.
//
// Every field is optional; pass nil to take the defaults.
type AccessRequestsCatalogParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Catalog: Permissions a request may ask for
//
// The server's permission catalog plus the subset the caller already holds and
// the bounds on grant length. Served rather than hard-coded in clients so a
// picker cannot drift from what the server will accept.
//
// _Requires permission: `access:read`._
//
// GET /api/org/{orgId}/access-requests/catalog
func (n *AccessRequestsNamespace) Catalog(ctx context.Context, params *AccessRequestsCatalogParams, opts ...RequestOption) (*AccessRequestCatalog, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/access-requests/catalog")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *AccessRequestCatalog
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccessRequestsCreateParams holds the parameters for
// `client.accessRequests.create`.
//
// Every field is optional; pass nil to take the defaults.
type AccessRequestsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *AccessRequestCreate
}

// Create: Request elevated access
//
// Ask for specific permissions, for a specific number of minutes, with a reason.
// Rejected with 400 when the caller's role already grants every permission asked
// for — that is almost always a wrong permission string rather than a real
// request. Fans out to push, Slack (with Approve/Deny buttons) and Microsoft
// Teams under the Pages opt-in. Audit-logged.
//
// _Requires permission: `access:request`._
//
// POST /api/org/{orgId}/access-requests
//
// Raises on 400: Bad request
func (n *AccessRequestsNamespace) Create(ctx context.Context, params *AccessRequestsCreateParams, opts ...RequestOption) (*AccessRequest, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/access-requests")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *AccessRequest
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccessRequestsDenyParams holds the parameters for
// `client.accessRequests.deny`.
type AccessRequestsDenyParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	RequestID string
	// Body: the JSON request body.
	Body *AccessDecision
}

// Deny: Deny an access request
//
// Records the refusal. Two rules are enforced here and cannot be bypassed: you
// cannot decide your own request (403 `self_approval`), and you cannot grant a
// permission you do not hold yourself (403 `exceeds_approver`) — denying
// something aimed higher than you is allowed. Deciding a request that has
// already been decided or has timed out is a 409. Audit-logged.
//
// _Requires permission: `access:approve`._
//
// POST /api/org/{orgId}/access-requests/{requestId}/deny
//
// Raises on 400: Bad request
//
// Raises on 403: Self-approval, or granting beyond the approver's own
// permissions
//
// Raises on 404: Not found
//
// Raises on 409: Already decided, or the request timed out
func (n *AccessRequestsNamespace) Deny(ctx context.Context, params AccessRequestsDenyParams, opts ...RequestOption) (*AccessRequest, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/access-requests/{requestId}/deny")
	r.setPath("orgId", params.OrgID)
	r.setPath("requestId", params.RequestID)
	r.setJSONBody(params.Body)
	var out *AccessRequest
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccessRequestsListParams holds the parameters for
// `client.accessRequests.list`.
//
// Every field is optional; pass nil to take the defaults.
type AccessRequestsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Status: `pending` (awaiting a decision), `approved`, `denied`, or
	// `expired` (nobody decided in time, or the requester withdrew it). An
	// approved row is only *granting* permissions while `active` is true.
	//
	// One of "pending", "approved", "denied", "expired".
	Status *string
	// Mine: Only the caller's own requests.
	//
	// One of "1".
	Mine *string
	// Active: Only rows granting permissions right now.
	//
	// One of "1".
	Active *string
}

// List: List access requests
//
// The organization's break-glass requests, newest first. A `pending` listing
// hides rows whose timeout has already passed, so the queue never offers a
// decision that would immediately be refused.
//
// _Requires permission: `access:read`._
//
// GET /api/org/{orgId}/access-requests
//
// Raises on 400: Bad request
func (n *AccessRequestsNamespace) List(ctx context.Context, params *AccessRequestsListParams, opts ...RequestOption) ([]AccessRequest, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/access-requests")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("status", params.Status)
		r.addQuery("mine", params.Mine)
		r.addQuery("active", params.Active)
	}
	var out []AccessRequest
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccessRequestsRevokeParams holds the parameters for
// `client.accessRequests.revoke`.
type AccessRequestsRevokeParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	RequestID string
}

// Revoke: End a live elevation early
//
// Allowed for anyone with `access:approve` and for the holder — giving back an
// elevation you no longer need must never require finding an approver. Applies
// from the next permission resolution; nothing is cached. Audit-logged.
//
// POST /api/org/{orgId}/access-requests/{requestId}/revoke
//
// Raises on 404: Not found
//
// Raises on 409: The grant is not active
func (n *AccessRequestsNamespace) Revoke(ctx context.Context, params AccessRequestsRevokeParams, opts ...RequestOption) (*AccessRequest, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/access-requests/{requestId}/revoke")
	r.setPath("orgId", params.OrgID)
	r.setPath("requestId", params.RequestID)
	var out *AccessRequest
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccessRequestsWithdrawParams holds the parameters for
// `client.accessRequests.withdraw`.
type AccessRequestsWithdrawParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	RequestID string
}

// Withdraw: Withdraw your own pending request
//
// Its own operation rather than a self-denial, so the audit trail distinguishes
// 'nobody would approve this' from 'they decided they didn't need it'.
// Audit-logged.
//
// _Requires permission: `access:request`._
//
// POST /api/org/{orgId}/access-requests/{requestId}/withdraw
//
// Raises on 404: Not found
//
// Raises on 409: Already decided or expired
func (n *AccessRequestsNamespace) Withdraw(ctx context.Context, params AccessRequestsWithdrawParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/access-requests/{requestId}/withdraw")
	r.setPath("orgId", params.OrgID)
	r.setPath("requestId", params.RequestID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsNamespace is `client.accounts`.
type AccountsNamespace struct {
	t *transport

	// Credentials: `client.accounts.credentials`.
	Credentials *AccountsCredentialsNamespace
	// Plugins: `client.accounts.plugins`.
	Plugins *AccountsPluginsNamespace
	// Preflight: `client.accounts.preflight`.
	Preflight *AccountsPreflightNamespace
	// SyncType: `client.accounts.syncType`.
	SyncType *AccountsSyncTypeNamespace
}

func newAccountsNamespace(t *transport) *AccountsNamespace {
	n := &AccountsNamespace{t: t}
	n.Credentials = newAccountsCredentialsNamespace(t)
	n.Plugins = newAccountsPluginsNamespace(t)
	n.Preflight = newAccountsPreflightNamespace(t)
	n.SyncType = newAccountsSyncTypeNamespace(t)
	return n
}

// AccountsCreateParams holds the parameters for `client.accounts.create`.
type AccountsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body AccountsCreateRequest
}

// Create: Create an account
//
// Stores encrypted credentials and triggers a first sync. `syncError` is set if
// the initial sync failed (the account row is still created).
//
// _Requires permission: `accounts:write`._
//
// POST /api/org/{orgId}/accounts
//
// Raises on 400: Bad request
//
// Raises on 402: Payment required — the organization's plan does not include
// this
func (n *AccountsNamespace) Create(ctx context.Context, params AccountsCreateParams, opts ...RequestOption) (*CreateAccountResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/accounts")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CreateAccountResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsDeleteParams holds the parameters for `client.accounts.delete`.
type AccountsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete an account
//
// _Requires permission: `accounts:delete`._
//
// DELETE /api/org/{orgId}/accounts/{id}
func (n *AccountsNamespace) Delete(ctx context.Context, params AccountsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/accounts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsDetailParams holds the parameters for `client.accounts.detail`.
type AccountsDetailParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Detail: Account metadata + resource type list
//
// _Requires permission: `accounts:read`._
//
// GET /api/org/{orgId}/accounts/{id}/detail
//
// Raises on 404: Not found
func (n *AccountsNamespace) Detail(ctx context.Context, params AccountsDetailParams, opts ...RequestOption) (*AccountDetail, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/accounts/{id}/detail")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *AccountDetail
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsExportTerraformParams holds the parameters for
// `client.accounts.exportTerraform`.
type AccountsExportTerraformParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// ExportTerraform: Generate Terraform HCL for the account's stored inventory
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/accounts/{id}/export-terraform
//
// Raises on 404: Not found
func (n *AccountsNamespace) ExportTerraform(ctx context.Context, params AccountsExportTerraformParams, opts ...RequestOption) (*TerraformExport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/accounts/{id}/export-terraform")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *TerraformExport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsListParams holds the parameters for `client.accounts.list`.
//
// Every field is optional; pass nil to take the defaults.
type AccountsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List accounts in this organization
//
// _Requires permission: `accounts:read`._
//
// GET /api/org/{orgId}/accounts
func (n *AccountsNamespace) List(ctx context.Context, params *AccountsListParams, opts ...RequestOption) ([]Account, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/accounts")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []Account
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsResourcesParams holds the parameters for `client.accounts.resources`.
type AccountsResourcesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// TopLevelOnly: If `true`, only resources with no `parentResourceId` are
	// returned.
	//
	// One of "true", "false".
	TopLevelOnly *string
}

// Resources: List cached resources for an account
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/accounts/{id}/resources
func (n *AccountsNamespace) Resources(ctx context.Context, params AccountsResourcesParams, opts ...RequestOption) ([]Resource, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/accounts/{id}/resources")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.addQuery("topLevelOnly", params.TopLevelOnly)
	var out []Resource
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsSyncParams holds the parameters for `client.accounts.sync`.
type AccountsSyncParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Sync: Sync all resource types for an account
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/accounts/{id}/sync
func (n *AccountsNamespace) Sync(ctx context.Context, params AccountsSyncParams, opts ...RequestOption) (*SyncResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/accounts/{id}/sync")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *SyncResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsUpdateParams holds the parameters for `client.accounts.update`.
type AccountsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body UpdateAccountRequest
}

// Update: Update an account (rename and/or change bastion binding)
//
// _Requires permission: `accounts:write`._
//
// PATCH /api/org/{orgId}/accounts/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *AccountsNamespace) Update(ctx context.Context, params AccountsUpdateParams, opts ...RequestOption) (*UpdatedAccount, error) {
	r := newRequest(http.MethodPatch, "/api/org/{orgId}/accounts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *UpdatedAccount
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsCredentialsNamespace is `client.accounts.credentials`.
type AccountsCredentialsNamespace struct {
	t *transport
}

func newAccountsCredentialsNamespace(t *transport) *AccountsCredentialsNamespace {
	n := &AccountsCredentialsNamespace{t: t}
	return n
}

// AccountsCredentialsGetParams holds the parameters for
// `client.accounts.credentials.get`.
type AccountsCredentialsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Fetch the decrypted credentials for an account
//
// Returns the credentials map as it was originally submitted. Sensitive — gate
// access carefully.
//
// _Requires permission: `secrets:read`._
//
// GET /api/org/{orgId}/accounts/{id}/credentials
//
// Raises on 404: Not found
func (n *AccountsCredentialsNamespace) Get(ctx context.Context, params AccountsCredentialsGetParams, opts ...RequestOption) (map[string]string, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/accounts/{id}/credentials")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out map[string]string
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsCredentialsUpdateParams holds the parameters for
// `client.accounts.credentials.update`.
type AccountsCredentialsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body AccountsCredentialsUpdateRequest
}

// Update: Rotate the credentials an account uses to talk to the upstream
// provider
//
// Replaces the encrypted credentials blob in place. Used to swap a stale or
// narrowly-scoped token for a freshly-minted one without recreating the account
// (preserves existing resources, pins, dashboards, sync history).
//
// _Requires permission: `secrets:write`._
//
// PUT /api/org/{orgId}/accounts/{id}/credentials
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *AccountsCredentialsNamespace) Update(ctx context.Context, params AccountsCredentialsUpdateParams, opts ...RequestOption) (*AccountsCredentialsUpdateResponse, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/accounts/{id}/credentials")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *AccountsCredentialsUpdateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsPluginsNamespace is `client.accounts.plugins`.
type AccountsPluginsNamespace struct {
	t *transport
}

func newAccountsPluginsNamespace(t *transport) *AccountsPluginsNamespace {
	n := &AccountsPluginsNamespace{t: t}
	return n
}

// AccountsPluginsListParams holds the parameters for
// `client.accounts.plugins.list`.
//
// Every field is optional; pass nil to take the defaults.
type AccountsPluginsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List installed plugins and their credential fields
//
// _Requires permission: `accounts:read`._
//
// GET /api/org/{orgId}/accounts/plugins
func (n *AccountsPluginsNamespace) List(ctx context.Context, params *AccountsPluginsListParams, opts ...RequestOption) ([]PluginSummary, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/accounts/plugins")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []PluginSummary
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsPluginsPolicyTemplateParams holds the parameters for
// `client.accounts.plugins.policyTemplate`.
type AccountsPluginsPolicyTemplateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	// Capabilities: Comma-separated capability ids, e.g. `resources,costs`.
	Capabilities *string
}

// PolicyTemplate: Generate a least-privilege credential template for a plugin
//
// Returns the paste-ready credential document (IAM policy JSON, custom role
// YAML, token template…) scoped to the requested capability ids. Omitting
// `capabilities` (or sending it empty) selects every declared capability; any
// unknown capability id is rejected with 400. 400 also for plugins that don't
// provide a template.
//
// _Requires permission: `accounts:read`._
//
// GET /api/org/{orgId}/accounts/plugins/{pluginId}/policy-template
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *AccountsPluginsNamespace) PolicyTemplate(ctx context.Context, params AccountsPluginsPolicyTemplateParams, opts ...RequestOption) (*PolicyTemplateResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/accounts/plugins/{pluginId}/policy-template")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.addQuery("capabilities", params.Capabilities)
	var out *PolicyTemplateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsPreflightNamespace is `client.accounts.preflight`.
type AccountsPreflightNamespace struct {
	t *transport
}

func newAccountsPreflightNamespace(t *transport) *AccountsPreflightNamespace {
	n := &AccountsPreflightNamespace{t: t}
	return n
}

// AccountsPreflightCreateParams holds the parameters for
// `client.accounts.preflight.create`.
type AccountsPreflightCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body PreflightRequest
}

// Create: Probe credentials before creating an account
//
// Runs the plugin's per-capability permission checks against the submitted
// credentials. Nothing is stored — use it from the add-account flow before
// committing.
//
// _Requires permission: `accounts:write`._
//
// POST /api/org/{orgId}/accounts/preflight
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *AccountsPreflightNamespace) Create(ctx context.Context, params AccountsPreflightCreateParams, opts ...RequestOption) (*PreflightReport, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/accounts/preflight")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *PreflightReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsPreflightPostOrgOrgIDAccountsIDPreflightParams holds the parameters
// for `client.accounts.preflight.postOrgOrgIdAccountsIdPreflight`.
type AccountsPreflightPostOrgOrgIDAccountsIDPreflightParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// PostOrgOrgIDAccountsIDPreflight: Re-run credential preflight on a stored
// account
//
// _Requires permission: `accounts:write`._
//
// POST /api/org/{orgId}/accounts/{id}/preflight
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *AccountsPreflightNamespace) PostOrgOrgIDAccountsIDPreflight(ctx context.Context, params AccountsPreflightPostOrgOrgIDAccountsIDPreflightParams, opts ...RequestOption) (*PreflightReport, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/accounts/{id}/preflight")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *PreflightReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AccountsSyncTypeNamespace is `client.accounts.syncType`.
type AccountsSyncTypeNamespace struct {
	t *transport
}

func newAccountsSyncTypeNamespace(t *transport) *AccountsSyncTypeNamespace {
	n := &AccountsSyncTypeNamespace{t: t}
	return n
}

// AccountsSyncTypeCreateParams holds the parameters for
// `client.accounts.syncType.create`.
type AccountsSyncTypeCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID  *string
	ID     string
	TypeID ResourceTypeID
}

// Create: Sync a single resource type and return its resources
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/accounts/{id}/sync-type/{typeId}
//
// Raises on 404: Not found
//
// Raises on 500: Server error
func (n *AccountsSyncTypeNamespace) Create(ctx context.Context, params AccountsSyncTypeCreateParams, opts ...RequestOption) ([]SyncedResource, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/accounts/{id}/sync-type/{typeId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setPath("typeId", params.TypeID)
	var out []SyncedResource
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsNamespace is `client.agents`.
type AgentsNamespace struct {
	t *transport

	// Sessions: `client.agents.sessions`.
	Sessions *AgentsSessionsNamespace
	// Settings: `client.agents.settings`.
	Settings *AgentsSettingsNamespace
}

func newAgentsNamespace(t *transport) *AgentsNamespace {
	n := &AgentsNamespace{t: t}
	n.Sessions = newAgentsSessionsNamespace(t)
	n.Settings = newAgentsSettingsNamespace(t)
	return n
}

// AgentsAccountsParams holds the parameters for `client.agents.accounts`.
//
// Every field is optional; pass nil to take the defaults.
type AgentsAccountsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Accounts: List accounts whose plugins can create agent VMs
//
// _Requires permission: `accounts:read`._
//
// GET /api/org/{orgId}/agents/accounts
func (n *AgentsNamespace) Accounts(ctx context.Context, params *AgentsAccountsParams, opts ...RequestOption) ([]AgentVMAccount, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/agents/accounts")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []AgentVMAccount
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsSessionsNamespace is `client.agents.sessions`.
type AgentsSessionsNamespace struct {
	t *transport
}

func newAgentsSessionsNamespace(t *transport) *AgentsSessionsNamespace {
	n := &AgentsSessionsNamespace{t: t}
	return n
}

// AgentsSessionsCreateParams holds the parameters for
// `client.agents.sessions.create`.
type AgentsSessionsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateAgentSession
}

// Create: Create an agent session
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/agents/sessions
//
// Raises on 400: Bad request
func (n *AgentsSessionsNamespace) Create(ctx context.Context, params AgentsSessionsCreateParams, opts ...RequestOption) (*AgentSession, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/agents/sessions")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *AgentSession
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsSessionsDeleteParams holds the parameters for
// `client.agents.sessions.delete`.
type AgentsSessionsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete an agent session and destroy its VM
//
// _Requires permission: `resources:delete`._
//
// DELETE /api/org/{orgId}/agents/sessions/{id}
//
// Raises on 404: Not found
//
// Raises on 502: The provider refused to delete the VM
func (n *AgentsSessionsNamespace) Delete(ctx context.Context, params AgentsSessionsDeleteParams, opts ...RequestOption) (*AgentsSessionsDeleteResponse, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/agents/sessions/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *AgentsSessionsDeleteResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsSessionsListParams holds the parameters for
// `client.agents.sessions.list`.
//
// Every field is optional; pass nil to take the defaults.
type AgentsSessionsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List agent sessions
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/agents/sessions
func (n *AgentsSessionsNamespace) List(ctx context.Context, params *AgentsSessionsListParams, opts ...RequestOption) ([]AgentSession, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/agents/sessions")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []AgentSession
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsSessionsOpenParams holds the parameters for
// `client.agents.sessions.open`.
type AgentsSessionsOpenParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Open: Return the command and working directory for an agent session
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/agents/sessions/{id}/open
//
// Raises on 404: Not found
func (n *AgentsSessionsNamespace) Open(ctx context.Context, params AgentsSessionsOpenParams, opts ...RequestOption) (*AgentsSessionsOpenResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/agents/sessions/{id}/open")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *AgentsSessionsOpenResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsSessionsReconcileParams holds the parameters for
// `client.agents.sessions.reconcile`.
type AgentsSessionsReconcileParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Reconcile: Return reconciliation branch metadata
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/agents/sessions/{id}/reconcile
//
// Raises on 404: Not found
func (n *AgentsSessionsNamespace) Reconcile(ctx context.Context, params AgentsSessionsReconcileParams, opts ...RequestOption) (*AgentsSessionsReconcileResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/agents/sessions/{id}/reconcile")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *AgentsSessionsReconcileResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsSettingsNamespace is `client.agents.settings`.
type AgentsSettingsNamespace struct {
	t *transport
}

func newAgentsSettingsNamespace(t *transport) *AgentsSettingsNamespace {
	n := &AgentsSettingsNamespace{t: t}
	return n
}

// AgentsSettingsGetParams holds the parameters for `client.agents.settings.get`.
//
// Every field is optional; pass nil to take the defaults.
type AgentsSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get saved Agents defaults
//
// _Requires permission: `accounts:read`._
//
// GET /api/org/{orgId}/agents/settings
func (n *AgentsSettingsNamespace) Get(ctx context.Context, params *AgentsSettingsGetParams, opts ...RequestOption) (*AgentSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/agents/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *AgentSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AgentsSettingsUpdateParams holds the parameters for
// `client.agents.settings.update`.
type AgentsSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *AgentSettings
}

// Update: Save Agents defaults
//
// _Requires permission: `accounts:write`._
//
// PUT /api/org/{orgId}/agents/settings
func (n *AgentsSettingsNamespace) Update(ctx context.Context, params AgentsSettingsUpdateParams, opts ...RequestOption) (*AgentSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/agents/settings")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *AgentSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AlertRulesNamespace is `client.alertRules`.
type AlertRulesNamespace struct {
	t *transport

	// Deliveries: `client.alertRules.deliveries`.
	Deliveries *AlertRulesDeliveriesNamespace
}

func newAlertRulesNamespace(t *transport) *AlertRulesNamespace {
	n := &AlertRulesNamespace{t: t}
	n.Deliveries = newAlertRulesDeliveriesNamespace(t)
	return n
}

// AlertRulesAdoptDefaultsParams holds the parameters for
// `client.alertRules.adoptDefaults`.
//
// Every field is optional; pass nil to take the defaults.
type AlertRulesAdoptDefaultsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// AdoptDefaults: Persist the default rule so it can be edited
//
// A no-op when the organization already has rules.
//
// POST /api/org/{orgId}/alert-rules/adopt-defaults
//
// Raises on 403: Forbidden
func (n *AlertRulesNamespace) AdoptDefaults(ctx context.Context, params *AlertRulesAdoptDefaultsParams, opts ...RequestOption) (*AlertRulesAdoptDefaultsResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/alert-rules/adopt-defaults")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *AlertRulesAdoptDefaultsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AlertRulesGetParams holds the parameters for `client.alertRules.get`.
//
// Every field is optional; pass nil to take the defaults.
type AlertRulesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's alert routing rules
//
// Returns the rules in evaluation order, plus the channels and accounts a rule
// can name so a client can render destinations by name. An organization that has
// saved no rules gets the synthesized default with `usingDefaults: true`.
//
// GET /api/org/{orgId}/alert-rules
//
// Raises on 403: Forbidden
func (n *AlertRulesNamespace) Get(ctx context.Context, params *AlertRulesGetParams, opts ...RequestOption) (*AlertRulesResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/alert-rules")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *AlertRulesResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AlertRulesUpdateParams holds the parameters for `client.alertRules.update`.
//
// Every field is optional; pass nil to take the defaults.
type AlertRulesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *AlertRulesUpdateRequest
}

// Update: Replace the organization's alert routing rules
//
// Whole-list replacement in one transaction. Order is part of the meaning — a
// rule is only correct relative to the ones above it — so a reorder applied as
// several requests would leave a window in which alerts route somewhere nobody
// asked for. Positions are re-derived from array order.
//
// PUT /api/org/{orgId}/alert-rules
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
func (n *AlertRulesNamespace) Update(ctx context.Context, params *AlertRulesUpdateParams, opts ...RequestOption) (*AlertRulesUpdateResponse, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/alert-rules")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *AlertRulesUpdateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AlertRulesDeliveriesNamespace is `client.alertRules.deliveries`.
type AlertRulesDeliveriesNamespace struct {
	t *transport
}

func newAlertRulesDeliveriesNamespace(t *transport) *AlertRulesDeliveriesNamespace {
	n := &AlertRulesDeliveriesNamespace{t: t}
	return n
}

// AlertRulesDeliveriesAckParams holds the parameters for
// `client.alertRules.deliveries.ack`.
type AlertRulesDeliveriesAckParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Ack: Acknowledge an alert, cancelling its escalation
//
// A conditional update: only a delivery still in `awaiting_ack` can move, so two
// people pressing at once produce one acknowledgement and an alert that already
// escalated cannot be retroactively silenced.
//
// POST /api/org/{orgId}/alert-rules/deliveries/{id}/ack
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *AlertRulesDeliveriesNamespace) Ack(ctx context.Context, params AlertRulesDeliveriesAckParams, opts ...RequestOption) (*AlertRulesDeliveriesAckResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/alert-rules/deliveries/{id}/ack")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *AlertRulesDeliveriesAckResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AlertRulesDeliveriesCancelParams holds the parameters for
// `client.alertRules.deliveries.cancel`.
//
// Every field is optional; pass nil to take the defaults.
type AlertRulesDeliveriesCancelParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *AlertRulesDeliveriesCancelRequest
}

// Cancel: Drop held or awaiting-acknowledgement deliveries
//
// POST /api/org/{orgId}/alert-rules/deliveries/cancel
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
func (n *AlertRulesDeliveriesNamespace) Cancel(ctx context.Context, params *AlertRulesDeliveriesCancelParams, opts ...RequestOption) (*AlertRulesDeliveriesCancelResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/alert-rules/deliveries/cancel")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *AlertRulesDeliveriesCancelResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AlertRulesDeliveriesListParams holds the parameters for
// `client.alertRules.deliveries.list`.
//
// Every field is optional; pass nil to take the defaults.
type AlertRulesDeliveriesListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	Limit *int64
}

// List: List recent held and escalating alerts
//
// Only alerts a rule created follow-up work for appear here: one held by quiet
// hours, or one waiting on an acknowledgement. An alert that went straight out
// leaves no row.
//
// GET /api/org/{orgId}/alert-rules/deliveries
//
// Raises on 403: Forbidden
func (n *AlertRulesDeliveriesNamespace) List(ctx context.Context, params *AlertRulesDeliveriesListParams, opts ...RequestOption) ([]AlertDelivery, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/alert-rules/deliveries")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("limit", params.Limit)
	}
	var out []AlertDelivery
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// APIKeysNamespace is `client.apiKeys`.
type APIKeysNamespace struct {
	t *transport
}

func newAPIKeysNamespace(t *transport) *APIKeysNamespace {
	n := &APIKeysNamespace{t: t}
	return n
}

// APIKeysCreateParams holds the parameters for `client.apiKeys.create`.
type APIKeysCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateAPIKeyRequest
}

// Create: Create an API key (plaintext returned once)
//
// POST /api/org/{orgId}/api-keys
func (n *APIKeysNamespace) Create(ctx context.Context, params APIKeysCreateParams, opts ...RequestOption) (*CreatedAPIKey, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/api-keys")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CreatedAPIKey
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// APIKeysListParams holds the parameters for `client.apiKeys.list`.
//
// Every field is optional; pass nil to take the defaults.
type APIKeysListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List API keys (no plaintext)
//
// GET /api/org/{orgId}/api-keys
func (n *APIKeysNamespace) List(ctx context.Context, params *APIKeysListParams, opts ...RequestOption) ([]APIKey, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/api-keys")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []APIKey
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// APIKeysRevokeParams holds the parameters for `client.apiKeys.revoke`.
type APIKeysRevokeParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Revoke: Revoke an API key
//
// POST /api/org/{orgId}/api-keys/{id}/revoke
func (n *APIKeysNamespace) Revoke(ctx context.Context, params APIKeysRevokeParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/api-keys/{id}/revoke")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// APIKeysRotateParams holds the parameters for `client.apiKeys.rotate`.
type APIKeysRotateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Rotate: Rotate an API key (revokes old, returns new)
//
// POST /api/org/{orgId}/api-keys/{id}/rotate
//
// Raises on 404: Not found
func (n *APIKeysNamespace) Rotate(ctx context.Context, params APIKeysRotateParams, opts ...RequestOption) (*CreatedAPIKey, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/api-keys/{id}/rotate")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CreatedAPIKey
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ArtifactsNamespace is `client.artifacts`.
type ArtifactsNamespace struct {
	t *transport
}

func newArtifactsNamespace(t *transport) *ArtifactsNamespace {
	n := &ArtifactsNamespace{t: t}
	return n
}

// ArtifactsListParams holds the parameters for `client.artifacts.list`.
type ArtifactsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ArtifactsListRequest
}

// List: List artifact-registry items for a resource
//
// _Requires permission: `storage:read`._
//
// POST /api/org/{orgId}/artifacts/list
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 500: Server error
func (n *ArtifactsNamespace) List(ctx context.Context, params ArtifactsListParams, opts ...RequestOption) (JSONObject, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/artifacts/list")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out JSONObject
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AssociationsNamespace is `client.associations`.
type AssociationsNamespace struct {
	t *transport
}

func newAssociationsNamespace(t *transport) *AssociationsNamespace {
	n := &AssociationsNamespace{t: t}
	return n
}

// AssociationsCreateParams holds the parameters for
// `client.associations.create`.
type AssociationsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body AssociationRequest
}

// Create: Wire one resource's output into another resource's secret field
//
// _Requires permission: `secrets:write`._
//
// POST /api/org/{orgId}/associations
//
// Raises on 404: Not found
func (n *AssociationsNamespace) Create(ctx context.Context, params AssociationsCreateParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/associations")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AssociationsLiteralParams holds the parameters for
// `client.associations.literal`.
type AssociationsLiteralParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body LiteralAssociationRequest
}

// Literal: Set a secret field to a literal plaintext value
//
// _Requires permission: `secrets:write`._
//
// POST /api/org/{orgId}/associations/literal
//
// Raises on 404: Not found
func (n *AssociationsNamespace) Literal(ctx context.Context, params AssociationsLiteralParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/associations/literal")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AuditLogsNamespace is `client.auditLogs`.
type AuditLogsNamespace struct {
	t *transport
}

func newAuditLogsNamespace(t *transport) *AuditLogsNamespace {
	n := &AuditLogsNamespace{t: t}
	return n
}

// AuditLogsGetParams holds the parameters for `client.auditLogs.get`.
//
// Every field is optional; pass nil to take the defaults.
type AuditLogsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	Page       *int64
	PageSize   *int64
	Action     *string
	EntityType *string
	UserID     *string
	APIKeyID   *string
	From       *string
	To         *string
}

// Get: List audit log entries (paginated, filterable)
//
// _Requires permission: `audit:read`._
//
// GET /api/org/{orgId}/audit-logs
func (n *AuditLogsNamespace) Get(ctx context.Context, params *AuditLogsGetParams, opts ...RequestOption) (*AuditResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/audit-logs")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("page", params.Page)
		r.addQuery("pageSize", params.PageSize)
		r.addQuery("action", params.Action)
		r.addQuery("entityType", params.EntityType)
		r.addQuery("userId", params.UserID)
		r.addQuery("apiKeyId", params.APIKeyID)
		r.addQuery("from", params.From)
		r.addQuery("to", params.To)
	}
	var out *AuditResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// AuthNamespace is `client.auth`.
type AuthNamespace struct {
	t *transport
}

func newAuthNamespace(t *transport) *AuthNamespace {
	n := &AuthNamespace{t: t}
	return n
}

// Me: Current session + onboarding status
//
// GET /api/auth/me
//
// Raises on 401: Unauthenticated
func (n *AuthNamespace) Me(ctx context.Context, opts ...RequestOption) (*Session, error) {
	r := newRequest(http.MethodGet, "/api/auth/me")
	var out *Session
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// Orgs: Organizations the current user belongs to
//
// GET /api/auth/orgs
//
// Raises on 401: Unauthenticated
func (n *AuthNamespace) Orgs(ctx context.Context, opts ...RequestOption) ([]OrgMembership, error) {
	r := newRequest(http.MethodGet, "/api/auth/orgs")
	var out []OrgMembership
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BastionsNamespace is `client.bastions`.
type BastionsNamespace struct {
	t *transport
}

func newBastionsNamespace(t *transport) *BastionsNamespace {
	n := &BastionsNamespace{t: t}
	return n
}

// BastionsCreateParams holds the parameters for `client.bastions.create`.
type BastionsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateBastionRequest
}

// Create: Register a new bastion (returns the enrollment token once)
//
// _Requires permission: `bastions:write`._
//
// POST /api/org/{orgId}/bastions
//
// Raises on 400: Bad request
func (n *BastionsNamespace) Create(ctx context.Context, params BastionsCreateParams, opts ...RequestOption) (*CreateBastionResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/bastions")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CreateBastionResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BastionsDeleteParams holds the parameters for `client.bastions.delete`.
type BastionsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Revoke a bastion — accounts referencing it have their bastion binding
// cleared
//
// _Requires permission: `bastions:write`._
//
// DELETE /api/org/{orgId}/bastions/{id}
//
// Raises on 404: Not found
func (n *BastionsNamespace) Delete(ctx context.Context, params BastionsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/bastions/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BastionsListParams holds the parameters for `client.bastions.list`.
//
// Every field is optional; pass nil to take the defaults.
type BastionsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List bastion agents registered to this org
//
// _Requires permission: `bastions:read`._
//
// GET /api/org/{orgId}/bastions
func (n *BastionsNamespace) List(ctx context.Context, params *BastionsListParams, opts ...RequestOption) ([]Bastion, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/bastions")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []Bastion
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingNamespace is `client.billing`.
type BillingNamespace struct {
	t *transport

	// Capacity: `client.billing.capacity`.
	Capacity *BillingCapacityNamespace
}

func newBillingNamespace(t *transport) *BillingNamespace {
	n := &BillingNamespace{t: t}
	n.Capacity = newBillingCapacityNamespace(t)
	return n
}

// BillingCheckoutParams holds the parameters for `client.billing.checkout`.
//
// Every field is optional; pass nil to take the defaults.
type BillingCheckoutParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Checkout: Start a Stripe Checkout session
//
// Rejected with 400 for complimentary organizations — they are never billed.
//
// _Requires permission: `billing:write`._
//
// POST /api/org/{orgId}/billing/checkout
//
// Raises on 400: Bad request
//
// Raises on 500: Server error
func (n *BillingNamespace) Checkout(ctx context.Context, params *BillingCheckoutParams, opts ...RequestOption) (*StripeRedirectURL, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/billing/checkout")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *StripeRedirectURL
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingPortalParams holds the parameters for `client.billing.portal`.
//
// Every field is optional; pass nil to take the defaults.
type BillingPortalParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Portal: Get a Stripe customer portal URL
//
// _Requires permission: `billing:write`._
//
// POST /api/org/{orgId}/billing/portal
//
// Raises on 404: Not found
func (n *BillingNamespace) Portal(ctx context.Context, params *BillingPortalParams, opts ...RequestOption) (*StripeRedirectURL, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/billing/portal")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *StripeRedirectURL
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingStatusParams holds the parameters for `client.billing.status`.
//
// Every field is optional; pass nil to take the defaults.
type BillingStatusParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Status: Get the org's billing status (complimentary flag + subscription or
// `null`)
//
// _Requires permission: `billing:read`._
//
// GET /api/org/{orgId}/billing/status
func (n *BillingNamespace) Status(ctx context.Context, params *BillingStatusParams, opts ...RequestOption) (*BillingStatus, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/billing/status")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *BillingStatus
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingCapacityNamespace is `client.billing.capacity`.
type BillingCapacityNamespace struct {
	t *transport
}

func newBillingCapacityNamespace(t *transport) *BillingCapacityNamespace {
	n := &BillingCapacityNamespace{t: t}
	return n
}

// BillingCapacityCheckoutParams holds the parameters for
// `client.billing.capacity.checkout`.
//
// Every field is optional; pass nil to take the defaults.
type BillingCapacityCheckoutParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *CapacityCheckoutRequest
}

// Checkout: Start a Stripe Checkout session for prepaid capacity slots
//
// A capacity slot is one seat bought outright for a fixed term instead of rented
// monthly, and it grants paid-plan access on its own. This is a one-time
// payment, so the seats are granted by the `checkout.session.completed` webhook
// once Stripe confirms the payment — a 200 here only means the buyer was sent to
// a payment page. Rejected with 400 for complimentary organizations, and 503
// when the deployment has no one-time capacity price configured.
//
// POST /api/org/{orgId}/billing/capacity/checkout
//
// Raises on 400: Bad request
//
// Raises on 500: Server error
//
// Raises on 503: A backing service this endpoint depends on is not available
func (n *BillingCapacityNamespace) Checkout(ctx context.Context, params *BillingCapacityCheckoutParams, opts ...RequestOption) (*StripeRedirectURL, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/billing/capacity/checkout")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *StripeRedirectURL
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingRulesNamespace is `client.billingRules`.
type BillingRulesNamespace struct {
	t *transport
}

func newBillingRulesNamespace(t *transport) *BillingRulesNamespace {
	n := &BillingRulesNamespace{t: t}
	return n
}

// BillingRulesCreateParams holds the parameters for
// `client.billingRules.create`.
type BillingRulesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body BillingRuleInput
}

// Create: Create a billing rule
//
// Requires `org:settings:write` rather than `costs:write`: a billing rule
// changes every figure the organisation reports about itself, which is a
// governance act on the scale of stating an exchange rate, not the scale of
// saving a report.
//
// _Requires permission: `org:settings:write`._
//
// POST /api/org/{orgId}/billing-rules
//
// Raises on 400: Bad request
//
// Raises on 409: Conflict
func (n *BillingRulesNamespace) Create(ctx context.Context, params BillingRulesCreateParams, opts ...RequestOption) (*BillingRule, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/billing-rules")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *BillingRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingRulesDeleteParams holds the parameters for
// `client.billingRules.delete`.
type BillingRulesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a billing rule
//
// Nothing cascades and nothing is restated: no adjustment was ever written into
// stored cost data, so the next read simply computes without this rule.
//
// _Requires permission: `org:settings:write`._
//
// DELETE /api/org/{orgId}/billing-rules/{id}
//
// Raises on 404: Not found
func (n *BillingRulesNamespace) Delete(ctx context.Context, params BillingRulesDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/billing-rules/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingRulesGetParams holds the parameters for `client.billingRules.get`.
type BillingRulesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a billing rule
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/billing-rules/{id}
//
// Raises on 404: Not found
func (n *BillingRulesNamespace) Get(ctx context.Context, params BillingRulesGetParams, opts ...RequestOption) (*BillingRule, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/billing-rules/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *BillingRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingRulesListParams holds the parameters for `client.billingRules.list`.
//
// Every field is optional; pass nil to take the defaults.
type BillingRulesListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List billing rules in evaluation order
//
// Billing rules are the organisation's own adjustments to collected spend — a
// markup that recovers shared overhead, a discount negotiated outside the
// provider's pricing, a shared cluster reallocated onto the teams that use it.
//
// **They are applied at query time and never written into stored cost data.**
// Collected spend stays exactly what the provider reported, so it can still be
// reconciled against an invoice, and editing or deleting a rule restates
// nothing.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/billing-rules
func (n *BillingRulesNamespace) List(ctx context.Context, params *BillingRulesListParams, opts ...RequestOption) ([]BillingRule, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/billing-rules")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []BillingRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BillingRulesUpdateParams holds the parameters for
// `client.billingRules.update`.
type BillingRulesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body BillingRuleInput
}

// Update: Update a billing rule
//
// A full replace, `enabled` included — switching a markup off is an edit of the
// rule, so there is one audited action for “this rule changed” rather than two.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/billing-rules/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *BillingRulesNamespace) Update(ctx context.Context, params BillingRulesUpdateParams, opts ...RequestOption) (*BillingRule, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/billing-rules/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *BillingRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BlastRadiusNamespace is `client.blastRadius`.
type BlastRadiusNamespace struct {
	t *transport
}

func newBlastRadiusNamespace(t *transport) *BlastRadiusNamespace {
	n := &BlastRadiusNamespace{t: t}
	return n
}

// BlastRadiusGetParams holds the parameters for `client.blastRadius.get`.
type BlastRadiusGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ResourceID ResourceID
}

// Get: What breaks if this resource is deleted
//
// An impact report for one resource, assembled from the dependency graph walked
// inbound, network flow attribution, and the org objects that name the resource
// without depending on it (dashboards, custom graphs, probes, status pages,
// metric alerts, leases, schedules, saved log queries, workflows, and its
// recorded owner).
//
// The endpoint answers 200 with a partial report rather than failing when a
// source is unavailable; `unchecked` says which, in prose.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/blast-radius
//
// Raises on 400: Missing resourceId
func (n *BlastRadiusNamespace) Get(ctx context.Context, params BlastRadiusGetParams, opts ...RequestOption) (*BlastRadiusReport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/blast-radius")
	r.setPath("orgId", params.OrgID)
	r.addQuery("resourceId", params.ResourceID)
	var out *BlastRadiusReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BudgetsNamespace is `client.budgets`.
type BudgetsNamespace struct {
	t *transport
}

func newBudgetsNamespace(t *transport) *BudgetsNamespace {
	n := &BudgetsNamespace{t: t}
	return n
}

// BudgetsCreateParams holds the parameters for `client.budgets.create`.
type BudgetsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body BudgetInput
}

// Create: Create a budget
//
// POST /api/org/{orgId}/budgets
//
// Raises on 400: Bad request
func (n *BudgetsNamespace) Create(ctx context.Context, params BudgetsCreateParams, opts ...RequestOption) (*BudgetFull, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/budgets")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *BudgetFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BudgetsDeleteParams holds the parameters for `client.budgets.delete`.
type BudgetsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a budget
//
// DELETE /api/org/{orgId}/budgets/{id}
//
// Raises on 404: Not found
func (n *BudgetsNamespace) Delete(ctx context.Context, params BudgetsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/budgets/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BudgetsEventsParams holds the parameters for `client.budgets.events`.
type BudgetsEventsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Events: Alert event history for a budget
//
// GET /api/org/{orgId}/budgets/{id}/events
//
// Raises on 404: Not found
func (n *BudgetsNamespace) Events(ctx context.Context, params BudgetsEventsParams, opts ...RequestOption) ([]BudgetAlertEvent, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/budgets/{id}/events")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out []BudgetAlertEvent
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BudgetsGetParams holds the parameters for `client.budgets.get`.
type BudgetsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a budget with current-month status
//
// GET /api/org/{orgId}/budgets/{id}
//
// Raises on 404: Not found
func (n *BudgetsNamespace) Get(ctx context.Context, params BudgetsGetParams, opts ...RequestOption) (*BudgetFull, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/budgets/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *BudgetFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BudgetsListParams holds the parameters for `client.budgets.list`.
//
// Every field is optional; pass nil to take the defaults.
type BudgetsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List budgets with current-month actuals and forecasts
//
// GET /api/org/{orgId}/budgets
func (n *BudgetsNamespace) List(ctx context.Context, params *BudgetsListParams, opts ...RequestOption) ([]BudgetWithStatus, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/budgets")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []BudgetWithStatus
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BudgetsUpdateParams holds the parameters for `client.budgets.update`.
type BudgetsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body BudgetInput
}

// Update: Update a budget
//
// PUT /api/org/{orgId}/budgets/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *BudgetsNamespace) Update(ctx context.Context, params BudgetsUpdateParams, opts ...RequestOption) (*BudgetFull, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/budgets/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *BudgetFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsNamespace is `client.businessMetrics`.
type BusinessMetricsNamespace struct {
	t *transport

	// Get: `client.businessMetrics.get`.
	Get *BusinessMetricsGetNamespace
	// Values: `client.businessMetrics.values`.
	Values *BusinessMetricsValuesNamespace
}

func newBusinessMetricsNamespace(t *transport) *BusinessMetricsNamespace {
	n := &BusinessMetricsNamespace{t: t}
	n.Get = newBusinessMetricsGetNamespace(t)
	n.Values = newBusinessMetricsValuesNamespace(t)
	return n
}

// BusinessMetricsCreateParams holds the parameters for
// `client.businessMetrics.create`.
type BusinessMetricsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body BusinessMetricInput
}

// Create: Create a business metric
//
// Keys must be unique per organization among live metrics — they are how
// workflows and the CLI address the metric. A key collision is a 409.
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/business-metrics
//
// Raises on 400: Bad request
//
// Raises on 409: A live metric already uses this key.
func (n *BusinessMetricsNamespace) Create(ctx context.Context, params BusinessMetricsCreateParams, opts ...RequestOption) (*BusinessMetric, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/business-metrics")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *BusinessMetric
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsDeleteParams holds the parameters for
// `client.businessMetrics.delete`.
type BusinessMetricsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Metric id or key
	ID string
}

// Delete: Delete a business metric
//
// Soft delete. Not refused when a dashboard card references the metric, unlike a
// saved cost filter: a unit-cost card whose metric is gone fails its query and
// says so, whereas a card that quietly reverted to plain spend would be a chart
// claiming to be something it is not.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/business-metrics/{id}
//
// Raises on 404: Not found
func (n *BusinessMetricsNamespace) Delete(ctx context.Context, params BusinessMetricsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/business-metrics/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsUnitCostsParams holds the parameters for
// `client.businessMetrics.unitCosts`.
type BusinessMetricsUnitCostsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Metric id or key
	ID string
	// Body: the JSON request body.
	Body UnitCostQueryRequest
}

// UnitCosts: Query unit costs or margin
//
// Divide spend by the metric, bucketed as asked. Three properties of the answer
// are worth knowing before reading it:
//
// - **The ratio is computed at the requested bucket**, from a summed numerator
// and a summed denominator — never a mean of daily ratios, which weights a quiet
// day as heavily as a peak one. The same holds for `overallValue`. - **A missing or non-positive denominator is a gap** (`value: null` with a
// `gap` reason), never 0 and never infinite. - **Currencies are never merged.** Spend in a currency with no stated rate
// keeps its own series rather than being dropped or added to another.
//
// There is no `groupBy`: a per-group ratio would need a per-group denominator,
// and dividing each service's spend by the whole customer count produces numbers
// that do not sum to the real one.
//
// _Requires permission: `costs:read`._
//
// POST /api/org/{orgId}/business-metrics/{id}/unit-costs
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *BusinessMetricsNamespace) UnitCosts(ctx context.Context, params BusinessMetricsUnitCostsParams, opts ...RequestOption) (*UnitCostQueryResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/business-metrics/{id}/unit-costs")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *UnitCostQueryResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsUpdateParams holds the parameters for
// `client.businessMetrics.update`.
type BusinessMetricsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Metric id or key
	ID string
	// Body: the JSON request body.
	Body BusinessMetricInput
}

// Update: Update a business metric
//
// Replaces the whole definition. Changing `key` never orphans history — values
// are keyed on the metric's id — but it does break a workflow still writing to
// the old key, which is why the key is separate from the display name in the
// first place.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/business-metrics/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: A live metric already uses this key.
func (n *BusinessMetricsNamespace) Update(ctx context.Context, params BusinessMetricsUpdateParams, opts ...RequestOption) (*BusinessMetric, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/business-metrics/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *BusinessMetric
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsGetNamespace is `client.businessMetrics.get`.
type BusinessMetricsGetNamespace struct {
	t *transport
}

func newBusinessMetricsGetNamespace(t *transport) *BusinessMetricsGetNamespace {
	n := &BusinessMetricsGetNamespace{t: t}
	return n
}

// BusinessMetricsGetGetParams holds the parameters for
// `client.businessMetrics.get.get`.
//
// Every field is optional; pass nil to take the defaults.
type BusinessMetricsGetGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List business metrics
//
// The organization's declared denominators, by key, each with the range of days
// it has values for.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/business-metrics
func (n *BusinessMetricsGetNamespace) Get(ctx context.Context, params *BusinessMetricsGetGetParams, opts ...RequestOption) (*BusinessMetricsGetGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/business-metrics")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *BusinessMetricsGetGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsGetGetOrgOrgIDBusinessMetricsIDParams holds the parameters for
// `client.businessMetrics.get.getOrgOrgIdBusinessMetricsId`.
type BusinessMetricsGetGetOrgOrgIDBusinessMetricsIDParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Metric id or key
	ID string
}

// GetOrgOrgIDBusinessMetricsID: Get a business metric
//
// `id` accepts either the metric's id or its key.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/business-metrics/{id}
//
// Raises on 404: Not found
func (n *BusinessMetricsGetNamespace) GetOrgOrgIDBusinessMetricsID(ctx context.Context, params BusinessMetricsGetGetOrgOrgIDBusinessMetricsIDParams, opts ...RequestOption) (*BusinessMetric, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/business-metrics/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *BusinessMetric
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsValuesNamespace is `client.businessMetrics.values`.
type BusinessMetricsValuesNamespace struct {
	t *transport
}

func newBusinessMetricsValuesNamespace(t *transport) *BusinessMetricsValuesNamespace {
	n := &BusinessMetricsValuesNamespace{t: t}
	return n
}

// BusinessMetricsValuesCreateParams holds the parameters for
// `client.businessMetrics.values.create`.
type BusinessMetricsValuesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Metric id or key
	ID string
	// Body: the JSON request body.
	Body BusinessMetricValuesInput
}

// Create: Report metric values
//
// Write a batch of days. **Re-reporting a day restates it rather than
// accumulating**, which is what makes a nightly job safe to retry. Nothing lands
// unless the whole batch validates, so a bad row is a 400 rather than half a
// month restated. The same guarantees back `infra.businessMetrics.write(...)` in
// a workflow — both go through one validator.
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/business-metrics/{id}/values
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *BusinessMetricsValuesNamespace) Create(ctx context.Context, params BusinessMetricsValuesCreateParams, opts ...RequestOption) (*BusinessMetricsValuesCreateResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/business-metrics/{id}/values")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *BusinessMetricsValuesCreateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// BusinessMetricsValuesGetParams holds the parameters for
// `client.businessMetrics.values.get`.
type BusinessMetricsValuesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Metric id or key
	ID string
	// Limit: Default 90.
	Limit *int64
}

// Get: List a metric's reported values
//
// Newest day first.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/business-metrics/{id}/values
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *BusinessMetricsValuesNamespace) Get(ctx context.Context, params BusinessMetricsValuesGetParams, opts ...RequestOption) (*BusinessMetricsValuesGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/business-metrics/{id}/values")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.addQuery("limit", params.Limit)
	var out *BusinessMetricsValuesGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangeFreezesNamespace is `client.changeFreezes`.
type ChangeFreezesNamespace struct {
	t *transport
}

func newChangeFreezesNamespace(t *transport) *ChangeFreezesNamespace {
	n := &ChangeFreezesNamespace{t: t}
	return n
}

// ChangeFreezesCreateParams holds the parameters for
// `client.changeFreezes.create`.
type ChangeFreezesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ChangeFreezeInput
}

// Create: Declare a change freeze window
//
// While the freeze is in effect, destructive actions (resource deletion,
// destructive plugin actions, secret-version destroys, deployment rollbacks)
// return `423` unless explicitly overridden by a caller with `freezes:override`.
//
// _Requires permission: `freezes:write`._
//
// POST /api/org/{orgId}/change-freezes
//
// Raises on 400: Bad request
func (n *ChangeFreezesNamespace) Create(ctx context.Context, params ChangeFreezesCreateParams, opts ...RequestOption) (*ChangeFreeze, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/change-freezes")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ChangeFreeze
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangeFreezesDeleteParams holds the parameters for
// `client.changeFreezes.delete`.
type ChangeFreezesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a change freeze window
//
// _Requires permission: `freezes:write`._
//
// DELETE /api/org/{orgId}/change-freezes/{id}
//
// Raises on 404: Not found
func (n *ChangeFreezesNamespace) Delete(ctx context.Context, params ChangeFreezesDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/change-freezes/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangeFreezesEndParams holds the parameters for `client.changeFreezes.end`.
type ChangeFreezesEndParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// End: End a change freeze now
//
// _Requires permission: `freezes:write`._
//
// POST /api/org/{orgId}/change-freezes/{id}/end
//
// Raises on 404: Not found
func (n *ChangeFreezesNamespace) End(ctx context.Context, params ChangeFreezesEndParams, opts ...RequestOption) (*ChangeFreeze, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/change-freezes/{id}/end")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *ChangeFreeze
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangeFreezesListParams holds the parameters for `client.changeFreezes.list`.
//
// Every field is optional; pass nil to take the defaults.
type ChangeFreezesListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List change freeze windows, newest first
//
// _Requires permission: `freezes:read`._
//
// GET /api/org/{orgId}/change-freezes
func (n *ChangeFreezesNamespace) List(ctx context.Context, params *ChangeFreezesListParams, opts ...RequestOption) ([]ChangeFreeze, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/change-freezes")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []ChangeFreeze
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangeFreezesStatusParams holds the parameters for
// `client.changeFreezes.status`.
//
// Every field is optional; pass nil to take the defaults.
type ChangeFreezesStatusParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Status: The freeze currently in effect, if any
//
// Returns the active freeze window (active, started, not yet past its end time)
// or `freeze: null`. Clients poll this to show the freeze banner and pre-warn
// before destructive actions.
//
// _Requires permission: `freezes:read`._
//
// GET /api/org/{orgId}/change-freezes/status
func (n *ChangeFreezesNamespace) Status(ctx context.Context, params *ChangeFreezesStatusParams, opts ...RequestOption) (*ChangeFreezeStatus, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/change-freezes/status")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *ChangeFreezeStatus
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangeFreezesUpdateParams holds the parameters for
// `client.changeFreezes.update`.
type ChangeFreezesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body ChangeFreezeInput
}

// Update: Update a change freeze window
//
// _Requires permission: `freezes:write`._
//
// PUT /api/org/{orgId}/change-freezes/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ChangeFreezesNamespace) Update(ctx context.Context, params ChangeFreezesUpdateParams, opts ...RequestOption) (*ChangeFreeze, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/change-freezes/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *ChangeFreeze
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangesNamespace is `client.changes`.
type ChangesNamespace struct {
	t *transport

	// AlertSettings: `client.changes.alertSettings`.
	AlertSettings *ChangesAlertSettingsNamespace
}

func newChangesNamespace(t *transport) *ChangesNamespace {
	n := &ChangesNamespace{t: t}
	n.AlertSettings = newChangesAlertSettingsNamespace(t)
	return n
}

// ChangesGetParams holds the parameters for `client.changes.get`.
//
// Every field is optional; pass nil to take the defaults.
type ChangesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	Page       *int64
	PageSize   *int64
	AccountID  *string
	ResourceID *string
	Kind       *ResourceChangeKind
	From       *string
	To         *string
}

// Get: Org-wide change timeline (paginated, filterable)
//
// Change events recorded by the resource poller: each poll cycle diffs the
// freshly fetched state against the stored snapshot and records resources that
// appeared, changed a stored field, or disappeared upstream. Cross-provider by
// construction — the diff runs on the generic stored record, so every plugin's
// resources show up here.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/changes
//
// Raises on 400: Bad request
func (n *ChangesNamespace) Get(ctx context.Context, params *ChangesGetParams, opts ...RequestOption) (*ResourceChangeFeedResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/changes")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("page", params.Page)
		r.addQuery("pageSize", params.PageSize)
		r.addQuery("accountId", params.AccountID)
		r.addQuery("resourceId", params.ResourceID)
		r.addQuery("kind", params.Kind)
		r.addQuery("from", params.From)
		r.addQuery("to", params.To)
	}
	var out *ResourceChangeFeedResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangesResourceParams holds the parameters for `client.changes.resource`.
type ChangesResourceParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ResourceID string
	Limit      *int64
}

// Resource: Change timeline for one resource
//
// Recent change events for a single resource, newest first. The resource id
// travels as a query parameter because composite ids contain slashes and colons.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/changes/resource
//
// Raises on 400: Bad request
func (n *ChangesNamespace) Resource(ctx context.Context, params ChangesResourceParams, opts ...RequestOption) (*ResourceChangeListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/changes/resource")
	r.setPath("orgId", params.OrgID)
	r.addQuery("resourceId", params.ResourceID)
	r.addQuery("limit", params.Limit)
	var out *ResourceChangeListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangesAlertSettingsNamespace is `client.changes.alertSettings`.
type ChangesAlertSettingsNamespace struct {
	t *transport
}

func newChangesAlertSettingsNamespace(t *transport) *ChangesAlertSettingsNamespace {
	n := &ChangesAlertSettingsNamespace{t: t}
	return n
}

// ChangesAlertSettingsGetParams holds the parameters for
// `client.changes.alertSettings.get`.
//
// Every field is optional; pass nil to take the defaults.
type ChangesAlertSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's resource-drift alert filter
//
// Drift notifications are batched: at most one message per organization per
// `cooldownMinutes`, covering every change since the previous one. These
// settings decide which changes count and how often a message may go out. Who
// receives it is the `resourceDrift` opt-in on push preferences, Slack channels
// and Teams webhooks — off by default on all three.
//
// GET /api/org/{orgId}/changes/alert-settings
func (n *ChangesAlertSettingsNamespace) Get(ctx context.Context, params *ChangesAlertSettingsGetParams, opts ...RequestOption) (*DriftAlertSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/changes/alert-settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *DriftAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ChangesAlertSettingsUpdateParams holds the parameters for
// `client.changes.alertSettings.update`.
//
// Every field is optional; pass nil to take the defaults.
type ChangesAlertSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *DriftAlertSettingsUpdate
}

// Update: Update the organization's resource-drift alert filter
//
// Every field is optional so a single toggle can be saved on its own.
// `cooldownMinutes` is floored at 5: below the poller's own cycle the
// notification rate would follow the sync rate again, which is what the batching
// exists to prevent.
//
// PUT /api/org/{orgId}/changes/alert-settings
//
// Raises on 400: Bad request
func (n *ChangesAlertSettingsNamespace) Update(ctx context.Context, params *ChangesAlertSettingsUpdateParams, opts ...RequestOption) (*DriftAlertSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/changes/alert-settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *DriftAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CommitmentsNamespace is `client.commitments`.
type CommitmentsNamespace struct {
	t *transport
}

func newCommitmentsNamespace(t *transport) *CommitmentsNamespace {
	n := &CommitmentsNamespace{t: t}
	return n
}

// CommitmentsGetParams holds the parameters for `client.commitments.get`.
//
// Every field is optional; pass nil to take the defaults.
type CommitmentsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Reservations, savings plans and committed-use discounts
//
// The organization's purchased commitments — reserved instances, savings plans,
// committed-use discounts — with three derived readings.
//
// **Coverage** is a range, not a number: the broad ratio counts every uncovered
// usage dollar in the denominator (a lower bound — egress and per-request
// charges can never be committed against), the narrow ratio only uncovered usage
// in cells where a commitment demonstrably landed (an upper bound). Accounts
// whose plugin cannot distinguish charge types are excluded and listed; a scope
// where every account is excluded reports unavailable, not 0%.
//
// **Utilization** is measured only over days cost data was actually collected —
// a collection gap is reported as missing days, never counted as idle
// commitment. Unit-denominated commitments (GCP) report null with a reason,
// never 0%. Azure's own reported utilization rides on each holding separately
// and is never blended with the derived figure.
//
// **The planner** recommends committing at the p10 floor of daily uncovered
// spend, gated on presence, trend, floor and materiality. Savings are quoted
// against published "up to" discount rates and marked as such. Nothing is ever
// purchased automatically.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/commitments
func (n *CommitmentsNamespace) Get(ctx context.Context, params *CommitmentsGetParams, opts ...RequestOption) (*CommitmentsFeed, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/commitments")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CommitmentsFeed
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ConfigNamespace is `client.config`.
type ConfigNamespace struct {
	t *transport
}

func newConfigNamespace(t *transport) *ConfigNamespace {
	n := &ConfigNamespace{t: t}
	return n
}

// ConfigApplyParams holds the parameters for `client.config.apply`.
type ConfigApplyParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body OrgConfigRequest
}

// Apply: Apply a configuration document
//
// Applies the document in a single transaction and returns the plan that was
// executed — all or nothing, so a failure never leaves the organization halfway
// between two configurations.
//
// Requires the write permission of every section the document carries, so this
// cannot be used to reach past a role that withholds one.
//
// _Requires permission: `config:write`._
//
// POST /api/org/{orgId}/config/apply
//
// Raises on 400: Bad request
//
// Raises on 402: Payment required — the organization's plan does not include
// this
//
// Raises on 403: Forbidden
func (n *ConfigNamespace) Apply(ctx context.Context, params ConfigApplyParams, opts ...RequestOption) (*OrgConfigApplyResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/config/apply")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OrgConfigApplyResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ConfigExportParams holds the parameters for `client.config.export`.
//
// Every field is optional; pass nil to take the defaults.
type ConfigExportParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Sections: Comma-separated subset of sections to export. Defaults to all
	// of: budgets, customGraphs, workflows, dashboards, metricAlerts, probes,
	// costCentres, tagPolicy, alertSettings.
	Sections *string
}

// Export: Export the organization's configuration as one document
//
// Dashboards, workflows, custom graphs, budgets, metric alerts, synthetic
// probes, cost centres, the tag policy and the org-wide alert settings,
// addressed by stable keys rather than row ids so the result applies to any
// organization.
//
// Credentials, accounts, resources and workflow signing secrets are never
// included. Ordering is stable, so re-exporting an unchanged organization
// produces the same bytes — commit it to git and the diff is the change.
//
// Requires the read permission of every section exported; it refuses rather than
// silently omitting one, because a partial document applied in `replace` mode
// would delete what the exporter could not see.
//
// _Requires permission: `config:read`._
//
// GET /api/org/{orgId}/config/export
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
func (n *ConfigNamespace) Export(ctx context.Context, params *ConfigExportParams, opts ...RequestOption) (*OrgConfigDocument, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/config/export")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("sections", params.Sections)
	}
	var out *OrgConfigDocument
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ConfigPlanParams holds the parameters for `client.config.plan`.
type ConfigPlanParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body OrgConfigRequest
}

// Plan: Preview what applying a document would do
//
// The dry run: validates the document, resolves its cross-references against
// this organization, and returns the create/update/delete/unchanged plan without
// writing anything. Read-only, so a reviewer with read access can run it on a
// pull request.
//
// _Requires permission: `config:read`._
//
// POST /api/org/{orgId}/config/plan
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
func (n *ConfigNamespace) Plan(ctx context.Context, params ConfigPlanParams, opts ...RequestOption) (*OrgConfigPlan, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/config/plan")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OrgConfigPlan
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ConnectNamespace is `client.connect`.
type ConnectNamespace struct {
	t *transport
}

func newConnectNamespace(t *transport) *ConnectNamespace {
	n := &ConnectNamespace{t: t}
	return n
}

// ConnectEnvDeployParams holds the parameters for `client.connect.envDeploy`.
type ConnectEnvDeployParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ConnectEnvDeployRequest
}

// EnvDeploy: Deploy env vars from a source resource to an SSH target
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/connect/env-deploy
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ConnectNamespace) EnvDeploy(ctx context.Context, params ConnectEnvDeployParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/connect/env-deploy")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ConnectSecretExportParams holds the parameters for
// `client.connect.secretExport`.
type ConnectSecretExportParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ConnectSecretExportRequest
}

// SecretExport: Materialize source outputs as a secret in the target (e.g. K8s)
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/connect/secret-export
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ConnectNamespace) SecretExport(ctx context.Context, params ConnectSecretExportParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/connect/secret-export")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ConnectTemplatesParams holds the parameters for `client.connect.templates`.
type ConnectTemplatesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ConnectTemplatesRequest
}

// Templates: List secret-export templates and target capabilities
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/connect/templates
//
// Raises on 404: Not found
func (n *ConnectNamespace) Templates(ctx context.Context, params ConnectTemplatesParams, opts ...RequestOption) (*ConnectTemplatesResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/connect/templates")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ConnectTemplatesResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAlertsNamespace is `client.costAlerts`.
type CostAlertsNamespace struct {
	t *transport

	// Get: `client.costAlerts.get`.
	Get *CostAlertsGetNamespace
}

func newCostAlertsNamespace(t *transport) *CostAlertsNamespace {
	n := &CostAlertsNamespace{t: t}
	n.Get = newCostAlertsGetNamespace(t)
	return n
}

// CostAlertsCreateParams holds the parameters for `client.costAlerts.create`.
type CostAlertsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostAlertInput
}

// Create: Create a change-based cost alert
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-alerts
//
// Raises on 400: Bad request
func (n *CostAlertsNamespace) Create(ctx context.Context, params CostAlertsCreateParams, opts ...RequestOption) (*CostAlert, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-alerts")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostAlert
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAlertsDeleteParams holds the parameters for `client.costAlerts.delete`.
type CostAlertsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a cost alert
//
// Soft delete. Fired events disappear from the org-wide event feed with it.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/cost-alerts/{id}
//
// Raises on 404: Not found
func (n *CostAlertsNamespace) Delete(ctx context.Context, params CostAlertsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-alerts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAlertsEventsParams holds the parameters for `client.costAlerts.events`.
//
// Every field is optional; pass nil to take the defaults.
type CostAlertsEventsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	AlertID *string
	Limit   *int64
}

// Events: List recently fired cost-alert events
//
// Newest first. Optionally scoped to one alert with ?alertId=; an unknown
// alertId is a 404, distinct from an alert that simply has no events yet.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-alerts/events
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostAlertsNamespace) Events(ctx context.Context, params *CostAlertsEventsParams, opts ...RequestOption) (*CostAlertsEventsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-alerts/events")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("alertId", params.AlertID)
		r.addQuery("limit", params.Limit)
	}
	var out *CostAlertsEventsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAlertsUpdateParams holds the parameters for `client.costAlerts.update`.
type CostAlertsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CostAlertInput
}

// Update: Update a cost alert
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/cost-alerts/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostAlertsNamespace) Update(ctx context.Context, params CostAlertsUpdateParams, opts ...RequestOption) (*CostAlert, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-alerts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CostAlert
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAlertsGetNamespace is `client.costAlerts.get`.
type CostAlertsGetNamespace struct {
	t *transport
}

func newCostAlertsGetNamespace(t *transport) *CostAlertsGetNamespace {
	n := &CostAlertsGetNamespace{t: t}
	return n
}

// CostAlertsGetGetParams holds the parameters for `client.costAlerts.get.get`.
//
// Every field is optional; pass nil to take the defaults.
type CostAlertsGetGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List change-based cost alerts
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-alerts
func (n *CostAlertsGetNamespace) Get(ctx context.Context, params *CostAlertsGetGetParams, opts ...RequestOption) (*CostAlertsGetGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-alerts")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CostAlertsGetGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAlertsGetGetOrgOrgIDCostAlertsIDParams holds the parameters for
// `client.costAlerts.get.getOrgOrgIdCostAlertsId`.
type CostAlertsGetGetOrgOrgIDCostAlertsIDParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// GetOrgOrgIDCostAlertsID: Get a cost alert
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-alerts/{id}
//
// Raises on 404: Not found
func (n *CostAlertsGetNamespace) GetOrgOrgIDCostAlertsID(ctx context.Context, params CostAlertsGetGetOrgOrgIDCostAlertsIDParams, opts ...RequestOption) (*CostAlert, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-alerts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CostAlert
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAnnotationsNamespace is `client.costAnnotations`.
type CostAnnotationsNamespace struct {
	t *transport
}

func newCostAnnotationsNamespace(t *transport) *CostAnnotationsNamespace {
	n := &CostAnnotationsNamespace{t: t}
	return n
}

// CostAnnotationsCreateParams holds the parameters for
// `client.costAnnotations.create`.
type CostAnnotationsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostAnnotationInput
}

// Create: Create a cost annotation
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-annotations
//
// Raises on 400: Bad request
func (n *CostAnnotationsNamespace) Create(ctx context.Context, params CostAnnotationsCreateParams, opts ...RequestOption) (*CostAnnotation, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-annotations")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostAnnotation
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAnnotationsDeleteParams holds the parameters for
// `client.costAnnotations.delete`.
type CostAnnotationsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a cost annotation
//
// A hard delete. A withdrawn explanation should stop being drawn, and nothing
// references a note by id.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/cost-annotations/{id}
//
// Raises on 404: Not found
func (n *CostAnnotationsNamespace) Delete(ctx context.Context, params CostAnnotationsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-annotations/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAnnotationsGetParams holds the parameters for
// `client.costAnnotations.get`.
//
// Every field is optional; pass nil to take the defaults.
type CostAnnotationsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ReportID: Scope to the notes a chart for this report should draw.
	ReportID *string
}

// Get: List cost annotations
//
// Dated notes drawn over cost charts. With `reportId`, the set a chart for that
// report draws: the org-wide notes plus that report's own. Without it, every
// annotation in the org. Annotations are an overlay — they never appear in a
// series, a total, or an axis.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-annotations
func (n *CostAnnotationsNamespace) Get(ctx context.Context, params *CostAnnotationsGetParams, opts ...RequestOption) (*CostAnnotationsGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-annotations")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("reportId", params.ReportID)
	}
	var out *CostAnnotationsGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostAnnotationsUpdateParams holds the parameters for
// `client.costAnnotations.update`.
type CostAnnotationsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CostAnnotationInput
}

// Update: Update a cost annotation
//
// Replaces the note's dates, text and scope. Moving a note between org-wide and
// one report is this same PUT with a different `costReportId`.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/cost-annotations/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostAnnotationsNamespace) Update(ctx context.Context, params CostAnnotationsUpdateParams, opts ...RequestOption) (*CostAnnotation, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-annotations/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CostAnnotation
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresNamespace is `client.costCentres`.
type CostCentresNamespace struct {
	t *transport

	// Rules: `client.costCentres.rules`.
	Rules *CostCentresRulesNamespace
}

func newCostCentresNamespace(t *transport) *CostCentresNamespace {
	n := &CostCentresNamespace{t: t}
	n.Rules = newCostCentresRulesNamespace(t)
	return n
}

// CostCentresCreateParams holds the parameters for `client.costCentres.create`.
type CostCentresCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostCentreInput
}

// Create: Create a cost centre
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-centres
//
// Raises on 400: Bad request
func (n *CostCentresNamespace) Create(ctx context.Context, params CostCentresCreateParams, opts ...RequestOption) (*CostCentre, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-centres")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostCentre
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresDeleteParams holds the parameters for `client.costCentres.delete`.
type CostCentresDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a cost centre (its allocation rules go with it)
//
// The centre's allocation rules are deleted with it, so the spend they claimed
// falls through to the next matching rule or to "Unallocated". Child centres are
// not deleted: they are re-parented onto the deleted centre's own parent (a
// root's children become roots), so a subtree keeps its shape and no ancestor's
// subtree total moves unexpectedly.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/cost-centres/{id}
//
// Raises on 404: Not found
func (n *CostCentresNamespace) Delete(ctx context.Context, params CostCentresDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-centres/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresListParams holds the parameters for `client.costCentres.list`.
//
// Every field is optional; pass nil to take the defaults.
type CostCentresListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List cost centres
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-centres
func (n *CostCentresNamespace) List(ctx context.Context, params *CostCentresListParams, opts ...RequestOption) ([]CostCentre, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-centres")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []CostCentre
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresUpdateParams holds the parameters for `client.costCentres.update`.
type CostCentresUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CostCentreInput
}

// Update: Update or move a cost centre
//
// Renames, redescribes, and/or moves a centre. Moving is `parentId` changing;
// omitting the field leaves the centre where it is. 400 when the move would
// cycle or breach the depth cap.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/cost-centres/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostCentresNamespace) Update(ctx context.Context, params CostCentresUpdateParams, opts ...RequestOption) (*CostCentre, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-centres/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CostCentre
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresRulesNamespace is `client.costCentres.rules`.
type CostCentresRulesNamespace struct {
	t *transport
}

func newCostCentresRulesNamespace(t *transport) *CostCentresRulesNamespace {
	n := &CostCentresRulesNamespace{t: t}
	return n
}

// CostCentresRulesCreateParams holds the parameters for
// `client.costCentres.rules.create`.
type CostCentresRulesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body AllocationRuleInput
}

// Create: Create an allocation rule
//
// Maps spend onto a cost centre. Rules evaluate first-match-wins by ascending
// priority against each cost row's tags, account, provider, and service.
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-centres/rules
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostCentresRulesNamespace) Create(ctx context.Context, params CostCentresRulesCreateParams, opts ...RequestOption) (*AllocationRule, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-centres/rules")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *AllocationRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresRulesDeleteParams holds the parameters for
// `client.costCentres.rules.delete`.
type CostCentresRulesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete an allocation rule
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/cost-centres/rules/{id}
//
// Raises on 404: Not found
func (n *CostCentresRulesNamespace) Delete(ctx context.Context, params CostCentresRulesDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-centres/rules/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresRulesListParams holds the parameters for
// `client.costCentres.rules.list`.
//
// Every field is optional; pass nil to take the defaults.
type CostCentresRulesListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List allocation rules in evaluation order
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-centres/rules
func (n *CostCentresRulesNamespace) List(ctx context.Context, params *CostCentresRulesListParams, opts ...RequestOption) ([]AllocationRule, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-centres/rules")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []AllocationRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresRulesSwapParams holds the parameters for
// `client.costCentres.rules.swap`.
type CostCentresRulesSwapParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SwapAllocationRulesBody
}

// Swap: Swap the priorities of two allocation rules
//
// Atomically swaps priorities so first-match-wins order can be edited without a
// half-applied pair of independent updates.
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-centres/rules/swap
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostCentresRulesNamespace) Swap(ctx context.Context, params CostCentresRulesSwapParams, opts ...RequestOption) ([]AllocationRule, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-centres/rules/swap")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out []AllocationRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostCentresRulesUpdateParams holds the parameters for
// `client.costCentres.rules.update`.
type CostCentresRulesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body AllocationRuleInput
}

// Update: Update an allocation rule
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/cost-centres/rules/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostCentresRulesNamespace) Update(ctx context.Context, params CostCentresRulesUpdateParams, opts ...RequestOption) (*AllocationRule, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-centres/rules/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *AllocationRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostExportsNamespace is `client.costExports`.
type CostExportsNamespace struct {
	t *transport
}

func newCostExportsNamespace(t *transport) *CostExportsNamespace {
	n := &CostExportsNamespace{t: t}
	return n
}

// CostExportsCreateParams holds the parameters for `client.costExports.create`.
type CostExportsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostExportInput
}

// Create: Create a cost export
//
// Credentials are required on create. They are encrypted at rest and no route
// ever returns them; responses carry a redacted `credentialHint` instead.
//
// _Requires permission: `org:settings:write`._
//
// POST /api/org/{orgId}/cost-exports
//
// Raises on 400: Bad request
func (n *CostExportsNamespace) Create(ctx context.Context, params CostExportsCreateParams, opts ...RequestOption) (*CostExport, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-exports")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostExport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostExportsDeleteParams holds the parameters for `client.costExports.delete`.
type CostExportsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a cost export
//
// Soft delete. Objects already written to the destination are left alone.
//
// _Requires permission: `org:settings:write`._
//
// DELETE /api/org/{orgId}/cost-exports/{id}
//
// Raises on 404: Not found
func (n *CostExportsNamespace) Delete(ctx context.Context, params CostExportsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-exports/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostExportsGetParams holds the parameters for `client.costExports.get`.
type CostExportsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a cost export
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-exports/{id}
//
// Raises on 404: Not found
func (n *CostExportsNamespace) Get(ctx context.Context, params CostExportsGetParams, opts ...RequestOption) (*CostExport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-exports/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CostExport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostExportsListParams holds the parameters for `client.costExports.list`.
//
// Every field is optional; pass nil to take the defaults.
type CostExportsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List scheduled cost exports
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-exports
func (n *CostExportsNamespace) List(ctx context.Context, params *CostExportsListParams, opts ...RequestOption) ([]CostExport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-exports")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []CostExport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostExportsRunParams holds the parameters for `client.costExports.run`.
type CostExportsRunParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Run: Run a cost export now
//
// Runs the export immediately against the same code path the poller uses,
// writing every period in the restatement window. Answers 200 with `status:
// "failed"` and a message rather than an error status when the destination
// rejects the write — the caller wants the reason, and the same failure is
// recorded on the export.
//
// _Requires permission: `org:settings:write`._
//
// POST /api/org/{orgId}/cost-exports/{id}/run
//
// Raises on 404: Not found
func (n *CostExportsNamespace) Run(ctx context.Context, params CostExportsRunParams, opts ...RequestOption) (*CostExportRunResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-exports/{id}/run")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CostExportRunResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostExportsUpdateParams holds the parameters for `client.costExports.update`.
type CostExportsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CostExportInput
}

// Update: Update a cost export
//
// Replaces everything but the credential. Omit
// `accessKeyId`/`secretAccessKey`/`url` to keep the stored credential; changing
// the destination type requires supplying a new one. Saving reschedules the
// export from now.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/cost-exports/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostExportsNamespace) Update(ctx context.Context, params CostExportsUpdateParams, opts ...RequestOption) (*CostExport, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-exports/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CostExport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportFoldersNamespace is `client.costReportFolders`.
type CostReportFoldersNamespace struct {
	t *transport
}

func newCostReportFoldersNamespace(t *transport) *CostReportFoldersNamespace {
	n := &CostReportFoldersNamespace{t: t}
	return n
}

// CostReportFoldersCreateParams holds the parameters for
// `client.costReportFolders.create`.
type CostReportFoldersCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostReportFolderInput
}

// Create: Create a cost-report folder
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-report-folders
//
// Raises on 400: Bad request
func (n *CostReportFoldersNamespace) Create(ctx context.Context, params CostReportFoldersCreateParams, opts ...RequestOption) (*CostReportFolder, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-report-folders")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostReportFolder
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportFoldersDeleteParams holds the parameters for
// `client.costReportFolders.delete`.
type CostReportFoldersDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a cost-report folder
//
// Never blocked by contents and never destructive to them: the folder's reports
// and immediate subfolders fall back to the top level. Deleting a folder cannot
// delete a report.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/cost-report-folders/{id}
//
// Raises on 404: Not found
func (n *CostReportFoldersNamespace) Delete(ctx context.Context, params CostReportFoldersDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-report-folders/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportFoldersListParams holds the parameters for
// `client.costReportFolders.list`.
//
// Every field is optional; pass nil to take the defaults.
type CostReportFoldersListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List cost-report folders
//
// The org's report folders as a flat list — build the tree from
// `parentFolderId`. Folders organize the Reports list and nothing else; a
// report's id, URL and dashboard cards are unchanged by where it is filed.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-report-folders
func (n *CostReportFoldersNamespace) List(ctx context.Context, params *CostReportFoldersListParams, opts ...RequestOption) ([]CostReportFolder, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-report-folders")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []CostReportFolder
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportFoldersUpdateParams holds the parameters for
// `client.costReportFolders.update`.
type CostReportFoldersUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CostReportFolderInput
}

// Update: Update a cost-report folder
//
// Rename and/or reparent. Filing a *report* is not here — that is `PUT
// /cost-reports/{id}` with a different `folderId`. Reparenting past the 3-level
// depth limit, or under the folder's own subtree, is a 400.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/cost-report-folders/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostReportFoldersNamespace) Update(ctx context.Context, params CostReportFoldersUpdateParams, opts ...RequestOption) (*CostReportFolder, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-report-folders/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CostReportFolder
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportNotificationsNamespace is `client.costReportNotifications`.
type CostReportNotificationsNamespace struct {
	t *transport
}

func newCostReportNotificationsNamespace(t *transport) *CostReportNotificationsNamespace {
	n := &CostReportNotificationsNamespace{t: t}
	return n
}

// CostReportNotificationsListParams holds the parameters for
// `client.costReportNotifications.list`.
//
// Every field is optional; pass nil to take the defaults.
type CostReportNotificationsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List every delivery schedule in the organization
//
// All reports' schedules in one call — what the CLI's schedules column reads.
// Schedules of deleted reports are excluded.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-report-notifications
func (n *CostReportNotificationsNamespace) List(ctx context.Context, params *CostReportNotificationsListParams, opts ...RequestOption) ([]ReportNotification, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-report-notifications")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []ReportNotification
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsNamespace is `client.costReports`.
type CostReportsNamespace struct {
	t *transport

	// Notifications: `client.costReports.notifications`.
	Notifications *CostReportsNotificationsNamespace
}

func newCostReportsNamespace(t *transport) *CostReportsNamespace {
	n := &CostReportsNamespace{t: t}
	n.Notifications = newCostReportsNotificationsNamespace(t)
	return n
}

// CostReportsCreateParams holds the parameters for `client.costReports.create`.
type CostReportsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostReportInput
}

// Create: Create a cost report
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-reports
//
// Raises on 400: Bad request
func (n *CostReportsNamespace) Create(ctx context.Context, params CostReportsCreateParams, opts ...RequestOption) (*CostReport, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-reports")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsDeleteParams holds the parameters for `client.costReports.delete`.
type CostReportsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a cost report
//
// Soft delete. Every dashboard card pointing at the report is removed with it —
// a card whose report is gone could only ever render as an unavailable tile.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/cost-reports/{id}
//
// Raises on 404: Not found
func (n *CostReportsNamespace) Delete(ctx context.Context, params CostReportsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-reports/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsGetParams holds the parameters for `client.costReports.get`.
type CostReportsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a cost report
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-reports/{id}
//
// Raises on 404: Not found
func (n *CostReportsNamespace) Get(ctx context.Context, params CostReportsGetParams, opts ...RequestOption) (*CostReport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-reports/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CostReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsListParams holds the parameters for `client.costReports.list`.
//
// Every field is optional; pass nil to take the defaults.
type CostReportsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List saved cost reports
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-reports
func (n *CostReportsNamespace) List(ctx context.Context, params *CostReportsListParams, opts ...RequestOption) ([]CostReport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-reports")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []CostReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsRunParams holds the parameters for `client.costReports.run`.
type CostReportsRunParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Run: Run a cost report
//
// Executes the report's saved config and returns the series, along with the
// inclusive window a relative preset resolved to. Takes no body: the report *is*
// the query, so a caller never has to reassemble its config to get the numbers.
//
// _Requires permission: `costs:read`._
//
// POST /api/org/{orgId}/cost-reports/{id}/run
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostReportsNamespace) Run(ctx context.Context, params CostReportsRunParams, opts ...RequestOption) (*CostReportRunResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-reports/{id}/run")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CostReportRunResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsUpdateParams holds the parameters for `client.costReports.update`.
type CostReportsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CostReportInput
}

// Update: Update a cost report
//
// Replaces the report's name, description, config and folder. Every dashboard
// showing the report picks up the new config — that is what referencing a report
// by id buys.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/cost-reports/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostReportsNamespace) Update(ctx context.Context, params CostReportsUpdateParams, opts ...RequestOption) (*CostReport, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-reports/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CostReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsNotificationsNamespace is `client.costReports.notifications`.
type CostReportsNotificationsNamespace struct {
	t *transport
}

func newCostReportsNotificationsNamespace(t *transport) *CostReportsNotificationsNamespace {
	n := &CostReportsNotificationsNamespace{t: t}
	return n
}

// CostReportsNotificationsCreateParams holds the parameters for
// `client.costReports.notifications.create`.
type CostReportsNotificationsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body ReportNotificationInput
}

// Create: Create a delivery schedule
//
// On its cadence the server runs the report and sends a composed text summary —
// period total (converted to the org's display currency where configured, with
// the conversion caveat), change vs the previous period, top groups, and a deep
// link. No chart images. An empty result still sends, saying so.
//
// _Requires permission: `org:settings:write`._
//
// POST /api/org/{orgId}/cost-reports/{id}/notifications
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostReportsNotificationsNamespace) Create(ctx context.Context, params CostReportsNotificationsCreateParams, opts ...RequestOption) (*ReportNotification, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-reports/{id}/notifications")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *ReportNotification
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsNotificationsDeleteParams holds the parameters for
// `client.costReports.notifications.delete`.
type CostReportsNotificationsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID          *string
	ID             string
	NotificationID string
}

// Delete: Delete a delivery schedule
//
// _Requires permission: `org:settings:write`._
//
// DELETE /api/org/{orgId}/cost-reports/{id}/notifications/{notificationId}
//
// Raises on 404: Not found
func (n *CostReportsNotificationsNamespace) Delete(ctx context.Context, params CostReportsNotificationsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-reports/{id}/notifications/{notificationId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setPath("notificationId", params.NotificationID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsNotificationsListParams holds the parameters for
// `client.costReports.notifications.list`.
type CostReportsNotificationsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// List: List a report's delivery schedules
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-reports/{id}/notifications
//
// Raises on 404: Not found
func (n *CostReportsNotificationsNamespace) List(ctx context.Context, params CostReportsNotificationsListParams, opts ...RequestOption) ([]ReportNotification, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-reports/{id}/notifications")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out []ReportNotification
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsNotificationsSendParams holds the parameters for
// `client.costReports.notifications.send`.
type CostReportsNotificationsSendParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID          *string
	ID             string
	NotificationID string
}

// Send: Send a schedule's report now
//
// Runs the report and delivers it to this schedule's destinations immediately,
// ignoring the schedule and its enabled flag. Fails with a 400 naming the reason
// when nothing could be delivered. A successful manual send clears a parked
// failure — it is the documented recovery for a partial delivery.
//
// _Requires permission: `org:settings:write`._
//
// POST /api/org/{orgId}/cost-reports/{id}/notifications/{notificationId}/send
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostReportsNotificationsNamespace) Send(ctx context.Context, params CostReportsNotificationsSendParams, opts ...RequestOption) (*ReportNotificationSendResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-reports/{id}/notifications/{notificationId}/send")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setPath("notificationId", params.NotificationID)
	var out *ReportNotificationSendResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsNotificationsTargetsParams holds the parameters for
// `client.costReports.notifications.targets`.
type CostReportsNotificationsTargetsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Targets: List the destinations a schedule can deliver to
//
// The org's live Slack channels and Teams webhooks, and whether this deployment
// can send mail. Destinations are picked from here — a schedule can only point
// at surfaces the org already connected.
//
// _Requires permission: `org:settings:write`._
//
// GET /api/org/{orgId}/cost-reports/{id}/notifications/targets
func (n *CostReportsNotificationsNamespace) Targets(ctx context.Context, params CostReportsNotificationsTargetsParams, opts ...RequestOption) (*ReportDeliveryTargets, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-reports/{id}/notifications/targets")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *ReportDeliveryTargets
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostReportsNotificationsUpdateParams holds the parameters for
// `client.costReports.notifications.update`.
type CostReportsNotificationsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID          *string
	ID             string
	NotificationID string
	// Body: the JSON request body.
	Body ReportNotificationInput
}

// Update: Update a delivery schedule
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/cost-reports/{id}/notifications/{notificationId}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *CostReportsNotificationsNamespace) Update(ctx context.Context, params CostReportsNotificationsUpdateParams, opts ...RequestOption) (*ReportNotification, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-reports/{id}/notifications/{notificationId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setPath("notificationId", params.NotificationID)
	r.setJSONBody(params.Body)
	var out *ReportNotification
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostScenariosNamespace is `client.costScenarios`.
type CostScenariosNamespace struct {
	t *transport

	// Get: `client.costScenarios.get`.
	Get *CostScenariosGetNamespace
}

func newCostScenariosNamespace(t *transport) *CostScenariosNamespace {
	n := &CostScenariosNamespace{t: t}
	n.Get = newCostScenariosGetNamespace(t)
	return n
}

// CostScenariosCreateParams holds the parameters for
// `client.costScenarios.create`.
type CostScenariosCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostScenarioModelInput
}

// Create: Create a scenario model
//
// Names must be unique per organization (case-insensitively) — the name is what
// a chart prints under its scenario line and what the CLI's `--scenario <name>`
// addresses, so two models sharing one would make both meaningless. A model
// needs at least one adjustment: an empty model changes nothing, which is the
// same as applying no scenario.
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/cost-scenarios
//
// Raises on 400: Bad request
//
// Raises on 409: A live scenario model already uses this name.
func (n *CostScenariosNamespace) Create(ctx context.Context, params CostScenariosCreateParams, opts ...RequestOption) (*CostScenarioModel, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/cost-scenarios")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostScenarioModel
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostScenariosDeleteParams holds the parameters for
// `client.costScenarios.delete`.
type CostScenariosDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a scenario model
//
// Soft delete — **refused with a 409 while anything references the model**, with
// the referents in the body. For a chart, deleting would silently drop the
// assumptions from a projection somebody is reading; for a budget it would move
// the forecast thresholds back to the bare trend, changing when people get
// paged. Detaching the referents is a deliberate step, never a side effect of
// deletion.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/cost-scenarios/{id}
//
// Raises on 404: Not found
//
// Raises on 409: Still referenced — the body lists every referent.
func (n *CostScenariosNamespace) Delete(ctx context.Context, params CostScenariosDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/cost-scenarios/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostScenariosReferentsParams holds the parameters for
// `client.costScenarios.referents`.
type CostScenariosReferentsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Referents: List a scenario model's referents
//
// Every budget, cost report and dashboard cost graph referencing this model —
// what an edit will change, and what a delete would be refused over. Budgets
// come first: they are the referents that page people.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-scenarios/{id}/referents
//
// Raises on 404: Not found
func (n *CostScenariosNamespace) Referents(ctx context.Context, params CostScenariosReferentsParams, opts ...RequestOption) (*CostScenariosReferentsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-scenarios/{id}/referents")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CostScenariosReferentsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostScenariosUpdateParams holds the parameters for
// `client.costScenarios.update`.
type CostScenariosUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CostScenarioModelInput
}

// Update: Update a scenario model
//
// Replaces the whole model. This is the high-leverage write: every chart drawing
// it, and **every budget whose forecast thresholds are measured against it**,
// uses the new numbers on its next evaluation — which for a budget can change
// which alerts fire. `GET /{id}/referents` names what a change will touch.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/cost-scenarios/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: A live scenario model already uses this name.
func (n *CostScenariosNamespace) Update(ctx context.Context, params CostScenariosUpdateParams, opts ...RequestOption) (*CostScenarioModel, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/cost-scenarios/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CostScenarioModel
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostScenariosGetNamespace is `client.costScenarios.get`.
type CostScenariosGetNamespace struct {
	t *transport
}

func newCostScenariosGetNamespace(t *transport) *CostScenariosGetNamespace {
	n := &CostScenariosGetNamespace{t: t}
	return n
}

// CostScenariosGetGetParams holds the parameters for
// `client.costScenarios.get.get`.
//
// Every field is optional; pass nil to take the defaults.
type CostScenariosGetGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List scenario models
//
// Named, reusable sets of adjustments an organization overlays on a cost
// forecast — the **known future cost a trend fit cannot see**. Pass an id as
// `POST /costs/query`'s `scenarioModelId` (alongside `forecast: true`) to get
// the adjusted projection back *beside* the unadjusted one, never instead of it.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-scenarios
func (n *CostScenariosGetNamespace) Get(ctx context.Context, params *CostScenariosGetGetParams, opts ...RequestOption) (*CostScenariosGetGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-scenarios")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CostScenariosGetGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostScenariosGetGetOrgOrgIDCostScenariosIDParams holds the parameters for
// `client.costScenarios.get.getOrgOrgIdCostScenariosId`.
type CostScenariosGetGetOrgOrgIDCostScenariosIDParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// GetOrgOrgIDCostScenariosID: Get a scenario model
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/cost-scenarios/{id}
//
// Raises on 404: Not found
func (n *CostScenariosGetNamespace) GetOrgOrgIDCostScenariosID(ctx context.Context, params CostScenariosGetGetOrgOrgIDCostScenariosIDParams, opts ...RequestOption) (*CostScenarioModel, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/cost-scenarios/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CostScenarioModel
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsNamespace is `client.costs`.
type CostsNamespace struct {
	t *transport

	// Anomalies: `client.costs.anomalies`.
	Anomalies *CostsAnomaliesNamespace
	// AnomalySettings: `client.costs.anomalySettings`.
	AnomalySettings *CostsAnomalySettingsNamespace
	// EfficiencyAlertSettings: `client.costs.efficiencyAlertSettings`.
	EfficiencyAlertSettings *CostsEfficiencyAlertSettingsNamespace
}

func newCostsNamespace(t *transport) *CostsNamespace {
	n := &CostsNamespace{t: t}
	n.Anomalies = newCostsAnomaliesNamespace(t)
	n.AnomalySettings = newCostsAnomalySettingsNamespace(t)
	n.EfficiencyAlertSettings = newCostsEfficiencyAlertSettingsNamespace(t)
	return n
}

// CostsDimensionsParams holds the parameters for `client.costs.dimensions`.
type CostsDimensionsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment", "tag-keys".
	Dimension string
	TagKey    *string
}

// Dimensions: List distinct values for a cost dimension
//
// Feeds the filter and group-by pickers. Pass dimension=tag-keys for tag keys;
// dimension=tag requires tagKey. `charge_type` answers from the fixed set of
// charge types rather than from the stored data, so the picker is populated
// before any provider has reported one.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/dimensions
//
// Raises on 400: Bad request
func (n *CostsNamespace) Dimensions(ctx context.Context, params CostsDimensionsParams, opts ...RequestOption) (*CostDimensionValues, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/dimensions")
	r.setPath("orgId", params.OrgID)
	r.addQuery("dimension", params.Dimension)
	r.addQuery("tagKey", params.TagKey)
	var out *CostDimensionValues
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsEfficiencyAlertsParams holds the parameters for
// `client.costs.efficiencyAlerts`.
//
// Every field is optional; pass nil to take the defaults.
type CostsEfficiencyAlertsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Kind: Restrict to one detector. Omitted returns all three, interleaved by
	// time.
	//
	// One of "commitment_expiry", "commitment_idle", "unit_cost_regression".
	Kind *string
	// Limit: Rows to return, newest first. Defaults to 50.
	Limit *int64
}

// EfficiencyAlerts: Recently fired efficiency alerts
//
// The three slow-lane cost alerts in one feed, newest first: commitments about
// to lapse, commitments that are not being used, and business metrics whose cost
// per unit rose. Unlike budgets, anomalies and change alerts — all of which
// compare a spend total against another spend total — these read the commitment
// calendar and the volume the spend bought, so they see the two surprises the
// other three structurally cannot.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/efficiency-alerts
//
// Raises on 400: Bad request
func (n *CostsNamespace) EfficiencyAlerts(ctx context.Context, params *CostsEfficiencyAlertsParams, opts ...RequestOption) (*CostsEfficiencyAlertsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/efficiency-alerts")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("kind", params.Kind)
		r.addQuery("limit", params.Limit)
	}
	var out *CostsEfficiencyAlertsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsQueryParams holds the parameters for `client.costs.query`.
type CostsQueryParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostQueryRequest
}

// Query: Query aggregated cost series
//
// Aggregates collected provider spend into per-bucket, per-group series for cost
// graphs. Currencies are never merged; mixed-currency orgs get one series per
// currency. Optionally returns a previous-period comparison and a trend
// forecast.
//
// `costBasis` chooses between cash and amortized money, and `chargeTypes`
// narrows which kinds of charge count. Both the comparison period and the
// forecast are computed on the same basis and charge types as the series itself.
//
// The filter can be sent structurally as `filters` or as text in the cost query
// language via `query` (`provider = 'aws' AND tag['env'] != 'dev'`). They are
// two spellings of one filter: sending both is a 400, and a query that does not
// parse is a 400 carrying the offset of the mistake.
//
// _Requires permission: `costs:read`._
//
// POST /api/org/{orgId}/costs/query
//
// Raises on 400: Bad request
func (n *CostsNamespace) Query(ctx context.Context, params CostsQueryParams, opts ...RequestOption) (*CostQueryResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/costs/query")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostQueryResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsRowsParams holds the parameters for `client.costs.rows`.
type CostsRowsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostPushRequest
}

// Rows: Push cost rows from your own systems
//
// Reports spend Infrawrench has no provider plugin for — a parsed SaaS invoice,
// an internal chargeback, a colo bill — into the same store the provider
// collectors write to, so it appears in cost graphs, dimension filters, and
// budgets alongside everything else.
//
// Rows are grouped under a caller-chosen `source`. Writes are idempotent per
// `(source, day, service, region, resourceId, tags, currency)`: pushing the same
// day again restates that day rather than adding to it, so a nightly job can
// safely re-push a trailing window. Rows pushed under a source can never
// overwrite rows a provider collector wrote.
//
// The whole batch is validated before anything is stored, so a 400 means nothing
// was written.
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/costs/rows
//
// Raises on 400: Bad request
func (n *CostsNamespace) Rows(ctx context.Context, params CostsRowsParams, opts ...RequestOption) (*CostPushResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/costs/rows")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostPushResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsShowbackParams holds the parameters for `client.costs.showback`.
//
// Every field is optional; pass nil to take the defaults.
type CostsShowbackParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// From: Defaults to 30 days ago.
	From *string
	// To: Defaults to today.
	To *string
	// Basis: Which money to sum. `cash` (the default) is what the provider
	// charged on the day it charged it; `amortized` spreads a commitment's
	// up-front fee across the term it buys. Providers that report no amortized
	// amount fall back to their cash amount.
	//
	// One of "cash", "amortized".
	Basis *string
	// Adjusted: Apply the organization's billing rules (see /billing-rules):
	// markups multiply, and a reallocation moves a centre's spend onto another
	// centre. Off by default — a chargeback report that silently showed
	// marked-up numbers is one the receiving team could not reconcile. On, the
	// response carries `adjustment` with the collected totals beside the
	// adjusted ones. Fixed-amount rules are booked onto the cost centre they
	// name (or "Unallocated" when they name none), pro-rated across the period.
	//
	// One of "true", "false".
	Adjusted *string
}

// Showback: Spend grouped by cost centre (showback)
//
// Runs the org's allocation rules over collected spend and sums per cost centre
// and currency. Spend no rule claims comes back as the "Unallocated" bucket;
// every defined centre appears even with zero spend.
//
// Cost centres nest, so the list is a depth-first tree. Each entry carries
// `totals` (spend allocated directly to it) and `subtreeTotals` (its own plus
// every descendant's) — "Engineering, of which Platform" needs both. Rules still
// evaluate first-match-wins by ascending priority against a flat list, so a row
// is allocated exactly once even when a rule targets a parent and another
// targets its child; at equal priority the more deeply nested centre wins.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/showback
//
// Raises on 400: Bad request
func (n *CostsNamespace) Showback(ctx context.Context, params *CostsShowbackParams, opts ...RequestOption) (*ShowbackReport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/showback")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("from", params.From)
		r.addQuery("to", params.To)
		r.addQuery("basis", params.Basis)
		r.addQuery("adjusted", params.Adjusted)
	}
	var out *ShowbackReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsStatusParams holds the parameters for `client.costs.status`.
//
// Every field is optional; pass nil to take the defaults.
type CostsStatusParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Status: Per-account cost collection status
//
// Which accounts support cost collection, whether their history backfill has
// completed, and the ingested date coverage.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/status
func (n *CostsNamespace) Status(ctx context.Context, params *CostsStatusParams, opts ...RequestOption) (*CostsStatusResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/status")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CostsStatusResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsUntaggedParams holds the parameters for `client.costs.untagged`.
//
// Every field is optional; pass nil to take the defaults.
type CostsUntaggedParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// From: Defaults to 30 days ago.
	From *string
	// To: Defaults to today.
	To *string
	// Basis: Which money to sum. `cash` (the default) is what the provider
	// charged on the day it charged it; `amortized` spreads a commitment's
	// up-front fee across the term it buys. Providers that report no amortized
	// amount fall back to their cash amount.
	//
	// One of "cash", "amortized".
	Basis *string
}

// Untagged: Untagged spend over the required tag keys
//
// Spend on cost rows missing at least one of the org's required tag keys,
// overall and per key, plus the largest untagged (account, service) buckets.
// Empty when no tag policy is configured — untagged is only meaningful against a
// policy.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/untagged
//
// Raises on 400: Bad request
func (n *CostsNamespace) Untagged(ctx context.Context, params *CostsUntaggedParams, opts ...RequestOption) (*UntaggedSpendReport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/untagged")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("from", params.From)
		r.addQuery("to", params.To)
		r.addQuery("basis", params.Basis)
	}
	var out *UntaggedSpendReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsAnomaliesNamespace is `client.costs.anomalies`.
type CostsAnomaliesNamespace struct {
	t *transport
}

func newCostsAnomaliesNamespace(t *transport) *CostsAnomaliesNamespace {
	n := &CostsAnomaliesNamespace{t: t}
	return n
}

// CostsAnomaliesAcknowledgeParams holds the parameters for
// `client.costs.anomalies.acknowledge`.
type CostsAnomaliesAcknowledgeParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	AnomalyID string
	// Body: the JSON request body.
	Body *CostsAnomaliesAcknowledgeRequest
}

// Acknowledge: Explain a detected cost anomaly
//
// Record what a finding actually was, and publish that sentence as a cost
// annotation on **every** chart covering the anomalous day — the point being
// that 'we migrated the fleet' is not a fact about whichever report somebody
// happened to open. The note's date (the anomalous day) and its org-wide scope
// are derived from the anomaly and are not the caller's to choose.
//
// The reply is the updated anomaly, carrying `acknowledgement` with the id of
// the note it created. Sending it again replaces the sentence and rewords that
// note rather than filing a second one; it will not recreate a note that has
// since been deleted, since deleting a note is a deliberate act and the finding
// stays explained without it.
//
// This does not suppress detection. If the same provider or service spikes again
// on a later day, that is a new anomaly and it is detected and alerted on as
// normal.
//
// POST /api/org/{orgId}/costs/anomalies/{anomalyId}/acknowledge
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *CostsAnomaliesNamespace) Acknowledge(ctx context.Context, params CostsAnomaliesAcknowledgeParams, opts ...RequestOption) (*CostAnomaly, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/costs/anomalies/{anomalyId}/acknowledge")
	r.setPath("orgId", params.OrgID)
	r.setPath("anomalyId", params.AnomalyID)
	r.setJSONBody(params.Body)
	var out *CostAnomaly
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsAnomaliesGetParams holds the parameters for `client.costs.anomalies.get`.
//
// Every field is optional; pass nil to take the defaults.
type CostsAnomaliesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Days: Window in days over anomalous days, 1-90. Defaults to 30.
	Days *string
}

// Get: List recently detected cost anomalies
//
// Spend anomalies detected by the daily background pass. Two kinds share the
// list: a `spike`, where a provider's or service's spend exceeded its trailing
// 28-day baseline by a statistical threshold (mean + N·stddev, with an absolute
// floor to ignore penny-scale noise), and a `new_source`, where a provider or
// service with no spend at all across that window suddenly billed a material
// amount. Thresholds are per organization — see GET /costs/anomaly-settings.
// Newest day first, capped at 200 rows.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/anomalies
//
// Raises on 400: Bad request
func (n *CostsAnomaliesNamespace) Get(ctx context.Context, params *CostsAnomaliesGetParams, opts ...RequestOption) (*CostsAnomaliesGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/anomalies")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("days", params.Days)
	}
	var out *CostsAnomaliesGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsAnomalySettingsNamespace is `client.costs.anomalySettings`.
type CostsAnomalySettingsNamespace struct {
	t *transport
}

func newCostsAnomalySettingsNamespace(t *transport) *CostsAnomalySettingsNamespace {
	n := &CostsAnomalySettingsNamespace{t: t}
	return n
}

// CostsAnomalySettingsGetParams holds the parameters for
// `client.costs.anomalySettings.get`.
//
// Every field is optional; pass nil to take the defaults.
type CostsAnomalySettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's anomaly detection thresholds
//
// The tunable part of cost anomaly detection. Everything else about the model —
// the 28-day baseline, the 7-day notification cooldown, the minimum history a
// baseline needs — is fixed. An organization that has never changed a threshold
// reads back the defaults. The response also carries the derived, read-only
// `smsConfigured`.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/anomaly-settings
func (n *CostsAnomalySettingsNamespace) Get(ctx context.Context, params *CostsAnomalySettingsGetParams, opts ...RequestOption) (*CostAnomalySettingsView, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/anomaly-settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CostAnomalySettingsView
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsAnomalySettingsUpdateParams holds the parameters for
// `client.costs.anomalySettings.update`.
type CostsAnomalySettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostAnomalySettings
}

// Update: Update the organization's anomaly detection thresholds
//
// Takes effect on the next detection pass (which runs after each cost
// collection). Anomalies already stored are not re-judged. All four fields are
// required — this is a PUT of the whole settings object, not a patch — and
// `smsAlerts` deliberately has no server-side default, so a client that omits it
// is rejected rather than silently switching an organization's SMS paging back
// off. `smsConfigured` is derived and is not accepted here.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/costs/anomaly-settings
//
// Raises on 400: Bad request
func (n *CostsAnomalySettingsNamespace) Update(ctx context.Context, params CostsAnomalySettingsUpdateParams, opts ...RequestOption) (*CostAnomalySettingsView, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/costs/anomaly-settings")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostAnomalySettingsView
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsEfficiencyAlertSettingsNamespace is
// `client.costs.efficiencyAlertSettings`.
type CostsEfficiencyAlertSettingsNamespace struct {
	t *transport
}

func newCostsEfficiencyAlertSettingsNamespace(t *transport) *CostsEfficiencyAlertSettingsNamespace {
	n := &CostsEfficiencyAlertSettingsNamespace{t: t}
	return n
}

// CostsEfficiencyAlertSettingsGetParams holds the parameters for
// `client.costs.efficiencyAlertSettings.get`.
//
// Every field is optional; pass nil to take the defaults.
type CostsEfficiencyAlertSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's efficiency alert tuning
//
// Thresholds for the commitment-expiry, idle-commitment and unit-cost-regression
// detectors. An organization that has never changed one reads back the defaults,
// which are chosen to work with no setup.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/costs/efficiency-alert-settings
func (n *CostsEfficiencyAlertSettingsNamespace) Get(ctx context.Context, params *CostsEfficiencyAlertSettingsGetParams, opts ...RequestOption) (*CostEfficiencySettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/costs/efficiency-alert-settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CostEfficiencySettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CostsEfficiencyAlertSettingsUpdateParams holds the parameters for
// `client.costs.efficiencyAlertSettings.update`.
type CostsEfficiencyAlertSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostEfficiencySettings
}

// Update: Update the organization's efficiency alert tuning
//
// Takes effect on the next evaluation pass (which runs after each cost
// collection). Already-fired alerts are not re-judged, and horizons that have
// already fired for a commitment's current term do not fire again — widening the
// horizon list warns about future crossings, not past ones. A PUT of the whole
// object, not a patch.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/costs/efficiency-alert-settings
//
// Raises on 400: Bad request
func (n *CostsEfficiencyAlertSettingsNamespace) Update(ctx context.Context, params CostsEfficiencyAlertSettingsUpdateParams, opts ...RequestOption) (*CostEfficiencySettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/costs/efficiency-alert-settings")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CostEfficiencySettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CredentialHygieneNamespace is `client.credentialHygiene`.
type CredentialHygieneNamespace struct {
	t *transport
}

func newCredentialHygieneNamespace(t *transport) *CredentialHygieneNamespace {
	n := &CredentialHygieneNamespace{t: t}
	return n
}

// CredentialHygieneGetParams holds the parameters for
// `client.credentialHygiene.get`.
//
// Every field is optional; pass nil to take the defaults.
type CredentialHygieneGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// WindowDays: Activity window. Defaults to 90.
	WindowDays *int64
}

// Get: Credential hygiene report
//
// API keys nobody uses, SSH keys nothing references, and members holding write
// permissions they have never exercised — derived entirely from data the server
// already holds. No provider call and nothing to enable.
//
// **The audit log only witnesses writes.** Reading a resource list or a cost
// graph leaves no audit row by design, so this report draws no conclusion about
// read permissions: an absence of evidence about them proves nothing.
// `permissionFindingsWithheld` is set when the organization does not yet have
// enough audit history for the unused-permission finding to be meaningful. Both
// are load-bearing — a governance report that overclaims is worse than none.
//
// Gated on `audit:read` rather than a permission of its own: every fact here is
// already reachable by anyone who can read the audit log, so this is a lens
// rather than a new disclosure.
//
// _Requires permission: `audit:read`._
//
// GET /api/org/{orgId}/credential-hygiene
//
// Raises on 400: Bad request
func (n *CredentialHygieneNamespace) Get(ctx context.Context, params *CredentialHygieneGetParams, opts ...RequestOption) (*HygieneReport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/credential-hygiene")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("windowDays", params.WindowDays)
	}
	var out *HygieneReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CreditsNamespace is `client.credits`.
type CreditsNamespace struct {
	t *transport
}

func newCreditsNamespace(t *transport) *CreditsNamespace {
	n := &CreditsNamespace{t: t}
	return n
}

// CreditsGetParams holds the parameters for `client.credits.get`.
//
// Every field is optional; pass nil to take the defaults.
type CreditsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Prepaid credit balances, burn rate and runway
//
// Every prepaid pot the organization holds, most urgent first. A provider that
// bills in arrears sends an invoice you can argue with; a prepaid pot that
// empties simply stops answering — so this is an availability number as much as
// a finance one.
//
// The burn rate is measured from the server's own series of readings rather than
// reported by the provider, and it is the sum of the **decreases** between
// consecutive readings: a top-up inside the window is recorded separately, never
// netted off. The runway is bounded by both the burn and the credit's own
// expiry, whichever comes first.
//
// Only providers that expose a balance appear here; most bill in arrears and
// have no pot.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/credits
func (n *CreditsNamespace) Get(ctx context.Context, params *CreditsGetParams, opts ...RequestOption) (*CreditBurndown, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/credits")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CreditBurndown
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CurrencyNamespace is `client.currency`.
type CurrencyNamespace struct {
	t *transport

	// Rates: `client.currency.rates`.
	Rates *CurrencyRatesNamespace
}

func newCurrencyNamespace(t *transport) *CurrencyNamespace {
	n := &CurrencyNamespace{t: t}
	n.Rates = newCurrencyRatesNamespace(t)
	return n
}

// CurrencyGetParams holds the parameters for `client.currency.get`.
//
// Every field is optional; pass nil to take the defaults.
type CurrencyGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: The org's display currency and exchange rate table
//
// Readable with `costs:read` rather than a settings permission: anyone who can
// see a converted total has to be able to see what it was converted at, or the
// number is unauditable.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/currency
func (n *CurrencyNamespace) Get(ctx context.Context, params *CurrencyGetParams, opts ...RequestOption) (*CurrencyConfig, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/currency")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *CurrencyConfig
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CurrencyUpdateParams holds the parameters for `client.currency.update`.
type CurrencyUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CurrencySettings
}

// Update: Set or clear the org's display currency
//
// Setting a currency opts the organization into converted totals; `null` turns
// conversion off everywhere and restores the per-currency view. Clearing does
// not delete the rate table, so conversion can be turned back on without
// re-stating anything. Only currencies with a configured rate are converted —
// Infrawrench never fetches live exchange rates.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/currency
//
// Raises on 400: Bad request
func (n *CurrencyNamespace) Update(ctx context.Context, params CurrencyUpdateParams, opts ...RequestOption) (*CurrencySettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/currency")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CurrencySettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CurrencyRatesNamespace is `client.currency.rates`.
type CurrencyRatesNamespace struct {
	t *transport
}

func newCurrencyRatesNamespace(t *transport) *CurrencyRatesNamespace {
	n := &CurrencyRatesNamespace{t: t}
	return n
}

// CurrencyRatesDeleteParams holds the parameters for
// `client.currency.rates.delete`.
type CurrencyRatesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID  *string
	RateID string
}

// Delete: Delete one exchange rate
//
// Removing a rate makes the days it covered fall back to the next-older rate, or
// to unconverted if none remains. Spend never disappears — it reverts to its own
// currency.
//
// _Requires permission: `org:settings:write`._
//
// DELETE /api/org/{orgId}/currency/rates/{rateId}
//
// Raises on 404: Not found
func (n *CurrencyRatesNamespace) Delete(ctx context.Context, params CurrencyRatesDeleteParams, opts ...RequestOption) (*CurrencyRatesDeleteResponse, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/currency/rates/{rateId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("rateId", params.RateID)
	var out *CurrencyRatesDeleteResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CurrencyRatesUpdateParams holds the parameters for
// `client.currency.rates.update`.
type CurrencyRatesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ExchangeRateInput
}

// Update: Create or replace one exchange rate
//
// Upserts on (`fromCurrency`, `toCurrency`, `effectiveFrom`) — one rate per pair
// per day, so correcting a rate replaces it rather than adding a second one
// whose precedence a reader would have to guess. Rates are stated to the display
// currency in one hop: nothing inverts a rate or chains two, because both
// produce a number you never stated.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/currency/rates
//
// Raises on 400: Bad request
func (n *CurrencyRatesNamespace) Update(ctx context.Context, params CurrencyRatesUpdateParams, opts ...RequestOption) (*ExchangeRate, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/currency/rates")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ExchangeRate
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CustomGraphsNamespace is `client.customGraphs`.
type CustomGraphsNamespace struct {
	t *transport
}

func newCustomGraphsNamespace(t *transport) *CustomGraphsNamespace {
	n := &CustomGraphsNamespace{t: t}
	return n
}

// CustomGraphsCheckParams holds the parameters for `client.customGraphs.check`.
type CustomGraphsCheckParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CustomGraphCheckRequest
}

// Check: Type-check custom-graph source without saving it
//
// _Requires permission: `dashboards:read`._
//
// POST /api/org/{orgId}/custom-graphs/check
//
// Raises on 400: Bad request
func (n *CustomGraphsNamespace) Check(ctx context.Context, params CustomGraphsCheckParams, opts ...RequestOption) (*CustomGraphCheckResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/custom-graphs/check")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CustomGraphCheckResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CustomGraphsCreateParams holds the parameters for
// `client.customGraphs.create`.
type CustomGraphsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CustomGraphInput
}

// Create: Create a custom graph (paid plan required)
//
// _Requires permission: `dashboards:write`._
//
// POST /api/org/{orgId}/custom-graphs
//
// Raises on 400: Bad request
//
// Raises on 402: Payment required — the organization's plan does not include
// this
func (n *CustomGraphsNamespace) Create(ctx context.Context, params CustomGraphsCreateParams, opts ...RequestOption) (*CustomGraphFull, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/custom-graphs")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CustomGraphFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CustomGraphsDeleteParams holds the parameters for
// `client.customGraphs.delete`.
type CustomGraphsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a custom graph (and its dashboard cards)
//
// _Requires permission: `dashboards:write`._
//
// DELETE /api/org/{orgId}/custom-graphs/{id}
//
// Raises on 404: Not found
func (n *CustomGraphsNamespace) Delete(ctx context.Context, params CustomGraphsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/custom-graphs/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CustomGraphsGetParams holds the parameters for `client.customGraphs.get`.
type CustomGraphsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a custom graph (including source)
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/custom-graphs/{id}
//
// Raises on 404: Not found
func (n *CustomGraphsNamespace) Get(ctx context.Context, params CustomGraphsGetParams, opts ...RequestOption) (*CustomGraphFull, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/custom-graphs/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *CustomGraphFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CustomGraphsListParams holds the parameters for `client.customGraphs.list`.
//
// Every field is optional; pass nil to take the defaults.
type CustomGraphsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List custom graphs
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/custom-graphs
func (n *CustomGraphsNamespace) List(ctx context.Context, params *CustomGraphsListParams, opts ...RequestOption) ([]CustomGraphSummary, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/custom-graphs")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []CustomGraphSummary
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CustomGraphsRenderParams holds the parameters for
// `client.customGraphs.render`.
type CustomGraphsRenderParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body *CustomGraphRenderRequest
}

// Render: Run the graph's script and return its render spec (paid plan required)
//
// _Requires permission: `dashboards:read`._
//
// POST /api/org/{orgId}/custom-graphs/{id}/render
//
// Raises on 400: Bad request
//
// Raises on 402: Payment required — the organization's plan does not include
// this
//
// Raises on 404: Not found
func (n *CustomGraphsNamespace) Render(ctx context.Context, params CustomGraphsRenderParams, opts ...RequestOption) (*CustomGraphRenderResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/custom-graphs/{id}/render")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CustomGraphRenderResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// CustomGraphsTypingsParams holds the parameters for
// `client.customGraphs.typings`.
//
// Every field is optional; pass nil to take the defaults.
type CustomGraphsTypingsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Typings: The ambient graph.d.ts for custom-graph source
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/custom-graphs/typings
func (n *CustomGraphsNamespace) Typings(ctx context.Context, params *CustomGraphsTypingsParams, opts ...RequestOption) (io.ReadCloser, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/custom-graphs/typings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	return n.t.stream(ctx, r, opts)
}

// CustomGraphsUpdateParams holds the parameters for
// `client.customGraphs.update`.
type CustomGraphsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body CustomGraphUpdate
}

// Update: Update a custom graph (paid plan required)
//
// _Requires permission: `dashboards:write`._
//
// PUT /api/org/{orgId}/custom-graphs/{id}
//
// Raises on 400: Bad request
//
// Raises on 402: Payment required — the organization's plan does not include
// this
//
// Raises on 404: Not found
func (n *CustomGraphsNamespace) Update(ctx context.Context, params CustomGraphsUpdateParams, opts ...RequestOption) (*CustomGraphFull, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/custom-graphs/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *CustomGraphFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsNamespace is `client.dashboards`.
type DashboardsNamespace struct {
	t *transport

	// Default: `client.dashboards.default`.
	Default *DashboardsDefaultNamespace
	// Pin: `client.dashboards.pin`.
	Pin *DashboardsPinNamespace
	// Widgets: `client.dashboards.widgets`.
	Widgets *DashboardsWidgetsNamespace
}

func newDashboardsNamespace(t *transport) *DashboardsNamespace {
	n := &DashboardsNamespace{t: t}
	n.Default = newDashboardsDefaultNamespace(t)
	n.Pin = newDashboardsPinNamespace(t)
	n.Widgets = newDashboardsWidgetsNamespace(t)
	return n
}

// DashboardsCreateParams holds the parameters for `client.dashboards.create`.
type DashboardsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body DashboardsCreateRequest
}

// Create: Create a dashboard
//
// _Requires permission: `dashboards:write`._
//
// POST /api/org/{orgId}/dashboards
func (n *DashboardsNamespace) Create(ctx context.Context, params DashboardsCreateParams, opts ...RequestOption) (*DashboardFull, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *DashboardFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsDeleteParams holds the parameters for `client.dashboards.delete`.
type DashboardsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a dashboard
//
// Cannot delete the default dashboard.
//
// _Requires permission: `dashboards:write`._
//
// DELETE /api/org/{orgId}/dashboards/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *DashboardsNamespace) Delete(ctx context.Context, params DashboardsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/dashboards/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsGetParams holds the parameters for `client.dashboards.get`.
type DashboardsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a dashboard with its pins
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/dashboards/{id}
//
// Raises on 404: Not found
func (n *DashboardsNamespace) Get(ctx context.Context, params DashboardsGetParams, opts ...RequestOption) (*DashboardWithPins, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/dashboards/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *DashboardWithPins
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsListParams holds the parameters for `client.dashboards.list`.
//
// Every field is optional; pass nil to take the defaults.
type DashboardsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List dashboards
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/dashboards
func (n *DashboardsNamespace) List(ctx context.Context, params *DashboardsListParams, opts ...RequestOption) ([]Dashboard, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/dashboards")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []Dashboard
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsProbeParams holds the parameters for `client.dashboards.probe`.
type DashboardsProbeParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ProbeRequest
}

// Probe: Read cached stats/metrics for dashboard cards
//
// _Requires permission: `dashboards:read`._
//
// POST /api/org/{orgId}/dashboards/probe
func (n *DashboardsNamespace) Probe(ctx context.Context, params DashboardsProbeParams, opts ...RequestOption) (map[string]ProbeStatus, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/probe")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out map[string]ProbeStatus
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsRenameParams holds the parameters for `client.dashboards.rename`.
type DashboardsRenameParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body DashboardsRenameRequest
}

// Rename: Rename a dashboard
//
// _Requires permission: `dashboards:write`._
//
// POST /api/org/{orgId}/dashboards/{id}/rename
func (n *DashboardsNamespace) Rename(ctx context.Context, params DashboardsRenameParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/{id}/rename")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsReorderParams holds the parameters for `client.dashboards.reorder`.
type DashboardsReorderParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body ReorderRequest
}

// Reorder: Reorder dashboard cards
//
// Persists the order of a dashboard's grid. Pass `cards` to order resource pins,
// workflow pins, and widgets as one sequence; `resourceIds` orders resource pins
// alone.
//
// _Requires permission: `dashboards:write`._
//
// POST /api/org/{orgId}/dashboards/{id}/reorder
//
// Raises on 404: Not found
func (n *DashboardsNamespace) Reorder(ctx context.Context, params DashboardsReorderParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/{id}/reorder")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsUnpinParams holds the parameters for `client.dashboards.unpin`.
type DashboardsUnpinParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body UnpinRequest
}

// Unpin: Unpin a resource
//
// _Requires permission: `dashboards:write`._
//
// POST /api/org/{orgId}/dashboards/unpin
//
// Raises on 404: Not found
func (n *DashboardsNamespace) Unpin(ctx context.Context, params DashboardsUnpinParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/unpin")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsValidateTabsParams holds the parameters for
// `client.dashboards.validateTabs`.
type DashboardsValidateTabsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ValidateTabsRequest
}

// ValidateTabs: Validate workspace tab targets still exist
//
// _Requires permission: `dashboards:read`._
//
// POST /api/org/{orgId}/dashboards/validate-tabs
func (n *DashboardsNamespace) ValidateTabs(ctx context.Context, params DashboardsValidateTabsParams, opts ...RequestOption) (*ValidateTabsResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/validate-tabs")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ValidateTabsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsWorkflowPinParams holds the parameters for
// `client.dashboards.workflowPin`.
type DashboardsWorkflowPinParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body WorkflowPinRequest
}

// WorkflowPin: Pin a workflow's metrics to a dashboard
//
// POST /api/org/{orgId}/dashboards/workflow-pin
//
// Raises on 404: Not found
func (n *DashboardsNamespace) WorkflowPin(ctx context.Context, params DashboardsWorkflowPinParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/workflow-pin")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsWorkflowUnpinParams holds the parameters for
// `client.dashboards.workflowUnpin`.
type DashboardsWorkflowUnpinParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body WorkflowPinRequest
}

// WorkflowUnpin: Unpin a workflow from a dashboard
//
// POST /api/org/{orgId}/dashboards/workflow-unpin
//
// Raises on 404: Not found
func (n *DashboardsNamespace) WorkflowUnpin(ctx context.Context, params DashboardsWorkflowUnpinParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/workflow-unpin")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsDefaultNamespace is `client.dashboards.default`.
type DashboardsDefaultNamespace struct {
	t *transport
}

func newDashboardsDefaultNamespace(t *transport) *DashboardsDefaultNamespace {
	n := &DashboardsDefaultNamespace{t: t}
	return n
}

// DashboardsDefaultFullParams holds the parameters for
// `client.dashboards.default.full`.
//
// Every field is optional; pass nil to take the defaults.
type DashboardsDefaultFullParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Full: Get-or-create the default dashboard with its pins
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/dashboards/default/full
func (n *DashboardsDefaultNamespace) Full(ctx context.Context, params *DashboardsDefaultFullParams, opts ...RequestOption) (*DashboardWithPins, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/dashboards/default/full")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *DashboardWithPins
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsPinNamespace is `client.dashboards.pin`.
type DashboardsPinNamespace struct {
	t *transport
}

func newDashboardsPinNamespace(t *transport) *DashboardsPinNamespace {
	n := &DashboardsPinNamespace{t: t}
	return n
}

// DashboardsPinCreateParams holds the parameters for
// `client.dashboards.pin.create`.
type DashboardsPinCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body PinRequest
}

// Create: Pin a resource to a dashboard
//
// _Requires permission: `dashboards:write`._
//
// POST /api/org/{orgId}/dashboards/pin
//
// Raises on 404: Not found
func (n *DashboardsPinNamespace) Create(ctx context.Context, params DashboardsPinCreateParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/pin")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsPinGetParams holds the parameters for `client.dashboards.pin.get`.
type DashboardsPinGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	PinID string
}

// Get: Full enriched pin data + cached probe status
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/dashboards/pin/{pinId}
//
// Raises on 404: Not found
func (n *DashboardsPinNamespace) Get(ctx context.Context, params DashboardsPinGetParams, opts ...RequestOption) (*PinFull, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/dashboards/pin/{pinId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("pinId", params.PinID)
	var out *PinFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsPinRangeParams holds the parameters for
// `client.dashboards.pin.range`.
type DashboardsPinRangeParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID  *string
	PinID  string
	FromMs *int64
	ToMs   *int64
}

// Range: Historical metric series for a pinned resource
//
// Returns per-series metric points between fromMs and toMs. The backend
// auto-routes between raw, 1-minute, and 1-hour rollups based on span: ≤2h raw,
// ≤7d 1m, >7d 1h.
//
// GET /api/org/{orgId}/dashboards/pin/{pinId}/range
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *DashboardsPinNamespace) Range(ctx context.Context, params DashboardsPinRangeParams, opts ...RequestOption) (*PinRangeResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/dashboards/pin/{pinId}/range")
	r.setPath("orgId", params.OrgID)
	r.setPath("pinId", params.PinID)
	r.addQuery("fromMs", params.FromMs)
	r.addQuery("toMs", params.ToMs)
	var out *PinRangeResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsWidgetsNamespace is `client.dashboards.widgets`.
type DashboardsWidgetsNamespace struct {
	t *transport
}

func newDashboardsWidgetsNamespace(t *transport) *DashboardsWidgetsNamespace {
	n := &DashboardsWidgetsNamespace{t: t}
	return n
}

// DashboardsWidgetsCreateParams holds the parameters for
// `client.dashboards.widgets.create`.
type DashboardsWidgetsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateWidgetRequest
}

// Create: Add a cost-graph or budget widget to a dashboard
//
// POST /api/org/{orgId}/dashboards/widgets
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *DashboardsWidgetsNamespace) Create(ctx context.Context, params DashboardsWidgetsCreateParams, opts ...RequestOption) (*DashboardWidgetFull, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/dashboards/widgets")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *DashboardWidgetFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsWidgetsDeleteParams holds the parameters for
// `client.dashboards.widgets.delete`.
type DashboardsWidgetsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	WidgetID string
}

// Delete: Remove a widget from a dashboard
//
// DELETE /api/org/{orgId}/dashboards/widgets/{widgetId}
//
// Raises on 404: Not found
func (n *DashboardsWidgetsNamespace) Delete(ctx context.Context, params DashboardsWidgetsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/dashboards/widgets/{widgetId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("widgetId", params.WidgetID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DashboardsWidgetsUpdateParams holds the parameters for
// `client.dashboards.widgets.update`.
type DashboardsWidgetsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	WidgetID string
	// Body: the JSON request body.
	Body UpdateWidgetRequest
}

// Update: Update a widget's title, config, or layout
//
// PATCH /api/org/{orgId}/dashboards/widgets/{widgetId}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *DashboardsWidgetsNamespace) Update(ctx context.Context, params DashboardsWidgetsUpdateParams, opts ...RequestOption) (*DashboardWidgetFull, error) {
	r := newRequest(http.MethodPatch, "/api/org/{orgId}/dashboards/widgets/{widgetId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("widgetId", params.WidgetID)
	r.setJSONBody(params.Body)
	var out *DashboardWidgetFull
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DependencyGraphNamespace is `client.dependencyGraph`.
type DependencyGraphNamespace struct {
	t *transport
}

func newDependencyGraphNamespace(t *transport) *DependencyGraphNamespace {
	n := &DependencyGraphNamespace{t: t}
	return n
}

// DependencyGraphGetParams holds the parameters for
// `client.dependencyGraph.get`.
//
// Every field is optional; pass nil to take the defaults.
type DependencyGraphGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ResourceID *ResourceID
}

// Get: The org's resource dependency graph, from synced cloud data and output
// references
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/dependency-graph
func (n *DependencyGraphNamespace) Get(ctx context.Context, params *DependencyGraphGetParams, opts ...RequestOption) (*DependencyGraphResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/dependency-graph")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("resourceId", params.ResourceID)
	}
	var out *DependencyGraphResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsNamespace is `client.deployments`.
type DeploymentsNamespace struct {
	t *transport

	// Runs: `client.deployments.runs`.
	Runs *DeploymentsRunsNamespace
	// Triggers: `client.deployments.triggers`.
	Triggers *DeploymentsTriggersNamespace
}

func newDeploymentsNamespace(t *transport) *DeploymentsNamespace {
	n := &DeploymentsNamespace{t: t}
	n.Runs = newDeploymentsRunsNamespace(t)
	n.Triggers = newDeploymentsTriggersNamespace(t)
	return n
}

// DeploymentsEnvsParams holds the parameters for `client.deployments.envs`.
//
// Every field is optional; pass nil to take the defaults.
type DeploymentsEnvsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *DeployEnvsInput
}

// Envs: List the environments a repository's Infrafile declares
//
// Reads `Infrafile` at the branch head and returns its declared environments.
// The file is parsed, not executed.
//
// _Requires permission: `deployments:read`._
//
// POST /api/org/{orgId}/deployments/envs
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *DeploymentsNamespace) Envs(ctx context.Context, params *DeploymentsEnvsParams, opts ...RequestOption) (*DeployEnvs, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/deployments/envs")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *DeployEnvs
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsPlanParams holds the parameters for `client.deployments.plan`.
//
// Every field is optional; pass nil to take the defaults.
type DeploymentsPlanParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *DeployPlanInput
}

// Plan: Preview a deploy without building
//
// Runs the Infrafile's `plan()` and renders its Dockerfile, then stops. Nothing
// is built or deployed.
//
// _Requires permission: `deployments:plan`._
//
// POST /api/org/{orgId}/deployments/plan
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *DeploymentsNamespace) Plan(ctx context.Context, params *DeploymentsPlanParams, opts ...RequestOption) (*DeployPlanResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/deployments/plan")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *DeployPlanResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsReposParams holds the parameters for `client.deployments.repos`.
//
// Every field is optional; pass nil to take the defaults.
type DeploymentsReposParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Repos: List repositories this organization can deploy from
//
// Repositories visible to the organization's GitHub App installations. Empty
// when the app is not configured.
//
// _Requires permission: `deployments:read`._
//
// GET /api/org/{orgId}/deployments/repos
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
func (n *DeploymentsNamespace) Repos(ctx context.Context, params *DeploymentsReposParams, opts ...RequestOption) ([]DeployRepo, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/deployments/repos")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []DeployRepo
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsRunsNamespace is `client.deployments.runs`.
type DeploymentsRunsNamespace struct {
	t *transport
}

func newDeploymentsRunsNamespace(t *transport) *DeploymentsRunsNamespace {
	n := &DeploymentsRunsNamespace{t: t}
	return n
}

// DeploymentsRunsCreateParams holds the parameters for
// `client.deployments.runs.create`.
//
// Every field is optional; pass nil to take the defaults.
type DeploymentsRunsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *DeploymentRunInput
}

// Create: Record a deployment that ran elsewhere
//
// The CLI builds on the operator's own machine, so the server never sees that
// run. Reporting it here keeps one history across both origins.
//
// _Requires permission: `deployments:write`._
//
// POST /api/org/{orgId}/deployments/runs
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
func (n *DeploymentsRunsNamespace) Create(ctx context.Context, params *DeploymentsRunsCreateParams, opts ...RequestOption) (*DeploymentsRunsCreateResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/deployments/runs")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *DeploymentsRunsCreateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsRunsGetParams holds the parameters for
// `client.deployments.runs.get`.
type DeploymentsRunsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get one deployment run, with its logs and rendered Dockerfile
//
// _Requires permission: `deployments:read`._
//
// GET /api/org/{orgId}/deployments/runs/{id}
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *DeploymentsRunsNamespace) Get(ctx context.Context, params DeploymentsRunsGetParams, opts ...RequestOption) (*DeploymentRun, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/deployments/runs/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *DeploymentRun
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsRunsListParams holds the parameters for
// `client.deployments.runs.list`.
//
// Every field is optional; pass nil to take the defaults.
type DeploymentsRunsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	Env   *string
	Limit *int64
}

// List: List deployment runs
//
// _Requires permission: `deployments:read`._
//
// GET /api/org/{orgId}/deployments/runs
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
func (n *DeploymentsRunsNamespace) List(ctx context.Context, params *DeploymentsRunsListParams, opts ...RequestOption) ([]DeploymentRun, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/deployments/runs")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("env", params.Env)
		r.addQuery("limit", params.Limit)
	}
	var out []DeploymentRun
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsRunsRollbackParams holds the parameters for
// `client.deployments.runs.rollback`.
type DeploymentsRunsRollbackParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body *DeployRollbackInput
}

// Rollback: Roll back to a previous deployment
//
// Re-runs that run's `deploy()` with the image and plan it recorded, building
// nothing — the exact artifact that was known good ships again. The Infrafile is
// read at the commit that run deployed, not at the branch head. Only a
// successful run that produced an image can be rolled back to. With
// `deleteCreated`, resources that runs after the target created through
// `infra.accounts` are deleted once the rollback has succeeded — undoing the
// provisioning, not just the shipping. Deletions are best-effort and reported in
// the result's notes.
//
// _Requires permission: `deployments:write`._
//
// POST /api/org/{orgId}/deployments/runs/{id}/rollback
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 402: Payment required — the organization's plan does not include
// this
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
//
// Raises on 423: Blocked by an active change freeze. Retry with the
// `x-change-freeze-override: true` header if you hold `freezes:override`; both
// blocks and overrides are audit-logged.
func (n *DeploymentsRunsNamespace) Rollback(ctx context.Context, params DeploymentsRunsRollbackParams, opts ...RequestOption) (*DeployPlanResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/deployments/runs/{id}/rollback")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *DeployPlanResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsTriggersNamespace is `client.deployments.triggers`.
type DeploymentsTriggersNamespace struct {
	t *transport
}

func newDeploymentsTriggersNamespace(t *transport) *DeploymentsTriggersNamespace {
	n := &DeploymentsTriggersNamespace{t: t}
	return n
}

// DeploymentsTriggersCreateParams holds the parameters for
// `client.deployments.triggers.create`.
//
// Every field is optional; pass nil to take the defaults.
type DeploymentsTriggersCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *DeployTriggerInput
}

// Create: Deploy an environment whenever a branch moves
//
// Arming a trigger records the branch's current commit WITHOUT deploying it —
// the trigger fires on the next push, not on the state at the moment it was
// created. The environment is validated against the Infrafile at that branch
// head, so a typo fails here rather than silently never firing.
//
// _Requires permission: `deployments:write`._
//
// POST /api/org/{orgId}/deployments/triggers
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *DeploymentsTriggersNamespace) Create(ctx context.Context, params *DeploymentsTriggersCreateParams, opts ...RequestOption) (*DeployTrigger, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/deployments/triggers")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *DeployTrigger
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsTriggersDeleteParams holds the parameters for
// `client.deployments.triggers.delete`.
type DeploymentsTriggersDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a deploy trigger
//
// _Requires permission: `deployments:write`._
//
// DELETE /api/org/{orgId}/deployments/triggers/{id}
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
func (n *DeploymentsTriggersNamespace) Delete(ctx context.Context, params DeploymentsTriggersDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/deployments/triggers/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsTriggersListParams holds the parameters for
// `client.deployments.triggers.list`.
//
// Every field is optional; pass nil to take the defaults.
type DeploymentsTriggersListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List deploy-on-push triggers
//
// _Requires permission: `deployments:read`._
//
// GET /api/org/{orgId}/deployments/triggers
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
func (n *DeploymentsTriggersNamespace) List(ctx context.Context, params *DeploymentsTriggersListParams, opts ...RequestOption) ([]DeployTrigger, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/deployments/triggers")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []DeployTrigger
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeploymentsTriggersUpdateParams holds the parameters for
// `client.deployments.triggers.update`.
type DeploymentsTriggersUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body *DeploymentsTriggersUpdateRequest
}

// Update: Enable or disable a deploy trigger
//
// _Requires permission: `deployments:write`._
//
// PATCH /api/org/{orgId}/deployments/triggers/{id}
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *DeploymentsTriggersNamespace) Update(ctx context.Context, params DeploymentsTriggersUpdateParams, opts ...RequestOption) (*DeployTrigger, error) {
	r := newRequest(http.MethodPatch, "/api/org/{orgId}/deployments/triggers/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *DeployTrigger
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DigestNamespace is `client.digest`.
type DigestNamespace struct {
	t *transport

	// Recipients: `client.digest.recipients`.
	Recipients *DigestRecipientsNamespace
}

func newDigestNamespace(t *transport) *DigestNamespace {
	n := &DigestNamespace{t: t}
	n.Recipients = newDigestRecipientsNamespace(t)
	return n
}

// DigestGetParams holds the parameters for `client.digest.get`.
//
// Every field is optional; pass nil to take the defaults.
type DigestGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's weekly digest settings
//
// The weekly digest is a summary of the last complete Monday-to-Sunday week's
// spend (with week-over-week movers), sync incidents, and resource churn,
// delivered to the Slack channels and Teams webhooks opted into the weeklyDigest
// trigger and to the organization's digest email recipients. The response also
// carries the outcome of the most recent delivery attempt so a silently failing
// digest is visible.
//
// GET /api/org/{orgId}/digest
func (n *DigestNamespace) Get(ctx context.Context, params *DigestGetParams, opts ...RequestOption) (*DigestSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/digest")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *DigestSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DigestSendParams holds the parameters for `client.digest.send`.
//
// Every field is optional; pass nil to take the defaults.
type DigestSendParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Send: Compose and send last week's digest now
//
// Ignores the schedule and the enabled flag — composes the digest for the last
// complete week and sends it to every opted-in channel and email recipient. This
// is also the manual recovery for a partial delivery, which is never retried
// automatically. Fails when nothing is routed to receive the digest, or when
// every destination rejected it.
//
// POST /api/org/{orgId}/digest/send
//
// Raises on 400: Bad request
func (n *DigestNamespace) Send(ctx context.Context, params *DigestSendParams, opts ...RequestOption) (*DigestSendResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/digest/send")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *DigestSendResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DigestUpdateParams holds the parameters for `client.digest.update`.
//
// Every field is optional; pass nil to take the defaults.
type DigestUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *DigestSettingsUpdate
}

// Update: Update the weekly digest settings
//
// Every field is optional. Enabling schedules the first digest for the next
// configured send time rather than sending immediately — use POST /digest/send
// for an immediate one. The week boundary follows `timezone`, so the reported
// window is always the organization's own local Monday-to-Sunday week. Changing
// the schedule clears any parked failure state but never replays a week that
// already went out.
//
// PUT /api/org/{orgId}/digest
//
// Raises on 400: Bad request
func (n *DigestNamespace) Update(ctx context.Context, params *DigestUpdateParams, opts ...RequestOption) (*DigestSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/digest")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *DigestSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DigestRecipientsNamespace is `client.digest.recipients`.
type DigestRecipientsNamespace struct {
	t *transport
}

func newDigestRecipientsNamespace(t *transport) *DigestRecipientsNamespace {
	n := &DigestRecipientsNamespace{t: t}
	return n
}

// DigestRecipientsCreateParams holds the parameters for
// `client.digest.recipients.create`.
//
// Every field is optional; pass nil to take the defaults.
type DigestRecipientsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *DigestEmailRecipientCreate
}

// Create: Add a digest email recipient
//
// Adding an address the organization already has is a no-op that returns the
// existing entry, so a double submit cannot double-deliver.
//
// POST /api/org/{orgId}/digest/recipients
//
// Raises on 400: Bad request
func (n *DigestRecipientsNamespace) Create(ctx context.Context, params *DigestRecipientsCreateParams, opts ...RequestOption) (*DigestEmailRecipient, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/digest/recipients")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *DigestEmailRecipient
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DigestRecipientsDeleteParams holds the parameters for
// `client.digest.recipients.delete`.
type DigestRecipientsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// RecipientID: Recipient id
	RecipientID string
}

// Delete: Remove a digest email recipient
//
// DELETE /api/org/{orgId}/digest/recipients/{recipientId}
//
// Raises on 404: Not found
func (n *DigestRecipientsNamespace) Delete(ctx context.Context, params DigestRecipientsDeleteParams, opts ...RequestOption) (*DigestRecipientsDeleteResponse, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/digest/recipients/{recipientId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("recipientId", params.RecipientID)
	var out *DigestRecipientsDeleteResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DigestRecipientsGetParams holds the parameters for
// `client.digest.recipients.get`.
//
// Every field is optional; pass nil to take the defaults.
type DigestRecipientsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List the organization's digest email recipients
//
// Email is a digest-only transport, so its destinations are an
// organization-level address list rather than a per-channel trigger. Addresses
// need not belong to Infrawrench users — a finance alias is a valid recipient.
//
// GET /api/org/{orgId}/digest/recipients
func (n *DigestRecipientsNamespace) Get(ctx context.Context, params *DigestRecipientsGetParams, opts ...RequestOption) (*DigestEmailRecipientList, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/digest/recipients")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *DigestEmailRecipientList
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DNSNamespace is `client.dns`.
type DNSNamespace struct {
	t *transport
}

func newDNSNamespace(t *transport) *DNSNamespace {
	n := &DNSNamespace{t: t}
	return n
}

// DNSGetParams holds the parameters for `client.dns.get`.
//
// Every field is optional; pass nil to take the defaults.
type DNSGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List every DNS zone and record, with dangling targets flagged
//
// One view over every zone and record across the connected DNS providers
// (Cloudflare, Route 53, Cloud DNS, DigitalOcean, Netlify, Azure DNS, Vercel),
// with each record target classified against the rest of the workspace. No
// provider API calls are made and no DNS is resolved — results reflect the last
// sync.
//
// A `dangling` target is a subdomain-takeover candidate: the record points into
// a provider namespace this workspace manages and nothing synced claims it. The
// same records surface as `dns-dangling-target` findings on `GET /posture` and
// alert through the posture channel, so there is no separate DNS alert setting.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/dns
func (n *DNSNamespace) Get(ctx context.Context, params *DNSGetParams, opts ...RequestOption) (*DNSInventoryResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/dns")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *DNSInventoryResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DockerNamespace is `client.docker`.
type DockerNamespace struct {
	t *transport
}

func newDockerNamespace(t *transport) *DockerNamespace {
	n := &DockerNamespace{t: t}
	return n
}

// DockerCommandParams holds the parameters for `client.docker.command`.
type DockerCommandParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body DockerCommandRequest
}

// Command: Run a Docker daemon operation
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/docker/command
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *DockerNamespace) Command(ctx context.Context, params DockerCommandParams, opts ...RequestOption) (*DockerCommandResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/docker/command")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *DockerCommandResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// EnvironmentDiffNamespace is `client.environmentDiff`.
type EnvironmentDiffNamespace struct {
	t *transport
}

func newEnvironmentDiffNamespace(t *transport) *EnvironmentDiffNamespace {
	n := &EnvironmentDiffNamespace{t: t}
	return n
}

// EnvironmentDiffGetParams holds the parameters for
// `client.environmentDiff.get`.
type EnvironmentDiffGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// A: Baseline account id — by convention the environment that works.
	A string
	// B: Compared account id. Must differ from `a` and use the same provider.
	B string
	// ResourceTypeID: Compare one resource type only.
	ResourceTypeID *string
	// IncludeIdentityFields: Compare identity and timestamp fields too, instead
	// of filtering them out.
	//
	// One of "true", "false".
	IncludeIdentityFields *string
}

// Get: Compare two accounts' resource inventories
//
// Compares two accounts of the same provider — typically staging against
// production — over already-synced state: which resource types exist in one and
// not the other, the per-type count deltas, and the fields on which two
// corresponding resources disagree (instance class, engine version, feature
// flags).
//
// Resources are paired by resource type plus name with environment words
// removed, so `api-staging` lines up with `api-prod` without any naming
// convention to configure. By default the comparison hides divergences that are
// artefacts of being two different resources — ids, links, network addresses and
// timestamps — because every resource has different ones; pass
// `includeIdentityFields=true` to see them.
//
// Read-only and cheap: no provider API calls are made, so results reflect the
// last sync.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/environment-diff
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *EnvironmentDiffNamespace) Get(ctx context.Context, params EnvironmentDiffGetParams, opts ...RequestOption) (*EnvironmentDiffResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/environment-diff")
	r.setPath("orgId", params.OrgID)
	r.addQuery("a", params.A)
	r.addQuery("b", params.B)
	r.addQuery("resourceTypeId", params.ResourceTypeID)
	r.addQuery("includeIdentityFields", params.IncludeIdentityFields)
	var out *EnvironmentDiffResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ExpiringNamespace is `client.expiring`.
type ExpiringNamespace struct {
	t *transport

	// Settings: `client.expiring.settings`.
	Settings *ExpiringSettingsNamespace
}

func newExpiringNamespace(t *transport) *ExpiringNamespace {
	n := &ExpiringNamespace{t: t}
	n.Settings = newExpiringSettingsNamespace(t)
	return n
}

// ExpiringGetParams holds the parameters for `client.expiring.get`.
//
// Every field is optional; pass nil to take the defaults.
type ExpiringGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List approaching deadlines on synced resources
//
// One cross-provider countdown of everything with a clock on it: TLS certificate
// expiries, domain registrations, API token expirations, access keys past their
// rotation budget, Kubernetes/SSH credential ages. Plugins declare which synced
// fields carry deadlines; the feed is computed over already-stored state, so no
// provider API calls are made and results reflect the last sync. Items are
// sorted soonest first and bucketed by severity against the organization's lead
// time.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/expiring
func (n *ExpiringNamespace) Get(ctx context.Context, params *ExpiringGetParams, opts ...RequestOption) (*ExpiryListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/expiring")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *ExpiryListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ExpiringSettingsNamespace is `client.expiring.settings`.
type ExpiringSettingsNamespace struct {
	t *transport
}

func newExpiringSettingsNamespace(t *transport) *ExpiringSettingsNamespace {
	n := &ExpiringSettingsNamespace{t: t}
	return n
}

// ExpiringSettingsGetParams holds the parameters for
// `client.expiring.settings.get`.
//
// Every field is optional; pass nil to take the defaults.
type ExpiringSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's expiry alert settings
//
// The lead time feeds both the feed's `upcoming` bucket and the poller's daily
// alert scan. An organization that never saved reads the shipped defaults
// (enabled, 60 days).
//
// _Requires permission: `org:settings:write`._
//
// GET /api/org/{orgId}/expiring/settings
func (n *ExpiringSettingsNamespace) Get(ctx context.Context, params *ExpiringSettingsGetParams, opts ...RequestOption) (*ExpiryAlertSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/expiring/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *ExpiryAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ExpiringSettingsUpdateParams holds the parameters for
// `client.expiring.settings.update`.
//
// Every field is optional; pass nil to take the defaults.
type ExpiringSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *ExpiryAlertSettingsUpdate
}

// Update: Update the expiry alert settings
//
// Every field is optional so a single toggle can be saved on its own. `leadDays`
// must be a whole number from 1 to 365. Saving never resets the alert cooldown.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/expiring/settings
//
// Raises on 400: Bad request
func (n *ExpiringSettingsNamespace) Update(ctx context.Context, params *ExpiringSettingsUpdateParams, opts ...RequestOption) (*ExpiryAlertSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/expiring/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *ExpiryAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvitationsNamespace is `client.invitations`.
type InvitationsNamespace struct {
	t *transport

	// ByToken: `client.invitations.byToken`.
	ByToken *InvitationsByTokenNamespace
}

func newInvitationsNamespace(t *transport) *InvitationsNamespace {
	n := &InvitationsNamespace{t: t}
	n.ByToken = newInvitationsByTokenNamespace(t)
	return n
}

// InvitationsAcceptParams holds the parameters for `client.invitations.accept`.
type InvitationsAcceptParams struct {
	// Body: the JSON request body.
	Body AcceptInvitationRequest
}

// Accept: Accept an invitation
//
// POST /api/invitations/accept
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
func (n *InvitationsNamespace) Accept(ctx context.Context, params InvitationsAcceptParams, opts ...RequestOption) (*AcceptInvitationResponse, error) {
	r := newRequest(http.MethodPost, "/api/invitations/accept")
	r.setJSONBody(params.Body)
	var out *AcceptInvitationResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvitationsByTokenNamespace is `client.invitations.byToken`.
type InvitationsByTokenNamespace struct {
	t *transport
}

func newInvitationsByTokenNamespace(t *transport) *InvitationsByTokenNamespace {
	n := &InvitationsByTokenNamespace{t: t}
	return n
}

// InvitationsByTokenGetParams holds the parameters for
// `client.invitations.byToken.get`.
type InvitationsByTokenGetParams struct {
	Token string
}

// Get: Get invitation details by token
//
// GET /api/invitations/by-token/{token}
//
// Raises on 404: Not found
func (n *InvitationsByTokenNamespace) Get(ctx context.Context, params InvitationsByTokenGetParams, opts ...RequestOption) (*InvitationDetail, error) {
	r := newRequest(http.MethodGet, "/api/invitations/by-token/{token}")
	r.setPath("token", params.Token)
	var out *InvitationDetail
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesNamespace is `client.invoices`.
type InvoicesNamespace struct {
	t *transport
}

func newInvoicesNamespace(t *transport) *InvoicesNamespace {
	n := &InvoicesNamespace{t: t}
	return n
}

// InvoicesApproveParams holds the parameters for `client.invoices.approve`.
type InvoicesApproveParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Approve: Approve an invoice — freeze its figures
//
// Computes the figures one last time and writes them onto the invoice together
// with the exchange rates, the day they were read, the billing rules in force
// and the names everything in scope had. From here the invoice is a document,
// not a query.
//
// A distinct act from generation, on a distinct permission (`invoices:issue`),
// with its own audit entry recording who approved what.
//
// Refused with 409 when a currency in the invoice has no stated exchange rate:
// an approved invoice has to be quotable as one number in the customer's
// currency.
//
// Refused with 409, too, when the draft or its customer changed while the
// figures were being computed — a different period, scope, currency, cost basis
// or billing-rules setting. Nothing is approved in that case: freezing figures
// that describe a different question would be worse than making the caller look
// again.
//
// _Requires permission: `invoices:issue`._
//
// POST /api/org/{orgId}/invoices/{id}/approve
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *InvoicesNamespace) Approve(ctx context.Context, params InvoicesApproveParams, opts ...RequestOption) (*Invoice, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/invoices/{id}/approve")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *Invoice
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesCreateParams holds the parameters for `client.invoices.create`.
type InvoicesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body InvoiceInput
}

// Create: Raise a draft invoice
//
// Always lands in `draft`. Generating and issuing are two acts on two
// permissions: a mistyped period must not be able to reach a customer without
// anyone having read the numbers.
//
// _Requires permission: `invoices:write`._
//
// POST /api/org/{orgId}/invoices
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *InvoicesNamespace) Create(ctx context.Context, params InvoicesCreateParams, opts ...RequestOption) (*Invoice, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/invoices")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *Invoice
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesDeleteParams holds the parameters for `client.invoices.delete`.
type InvoicesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a draft invoice
//
// Draft only, and refused with 409 otherwise. An issued invoice is voided;
// deleting one would erase a document a customer holds a copy of.
//
// _Requires permission: `invoices:write`._
//
// DELETE /api/org/{orgId}/invoices/{id}
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *InvoicesNamespace) Delete(ctx context.Context, params InvoicesDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/invoices/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesExportParams holds the parameters for `client.invoices.export`.
type InvoicesExportParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Export: Download an invoice as CSV
//
// The derivation, not a rendered document: what was collected, what the rules
// added, the rate and the day it was read, and the final figure — every column
// an accounts-payable clerk needs to check the arithmetic. Same RFC 4180 quoting
// as the scheduled cost exports.
//
// _Requires permission: `invoices:read`._
//
// GET /api/org/{orgId}/invoices/{id}/export
//
// Raises on 404: Not found
func (n *InvoicesNamespace) Export(ctx context.Context, params InvoicesExportParams, opts ...RequestOption) (io.ReadCloser, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/invoices/{id}/export")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	return n.t.stream(ctx, r, opts)
}

// InvoicesGetParams holds the parameters for `client.invoices.get`.
type InvoicesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get an invoice
//
// **A draft recomputes from live spend; an approved, sent or void invoice does
// not.** `live` says which happened. A frozen invoice returns the figures
// written at approval and does not read cost data at all.
//
// _Requires permission: `invoices:read`._
//
// GET /api/org/{orgId}/invoices/{id}
//
// Raises on 404: Not found
func (n *InvoicesNamespace) Get(ctx context.Context, params InvoicesGetParams, opts ...RequestOption) (*Invoice, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/invoices/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *Invoice
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesListParams holds the parameters for `client.invoices.list`.
//
// Every field is optional; pass nil to take the defaults.
type InvoicesListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID            *string
	ManagedAccountID *string
	Status           *InvoiceStatus
}

// List: List invoices
//
// Summaries, newest period first. A draft's `totals` is null here rather than
// recomputed — recomputing every draft would make opening the list one cost-data
// scan per draft, and zero would be a lie the reader cannot detect.
//
// _Requires permission: `invoices:read`._
//
// GET /api/org/{orgId}/invoices
//
// Raises on 400: Bad request
func (n *InvoicesNamespace) List(ctx context.Context, params *InvoicesListParams, opts ...RequestOption) ([]InvoiceSummary, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/invoices")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("managedAccountId", params.ManagedAccountID)
		r.addQuery("status", params.Status)
	}
	var out []InvoiceSummary
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesSendParams holds the parameters for `client.invoices.send`.
type InvoicesSendParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body *InvoiceSendRequest
}

// Send: Send an invoice to its customer
//
// Changes no figure — the document was frozen at approval. It records the
// **release** (this may go to the customer, and this person said so), then
// emails the invoice to the customer's contact addresses with the CSV attached.
//
// **200 even when delivery failed.** The release happened and is recorded either
// way; `delivery` says what became of the transport. An error status would leave
// the caller unable to tell which of the two failed. A failed delivery is
// visible, and re-sending retries it.
//
// Sending again needs `resend: true` only when the last attempt reached somebody
// — see `InvoiceSendRequest`. The body may be omitted entirely for a first send.
//
// _Requires permission: `invoices:issue`._
//
// POST /api/org/{orgId}/invoices/{id}/send
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *InvoicesNamespace) Send(ctx context.Context, params InvoicesSendParams, opts ...RequestOption) (*Invoice, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/invoices/{id}/send")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *Invoice
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesUpdateParams holds the parameters for `client.invoices.update`.
type InvoicesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body InvoiceUpdate
}

// Update: Edit a draft invoice
//
// Draft only. An approved, sent or void invoice is refused with 409 by the
// service, not merely hidden by the UI — an issued invoice that silently changed
// after the customer received it is the worst outcome this feature could
// produce.
//
// _Requires permission: `invoices:write`._
//
// PUT /api/org/{orgId}/invoices/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *InvoicesNamespace) Update(ctx context.Context, params InvoicesUpdateParams, opts ...RequestOption) (*Invoice, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/invoices/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *Invoice
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// InvoicesVoidParams holds the parameters for `client.invoices.void`.
type InvoicesVoidParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body InvoiceVoidRequest
}

// Void: Void an issued invoice
//
// The only correction there is. The original keeps every figure it was sent with
// — “we billed you this, it was wrong, here is the corrected one” is a story a
// customer can follow, and “we changed the invoice” is not.
//
// With `supersede`, the void, the corrective draft and both directions of the
// link between them are one transaction. Void is irreversible, so a half-applied
// correction would leave a withdrawn invoice with no way forward; this call
// either applies whole or not at all.
//
// _Requires permission: `invoices:issue`._
//
// POST /api/org/{orgId}/invoices/{id}/void
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *InvoicesNamespace) Void(ctx context.Context, params InvoicesVoidParams, opts ...RequestOption) (*InvoiceVoidResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/invoices/{id}/void")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *InvoiceVoidResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraNamespace is `client.jira`.
type JiraNamespace struct {
	t *transport

	// Projects: `client.jira.projects`.
	Projects *JiraProjectsNamespace
}

func newJiraNamespace(t *transport) *JiraNamespace {
	n := &JiraNamespace{t: t}
	n.Projects = newJiraProjectsNamespace(t)
	return n
}

// JiraDeleteParams holds the parameters for `client.jira.delete`.
//
// Every field is optional; pass nil to take the defaults.
type JiraDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Delete: Disconnect Jira
//
// Issue links already recorded are kept, so filed findings stay marked as filed.
//
// _Requires permission: `jira:write`._
//
// DELETE /api/org/{orgId}/jira
//
// Raises on 404: Not found
func (n *JiraNamespace) Delete(ctx context.Context, params *JiraDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/jira")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraGetParams holds the parameters for `client.jira.get`.
//
// Every field is optional; pass nil to take the defaults.
type JiraGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the org's Jira connection
//
// The stored API token is never returned; `tokenHint` stands in for it.
//
// _Requires permission: `jira:read`._
//
// GET /api/org/{orgId}/jira
func (n *JiraNamespace) Get(ctx context.Context, params *JiraGetParams, opts ...RequestOption) (*JiraGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/jira")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *JiraGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraIssuesParams holds the parameters for `client.jira.issues`.
type JiraIssuesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateJiraIssueInput
}

// Issues: File a finding as a Jira issue
//
// Creates the issue, then records the link between it and the finding. The link
// is what lets a list view show "already filed" instead of offering the button
// again.
//
// _Requires permission: `jira:write`._
//
// POST /api/org/{orgId}/jira/issues
//
// Raises on 400: Bad request
//
// Raises on 502: Jira refused to create the issue, or was unreachable
func (n *JiraNamespace) Issues(ctx context.Context, params JiraIssuesParams, opts ...RequestOption) (*CreateJiraIssueResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/jira/issues")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CreateJiraIssueResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraLinksParams holds the parameters for `client.jira.links`.
//
// Every field is optional; pass nil to take the defaults.
type JiraLinksParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	SourceKind *JiraSourceKind
	// SourceID: Repeat to narrow to specific findings. Omit to return every link
	// of the kind — this is the batch lookup a list view makes once instead of
	// one request per row.
	SourceID []string
}

// Links: Look up filed issues for a set of findings
//
// _Requires permission: `jira:read`._
//
// GET /api/org/{orgId}/jira/links
//
// Raises on 400: Bad request
func (n *JiraNamespace) Links(ctx context.Context, params *JiraLinksParams, opts ...RequestOption) ([]JiraIssueLink, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/jira/links")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("sourceKind", params.SourceKind)
		r.addQuery("sourceId", params.SourceID)
	}
	var out []JiraIssueLink
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraUpdateParams holds the parameters for `client.jira.update`.
type JiraUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body JiraIntegrationInput
}

// Update: Connect Jira, or update the connection
//
// _Requires permission: `jira:write`._
//
// PUT /api/org/{orgId}/jira
//
// Raises on 400: Bad request
func (n *JiraNamespace) Update(ctx context.Context, params JiraUpdateParams, opts ...RequestOption) (*JiraIntegration, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/jira")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *JiraIntegration
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraVerifyParams holds the parameters for `client.jira.verify`.
//
// Every field is optional; pass nil to take the defaults.
type JiraVerifyParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *JiraVerifyInput
}

// Verify: Check Jira credentials
//
// Calls GET /rest/api/3/myself on the site, so a wrong email or a revoked token
// is reported on the settings form rather than on the first attempt to file an
// issue.
//
// _Requires permission: `jira:write`._
//
// POST /api/org/{orgId}/jira/verify
//
// Raises on 400: Bad request
//
// Raises on 502: Jira rejected the credentials or was unreachable
func (n *JiraNamespace) Verify(ctx context.Context, params *JiraVerifyParams, opts ...RequestOption) (*JiraVerifyResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/jira/verify")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *JiraVerifyResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraProjectsNamespace is `client.jira.projects`.
type JiraProjectsNamespace struct {
	t *transport
}

func newJiraProjectsNamespace(t *transport) *JiraProjectsNamespace {
	n := &JiraProjectsNamespace{t: t}
	return n
}

// JiraProjectsIssueTypesParams holds the parameters for
// `client.jira.projects.issueTypes`.
type JiraProjectsIssueTypesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	Key   string
}

// IssueTypes: List issue types valid in a project
//
// Reads the project's own create metadata rather than the global issue-type
// list, so the picker cannot offer a type the project's scheme would reject.
// Subtasks are excluded.
//
// _Requires permission: `jira:read`._
//
// GET /api/org/{orgId}/jira/projects/{key}/issue-types
//
// Raises on 400: Bad request
func (n *JiraProjectsNamespace) IssueTypes(ctx context.Context, params JiraProjectsIssueTypesParams, opts ...RequestOption) ([]JiraIssueType, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/jira/projects/{key}/issue-types")
	r.setPath("orgId", params.OrgID)
	r.setPath("key", params.Key)
	var out []JiraIssueType
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// JiraProjectsListParams holds the parameters for `client.jira.projects.list`.
//
// Every field is optional; pass nil to take the defaults.
type JiraProjectsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List Jira projects
//
// Backs the project picker, so nobody has to know a project key by hand.
//
// _Requires permission: `jira:read`._
//
// GET /api/org/{orgId}/jira/projects
//
// Raises on 400: Bad request
func (n *JiraProjectsNamespace) List(ctx context.Context, params *JiraProjectsListParams, opts ...RequestOption) ([]JiraProject, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/jira/projects")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []JiraProject
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// KVNamespace is `client.kv`.
type KVNamespace struct {
	t *transport
}

func newKVNamespace(t *transport) *KVNamespace {
	n := &KVNamespace{t: t}
	return n
}

// KVCommandParams holds the parameters for `client.kv.command`.
type KVCommandParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body KVCommandRequest
}

// Command: Run a Redis-style KV command
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/kv/command
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *KVNamespace) Command(ctx context.Context, params KVCommandParams, opts ...RequestOption) (*KVCommandResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/kv/command")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *KVCommandResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LeasesNamespace is `client.leases`.
type LeasesNamespace struct {
	t *transport
}

func newLeasesNamespace(t *transport) *LeasesNamespace {
	n := &LeasesNamespace{t: t}
	return n
}

// LeasesCancelParams holds the parameters for `client.leases.cancel`.
type LeasesCancelParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	LeaseID string
}

// Cancel: Cancel a lease
//
// Stop the countdown — the resource stays, the lease goes `canceled` and leaves
// the expiry radar. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/leases/{leaseId}/cancel
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *LeasesNamespace) Cancel(ctx context.Context, params LeasesCancelParams, opts ...RequestOption) (*ResourceLease, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/leases/{leaseId}/cancel")
	r.setPath("orgId", params.OrgID)
	r.setPath("leaseId", params.LeaseID)
	var out *ResourceLease
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LeasesCreateParams holds the parameters for `client.leases.create`.
//
// Every field is optional; pass nil to take the defaults.
type LeasesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *ResourceLeaseCreate
}

// Create: Create a resource lease
//
// Attach an expiry to a resource — 'give me a test cluster for 3 days'. One
// lease per resource (an active lease conflicts; a terminal one is replaced).
// `autoDelete: true` opts into deletion at expiry — the poller announces it
// twice first, defers during change freezes, and requires the caller to hold
// `resources:delete`. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/leases
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: The resource already has an active lease
func (n *LeasesNamespace) Create(ctx context.Context, params *LeasesCreateParams, opts ...RequestOption) (*ResourceLease, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/leases")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *ResourceLease
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LeasesDeleteParams holds the parameters for `client.leases.delete`.
type LeasesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	LeaseID string
}

// Delete: Delete a lease row
//
// Remove the lease record entirely (including terminal rows). The resource is
// not touched. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// DELETE /api/org/{orgId}/leases/{leaseId}
//
// Raises on 404: Not found
func (n *LeasesNamespace) Delete(ctx context.Context, params LeasesDeleteParams, opts ...RequestOption) error {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/leases/{leaseId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("leaseId", params.LeaseID)
	return n.t.do(ctx, r, nil, opts)
}

// LeasesGetParams holds the parameters for `client.leases.get`.
//
// Every field is optional; pass nil to take the defaults.
type LeasesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List resource leases
//
// Every lease in the organization, soonest deadline first. Active leases also
// appear on the expiry radar (`GET /expiring`) as kind `lease` items, so the
// owner is nagged through the existing expiry alerts.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/leases
func (n *LeasesNamespace) Get(ctx context.Context, params *LeasesGetParams, opts ...RequestOption) (*ResourceLeaseList, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/leases")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *ResourceLeaseList
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LeasesResourceParams holds the parameters for `client.leases.resource`.
type LeasesResourceParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ResourceID string
}

// Resource: Get one resource's lease
//
// The (unique) lease on a resource, whatever its status, or null.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/leases/resource
//
// Raises on 400: Bad request
func (n *LeasesNamespace) Resource(ctx context.Context, params LeasesResourceParams, opts ...RequestOption) (*ResourceLeaseLookup, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/leases/resource")
	r.setPath("orgId", params.OrgID)
	r.addQuery("resourceId", params.ResourceID)
	var out *ResourceLeaseLookup
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LeasesUpdateParams holds the parameters for `client.leases.update`.
type LeasesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	LeaseID string
	// Body: the JSON request body.
	Body *ResourceLeaseUpdate
}

// Update: Update a lease
//
// Edit the deadline, the auto-delete opt-in and/or the note of an active lease.
// Changing the deadline or the auto-delete flag re-arms the two-announcement
// schedule. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// PUT /api/org/{orgId}/leases/{leaseId}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *LeasesNamespace) Update(ctx context.Context, params LeasesUpdateParams, opts ...RequestOption) (*ResourceLease, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/leases/{leaseId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("leaseId", params.LeaseID)
	r.setJSONBody(params.Body)
	var out *ResourceLease
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LinearNamespace is `client.linear`.
type LinearNamespace struct {
	t *transport
}

func newLinearNamespace(t *transport) *LinearNamespace {
	n := &LinearNamespace{t: t}
	return n
}

// LinearDeleteParams holds the parameters for `client.linear.delete`.
//
// Every field is optional; pass nil to take the defaults.
type LinearDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Delete: Disconnect Linear
//
// Issue links already recorded are kept, so filed findings stay marked as filed.
//
// _Requires permission: `linear:write`._
//
// DELETE /api/org/{orgId}/linear
//
// Raises on 404: Not found
func (n *LinearNamespace) Delete(ctx context.Context, params *LinearDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/linear")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LinearGetParams holds the parameters for `client.linear.get`.
//
// Every field is optional; pass nil to take the defaults.
type LinearGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the org's Linear connection
//
// The stored API key is never returned; `keyHint` stands in for it.
//
// _Requires permission: `linear:read`._
//
// GET /api/org/{orgId}/linear
func (n *LinearNamespace) Get(ctx context.Context, params *LinearGetParams, opts ...RequestOption) (*LinearGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/linear")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *LinearGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LinearIssuesParams holds the parameters for `client.linear.issues`.
type LinearIssuesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateLinearIssueInput
}

// Issues: File a finding as a Linear issue
//
// Creates the issue via the issueCreate mutation, then records the link between
// it and the finding. The link is what lets a list view show "already filed"
// instead of offering the button again.
//
// _Requires permission: `linear:write`._
//
// POST /api/org/{orgId}/linear/issues
//
// Raises on 400: Bad request
//
// Raises on 502: Linear refused to create the issue, or was unreachable
func (n *LinearNamespace) Issues(ctx context.Context, params LinearIssuesParams, opts ...RequestOption) (*CreateLinearIssueResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/linear/issues")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CreateLinearIssueResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LinearLinksParams holds the parameters for `client.linear.links`.
//
// Every field is optional; pass nil to take the defaults.
type LinearLinksParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	SourceKind *LinearSourceKind
	// SourceID: Repeat to narrow to specific findings. Omit to return every link
	// of the kind — this is the batch lookup a list view makes once instead of
	// one request per row.
	SourceID []string
}

// Links: Look up filed issues for a set of findings
//
// _Requires permission: `linear:read`._
//
// GET /api/org/{orgId}/linear/links
//
// Raises on 400: Bad request
func (n *LinearNamespace) Links(ctx context.Context, params *LinearLinksParams, opts ...RequestOption) ([]LinearIssueLink, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/linear/links")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("sourceKind", params.SourceKind)
		r.addQuery("sourceId", params.SourceID)
	}
	var out []LinearIssueLink
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LinearTeamsParams holds the parameters for `client.linear.teams`.
//
// Every field is optional; pass nil to take the defaults.
type LinearTeamsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Teams: List Linear teams
//
// Backs the team picker, so nobody has to know a team id by hand — issueCreate
// requires one, and every issue belongs to exactly one team.
//
// _Requires permission: `linear:read`._
//
// GET /api/org/{orgId}/linear/teams
//
// Raises on 400: Bad request
func (n *LinearNamespace) Teams(ctx context.Context, params *LinearTeamsParams, opts ...RequestOption) ([]LinearTeam, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/linear/teams")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []LinearTeam
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LinearUpdateParams holds the parameters for `client.linear.update`.
type LinearUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body LinearIntegrationInput
}

// Update: Connect Linear, or update the connection
//
// _Requires permission: `linear:write`._
//
// PUT /api/org/{orgId}/linear
//
// Raises on 400: Bad request
func (n *LinearNamespace) Update(ctx context.Context, params LinearUpdateParams, opts ...RequestOption) (*LinearIntegration, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/linear")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *LinearIntegration
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LinearVerifyParams holds the parameters for `client.linear.verify`.
//
// Every field is optional; pass nil to take the defaults.
type LinearVerifyParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *LinearVerifyInput
}

// Verify: Check Linear credentials
//
// Runs the `viewer` query against the Linear GraphQL API, so a mistyped or
// revoked key is reported on the settings form rather than on the first attempt
// to file an issue.
//
// _Requires permission: `linear:write`._
//
// POST /api/org/{orgId}/linear/verify
//
// Raises on 400: Bad request
//
// Raises on 502: Linear rejected the key or was unreachable
func (n *LinearNamespace) Verify(ctx context.Context, params *LinearVerifyParams, opts ...RequestOption) (*LinearVerifyResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/linear/verify")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *LinearVerifyResult
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LogWorkspacesNamespace is `client.logWorkspaces`.
type LogWorkspacesNamespace struct {
	t *transport
}

func newLogWorkspacesNamespace(t *transport) *LogWorkspacesNamespace {
	n := &LogWorkspacesNamespace{t: t}
	return n
}

// LogWorkspacesCreateParams holds the parameters for
// `client.logWorkspaces.create`.
type LogWorkspacesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body LogWorkspaceQueryCreate
}

// Create: Save a log workspace query
//
// Save a named multi-resource tail: up to 8 log streams plus a search
// expression, so the workspace can be reopened. With `alertEnabled` the poller
// evaluates the query every few minutes over a bounded tail window and notifies
// (push/Slack/Teams, `logMatchAlerts` trigger) when a line matches, with a
// cooldown between alerts. Alerting requires a non-empty search expression.
// Audit-logged.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/log-workspaces
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: A saved query with this name already exists
func (n *LogWorkspacesNamespace) Create(ctx context.Context, params LogWorkspacesCreateParams, opts ...RequestOption) (*LogWorkspaceQuery, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/log-workspaces")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *LogWorkspaceQuery
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LogWorkspacesDeleteParams holds the parameters for
// `client.logWorkspaces.delete`.
type LogWorkspacesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	QueryID string
}

// Delete: Delete a saved log query
//
// Remove the saved query and stop any alerting it carried. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// DELETE /api/org/{orgId}/log-workspaces/{queryId}
//
// Raises on 404: Not found
func (n *LogWorkspacesNamespace) Delete(ctx context.Context, params LogWorkspacesDeleteParams, opts ...RequestOption) error {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/log-workspaces/{queryId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("queryId", params.QueryID)
	return n.t.do(ctx, r, nil, opts)
}

// LogWorkspacesGetParams holds the parameters for `client.logWorkspaces.get`.
//
// Every field is optional; pass nil to take the defaults.
type LogWorkspacesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List saved log queries
//
// Every saved log-workspace query in the organization: its name, the set of log
// streams it tails, the search expression, the alert flag and the alert pass's
// last evaluation state. Log text itself is fetched per resource via `POST
// /api/org/{orgId}/resources/{pluginId}/{typeId}/logs`.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/log-workspaces
func (n *LogWorkspacesNamespace) Get(ctx context.Context, params *LogWorkspacesGetParams, opts ...RequestOption) (*LogWorkspaceQueryList, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/log-workspaces")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *LogWorkspaceQueryList
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LogWorkspacesResourcesParams holds the parameters for
// `client.logWorkspaces.resources`.
//
// Every field is optional; pass nil to take the defaults.
type LogWorkspacesResourcesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Resources: List log-capable resources
//
// Synced resources whose rendered detail declares the logs capability — the
// candidates a log workspace can tail — plus sidecar streams reached through a
// peer integration (pods and workloads inside a managed cluster, listed live
// from the provider and marked with `parentResourceId`). Discovered from the
// plugin contract (never a hardcoded provider list), capped at 500 results.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/log-workspaces/resources
func (n *LogWorkspacesNamespace) Resources(ctx context.Context, params *LogWorkspacesResourcesParams, opts ...RequestOption) (*LogCapableResourceList, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/log-workspaces/resources")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *LogCapableResourceList
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// LogWorkspacesUpdateParams holds the parameters for
// `client.logWorkspaces.update`.
type LogWorkspacesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	QueryID string
	// Body: the JSON request body.
	Body LogWorkspaceQueryUpdate
}

// Update: Update a saved log query
//
// Edit the name, resource set, search expression and/or the alert toggle.
// Changing the search or the resources resets the alert pass's evaluation state;
// turning the alert on makes the query due for evaluation immediately.
// Audit-logged.
//
// _Requires permission: `resources:write`._
//
// PUT /api/org/{orgId}/log-workspaces/{queryId}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: A saved query with this name already exists
func (n *LogWorkspacesNamespace) Update(ctx context.Context, params LogWorkspacesUpdateParams, opts ...RequestOption) (*LogWorkspaceQuery, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/log-workspaces/{queryId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("queryId", params.QueryID)
	r.setJSONBody(params.Body)
	var out *LogWorkspaceQuery
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ManagedAccountsNamespace is `client.managedAccounts`.
type ManagedAccountsNamespace struct {
	t *transport
}

func newManagedAccountsNamespace(t *transport) *ManagedAccountsNamespace {
	n := &ManagedAccountsNamespace{t: t}
	return n
}

// ManagedAccountsCreateParams holds the parameters for
// `client.managedAccounts.create`.
type ManagedAccountsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ManagedAccountInput
}

// Create: Create a managed account
//
// Refused with 409 when a cost centre or account named here is already billed to
// another customer. The error names the other customer, because “it conflicts”
// without saying with whom sends the caller hunting.
//
// _Requires permission: `invoices:write`._
//
// POST /api/org/{orgId}/managed-accounts
//
// Raises on 400: Bad request
//
// Raises on 409: Conflict
func (n *ManagedAccountsNamespace) Create(ctx context.Context, params ManagedAccountsCreateParams, opts ...RequestOption) (*ManagedAccount, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/managed-accounts")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ManagedAccount
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ManagedAccountsDeleteParams holds the parameters for
// `client.managedAccounts.delete`.
type ManagedAccountsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Retire a managed account
//
// A soft delete: an issued invoice names its customer, and an invoice whose
// customer stopped resolving is exactly the unreconcilable document this feature
// exists to prevent. Draft invoices are removed with it — a draft was never
// issued.
//
// _Requires permission: `invoices:write`._
//
// DELETE /api/org/{orgId}/managed-accounts/{id}
//
// Raises on 404: Not found
func (n *ManagedAccountsNamespace) Delete(ctx context.Context, params ManagedAccountsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/managed-accounts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ManagedAccountsGetParams holds the parameters for
// `client.managedAccounts.get`.
type ManagedAccountsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a managed account
//
// _Requires permission: `invoices:read`._
//
// GET /api/org/{orgId}/managed-accounts/{id}
//
// Raises on 404: Not found
func (n *ManagedAccountsNamespace) Get(ctx context.Context, params ManagedAccountsGetParams, opts ...RequestOption) (*ManagedAccount, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/managed-accounts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *ManagedAccount
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ManagedAccountsListParams holds the parameters for
// `client.managedAccounts.list`.
//
// Every field is optional; pass nil to take the defaults.
type ManagedAccountsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List managed accounts
//
// The customers a managed service provider bills. A managed account references
// existing cost centres rather than defining its own matching rules, so the
// spend on an invoice is the same spend the showback report attributes to those
// centres.
//
// _Requires permission: `invoices:read`._
//
// GET /api/org/{orgId}/managed-accounts
func (n *ManagedAccountsNamespace) List(ctx context.Context, params *ManagedAccountsListParams, opts ...RequestOption) ([]ManagedAccount, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/managed-accounts")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []ManagedAccount
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ManagedAccountsUpdateParams holds the parameters for
// `client.managedAccounts.update`.
type ManagedAccountsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body ManagedAccountInput
}

// Update: Update a managed account
//
// A full replace. Editing the scope changes what **future** drafts are drawn
// over and nothing else: every approved invoice holds its own copy of the scope,
// so moving a cost centre between customers cannot re-bill a period that has
// already been invoiced.
//
// _Requires permission: `invoices:write`._
//
// PUT /api/org/{orgId}/managed-accounts/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *ManagedAccountsNamespace) Update(ctx context.Context, params ManagedAccountsUpdateParams, opts ...RequestOption) (*ManagedAccount, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/managed-accounts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *ManagedAccount
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsNamespace is `client.metricAlerts`.
type MetricAlertsNamespace struct {
	t *transport
}

func newMetricAlertsNamespace(t *transport) *MetricAlertsNamespace {
	n := &MetricAlertsNamespace{t: t}
	return n
}

// MetricAlertsCreateParams holds the parameters for
// `client.metricAlerts.create`.
type MetricAlertsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body MetricAlertRuleInput
}

// Create: Create a metric alert rule
//
// Rules select resources by query (plugin + resource type + tag), never by id
// list, so a rule automatically covers resources created after it was written.
// The poller evaluates enabled rules about once a minute and alerts when the
// condition held for the whole trailing window.
//
// POST /api/org/{orgId}/metric-alerts
//
// Raises on 400: Bad request
func (n *MetricAlertsNamespace) Create(ctx context.Context, params MetricAlertsCreateParams, opts ...RequestOption) (*MetricAlertRule, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/metric-alerts")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *MetricAlertRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsDeleteParams holds the parameters for
// `client.metricAlerts.delete`.
type MetricAlertsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a metric alert rule
//
// Soft delete. The rule's firing history stays readable via
// /metric-alerts/events.
//
// DELETE /api/org/{orgId}/metric-alerts/{id}
//
// Raises on 404: Not found
func (n *MetricAlertsNamespace) Delete(ctx context.Context, params MetricAlertsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/metric-alerts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsEventsParams holds the parameters for
// `client.metricAlerts.events`.
//
// Every field is optional; pass nil to take the defaults.
type MetricAlertsEventsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID  *string
	RuleID *string
	Limit  *int64
}

// Events: Recent metric alert firings
//
// GET /api/org/{orgId}/metric-alerts/events
func (n *MetricAlertsNamespace) Events(ctx context.Context, params *MetricAlertsEventsParams, opts ...RequestOption) ([]MetricAlertEvent, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/metric-alerts/events")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("ruleId", params.RuleID)
		r.addQuery("limit", params.Limit)
	}
	var out []MetricAlertEvent
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsGetParams holds the parameters for `client.metricAlerts.get`.
type MetricAlertsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a metric alert rule
//
// GET /api/org/{orgId}/metric-alerts/{id}
//
// Raises on 404: Not found
func (n *MetricAlertsNamespace) Get(ctx context.Context, params MetricAlertsGetParams, opts ...RequestOption) (*MetricAlertRule, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/metric-alerts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *MetricAlertRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsListParams holds the parameters for `client.metricAlerts.list`.
//
// Every field is optional; pass nil to take the defaults.
type MetricAlertsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List metric alert rules with live firing status
//
// GET /api/org/{orgId}/metric-alerts
func (n *MetricAlertsNamespace) List(ctx context.Context, params *MetricAlertsListParams, opts ...RequestOption) ([]MetricAlertRuleWithStatus, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/metric-alerts")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []MetricAlertRuleWithStatus
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsMetricKeysParams holds the parameters for
// `client.metricAlerts.metricKeys`.
//
// Every field is optional; pass nil to take the defaults.
type MetricAlertsMetricKeysParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID          *string
	PluginID       *string
	ResourceTypeID *string
}

// MetricKeys: List metric series that actually exist
//
// The series labels resources reported in the last 7 days, optionally narrowed
// to one plugin and resource type — what the rule builder's metric picker is fed
// from.
//
// GET /api/org/{orgId}/metric-alerts/metric-keys
func (n *MetricAlertsNamespace) MetricKeys(ctx context.Context, params *MetricAlertsMetricKeysParams, opts ...RequestOption) ([]MetricSeriesKey, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/metric-alerts/metric-keys")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("pluginId", params.PluginID)
		r.addQuery("resourceTypeId", params.ResourceTypeID)
	}
	var out []MetricSeriesKey
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsSelectorOptionsParams holds the parameters for
// `client.metricAlerts.selectorOptions`.
//
// Every field is optional; pass nil to take the defaults.
type MetricAlertsSelectorOptionsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// SelectorOptions: List what the organization's resources offer to select on
//
// GET /api/org/{orgId}/metric-alerts/selector-options
func (n *MetricAlertsNamespace) SelectorOptions(ctx context.Context, params *MetricAlertsSelectorOptionsParams, opts ...RequestOption) (*MetricAlertSelectorOptions, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/metric-alerts/selector-options")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *MetricAlertSelectorOptions
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsSelectorPreviewParams holds the parameters for
// `client.metricAlerts.selectorPreview`.
//
// Every field is optional; pass nil to take the defaults.
type MetricAlertsSelectorPreviewParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID          *string
	PluginID       *string
	ResourceTypeID *string
	TagKey         *string
	TagValue       *string
}

// SelectorPreview: Preview which resources a selector matches right now
//
// GET /api/org/{orgId}/metric-alerts/selector-preview
//
// Raises on 400: Bad request
func (n *MetricAlertsNamespace) SelectorPreview(ctx context.Context, params *MetricAlertsSelectorPreviewParams, opts ...RequestOption) (*MetricAlertSelectorPreview, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/metric-alerts/selector-preview")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("pluginId", params.PluginID)
		r.addQuery("resourceTypeId", params.ResourceTypeID)
		r.addQuery("tagKey", params.TagKey)
		r.addQuery("tagValue", params.TagValue)
	}
	var out *MetricAlertSelectorPreview
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MetricAlertsUpdateParams holds the parameters for
// `client.metricAlerts.update`.
type MetricAlertsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body MetricAlertRuleInput
}

// Update: Update a metric alert rule
//
// PUT /api/org/{orgId}/metric-alerts/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *MetricAlertsNamespace) Update(ctx context.Context, params MetricAlertsUpdateParams, opts ...RequestOption) (*MetricAlertRule, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/metric-alerts/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *MetricAlertRule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MomentNamespace is `client.moment`.
type MomentNamespace struct {
	t *transport
}

func newMomentNamespace(t *transport) *MomentNamespace {
	n := &MomentNamespace{t: t}
	return n
}

// MomentGetParams holds the parameters for `client.moment.get`.
//
// Every field is optional; pass nil to take the defaults.
type MomentGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// At: Centre of the window. Defaults to now.
	At *string
	// Window: Half-window in minutes (the ± around `at`). Default 60, max 4320
	// (±3 days).
	Window *int64
}

// Get: Everything that happened around a timestamp
//
// "What changed around 03:14?" — one merged, chronological narrative of
// everything the platform knows happened in a window: resource changes
// (including sleep/wake schedule attribution), provider status incidents that
// started/resolved in or overlap the window, cost anomalies, workflow runs,
// deployments, audit-log entries, change freezes, and the drift/expiry alert
// deliveries. Each feed is gated on the same permission its own endpoint
// requires; feeds the caller cannot read are reported as `omitted`, and a feed
// whose query fails is reported as `error` without blanking the rest of the
// response.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/moment
//
// Raises on 400: Bad request
func (n *MomentNamespace) Get(ctx context.Context, params *MomentGetParams, opts ...RequestOption) (*MomentResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/moment")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("at", params.At)
		r.addQuery("window", params.Window)
	}
	var out *MomentResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MsteamsNamespace is `client.msteams`.
type MsteamsNamespace struct {
	t *transport

	// Webhooks: `client.msteams.webhooks`.
	Webhooks *MsteamsWebhooksNamespace
}

func newMsteamsNamespace(t *transport) *MsteamsNamespace {
	n := &MsteamsNamespace{t: t}
	n.Webhooks = newMsteamsWebhooksNamespace(t)
	return n
}

// MsteamsStatusParams holds the parameters for `client.msteams.status`.
//
// Every field is optional; pass nil to take the defaults.
type MsteamsStatusParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Status: List the organization's Teams channels
//
// Returns the Teams channels alerts can be routed to. Which alerts reach each
// one is decided by /alert-rules. Webhook URLs are never included.
//
// GET /api/org/{orgId}/msteams/status
func (n *MsteamsNamespace) Status(ctx context.Context, params *MsteamsStatusParams, opts ...RequestOption) (*MsTeamsStatus, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/msteams/status")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *MsTeamsStatus
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MsteamsTestParams holds the parameters for `client.msteams.test`.
//
// Every field is optional; pass nil to take the defaults.
type MsteamsTestParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Test: Post a test card to every configured Teams channel
//
// Ignores routing rules — every channel gets the test. Fails with the error
// Microsoft returned when nothing could be delivered (HTTP 404 usually means the
// Workflow was deleted or turned off).
//
// POST /api/org/{orgId}/msteams/test
//
// Raises on 400: Bad request
func (n *MsteamsNamespace) Test(ctx context.Context, params *MsteamsTestParams, opts ...RequestOption) (*MsteamsTestResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/msteams/test")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *MsteamsTestResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MsteamsWebhooksNamespace is `client.msteams.webhooks`.
type MsteamsWebhooksNamespace struct {
	t *transport
}

func newMsteamsWebhooksNamespace(t *transport) *MsteamsWebhooksNamespace {
	n := &MsteamsWebhooksNamespace{t: t}
	return n
}

// MsteamsWebhooksCreateParams holds the parameters for
// `client.msteams.webhooks.create`.
//
// Every field is optional; pass nil to take the defaults.
type MsteamsWebhooksCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *MsTeamsWebhookCreate
}

// Create: Connect a Teams channel as an alert destination
//
// Adds a channel by webhook URL, or updates the one already holding that URL.
// Which alerts reach it is decided by /alert-rules — connecting a channel routes
// nothing to it on its own. Responds 400 when the URL is not https or its host
// is not Microsoft-operated.
//
// POST /api/org/{orgId}/msteams/webhooks
//
// Raises on 400: Bad request
func (n *MsteamsWebhooksNamespace) Create(ctx context.Context, params *MsteamsWebhooksCreateParams, opts ...RequestOption) (*MsTeamsWebhook, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/msteams/webhooks")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *MsTeamsWebhook
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MsteamsWebhooksDeleteParams holds the parameters for
// `client.msteams.webhooks.delete`.
type MsteamsWebhooksDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Disconnect a Teams channel
//
// DELETE /api/org/{orgId}/msteams/webhooks/{id}
//
// Raises on 404: Not found
func (n *MsteamsWebhooksNamespace) Delete(ctx context.Context, params MsteamsWebhooksDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/msteams/webhooks/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// MsteamsWebhooksUpdateParams holds the parameters for
// `client.msteams.webhooks.update`.
type MsteamsWebhooksUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body *MsTeamsWebhookUpdate
}

// Update: Rename a Teams channel
//
// The webhook URL is immutable — remove the channel and re-add it to change it.
//
// PATCH /api/org/{orgId}/msteams/webhooks/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *MsteamsWebhooksNamespace) Update(ctx context.Context, params MsteamsWebhooksUpdateParams, opts ...RequestOption) (*MsTeamsWebhook, error) {
	r := newRequest(http.MethodPatch, "/api/org/{orgId}/msteams/webhooks/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *MsTeamsWebhook
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// NetworkFlowsNamespace is `client.networkFlows`.
type NetworkFlowsNamespace struct {
	t *transport

	// Settings: `client.networkFlows.settings`.
	Settings *NetworkFlowsSettingsNamespace
}

func newNetworkFlowsNamespace(t *transport) *NetworkFlowsNamespace {
	n := &NetworkFlowsNamespace{t: t}
	n.Settings = newNetworkFlowsSettingsNamespace(t)
	return n
}

// NetworkFlowsGetParams holds the parameters for `client.networkFlows.get`.
//
// Every field is optional; pass nil to take the defaults.
type NetworkFlowsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// From: Inclusive start day. Defaults to 13 days ago.
	From *string
	// To: Inclusive end day. Defaults to today.
	To *string
	// Scope: Narrow to one billing boundary.
	//
	// One of "intra_zone", "cross_zone", "cross_region", "internet_egress",
	// "internet_ingress", "provider_service", "nat_gateway",
	// "private_interconnect", "unknown".
	Scope *string
	// AccountID: Narrow to one connected account.
	AccountID *string
	// Limit: Pairs to return in `topFlows`, largest cost first. Defaults to 50.
	Limit *int64
}

// Get: Priced source→destination network flow attribution
//
// Which two things are talking, across which billing boundary, and what that
// costs. Answers the question the cost dimensions structurally cannot: every
// cost dimension is about one side of a transfer, and a network charge is about
// a pair.
//
// All figures are **estimates** and the `estimated` field says so
// unconditionally. Bytes come from the provider's flow logs (which sample, or
// drop records under capacity pressure) and are priced at published list rates
// with no free tier, no volume tier and no negotiated discount applied. Use the
// ranking; do not reconcile the total against an invoice line.
//
// Accounts whose provider has no readable flow source appear in `accounts` with
// `supportsFlows: false` and contribute nothing to the totals — never zero
// bytes.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/network-flows
//
// Raises on 400: Bad request
func (n *NetworkFlowsNamespace) Get(ctx context.Context, params *NetworkFlowsGetParams, opts ...RequestOption) (*NetworkFlowFeed, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/network-flows")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("from", params.From)
		r.addQuery("to", params.To)
		r.addQuery("scope", params.Scope)
		r.addQuery("accountId", params.AccountID)
		r.addQuery("limit", params.Limit)
	}
	var out *NetworkFlowFeed
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// NetworkFlowsSettingsNamespace is `client.networkFlows.settings`.
type NetworkFlowsSettingsNamespace struct {
	t *transport
}

func newNetworkFlowsSettingsNamespace(t *transport) *NetworkFlowsSettingsNamespace {
	n := &NetworkFlowsSettingsNamespace{t: t}
	return n
}

// NetworkFlowsSettingsGetParams holds the parameters for
// `client.networkFlows.settings.get`.
//
// Every field is optional; pass nil to take the defaults.
type NetworkFlowsSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Read the network flow collection switch
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/network-flows/settings
func (n *NetworkFlowsSettingsNamespace) Get(ctx context.Context, params *NetworkFlowsSettingsGetParams, opts ...RequestOption) (*NetworkFlowSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/network-flows/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *NetworkFlowSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// NetworkFlowsSettingsUpdateParams holds the parameters for
// `client.networkFlows.settings.update`.
type NetworkFlowsSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body NetworkFlowSettings
}

// Update: Turn network flow collection on or off
//
// Collection is **off by default**. Enabling it authorizes Infrawrench to run
// daily queries against the provider's log store — and on AWS those queries are
// billed to your own cloud account per GB of log data scanned, every day, until
// you turn them off. That is why the write is governed by `org:settings:write`
// rather than `costs:write`, and why it is audit-logged.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/network-flows/settings
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
func (n *NetworkFlowsSettingsNamespace) Update(ctx context.Context, params NetworkFlowsSettingsUpdateParams, opts ...RequestOption) (*NetworkFlowSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/network-flows/settings")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *NetworkFlowSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// OrgsNamespace is `client.orgs`.
type OrgsNamespace struct {
	t *transport
}

func newOrgsNamespace(t *transport) *OrgsNamespace {
	n := &OrgsNamespace{t: t}
	return n
}

// OrgsCreateParams holds the parameters for `client.orgs.create`.
type OrgsCreateParams struct {
	// Body: the JSON request body.
	Body CreateOrgRequest
}

// Create: Create a new organization
//
// The caller becomes the `owner` of the new organization.
//
// POST /api/orgs
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
func (n *OrgsNamespace) Create(ctx context.Context, params OrgsCreateParams, opts ...RequestOption) (*Organization, error) {
	r := newRequest(http.MethodPost, "/api/orgs")
	r.setJSONBody(params.Body)
	var out *Organization
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// OrphansNamespace is `client.orphans`.
type OrphansNamespace struct {
	t *transport
}

func newOrphansNamespace(t *transport) *OrphansNamespace {
	n := &OrphansNamespace{t: t}
	return n
}

// OrphansGetParams holds the parameters for `client.orphans.get`.
//
// Every field is optional; pass nil to take the defaults.
type OrphansGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List likely-orphaned and idle resources
//
// Scans the organization's already-synced resources against each plugin's
// declarative orphan heuristics — unattached volumes, unassigned
// floating/elastic IPs, reserved-but-unused static IPs — and returns the matches
// grouped by account, each with the plugin's reason. Purely a read over stored
// state: no provider API calls are made, so results reflect the last sync. Where
// the org's collected cost data has per-resource rows, matches are annotated
// with trailing spend.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/orphans
func (n *OrphansNamespace) Get(ctx context.Context, params *OrphansGetParams, opts ...RequestOption) (*OrphanListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/orphans")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *OrphanListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// OwnershipNamespace is `client.ownership`.
type OwnershipNamespace struct {
	t *transport
}

func newOwnershipNamespace(t *transport) *OwnershipNamespace {
	n := &OwnershipNamespace{t: t}
	return n
}

// OwnershipDeleteParams holds the parameters for `client.ownership.delete`.
type OwnershipDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ResourceID string
}

// Delete: Clear a resource's ownership
//
// Removes the ownership record. The resource itself is untouched.
//
// _Requires permission: `resources:write`._
//
// DELETE /api/org/{orgId}/ownership
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *OwnershipNamespace) Delete(ctx context.Context, params OwnershipDeleteParams, opts ...RequestOption) error {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/ownership")
	r.setPath("orgId", params.OrgID)
	r.addQuery("resourceId", params.ResourceID)
	return n.t.do(ctx, r, nil, opts)
}

// OwnershipGetParams holds the parameters for `client.ownership.get`.
//
// Every field is optional; pass nil to take the defaults.
type OwnershipGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List resource ownership records
//
// Every ownership record in the organization — owner, purpose and authorizing
// ticket, per resource. Only resources somebody has recorded something about
// appear; an absent record means the resource is unowned.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/ownership
func (n *OwnershipNamespace) Get(ctx context.Context, params *OwnershipGetParams, opts ...RequestOption) (*ResourceOwnershipListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/ownership")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *ResourceOwnershipListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// OwnershipMembersParams holds the parameters for `client.ownership.members`.
//
// Every field is optional; pass nil to take the defaults.
type OwnershipMembersParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Members: List people an owner can be set to
//
// Org members, as a minimal id/name/email projection for the owner picker.
// Requires only `resources:read`, deliberately not `team:read`: recording who
// owns a resource must not be reserved for whoever can also read roles and
// membership.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/ownership/members
func (n *OwnershipNamespace) Members(ctx context.Context, params *OwnershipMembersParams, opts ...RequestOption) (*OwnerCandidateListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/ownership/members")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *OwnerCandidateListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// OwnershipResourceParams holds the parameters for `client.ownership.resource`.
type OwnershipResourceParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ResourceID string
}

// Resource: Get one resource's ownership
//
// The ownership record for a single resource, or null when none is recorded.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/ownership/resource
//
// Raises on 400: Bad request
func (n *OwnershipNamespace) Resource(ctx context.Context, params OwnershipResourceParams, opts ...RequestOption) (*ResourceOwnershipEnvelope, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/ownership/resource")
	r.setPath("orgId", params.OrgID)
	r.addQuery("resourceId", params.ResourceID)
	var out *ResourceOwnershipEnvelope
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// OwnershipUpdateParams holds the parameters for `client.ownership.update`.
//
// Every field is optional; pass nil to take the defaults.
type OwnershipUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *ResourceOwnershipPatch
}

// Update: Set a resource's ownership
//
// Upsert keyed by `resourceId` — ownership is a property of the resource, so
// there is no separate create and update. Omitted fields keep their value and
// `null` clears one. Clearing every field removes the record entirely and the
// response is `null`, which is the new truth rather than an empty record. An
// `ownerUserId` must be a member of this organization: ownership that looks
// routable but reaches nobody is worse than none.
//
// _Requires permission: `resources:write`._
//
// PUT /api/org/{orgId}/ownership
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *OwnershipNamespace) Update(ctx context.Context, params *OwnershipUpdateParams, opts ...RequestOption) (any, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/ownership")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out any
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// PagesNamespace is `client.pages`.
type PagesNamespace struct {
	t *transport
}

func newPagesNamespace(t *transport) *PagesNamespace {
	n := &PagesNamespace{t: t}
	return n
}

// PagesCreateParams holds the parameters for `client.pages.create`.
type PagesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body PageRequest
}

// Create: Raise an alert to the organization's on-call transports
//
// Fans an alert out over whatever the org has configured — Twilio SMS (and voice
// on request), mobile push, Slack channels, and Microsoft Teams webhooks —
// honouring each recipient's opt-ins. This is the same alert a workflow raises
// with `infra.page(...)`, for code that runs somewhere Infrawrench does not: a
// health check, a deploy script, a cron on a box.
//
// Repeat pages under the same `(source, key)` are **suppressed, not rejected**:
// a monitor that fires every minute pages once and then gets `200` with
// `suppressed: true` and the `retryAt` at which the key can page again. A page
// that reached nobody does not start a cooldown, so the next call tries again.
//
// Recipients opt in per channel under the same setting that covers workflow
// pages.
//
// _Requires permission: `pages:write`._
//
// POST /api/org/{orgId}/pages
//
// Raises on 400: Bad request
func (n *PagesNamespace) Create(ctx context.Context, params PagesCreateParams, opts ...RequestOption) (*PageResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/pages")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *PageResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// PagesDeleteParams holds the parameters for `client.pages.delete`.
type PagesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Source: Stable name for the system raising the page: letters, digits, `.`,
	// `_` and `-`. It is the notification's sender, and it scopes the cooldown —
	// two services paging under the same key never throttle each other.
	Source string
	// Key: Defaults to `default`.
	Key *string
}

// Delete: Clear a page key's cooldown
//
// Drops the cooldown for one `(source, key)` so the next page under it delivers
// immediately. Call it when the condition you alerted on recovers — the workflow
// equivalent is `infra.page.clear(key)`. Clearing a key that was never paged is
// not an error.
//
// _Requires permission: `pages:write`._
//
// DELETE /api/org/{orgId}/pages
//
// Raises on 400: Bad request
func (n *PagesNamespace) Delete(ctx context.Context, params PagesDeleteParams, opts ...RequestOption) (*PageClearResponse, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/pages")
	r.setPath("orgId", params.OrgID)
	r.addQuery("source", params.Source)
	r.addQuery("key", params.Key)
	var out *PageClearResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// PostureNamespace is `client.posture`.
type PostureNamespace struct {
	t *transport

	// Dismissals: `client.posture.dismissals`.
	Dismissals *PostureDismissalsNamespace
	// Settings: `client.posture.settings`.
	Settings *PostureSettingsNamespace
}

func newPostureNamespace(t *transport) *PostureNamespace {
	n := &PostureNamespace{t: t}
	n.Dismissals = newPostureDismissalsNamespace(t)
	n.Settings = newPostureSettingsNamespace(t)
	return n
}

// PostureGetParams holds the parameters for `client.posture.get`.
//
// Every field is optional; pass nil to take the defaults.
type PostureGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List security posture findings on synced resources
//
// Plugin-declared security checks evaluated over already-synced resource state:
// public buckets, 0.0.0.0/0 ingress rules, unencrypted disks, publicly reachable
// database endpoints, stale credentials, missing deletion/backup protection. No
// provider API calls are made and results reflect the last sync. Findings are
// sorted worst severity first, with per-severity counts. Findings the
// organization has dismissed are reported separately under `dismissed` and are
// excluded from `findings`, `counts` and the posture alerts.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/posture
func (n *PostureNamespace) Get(ctx context.Context, params *PostureGetParams, opts ...RequestOption) (*PostureListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/posture")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *PostureListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// PostureDismissalsNamespace is `client.posture.dismissals`.
type PostureDismissalsNamespace struct {
	t *transport
}

func newPostureDismissalsNamespace(t *transport) *PostureDismissalsNamespace {
	n := &PostureDismissalsNamespace{t: t}
	return n
}

// PostureDismissalsCreateParams holds the parameters for
// `client.posture.dismissals.create`.
//
// Every field is optional; pass nil to take the defaults.
type PostureDismissalsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *PostureDismissalCreate
}

// Create: Dismiss a posture finding
//
// Accept a finding — the bucket really is meant to be public, the key really is
// rotated out of band. The finding leaves `findings` and stops feeding the daily
// posture alerts, but the rule keeps being evaluated and the finding is reported
// back under `dismissed` for as long as it still matches. Idempotent: dismissing
// an already-dismissed finding rewrites the note and the author.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/posture/dismissals
//
// Raises on 400: Bad request
func (n *PostureDismissalsNamespace) Create(ctx context.Context, params *PostureDismissalsCreateParams, opts ...RequestOption) (*PostureDismissal, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/posture/dismissals")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *PostureDismissal
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// PostureDismissalsDeleteParams holds the parameters for
// `client.posture.dismissals.delete`.
type PostureDismissalsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ResourceID: Infrawrench resource id the finding is on.
	ResourceID string
	// RuleID: The matched rule's id.
	RuleID string
}

// Delete: Restore a dismissed posture finding
//
// Undo a dismissal, putting the finding back on the list and back into the alert
// feed. The finding is identified by query parameters rather than path segments
// because resource ids are provider-native and routinely contain slashes.
//
// _Requires permission: `resources:write`._
//
// DELETE /api/org/{orgId}/posture/dismissals
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *PostureDismissalsNamespace) Delete(ctx context.Context, params PostureDismissalsDeleteParams, opts ...RequestOption) error {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/posture/dismissals")
	r.setPath("orgId", params.OrgID)
	r.addQuery("resourceId", params.ResourceID)
	r.addQuery("ruleId", params.RuleID)
	return n.t.do(ctx, r, nil, opts)
}

// PostureSettingsNamespace is `client.posture.settings`.
type PostureSettingsNamespace struct {
	t *transport
}

func newPostureSettingsNamespace(t *transport) *PostureSettingsNamespace {
	n := &PostureSettingsNamespace{t: t}
	return n
}

// PostureSettingsGetParams holds the parameters for
// `client.posture.settings.get`.
//
// Every field is optional; pass nil to take the defaults.
type PostureSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's posture alert settings
//
// Whether the poller's daily posture alert scan is enabled. An organization that
// never saved reads the shipped defaults (enabled).
//
// _Requires permission: `org:settings:write`._
//
// GET /api/org/{orgId}/posture/settings
func (n *PostureSettingsNamespace) Get(ctx context.Context, params *PostureSettingsGetParams, opts ...RequestOption) (*PostureAlertSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/posture/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *PostureAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// PostureSettingsUpdateParams holds the parameters for
// `client.posture.settings.update`.
//
// Every field is optional; pass nil to take the defaults.
type PostureSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *PostureAlertSettingsUpdate
}

// Update: Update the posture alert settings
//
// Saving never resets the alert cooldown.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/posture/settings
//
// Raises on 400: Bad request
func (n *PostureSettingsNamespace) Update(ctx context.Context, params *PostureSettingsUpdateParams, opts ...RequestOption) (*PostureAlertSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/posture/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *PostureAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProbesNamespace is `client.probes`.
type ProbesNamespace struct {
	t *transport
}

func newProbesNamespace(t *transport) *ProbesNamespace {
	n := &ProbesNamespace{t: t}
	return n
}

// ProbesCreateParams holds the parameters for `client.probes.create`.
//
// Every field is optional; pass nil to take the defaults.
type ProbesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *SyntheticProbeCreate
}

// Create: Create a probe
//
// Point an uptime/latency check at an endpoint. Numeric inputs are clamped into
// their allowed ranges rather than rejected; the first check runs within one
// poller tick. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/probes
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ProbesNamespace) Create(ctx context.Context, params *ProbesCreateParams, opts ...RequestOption) (*SyntheticProbe, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/probes")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *SyntheticProbe
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProbesDeleteParams holds the parameters for `client.probes.delete`.
type ProbesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	ProbeID string
}

// Delete: Delete a probe
//
// Remove the probe. Recorded series age out of the metric store. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// DELETE /api/org/{orgId}/probes/{probeId}
//
// Raises on 404: Not found
func (n *ProbesNamespace) Delete(ctx context.Context, params ProbesDeleteParams, opts ...RequestOption) error {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/probes/{probeId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("probeId", params.ProbeID)
	return n.t.do(ctx, r, nil, opts)
}

// ProbesGetParams holds the parameters for `client.probes.get`.
//
// Every field is optional; pass nil to take the defaults.
type ProbesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List synthetic probes
//
// Every probe in the organization with its live status, consecutive-failure
// count, last latency and trailing-24h uptime. Probes run on an interval from an
// edge proxy outside the cluster, so results reflect what an internet client
// would see.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/probes
func (n *ProbesNamespace) Get(ctx context.Context, params *ProbesGetParams, opts ...RequestOption) (*SyntheticProbeList, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/probes")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SyntheticProbeList
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProbesMetricsParams holds the parameters for `client.probes.metrics`.
type ProbesMetricsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	ProbeID string
	// StartMs: Range start, Unix epoch ms.
	StartMs *string
	// EndMs: Range end, Unix epoch ms.
	EndMs *string
}

// Metrics: Read a probe's recorded series
//
// The "Latency" (ms) and "Up" (1/0) series over a time range, from the shared
// metric store. Resolution auto-selects raw/1-minute/1-hour rollups by span.
// Defaults to the trailing 24 hours.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/probes/{probeId}/metrics
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 503: A backing service this endpoint depends on is not available
func (n *ProbesNamespace) Metrics(ctx context.Context, params ProbesMetricsParams, opts ...RequestOption) (*ProbeMetrics, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/probes/{probeId}/metrics")
	r.setPath("orgId", params.OrgID)
	r.setPath("probeId", params.ProbeID)
	r.addQuery("startMs", params.StartMs)
	r.addQuery("endMs", params.EndMs)
	var out *ProbeMetrics
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProbesSuggestionsParams holds the parameters for `client.probes.suggestions`.
//
// Every field is optional; pass nil to take the defaults.
type ProbesSuggestionsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Suggestions: Suggest endpoints from synced resources
//
// Endpoint candidates mined from the organization's synced resource outputs and
// fields (keys like url, endpoint, host, domain, publicIp). A cheap read over
// stored state — no provider API calls. Deduplicated by URL.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/probes/suggestions
func (n *ProbesNamespace) Suggestions(ctx context.Context, params *ProbesSuggestionsParams, opts ...RequestOption) (*ProbeSuggestions, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/probes/suggestions")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *ProbeSuggestions
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProbesUpdateParams holds the parameters for `client.probes.update`.
type ProbesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID   *string
	ProbeID string
	// Body: the JSON request body.
	Body *SyntheticProbeUpdate
}

// Update: Update or disable a probe
//
// Edit settings and/or toggle `enabled`. Changing the URL or method resets the
// probe's state to `unknown` — the history belongs to the old endpoint.
// Audit-logged.
//
// _Requires permission: `resources:write`._
//
// PUT /api/org/{orgId}/probes/{probeId}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ProbesNamespace) Update(ctx context.Context, params ProbesUpdateParams, opts ...RequestOption) (*SyntheticProbe, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/probes/{probeId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("probeId", params.ProbeID)
	r.setJSONBody(params.Body)
	var out *SyntheticProbe
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileNamespace is `client.profile`.
type ProfileNamespace struct {
	t *transport

	// EmailChange: `client.profile.emailChange`.
	EmailChange *ProfileEmailChangeNamespace
	// MFA: `client.profile.mfa`.
	MFA *ProfileMFANamespace
	// Sessions: `client.profile.sessions`.
	Sessions *ProfileSessionsNamespace
}

func newProfileNamespace(t *transport) *ProfileNamespace {
	n := &ProfileNamespace{t: t}
	n.EmailChange = newProfileEmailChangeNamespace(t)
	n.MFA = newProfileMFANamespace(t)
	n.Sessions = newProfileSessionsNamespace(t)
	return n
}

// Delete: Delete the signed-in user's account
//
// Irreversible. Organizations where the caller is the only member are deleted
// and their subscriptions cancelled; other memberships are simply removed.
// Refuses with `transfer_ownership_required` while the caller is the only owner
// of an organization other people belong to.
//
// DELETE /api/profile
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Recent sign-in required. Send the user through sign-in again
// and retry; the request itself was well-formed.
//
// Raises on 409: The caller still solely owns a shared organization; nothing was
// deleted.
//
// Raises on 502: A subscription could not be cancelled; nothing was deleted.
func (n *ProfileNamespace) Delete(ctx context.Context, opts ...RequestOption) (*AccountDeleted, error) {
	r := newRequest(http.MethodDelete, "/api/profile")
	var out *AccountDeleted
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// DeletionPreview: What deleting this account would do
//
// Read-only. Lets a confirmation screen name the organizations that go with the
// account, and the ones that must be handed over first.
//
// GET /api/profile/deletion-preview
//
// Raises on 401: Unauthenticated
func (n *ProfileNamespace) DeletionPreview(ctx context.Context, opts ...RequestOption) (*AccountDeletionPreview, error) {
	r := newRequest(http.MethodGet, "/api/profile/deletion-preview")
	var out *AccountDeletionPreview
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// Get: The signed-in user's account profile
//
// User-scoped, not organization-scoped: one WorkOS identity is shared across
// every organization the user belongs to.
//
// GET /api/profile
//
// Raises on 401: Unauthenticated
func (n *ProfileNamespace) Get(ctx context.Context, opts ...RequestOption) (*Profile, error) {
	r := newRequest(http.MethodGet, "/api/profile")
	var out *Profile
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// PasswordReset: Mint a password reset link for the signed-in user
//
// Returns a one-time AuthKit-hosted reset URL rather than emailing it — the
// caller already holds a valid session for the account. Also the way to set a
// first password on an SSO or OAuth-only account.
//
// POST /api/profile/password-reset
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Recent sign-in required. Send the user through sign-in again
// and retry; the request itself was well-formed.
func (n *ProfileNamespace) PasswordReset(ctx context.Context, opts ...RequestOption) (*ProfilePasswordResetResponse, error) {
	r := newRequest(http.MethodPost, "/api/profile/password-reset")
	var out *ProfilePasswordResetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SendVerificationEmail: Re-send the email verification message
//
// POST /api/profile/send-verification-email
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
func (n *ProfileNamespace) SendVerificationEmail(ctx context.Context, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/profile/send-verification-email")
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileUpdateParams holds the parameters for `client.profile.update`.
//
// Every field is optional; pass nil to take the defaults.
type ProfileUpdateParams struct {
	// Body: the JSON request body.
	Body *ProfileUpdateRequest
}

// Update: Update the signed-in user's name
//
// PATCH /api/profile
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
func (n *ProfileNamespace) Update(ctx context.Context, params *ProfileUpdateParams, opts ...RequestOption) (*ProfileSummary, error) {
	r := newRequest(http.MethodPatch, "/api/profile")
	if params != nil {
		r.setJSONBody(params.Body)
	}
	var out *ProfileSummary
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileEmailChangeNamespace is `client.profile.emailChange`.
type ProfileEmailChangeNamespace struct {
	t *transport
}

func newProfileEmailChangeNamespace(t *transport) *ProfileEmailChangeNamespace {
	n := &ProfileEmailChangeNamespace{t: t}
	return n
}

// ProfileEmailChangeConfirmParams holds the parameters for
// `client.profile.emailChange.confirm`.
//
// Every field is optional; pass nil to take the defaults.
type ProfileEmailChangeConfirmParams struct {
	// Body: the JSON request body.
	Body *ProfileEmailChangeConfirmRequest
}

// Confirm: Redeem an email change code
//
// On success the account's email is the new address and it is marked verified.
//
// POST /api/profile/email-change/confirm
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Recent sign-in required. Send the user through sign-in again
// and retry; the request itself was well-formed.
func (n *ProfileEmailChangeNamespace) Confirm(ctx context.Context, params *ProfileEmailChangeConfirmParams, opts ...RequestOption) (*ProfileEmailChangeConfirmResponse, error) {
	r := newRequest(http.MethodPost, "/api/profile/email-change/confirm")
	if params != nil {
		r.setJSONBody(params.Body)
	}
	var out *ProfileEmailChangeConfirmResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileEmailChangeCreateParams holds the parameters for
// `client.profile.emailChange.create`.
//
// Every field is optional; pass nil to take the defaults.
type ProfileEmailChangeCreateParams struct {
	// Body: the JSON request body.
	Body *ProfileEmailChangeCreateRequest
}

// Create: Send a confirmation code to a new email address
//
// Starts an email change. The code goes to the new address and the account keeps
// its current address until `/api/profile/email-change/confirm` redeems it, so
// an abandoned or mistyped change is harmless.
//
// POST /api/profile/email-change
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Recent sign-in required. Send the user through sign-in again
// and retry; the request itself was well-formed.
func (n *ProfileEmailChangeNamespace) Create(ctx context.Context, params *ProfileEmailChangeCreateParams, opts ...RequestOption) (*ProfileEmailChangeCreateResponse, error) {
	r := newRequest(http.MethodPost, "/api/profile/email-change")
	if params != nil {
		r.setJSONBody(params.Body)
	}
	var out *ProfileEmailChangeCreateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileMFANamespace is `client.profile.mfa`.
type ProfileMFANamespace struct {
	t *transport
}

func newProfileMFANamespace(t *transport) *ProfileMFANamespace {
	n := &ProfileMFANamespace{t: t}
	return n
}

// ProfileMFAChallengeParams holds the parameters for
// `client.profile.mfa.challenge`.
type ProfileMFAChallengeParams struct {
	FactorID string
}

// Challenge: Issue a fresh challenge for a factor
//
// POST /api/profile/mfa/{factorId}/challenge
//
// Raises on 401: Unauthenticated
//
// Raises on 404: Not found
func (n *ProfileMFANamespace) Challenge(ctx context.Context, params ProfileMFAChallengeParams, opts ...RequestOption) (*ProfileMfachallengeResponse, error) {
	r := newRequest(http.MethodPost, "/api/profile/mfa/{factorId}/challenge")
	r.setPath("factorId", params.FactorID)
	var out *ProfileMfachallengeResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// Create: Begin TOTP enrolment
//
// Creates the factor and a first challenge. The factor only becomes usable once
// a code is verified; abandon the flow by DELETEing the returned `factorId`.
//
// POST /api/profile/mfa
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Recent sign-in required. Send the user through sign-in again
// and retry; the request itself was well-formed.
func (n *ProfileMFANamespace) Create(ctx context.Context, opts ...RequestOption) (*TOTPEnrollment, error) {
	r := newRequest(http.MethodPost, "/api/profile/mfa")
	var out *TOTPEnrollment
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileMFADeleteParams holds the parameters for `client.profile.mfa.delete`.
type ProfileMFADeleteParams struct {
	FactorID string
}

// Delete: Remove an authentication factor
//
// DELETE /api/profile/mfa/{factorId}
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Recent sign-in required. Send the user through sign-in again
// and retry; the request itself was well-formed.
//
// Raises on 404: Not found
func (n *ProfileMFANamespace) Delete(ctx context.Context, params ProfileMFADeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/profile/mfa/{factorId}")
	r.setPath("factorId", params.FactorID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// List: List enrolled authentication factors
//
// Includes factors whose enrolment was never confirmed — WorkOS does not expose
// a verified flag.
//
// GET /api/profile/mfa
//
// Raises on 401: Unauthenticated
func (n *ProfileMFANamespace) List(ctx context.Context, opts ...RequestOption) ([]AuthFactor, error) {
	r := newRequest(http.MethodGet, "/api/profile/mfa")
	var out []AuthFactor
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileMFAVerifyParams holds the parameters for `client.profile.mfa.verify`.
type ProfileMFAVerifyParams struct {
	FactorID string
	// Body: the JSON request body.
	Body *ProfileMfaverifyRequest
}

// Verify: Verify a code against a challenge
//
// POST /api/profile/mfa/{factorId}/verify
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 404: Not found
func (n *ProfileMFANamespace) Verify(ctx context.Context, params ProfileMFAVerifyParams, opts ...RequestOption) (*ProfileMfaverifyResponse, error) {
	r := newRequest(http.MethodPost, "/api/profile/mfa/{factorId}/verify")
	r.setPath("factorId", params.FactorID)
	r.setJSONBody(params.Body)
	var out *ProfileMfaverifyResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ProfileSessionsNamespace is `client.profile.sessions`.
type ProfileSessionsNamespace struct {
	t *transport
}

func newProfileSessionsNamespace(t *transport) *ProfileSessionsNamespace {
	n := &ProfileSessionsNamespace{t: t}
	return n
}

// ProfileSessionsDeleteParams holds the parameters for
// `client.profile.sessions.delete`.
type ProfileSessionsDeleteParams struct {
	SessionID string
}

// Delete: Revoke one session
//
// Refuses the session making the request — use sign-out for that.
//
// DELETE /api/profile/sessions/{sessionId}
//
// Raises on 400: Bad request
//
// Raises on 401: Unauthenticated
//
// Raises on 404: Not found
func (n *ProfileSessionsNamespace) Delete(ctx context.Context, params ProfileSessionsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/profile/sessions/{sessionId}")
	r.setPath("sessionId", params.SessionID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// List: List the signed-in user's active sessions
//
// GET /api/profile/sessions
//
// Raises on 401: Unauthenticated
func (n *ProfileSessionsNamespace) List(ctx context.Context, opts ...RequestOption) ([]UserSession, error) {
	r := newRequest(http.MethodGet, "/api/profile/sessions")
	var out []UserSession
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// RevokeOthers: Revoke every session except the current one
//
// POST /api/profile/sessions/revoke-others
//
// Raises on 401: Unauthenticated
//
// Raises on 403: Recent sign-in required. Send the user through sign-in again
// and retry; the request itself was well-formed.
func (n *ProfileSessionsNamespace) RevokeOthers(ctx context.Context, opts ...RequestOption) (*ProfileSessionsRevokeOthersResponse, error) {
	r := newRequest(http.MethodPost, "/api/profile/sessions/revoke-others")
	var out *ProfileSessionsRevokeOthersResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// QuotasNamespace is `client.quotas`.
type QuotasNamespace struct {
	t *transport

	// Settings: `client.quotas.settings`.
	Settings *QuotasSettingsNamespace
}

func newQuotasNamespace(t *transport) *QuotasNamespace {
	n := &QuotasNamespace{t: t}
	n.Settings = newQuotasSettingsNamespace(t)
	return n
}

// QuotasGetParams holds the parameters for `client.quotas.get`.
//
// Every field is optional; pass nil to take the defaults.
type QuotasGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List provider quota utilisation across the organization
//
// How close each account is to the limits its provider enforces, with the trend
// fitted over the last 14 days of collected readings. Both halves of every row —
// the used figure and the limit — come from the provider; nothing is filled in
// from published defaults, so an account with an approved increase reads as
// having the headroom it has. This is a read over already-collected snapshots:
// no provider API calls are made here, and the readings are as fresh as the last
// collection pass (roughly six hours). A plugin that declares no quota
// capability contributes nothing rather than zero — see `unsupportedPluginIds`.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/quotas
func (n *QuotasNamespace) Get(ctx context.Context, params *QuotasGetParams, opts ...RequestOption) (*QuotaListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/quotas")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *QuotaListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// QuotasSettingsNamespace is `client.quotas.settings`.
type QuotasSettingsNamespace struct {
	t *transport
}

func newQuotasSettingsNamespace(t *transport) *QuotasSettingsNamespace {
	n := &QuotasSettingsNamespace{t: t}
	return n
}

// QuotasSettingsGetParams holds the parameters for `client.quotas.settings.get`.
//
// Every field is optional; pass nil to take the defaults.
type QuotasSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the organization's quota alert settings
//
// The threshold feeds both the feed's severity buckets and the poller's daily
// alert scan. An organization that never saved reads the shipped defaults
// (enabled, 0.8).
//
// _Requires permission: `org:settings:write`._
//
// GET /api/org/{orgId}/quotas/settings
func (n *QuotasSettingsNamespace) Get(ctx context.Context, params *QuotasSettingsGetParams, opts ...RequestOption) (*QuotaAlertSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/quotas/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *QuotaAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// QuotasSettingsUpdateParams holds the parameters for
// `client.quotas.settings.update`.
//
// Every field is optional; pass nil to take the defaults.
type QuotasSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *QuotaAlertSettingsUpdate
}

// Update: Update the quota alert settings
//
// Every field is optional so a single toggle can be saved on its own.
// `threshold` is a fraction from 0.5 to 0.99 and is rejected rather than clamped
// when out of range. Saving never resets the alert cooldown.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/quotas/settings
//
// Raises on 400: Bad request
func (n *QuotasSettingsNamespace) Update(ctx context.Context, params *QuotasSettingsUpdateParams, opts ...RequestOption) (*QuotaAlertSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/quotas/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *QuotaAlertSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesNamespace is `client.resources`.
type ResourcesNamespace struct {
	t *transport

	// Manifest: `client.resources.manifest`.
	Manifest *ResourcesManifestNamespace
	// SecretVersions: `client.resources.secretVersions`.
	SecretVersions *ResourcesSecretVersionsNamespace
}

func newResourcesNamespace(t *transport) *ResourcesNamespace {
	n := &ResourcesNamespace{t: t}
	n.Manifest = newResourcesManifestNamespace(t)
	n.SecretVersions = newResourcesSecretVersionsNamespace(t)
	return n
}

// ResourcesAttachParams holds the parameters for `client.resources.attach`.
type ResourcesAttachParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body AttachRequest
}

// Attach: Attach a resource onto another (e.g. disk → VM)
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/resources/attach
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) Attach(ctx context.Context, params ResourcesAttachParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/attach")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesCostEstimateParams holds the parameters for
// `client.resources.costEstimate`.
type ResourcesCostEstimateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CostEstimateRequest
}

// CostEstimate: Estimated monthly cost of a configuration
//
// Calls the plugin's `estimateCost` and returns a monthly total with the line
// items behind it. Price a proposed resource by passing `fields`, an existing
// one by passing `resourceId`, or a proposed change to an existing one by
// passing both — `fields` is merged over the resource's stored fields, so the
// caller only sends what changed. `estimate` is null when the plugin cannot
// price the configuration; that is not the same as an estimate of zero, and it
// should not be rendered as one.
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/cost-estimate
//
// Raises on 404: Not found
func (n *ResourcesNamespace) CostEstimate(ctx context.Context, params ResourcesCostEstimateParams, opts ...RequestOption) (*ResourcesCostEstimateResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/cost-estimate")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ResourcesCostEstimateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesCreateParams holds the parameters for `client.resources.create`.
type ResourcesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateResourceRequest
}

// Create: Create a new resource via its plugin
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/resources/create
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 422: Blocked by the organization's tag policy: the submitted fields
// are missing a required tag (or carry a disallowed value). Retry with the
// `x-tag-policy-override: true` header if you hold `tag-policy:override`; both
// blocks and overrides are audit-logged.
func (n *ResourcesNamespace) Create(ctx context.Context, params ResourcesCreateParams, opts ...RequestOption) (*CreateResourceResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/create")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *CreateResourceResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesCreateConfigParams holds the parameters for
// `client.resources.createConfig`.
type ResourcesCreateConfigParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateConfigRequest
}

// CreateConfig: Get the dynamic create form for a resource type
//
// Calls the plugin's `getCreateConfig`. The returned `CreateResourceConfig` is
// plugin-shaped — see `JsonObject`.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/resources/create-config
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) CreateConfig(ctx context.Context, params ResourcesCreateConfigParams, opts ...RequestOption) (JSONObject, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/create-config")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out JSONObject
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesCreatePricingParams holds the parameters for
// `client.resources.createPricing`.
type ResourcesCreatePricingParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreatePricingRequest
}

// CreatePricing: Pricing per size for a create form
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/create-pricing
func (n *ResourcesNamespace) CreatePricing(ctx context.Context, params ResourcesCreatePricingParams, opts ...RequestOption) (map[string]JSONObject, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/create-pricing")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out map[string]JSONObject
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesDeleteParams holds the parameters for `client.resources.delete`.
type ResourcesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID            *string
	PluginID         PluginID
	TypeID           ResourceTypeID
	ResourceID       ResourceID
	AccountID        string
	ParentResourceID *ResourceID
}

// Delete: Delete a resource via the plugin
//
// _Requires permission: `resources:delete`._
//
// DELETE /api/org/{orgId}/resources/{pluginId}/{typeId}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 423: Blocked by an active change freeze. Retry with the
// `x-change-freeze-override: true` header if you hold `freezes:override`; both
// blocks and overrides are audit-logged.
func (n *ResourcesNamespace) Delete(ctx context.Context, params ResourcesDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/resources/{pluginId}/{typeId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.addQuery("resourceId", params.ResourceID)
	r.addQuery("accountId", params.AccountID)
	r.addQuery("parentResourceId", params.ParentResourceID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesDescribeParams holds the parameters for `client.resources.describe`.
type ResourcesDescribeParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body DescribeRequest
}

// Describe: Get human-readable describe text for a resource
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/describe
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) Describe(ctx context.Context, params ResourcesDescribeParams, opts ...RequestOption) (*DescribeResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/describe")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *DescribeResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesDetailParams holds the parameters for `client.resources.detail`.
type ResourcesDetailParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID            *string
	PluginID         PluginID
	TypeID           ResourceTypeID
	ResourceID       ResourceID
	AccountID        *string
	ParentResourceID *ResourceID
	// IncludePeerPanes: Default true. If false, peer panes are returned as
	// stubs.
	//
	// One of "true", "false".
	IncludePeerPanes *string
}

// Detail: Full resource detail page payload
//
// Performs a live `listResources` against the provider, falls back to DB on
// failure, and returns the plugin's `renderDetail` schema plus host-derived
// flags (SQL/KV/SSH availability, child resources, peer panes, etc).
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/resources/{pluginId}/{typeId}/detail
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) Detail(ctx context.Context, params ResourcesDetailParams, opts ...RequestOption) (*ResourceDetail, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/resources/{pluginId}/{typeId}/detail")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.addQuery("resourceId", params.ResourceID)
	r.addQuery("accountId", params.AccountID)
	r.addQuery("parentResourceId", params.ParentResourceID)
	r.addQuery("includePeerPanes", params.IncludePeerPanes)
	var out *ResourceDetail
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesExportCredentialParams holds the parameters for
// `client.resources.exportCredential`.
type ResourcesExportCredentialParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body ExportCredentialRequest
}

// ExportCredential: Export a credential file for a resource (one-time reveal)
//
// _Requires permission: `secrets:read`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/export-credential
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) ExportCredential(ctx context.Context, params ResourcesExportCredentialParams, opts ...RequestOption) (*CredentialExport, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/export-credential")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *CredentialExport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesExportTerraformParams holds the parameters for
// `client.resources.exportTerraform`.
type ResourcesExportTerraformParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body ExportTerraformRequest
}

// ExportTerraform: Generate Terraform HCL for a resource (and its direct
// children) from stored state
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/export-terraform
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) ExportTerraform(ctx context.Context, params ResourcesExportTerraformParams, opts ...RequestOption) (*TerraformExport, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/export-terraform")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *TerraformExport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesFieldActionParams holds the parameters for
// `client.resources.fieldAction`.
type ResourcesFieldActionParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body FieldActionRequest
}

// FieldAction: Execute an in-form field action (e.g. generate an IAM role)
//
// Calls the plugin's `executeFieldAction`. Returns `{ value }` to assign to the
// field; for `select` fields the optional `option` should be spliced into the
// options list so the new value can be displayed.
//
// POST /api/org/{orgId}/resources/field-action
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) FieldAction(ctx context.Context, params ResourcesFieldActionParams, opts ...RequestOption) (*FieldActionResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/field-action")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *FieldActionResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesImportYAMLParams holds the parameters for
// `client.resources.importYaml`.
type ResourcesImportYAMLParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	// Body: the JSON request body.
	Body ImportYAMLRequest
}

// ImportYAML: Bulk-import resources from YAML (kubectl apply -f equivalent)
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/resources/{pluginId}/import-yaml
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) ImportYAML(ctx context.Context, params ResourcesImportYAMLParams, opts ...RequestOption) (JSONObject, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/import-yaml")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setJSONBody(params.Body)
	var out JSONObject
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesInvokeActionParams holds the parameters for
// `client.resources.invokeAction`.
type ResourcesInvokeActionParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body InvokeActionRequest
}

// InvokeAction: Invoke a plugin-defined action on a resource
//
// Actions the plugin marks `destructive: true` in its detail schema are blocked
// with `423` while an org change freeze is in effect.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/resources/invoke-action
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 423: Blocked by an active change freeze. Retry with the
// `x-change-freeze-override: true` header if you hold `freezes:override`; both
// blocks and overrides are audit-logged.
func (n *ResourcesNamespace) InvokeAction(ctx context.Context, params ResourcesInvokeActionParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/invoke-action")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesLogsParams holds the parameters for `client.resources.logs`.
type ResourcesLogsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body LogsRequest
}

// Logs: Fetch logs for a resource
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/logs
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) Logs(ctx context.Context, params ResourcesLogsParams, opts ...RequestOption) (*LogsResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/logs")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *LogsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesMetricsParams holds the parameters for `client.resources.metrics`.
type ResourcesMetricsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body MetricsRequest
}

// Metrics: Fetch metric series for a resource
//
// Historical points from the metrics store when the resource has accumulated any
// (resources pinned to a dashboard are polled continuously); otherwise the
// series are fetched live from the provider on demand.
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/metrics
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) Metrics(ctx context.Context, params ResourcesMetricsParams, opts ...RequestOption) (*MetricsResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/metrics")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *MetricsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesNoSQLCommandParams holds the parameters for
// `client.resources.nosqlCommand`.
type ResourcesNoSQLCommandParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body NoSQLCommandRequest
}

// NoSQLCommand: Run a NoSQL document-browser command (e.g. MongoDB shell)
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/resources/nosql-command
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesNamespace) NoSQLCommand(ctx context.Context, params ResourcesNoSQLCommandParams, opts ...RequestOption) (*ResourcesNoSqlcommandResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/nosql-command")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ResourcesNoSqlcommandResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesPeerPanesParams holds the parameters for
// `client.resources.peerPanes`.
type ResourcesPeerPanesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body PeerPanesRequest
}

// PeerPanes: Lazy-fetch peer-integration panes for a resource
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/peer-panes
//
// Raises on 404: Not found
func (n *ResourcesNamespace) PeerPanes(ctx context.Context, params ResourcesPeerPanesParams, opts ...RequestOption) ([]PeerPane, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/peer-panes")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out []PeerPane
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesPickerResourcesParams holds the parameters for
// `client.resources.pickerResources`.
type ResourcesPickerResourcesParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body PickerResourcesRequest
}

// PickerResources: Fetch options for a `resource-picker` field
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/picker-resources
func (n *ResourcesNamespace) PickerResources(ctx context.Context, params ResourcesPickerResourcesParams, opts ...RequestOption) ([]PickerResource, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/picker-resources")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out []PickerResource
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesUpdateParams holds the parameters for `client.resources.update`.
type ResourcesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body UpdateResourceRequest
}

// Update: Update a resource via its plugin
//
// Applies the supplied field changes upstream and persists the refreshed
// fields/display name to the DB. The body's `fields` map only carries the keys
// the caller actually changed. Blocked with `423` while an org change freeze is
// in effect (this is also the path that applies right-sizing recommendations);
// every applied update is audit-logged.
//
// POST /api/org/{orgId}/resources/update
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 423: Blocked by an active change freeze. Retry with the
// `x-change-freeze-override: true` header if you hold `freezes:override`; both
// blocks and overrides are audit-logged.
func (n *ResourcesNamespace) Update(ctx context.Context, params ResourcesUpdateParams, opts ...RequestOption) (*UpdateResourceResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/update")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *UpdateResourceResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesManifestNamespace is `client.resources.manifest`.
type ResourcesManifestNamespace struct {
	t *transport
}

func newResourcesManifestNamespace(t *transport) *ResourcesManifestNamespace {
	n := &ResourcesManifestNamespace{t: t}
	return n
}

// ResourcesManifestCreateParams holds the parameters for
// `client.resources.manifest.create`.
type ResourcesManifestCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body ApplyManifestRequest
}

// Create: Apply an edited manifest to a resource
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/manifest
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesManifestNamespace) Create(ctx context.Context, params ResourcesManifestCreateParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/manifest")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesManifestGetParams holds the parameters for
// `client.resources.manifest.get`.
type ResourcesManifestGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID            *string
	PluginID         PluginID
	TypeID           ResourceTypeID
	ResourceID       ResourceID
	AccountID        string
	ParentResourceID *ResourceID
}

// Get: Fetch the raw manifest (YAML/JSON) for a resource
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/resources/{pluginId}/{typeId}/manifest
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesManifestNamespace) Get(ctx context.Context, params ResourcesManifestGetParams, opts ...RequestOption) (*Manifest, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/resources/{pluginId}/{typeId}/manifest")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.addQuery("resourceId", params.ResourceID)
	r.addQuery("accountId", params.AccountID)
	r.addQuery("parentResourceId", params.ParentResourceID)
	var out *Manifest
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesSecretVersionsNamespace is `client.resources.secretVersions`.
type ResourcesSecretVersionsNamespace struct {
	t *transport
}

func newResourcesSecretVersionsNamespace(t *transport) *ResourcesSecretVersionsNamespace {
	n := &ResourcesSecretVersionsNamespace{t: t}
	return n
}

// ResourcesSecretVersionsAccessParams holds the parameters for
// `client.resources.secretVersions.access`.
type ResourcesSecretVersionsAccessParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body SecretAccessRequest
}

// Access: Reveal the plaintext value of a specific version (one-time)
//
// _Requires permission: `secrets:read`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions/access
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesSecretVersionsNamespace) Access(ctx context.Context, params ResourcesSecretVersionsAccessParams, opts ...RequestOption) (*SecretAccessResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions/access")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *SecretAccessResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesSecretVersionsAddParams holds the parameters for
// `client.resources.secretVersions.add`.
type ResourcesSecretVersionsAddParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body SecretAddRequest
}

// Add: Add a new secret version
//
// _Requires permission: `secrets:write`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions/add
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesSecretVersionsNamespace) Add(ctx context.Context, params ResourcesSecretVersionsAddParams, opts ...RequestOption) (*SecretVersionResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions/add")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *SecretVersionResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesSecretVersionsGetParams holds the parameters for
// `client.resources.secretVersions.get`.
type ResourcesSecretVersionsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID            *string
	PluginID         PluginID
	TypeID           ResourceTypeID
	ResourceID       ResourceID
	AccountID        string
	ParentResourceID *ResourceID
}

// Get: List secret versions for a versioned-secret resource
//
// _Requires permission: `secrets:read`._
//
// GET /api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *ResourcesSecretVersionsNamespace) Get(ctx context.Context, params ResourcesSecretVersionsGetParams, opts ...RequestOption) (*SecretVersionsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.addQuery("resourceId", params.ResourceID)
	r.addQuery("accountId", params.AccountID)
	r.addQuery("parentResourceId", params.ParentResourceID)
	var out *SecretVersionsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// ResourcesSecretVersionsModifyParams holds the parameters for
// `client.resources.secretVersions.modify`.
type ResourcesSecretVersionsModifyParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID    *string
	PluginID PluginID
	TypeID   ResourceTypeID
	// Body: the JSON request body.
	Body SecretModifyRequest
}

// Modify: Enable/disable/destroy a secret version
//
// _Requires permission: `secrets:write`._
//
// POST /api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions/modify
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 423: Blocked by an active change freeze. Retry with the
// `x-change-freeze-override: true` header if you hold `freezes:override`; both
// blocks and overrides are audit-logged.
func (n *ResourcesSecretVersionsNamespace) Modify(ctx context.Context, params ResourcesSecretVersionsModifyParams, opts ...RequestOption) (*SecretVersionResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/{pluginId}/{typeId}/secret-versions/modify")
	r.setPath("orgId", params.OrgID)
	r.setPath("pluginId", params.PluginID)
	r.setPath("typeId", params.TypeID)
	r.setJSONBody(params.Body)
	var out *SecretVersionResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// RightsizingNamespace is `client.rightsizing`.
type RightsizingNamespace struct {
	t *transport
}

func newRightsizingNamespace(t *transport) *RightsizingNamespace {
	n := &RightsizingNamespace{t: t}
	return n
}

// RightsizingGetParams holds the parameters for `client.rightsizing.get`.
//
// Every field is optional; pass nil to take the defaults.
type RightsizingGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Refresh: Bypass the short server-side cache and recompute now.
	//
	// One of "true", "false".
	Refresh *string
}

// Get: List oversized resources with resize recommendations
//
// Computes p95 CPU/memory utilisation over the last 14 days of stored metrics
// for every resource whose plugin declares right-sizing support, and matches
// under-utilised ones against the plugin's real size catalog (the create form's
// size options, live-priced). Each recommendation names the cheapest smaller
// size that still clears a headroom margin and quotes the monthly saving. Apply
// one by submitting `sizeFieldKey` with the recommended size id through the
// resource-update endpoint — which enforces change freezes and writes the audit
// trail. Results are cached for a few minutes; pass `refresh=true` to recompute.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/rightsizing
func (n *RightsizingNamespace) Get(ctx context.Context, params *RightsizingGetParams, opts ...RequestOption) (*RightsizingListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/rightsizing")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("refresh", params.Refresh)
	}
	var out *RightsizingListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SavedCostFiltersNamespace is `client.savedCostFilters`.
type SavedCostFiltersNamespace struct {
	t *transport
}

func newSavedCostFiltersNamespace(t *transport) *SavedCostFiltersNamespace {
	n := &SavedCostFiltersNamespace{t: t}
	return n
}

// SavedCostFiltersCreateParams holds the parameters for
// `client.savedCostFilters.create`.
type SavedCostFiltersCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SavedCostFilterInput
}

// Create: Create a saved cost filter
//
// Names must be unique per organization (case-insensitively) — they are how the
// CLI's `--filter <name>` and humans address the filter. A name collision is a 409.
//
// _Requires permission: `costs:write`._
//
// POST /api/org/{orgId}/saved-cost-filters
//
// Raises on 400: Bad request
//
// Raises on 409: A live saved filter already uses this name.
func (n *SavedCostFiltersNamespace) Create(ctx context.Context, params SavedCostFiltersCreateParams, opts ...RequestOption) (*SavedCostFilter, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/saved-cost-filters")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *SavedCostFilter
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SavedCostFiltersDeleteParams holds the parameters for
// `client.savedCostFilters.delete`.
type SavedCostFiltersDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a saved cost filter
//
// Soft delete — **refused with a 409 while anything references the filter**,
// with the referents in the body. Deleting a referenced filter would silently
// widen every referent's scope to all spend; for a budget that can fire or
// suppress alerts, so detaching the referents is a deliberate step, never a side
// effect of deletion.
//
// _Requires permission: `costs:write`._
//
// DELETE /api/org/{orgId}/saved-cost-filters/{id}
//
// Raises on 404: Not found
//
// Raises on 409: Still referenced — the body lists every referent.
func (n *SavedCostFiltersNamespace) Delete(ctx context.Context, params SavedCostFiltersDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/saved-cost-filters/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SavedCostFiltersGetParams holds the parameters for
// `client.savedCostFilters.get`.
type SavedCostFiltersGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Get: Get a saved cost filter
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/saved-cost-filters/{id}
//
// Raises on 404: Not found
func (n *SavedCostFiltersNamespace) Get(ctx context.Context, params SavedCostFiltersGetParams, opts ...RequestOption) (*SavedCostFilter, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/saved-cost-filters/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *SavedCostFilter
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SavedCostFiltersListParams holds the parameters for
// `client.savedCostFilters.list`.
//
// Every field is optional; pass nil to take the defaults.
type SavedCostFiltersListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List saved cost filters
//
// Named, reusable cost filter sets. Graphs, reports and budgets reference one
// **by id** (`savedFilterId` in their configs and in `POST /costs/query`), and
// the server resolves the reference at query time — so editing a saved filter
// changes every referent at once, and nothing ever holds a copy.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/saved-cost-filters
func (n *SavedCostFiltersNamespace) List(ctx context.Context, params *SavedCostFiltersListParams, opts ...RequestOption) ([]SavedCostFilter, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/saved-cost-filters")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []SavedCostFilter
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SavedCostFiltersReferentsParams holds the parameters for
// `client.savedCostFilters.referents`.
type SavedCostFiltersReferentsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Referents: List a saved filter's referents
//
// Every budget, cost report and dashboard cost graph referencing this filter —
// what an edit will re-scope, and what a delete would be refused over.
//
// _Requires permission: `costs:read`._
//
// GET /api/org/{orgId}/saved-cost-filters/{id}/referents
//
// Raises on 404: Not found
func (n *SavedCostFiltersNamespace) Referents(ctx context.Context, params SavedCostFiltersReferentsParams, opts ...RequestOption) (*SavedCostFiltersReferentsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/saved-cost-filters/{id}/referents")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *SavedCostFiltersReferentsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SavedCostFiltersUpdateParams holds the parameters for
// `client.savedCostFilters.update`.
type SavedCostFiltersUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body SavedCostFilterInput
}

// Update: Update a saved cost filter
//
// Replaces the filter's name, description and terms. This is the high-leverage
// write: every graph, report and budget referencing the filter runs the new
// terms on its next query — re-scoping a referenced budget can change which
// alerts fire. `GET /{id}/referents` names what a change will touch.
//
// _Requires permission: `costs:write`._
//
// PUT /api/org/{orgId}/saved-cost-filters/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: A live saved filter already uses this name.
func (n *SavedCostFiltersNamespace) Update(ctx context.Context, params SavedCostFiltersUpdateParams, opts ...RequestOption) (*SavedCostFilter, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/saved-cost-filters/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *SavedCostFilter
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SchedulesNamespace is `client.schedules`.
type SchedulesNamespace struct {
	t *transport
}

func newSchedulesNamespace(t *transport) *SchedulesNamespace {
	n := &SchedulesNamespace{t: t}
	return n
}

// SchedulesCreateParams holds the parameters for `client.schedules.create`.
//
// Every field is optional; pass nil to take the defaults.
type SchedulesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *SleepScheduleCreate
}

// Create: Create a sleep/wake schedule
//
// Attach an off-at/on-at weekly window to a resource. The resource's type must
// declare lifecycle start/stop actions (see the resource type metadata); one
// schedule per resource. Times are wall-clock in the given IANA timezone and
// remain correct across DST. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/schedules
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: The resource already has a schedule
func (n *SchedulesNamespace) Create(ctx context.Context, params *SchedulesCreateParams, opts ...RequestOption) (*SleepSchedule, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/schedules")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *SleepSchedule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SchedulesDeleteParams holds the parameters for `client.schedules.delete`.
type SchedulesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ScheduleID string
}

// Delete: Delete a schedule
//
// Remove the schedule. The resource is left in whatever state it is in.
// Audit-logged.
//
// _Requires permission: `resources:write`._
//
// DELETE /api/org/{orgId}/schedules/{scheduleId}
//
// Raises on 404: Not found
func (n *SchedulesNamespace) Delete(ctx context.Context, params SchedulesDeleteParams, opts ...RequestOption) error {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/schedules/{scheduleId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("scheduleId", params.ScheduleID)
	return n.t.do(ctx, r, nil, opts)
}

// SchedulesGetParams holds the parameters for `client.schedules.get`.
//
// Every field is optional; pass nil to take the defaults.
type SchedulesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List sleep/wake schedules
//
// Every schedule in the organization with its next transition, last run outcome
// and a projected monthly saving computed from trailing per-resource spend and
// the weekly off-hours fraction. Schedules attach to resources whose plugin
// declares lifecycle start/stop actions; the poller executes due transitions
// server-side.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/schedules
func (n *SchedulesNamespace) Get(ctx context.Context, params *SchedulesGetParams, opts ...RequestOption) (*SleepScheduleList, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/schedules")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SleepScheduleList
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SchedulesPreviewParams holds the parameters for `client.schedules.preview`.
//
// Every field is optional; pass nil to take the defaults.
type SchedulesPreviewParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *SleepSchedulePreviewRequest
}

// Preview: Preview a schedule's projected saving
//
// Quote a timing against a resource before saving: the weekly off-hours
// fraction, the resource's trailing spend normalized to a month, the projected
// monthly saving, and the next few transitions. Makes no provider API calls and
// changes nothing.
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/schedules/preview
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SchedulesNamespace) Preview(ctx context.Context, params *SchedulesPreviewParams, opts ...RequestOption) (*SleepSchedulePreview, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/schedules/preview")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *SleepSchedulePreview
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SchedulesUpdateParams holds the parameters for `client.schedules.update`.
type SchedulesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	ScheduleID string
	// Body: the JSON request body.
	Body *SleepScheduleUpdate
}

// Update: Update or pause a schedule
//
// Edit the timing and/or toggle `paused`. Any change recomputes the next
// transition; pausing clears it. Audit-logged.
//
// _Requires permission: `resources:write`._
//
// PUT /api/org/{orgId}/schedules/{scheduleId}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SchedulesNamespace) Update(ctx context.Context, params SchedulesUpdateParams, opts ...RequestOption) (*SleepSchedule, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/schedules/{scheduleId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("scheduleId", params.ScheduleID)
	r.setJSONBody(params.Body)
	var out *SleepSchedule
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SearchNamespace is `client.search`.
type SearchNamespace struct {
	t *transport
}

func newSearchNamespace(t *transport) *SearchNamespace {
	n := &SearchNamespace{t: t}
	return n
}

// SearchListParams holds the parameters for `client.search.list`.
//
// Every field is optional; pass nil to take the defaults.
type SearchListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	Q     *string
}

// List: Search resources (capped at 50 hits) and workflows across the org
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/search
func (n *SearchNamespace) List(ctx context.Context, params *SearchListParams, opts ...RequestOption) ([]SearchHit, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/search")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("q", params.Q)
	}
	var out []SearchHit
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SessionRecordingsNamespace is `client.sessionRecordings`.
type SessionRecordingsNamespace struct {
	t *transport

	// Settings: `client.sessionRecordings.settings`.
	Settings *SessionRecordingsSettingsNamespace
}

func newSessionRecordingsNamespace(t *transport) *SessionRecordingsNamespace {
	n := &SessionRecordingsNamespace{t: t}
	n.Settings = newSessionRecordingsSettingsNamespace(t)
	return n
}

// SessionRecordingsCastParams holds the parameters for
// `client.sessionRecordings.cast`.
type SessionRecordingsCastParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID       *string
	RecordingID string
	// Download: Force an attachment disposition.
	//
	// One of "1".
	Download *string
}

// Cast: Download a recording as an asciicast
//
// The session as an [asciicast
// v2](https://docs.asciinema.org/manual/asciicast/v2/) document: a JSON header
// line followed by one `[time, code, data]` event per line. Deliberately
// somebody else's format — the same bytes play in `asciinema play` and in the
// reference web player, so a recording is useful to an auditor who has never
// seen this product. `?download=1` returns it as an attachment. **Every fetch is
// audit-logged**, including this one: an investigator has to be able to answer
// who has watched a given tape.
//
// _Requires permission: `session-recordings:read`._
//
// GET /api/org/{orgId}/session-recordings/{recordingId}/cast
//
// Raises on 404: Not found
func (n *SessionRecordingsNamespace) Cast(ctx context.Context, params SessionRecordingsCastParams, opts ...RequestOption) (io.ReadCloser, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/session-recordings/{recordingId}/cast")
	r.setPath("orgId", params.OrgID)
	r.setPath("recordingId", params.RecordingID)
	r.addQuery("download", params.Download)
	return n.t.stream(ctx, r, opts)
}

// SessionRecordingsDeleteParams holds the parameters for
// `client.sessionRecordings.delete`.
type SessionRecordingsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID       *string
	RecordingID string
}

// Delete: Delete a recording
//
// Removes the recording and its stored chunks. Audit-logged.
//
// _Requires permission: `session-recordings:write`._
//
// DELETE /api/org/{orgId}/session-recordings/{recordingId}
//
// Raises on 404: Not found
func (n *SessionRecordingsNamespace) Delete(ctx context.Context, params SessionRecordingsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/session-recordings/{recordingId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("recordingId", params.RecordingID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SessionRecordingsGetParams holds the parameters for
// `client.sessionRecordings.get`.
type SessionRecordingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID       *string
	RecordingID string
}

// Get: Get one recording's metadata
//
// _Requires permission: `session-recordings:read`._
//
// GET /api/org/{orgId}/session-recordings/{recordingId}
//
// Raises on 404: Not found
func (n *SessionRecordingsNamespace) Get(ctx context.Context, params SessionRecordingsGetParams, opts ...RequestOption) (*SessionRecording, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/session-recordings/{recordingId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("recordingId", params.RecordingID)
	var out *SessionRecording
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SessionRecordingsListParams holds the parameters for
// `client.sessionRecordings.list`.
//
// Every field is optional; pass nil to take the defaults.
type SessionRecordingsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Status: `recording` (live), `complete` (closed cleanly), `truncated` (hit
	// the per-session capture ceiling — the tape is a genuine partial and says
	// so), or `abandoned` (the server handling the session went away before it
	// could close the row).
	//
	// One of "recording", "complete", "truncated", "abandoned".
	Status     *string
	UserID     *string
	ResourceID *string
	AccountID  *string
	// Since: Inclusive lower bound on `startedAt`.
	Since *string
	// Until: Exclusive upper bound on `startedAt`.
	Until *string
	Limit *int64
}

// List: List recorded SSH sessions
//
// Recorded sessions, newest first. Only SSH opened through the cloud is recorded
// — those sessions are already proxied by the server, so recording tees a stream
// it holds rather than requiring an agent on the host. A desktop session that
// dials a host directly never reaches the server and cannot appear here.
//
// _Requires permission: `session-recordings:read`._
//
// GET /api/org/{orgId}/session-recordings
//
// Raises on 400: Bad request
func (n *SessionRecordingsNamespace) List(ctx context.Context, params *SessionRecordingsListParams, opts ...RequestOption) ([]SessionRecording, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/session-recordings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("status", params.Status)
		r.addQuery("userId", params.UserID)
		r.addQuery("resourceId", params.ResourceID)
		r.addQuery("accountId", params.AccountID)
		r.addQuery("since", params.Since)
		r.addQuery("until", params.Until)
		r.addQuery("limit", params.Limit)
	}
	var out []SessionRecording
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SessionRecordingsSettingsNamespace is `client.sessionRecordings.settings`.
type SessionRecordingsSettingsNamespace struct {
	t *transport
}

func newSessionRecordingsSettingsNamespace(t *transport) *SessionRecordingsSettingsNamespace {
	n := &SessionRecordingsSettingsNamespace{t: t}
	return n
}

// SessionRecordingsSettingsGetParams holds the parameters for
// `client.sessionRecordings.settings.get`.
//
// Every field is optional; pass nil to take the defaults.
type SessionRecordingsSettingsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Get the recording policy
//
// The organization's recording policy plus what it currently stores. Usage rides
// along with the policy because the only question anyone asks about retention is
// what it costs.
//
// _Requires permission: `session-recordings:read`._
//
// GET /api/org/{orgId}/session-recordings/settings
func (n *SessionRecordingsSettingsNamespace) Get(ctx context.Context, params *SessionRecordingsSettingsGetParams, opts ...RequestOption) (*SessionRecordingSettings, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/session-recordings/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SessionRecordingSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SessionRecordingsSettingsUpdateParams holds the parameters for
// `client.sessionRecordings.settings.update`.
//
// Every field is optional; pass nil to take the defaults.
type SessionRecordingsSettingsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *SessionRecordingSettingsUpdate
}

// Update: Update the recording policy
//
// Partial update — omitted fields keep their current value. Recording is opt-in
// and off by default. Audit-logged with the before/after policy.
//
// _Requires permission: `session-recordings:write`._
//
// PUT /api/org/{orgId}/session-recordings/settings
//
// Raises on 400: Bad request
func (n *SessionRecordingsSettingsNamespace) Update(ctx context.Context, params *SessionRecordingsSettingsUpdateParams, opts ...RequestOption) (*SessionRecordingSettings, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/session-recordings/settings")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *SessionRecordingSettings
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SFTPNamespace is `client.sftp`.
type SFTPNamespace struct {
	t *transport
}

func newSFTPNamespace(t *transport) *SFTPNamespace {
	n := &SFTPNamespace{t: t}
	return n
}

// SFTPDeleteParams holds the parameters for `client.sftp.delete`.
type SFTPDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SFTPDeleteRequest
}

// Delete: Delete a file or directory over SFTP
//
// _Requires permission: `storage:write`._
//
// POST /api/org/{orgId}/sftp/delete
//
// Raises on 404: Not found
//
// Raises on 500: Server error
func (n *SFTPNamespace) Delete(ctx context.Context, params SFTPDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/sftp/delete")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SFTPDownloadParams holds the parameters for `client.sftp.download`.
type SFTPDownloadParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	AccountID string
	// Paths: JSON-encoded array of remote paths
	Paths       string
	BasePath    *string
	SSHKeyID    *string
	SSHHost     *string
	SSHUsername *string
}

// Download: Download one or many files via SFTP (zipped if more than one)
//
// _Requires permission: `storage:read`._
//
// GET /api/org/{orgId}/v1/sftp/download
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 500: Server error
func (n *SFTPNamespace) Download(ctx context.Context, params SFTPDownloadParams, opts ...RequestOption) (io.ReadCloser, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/v1/sftp/download")
	r.setPath("orgId", params.OrgID)
	r.addQuery("accountId", params.AccountID)
	r.addQuery("paths", params.Paths)
	r.addQuery("basePath", params.BasePath)
	r.addQuery("sshKeyId", params.SSHKeyID)
	r.addQuery("sshHost", params.SSHHost)
	r.addQuery("sshUsername", params.SSHUsername)
	return n.t.stream(ctx, r, opts)
}

// SFTPListParams holds the parameters for `client.sftp.list`.
type SFTPListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SFTPListRequest
}

// List: List a directory over SFTP
//
// _Requires permission: `storage:read`._
//
// POST /api/org/{orgId}/sftp/list
//
// Raises on 404: Not found
//
// Raises on 500: Server error
func (n *SFTPNamespace) List(ctx context.Context, params SFTPListParams, opts ...RequestOption) ([]SFTPEntry, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/sftp/list")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out []SFTPEntry
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SFTPMkdirParams holds the parameters for `client.sftp.mkdir`.
type SFTPMkdirParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SFTPPathRequest
}

// Mkdir: Create a directory over SFTP
//
// _Requires permission: `storage:write`._
//
// POST /api/org/{orgId}/sftp/mkdir
//
// Raises on 404: Not found
//
// Raises on 500: Server error
func (n *SFTPNamespace) Mkdir(ctx context.Context, params SFTPMkdirParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/sftp/mkdir")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SFTPUploadParams holds the parameters for `client.sftp.upload`.
type SFTPUploadParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: sent as `multipart/form-data`; the `io.Reader` field is the file.
	Body SFTPUploadForm
}

// Upload: Upload a file via SFTP
//
// _Requires permission: `storage:write`._
//
// POST /api/org/{orgId}/v1/sftp/upload
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SFTPNamespace) Upload(ctx context.Context, params SFTPUploadParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/v1/sftp/upload")
	r.setPath("orgId", params.OrgID)
	r.setFormBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesNamespace is `client.sharedConsoles`.
type SharedConsolesNamespace struct {
	t *transport

	// Invites: `client.sharedConsoles.invites`.
	Invites *SharedConsolesInvitesNamespace
	// Participants: `client.sharedConsoles.participants`.
	Participants *SharedConsolesParticipantsNamespace
}

func newSharedConsolesNamespace(t *transport) *SharedConsolesNamespace {
	n := &SharedConsolesNamespace{t: t}
	n.Invites = newSharedConsolesInvitesNamespace(t)
	n.Participants = newSharedConsolesParticipantsNamespace(t)
	return n
}

// SharedConsolesCreateParams holds the parameters for
// `client.sharedConsoles.create`.
//
// Every field is optional; pass nil to take the defaults.
type SharedConsolesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *CreateSharedConsole
}

// Create: Share a live SSH session
//
// Opens a share on a session you already have running and mints its first
// invite. You become the driver.
//
// Returns 409 `console_not_here` when the pty is held by a different server
// replica than the one answering this call — reopen the terminal and share
// again. Writing the share anyway would produce a link that authorises correctly
// and then finds nothing to attach to.
//
// Requires `resources:execute` — the same permission as opening the terminal.
// Closed to API keys: sharing a shell is an act a person performs.
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/shared-consoles
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *SharedConsolesNamespace) Create(ctx context.Context, params *SharedConsolesCreateParams, opts ...RequestOption) (*SharedConsoleCreated, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/shared-consoles")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *SharedConsoleCreated
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesDeleteParams holds the parameters for
// `client.sharedConsoles.delete`.
type SharedConsolesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
}

// Delete: Revoke a share
//
// Disconnects every guest and stops the fan-out. The sharer's own SSH session
// carries on — revoking a share is not killing a terminal.
//
// The sharer or a holder of `org:settings:write`. Deliberately does **not**
// require `resources:execute`: ending access must never be gated on still
// holding the access, or an owner whose role was narrowed mid-incident could not
// close the session they opened.
//
// DELETE /api/org/{orgId}/shared-consoles/{consoleId}
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *SharedConsolesNamespace) Delete(ctx context.Context, params SharedConsolesDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/shared-consoles/{consoleId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesGetParams holds the parameters for `client.sharedConsoles.get`.
type SharedConsolesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
}

// Get: Get one shared console
//
// Visible to participants and to anyone who could revoke it (the sharer, or a
// holder of `org:settings:write`). Others get 404 — that a named colleague has a
// root shell open on a named production host right now is operational
// information.
//
// _Requires permission: `resources:execute`._
//
// GET /api/org/{orgId}/shared-consoles/{consoleId}
//
// Raises on 404: Not found
func (n *SharedConsolesNamespace) Get(ctx context.Context, params SharedConsolesGetParams, opts ...RequestOption) (*SharedConsoleState, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/shared-consoles/{consoleId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	var out *SharedConsoleState
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesHandoverParams holds the parameters for
// `client.sharedConsoles.handover`.
type SharedConsolesHandoverParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
	// Body: the JSON request body.
	Body *SharedConsolesHandoverRequest
}

// Handover: Move the keyboard to another participant
//
// Authorised by the **current driver** (the keyboard is theirs to give) or by
// the **sharer** (it is their box, and asking permission from somebody who has
// stopped responding is not a control). An observer cannot promote themselves —
// that is `/request-driver`.
//
// Two simultaneous grants cannot both win: the database's partial unique index
// decides the order, and the loser gets 409 `driver-race-lost`.
//
// The pty resizes to the new driver's viewport; everyone else letterboxes.
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/shared-consoles/{consoleId}/handover
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *SharedConsolesNamespace) Handover(ctx context.Context, params SharedConsolesHandoverParams, opts ...RequestOption) (*SharedConsoleState, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/shared-consoles/{consoleId}/handover")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	r.setJSONBody(params.Body)
	var out *SharedConsoleState
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesJoinParams holds the parameters for
// `client.sharedConsoles.join`.
type SharedConsolesJoinParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
	// Body: the JSON request body.
	Body *SharedConsolesJoinRequest
}

// Join: Redeem an invite and join
//
// Admission needs live org membership **and** `resources:execute` — the invite
// is a locator, never a capability, so a leaked link admits nobody who could not
// have opened the shell themselves.
//
// The invite is consumed by the first person it admits. Somebody already on the
// console resumes their own row without a token, so a reload costs them nothing
// and obliges the sharer to mint nothing. New joiners always start as observers
// whatever the link said.
//
// Audit-logged as `shared_console.join`, and written onto the recording's
// timeline as an asciicast marker.
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/shared-consoles/{consoleId}/join
//
// Raises on 400: Bad request
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *SharedConsolesNamespace) Join(ctx context.Context, params SharedConsolesJoinParams, opts ...RequestOption) (*SharedConsoleJoined, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/shared-consoles/{consoleId}/join")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	r.setJSONBody(params.Body)
	var out *SharedConsoleJoined
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesLeaveParams holds the parameters for
// `client.sharedConsoles.leave`.
type SharedConsolesLeaveParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
}

// Leave: Leave a shared console
//
// Steps you off without ending the session. Your row survives, so the same
// invite is not needed again. Deliberately does not require `resources:execute`:
// giving access up must never be gated on still holding it.
//
// POST /api/org/{orgId}/shared-consoles/{consoleId}/leave
//
// Raises on 404: Not found
func (n *SharedConsolesNamespace) Leave(ctx context.Context, params SharedConsolesLeaveParams, opts ...RequestOption) (*SharedConsoleState, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/shared-consoles/{consoleId}/leave")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	var out *SharedConsoleState
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesListParams holds the parameters for
// `client.sharedConsoles.list`.
//
// Every field is optional; pass nil to take the defaults.
type SharedConsolesListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List sessions currently shared
//
// Live shared SSH sessions in this organization, with who is on each. Only cloud
// SSH can be shared: those sessions are already proxied by the server, so
// fanning the pty out to a second socket is a consumer of a stream it holds. A
// desktop session dialling a host directly never reaches the server and cannot
// be shared.
//
// _Requires permission: `resources:execute`._
//
// GET /api/org/{orgId}/shared-consoles
func (n *SharedConsolesNamespace) List(ctx context.Context, params *SharedConsolesListParams, opts ...RequestOption) ([]SharedConsoleSummary, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/shared-consoles")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []SharedConsoleSummary
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesRequestDriverParams holds the parameters for
// `client.sharedConsoles.requestDriver`.
type SharedConsolesRequestDriverParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
}

// RequestDriver: Ask for the keyboard
//
// Raises a flag the driver and the sharer can see. Grants nothing on its own —
// that is the point.
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/shared-consoles/{consoleId}/request-driver
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *SharedConsolesNamespace) RequestDriver(ctx context.Context, params SharedConsolesRequestDriverParams, opts ...RequestOption) (*SharedConsoleState, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/shared-consoles/{consoleId}/request-driver")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	var out *SharedConsoleState
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesInvitesNamespace is `client.sharedConsoles.invites`.
type SharedConsolesInvitesNamespace struct {
	t *transport
}

func newSharedConsolesInvitesNamespace(t *transport) *SharedConsolesInvitesNamespace {
	n := &SharedConsolesInvitesNamespace{t: t}
	return n
}

// SharedConsolesInvitesCreateParams holds the parameters for
// `client.sharedConsoles.invites.create`.
type SharedConsolesInvitesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
	// Body: the JSON request body.
	Body *SharedConsolesInvitesCreateRequest
}

// Create: Mint a replacement invite
//
// An invite is spent by the first person it admits, so inviting a second guest
// means minting a second link. Replaces any outstanding one. Sharer or
// `org:settings:write`.
//
// POST /api/org/{orgId}/shared-consoles/{consoleId}/invites
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *SharedConsolesInvitesNamespace) Create(ctx context.Context, params SharedConsolesInvitesCreateParams, opts ...RequestOption) (*SharedConsoleCreated, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/shared-consoles/{consoleId}/invites")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	r.setJSONBody(params.Body)
	var out *SharedConsoleCreated
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesInvitesDeleteParams holds the parameters for
// `client.sharedConsoles.invites.delete`.
type SharedConsolesInvitesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	ConsoleID string
}

// Delete: Withdraw the outstanding invite
//
// Kills the link without touching the session or anyone already on it.
//
// DELETE /api/org/{orgId}/shared-consoles/{consoleId}/invites
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
func (n *SharedConsolesInvitesNamespace) Delete(ctx context.Context, params SharedConsolesInvitesDeleteParams, opts ...RequestOption) (*SharedConsoleState, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/shared-consoles/{consoleId}/invites")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	var out *SharedConsoleState
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesInvitesGetParams holds the parameters for
// `client.sharedConsoles.invites.get`.
type SharedConsolesInvitesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	Token string
}

// Get: Preview what an invite link points at
//
// What the join screen shows before anyone commits: which host, whose session,
// and whether you may join it. Reachable with a valid token by a signed-in
// member who already holds `resources:execute` — the token says *which* session,
// never *whether*. Returns nothing from the session itself.
//
// _Requires permission: `resources:execute`._
//
// GET /api/org/{orgId}/shared-consoles/invites/{token}
//
// Raises on 404: Not found
func (n *SharedConsolesInvitesNamespace) Get(ctx context.Context, params SharedConsolesInvitesGetParams, opts ...RequestOption) (*SharedConsoleInvitePreview, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/shared-consoles/invites/{token}")
	r.setPath("orgId", params.OrgID)
	r.setPath("token", params.Token)
	var out *SharedConsoleInvitePreview
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SharedConsolesParticipantsNamespace is `client.sharedConsoles.participants`.
type SharedConsolesParticipantsNamespace struct {
	t *transport
}

func newSharedConsolesParticipantsNamespace(t *transport) *SharedConsolesParticipantsNamespace {
	n := &SharedConsolesParticipantsNamespace{t: t}
	return n
}

// SharedConsolesParticipantsDeleteParams holds the parameters for
// `client.sharedConsoles.participants.delete`.
type SharedConsolesParticipantsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID         *string
	ConsoleID     string
	ParticipantID string
}

// Delete: Remove somebody from a shared console
//
// Their socket is closed immediately on the replica holding the pty, and within
// one two-second sweep on any other. They are marked `removed` rather than
// `left`, so they cannot resume without a fresh invite. The sharer cannot be
// removed — revoke the share.
//
// DELETE
// /api/org/{orgId}/shared-consoles/{consoleId}/participants/{participantId}
//
// Raises on 403: Forbidden
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *SharedConsolesParticipantsNamespace) Delete(ctx context.Context, params SharedConsolesParticipantsDeleteParams, opts ...RequestOption) (*SharedConsoleState, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/shared-consoles/{consoleId}/participants/{participantId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("consoleId", params.ConsoleID)
	r.setPath("participantId", params.ParticipantID)
	var out *SharedConsoleState
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackNamespace is `client.slack`.
type SlackNamespace struct {
	t *transport

	// Channels: `client.slack.channels`.
	Channels *SlackChannelsNamespace
	// Installations: `client.slack.installations`.
	Installations *SlackInstallationsNamespace
}

func newSlackNamespace(t *transport) *SlackNamespace {
	n := &SlackNamespace{t: t}
	n.Channels = newSlackChannelsNamespace(t)
	n.Installations = newSlackInstallationsNamespace(t)
	return n
}

// SlackInstallURLParams holds the parameters for `client.slack.installUrl`.
//
// Every field is optional; pass nil to take the defaults.
type SlackInstallURLParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// InstallURL: Get the Add to Slack URL
//
// Returns a slack.com/oauth/v2/authorize URL carrying a signed `state` that
// binds the resulting install to this organization. Send the user's browser
// there; Slack redirects back to /api/slack/oauth/callback.
//
// GET /api/org/{orgId}/slack/install-url
//
// Raises on 400: Bad request
func (n *SlackNamespace) InstallURL(ctx context.Context, params *SlackInstallURLParams, opts ...RequestOption) (*SlackInstallUrlresponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/slack/install-url")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SlackInstallUrlresponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackStatusParams holds the parameters for `client.slack.status`.
//
// Every field is optional; pass nil to take the defaults.
type SlackStatusParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Status: Get the organization's Slack connection
//
// Reports whether the server has a Slack app registered, which workspaces this
// organization has connected, and which channels alerts are routed to.
//
// GET /api/org/{orgId}/slack/status
func (n *SlackNamespace) Status(ctx context.Context, params *SlackStatusParams, opts ...RequestOption) (*SlackStatus, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/slack/status")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SlackStatus
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackTestParams holds the parameters for `client.slack.test`.
//
// Every field is optional; pass nil to take the defaults.
type SlackTestParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Test: Post a test message to every configured channel
//
// Ignores routing rules — every channel gets the test. Fails with the Slack
// error when nothing could be delivered (`not_in_channel` means the bot needs
// inviting to a private channel).
//
// POST /api/org/{orgId}/slack/test
//
// Raises on 400: Bad request
func (n *SlackNamespace) Test(ctx context.Context, params *SlackTestParams, opts ...RequestOption) (*SlackTestResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/slack/test")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SlackTestResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackChannelsNamespace is `client.slack.channels`.
type SlackChannelsNamespace struct {
	t *transport
}

func newSlackChannelsNamespace(t *transport) *SlackChannelsNamespace {
	n := &SlackChannelsNamespace{t: t}
	return n
}

// SlackChannelsCreateParams holds the parameters for
// `client.slack.channels.create`.
//
// Every field is optional; pass nil to take the defaults.
type SlackChannelsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *SlackChannelCreate
}

// Create: Connect a Slack channel as an alert destination
//
// Adds a channel as a possible destination, or refreshes the cached name of one
// already added. Which alerts reach it is decided by /alert-rules; an
// organization with no rules falls back to the default (everything except drift,
// everywhere), so a freshly added channel starts receiving alerts without a
// second step.
//
// POST /api/org/{orgId}/slack/channels
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SlackChannelsNamespace) Create(ctx context.Context, params *SlackChannelsCreateParams, opts ...RequestOption) (*SlackChannel, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/slack/channels")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *SlackChannel
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackChannelsDeleteParams holds the parameters for
// `client.slack.channels.delete`.
type SlackChannelsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Disconnect a channel
//
// DELETE /api/org/{orgId}/slack/channels/{id}
//
// Raises on 404: Not found
func (n *SlackChannelsNamespace) Delete(ctx context.Context, params SlackChannelsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/slack/channels/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackChannelsUpdateParams holds the parameters for
// `client.slack.channels.update`.
type SlackChannelsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body *SlackChannelUpdate
}

// Update: Refresh a channel's cached name
//
// PATCH /api/org/{orgId}/slack/channels/{id}
//
// Raises on 404: Not found
func (n *SlackChannelsNamespace) Update(ctx context.Context, params SlackChannelsUpdateParams, opts ...RequestOption) (*SlackChannel, error) {
	r := newRequest(http.MethodPatch, "/api/org/{orgId}/slack/channels/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *SlackChannel
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackInstallationsNamespace is `client.slack.installations`.
type SlackInstallationsNamespace struct {
	t *transport
}

func newSlackInstallationsNamespace(t *transport) *SlackInstallationsNamespace {
	n := &SlackInstallationsNamespace{t: t}
	return n
}

// SlackInstallationsAvailableChannelsParams holds the parameters for
// `client.slack.installations.availableChannels`.
type SlackInstallationsAvailableChannelsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID          *string
	InstallationID string
}

// AvailableChannels: List channels the connected workspace can see
//
// Live call to Slack's conversations.list, for populating a channel picker.
// Returns non-archived public and private channels visible to the bot.
//
// GET /api/org/{orgId}/slack/installations/{installationId}/available-channels
//
// Raises on 400: Bad request
func (n *SlackInstallationsNamespace) AvailableChannels(ctx context.Context, params SlackInstallationsAvailableChannelsParams, opts ...RequestOption) (*SlackInstallationsAvailableChannelsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/slack/installations/{installationId}/available-channels")
	r.setPath("orgId", params.OrgID)
	r.setPath("installationId", params.InstallationID)
	var out *SlackInstallationsAvailableChannelsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SlackInstallationsDeleteParams holds the parameters for
// `client.slack.installations.delete`.
type SlackInstallationsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID          *string
	InstallationID string
}

// Delete: Disconnect a Slack workspace
//
// Stops all delivery to this workspace. The channel routing is retained, so
// re-installing restores it.
//
// DELETE /api/org/{orgId}/slack/installations/{installationId}
//
// Raises on 404: Not found
func (n *SlackInstallationsNamespace) Delete(ctx context.Context, params SlackInstallationsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/slack/installations/{installationId}")
	r.setPath("orgId", params.OrgID)
	r.setPath("installationId", params.InstallationID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SQLNamespace is `client.sql`.
type SQLNamespace struct {
	t *transport
}

func newSQLNamespace(t *transport) *SQLNamespace {
	n := &SQLNamespace{t: t}
	return n
}

// SQLEstimateParams holds the parameters for `client.sql.estimate`.
type SQLEstimateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SQLEstimateRequest
}

// Estimate: Dry-run cost estimate (e.g. BigQuery byte scan)
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/sql/estimate
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SQLNamespace) Estimate(ctx context.Context, params SQLEstimateParams, opts ...RequestOption) (JSONObject, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/sql/estimate")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out JSONObject
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SQLExecuteParams holds the parameters for `client.sql.execute`.
type SQLExecuteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SQLExecuteRequest
}

// Execute: Run an INSERT/UPDATE/DELETE/DDL statement
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/sql/execute
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SQLNamespace) Execute(ctx context.Context, params SQLExecuteParams, opts ...RequestOption) (*SQLExecuteResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/sql/execute")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *SQLExecuteResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SQLQueryParams holds the parameters for `client.sql.query`.
type SQLQueryParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SQLQueryRequest
}

// Query: Run a read-only SQL query
//
// Routes to the right driver: REST `executeQuery` (BigQuery, Databricks),
// per-resource SQL driver (Neon, Turso) or account-level SQL driver (Postgres,
// MySQL).
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/sql/query
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SQLNamespace) Query(ctx context.Context, params SQLQueryParams, opts ...RequestOption) (any, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/sql/query")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out any
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHFanoutNamespace is `client.sshFanout`.
type SSHFanoutNamespace struct {
	t *transport

	// Snippets: `client.sshFanout.snippets`.
	Snippets *SSHFanoutSnippetsNamespace
}

func newSSHFanoutNamespace(t *transport) *SSHFanoutNamespace {
	n := &SSHFanoutNamespace{t: t}
	n.Snippets = newSSHFanoutSnippetsNamespace(t)
	return n
}

// SSHFanoutRunParams holds the parameters for `client.sshFanout.run`.
type SSHFanoutRunParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SSHFanoutRunRequest
}

// Run: Run one command across many SSH hosts
//
// Executes the command on every selected target under a concurrency cap (default
// 8, max 16). Per-host results carry stdout, stderr, and exit code; transport
// failures (unreachable, untrusted host key, blocked internal host) are per-host
// too. Resource targets need `sshKeyId` (an org SSH key owned by the caller).
// Blocked with HTTP 423 while a change freeze is in effect; audit-logged.
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/ssh-fanout/run
//
// Raises on 400: Bad request
//
// Raises on 423: Blocked by an active change freeze. Retry with the
// `x-change-freeze-override: true` header if you hold `freezes:override`; both
// blocks and overrides are audit-logged.
func (n *SSHFanoutNamespace) Run(ctx context.Context, params SSHFanoutRunParams, opts ...RequestOption) (*SSHFanoutRunResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-fanout/run")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *SSHFanoutRunResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHFanoutTargetsParams holds the parameters for `client.sshFanout.targets`.
//
// Every field is optional; pass nil to take the defaults.
type SSHFanoutTargetsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Targets: List SSH-capable fan-out targets
//
// Every SSH-capable target in the org: `ssh` plugin accounts (native
// credentials) plus resources whose type declares an sshEndpoint with a
// resolvable host (EC2 instances, droplets, Hetzner servers, …).
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/ssh-fanout/targets
func (n *SSHFanoutNamespace) Targets(ctx context.Context, params *SSHFanoutTargetsParams, opts ...RequestOption) (*SSHFanoutTargetsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/ssh-fanout/targets")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SSHFanoutTargetsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHFanoutSnippetsNamespace is `client.sshFanout.snippets`.
type SSHFanoutSnippetsNamespace struct {
	t *transport
}

func newSSHFanoutSnippetsNamespace(t *transport) *SSHFanoutSnippetsNamespace {
	n := &SSHFanoutSnippetsNamespace{t: t}
	return n
}

// SSHFanoutSnippetsCreateParams holds the parameters for
// `client.sshFanout.snippets.create`.
type SSHFanoutSnippetsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SSHSnippetInput
}

// Create: Save a command snippet
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/ssh-fanout/snippets
//
// Raises on 400: Bad request
//
// Raises on 409: Conflict
func (n *SSHFanoutSnippetsNamespace) Create(ctx context.Context, params SSHFanoutSnippetsCreateParams, opts ...RequestOption) (*SshfanoutSnippetsCreateResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-fanout/snippets")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *SshfanoutSnippetsCreateResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHFanoutSnippetsDeleteParams holds the parameters for
// `client.sshFanout.snippets.delete`.
type SSHFanoutSnippetsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a saved command snippet
//
// _Requires permission: `resources:execute`._
//
// DELETE /api/org/{orgId}/ssh-fanout/snippets/{id}
//
// Raises on 404: Not found
func (n *SSHFanoutSnippetsNamespace) Delete(ctx context.Context, params SSHFanoutSnippetsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/ssh-fanout/snippets/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHFanoutSnippetsGetParams holds the parameters for
// `client.sshFanout.snippets.get`.
//
// Every field is optional; pass nil to take the defaults.
type SSHFanoutSnippetsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List saved command snippets
//
// Org-shared saved commands for reuse from the fan-out screen and CLI.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/ssh-fanout/snippets
func (n *SSHFanoutSnippetsNamespace) Get(ctx context.Context, params *SSHFanoutSnippetsGetParams, opts ...RequestOption) (*SshfanoutSnippetsGetResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/ssh-fanout/snippets")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *SshfanoutSnippetsGetResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHFanoutSnippetsUpdateParams holds the parameters for
// `client.sshFanout.snippets.update`.
type SSHFanoutSnippetsUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body SSHSnippetInput
}

// Update: Update a saved command snippet
//
// _Requires permission: `resources:execute`._
//
// PUT /api/org/{orgId}/ssh-fanout/snippets/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SSHFanoutSnippetsNamespace) Update(ctx context.Context, params SSHFanoutSnippetsUpdateParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/ssh-fanout/snippets/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHKeysNamespace is `client.sshKeys`.
type SSHKeysNamespace struct {
	t *transport
}

func newSSHKeysNamespace(t *transport) *SSHKeysNamespace {
	n := &SSHKeysNamespace{t: t}
	return n
}

// SSHKeysCreateParams holds the parameters for `client.sshKeys.create`.
type SSHKeysCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body GenerateSSHKeyRequest
}

// Create: Generate a new Ed25519 keypair (private key returned once)
//
// _Requires permission: `ssh-keys:write`._
//
// POST /api/org/{orgId}/ssh-keys
//
// Raises on 400: Bad request
func (n *SSHKeysNamespace) Create(ctx context.Context, params SSHKeysCreateParams, opts ...RequestOption) (*GeneratedSSHKey, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-keys")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *GeneratedSSHKey
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHKeysDeleteParams holds the parameters for `client.sshKeys.delete`.
type SSHKeysDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete an SSH key (owner only)
//
// _Requires permission: `ssh-keys:write`._
//
// DELETE /api/org/{orgId}/ssh-keys/{id}
func (n *SSHKeysNamespace) Delete(ctx context.Context, params SSHKeysDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/ssh-keys/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHKeysImportParams holds the parameters for `client.sshKeys.import`.
type SSHKeysImportParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body ImportSSHKeyRequest
}

// Import: Import an existing public key
//
// _Requires permission: `ssh-keys:write`._
//
// POST /api/org/{orgId}/ssh-keys/import
//
// Raises on 400: Bad request
func (n *SSHKeysNamespace) Import(ctx context.Context, params SSHKeysImportParams, opts ...RequestOption) (*ImportedSSHKey, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-keys/import")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ImportedSSHKey
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHKeysListParams holds the parameters for `client.sshKeys.list`.
//
// Every field is optional; pass nil to take the defaults.
type SSHKeysListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List org SSH keys
//
// _Requires permission: `ssh-keys:read`._
//
// GET /api/org/{orgId}/ssh-keys
func (n *SSHKeysNamespace) List(ctx context.Context, params *SSHKeysListParams, opts ...RequestOption) ([]SSHKey, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/ssh-keys")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []SSHKey
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHTunnelsNamespace is `client.sshTunnels`.
type SSHTunnelsNamespace struct {
	t *transport
}

func newSSHTunnelsNamespace(t *transport) *SSHTunnelsNamespace {
	n := &SSHTunnelsNamespace{t: t}
	return n
}

// SSHTunnelsActiveParams holds the parameters for `client.sshTunnels.active`.
//
// Every field is optional; pass nil to take the defaults.
type SSHTunnelsActiveParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Active: List active tunnels for this org
//
// _Requires permission: `resources:execute`._
//
// GET /api/org/{orgId}/ssh-tunnels/active
func (n *SSHTunnelsNamespace) Active(ctx context.Context, params *SSHTunnelsActiveParams, opts ...RequestOption) (map[string]ActiveTunnel, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/ssh-tunnels/active")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out map[string]ActiveTunnel
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHTunnelsCloseParams holds the parameters for `client.sshTunnels.close`.
type SSHTunnelsCloseParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SshtunnelsCloseRequest
}

// Close: Close a tunnel by id
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/ssh-tunnels/close
func (n *SSHTunnelsNamespace) Close(ctx context.Context, params SSHTunnelsCloseParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-tunnels/close")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHTunnelsCreateAccountParams holds the parameters for
// `client.sshTunnels.createAccount`.
type SSHTunnelsCreateAccountParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SSHTunnelCreateAccountRequest
}

// CreateAccount: Create an account whose traffic is tunneled over SSH
//
// Verifies the SSH connection works before persisting.
//
// _Requires permission: `accounts:write`._
//
// POST /api/org/{orgId}/ssh-tunnels/create-account
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SSHTunnelsNamespace) CreateAccount(ctx context.Context, params SSHTunnelsCreateAccountParams, opts ...RequestOption) (*SSHTunnelCreateAccountResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-tunnels/create-account")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *SSHTunnelCreateAccountResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHTunnelsExecParams holds the parameters for `client.sshTunnels.exec`.
type SSHTunnelsExecParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SSHExecRequest
}

// Exec: Run a command over SSH using an org SSH key
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/ssh-tunnels/exec
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *SSHTunnelsNamespace) Exec(ctx context.Context, params SSHTunnelsExecParams, opts ...RequestOption) (*SSHExecResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-tunnels/exec")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *SSHExecResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// SSHTunnelsOpenParams holds the parameters for `client.sshTunnels.open`.
type SSHTunnelsOpenParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body SshtunnelsOpenRequest
}

// Open: Re-open the tunnel for an existing account
//
// _Requires permission: `resources:execute`._
//
// POST /api/org/{orgId}/ssh-tunnels/open
//
// Raises on 404: Not found
func (n *SSHTunnelsNamespace) Open(ctx context.Context, params SSHTunnelsOpenParams, opts ...RequestOption) (*SshtunnelsOpenResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/ssh-tunnels/open")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *SshtunnelsOpenResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StatusNamespace is `client.status`.
type StatusNamespace struct {
	t *transport
}

func newStatusNamespace(t *transport) *StatusNamespace {
	n := &StatusNamespace{t: t}
	return n
}

// StatusGetParams holds the parameters for `client.status.get`.
type StatusGetParams struct {
	Slug string
}

// Get: Read a public status page
//
// **Unauthenticated.** The only endpoint in this API that takes no credentials —
// a status page exists for people with no account. The payload carries labels,
// states and uptime history only: probe URLs, resource and account ids, the
// organization id and error detail are never included. An unpublished page and
// an unknown slug both answer 404, so the endpoint cannot be used to confirm
// that a slug is real.
//
// GET /api/status/{slug}
//
// Raises on 404: Not found
func (n *StatusNamespace) Get(ctx context.Context, params StatusGetParams, opts ...RequestOption) (*PublicStatusPage, error) {
	r := newRequest(http.MethodGet, "/api/status/{slug}")
	r.setPath("slug", params.Slug)
	var out *PublicStatusPage
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StatusIncidentsNamespace is `client.statusIncidents`.
type StatusIncidentsNamespace struct {
	t *transport
}

func newStatusIncidentsNamespace(t *transport) *StatusIncidentsNamespace {
	n := &StatusIncidentsNamespace{t: t}
	return n
}

// StatusIncidentsGetParams holds the parameters for
// `client.statusIncidents.get`.
//
// Every field is optional; pass nil to take the defaults.
type StatusIncidentsGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: Provider incidents overlapping your resources
//
// The "is it me or is it them?" feed. The poller watches each provider plugin's
// public status feed (declared on its manifest — zero credentials, zero
// rate-limit risk), caches active incidents, and this endpoint correlates them
// against the resources the organization holds: an incident matches a resource
// when it is provider-wide, names the resource's region, or names its resource
// type. Includes incidents resolved within the last 24 hours so recent drift can
// still be correlated. Active incidents first, most severe first.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/status-incidents
//
// Raises on 400: Bad request
func (n *StatusIncidentsNamespace) Get(ctx context.Context, params *StatusIncidentsGetParams, opts ...RequestOption) (*OrgStatusIncidentsResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/status-incidents")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *OrgStatusIncidentsResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StatusPagesNamespace is `client.statusPages`.
type StatusPagesNamespace struct {
	t *transport
}

func newStatusPagesNamespace(t *transport) *StatusPagesNamespace {
	n := &StatusPagesNamespace{t: t}
	return n
}

// StatusPagesCreateParams holds the parameters for `client.statusPages.create`.
//
// Every field is optional; pass nil to take the defaults.
type StatusPagesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body *StatusPageCreate
}

// Create: Create a status page
//
// Creates a page with a freshly generated slug. `published` defaults to false,
// so creating a page never exposes anything — publish it as a separate,
// deliberate step.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/status-pages
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *StatusPagesNamespace) Create(ctx context.Context, params *StatusPagesCreateParams, opts ...RequestOption) (*StatusPage, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/status-pages")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.setJSONBody(params.Body)
	}
	var out *StatusPage
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StatusPagesDeleteParams holds the parameters for `client.statusPages.delete`.
type StatusPagesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a status page
//
// The page's link stops working. The probes it published are untouched.
//
// _Requires permission: `resources:write`._
//
// DELETE /api/org/{orgId}/status-pages/{id}
//
// Raises on 404: Not found
func (n *StatusPagesNamespace) Delete(ctx context.Context, params StatusPagesDeleteParams, opts ...RequestOption) error {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/status-pages/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	return n.t.do(ctx, r, nil, opts)
}

// StatusPagesGetParams holds the parameters for `client.statusPages.get`.
//
// Every field is optional; pass nil to take the defaults.
type StatusPagesGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: List status pages
//
// Every status page in the organization, with the probes each publishes and
// whether it is currently reachable.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/status-pages
func (n *StatusPagesNamespace) Get(ctx context.Context, params *StatusPagesGetParams, opts ...RequestOption) (*StatusPageListResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/status-pages")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *StatusPageListResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StatusPagesRotateSlugParams holds the parameters for
// `client.statusPages.rotateSlug`.
type StatusPagesRotateSlugParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// RotateSlug: Issue a new public link
//
// Replaces the slug, revoking the current public URL immediately — the reroll
// for a link that ended up somewhere unintended. The page stays published.
//
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/status-pages/{id}/rotate-slug
//
// Raises on 404: Not found
func (n *StatusPagesNamespace) RotateSlug(ctx context.Context, params StatusPagesRotateSlugParams, opts ...RequestOption) (*StatusPage, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/status-pages/{id}/rotate-slug")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *StatusPage
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StatusPagesUpdateParams holds the parameters for `client.statusPages.update`.
type StatusPagesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body *StatusPagePatch
}

// Update: Update a status page
//
// Omitted fields keep their value. `components`, when present, replaces the
// whole ordered set — which is also how a reorder is expressed.
//
// _Requires permission: `resources:write`._
//
// PUT /api/org/{orgId}/status-pages/{id}
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *StatusPagesNamespace) Update(ctx context.Context, params StatusPagesUpdateParams, opts ...RequestOption) (*StatusPage, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/status-pages/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *StatusPage
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StorageNamespace is `client.storage`.
type StorageNamespace struct {
	t *transport
}

func newStorageNamespace(t *transport) *StorageNamespace {
	n := &StorageNamespace{t: t}
	return n
}

// StorageDeleteParams holds the parameters for `client.storage.delete`.
type StorageDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body StoragePathRequest
}

// Delete: Delete a storage object
//
// _Requires permission: `storage:write`._
//
// POST /api/org/{orgId}/storage/delete
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *StorageNamespace) Delete(ctx context.Context, params StorageDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/storage/delete")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StorageDownloadParams holds the parameters for `client.storage.download`.
type StorageDownloadParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID     *string
	AccountID string
	Bucket    string
	// Keys: JSON-encoded array of object keys, e.g. `["a.txt","b.txt"]`
	Keys string
}

// Download: Download one or many objects (zipped if more than one)
//
// _Requires permission: `storage:read`._
//
// GET /api/org/{orgId}/v1/storage/download
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
//
// Raises on 500: Server error
func (n *StorageNamespace) Download(ctx context.Context, params StorageDownloadParams, opts ...RequestOption) (io.ReadCloser, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/v1/storage/download")
	r.setPath("orgId", params.OrgID)
	r.addQuery("accountId", params.AccountID)
	r.addQuery("bucket", params.Bucket)
	r.addQuery("keys", params.Keys)
	return n.t.stream(ctx, r, opts)
}

// StorageListParams holds the parameters for `client.storage.list`.
type StorageListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body StorageListRequest
}

// List: List objects in a bucket / prefix
//
// _Requires permission: `storage:read`._
//
// POST /api/org/{orgId}/storage/list
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *StorageNamespace) List(ctx context.Context, params StorageListParams, opts ...RequestOption) ([]StorageObject, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/storage/list")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out []StorageObject
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StorageMkdirParams holds the parameters for `client.storage.mkdir`.
type StorageMkdirParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body StoragePathRequest
}

// Mkdir: Create a folder marker in a bucket
//
// _Requires permission: `storage:write`._
//
// POST /api/org/{orgId}/storage/mkdir
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *StorageNamespace) Mkdir(ctx context.Context, params StorageMkdirParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/storage/mkdir")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// StorageUploadParams holds the parameters for `client.storage.upload`.
type StorageUploadParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: sent as `multipart/form-data`; the `io.Reader` field is the file.
	Body StorageUploadForm
}

// Upload: Upload a file to object storage
//
// Multipart/form-data. Plugin must implement `uploadStorageObject`.
//
// _Requires permission: `storage:write`._
//
// POST /api/org/{orgId}/v1/storage/upload
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *StorageNamespace) Upload(ctx context.Context, params StorageUploadParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/v1/storage/upload")
	r.setPath("orgId", params.OrgID)
	r.setFormBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TagPolicyNamespace is `client.tagPolicy`.
type TagPolicyNamespace struct {
	t *transport
}

func newTagPolicyNamespace(t *transport) *TagPolicyNamespace {
	n := &TagPolicyNamespace{t: t}
	return n
}

// TagPolicyComplianceParams holds the parameters for
// `client.tagPolicy.compliance`.
//
// Every field is optional; pass nil to take the defaults.
type TagPolicyComplianceParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Compliance: Per-account tag compliance scores
//
// For each account: how many of its resources expose tags and how many of those
// carry every required tag with an allowed value. `score` is over the evaluated
// (tag-capable) set so untaggable resource types don't drag it.
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/tag-policy/compliance
func (n *TagPolicyNamespace) Compliance(ctx context.Context, params *TagPolicyComplianceParams, opts ...RequestOption) (*TagComplianceReport, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/tag-policy/compliance")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *TagComplianceReport
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TagPolicyGetParams holds the parameters for `client.tagPolicy.get`.
//
// Every field is optional; pass nil to take the defaults.
type TagPolicyGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Get: The org's required-tag policy
//
// _Requires permission: `resources:read`._
//
// GET /api/org/{orgId}/tag-policy
func (n *TagPolicyNamespace) Get(ctx context.Context, params *TagPolicyGetParams, opts ...RequestOption) (*TagPolicy, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/tag-policy")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *TagPolicy
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TagPolicyUpdateParams holds the parameters for `client.tagPolicy.update`.
type TagPolicyUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body TagPolicy
}

// Update: Replace the org's tag policy
//
// Sets the required tag keys (each optionally restricted to allowed values) and
// whether resource creation is blocked when they are missing. Keys are matched
// case-insensitively against the generic `tags`/`labels` field convention.
//
// _Requires permission: `org:settings:write`._
//
// PUT /api/org/{orgId}/tag-policy
//
// Raises on 400: Bad request
func (n *TagPolicyNamespace) Update(ctx context.Context, params TagPolicyUpdateParams, opts ...RequestOption) (*TagPolicy, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/tag-policy")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *TagPolicy
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamNamespace is `client.team`.
type TeamNamespace struct {
	t *transport

	// Invitations: `client.team.invitations`.
	Invitations *TeamInvitationsNamespace
	// Members: `client.team.members`.
	Members *TeamMembersNamespace
	// Roles: `client.team.roles`.
	Roles *TeamRolesNamespace
}

func newTeamNamespace(t *transport) *TeamNamespace {
	n := &TeamNamespace{t: t}
	n.Invitations = newTeamInvitationsNamespace(t)
	n.Members = newTeamMembersNamespace(t)
	n.Roles = newTeamRolesNamespace(t)
	return n
}

// TeamMeParams holds the parameters for `client.team.me`.
//
// Every field is optional; pass nil to take the defaults.
type TeamMeParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Me: Current user's effective permissions and role
//
// GET /api/org/{orgId}/team/me
func (n *TeamNamespace) Me(ctx context.Context, params *TeamMeParams, opts ...RequestOption) (*MeResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/team/me")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *MeResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamPermissionsParams holds the parameters for `client.team.permissions`.
//
// Every field is optional; pass nil to take the defaults.
type TeamPermissionsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Permissions: List all permission strings the server recognises
//
// _Requires permission: `team:read`._
//
// GET /api/org/{orgId}/team/permissions
func (n *TeamNamespace) Permissions(ctx context.Context, params *TeamPermissionsParams, opts ...RequestOption) (*PermissionCatalog, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/team/permissions")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out *PermissionCatalog
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamInvitationsNamespace is `client.team.invitations`.
type TeamInvitationsNamespace struct {
	t *transport
}

func newTeamInvitationsNamespace(t *transport) *TeamInvitationsNamespace {
	n := &TeamInvitationsNamespace{t: t}
	return n
}

// TeamInvitationsCreateParams holds the parameters for
// `client.team.invitations.create`.
type TeamInvitationsCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body InviteRequest
}

// Create: Create an invitation (token valid for 7 days)
//
// _Requires permission: `team:invite`._
//
// POST /api/org/{orgId}/team/invitations
//
// Raises on 402: Payment required — the organization's plan does not include
// this
//
// Raises on 403: The role would grant permissions the caller does not hold, or
// the caller is not an owner and tried to invite an owner
//
// Raises on 409: All seats are in use; retry with addSeat to buy one more
//
// Raises on 502: Buying the extra seat failed; the invitation was not sent
func (n *TeamInvitationsNamespace) Create(ctx context.Context, params TeamInvitationsCreateParams, opts ...RequestOption) (*InviteResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/team/invitations")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *InviteResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamInvitationsDeleteParams holds the parameters for
// `client.team.invitations.delete`.
type TeamInvitationsDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Revoke a pending invitation
//
// _Requires permission: `team:invite`._
//
// DELETE /api/org/{orgId}/team/invitations/{id}
//
// Raises on 404: Not found
func (n *TeamInvitationsNamespace) Delete(ctx context.Context, params TeamInvitationsDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/team/invitations/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamInvitationsListParams holds the parameters for
// `client.team.invitations.list`.
//
// Every field is optional; pass nil to take the defaults.
type TeamInvitationsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List pending and historical invitations
//
// _Requires permission: `team:read`._
//
// GET /api/org/{orgId}/team/invitations
func (n *TeamInvitationsNamespace) List(ctx context.Context, params *TeamInvitationsListParams, opts ...RequestOption) ([]Invitation, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/team/invitations")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []Invitation
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamMembersNamespace is `client.team.members`.
type TeamMembersNamespace struct {
	t *transport
}

func newTeamMembersNamespace(t *transport) *TeamMembersNamespace {
	n := &TeamMembersNamespace{t: t}
	return n
}

// TeamMembersDeleteParams holds the parameters for `client.team.members.delete`.
type TeamMembersDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Remove a member from the org
//
// _Requires permission: `team:remove`._
//
// DELETE /api/org/{orgId}/team/members/{id}
func (n *TeamMembersNamespace) Delete(ctx context.Context, params TeamMembersDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/team/members/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamMembersListParams holds the parameters for `client.team.members.list`.
//
// Every field is optional; pass nil to take the defaults.
type TeamMembersListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List org members
//
// _Requires permission: `team:read`._
//
// GET /api/org/{orgId}/team/members
func (n *TeamMembersNamespace) List(ctx context.Context, params *TeamMembersListParams, opts ...RequestOption) ([]OrgMember, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/team/members")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []OrgMember
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamMembersRoleParams holds the parameters for `client.team.members.role`.
type TeamMembersRoleParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body RoleChangeRequest
}

// Role: Change a member's role
//
// _Requires permission: `team:role:write`._
//
// PATCH /api/org/{orgId}/team/members/{id}/role
func (n *TeamMembersNamespace) Role(ctx context.Context, params TeamMembersRoleParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodPatch, "/api/org/{orgId}/team/members/{id}/role")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamRolesNamespace is `client.team.roles`.
type TeamRolesNamespace struct {
	t *transport
}

func newTeamRolesNamespace(t *transport) *TeamRolesNamespace {
	n := &TeamRolesNamespace{t: t}
	return n
}

// TeamRolesCreateParams holds the parameters for `client.team.roles.create`.
type TeamRolesCreateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body RoleCreateRequest
}

// Create: Create a custom role
//
// _Requires permission: `team:role:write`._
//
// POST /api/org/{orgId}/team/roles
func (n *TeamRolesNamespace) Create(ctx context.Context, params TeamRolesCreateParams, opts ...RequestOption) (*Role, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/team/roles")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *Role
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamRolesDeleteParams holds the parameters for `client.team.roles.delete`.
type TeamRolesDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Delete: Delete a custom role (must have no members or pending invitations)
//
// _Requires permission: `team:role:write`._
//
// DELETE /api/org/{orgId}/team/roles/{id}
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
//
// Raises on 422: Bad request
func (n *TeamRolesNamespace) Delete(ctx context.Context, params TeamRolesDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/team/roles/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamRolesListParams holds the parameters for `client.team.roles.list`.
//
// Every field is optional; pass nil to take the defaults.
type TeamRolesListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// List: List roles (system + custom)
//
// _Requires permission: `team:read`._
//
// GET /api/org/{orgId}/team/roles
func (n *TeamRolesNamespace) List(ctx context.Context, params *TeamRolesListParams, opts ...RequestOption) ([]Role, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/team/roles")
	if params != nil {
		r.setPath("orgId", params.OrgID)
	}
	var out []Role
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// TeamRolesUpdateParams holds the parameters for `client.team.roles.update`.
type TeamRolesUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
	// Body: the JSON request body.
	Body RoleUpdateRequest
}

// Update: Edit a custom role
//
// _Requires permission: `team:role:write`._
//
// PATCH /api/org/{orgId}/team/roles/{id}
//
// Raises on 404: Not found
//
// Raises on 422: Bad request
func (n *TeamRolesNamespace) Update(ctx context.Context, params TeamRolesUpdateParams, opts ...RequestOption) (*Role, error) {
	r := newRequest(http.MethodPatch, "/api/org/{orgId}/team/roles/{id}")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *Role
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// WorkflowApprovalsNamespace is `client.workflowApprovals`.
type WorkflowApprovalsNamespace struct {
	t *transport
}

func newWorkflowApprovalsNamespace(t *transport) *WorkflowApprovalsNamespace {
	n := &WorkflowApprovalsNamespace{t: t}
	return n
}

// WorkflowApprovalsApproveParams holds the parameters for
// `client.workflowApprovals.approve`.
type WorkflowApprovalsApproveParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Approve: Approve a pending workflow approval request
//
// The suspended run resumes within a few seconds of the decision landing.
//
// _Requires permission: `workflows:approve`._
//
// POST /api/org/{orgId}/workflow-approvals/{id}/approve
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *WorkflowApprovalsNamespace) Approve(ctx context.Context, params WorkflowApprovalsApproveParams, opts ...RequestOption) (*WorkflowApproval, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/workflow-approvals/{id}/approve")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *WorkflowApproval
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// WorkflowApprovalsDenyParams holds the parameters for
// `client.workflowApprovals.deny`.
type WorkflowApprovalsDenyParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	ID    string
}

// Deny: Deny a pending workflow approval request
//
// Denial fails the waiting `infra.waitForApproval(...)` call in the run.
//
// _Requires permission: `workflows:approve`._
//
// POST /api/org/{orgId}/workflow-approvals/{id}/deny
//
// Raises on 404: Not found
//
// Raises on 409: Conflict
func (n *WorkflowApprovalsNamespace) Deny(ctx context.Context, params WorkflowApprovalsDenyParams, opts ...RequestOption) (*WorkflowApproval, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/workflow-approvals/{id}/deny")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *WorkflowApproval
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// WorkflowApprovalsListParams holds the parameters for
// `client.workflowApprovals.list`.
//
// Every field is optional; pass nil to take the defaults.
type WorkflowApprovalsListParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID      *string
	Status     *WorkflowApprovalStatus
	WorkflowID *string
	RunID      *string
}

// List: List workflow approval requests
//
// Approval requests raised by `infra.waitForApproval(...)` inside workflow runs,
// newest first. Filter with `status=pending` to build an approvals inbox.
//
// _Requires permission: `workflows:read`._
//
// GET /api/org/{orgId}/workflow-approvals
//
// Raises on 400: Bad request
func (n *WorkflowApprovalsNamespace) List(ctx context.Context, params *WorkflowApprovalsListParams, opts ...RequestOption) ([]WorkflowApproval, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/workflow-approvals")
	if params != nil {
		r.setPath("orgId", params.OrgID)
		r.addQuery("status", params.Status)
		r.addQuery("workflowId", params.WorkflowID)
		r.addQuery("runId", params.RunID)
	}
	var out []WorkflowApproval
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// WorkflowsNamespace is `client.workflows`.
type WorkflowsNamespace struct {
	t *transport

	// Schedule: `client.workflows.schedule`.
	Schedule *WorkflowsScheduleNamespace
}

func newWorkflowsNamespace(t *transport) *WorkflowsNamespace {
	n := &WorkflowsNamespace{t: t}
	n.Schedule = newWorkflowsScheduleNamespace(t)
	return n
}

// WorkflowsScheduleNamespace is `client.workflows.schedule`.
type WorkflowsScheduleNamespace struct {
	t *transport
}

func newWorkflowsScheduleNamespace(t *transport) *WorkflowsScheduleNamespace {
	n := &WorkflowsScheduleNamespace{t: t}
	return n
}

// WorkflowsScheduleDeleteParams holds the parameters for
// `client.workflows.schedule.delete`.
type WorkflowsScheduleDeleteParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Workflow id
	ID string
}

// Delete: Remove a workflow's cron schedule
//
// Reverts the workflow's trigger to manual and clears the pending fire time. A
// no-op when the trigger is not cron.
//
// _Requires permission: `dashboards:write`._
//
// DELETE /api/org/{orgId}/workflows/{id}/schedule
//
// Raises on 404: Not found
func (n *WorkflowsScheduleNamespace) Delete(ctx context.Context, params WorkflowsScheduleDeleteParams, opts ...RequestOption) (*OK, error) {
	r := newRequest(http.MethodDelete, "/api/org/{orgId}/workflows/{id}/schedule")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *OK
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// WorkflowsScheduleGetParams holds the parameters for
// `client.workflows.schedule.get`.
type WorkflowsScheduleGetParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Workflow id
	ID string
}

// Get: Get a workflow's cron schedule
//
// The schedule view of the workflow's trigger, with the next few computed fire
// times. `schedule` is null when the workflow is triggered some other way
// (manual, git, budget).
//
// _Requires permission: `dashboards:read`._
//
// GET /api/org/{orgId}/workflows/{id}/schedule
//
// Raises on 404: Not found
func (n *WorkflowsScheduleNamespace) Get(ctx context.Context, params WorkflowsScheduleGetParams, opts ...RequestOption) (*WorkflowScheduleResponse, error) {
	r := newRequest(http.MethodGet, "/api/org/{orgId}/workflows/{id}/schedule")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	var out *WorkflowScheduleResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}

// WorkflowsScheduleUpdateParams holds the parameters for
// `client.workflows.schedule.update`.
type WorkflowsScheduleUpdateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// ID: Workflow id
	ID string
	// Body: the JSON request body.
	Body WorkflowScheduleInput
}

// Update: Create or replace a workflow's cron schedule
//
// Sets the workflow's trigger to cron with the given expression and timezone,
// validating both, and computes the next fire time. The workflow fires at the
// schedule's next occurrence — never immediately on save.
//
// _Requires permission: `dashboards:write`._
//
// PUT /api/org/{orgId}/workflows/{id}/schedule
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
func (n *WorkflowsScheduleNamespace) Update(ctx context.Context, params WorkflowsScheduleUpdateParams, opts ...RequestOption) (*WorkflowScheduleResponse, error) {
	r := newRequest(http.MethodPut, "/api/org/{orgId}/workflows/{id}/schedule")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
	r.setJSONBody(params.Body)
	var out *WorkflowScheduleResponse
	if err := n.t.do(ctx, r, &out, opts); err != nil {
		return out, err
	}
	return out, nil
}
