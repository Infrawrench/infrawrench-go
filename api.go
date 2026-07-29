// github.com/Infrawrench/infrawrench-go v0.14.0 | MIT | Copyright (c) 2026 Infrawrench LLC
// https://github.com/Infrawrench/Infrawrench
//
// Generated from the Infrawrench API OpenAPI 3.1 spec (API version 0.14.0).
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

	// Accounts: `client.accounts`.
	Accounts *AccountsNamespace
	// Agents: `client.agents`.
	Agents *AgentsNamespace
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
	// Budgets: `client.budgets`.
	Budgets *BudgetsNamespace
	// Connect: `client.connect`.
	Connect *ConnectNamespace
	// Costs: `client.costs`.
	Costs *CostsNamespace
	// Dashboards: `client.dashboards`.
	Dashboards *DashboardsNamespace
	// Deployments: `client.deployments`.
	Deployments *DeploymentsNamespace
	// Docker: `client.docker`.
	Docker *DockerNamespace
	// Invitations: `client.invitations`.
	Invitations *InvitationsNamespace
	// KV: `client.kv`.
	KV *KVNamespace
	// Msteams: `client.msteams`.
	Msteams *MsteamsNamespace
	// Orgs: `client.orgs`.
	Orgs *OrgsNamespace
	// Pages: `client.pages`.
	Pages *PagesNamespace
	// Profile: `client.profile`.
	Profile *ProfileNamespace
	// Resources: `client.resources`.
	Resources *ResourcesNamespace
	// Search: `client.search`.
	Search *SearchNamespace
	// SFTP: `client.sftp`.
	SFTP *SFTPNamespace
	// Slack: `client.slack`.
	Slack *SlackNamespace
	// SQL: `client.sql`.
	SQL *SQLNamespace
	// SSHKeys: `client.sshKeys`.
	SSHKeys *SSHKeysNamespace
	// SSHTunnels: `client.sshTunnels`.
	SSHTunnels *SSHTunnelsNamespace
	// Storage: `client.storage`.
	Storage *StorageNamespace
	// Team: `client.team`.
	Team *TeamNamespace
}

// NewAPIV1Client builds a client. With no options it talks to
// https://app.infrawrench.com anonymously, which is rarely what you want: pass
// WithAPIKey, and WithOrgID if you would rather not repeat the organization id
// on every call.
func NewAPIV1Client(opts ...ClientOption) *APIV1Client {
	t := newTransport(opts)
	c := &APIV1Client{t: t}
	c.Accounts = newAccountsNamespace(t)
	c.Agents = newAgentsNamespace(t)
	c.APIKeys = newAPIKeysNamespace(t)
	c.Artifacts = newArtifactsNamespace(t)
	c.Associations = newAssociationsNamespace(t)
	c.AuditLogs = newAuditLogsNamespace(t)
	c.Auth = newAuthNamespace(t)
	c.Bastions = newBastionsNamespace(t)
	c.Billing = newBillingNamespace(t)
	c.Budgets = newBudgetsNamespace(t)
	c.Connect = newConnectNamespace(t)
	c.Costs = newCostsNamespace(t)
	c.Dashboards = newDashboardsNamespace(t)
	c.Deployments = newDeploymentsNamespace(t)
	c.Docker = newDockerNamespace(t)
	c.Invitations = newInvitationsNamespace(t)
	c.KV = newKVNamespace(t)
	c.Msteams = newMsteamsNamespace(t)
	c.Orgs = newOrgsNamespace(t)
	c.Pages = newPagesNamespace(t)
	c.Profile = newProfileNamespace(t)
	c.Resources = newResourcesNamespace(t)
	c.Search = newSearchNamespace(t)
	c.SFTP = newSFTPNamespace(t)
	c.Slack = newSlackNamespace(t)
	c.SQL = newSQLNamespace(t)
	c.SSHKeys = newSSHKeysNamespace(t)
	c.SSHTunnels = newSSHTunnelsNamespace(t)
	c.Storage = newStorageNamespace(t)
	c.Team = newTeamNamespace(t)
	return c
}

// BaseURL reports the normalized base URL every call is sent to.
func (c *APIV1Client) BaseURL() string {
	return c.t.baseURL()
}

// AccountsNamespace is `client.accounts`.
type AccountsNamespace struct {
	t *transport

	// Credentials: `client.accounts.credentials`.
	Credentials *AccountsCredentialsNamespace
	// SyncType: `client.accounts.syncType`.
	SyncType *AccountsSyncTypeNamespace
}

func newAccountsNamespace(t *transport) *AccountsNamespace {
	n := &AccountsNamespace{t: t}
	n.Credentials = newAccountsCredentialsNamespace(t)
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

// AccountsPluginsParams holds the parameters for `client.accounts.plugins`.
//
// Every field is optional; pass nil to take the defaults.
type AccountsPluginsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
}

// Plugins: List installed plugins and their credential fields
//
// _Requires permission: `accounts:read`._
//
// GET /api/org/{orgId}/accounts/plugins
func (n *AccountsNamespace) Plugins(ctx context.Context, params *AccountsPluginsParams, opts ...RequestOption) ([]PluginSummary, error) {
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
}

func newBillingNamespace(t *transport) *BillingNamespace {
	n := &BillingNamespace{t: t}
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

// CostsNamespace is `client.costs`.
type CostsNamespace struct {
	t *transport
}

func newCostsNamespace(t *transport) *CostsNamespace {
	n := &CostsNamespace{t: t}
	return n
}

// CostsDimensionsParams holds the parameters for `client.costs.dimensions`.
type CostsDimensionsParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "tag-keys".
	Dimension string
	TagKey    *string
}

// Dimensions: List distinct values for a cost dimension
//
// Feeds the filter and group-by pickers. Pass dimension=tag-keys for tag keys;
// dimension=tag requires tagKey.
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
}

// Rollback: Roll back to a previous deployment
//
// Re-runs that run's `deploy()` with the image and plan it recorded, building
// nothing — the exact artifact that was known good ships again. The Infrafile is
// read at the commit that run deployed, not at the branch head. Only a
// successful run that produced an image can be rolled back to.
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
func (n *DeploymentsRunsNamespace) Rollback(ctx context.Context, params DeploymentsRunsRollbackParams, opts ...RequestOption) (*DeployPlanResult, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/deployments/runs/{id}/rollback")
	r.setPath("orgId", params.OrgID)
	r.setPath("id", params.ID)
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
// Returns the Teams channels alerts are routed to and which triggers each takes.
// Webhook URLs are never included.
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
// Ignores trigger opt-ins — every channel gets the test. Fails with the error
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

// Create: Route alerts to a Teams channel
//
// Adds a channel by webhook URL, or updates the one already holding that URL.
// Each trigger defaults to enabled. Responds 400 when the URL is not https or
// its host is not Microsoft-operated.
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

// Delete: Stop routing alerts to a Teams channel
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

// Update: Rename a Teams channel or change which alerts it receives
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

// ResourcesCreateCostEstimateParams holds the parameters for
// `client.resources.createCostEstimate`.
type ResourcesCreateCostEstimateParams struct {
	// OrgID: Organization id
	//
	// Falls back to the client's `orgId` when omitted.
	OrgID *string
	// Body: the JSON request body.
	Body CreateCostEstimateRequest
}

// CreateCostEstimate: Cost estimate for the current create form values
//
// _Requires permission: `resources:read`._
//
// POST /api/org/{orgId}/resources/create-cost-estimate
func (n *ResourcesNamespace) CreateCostEstimate(ctx context.Context, params ResourcesCreateCostEstimateParams, opts ...RequestOption) (*ResourcesCreateCostEstimateResponse, error) {
	r := newRequest(http.MethodPost, "/api/org/{orgId}/resources/create-cost-estimate")
	r.setPath("orgId", params.OrgID)
	r.setJSONBody(params.Body)
	var out *ResourcesCreateCostEstimateResponse
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
// _Requires permission: `resources:write`._
//
// POST /api/org/{orgId}/resources/invoke-action
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
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
// the caller actually changed.
//
// POST /api/org/{orgId}/resources/update
//
// Raises on 400: Bad request
//
// Raises on 404: Not found
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
// Ignores trigger opt-ins — every channel gets the test. Fails with the Slack
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

// Create: Route alerts to a Slack channel
//
// Adds a channel, or updates the trigger opt-ins of one already added. Each
// trigger defaults to enabled.
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

// Delete: Stop routing alerts to a channel
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

// Update: Change which alerts a channel receives
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
