// github.com/Infrawrench/infrawrench-go v1.39.0 | MIT | Copyright (c) 2026 Infrawrench LLC
// https://github.com/Infrawrench/Infrawrench
//
// Generated from the Infrawrench API OpenAPI 3.1 spec (API version 1.39.0).
//
// DO NOT EDIT. Regenerate with:
//   pnpm --filter @infrawrench/web generate:sdk
//
// Internal routes are absent by construction: the generator consumes the same
// published spec that /openapi.json serves, which drops every operation
// marked x-internal.

package infrawrench

import (
	"io"
)

// Wire models.
//
// A field is a pointer whenever the wire may omit it or send null. `omitempty`
// alone would not do: it also drops false, 0 and the empty string, so a caller
// could never turn a flag off or send an explicit zero. Slices and maps keep
// their own nil instead of gaining a second one.
//
// Fields carry the spec's own property names in their json tags, so the Go
// spelling is free to follow Go's conventions.

// AcceptInvitationRequest is the `AcceptInvitationRequest` schema.
type AcceptInvitationRequest struct {
	Token string `json:"token"`
}

// AcceptInvitationResponse is the `AcceptInvitationResponse` schema.
type AcceptInvitationResponse struct {
	Organization AcceptInvitationResponseOrganization `json:"organization"`
}

// AccessDecision is the `AccessDecision` schema.
type AccessDecision struct {
	// Note: Shown on the request and in the audit log.
	Note *string `json:"note,omitempty"`
}

// AccessDecisionConflict is the `AccessDecisionConflict` schema.
type AccessDecisionConflict struct {
	Error string `json:"error"`
}

// AccessDecisionForbidden is the `AccessDecisionForbidden` schema.
type AccessDecisionForbidden struct {
	Error string `json:"error"`
	// Code: One of "self_approval", "exceeds_approver".
	Code string `json:"code"`
	// Missing: For `exceeds_approver`: the permissions the approver does not
	// hold.
	Missing []string `json:"missing,omitempty"`
}

// AccessFinding is the `AccessFinding` schema.
type AccessFinding struct {
	// ResourceID: Infrawrench resource id the finding is on.
	ResourceID string `json:"resourceId"`
	// RuleID: Which rule was raised. Half of a dismissal's key, alongside the
	// resource id. The `access-review:` prefix is reserved so these can share
	// the posture dismissal store without colliding with plugin-declared posture
	// rule ids.
	//
	// One of "access-review:stale-principal", "access-review:admin-principal",
	// "access-review:key-past-rotation", "access-review:no-recorded-owner",
	// "access-review:no-mfa".
	RuleID string `json:"ruleId"`
	Title  string `json:"title"`
	// Severity: How bad the finding is. `critical` and `high` findings ride the
	// posture alert window; `medium` and `low` are review work surfaced on the
	// access review screen and in the weekly digest only.
	//
	// One of "critical", "high", "medium", "low".
	Severity string `json:"severity"`
	// Reason: Why this principal is flagged, in a sentence.
	Reason    string          `json:"reason"`
	Principal AccessPrincipal `json:"principal"`
}

// AccessPrincipal is the `AccessPrincipal` schema.
type AccessPrincipal struct {
	// ResourceID: Infrawrench resource id.
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	DisplayName      string   `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
	// Role: What kind of identity the principal is, from the resource type's
	// `principalRole` declaration. Grouping and labels only — it is not a
	// permission model.
	//
	// One of "user", "group", "role", "service-account", "key", "binding".
	Role string `json:"role"`
	// LastUsedAt: When the principal was last used, or null when the review has
	// no evidence.
	LastUsedAt        *string `json:"lastUsedAt"`
	DaysSinceLastUsed *int64  `json:"daysSinceLastUsed"`
	// Activity: What could be established about the principal's last use.
	// `unknown` means the resource type declares no last-used field, or the
	// provider stored nothing parseable — it is a first-class answer and is
	// never reported as `stale`.
	//
	// One of "active", "stale", "unknown".
	Activity  string  `json:"activity"`
	CreatedAt *string `json:"createdAt"`
	AgeDays   *int64  `json:"ageDays"`
	// Admin: True when the type's declared admin indicator matched; null when
	// the type declares none.
	Admin *bool `json:"admin"`
	// MFA: Multi-factor state, only on types that declare an MFA field. Null
	// everywhere else — "not synced" is not "MFA is off".
	MFA *bool `json:"mfa"`
	// Parent: The principal this one hangs off — a key's owner, a binding's
	// subject.
	Parent *string               `json:"parent"`
	Owner  *AccessPrincipalOwner `json:"owner"`
	// RevokeActionID: The plugin action that revokes this principal, when the
	// type declares one. Dispatch it through POST /resources/invoke-action; null
	// means the provider offers no revocation Infrawrench can invoke.
	RevokeActionID *string `json:"revokeActionId"`
}

// AccessPrincipalOwner: Who owns the resource, from the resource-ownership
// record. Null when nobody is named.
//
// The API may send null in its place.
type AccessPrincipalOwner struct {
	// UserID: Infrawrench user id when the owner is a member.
	UserID *string `json:"userId"`
	// DisplayName: Member name, or the free-text owner label.
	DisplayName string `json:"displayName"`
	// IsLabel: True when the owner is a label rather than a routable member.
	IsLabel   bool    `json:"isLabel"`
	TicketURL *string `json:"ticketUrl"`
	Purpose   *string `json:"purpose"`
}

// AccessRequest is the `AccessRequest` schema.
type AccessRequest struct {
	ID       string  `json:"id"`
	UserID   string  `json:"userId"`
	UserName *string `json:"userName"`
	// Permissions: The permission strings being asked for.
	Permissions []string `json:"permissions"`
	Reason      string   `json:"reason"`
	// DurationMinutes: How long the elevation lasts once granted.
	DurationMinutes int64 `json:"durationMinutes"`
	// Status: `pending` (awaiting a decision), `approved`, `denied`, or
	// `expired` (nobody decided in time, or the requester withdrew it). An
	// approved row is only *granting* permissions while `active` is true.
	//
	// One of "pending", "approved", "denied", "expired".
	Status string `json:"status"`
	// ExpiresAt: When an undecided request stops being decidable.
	ExpiresAt       string  `json:"expiresAt"`
	DecidedAt       *string `json:"decidedAt"`
	DecidedByUserID *string `json:"decidedByUserId"`
	DecidedByName   *string `json:"decidedByName"`
	DecisionNote    *string `json:"decisionNote"`
	GrantedAt       *string `json:"grantedAt"`
	// GrantExpiresAt: When the elevation lapses.
	GrantExpiresAt *string `json:"grantExpiresAt"`
	RevokedAt      *string `json:"revokedAt"`
	RevokedByName  *string `json:"revokedByName"`
	// Active: True when this row is granting permissions right now. Evaluated,
	// never swept — a grant stops applying the instant it lapses.
	Active    bool   `json:"active"`
	CreatedAt string `json:"createdAt"`
}

// AccessRequestCatalog is the `AccessRequestCatalog` schema.
type AccessRequestCatalog struct {
	Permissions []string `json:"permissions"`
	// Held: Permissions the caller already holds; asking for these changes
	// nothing.
	Held            []string `json:"held"`
	MinGrantMinutes int64    `json:"minGrantMinutes"`
	MaxGrantMinutes int64    `json:"maxGrantMinutes"`
}

// AccessRequestCreate is the `AccessRequestCreate` schema.
type AccessRequestCreate struct {
	Permissions     []string `json:"permissions"`
	Reason          string   `json:"reason"`
	DurationMinutes int64    `json:"durationMinutes"`
}

// AccessReviewDismissal is the `AccessReviewDismissal` schema.
type AccessReviewDismissal struct {
	ResourceID string `json:"resourceId"`
	RuleID     string `json:"ruleId"`
	// DismissedAt: When the finding was accepted.
	DismissedAt string `json:"dismissedAt"`
	// DismissedBy: Display name or email of whoever accepted it; null when
	// unknown.
	DismissedBy *string `json:"dismissedBy"`
	// Reason: The operator's note, when they left one.
	Reason *string `json:"reason"`
}

// AccessReviewDismissalCreate is the `AccessReviewDismissalCreate` schema.
type AccessReviewDismissalCreate struct {
	// ResourceID: Infrawrench resource id the finding is on.
	ResourceID string `json:"resourceId"`
	// RuleID: Which rule was raised. Half of a dismissal's key, alongside the
	// resource id. The `access-review:` prefix is reserved so these can share
	// the posture dismissal store without colliding with plugin-declared posture
	// rule ids.
	//
	// One of "access-review:stale-principal", "access-review:admin-principal",
	// "access-review:key-past-rotation", "access-review:no-recorded-owner",
	// "access-review:no-mfa".
	RuleID string `json:"ruleId"`
	// Reason: Why this finding is acceptable. Trimmed; an empty note is stored
	// as none.
	Reason *string `json:"reason,omitempty"`
}

// AccessReviewResponse is the `AccessReviewResponse` schema.
type AccessReviewResponse struct {
	// Principals: Every synced principal, by account then type then name. Never
	// filtered by dismissals — accepting a finding must not remove a principal
	// from the inventory.
	Principals []AccessPrincipal `json:"principals"`
	// Findings: Live findings, worst severity first. Dismissed findings are not
	// included.
	Findings []AccessFinding `json:"findings"`
	// TotalCount: Live finding count; dismissals excluded.
	TotalCount int64                      `json:"totalCount"`
	Counts     AccessReviewSeverityCounts `json:"counts"`
	ByRule     AccessReviewRuleCounts     `json:"byRule"`
	ByRole     AccessReviewRoleCounts     `json:"byRole"`
	// Dismissed: Findings a dismissal is currently suppressing, most recently
	// dismissed first. Only dismissals whose rule still matches appear.
	Dismissed      []DismissedAccessFinding `json:"dismissed"`
	DismissedCount int64                    `json:"dismissedCount"`
	// UnknownActivityCount: How many principals the review could establish no
	// last-use evidence for. Surfaces render this so "we found nothing" and "we
	// could not look" do not read the same.
	UnknownActivityCount int64 `json:"unknownActivityCount"`
	// StaleDays: The staleness window this review was computed against.
	StaleDays   int64  `json:"staleDays"`
	GeneratedAt string `json:"generatedAt"`
}

// AccessReviewRoleCounts is the `AccessReviewRoleCounts` schema.
type AccessReviewRoleCounts struct {
	User           int64 `json:"user"`
	Group          int64 `json:"group"`
	Role           int64 `json:"role"`
	ServiceAccount int64 `json:"service-account"`
	Key            int64 `json:"key"`
	Binding        int64 `json:"binding"`
}

// AccessReviewRuleCounts is the `AccessReviewRuleCounts` schema.
type AccessReviewRuleCounts struct {
	AccessReviewStalePrincipal  int64 `json:"access-review:stale-principal"`
	AccessReviewAdminPrincipal  int64 `json:"access-review:admin-principal"`
	AccessReviewKeyPastRotation int64 `json:"access-review:key-past-rotation"`
	AccessReviewNoRecordedOwner int64 `json:"access-review:no-recorded-owner"`
	AccessReviewNoMFA           int64 `json:"access-review:no-mfa"`
}

// AccessReviewSeverityCounts is the `AccessReviewSeverityCounts` schema.
type AccessReviewSeverityCounts struct {
	Critical int64 `json:"critical"`
	High     int64 `json:"high"`
	Medium   int64 `json:"medium"`
	Low      int64 `json:"low"`
}

// AccessRevokeConflict is the `AccessRevokeConflict` schema.
type AccessRevokeConflict struct {
	Error string `json:"error"`
}

// AccessWithdrawConflict is the `AccessWithdrawConflict` schema.
type AccessWithdrawConflict struct {
	Error string `json:"error"`
}

// Account is the `Account` schema.
type Account struct {
	ID          string `json:"id"`
	PluginID    string `json:"pluginId"`
	DisplayName string `json:"displayName"`
	// BastionID: Bastion this account's cloud-API egress is routed through.
	// `null` ⇒ direct egress.
	BastionID *string `json:"bastionId"`
	CreatedAt string  `json:"createdAt"`
}

// AccountDeleted is the `AccountDeleted` schema.
type AccountDeleted struct {
	OK                   bool  `json:"ok"`
	OrganizationsDeleted int64 `json:"organizationsDeleted"`
}

// AccountDeletionPreview is the `AccountDeletionPreview` schema.
type AccountDeletionPreview struct {
	// OrganizationsToDelete: Deleted with the account — the caller is their only
	// member.
	OrganizationsToDelete []OrganizationRef `json:"organizationsToDelete"`
	// OrganizationsToLeave: Survive; the caller's membership is removed.
	OrganizationsToLeave []OrganizationRef `json:"organizationsToLeave"`
	// Blockers: Non-empty means DELETE /api/profile will refuse until another
	// owner is promoted.
	Blockers []OwnershipBlocker `json:"blockers"`
}

// AccountDetail is the `AccountDetail` schema.
type AccountDetail struct {
	Account           AccountDetailAccount  `json:"account"`
	ResourceTypes     []ResourceTypeSummary `json:"resourceTypes"`
	PluginDisplayName string                `json:"pluginDisplayName"`
	PluginLogoSvg     string                `json:"pluginLogoSvg"`
}

// AccountTagCompliance is the `AccountTagCompliance` schema.
type AccountTagCompliance struct {
	AccountID      string `json:"accountId"`
	PluginID       string `json:"pluginId"`
	DisplayName    string `json:"displayName"`
	TotalResources int64  `json:"totalResources"`
	// Evaluated: Resources whose stored record exposes a tag map (the scoreable
	// set).
	Evaluated int64 `json:"evaluated"`
	Compliant int64 `json:"compliant"`
	// Score: Percent of evaluated resources carrying every required tag; null
	// when none.
	Score *int64 `json:"score"`
}

// ActiveTunnel is the `ActiveTunnel` schema.
type ActiveTunnel struct {
	LocalPort  int64  `json:"localPort"`
	SSHHost    string `json:"sshHost"`
	RemotePort int64  `json:"remotePort"`
}

// AgentClaimLookup is the `AgentClaimLookup` schema.
type AgentClaimLookup struct {
	RegistrationID   string `json:"registrationId"`
	WorkspaceName    string `json:"workspaceName"`
	TrialExpiresInMs *int64 `json:"trialExpiresInMs"`
	// MergeTargets: Organizations this user may merge the workspace into: ones
	// they already belong to AND hold `accounts:write` in. A merge writes cloud
	// credentials, so membership alone is not enough — the confirm route
	// enforces the same rule.
	MergeTargets []AgentClaimMergeTarget `json:"mergeTargets"`
}

// AgentClaimLookupRequest is the `AgentClaimLookupRequest` schema.
type AgentClaimLookupRequest struct {
	// Code: The `user_code` the agent showed its user.
	Code string `json:"code"`
}

// AgentClaimMergeTarget is the `AgentClaimMergeTarget` schema.
type AgentClaimMergeTarget struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

// AgentClaimRequest is the `AgentClaimRequest` schema.
type AgentClaimRequest struct {
	Code string `json:"code"`
	// Mode: `adopt` keeps the workspace as its own organization and stops the
	// clock. `merge` moves its cloud accounts into an organization you already
	// belong to and destroys the trial. Defaults to `adopt`.
	//
	// One of "adopt", "merge".
	Mode *string `json:"mode,omitempty"`
	// TargetOrganizationID: Required when `mode` is merge.
	TargetOrganizationID *string `json:"targetOrganizationId,omitempty"`
	// MoveHistory: Merge only: also re-parent the trial's metrics and cost
	// history. Off by default — it changes numbers the target organization may
	// already be reporting on. Needs `costs:write`.
	MoveHistory *bool `json:"moveHistory,omitempty"`
}

// AgentClaimResult is the `AgentClaimResult` schema.
type AgentClaimResult struct {
	// OrganizationID: The organization the agent acts in from now on.
	OrganizationID string `json:"organizationId"`
	// Mode: One of "adopt", "merge".
	Mode          string `json:"mode"`
	AccountsMoved int64  `json:"accountsMoved"`
	HistoryMoved  bool   `json:"historyMoved"`
}

// AgentClaimStarted is the `AgentClaimStarted` schema.
type AgentClaimStarted struct {
	// UserCode: Formatted as `XXXX-XXXX`. Show it to the user alongside
	// `verification_uri`.
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	// VerificationURIComplete: The verification page with the code pre-filled.
	// Convenient, but it puts a live bearer secret in a URL — prefer
	// `verification_uri` plus the code shown separately.
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresAt               string `json:"expires_at"`
	// Interval: Minimum seconds between status polls.
	Interval int64 `json:"interval"`
}

// AgentIdentity is the `AgentIdentity` schema.
type AgentIdentity struct {
	RegistrationID string `json:"registration_id"`
	OrganizationID string `json:"organization_id"`
	Claimed        bool   `json:"claimed"`
	// ClaimPending: A `user_code` is currently outstanding.
	ClaimPending bool `json:"claim_pending"`
	// TrialExpiresInMs: Milliseconds until deletion. Null once the workspace is
	// claimed.
	TrialExpiresInMs *int64 `json:"trial_expires_in_ms"`
}

// AgentRegisterRequest is the `AgentRegisterRequest` schema.
type AgentRegisterRequest struct {
	// Label: Short name for the workspace, shown to the user who claims it.
	Label *string `json:"label,omitempty"`
}

// AgentRegistration is the `AgentRegistration` schema.
type AgentRegistration struct {
	ID    string  `json:"id"`
	Label *string `json:"label"`
	// Kind: One of "anonymous", "service_auth".
	Kind string `json:"kind"`
	// Prefix: First 8 characters of the credential.
	Prefix          *string `json:"prefix"`
	ClaimedAt       *string `json:"claimedAt"`
	ClaimedByUserID *string `json:"claimedByUserId"`
	ClaimedByEmail  *string `json:"claimedByEmail"`
	LastSeenAt      *string `json:"lastSeenAt"`
	RevokedAt       *string `json:"revokedAt"`
	CreatedAt       string  `json:"createdAt"`
}

// AgentRevoked is the `AgentRevoked` schema.
type AgentRevoked struct {
	OK bool `json:"ok"`
	// Revoked: False when the registration was already revoked. The request
	// still succeeds — revocation is idempotent — but nothing changed.
	Revoked bool `json:"revoked"`
}

// AgentSession is the `AgentSession` schema.
type AgentSession struct {
	ID             string `json:"id"`
	Repo           string `json:"repo"`
	ProjectName    string `json:"projectName"`
	WorkspaceName  string `json:"workspaceName"`
	AccountID      string `json:"accountId"`
	PluginID       string `json:"pluginId"`
	ResourceTypeID string `json:"resourceTypeId"`
	// Tool: One of "codex", "claude-code".
	Tool string `json:"tool"`
	// Surface: One of "terminal", "t3-code".
	Surface    *string `json:"surface,omitempty"`
	BranchName string  `json:"branchName"`
	// Status: One of "pending", "provisioning", "setting-up", "up", "failed",
	// "stopped".
	Status       string   `json:"status"`
	VMResourceID *string  `json:"vmResourceId"`
	Logs         []string `json:"logs"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
}

// AgentSettings is the `AgentSettings` schema.
//
// The API may send null in its place.
type AgentSettings struct {
	AccountID      string `json:"accountId"`
	PluginID       string `json:"pluginId"`
	ResourceTypeID string `json:"resourceTypeId"`
	// Tool: One of "codex", "claude-code".
	Tool string `json:"tool"`
	// Surface: One of "terminal", "t3-code".
	Surface *string           `json:"surface,omitempty"`
	Fields  map[string]string `json:"fields"`
}

// AgentVMAccount is the `AgentVmAccount` schema.
//
// Spec schema: `AgentVmAccount`.
type AgentVMAccount struct {
	AccountID          string            `json:"accountId"`
	AccountName        string            `json:"accountName"`
	PluginID           string            `json:"pluginId"`
	PluginName         string            `json:"pluginName"`
	PluginLogoSvg      *string           `json:"pluginLogoSvg,omitempty"`
	ResourceTypeID     string            `json:"resourceTypeId"`
	ResourceTypeName   string            `json:"resourceTypeName"`
	DefaultUsername    string            `json:"defaultUsername"`
	DefaultFields      map[string]string `json:"defaultFields"`
	DefaultFieldLabels map[string]string `json:"defaultFieldLabels,omitempty"`
	CreateFields       []JSONObject      `json:"createFields,omitempty"`
	HiddenFieldKeys    []string          `json:"hiddenFieldKeys"`
}

// AlertCondition: One clause of a rule. A rule matches when every condition
// matches; 'or' is expressed by writing a second rule. A condition on a fact the
// alert does not carry never matches — in either direction, so `accountId notIn
// [x]` does not match an alert with no account.
type AlertCondition = any

// AlertDelivery is the `AlertDelivery` schema.
type AlertDelivery struct {
	ID       string        `json:"id"`
	Trigger  AlertTrigger  `json:"trigger"`
	Severity AlertSeverity `json:"severity"`
	Title    string        `json:"title"`
	Body     string        `json:"body"`
	RuleName *string       `json:"ruleName"`
	// State: One of "held", "awaiting_ack", "sent", "acknowledged", "escalated",
	// "expired".
	State     string `json:"state"`
	CreatedAt string `json:"createdAt"`
	// DeliverAfter: When a quiet-hours hold is released.
	DeliverAfter *string `json:"deliverAfter"`
	// EscalateAt: When an unacknowledged alert escalates.
	EscalateAt           *string `json:"escalateAt"`
	AcknowledgedAt       *string `json:"acknowledgedAt"`
	AcknowledgedByUserID *string `json:"acknowledgedByUserId"`
}

// AlertDestination: One place a matched alert goes. `push` reaches the
// organization's phones, still filtered by each member's own mutes — an
// organization rule decides whether the org is told, a member decides whether
// their phone rings.
//
// `on-call` resolves to one person at delivery time, so a rule reading "database
// alerts → whoever is on call" needs no edit at handover. A rotation that
// resolves to nobody — disabled, empty, not yet started — contributes nobody and
// the rule's **other** destinations still deliver: an alert lost to a
// misconfigured rotation would be the worst outcome the feature could have.
type AlertDestination = any

// AlertRule is the `AlertRule` schema.
type AlertRule struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	// Position: Ascending evaluation order
	Position   int64            `json:"position"`
	Conditions []AlertCondition `json:"conditions"`
	// Destinations: Empty is legal and meaningful: an enabled rule with no
	// destinations swallows matching alerts and shadows the rules below it.
	Destinations []AlertDestination `json:"destinations"`
	// ContinueOnMatch: False (the default) makes the list first-match-wins,
	// which is what lets a narrow rule sit above a broad one. True makes the
	// rule a tee that copies without shadowing.
	ContinueOnMatch bool              `json:"continueOnMatch"`
	QuietHours      *QuietHours       `json:"quietHours"`
	Escalation      *EscalationPolicy `json:"escalation"`
}

// AlertRuleInput is the `AlertRuleInput` schema.
type AlertRuleInput struct {
	// ID: Send the existing id to preserve it, which keeps in-flight held and
	// escalating deliveries pointing at their rule.
	ID              *string            `json:"id,omitempty"`
	Name            string             `json:"name"`
	Enabled         *bool              `json:"enabled,omitempty"`
	Conditions      []AlertCondition   `json:"conditions,omitempty"`
	Destinations    []AlertDestination `json:"destinations,omitempty"`
	ContinueOnMatch *bool              `json:"continueOnMatch,omitempty"`
	QuietHours      *QuietHours        `json:"quietHours,omitempty"`
	Escalation      *EscalationPolicy  `json:"escalation,omitempty"`
}

// AlertRulesResponse is the `AlertRulesResponse` schema.
type AlertRulesResponse struct {
	Rules []AlertRule `json:"rules"`
	// UsingDefaults: True when the organization has saved no rules and `rules`
	// is the synthesized default — everything except drift, to every connected
	// channel and to mobile push.
	UsingDefaults   bool                                `json:"usingDefaults"`
	SlackChannels   []AlertRulesResponseSlackChannels   `json:"slackChannels"`
	MsTeamsWebhooks []AlertRulesResponseMsTeamsWebhooks `json:"msTeamsWebhooks"`
	Accounts        []AlertRulesResponseAccounts        `json:"accounts"`
	// OnCallSchedules: Live on-call rotations, so the editor can offer 'whoever
	// is on call' as a destination. Disabled rotations are omitted for the same
	// reason a disconnected Slack install is: offering one would let the editor
	// build a rule that routes nowhere.
	OnCallSchedules []AlertRulesResponseOnCallSchedules `json:"onCallSchedules"`
}

// AlertSeverity: Alert severity, ordered info < warning < critical.
type AlertSeverity = string

// The values AlertSeverity takes.
const (
	AlertSeverityInfo     AlertSeverity = "info"
	AlertSeverityWarning  AlertSeverity = "warning"
	AlertSeverityCritical AlertSeverity = "critical"
)

// AlertTrigger: A kind of alert that can be routed.
type AlertTrigger = string

// The values AlertTrigger takes.
const (
	AlertTriggerSyncIncidents            AlertTrigger = "syncIncidents"
	AlertTriggerBudgetAlerts             AlertTrigger = "budgetAlerts"
	AlertTriggerAnomalyAlerts            AlertTrigger = "anomalyAlerts"
	AlertTriggerCostChangeAlerts         AlertTrigger = "costChangeAlerts"
	AlertTriggerCommitmentExpiryAlerts   AlertTrigger = "commitmentExpiryAlerts"
	AlertTriggerCommitmentIdleAlerts     AlertTrigger = "commitmentIdleAlerts"
	AlertTriggerUnitCostRegressionAlerts AlertTrigger = "unitCostRegressionAlerts"
	AlertTriggerMetricAlerts             AlertTrigger = "metricAlerts"
	AlertTriggerResourceDrift            AlertTrigger = "resourceDrift"
	AlertTriggerWorkflowPages            AlertTrigger = "workflowPages"
	AlertTriggerProviderIncidents        AlertTrigger = "providerIncidents"
	AlertTriggerExpiryAlerts             AlertTrigger = "expiryAlerts"
	AlertTriggerLogMatchAlerts           AlertTrigger = "logMatchAlerts"
	AlertTriggerPostureAlerts            AlertTrigger = "postureAlerts"
	AlertTriggerProbeAlerts              AlertTrigger = "probeAlerts"
	AlertTriggerQuotaAlerts              AlertTrigger = "quotaAlerts"
	AlertTriggerIncidentAlerts           AlertTrigger = "incidentAlerts"
	AlertTriggerWeeklyDigest             AlertTrigger = "weeklyDigest"
)

// AllocationRule is the `AllocationRule` schema.
type AllocationRule struct {
	ID           string              `json:"id"`
	CostCentreID string              `json:"costCentreId"`
	Priority     int64               `json:"priority"`
	Match        AllocationRuleMatch `json:"match"`
	CreatedAt    string              `json:"createdAt"`
	UpdatedAt    string              `json:"updatedAt"`
}

// AllocationRuleInput is the `AllocationRuleInput` schema.
type AllocationRuleInput struct {
	CostCentreID string `json:"costCentreId"`
	// Priority: Lower fires first; the first matching rule wins.
	Priority int64               `json:"priority"`
	Match    AllocationRuleMatch `json:"match"`
}

// AllocationRuleMatch: All set fields must match (AND). A rule with no fields is
// a catch-all that claims everything reaching it.
type AllocationRuleMatch struct {
	TagKey *string `json:"tagKey,omitempty"`
	// TagValue: Only meaningful with tagKey; alone, tagKey matches rows carrying
	// the key.
	TagValue  *string `json:"tagValue,omitempty"`
	AccountID *string `json:"accountId,omitempty"`
	PluginID  *string `json:"pluginId,omitempty"`
	Service   *string `json:"service,omitempty"`
}

// APIKey is the `ApiKey` schema.
//
// Spec schema: `ApiKey`.
type APIKey struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Prefix     string       `json:"prefix"`
	Scopes     []Permission `json:"scopes"`
	LastUsedAt *string      `json:"lastUsedAt"`
	ExpiresAt  *string      `json:"expiresAt"`
	RevokedAt  *string      `json:"revokedAt"`
	// LegacyHashSunsetAt: Cutover date past which a key still on the legacy
	// SHA-256 hash will be refused. Null once rehashed to HMAC.
	LegacyHashSunsetAt *string `json:"legacyHashSunsetAt"`
	// NeedsRotation: True when this key is still hashed with the legacy SHA-256
	// scheme and should be rotated before `legacyHashSunsetAt`.
	NeedsRotation bool   `json:"needsRotation"`
	CreatedAt     string `json:"createdAt"`
}

// ApplyManifestRequest is the `ApplyManifestRequest` schema.
type ApplyManifestRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceID       ResourceID  `json:"resourceId"`
	Manifest         string      `json:"manifest"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// ArtifactsListRequest is the `ArtifactsListRequest` schema.
type ArtifactsListRequest struct {
	AccountID      string     `json:"accountId"`
	ResourceID     ResourceID `json:"resourceId"`
	ResourceTypeID string     `json:"resourceTypeId"`
	PageToken      *string    `json:"pageToken,omitempty"`
	Prefix         *string    `json:"prefix,omitempty"`
}

// AssociationRequest is the `AssociationRequest` schema.
type AssociationRequest struct {
	ConsumerResourceID     ResourceID `json:"consumerResourceId"`
	ConsumerFieldKey       string     `json:"consumerFieldKey"`
	ProviderResourceID     ResourceID `json:"providerResourceId"`
	ProviderOutputKey      string     `json:"providerOutputKey"`
	ProviderPluginID       string     `json:"providerPluginId"`
	ProviderResourceTypeID string     `json:"providerResourceTypeId"`
	ProviderAccountID      string     `json:"providerAccountId"`
}

// AttachRequest is the `AttachRequest` schema.
type AttachRequest struct {
	PluginID         string     `json:"pluginId"`
	AccountID        string     `json:"accountId"`
	SourceTypeID     string     `json:"sourceTypeId"`
	SourceResourceID ResourceID `json:"sourceResourceId"`
	TargetTypeID     string     `json:"targetTypeId"`
	TargetResourceID ResourceID `json:"targetResourceId"`
}

// AuditEntry is the `AuditEntry` schema.
type AuditEntry struct {
	ID           string     `json:"id"`
	UserID       *string    `json:"userId"`
	APIKeyID     *string    `json:"apiKeyId"`
	Action       string     `json:"action"`
	EntityType   string     `json:"entityType"`
	EntityID     string     `json:"entityId"`
	Metadata     JSONObject `json:"metadata"`
	IPAddress    *string    `json:"ipAddress"`
	CreatedAt    string     `json:"createdAt"`
	UserName     *string    `json:"userName"`
	UserEmail    *string    `json:"userEmail"`
	APIKeyName   *string    `json:"apiKeyName"`
	APIKeyPrefix *string    `json:"apiKeyPrefix"`
}

// AuditResponse is the `AuditResponse` schema.
type AuditResponse struct {
	Entries []AuditEntry `json:"entries"`
	Total   int64        `json:"total"`
}

// AuthFactor is the `AuthFactor` schema.
type AuthFactor struct {
	ID string `json:"id"`
	// Type: One of "totp", "sms", "generic_otp".
	Type       string  `json:"type"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	TOTPIssuer *string `json:"totpIssuer"`
	TOTPUser   *string `json:"totpUser"`
}

// BackupCoverageResponse is the `BackupCoverageResponse` schema.
type BackupCoverageResponse struct {
	// Findings: Gaps, worst severity first.
	Findings    []BackupFinding       `json:"findings"`
	Counts      BackupSeverityCounts  `json:"counts"`
	KindCounts  BackupKindCounts      `json:"kindCounts"`
	TotalCount  int64                 `json:"totalCount"`
	Resources   []BackupCoverageRow   `json:"resources"`
	Summary     BackupCoverageSummary `json:"summary"`
	GeneratedAt string                `json:"generatedAt"`
}

// BackupCoverageRow is the `BackupCoverageRow` schema.
type BackupCoverageRow struct {
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	DisplayName      string   `json:"displayName"`
	ExternalID       *string  `json:"externalId"`
	// State: How the resource reads at a glance. `automated` means the provider
	// is taking backups we cannot enumerate, so there is a restore point but no
	// listable one. `unknown` means the resource type declares a provider-native
	// automated-backup signal but this instance's value could not be read — it
	// is unassessed, not a confirmed gap, and never produces a finding.
	//
	// One of "protected", "automated", "stale", "unknown", "unprotected".
	State string `json:"state"`
	// BackupCount: Backups in the inventory that protect this resource.
	BackupCount      int64    `json:"backupCount"`
	LatestBackupID   *string  `json:"latestBackupId"`
	LatestBackupName *string  `json:"latestBackupName"`
	LatestBackupAt   *string  `json:"latestBackupAt"`
	RpoHours         *float64 `json:"rpoHours"`
	// AutomatedBackups: Whether provider-native automated backups are on. Null
	// means the plugin syncs no signal either way — which never counts as
	// protection and never counts as a fault.
	AutomatedBackups *bool    `json:"automatedBackups"`
	RetentionDays    *float64 `json:"retentionDays"`
	// RpoPolicyID: The policy supplying `maxRpoHours` — the strictest RPO among
	// those selecting this resource. Tracked separately from the retention
	// policy because the two strictest demands routinely come from different
	// policies.
	RpoPolicyID   *string `json:"rpoPolicyId"`
	RpoPolicyName *string `json:"rpoPolicyName"`
	// RetentionPolicyID: The policy supplying `minRetentionDays`.
	RetentionPolicyID   *string `json:"retentionPolicyId"`
	RetentionPolicyName *string `json:"retentionPolicyName"`
	MaxRpoHours         *int64  `json:"maxRpoHours"`
	MinRetentionDays    *int64  `json:"minRetentionDays"`
}

// BackupCoverageSummary is the `BackupCoverageSummary` schema.
type BackupCoverageSummary struct {
	// StatefulCount: Stateful resources the plugin declarations can judge.
	StatefulCount  int64 `json:"statefulCount"`
	ProtectedCount int64 `json:"protectedCount"`
	// UnprotectedCount: Confirmed gaps. Excludes unassessed resources; this is
	// what the digest counts.
	UnprotectedCount int64 `json:"unprotectedCount"`
	// UnknownCount: Resources that could not be assessed: the type declares a
	// provider-native automated-backup signal but this instance's value was
	// absent or unrecognised. Reported separately so 'we found no gap' and 'we
	// could not tell' do not read alike.
	UnknownCount        int64 `json:"unknownCount"`
	BackupCount         int64 `json:"backupCount"`
	OrphanedBackupCount int64 `json:"orphanedBackupCount"`
	// UnattributableBackupCount: Backups whose source could not be determined —
	// the plugin syncs no source field, the field was empty, or more than one
	// resource answered to it. Reported rather than hidden: 'we found no
	// orphans' and 'we could not tell' are different answers.
	UnattributableBackupCount int64    `json:"unattributableBackupCount"`
	OrphanedGb                *float64 `json:"orphanedGb"`
	// OrphanedMonthlyCost: Null when billing data is unavailable or the orphans
	// span several currencies.
	OrphanedMonthlyCost *float64 `json:"orphanedMonthlyCost"`
	Currency            *string  `json:"currency"`
	// WorstRpoHours: Largest RPO across resources that have a datable backup at
	// all.
	WorstRpoHours *float64 `json:"worstRpoHours"`
}

// BackupFinding is the `BackupFinding` schema.
type BackupFinding struct {
	// ResourceID: Infrawrench resource id the finding is on.
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	DisplayName      string   `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
	// Kind: What the finding describes: nothing protects the resource; the
	// newest backup is older than the policy's RPO; the provider-native
	// retention window is shorter than the policy asks; or a backup whose source
	// resource no longer exists.
	//
	// One of "unprotected", "rpo-breach", "retention-below-policy",
	// "orphaned-snapshot".
	Kind string `json:"kind"`
	// Severity: How bad the gap is. Orphaned backups are always `low` — they
	// cost money, not data.
	//
	// One of "critical", "high", "medium", "low".
	Severity string `json:"severity"`
	Title    string `json:"title"`
	// Detail: Sentence explaining the gap and what would close it.
	Detail string `json:"detail"`
	// PolicyID: The policy supplying the objective this finding breaches — the
	// RPO policy for `rpo-breach`, the retention policy for
	// `retention-below-policy`. Null when no policy applies.
	PolicyID   *string `json:"policyId"`
	PolicyName *string `json:"policyName"`
	// RpoHours: Hours since the newest backup protecting the resource; null when
	// there is none.
	RpoHours *float64 `json:"rpoHours"`
	// MaxRpoHours: The policy's allowance, when one applied.
	MaxRpoHours *int64 `json:"maxRpoHours"`
	// RetentionDays: Provider-native retention window in days, when the plugin
	// syncs one.
	RetentionDays    *float64 `json:"retentionDays"`
	MinRetentionDays *int64   `json:"minRetentionDays"`
	LatestBackupID   *string  `json:"latestBackupId"`
	LatestBackupName *string  `json:"latestBackupName"`
	LatestBackupAt   *string  `json:"latestBackupAt"`
	// SizeGb: Size of an orphaned backup in GiB, when the plugin syncs one.
	SizeGb *float64 `json:"sizeGb"`
	// MonthlyCost: Trailing-30-day spend on an orphaned backup. Null means the
	// cost could not be determined — never that the backup is free.
	MonthlyCost *float64 `json:"monthlyCost"`
	Currency    *string  `json:"currency"`
}

// BackupKindCounts is the `BackupKindCounts` schema.
type BackupKindCounts struct {
	Unprotected          int64 `json:"unprotected"`
	RpoBreach            int64 `json:"rpo-breach"`
	RetentionBelowPolicy int64 `json:"retention-below-policy"`
	OrphanedSnapshot     int64 `json:"orphaned-snapshot"`
}

// BackupPolicy is the `BackupPolicy` schema.
type BackupPolicy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// ResourceTypeIDs: Resource types the policy selects; empty selects every
	// stateful type.
	ResourceTypeIDs []string `json:"resourceTypeIds"`
	// TagKey: Tag key that must be present. Matched case-insensitively.
	TagKey *string `json:"tagKey"`
	// TagValue: Required value of `tagKey`, matched exactly. Null means presence
	// is enough.
	TagValue *string `json:"tagValue"`
	// MaxRpoHours: The newest backup must be no older than this. Null means no
	// RPO demand.
	MaxRpoHours *int64 `json:"maxRpoHours"`
	// MinRetentionDays: Provider-native retention must be at least this. Null
	// means no demand.
	MinRetentionDays *int64 `json:"minRetentionDays"`
	Enabled          bool   `json:"enabled"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

// BackupPolicyCreate is the `BackupPolicyCreate` schema.
type BackupPolicyCreate struct {
	Name             string   `json:"name"`
	ResourceTypeIDs  []string `json:"resourceTypeIds,omitempty"`
	TagKey           *string  `json:"tagKey,omitempty"`
	TagValue         *string  `json:"tagValue,omitempty"`
	MaxRpoHours      *int64   `json:"maxRpoHours,omitempty"`
	MinRetentionDays *int64   `json:"minRetentionDays,omitempty"`
	Enabled          *bool    `json:"enabled,omitempty"`
}

// BackupPolicyList is the `BackupPolicyList` schema.
type BackupPolicyList struct {
	Policies []BackupPolicy `json:"policies"`
}

// BackupPolicyUpdate is the `BackupPolicyUpdate` schema.
type BackupPolicyUpdate struct {
	Name             *string  `json:"name,omitempty"`
	ResourceTypeIDs  []string `json:"resourceTypeIds,omitempty"`
	TagKey           *string  `json:"tagKey,omitempty"`
	TagValue         *string  `json:"tagValue,omitempty"`
	MaxRpoHours      *int64   `json:"maxRpoHours,omitempty"`
	MinRetentionDays *int64   `json:"minRetentionDays,omitempty"`
	Enabled          *bool    `json:"enabled,omitempty"`
}

// BackupSeverityCounts is the `BackupSeverityCounts` schema.
type BackupSeverityCounts struct {
	Critical int64 `json:"critical"`
	High     int64 `json:"high"`
	Medium   int64 `json:"medium"`
	Low      int64 `json:"low"`
}

// Bastion is the `Bastion` schema.
type Bastion struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	TokenPrefix     string        `json:"tokenPrefix"`
	AgentVersion    *string       `json:"agentVersion"`
	LastSeenAt      *string       `json:"lastSeenAt"`
	Status          BastionStatus `json:"status"`
	RevokedAt       *string       `json:"revokedAt"`
	CreatedAt       string        `json:"createdAt"`
	CreatedByUserID string        `json:"createdByUserId"`
	Connected       bool          `json:"connected"`
	AccountCount    int64         `json:"accountCount"`
}

// BastionStatus is the `BastionStatus` schema.
type BastionStatus = string

// The values BastionStatus takes.
const (
	BastionStatusPending BastionStatus = "pending"
	BastionStatusActive  BastionStatus = "active"
	BastionStatusRevoked BastionStatus = "revoked"
)

// BillingRule is the `BillingRule` schema.
type BillingRule struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	Enabled     bool                  `json:"enabled"`
	Priority    int64                 `json:"priority"`
	Match       BillingRuleMatch      `json:"match"`
	Adjustment  BillingRuleAdjustment `json:"adjustment"`
	CreatedAt   string                `json:"createdAt"`
	UpdatedAt   string                `json:"updatedAt"`
}

// BillingRuleAdjustment is the `BillingRuleAdjustment` schema.
type BillingRuleAdjustment struct {
	// Kind: `percentage` multiplies matched spend (every matching percentage
	// rule applies, so two 10% markups compound to 21%). `fixed` adds a flat
	// amount per period, pro-rated across the queried range, and is never
	// multiplied by anything. `reallocation` moves matched spend onto another
	// cost centre or account; the first matching reallocation rule wins, so a
	// row moves exactly once and the organisation's total is unchanged.
	//
	// One of "percentage", "fixed", "reallocation".
	Kind string `json:"kind"`
	// Percent: `percentage` only. Signed: +15 marks up by 15%, -10 discounts by
	// 10%. Bounded below at -100 because a discount larger than the cost would
	// turn spend into income.
	Percent *float64 `json:"percent,omitempty"`
	// Amount: `fixed` only, in the major unit of `currency`, per `period`.
	Amount   *float64 `json:"amount,omitempty"`
	Currency *string  `json:"currency,omitempty"`
	// Period: `fixed` only. A monthly amount is pro-rated across partial months:
	// a range covering ten days of a 31-day month contributes 10/31 of it.
	//
	// One of "daily", "monthly".
	Period *string `json:"period,omitempty"`
	// TargetKind: Required on `reallocation`, optional on `fixed` (where the
	// flat charge is booked), never set on `percentage`.
	//
	// One of "cost_centre", "account".
	TargetKind *string `json:"targetKind,omitempty"`
	TargetID   *string `json:"targetId,omitempty"`
}

// BillingRuleInput is the `BillingRuleInput` schema.
type BillingRuleInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Enabled: Disabled rules are kept and excluded from every query. Switching
	// a markup off for a quarter is an edit, not a delete.
	Enabled *bool `json:"enabled,omitempty"`
	// Priority: Lower evaluates first. Percentage rules all apply regardless of
	// order (multiplication commutes); reallocation is first-match-wins, so
	// priority decides which one moves a row.
	Priority   int64                 `json:"priority"`
	Match      BillingRuleMatch      `json:"match"`
	Adjustment BillingRuleAdjustment `json:"adjustment"`
}

// BillingRuleMatch: All set fields must match (AND); a rule with no fields
// matches all spend. The same vocabulary allocation rules use, plus chargeType.
type BillingRuleMatch struct {
	TagKey *string `json:"tagKey,omitempty"`
	// TagValue: Only meaningful with tagKey; alone, tagKey matches rows carrying
	// the key.
	TagValue  *string `json:"tagValue,omitempty"`
	AccountID *string `json:"accountId,omitempty"`
	PluginID  *string `json:"pluginId,omitempty"`
	Service   *string `json:"service,omitempty"`
	// ChargeType: Narrow to one kind of charge. A markup that recovers overhead
	// usually should not apply to credits, refunds or commitment purchases, and
	// this is how that is expressed.
	//
	// One of "usage", "commitment_covered_usage", "commitment_fee",
	// "commitment_discount", "credit", "tax", "refund", "adjustment", "support",
	// "other".
	ChargeType *string `json:"chargeType,omitempty"`
}

// BillingStatus is the `BillingStatus` schema.
type BillingStatus struct {
	// Complimentary: Platform-granted complimentary access: all paid perks,
	// uncapped AI chat, never billed.
	Complimentary bool           `json:"complimentary"`
	Subscription  *Subscription  `json:"subscription"`
	Capacity      CapacityStatus `json:"capacity"`
}

// BlastRadiusDependant is the `BlastRadiusDependant` schema.
type BlastRadiusDependant struct {
	Node *BlastRadiusNode `json:"node"`
	// Depth: Shortest hop count from the resource: 1 is a direct dependant, 2 or
	// more reached it through something else. The resource itself is never
	// listed.
	Depth int64 `json:"depth"`
	// Via: How a direct dependant reaches the resource. Absent for transitive
	// dependants, whose path is several edges and has no single caption.
	Via *BlastRadiusDependantVia `json:"via,omitempty"`
}

// BlastRadiusFlowPeer is the `BlastRadiusFlowPeer` schema.
type BlastRadiusFlowPeer struct {
	// Ref: The peer's flow ref — a provider resource id, or a class token like
	// `internet`.
	Ref   string `json:"ref"`
	Label string `json:"label"`
	// Direction: Relative to the resource being deleted, not to the row the
	// provider captured.
	//
	// One of "egress", "ingress".
	Direction string `json:"direction"`
	// Scope: The boundary the traffic crossed.
	Scope         string  `json:"scope"`
	Bytes         float64 `json:"bytes"`
	EstimatedCost float64 `json:"estimatedCost"`
	Currency      string  `json:"currency"`
	// Days: Days in the window this peer appeared on — a spike versus a standing
	// flow.
	Days       int64       `json:"days"`
	ResourceID *ResourceID `json:"resourceId"`
}

// BlastRadiusGap is the `BlastRadiusGap` schema.
type BlastRadiusGap struct {
	// Kind: One of "network-flows", "dependency-graph", "references",
	// "workflow-source", "custom-graph-source".
	Kind string `json:"kind"`
	// Reason: A full sentence, written to be rendered verbatim to the person
	// deleting.
	Reason string `json:"reason"`
}

// BlastRadiusNode: The resource itself, when it participates in the dependency
// graph.
//
// The API may send null in its place.
type BlastRadiusNode struct {
	ID                ResourceID `json:"id"`
	DisplayName       string     `json:"displayName"`
	PluginID          string     `json:"pluginId"`
	PluginDisplayName string     `json:"pluginDisplayName"`
	// PluginLogoSvg: Inline SVG markup; may be empty.
	PluginLogoSvg     string `json:"pluginLogoSvg"`
	ResourceTypeID    string `json:"resourceTypeId"`
	ResourceTypeLabel string `json:"resourceTypeLabel"`
	AccountID         string `json:"accountId"`
	AccountName       string `json:"accountName"`
}

// BlastRadiusReference is the `BlastRadiusReference` schema.
type BlastRadiusReference struct {
	// Kind: What kind of object names the resource.
	//
	// One of "dashboard", "custom-graph", "probe", "status-page",
	// "metric-alert", "lease", "schedule", "workflow", "log-query", "owner".
	Kind string `json:"kind"`
	// ID: The referring object's own id.
	ID   string `json:"id"`
	Name string `json:"name"`
	// Detail: One extra clause of context.
	Detail *string `json:"detail,omitempty"`
	// UserFacing: Set when the reference is visible outside the organization — a
	// published status page component, or the probe behind one. Any user-facing
	// reference makes the report high severity on its own.
	UserFacing *bool `json:"userFacing,omitempty"`
}

// BlastRadiusReport is the `BlastRadiusReport` schema.
type BlastRadiusReport struct {
	ResourceID ResourceID       `json:"resourceId"`
	Resource   *BlastRadiusNode `json:"resource"`
	// Dependants: Affected resources, direct first then by depth.
	Dependants      []BlastRadiusDependant `json:"dependants"`
	DirectCount     int64                  `json:"directCount"`
	TransitiveCount int64                  `json:"transitiveCount"`
	// References: Objects naming the resource without depending on it,
	// user-facing ones first.
	References []BlastRadiusReference `json:"references"`
	// FlowPeers: Measured network peers over the last 14 days, heaviest first.
	// Empty when flow collection is off — see `unchecked`.
	FlowPeers []BlastRadiusFlowPeer `json:"flowPeers"`
	// FlowTotals: Totals over `flowPeers`, or null when traffic could not be
	// measured at all. Zeroed totals mean collection is on and the resource is
	// quiet; null means nobody looked.
	FlowTotals *BlastRadiusReportFlowTotals `json:"flowTotals"`
	// Unchecked: What the report could not look at. An empty `dependants` list
	// with a non-empty `unchecked` list is not a clean bill of health, and
	// surfaces must not render it as one.
	Unchecked []BlastRadiusGap `json:"unchecked"`
	// Severity: `high` for anything user-facing or five or more direct
	// dependants; `unknown` when nothing was found but something could not be
	// checked.
	//
	// One of "none", "low", "medium", "high", "unknown".
	Severity string `json:"severity"`
	// Headline: One sentence, ready to render.
	Headline string `json:"headline"`
}

// BudgetAlertEvent is the `BudgetAlertEvent` schema.
type BudgetAlertEvent struct {
	ID    string `json:"id"`
	Month string `json:"month"`
	// ThresholdType: One of "actual", "forecast".
	ThresholdType       string `json:"thresholdType"`
	ThresholdPercent    int64  `json:"thresholdPercent"`
	ActualAmountCents   int64  `json:"actualAmountCents"`
	ForecastAmountCents *int64 `json:"forecastAmountCents"`
	TriggeredAt         string `json:"triggeredAt"`
}

// BudgetCostBasis: The basis `actualCents` and `forecastCents` were measured on.
type BudgetCostBasis = string

// The values BudgetCostBasis takes.
const (
	BudgetCostBasisCash      BudgetCostBasis = "cash"
	BudgetCostBasisAmortized BudgetCostBasis = "amortized"
)

// BudgetCostFilter is the `BudgetCostFilter` schema.
type BudgetCostFilter struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// BudgetFull is the `BudgetFull` schema.
type BudgetFull struct {
	ID             string             `json:"id"`
	OrganizationID string             `json:"organizationId"`
	Name           string             `json:"name"`
	AmountCents    int64              `json:"amountCents"`
	Currency       string             `json:"currency"`
	Filters        []BudgetCostFilter `json:"filters"`
	// SavedFilterID: A saved cost filter (see /saved-cost-filters) applied by
	// reference and AND-composed with `filters` when the budget is evaluated.
	// Updates are full replaces, so omitting it on PUT clears it. A reference
	// that fails to resolve errors the budget's evaluation rather than silently
	// measuring all spend.
	SavedFilterID *string `json:"savedFilterId"`
	// ScenarioModelID: A scenario model (see /cost-scenarios) this budget's
	// **forecast** thresholds are measured against. Null — the default, and the
	// value for every budget nobody deliberately opts in — keeps them on the
	// bare trend. Opting in is per-budget on purpose: a hypothesis somebody
	// typed into a form must not silently change when real people get paged.
	// `actual` thresholds are never affected; they measure money already spent.
	// Updates are full replaces, so omitting it on PUT clears the opt-in.
	ScenarioModelID *string           `json:"scenarioModelId"`
	Thresholds      []BudgetThreshold `json:"thresholds"`
	CostBasis       BudgetCostBasis   `json:"costBasis"`
	// UseAdjustedSpend: Measure this budget against billing-rule-adjusted spend
	// — the internal figure — instead of what the providers charged. False by
	// default, and for every budget nobody opted in. The default is a deliberate
	// refusal: a markup is organisation policy and a budget threshold pages a
	// real person, so adding one settings row must not be able to move every
	// on-call rota at once. Unlike a scenario this affects `actual` thresholds
	// too — an opted-in budget is measuring the internal number, and
	// month-to-date internal spend is as marked up as the forecast is. The alert
	// body says the figure is adjusted and names the collected one. Updates are
	// full replaces, so omitting it on PUT clears the opt-in.
	UseAdjustedSpend bool    `json:"useAdjustedSpend"`
	CreatedByUserID  *string `json:"createdByUserId"`
	DeletedAt        *string `json:"deletedAt"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}

// BudgetInput is the `BudgetInput` schema.
type BudgetInput struct {
	Name        string             `json:"name"`
	AmountCents int64              `json:"amountCents"`
	Currency    *string            `json:"currency,omitempty"`
	Filters     []BudgetCostFilter `json:"filters,omitempty"`
	// SavedFilterID: A saved cost filter (see /saved-cost-filters) applied by
	// reference and AND-composed with `filters` when the budget is evaluated.
	// Updates are full replaces, so omitting it on PUT clears it. A reference
	// that fails to resolve errors the budget's evaluation rather than silently
	// measuring all spend.
	SavedFilterID *string `json:"savedFilterId,omitempty"`
	// ScenarioModelID: A scenario model (see /cost-scenarios) this budget's
	// **forecast** thresholds are measured against. Null — the default, and the
	// value for every budget nobody deliberately opts in — keeps them on the
	// bare trend. Opting in is per-budget on purpose: a hypothesis somebody
	// typed into a form must not silently change when real people get paged.
	// `actual` thresholds are never affected; they measure money already spent.
	// Updates are full replaces, so omitting it on PUT clears the opt-in.
	ScenarioModelID *string           `json:"scenarioModelId,omitempty"`
	Thresholds      []BudgetThreshold `json:"thresholds"`
	CostBasis       *BudgetCostBasis  `json:"costBasis,omitempty"`
	// UseAdjustedSpend: Measure this budget against billing-rule-adjusted spend
	// — the internal figure — instead of what the providers charged. False by
	// default, and for every budget nobody opted in. The default is a deliberate
	// refusal: a markup is organisation policy and a budget threshold pages a
	// real person, so adding one settings row must not be able to move every
	// on-call rota at once. Unlike a scenario this affects `actual` thresholds
	// too — an opted-in budget is measuring the internal number, and
	// month-to-date internal spend is as marked up as the forecast is. The alert
	// body says the figure is adjusted and names the collected one. Updates are
	// full replaces, so omitting it on PUT clears the opt-in.
	UseAdjustedSpend *bool `json:"useAdjustedSpend,omitempty"`
}

// BudgetThreshold is the `BudgetThreshold` schema.
type BudgetThreshold struct {
	// Type: One of "actual", "forecast".
	Type    string `json:"type"`
	Percent int64  `json:"percent"`
}

// BudgetWithStatus is the `BudgetWithStatus` schema.
type BudgetWithStatus struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	AmountCents int64              `json:"amountCents"`
	Currency    string             `json:"currency"`
	Filters     []BudgetCostFilter `json:"filters"`
	Thresholds  []BudgetThreshold  `json:"thresholds"`
	CostBasis   BudgetCostBasis    `json:"costBasis"`
	// SavedFilterID: A saved cost filter (see /saved-cost-filters) applied by
	// reference and AND-composed with `filters` when the budget is evaluated.
	// Updates are full replaces, so omitting it on PUT clears it. A reference
	// that fails to resolve errors the budget's evaluation rather than silently
	// measuring all spend.
	SavedFilterID *string `json:"savedFilterId"`
	// ScenarioModelID: A scenario model (see /cost-scenarios) this budget's
	// **forecast** thresholds are measured against. Null — the default, and the
	// value for every budget nobody deliberately opts in — keeps them on the
	// bare trend. Opting in is per-budget on purpose: a hypothesis somebody
	// typed into a form must not silently change when real people get paged.
	// `actual` thresholds are never affected; they measure money already spent.
	// Updates are full replaces, so omitting it on PUT clears the opt-in.
	ScenarioModelID *string `json:"scenarioModelId"`
	// ScenarioModelName: The opted-into model's name, so a card can say whose
	// assumptions are in the number.
	ScenarioModelName *string `json:"scenarioModelName"`
	// UseAdjustedSpend: Measure this budget against billing-rule-adjusted spend
	// — the internal figure — instead of what the providers charged. False by
	// default, and for every budget nobody opted in. The default is a deliberate
	// refusal: a markup is organisation policy and a budget threshold pages a
	// real person, so adding one settings row must not be able to move every
	// on-call rota at once. Unlike a scenario this affects `actual` thresholds
	// too — an opted-in budget is measuring the internal number, and
	// month-to-date internal spend is as marked up as the forecast is. The alert
	// body says the figure is adjusted and names the collected one. Updates are
	// full replaces, so omitting it on PUT clears the opt-in.
	UseAdjustedSpend bool `json:"useAdjustedSpend"`
	// RawActualCents: Month-to-date **collected** spend, non-null only for a
	// budget measuring adjusted spend. Null on an unadjusted budget rather than
	// a copy of `actualCents`: "there is no separate collected figure because
	// this one is it" and "the collected figure happens to equal the adjusted
	// one" are different facts, and captioning every budget in the organisation
	// would make the adjusted ones invisible.
	RawActualCents *int64 `json:"rawActualCents"`
	Month          string `json:"month"`
	ActualCents    int64  `json:"actualCents"`
	// ForecastCents: The **unadjusted trend** forecast, whether or not a
	// scenario is applied — so both numbers are always comparable.
	ForecastCents *int64 `json:"forecastCents"`
	// ScenarioForecastCents: The scenario-adjusted month forecast, set only for
	// a budget that opted into a model, and the number its forecast thresholds
	// are judged against. Null means the thresholds used `forecastCents`.
	ScenarioForecastCents *int64                               `json:"scenarioForecastCents"`
	CurrentMonthEvents    []BudgetWithStatusCurrentMonthEvents `json:"currentMonthEvents"`
	Placements            []BudgetWithStatusPlacements         `json:"placements"`
}

// BusinessMetric is the `BusinessMetric` schema.
type BusinessMetric struct {
	ID              string                    `json:"id"`
	Key             string                    `json:"key"`
	Name            string                    `json:"name"`
	Unit            string                    `json:"unit"`
	Description     *string                   `json:"description"`
	Kind            BusinessMetricKind        `json:"kind"`
	Currency        *string                   `json:"currency"`
	CostScope       []BusinessMetricScopeTerm `json:"costScope"`
	SavedFilterID   *string                   `json:"savedFilterId"`
	CreatedByUserID *string                   `json:"createdByUserId"`
	CreatedAt       string                    `json:"createdAt"`
	UpdatedAt       string                    `json:"updatedAt"`
	Coverage        *BusinessMetricCoverage   `json:"coverage"`
}

// BusinessMetricCoverage: Null when the metric has no values at all — not an
// error, but every unit-cost chart drawn from it is one continuous gap.
//
// The API may send null in its place.
type BusinessMetricCoverage struct {
	// FirstDay: Earliest reported day, YYYY-MM-DD.
	FirstDay string `json:"firstDay"`
	LastDay  string `json:"lastDay"`
	// ReportedDays: Days carrying a value — compare against the span to spot a
	// sparse series.
	ReportedDays int64 `json:"reportedDays"`
}

// BusinessMetricInput is the `BusinessMetricInput` schema.
type BusinessMetricInput struct {
	// Key: Stable lowercase slug (letters, digits, `_ . -`) that workflows and
	// the CLI address the metric by. Unique per organization among live metrics,
	// and independent of `name` so a rename never breaks a running job.
	Key  string `json:"key"`
	Name string `json:"name"`
	// Unit: Singular unit label used for display — the noun in "USD per
	// customer".
	Unit        string             `json:"unit"`
	Description *string            `json:"description,omitempty"`
	Kind        BusinessMetricKind `json:"kind"`
	// Currency: ISO-4217 code. **Required when `kind` is `currency`, and
	// rejected otherwise** — a revenue metric with no currency cannot have
	// margin computed against it, and a count metric carrying one would suggest
	// its numbers are money when they are requests.
	Currency *string `json:"currency,omitempty"`
	// CostScope: The spend this metric divides, in the same filter vocabulary
	// cost graphs and budgets use. Empty (the default) is all of the
	// organization's spend. A unit-cost query may narrow this further but can
	// never widen it: the scope is part of what the metric means, and a caller
	// who could drop it would be answering a different question under the same
	// name.
	CostScope []BusinessMetricScopeTerm `json:"costScope,omitempty"`
	// SavedFilterID: A saved cost filter AND-composed with `costScope`, resolved
	// server-side at query time. A reference that fails to resolve errors the
	// unit-cost query rather than silently widening the numerator to all spend.
	SavedFilterID *string `json:"savedFilterId,omitempty"`
}

// BusinessMetricKind: What the metric's numbers are. `count` is a unit-less
// quantity (customers, requests, GB) and supports unit cost only. `currency` is
// money the business took in, denominated in the metric's own `currency`, and is
// the only kind margin can be computed against — `(revenue − cost) ÷ revenue`
// subtracts money from money and is undefined otherwise.
type BusinessMetricKind = string

// The values BusinessMetricKind takes.
const (
	BusinessMetricKindCount    BusinessMetricKind = "count"
	BusinessMetricKindCurrency BusinessMetricKind = "currency"
)

// BusinessMetricScopeTerm is the `BusinessMetricScopeTerm` schema.
type BusinessMetricScopeTerm struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// BusinessMetricValue is the `BusinessMetricValue` schema.
type BusinessMetricValue struct {
	// Day: UTC day, YYYY-MM-DD.
	Day   string  `json:"day"`
	Value float64 `json:"value"`
	// Source: One of "api", "workflow".
	Source    string `json:"source"`
	UpdatedAt string `json:"updatedAt"`
}

// BusinessMetricValuesInput is the `BusinessMetricValuesInput` schema.
type BusinessMetricValuesInput struct {
	// Values: Days to report. **Re-reporting a day restates it rather than
	// adding to it**, so an unattended nightly job is safe to retry — an
	// accumulating write would double every number the first time the job
	// re-ran. A batch naming the same day twice keeps the last value, applying
	// the same rule within a batch that restatement applies between them.
	Values []BusinessMetricValuesInputValues `json:"values"`
}

// CalendarEvent is the `CalendarEvent` schema.
type CalendarEvent struct {
	// ID: Stable across renders for the same underlying thing, because it
	// becomes the iCalendar UID. Recurring sources (sleep windows, cron runs)
	// key it by occurrence.
	ID string `json:"id"`
	// Kind: Which of the organization's own records the event was projected
	// from. The kinds are sources rather than a severity taxonomy: a reader
	// scanning a month wants to know that one bar is a freeze and another is a
	// certificate.
	//
	// One of "change-freeze", "sleep-schedule", "expiry", "commitment-expiry",
	// "workflow-schedule", "incident".
	Kind   string  `json:"kind"`
	Title  string  `json:"title"`
	Detail *string `json:"detail"`
	// StartsAt: Clamped to the requested window's lower bound when the
	// underlying span began earlier; `openEnded` says so.
	StartsAt string `json:"startsAt"`
	// EndsAt: Null means a point in time — a deadline, a scheduled run — or a
	// span whose end is not known. `openEnded` distinguishes the two.
	EndsAt *string `json:"endsAt"`
	// OpenEnded: The span continues past an edge of the window, or has no
	// declared end at all (a freeze held until further notice, an unresolved
	// incident).
	OpenEnded bool `json:"openEnded"`
	// AllDay: The event is meaningful only to the day — a deadline read off a
	// date field. Rendering such a thing at the provider's stored midnight would
	// be false precision.
	AllDay bool `json:"allDay"`
	// Severity: One of "critical", "warning", "info".
	Severity string            `json:"severity"`
	Link     CalendarEventLink `json:"link"`
}

// CalendarEventLink: Where opening the event should go — a hint rather than a
// URL, because each surface addresses its own pages differently.
//
// The API may send null in its place.
type CalendarEventLink = any

// CalendarResponse is the `CalendarResponse` schema.
type CalendarResponse struct {
	// Events: Soonest first; longer spans before shorter ones.
	Events []CalendarEvent `json:"events"`
	From   string          `json:"from"`
	To     string          `json:"to"`
	// EmptyKinds: Kinds that were asked for and produced no events in this
	// window.
	EmptyKinds []string `json:"emptyKinds"`
	// FailedKinds: Sources that threw. Reported rather than swallowed: 'nothing
	// scheduled' and 'we could not read it' are different answers, and one
	// failing source must not empty the page.
	FailedKinds []string `json:"failedKinds"`
	GeneratedAt string   `json:"generatedAt"`
}

// CalendarSubscription is the `CalendarSubscription` schema.
type CalendarSubscription struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Kinds: Kinds the feed carries. Empty means every kind, including ones
	// added later.
	Kinds []string `json:"kinds"`
	// URL: The subscription URL, returned **only** by the create call — the
	// token it contains is stored hashed and cannot be shown again. Lose it and
	// mint a new feed.
	URL       *string `json:"url,omitempty"`
	CreatedAt string  `json:"createdAt"`
	// LastAccessedAt: Last fetch, written at most hourly. Its purpose is
	// answering 'is anyone still using this?' before revoking, which an hour of
	// staleness cannot change.
	LastAccessedAt *string `json:"lastAccessedAt"`
	RevokedAt      *string `json:"revokedAt"`
}

// CalendarSubscriptionCreate is the `CalendarSubscriptionCreate` schema.
type CalendarSubscriptionCreate struct {
	Name  string   `json:"name"`
	Kinds []string `json:"kinds,omitempty"`
}

// CalendarSubscriptionList is the `CalendarSubscriptionList` schema.
type CalendarSubscriptionList struct {
	Subscriptions []CalendarSubscription `json:"subscriptions"`
}

// CapacityCheckoutRequest is the `CapacityCheckoutRequest` schema.
type CapacityCheckoutRequest struct {
	// Quantity: Slots to buy. Defaults to 1. The buyer can still adjust it in
	// Checkout.
	Quantity *int64 `json:"quantity,omitempty"`
}

// CapacitySlot is the `CapacitySlot` schema.
type CapacitySlot struct {
	ID string `json:"id"`
	// Quantity: Seats this purchase grants for the whole of its term.
	Quantity int64 `json:"quantity"`
	// Status: A slot is only granting capacity when it is `active` AND
	// `expiresAt` is still in the future.
	//
	// One of "active", "refunded".
	Status          string `json:"status"`
	StartsAt        string `json:"startsAt"`
	ExpiresAt       string `json:"expiresAt"`
	TermMonths      int64  `json:"termMonths"`
	AmountPaidCents *int64 `json:"amountPaidCents"`
}

// CapacityStatus is the `CapacityStatus` schema.
type CapacityStatus struct {
	// Purchasable: False when this deployment has no one-time capacity price
	// configured; the purchase route returns 503 and clients should hide the
	// offer.
	Purchasable bool  `json:"purchasable"`
	TermMonths  int64 `json:"termMonths"`
	// PriceUsd: List price of one slot in whole dollars, for display copy.
	PriceUsd int64 `json:"priceUsd"`
	// Seats: Seats from slots still inside their term, excluding lapsed and
	// refunded. ADDITIONAL to `subscription.seatCount` — an org's capacity is
	// the two summed, and an org can hold slots with no subscription at all.
	Seats int64 `json:"seats"`
	// Slots: Every purchase ever made, newest first, including lapsed and
	// refunded.
	Slots []CapacitySlot `json:"slots"`
}

// ChangeCostBasis: Which charge-type basis both windows are read on. `cash` (the
// default) is what the provider charged on the day it charged it; `amortized`
// spreads a commitment's up-front fee across the term it buys. It is echoed on
// every response because a delta whose basis is unstated is unreadable — an
// amortized 'after' against a cash 'before' looks exactly like a saving.
type ChangeCostBasis = string

// The values ChangeCostBasis takes.
const (
	ChangeCostBasisCash      ChangeCostBasis = "cash"
	ChangeCostBasisAmortized ChangeCostBasis = "amortized"
)

// ChangeCostImpact is the `ChangeCostImpact` schema.
type ChangeCostImpact struct {
	Status    ChangeCostImpactStatus `json:"status"`
	CostBasis ChangeCostBasis        `json:"costBasis"`
	// WindowDays: The half-window that was requested.
	WindowDays int64 `json:"windowDays"`
	// EffectiveWindowDays: The half-window the data supported. Clamped
	// symmetrically, so both means always average the same number of days.
	EffectiveWindowDays int64 `json:"effectiveWindowDays"`
	// EventDay: UTC day the change landed on. Excluded from both windows — it is
	// a mixed day.
	EventDay   string                     `json:"eventDay"`
	Before     *ChangeCostImpactWindow    `json:"before"`
	After      *ChangeCostImpactWindow    `json:"after"`
	Series     []ChangeCostImpactSeries   `json:"series"`
	Confidence ChangeCostImpactConfidence `json:"confidence"`
	Reasons    []ChangeCostImpactReason   `json:"reasons"`
	// OverlappingChanges: Other recorded changes to the same resource inside the
	// window. A delta is correlation, never causation; this is the number that
	// says how much else was going on.
	OverlappingChanges int64 `json:"overlappingChanges"`
}

// ChangeCostImpactAnnotationRequest is the `ChangeCostImpactAnnotationRequest`
// schema.
type ChangeCostImpactAnnotationRequest struct {
	// SubjectKind: One of "change", "deployment".
	SubjectKind string           `json:"subjectKind"`
	SubjectID   string           `json:"subjectId"`
	WindowDays  *int64           `json:"windowDays,omitempty"`
	CostBasis   *ChangeCostBasis `json:"costBasis,omitempty"`
}

// ChangeCostImpactAnnotationResponse is the `ChangeCostImpactAnnotationResponse`
// schema.
type ChangeCostImpactAnnotationResponse struct {
	AnnotationID string           `json:"annotationId"`
	Text         string           `json:"text"`
	Impact       ChangeCostImpact `json:"impact"`
}

// ChangeCostImpactConfidence: How much the delta is worth believing. Derived
// from the number of comparable days per side (7+ high, 4+ medium, otherwise
// low) and dropped one tier when other recorded changes touched the same
// resource inside the window.
type ChangeCostImpactConfidence = string

// The values ChangeCostImpactConfidence takes.
const (
	ChangeCostImpactConfidenceHigh   ChangeCostImpactConfidence = "high"
	ChangeCostImpactConfidenceMedium ChangeCostImpactConfidence = "medium"
	ChangeCostImpactConfidenceLow    ChangeCostImpactConfidence = "low"
	ChangeCostImpactConfidenceNone   ChangeCostImpactConfidence = "none"
)

// ChangeCostImpactEntry is the `ChangeCostImpactEntry` schema.
type ChangeCostImpactEntry struct {
	ChangeID   string           `json:"changeId"`
	ResourceID ResourceID       `json:"resourceId"`
	Impact     ChangeCostImpact `json:"impact"`
}

// ChangeCostImpactReason: Why the result reads the way it does. Every
// non-`measured` status carries at least one, and `measured` carries whatever
// lowered its confidence. `period_native_provider` is the notable one: a
// provider that dates a whole invoice period to the period's start cannot be
// read by a day-window comparison at all.
type ChangeCostImpactReason = string

// The values ChangeCostImpactReason takes.
const (
	ChangeCostImpactReasonNoCostIdentity       ChangeCostImpactReason = "no_cost_identity"
	ChangeCostImpactReasonPeriodNativeProvider ChangeCostImpactReason = "period_native_provider"
	ChangeCostImpactReasonNoCostData           ChangeCostImpactReason = "no_cost_data"
	ChangeCostImpactReasonNoCoverageBefore     ChangeCostImpactReason = "no_coverage_before"
	ChangeCostImpactReasonNoCoverageAfter      ChangeCostImpactReason = "no_coverage_after"
	ChangeCostImpactReasonShortWindow          ChangeCostImpactReason = "short_window"
	ChangeCostImpactReasonWindowClamped        ChangeCostImpactReason = "window_clamped"
	ChangeCostImpactReasonOverlappingChanges   ChangeCostImpactReason = "overlapping_changes"
)

// ChangeCostImpactSeries is the `ChangeCostImpactSeries` schema.
type ChangeCostImpactSeries struct {
	// Currency: ISO 4217 code. Currencies are never summed.
	Currency     string  `json:"currency"`
	BeforePerDay float64 `json:"beforePerDay"`
	AfterPerDay  float64 `json:"afterPerDay"`
	// DeltaPerDay: `afterPerDay - beforePerDay`. Positive means the change costs
	// more.
	DeltaPerDay float64 `json:"deltaPerDay"`
	// DeltaPercent: Null when the before window spent nothing — there is no
	// percentage.
	DeltaPercent *float64 `json:"deltaPercent"`
	BeforeTotal  float64  `json:"beforeTotal"`
	AfterTotal   float64  `json:"afterTotal"`
}

// ChangeCostImpactStatus: `measured` — both windows had collected data and the
// delta is real. `insufficient_data` — the windows exist but are too short to
// compare. `unknown` — nothing here can answer the question. **`unknown` is
// never zero**: a resource with no cost data reports that we cannot say, not
// that the change was free.
type ChangeCostImpactStatus = string

// The values ChangeCostImpactStatus takes.
const (
	ChangeCostImpactStatusMeasured         ChangeCostImpactStatus = "measured"
	ChangeCostImpactStatusInsufficientData ChangeCostImpactStatus = "insufficient_data"
	ChangeCostImpactStatusUnknown          ChangeCostImpactStatus = "unknown"
)

// ChangeCostImpactWindow is the `ChangeCostImpactWindow` schema.
//
// The API may send null in its place.
type ChangeCostImpactWindow struct {
	// From: Inclusive first UTC day, `YYYY-MM-DD`.
	From string `json:"from"`
	// To: Inclusive last UTC day.
	To string `json:"to"`
}

// ChangeCostImpactsRequest is the `ChangeCostImpactsRequest` schema.
type ChangeCostImpactsRequest struct {
	// ChangeIDs: Change ids from `GET /changes`. At most 50 — one feed page.
	ChangeIDs []string `json:"changeIds"`
	// WindowDays: Days either side of the change. Default 7; clamped
	// server-side.
	WindowDays *int64           `json:"windowDays,omitempty"`
	CostBasis  *ChangeCostBasis `json:"costBasis,omitempty"`
}

// ChangeCostImpactsResponse is the `ChangeCostImpactsResponse` schema.
type ChangeCostImpactsResponse struct {
	Impacts []ChangeCostImpactEntry `json:"impacts"`
}

// ChangeFreeze is the `ChangeFreeze` schema.
type ChangeFreeze struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Reason          *string `json:"reason"`
	StartsAt        string  `json:"startsAt"`
	EndsAt          *string `json:"endsAt"`
	Active          bool    `json:"active"`
	CreatedByUserID *string `json:"createdByUserId"`
	EndedByUserID   *string `json:"endedByUserId"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// ChangeFreezeBlocked is the `ChangeFreezeBlocked` schema.
type ChangeFreezeBlocked struct {
	Error string `json:"error"`
	// Code: One of "change_freeze_active".
	Code   string                    `json:"code"`
	Freeze ChangeFreezeBlockedFreeze `json:"freeze"`
}

// ChangeFreezeInput is the `ChangeFreezeInput` schema.
type ChangeFreezeInput struct {
	Name     string  `json:"name"`
	Reason   *string `json:"reason,omitempty"`
	StartsAt *string `json:"startsAt,omitempty"`
	EndsAt   *string `json:"endsAt,omitempty"`
}

// ChangeFreezeStatus is the `ChangeFreezeStatus` schema.
type ChangeFreezeStatus struct {
	Freeze any `json:"freeze"`
}

// ChatAskQuestionAnswer is the `ChatAskQuestionAnswer` schema.
type ChatAskQuestionAnswer struct {
	// QuestionID: Id of the question being answered.
	QuestionID string `json:"questionId"`
	// OptionID: Listed option id, or `other` when the user typed a custom value.
	OptionID *string `json:"optionId,omitempty"`
	// Text: Required for text questions and when optionId is `other`.
	Text *string `json:"text,omitempty"`
}

// ChatAskQuestionInput is the `ChatAskQuestionInput` schema.
type ChatAskQuestionInput struct {
	// Answers: One answer per question the agent asked.
	Answers []ChatAskQuestionAnswer `json:"answers"`
}

// ChatAskQuestionResult is the `ChatAskQuestionResult` schema.
type ChatAskQuestionResult struct {
	OK bool `json:"ok"`
	// AllResolved: True when every pending action and secret request on this
	// assistant message is resolved, so the caller may POST {resume: true}.
	AllResolved bool `json:"allResolved"`
}

// ChatSecretRequestResult is the `ChatSecretRequestResult` schema.
type ChatSecretRequestResult struct {
	OK          bool `json:"ok"`
	AllResolved bool `json:"allResolved"`
}

// ChildResourceRef is the `ChildResourceRef` schema.
type ChildResourceRef struct {
	ID             ResourceID `json:"id"`
	DisplayName    string     `json:"displayName"`
	ResourceTypeID string     `json:"resourceTypeId"`
	PluginID       string     `json:"pluginId"`
	AccountID      string     `json:"accountId"`
	Status         *StatusDot `json:"status,omitempty"`
	Fields         JSONObject `json:"fields,omitempty"`
}

// ChildTypeRef is the `ChildTypeRef` schema.
type ChildTypeRef struct {
	ID                string       `json:"id"`
	DisplayName       string       `json:"displayName"`
	PluralDisplayName *string      `json:"pluralDisplayName,omitempty"`
	SupportsCreate    bool         `json:"supportsCreate"`
	Fields            []JSONObject `json:"fields,omitempty"`
}

// CommitmentCoverage is the `CommitmentCoverage` schema.
type CommitmentCoverage struct {
	// Available: False when every in-scope account was excluded — 'we cannot
	// tell' reported as unavailable, never as 0%.
	Available  bool                         `json:"available"`
	Currencies []CommitmentCoverageCurrency `json:"currencies"`
	// ExcludedAccountIDs: Accounts whose plugin cannot tell usage from other
	// charge types; their rows would drag coverage down for reasons unrelated to
	// purchasing.
	ExcludedAccountIDs []string `json:"excludedAccountIds"`
}

// CommitmentCoverageCurrency is the `CommitmentCoverageCurrency` schema.
type CommitmentCoverageCurrency struct {
	Currency string `json:"currency"`
	// CoveredAmount: Usage spend on rows stamped with a commitment id.
	CoveredAmount   float64 `json:"coveredAmount"`
	UncoveredAmount float64 `json:"uncoveredAmount"`
	// UncoveredEligibleAmount: Uncovered usage in cells where a commitment
	// landed in the window — provider evidence of committability, not a
	// hand-maintained service table.
	UncoveredEligibleAmount float64 `json:"uncoveredEligibleAmount"`
	// BroadRatio: Lower bound: covered ÷ (covered + all uncovered usage).
	BroadRatio *float64 `json:"broadRatio"`
	// NarrowRatio: Upper bound: covered ÷ (covered + uncovered usage in eligible
	// cells).
	NarrowRatio *float64 `json:"narrowRatio"`
}

// CommitmentHolding is the `CommitmentHolding` schema.
type CommitmentHolding struct {
	AccountID   string   `json:"accountId"`
	AccountName string   `json:"accountName"`
	PluginID    PluginID `json:"pluginId"`
	// CommitmentID: Provider-native id — the join key against cost rows'
	// commitment dimension (an ARN where billing data carries ARNs, the bare id
	// where it does not).
	CommitmentID string `json:"commitmentId"`
	// Kind: One of "reservation", "savings_plan", "committed_use".
	Kind        string `json:"kind"`
	Description string `json:"description"`
	// Scope: Provider scope qualifier — an AZ, an instance family, 'Shared'.
	Scope *string `json:"scope"`
	// Region: Null means the commitment applies across regions (an AWS Compute
	// Savings Plan) — a real state, rendered as 'All regions', not missing data.
	Region    *string `json:"region"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
	// TermDays: Provider-reported term length — never derived from the dates,
	// which stop spanning the term once a commitment is split or merged.
	TermDays *int64 `json:"termDays"`
	// PaymentOption: One of "all_upfront", "partial_upfront", "no_upfront",
	// "monthly".
	PaymentOption *string `json:"paymentOption"`
	// Currency: Null when the provider reports no money at all for this record.
	Currency *string `json:"currency"`
	// UpfrontAmount: Null means the provider did not report a price (Azure's
	// list API reports none) — 'not reported', never rendered as 'free'.
	UpfrontAmount   *float64 `json:"upfrontAmount"`
	RecurringAmount *float64 `json:"recurringAmount"`
	// RecurringPeriod: Atomic with recurringAmount: an amount without a period
	// is a 730× ambiguity.
	//
	// One of "hour", "month".
	RecurringPeriod *string `json:"recurringPeriod"`
	// HourlyCommitmentAmount: Committed spend per hour — what utilization is
	// measured against.
	HourlyCommitmentAmount *float64 `json:"hourlyCommitmentAmount"`
	// UnitCommitments: Committed resource quantities for unit-denominated
	// commitments (GCP CUDs). A record has either this or hourlyCommitmentAmount
	// — the split decides which utilization question is even askable.
	UnitCommitments []CommitmentUnitAmount `json:"unitCommitments"`
	// State: One of "active", "expired", "queued".
	State string `json:"state"`
	// ProviderUtilization: The provider's own utilization aggregates (Azure
	// reservations only), verbatim — never blended with the derived utilization
	// below.
	ProviderUtilization []CommitmentProviderUtilization `json:"providerUtilization"`
	LastSeenAt          string                          `json:"lastSeenAt"`
	Utilization         CommitmentUtilization           `json:"utilization"`
}

// CommitmentPlanner is the `CommitmentPlanner` schema.
type CommitmentPlanner struct {
	// Available: False when the data window is under the 60-day minimum.
	Available       bool                       `json:"available"`
	WindowDayCount  int64                      `json:"windowDayCount"`
	Recommendations []CommitmentRecommendation `json:"recommendations"`
	Rejected        []CommitmentRejectedCell   `json:"rejected"`
}

// CommitmentPollFailure is the `CommitmentPollFailure` schema.
type CommitmentPollFailure struct {
	AccountID    string   `json:"accountId"`
	AccountName  string   `json:"accountName"`
	PluginID     PluginID `json:"pluginId"`
	Message      string   `json:"message"`
	FailureCount int64    `json:"failureCount"`
}

// CommitmentProviderUtilization is the `CommitmentProviderUtilization` schema.
type CommitmentProviderUtilization struct {
	// GrainDays: Trailing window the aggregate covers (1, 7, 30).
	GrainDays int64 `json:"grainDays"`
	// Percentage: Utilization percentage 0–100, exactly as the provider reports
	// it.
	Percentage float64 `json:"percentage"`
}

// CommitmentRecommendation is the `CommitmentRecommendation` schema.
type CommitmentRecommendation struct {
	PluginID PluginID `json:"pluginId"`
	Service  string   `json:"service"`
	Region   string   `json:"region"`
	Currency string   `json:"currency"`
	// RecommendedDailyCommitment: p10 of daily uncovered usage spend,
	// nearest-rank — the floor, not the average.
	RecommendedDailyCommitment  float64 `json:"recommendedDailyCommitment"`
	RecommendedHourlyCommitment float64 `json:"recommendedHourlyCommitment"`
	AnnualCommitment            float64 `json:"annualCommitment"`
	P50DailySpend               float64 `json:"p50DailySpend"`
	// SavingBasis: Published discounts are "up to" figures. `range` renders
	// "$X–$Y"; `upper_bound` renders "up to $Y" — never a bare "$Y".
	//
	// One of "range", "upper_bound".
	SavingBasis              string   `json:"savingBasis"`
	DiscountRateMin          *float64 `json:"discountRateMin,omitempty"`
	DiscountRateMax          float64  `json:"discountRateMax"`
	EstimatedAnnualSavingMin *float64 `json:"estimatedAnnualSavingMin,omitempty"`
	EstimatedAnnualSavingMax float64  `json:"estimatedAnnualSavingMax"`
	// BreakEvenUtilization: 1 − discount: below this utilization the commitment
	// loses to on-demand. Equivalently, the workload can shrink by the discount
	// before committing was a mistake.
	BreakEvenUtilization float64 `json:"breakEvenUtilization"`
	// AnnualLossIfUsageHalves: max(0, annualCommitment × (0.5 − discount)) at
	// the shallow end of the published discount — a ceiling on regret where no
	// floor rate is published.
	AnnualLossIfUsageHalves float64 `json:"annualLossIfUsageHalves"`
}

// CommitmentRejectedCell is the `CommitmentRejectedCell` schema.
type CommitmentRejectedCell struct {
	PluginID PluginID `json:"pluginId"`
	Service  string   `json:"service"`
	Region   string   `json:"region"`
	Currency string   `json:"currency"`
	// Gate: First gate the cell failed, in evaluation order — the most
	// actionable objection.
	//
	// One of "presence", "not_in_decline", "floor", "materiality".
	Gate string `json:"gate"`
}

// CommitmentUnitAmount is the `CommitmentUnitAmount` schema.
type CommitmentUnitAmount struct {
	// Unit: Provider-native unit label, untranslated — "VCPU", "MEMORY_MB",
	// "LOCAL_SSD_GB".
	Unit   string  `json:"unit"`
	Amount float64 `json:"amount"`
}

// CommitmentUtilization is the `CommitmentUtilization` schema.
type CommitmentUtilization struct {
	// Utilization: delivered ÷ obligation, unclamped (values above 1 mean spend
	// past the commitment). **Null means not measurable** — never 0, which would
	// read as 'unused'; the reason field says why.
	Utilization *float64 `json:"utilization"`
	// Reason: Why utilization is null: `unit_denominated` — the commitment is in
	// resource units (GCP CUDs) and cost rows cannot say how many ran;
	// `no_active_days` — the term does not intersect the window; `no_data_days`
	// — no cost data was collected on any active day; `unattributed_rows` — the
	// account's plugin does not stamp commitment ids onto cost rows, so
	// delivered spend would falsely read as zero.
	//
	// One of "unit_denominated", "no_active_days", "no_data_days",
	// "unattributed_rows".
	Reason *string `json:"reason,omitempty"`
	// ObligationAmount: hourlyCommitmentAmount × 24 × measuredDays, in the
	// commitment's currency.
	ObligationAmount *float64 `json:"obligationAmount"`
	DeliveredAmount  float64  `json:"deliveredAmount"`
	// ActiveDays: Days of the window the commitment was active.
	ActiveDays int64 `json:"activeDays"`
	// MeasuredDays: Active days with cost data — the only days in the
	// obligation. Counting a day the collection never ran would make a
	// fully-used plan read as under-utilized.
	MeasuredDays int64 `json:"measuredDays"`
	// MissingDays: Active days without cost data, reported rather than silently
	// counted.
	MissingDays int64 `json:"missingDays"`
	WindowDays  int64 `json:"windowDays"`
}

// CommitmentsFeed is the `CommitmentsFeed` schema.
type CommitmentsFeed struct {
	Holdings []CommitmentHolding     `json:"holdings"`
	Coverage CommitmentCoverage      `json:"coverage"`
	Planner  CommitmentPlanner       `json:"planner"`
	Failures []CommitmentPollFailure `json:"failures"`
	// PendingAccountIDs: Commitment-capable accounts never yet collected — named
	// rather than omitted.
	PendingAccountIDs     []string `json:"pendingAccountIds"`
	UtilizationWindowDays int64    `json:"utilizationWindowDays"`
	PlannerWindowDays     int64    `json:"plannerWindowDays"`
}

// ConnectEnvDeployRequest is the `ConnectEnvDeployRequest` schema.
type ConnectEnvDeployRequest struct {
	SourceAccountID      string            `json:"sourceAccountId"`
	SourceResourceID     ResourceID        `json:"sourceResourceId"`
	SourcePluginID       string            `json:"sourcePluginId"`
	SourceResourceTypeID string            `json:"sourceResourceTypeId"`
	SourceExternalID     *string           `json:"sourceExternalId,omitempty"`
	TargetSSHHost        string            `json:"targetSshHost"`
	SSHKeyID             string            `json:"sshKeyId"`
	SSHUsername          string            `json:"sshUsername"`
	TemplateID           string            `json:"templateId"`
	KeyOverrides         map[string]string `json:"keyOverrides"`
	// Format: One of "dotenv", "profile".
	Format   string `json:"format"`
	FilePath string `json:"filePath"`
	Append   bool   `json:"append"`
}

// ConnectSecretExportRequest is the `ConnectSecretExportRequest` schema.
type ConnectSecretExportRequest struct {
	SourceAccountID      string            `json:"sourceAccountId"`
	SourceResourceID     ResourceID        `json:"sourceResourceId"`
	SourcePluginID       string            `json:"sourcePluginId"`
	SourceResourceTypeID string            `json:"sourceResourceTypeId"`
	SourceExternalID     *string           `json:"sourceExternalId,omitempty"`
	TargetAccountID      string            `json:"targetAccountId"`
	TargetPluginID       string            `json:"targetPluginId"`
	TemplateID           string            `json:"templateId"`
	Namespace            string            `json:"namespace"`
	SecretName           string            `json:"secretName"`
	KeyOverrides         map[string]string `json:"keyOverrides"`
}

// ConnectTemplatesRequest is the `ConnectTemplatesRequest` schema.
type ConnectTemplatesRequest struct {
	SourcePluginID       string `json:"sourcePluginId"`
	SourceResourceTypeID string `json:"sourceResourceTypeId"`
	TargetAccountID      string `json:"targetAccountId"`
	TargetPluginID       string `json:"targetPluginId"`
}

// ConnectTemplatesResponse is the `ConnectTemplatesResponse` schema.
type ConnectTemplatesResponse struct {
	Templates               []SecretExportTemplate `json:"templates"`
	EffectiveResourceTypeID string                 `json:"effectiveResourceTypeId"`
	SupportsSecretImport    bool                   `json:"supportsSecretImport"`
	Namespaces              []string               `json:"namespaces"`
}

// CostAccountStatus is the `CostAccountStatus` schema.
type CostAccountStatus struct {
	AccountID     string   `json:"accountId"`
	PluginID      string   `json:"pluginId"`
	DisplayName   string   `json:"displayName"`
	SupportsCosts bool     `json:"supportsCosts"`
	PeriodNative  bool     `json:"periodNative"`
	Dimensions    []string `json:"dimensions"`
	// ChargeTypes: Whether this account's plugin can tell one kind of charge
	// from another. False means every row it writes is recorded as `usage` — not
	// that the provider only bills usage.
	ChargeTypes bool `json:"chargeTypes"`
	// Amortization: Whether this account's plugin reports an amortized amount
	// distinct from the cash amount. Clients offer the amortized cost basis only
	// when at least one account says yes; elsewhere the amortized view is the
	// cash numbers under another name.
	Amortization bool `json:"amortization"`
	// Estimated: Whether this account's amounts are derived by Infrawrench —
	// inventory priced against a rate card, or metered usage priced at published
	// list rates — rather than reported as billed spend. True means the series
	// cannot be reconciled against an invoice: resources deleted part-way
	// through a period are no longer in inventory to be priced, all rates are
	// list rather than negotiated, and credits, tax and refunds never appear.
	Estimated            bool    `json:"estimated"`
	CostLastPolledAt     *string `json:"costLastPolledAt"`
	CostBackfilledAt     *string `json:"costBackfilledAt"`
	CostPollFailureCount int64   `json:"costPollFailureCount"`
	// CostPollError: Last cost-collection failure for this account, cleared on
	// the next success. `helpLink` points at the provider page that fixes a
	// setup problem when the plugin can identify one (e.g. GCP's billing export
	// console).
	CostPollError *CostAccountStatusCostPollError `json:"costPollError"`
	Coverage      *CostAccountStatusCoverage      `json:"coverage"`
}

// CostAdjustmentSummary: What an adjusted answer did. Present whenever the
// request asked to be adjusted, even for an organisation with no rules — its
// absence means, and can only mean, that every figure in the response is exactly
// what the providers charged.
type CostAdjustmentSummary struct {
	// Rules: The enabled rules in force for this answer, in evaluation order.
	Rules []CostAdjustmentSummaryRules `json:"rules"`
	// RawTotals: The collected, unadjusted totals for exactly the same rows,
	// summed in the same scan. Always present on an adjusted answer — this is
	// the figure that reconciles against an invoice. Per-series raw figures are
	// deliberately not offered: after a reallocation the series are a different
	// partition of the same money.
	RawTotals map[string]float64 `json:"rawTotals"`
	// FixedTotals: Fixed-amount charges over the period, pro-rated. On a cost
	// query these are reported here and **not** folded into `totals`, which
	// stays the sum of the series; the figure an organisation reports internally
	// is the adjusted total plus this. On a showback report they are
	// additionally booked onto the cost centre the rule names.
	FixedTotals map[string]float64 `json:"fixedTotals"`
}

// CostAlert: A change-based cost alert: fires when spend on its scope moves more
// than the configured threshold versus the prior period. The third alert family
// alongside budgets (absolute monthly total) and anomaly detection (statistical
// outliers against a learned baseline).
type CostAlert struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Filters []CostAlertFilter `json:"filters"`
	// GroupBy: Per-group fan-out. Null watches the scope's one total; a
	// dimension watches each group against its own prior window, and each
	// offending group fires its own event.
	//
	// One of "provider", "account", "service", "region", "resource", "tag",
	// "charge_type", "commitment".
	GroupBy              *string             `json:"groupBy"`
	GroupByTagKey        *string             `json:"groupByTagKey"`
	Cadence              CostChangeCadence   `json:"cadence"`
	ThresholdPercent     *int64              `json:"thresholdPercent"`
	ThresholdAmountCents *int64              `json:"thresholdAmountCents"`
	Direction            CostChangeDirection `json:"direction"`
	Enabled              bool                `json:"enabled"`
	LastEvaluatedAt      *string             `json:"lastEvaluatedAt"`
	LastFiredAt          *string             `json:"lastFiredAt"`
	CreatedAt            string              `json:"createdAt"`
	UpdatedAt            string              `json:"updatedAt"`
}

// CostAlertEvent is the `CostAlertEvent` schema.
type CostAlertEvent struct {
	ID        string `json:"id"`
	AlertID   string `json:"alertId"`
	AlertName string `json:"alertName"`
	// PeriodKey: The cadence period the firing belongs to — a day, an ISO week
	// (2026-W32) or a month (2026-08). One period fires at most once per group
	// and currency.
	PeriodKey    string `json:"periodKey"`
	WindowFrom   string `json:"windowFrom"`
	WindowTo     string `json:"windowTo"`
	PreviousFrom string `json:"previousFrom"`
	PreviousTo   string `json:"previousTo"`
	// GroupKey: The offending group; empty when the alert watches one total.
	GroupKey            string `json:"groupKey"`
	Currency            string `json:"currency"`
	PreviousAmountCents int64  `json:"previousAmountCents"`
	CurrentAmountCents  int64  `json:"currentAmountCents"`
	// ChangePercent: Signed percent change. Null when the prior window had no
	// spend at all (new spend — the change is infinite); -100 when the group
	// vanished.
	ChangePercent *int64 `json:"changePercent"`
	// Direction: One of "increase", "decrease".
	Direction  string  `json:"direction"`
	FiredAt    string  `json:"firedAt"`
	NotifiedAt *string `json:"notifiedAt"`
}

// CostAlertFilter is the `CostAlertFilter` schema.
type CostAlertFilter struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// CostAlertInput is the `CostAlertInput` schema.
type CostAlertInput struct {
	Name    string            `json:"name"`
	Filters []CostAlertFilter `json:"filters,omitempty"`
	// GroupBy: Per-group fan-out. Null watches the scope's one total; a
	// dimension watches each group against its own prior window, and each
	// offending group fires its own event.
	//
	// One of "provider", "account", "service", "region", "resource", "tag",
	// "charge_type", "commitment".
	GroupBy *string `json:"groupBy,omitempty"`
	// GroupByTagKey: Required when groupBy is tag.
	GroupByTagKey *string           `json:"groupByTagKey,omitempty"`
	Cadence       CostChangeCadence `json:"cadence"`
	// ThresholdPercent: Percent of the prior window's spend the change must
	// reach. At least one of the two thresholds must be set; when both are, BOTH
	// must hold before the alert fires.
	ThresholdPercent *int64 `json:"thresholdPercent,omitempty"`
	// ThresholdAmountCents: Cents the change must reach.
	ThresholdAmountCents *int64              `json:"thresholdAmountCents,omitempty"`
	Direction            CostChangeDirection `json:"direction"`
	Enabled              *bool               `json:"enabled,omitempty"`
}

// CostAnnotation is the `CostAnnotation` schema.
type CostAnnotation struct {
	ID string `json:"id"`
	// StartDate: Inclusive first day (UTC) the note is about. Mapped to
	// whichever bucket holds it at the chart's binning — daily and cumulative
	// use the day itself, weekly the Monday that starts its week, monthly the
	// first of its month.
	StartDate string `json:"startDate"`
	// EndDate: Inclusive last day, or null for a note about a single moment. A
	// deploy is a moment; a migration is a week, and a week spelled as seven
	// notes misstates how many things happened. An end equal to the start is
	// stored as null — the same fact has one spelling.
	EndDate *string `json:"endDate"`
	Text    string  `json:"text"`
	// CostReportID: The report this note is scoped to, or null for **org-wide**.
	// Null is the useful default: an org-wide note is drawn on every cost chart,
	// because "we changed instance types" is not a fact about one report. An id
	// from another org is a 400.
	CostReportID    *string `json:"costReportId"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
	// CostAnomalyID: The detected cost anomaly this note was written to explain
	// (see POST /costs/anomalies/{anomalyId}/acknowledge), or null for a note
	// written by hand. The reverse of the anomaly's own
	// `acknowledgement.annotationId`, resolved from that same single link rather
	// than stored twice.
	CostAnomalyID *string `json:"costAnomalyId"`
}

// CostAnnotationInput is the `CostAnnotationInput` schema.
type CostAnnotationInput struct {
	// StartDate: Inclusive first day (UTC) the note is about. Mapped to
	// whichever bucket holds it at the chart's binning — daily and cumulative
	// use the day itself, weekly the Monday that starts its week, monthly the
	// first of its month.
	StartDate string `json:"startDate"`
	// EndDate: Inclusive last day, or null for a note about a single moment. A
	// deploy is a moment; a migration is a week, and a week spelled as seven
	// notes misstates how many things happened. An end equal to the start is
	// stored as null — the same fact has one spelling.
	EndDate *string `json:"endDate,omitempty"`
	Text    string  `json:"text"`
	// CostReportID: The report this note is scoped to, or null for **org-wide**.
	// Null is the useful default: an org-wide note is drawn on every cost chart,
	// because "we changed instance types" is not a fact about one report. An id
	// from another org is a 400.
	CostReportID *string `json:"costReportId,omitempty"`
}

// CostAnomaly is the `CostAnomaly` schema.
type CostAnomaly struct {
	ID string `json:"id"`
	// Day: The anomalous UTC day.
	Day string `json:"day"`
	// Kind: Which detection produced the row. `spike` is spend far above the
	// key's own trailing baseline; `new_source` is a provider or service with no
	// spend at all across the trailing window that suddenly has material spend —
	// it can never be a `spike`, since a zero baseline has no mean or deviation
	// to exceed. Rows written before new-source detection existed read as
	// `spike`.
	//
	// One of "spike", "new_source".
	Kind string `json:"kind"`
	// Dimension: One of "provider", "service".
	Dimension string `json:"dimension"`
	// DimensionKey: The dimension's value — a plugin id or a service name.
	DimensionKey string `json:"dimensionKey"`
	Currency     string `json:"currency"`
	ActualCents  int64  `json:"actualCents"`
	// BaselineCents: Mean daily spend over the trailing 28-day baseline, in
	// cents. Zero, or near it, for a `new_source` — clients must not compute a
	// percentage change from it.
	BaselineCents int64 `json:"baselineCents"`
	// ThresholdCents: The detection bar the day cleared, in cents: baseline mean + N·stddev for a `spike`, the new-source floor for a `new_source`.
	ThresholdCents int64  `json:"thresholdCents"`
	DetectedAt     string `json:"detectedAt"`
	// NotifiedAt: When the anomaly was delivered to a notification channel; null
	// when delivery failed or a recent anomaly for the same key suppressed it.
	NotifiedAt *string `json:"notifiedAt"`
	// Hints: Root-cause hints computed when the anomaly fired: human-readable
	// facts from the change timeline and audit log for the anomalous day and the
	// day before (e.g. "12 gce-instance resources appeared", a workflow run, a
	// lifted change freeze), ranked by likely relevance and capped at three.
	// Empty when nothing notable happened in the window or the anomaly predates
	// hint collection.
	Hints []string `json:"hints"`
	// Acknowledgement: Present once somebody has explained this finding, null
	// while it is still an open question. Acknowledging does not suppress
	// detection — the same key spiking again on a later day is a new anomaly and
	// fires as normal.
	Acknowledgement *CostAnomalyAcknowledgement `json:"acknowledgement"`
}

// CostAnomalySettings is the `CostAnomalySettings` schema.
type CostAnomalySettings struct {
	// Sigmas: Standard deviations above a key's own trailing mean that count as
	// a spike. Lower is more sensitive. Bounded at 1 — below that roughly a
	// third of ordinary days clear the bar — and at 10, above which nothing
	// short of a 10x jump fires. Defaults to 3.
	Sigmas float64 `json:"sigmas"`
	// MinDeltaCents: Minimum rise over the baseline mean before a spike alerts,
	// in USD cents (converted per series, so it means the same real amount in
	// every currency). Defaults to 1000 ($10).
	MinDeltaCents int64 `json:"minDeltaCents"`
	// NewSourceMinCents: Minimum first-day spend before a new spend source
	// alerts, in USD cents. A key with no prior spend has no statistical bar to
	// clear, so this absolute floor is the only thing keeping a new $0.02/day
	// service quiet. Defaults to 2500 ($25).
	NewSourceMinCents int64 `json:"newSourceMinCents"`
	// SmsAlerts: Which anomalies also text the organization's Twilio recipients.
	// Defaults to `off` — an organization with Twilio configured for budgets
	// does not start receiving anomaly texts until it asks to. `new_source`
	// texts only about spend appearing from nothing, which is what a leaked key
	// looks like on a bill; `all` adds spikes on existing lines. Delivery is
	// batched — one SMS per detection pass summarizing what it alerted on, at
	// most one every six hours per organization — and never places a voice call.
	// Push, Slack and Teams delivery is unaffected by this setting.
	//
	// One of "off", "new_source", "all".
	SmsAlerts string `json:"smsAlerts"`
}

// CostAnomalySettingsView is the `CostAnomalySettingsView` schema.
type CostAnomalySettingsView struct {
	// Sigmas: Standard deviations above a key's own trailing mean that count as
	// a spike. Lower is more sensitive. Bounded at 1 — below that roughly a
	// third of ordinary days clear the bar — and at 10, above which nothing
	// short of a 10x jump fires. Defaults to 3.
	Sigmas float64 `json:"sigmas"`
	// MinDeltaCents: Minimum rise over the baseline mean before a spike alerts,
	// in USD cents (converted per series, so it means the same real amount in
	// every currency). Defaults to 1000 ($10).
	MinDeltaCents int64 `json:"minDeltaCents"`
	// NewSourceMinCents: Minimum first-day spend before a new spend source
	// alerts, in USD cents. A key with no prior spend has no statistical bar to
	// clear, so this absolute floor is the only thing keeping a new $0.02/day
	// service quiet. Defaults to 2500 ($25).
	NewSourceMinCents int64 `json:"newSourceMinCents"`
	// SmsAlerts: Which anomalies also text the organization's Twilio recipients.
	// Defaults to `off` — an organization with Twilio configured for budgets
	// does not start receiving anomaly texts until it asks to. `new_source`
	// texts only about spend appearing from nothing, which is what a leaked key
	// looks like on a bill; `all` adds spikes on existing lines. Delivery is
	// batched — one SMS per detection pass summarizing what it alerted on, at
	// most one every six hours per organization — and never places a voice call.
	// Push, Slack and Teams delivery is unaffected by this setting.
	//
	// One of "off", "new_source", "all".
	SmsAlerts string `json:"smsAlerts"`
	// SmsConfigured: Whether an SMS raised right now could be delivered: paging
	// enabled for the organization, Twilio credentials and a from-number stored,
	// and at least one recipient opted into SMS. Read-only and derived — it is
	// not accepted on PUT.
	SmsConfigured bool `json:"smsConfigured"`
}

// CostBasis: Which number to sum. `cash` is what the provider charged on the day
// it charged it — the default, and what every query returned before this
// existed. `amortized` spreads a commitment's up-front fee across the term it
// buys, so a year of capacity bought on one day is counted on the days it
// covers. Providers that report no amortized amount fall back to their cash
// amount, so an amortized query over a mixed estate never drops their spend.
type CostBasis = string

// The values CostBasis takes.
const (
	CostBasisCash      CostBasis = "cash"
	CostBasisAmortized CostBasis = "amortized"
)

// CostCentre is the `CostCentre` schema.
type CostCentre struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	// ParentID: The centre this one sits under; null is a top-level centre.
	// Nesting is a reporting structure only — allocation still resolves each
	// cost row to exactly one centre.
	ParentID  *string `json:"parentId"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

// CostCentreInput is the `CostCentreInput` schema.
type CostCentreInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// ParentID: Cost centre to nest this one under; null is the top level. On an
	// update, moving a centre is this field changing — omitting it leaves the
	// centre where it is. Rejected with 400 when the parent is unknown, is the
	// centre itself or one of its own descendants, or when the resulting tree
	// would be more than 4 levels deep (measured over the whole subtree being
	// moved).
	ParentID *string `json:"parentId,omitempty"`
}

// CostChangeCadence: Which window is compared to which, in complete UTC days
// (the accruing current day never counts). daily: one complete day vs the same
// weekday one week earlier. weekly: the last 7 complete days vs the 7 before
// them. monthly: month-to-date vs the same number of days at the start of the
// prior month — never MTD vs the full prior month.
type CostChangeCadence = string

// The values CostChangeCadence takes.
const (
	CostChangeCadenceDaily   CostChangeCadence = "daily"
	CostChangeCadenceWeekly  CostChangeCadence = "weekly"
	CostChangeCadenceMonthly CostChangeCadence = "monthly"
)

// CostChangeDirection is the `CostChangeDirection` schema.
type CostChangeDirection = string

// The values CostChangeDirection takes.
const (
	CostChangeDirectionIncrease CostChangeDirection = "increase"
	CostChangeDirectionDecrease CostChangeDirection = "decrease"
	CostChangeDirectionBoth     CostChangeDirection = "both"
)

// CostChargeType is the `CostChargeType` schema.
type CostChargeType = string

// The values CostChargeType takes.
const (
	CostChargeTypeUsage                  CostChargeType = "usage"
	CostChargeTypeCommitmentCoveredUsage CostChargeType = "commitment_covered_usage"
	CostChargeTypeCommitmentFee          CostChargeType = "commitment_fee"
	CostChargeTypeCommitmentDiscount     CostChargeType = "commitment_discount"
	CostChargeTypeCredit                 CostChargeType = "credit"
	CostChargeTypeTax                    CostChargeType = "tax"
	CostChargeTypeRefund                 CostChargeType = "refund"
	CostChargeTypeAdjustment             CostChargeType = "adjustment"
	CostChargeTypeSupport                CostChargeType = "support"
	CostChargeTypeOther                  CostChargeType = "other"
)

// CostDateRange: A relative preset resolves against today every time the report
// runs, so a saved report keeps meaning 'the last 30 days'; an absolute range
// pins it to fixed dates.
type CostDateRange = any

// CostDimension is the `CostDimension` schema.
type CostDimension = string

// The values CostDimension takes.
const (
	CostDimensionProvider   CostDimension = "provider"
	CostDimensionAccount    CostDimension = "account"
	CostDimensionService    CostDimension = "service"
	CostDimensionRegion     CostDimension = "region"
	CostDimensionResource   CostDimension = "resource"
	CostDimensionTag        CostDimension = "tag"
	CostDimensionChargeType CostDimension = "charge_type"
	CostDimensionCommitment CostDimension = "commitment"
)

// CostDimensionValues is the `CostDimensionValues` schema.
type CostDimensionValues struct {
	Values []any `json:"values"`
}

// CostEfficiencySettings is the `CostEfficiencySettings` schema.
type CostEfficiencySettings struct {
	// CommitmentExpiryEnabled: Whether commitments approaching their term end
	// raise alerts. Defaults to true.
	CommitmentExpiryEnabled bool `json:"commitmentExpiryEnabled"`
	// CommitmentExpiryHorizonDays: Days of notice, each firing at most once per
	// commitment per term end. Defaults to [60, 30, 7]. A commitment fires at
	// the *smallest* horizon it has reached, so an account connected 30 days
	// before a term ends gets one alert, not two.
	CommitmentExpiryHorizonDays []int64 `json:"commitmentExpiryHorizonDays"`
	// CommitmentExpiryAlertOnExpired: Whether a commitment that lapsed without
	// any horizon warning having fired raises one alert anyway. Defaults to
	// true, and bounded to terms that ended within the last 90 days — connecting
	// an account with years of dead reservations produces one pass of recent
	// news, not an archive.
	CommitmentExpiryAlertOnExpired bool `json:"commitmentExpiryAlertOnExpired"`
	// CommitmentIdleEnabled: Whether under-used commitments raise alerts.
	// Defaults to true.
	CommitmentIdleEnabled bool `json:"commitmentIdleEnabled"`
	// CommitmentIdleThresholdPercent: Utilization percent the whole window must
	// stay under. Defaults to 70 — roughly where a 1-year no-upfront commitment
	// stops beating on-demand for the usage it covers.
	CommitmentIdleThresholdPercent int64 `json:"commitmentIdleThresholdPercent"`
	// CommitmentIdleWindowDays: Trailing days utilization is aggregated over.
	// Defaults to 30. Aggregated, never sampled per day: a weekday-only workload
	// reads about 71% over a month and does not fire, which is the point.
	CommitmentIdleWindowDays int64 `json:"commitmentIdleWindowDays"`
	// CommitmentIdleMinMeasuredDays: Window days that must carry cost data
	// before anything is judged. Defaults to 14. A commitment whose utilization
	// cannot be measured at all — a unit-denominated GCP CUD, or an account
	// whose plugin reports no commitment attribution — never alerts, regardless
	// of this value.
	CommitmentIdleMinMeasuredDays int64 `json:"commitmentIdleMinMeasuredDays"`
	// CommitmentIdleMinWasteCents: Least wasted money (obligation − delivered)
	// before alerting, in USD cents, restated per currency. Defaults to 5000
	// ($50).
	CommitmentIdleMinWasteCents int64 `json:"commitmentIdleMinWasteCents"`
	// UnitCostRegressionEnabled: Whether rising cost per business-metric unit
	// raises alerts. Defaults to true.
	UnitCostRegressionEnabled bool `json:"unitCostRegressionEnabled"`
	// UnitCostThresholdPercent: Percent the unit cost must rise versus the prior
	// window. Defaults to 20.
	UnitCostThresholdPercent int64 `json:"unitCostThresholdPercent"`
	// UnitCostWindowDays: Length of each of the two compared windows. Defaults
	// to 14 — two whole weekly cycles a side, so a weekday-shaped unit cost
	// compares like with like.
	UnitCostWindowDays int64 `json:"unitCostWindowDays"`
	// UnitCostMinReportedDays: Days inside **each** window that must carry a
	// reported, positive metric value. Defaults to 10. A day with no reported
	// value is a gap and contributes to neither the numerator nor the
	// denominator; a window that fails this bar produces no comparison at all
	// rather than a comparison against a gap.
	UnitCostMinReportedDays int64 `json:"unitCostMinReportedDays"`
	// UnitCostMinSpendCents: Least spend in the current window before alerting,
	// in USD cents, restated per currency. Defaults to 10000 ($100).
	UnitCostMinSpendCents int64 `json:"unitCostMinSpendCents"`
}

// CostEstimate is the `CostEstimate` schema.
//
// The API may send null in its place.
type CostEstimate struct {
	MonthlyAmount float64                `json:"monthlyAmount"`
	Currency      string                 `json:"currency"`
	LineItems     []CostEstimateLineItem `json:"lineItems"`
	Partial       *bool                  `json:"partial,omitempty"`
	Notes         []string               `json:"notes,omitempty"`
}

// CostEstimateLineItem is the `CostEstimateLineItem` schema.
type CostEstimateLineItem struct {
	Label         string   `json:"label"`
	MonthlyAmount float64  `json:"monthlyAmount"`
	Detail        *string  `json:"detail,omitempty"`
	Quantity      *float64 `json:"quantity,omitempty"`
	Unit          *string  `json:"unit,omitempty"`
}

// CostEstimateRequest is the `CostEstimateRequest` schema.
type CostEstimateRequest struct {
	AccountID        string            `json:"accountId"`
	ResourceTypeID   string            `json:"resourceTypeId"`
	Fields           map[string]string `json:"fields,omitempty"`
	ResourceID       *ResourceID       `json:"resourceId,omitempty"`
	PluginID         *string           `json:"pluginId,omitempty"`
	ParentResourceID *ResourceID       `json:"parentResourceId,omitempty"`
}

// CostExport is the `CostExport` schema.
type CostExport struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Format: One of "csv", "ndjson".
	Format string          `json:"format"`
	Query  CostExportQuery `json:"query"`
	// Cadence: One of "daily", "weekly", "monthly".
	Cadence         string                `json:"cadence"`
	Hour            int64                 `json:"hour"`
	Timezone        string                `json:"timezone"`
	RestatementDays int64                 `json:"restatementDays"`
	Enabled         bool                  `json:"enabled"`
	Destination     CostExportDestination `json:"destination"`
	HasCredentials  bool                  `json:"hasCredentials"`
	// CredentialHint: Redacted marker, e.g. `AKIA…7F2Q`. No route ever returns
	// the credential itself.
	CredentialHint *string `json:"credentialHint"`
	LastRunAt      *string `json:"lastRunAt"`
	// LastStatus: One of "pending", "succeeded", "failed".
	LastStatus string `json:"lastStatus"`
	// LastError: Why the last run failed, verbatim from the destination where
	// possible.
	LastError       *string `json:"lastError"`
	LastObjectCount *int64  `json:"lastObjectCount"`
	LastRowCount    *int64  `json:"lastRowCount"`
	NextRunAt       *string `json:"nextRunAt"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// CostExportDestination is the `CostExportDestination` schema.
type CostExportDestination = any

// CostExportFilter is the `CostExportFilter` schema.
type CostExportFilter struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// CostExportInput is the `CostExportInput` schema.
type CostExportInput struct {
	Name string `json:"name"`
	// Format: One of "csv", "ndjson".
	Format string          `json:"format"`
	Query  CostExportQuery `json:"query"`
	// Cadence: How often a run happens and — because a run writes one object per
	// period — what a period is: a calendar day, an ISO week (Monday-start), or
	// a calendar month.
	//
	// One of "daily", "weekly", "monthly".
	Cadence string `json:"cadence"`
	// Hour: Local hour in `timezone` a run fires at.
	Hour int64 `json:"hour"`
	// Timezone: IANA zone, e.g. `Europe/Berlin`. Validated against `Intl`.
	Timezone string `json:"timezone"`
	// RestatementDays: Trailing days of already-written periods each run
	// re-exports. Providers restate spend for days after the fact, so the object
	// written for yesterday is not final; every period overlapping this window
	// is rebuilt in full at its existing key, which overwrites rather than
	// duplicates. 0 disables it and is only correct for an org whose providers
	// never revise.
	RestatementDays int64                 `json:"restatementDays"`
	Enabled         bool                  `json:"enabled"`
	Destination     CostExportDestination `json:"destination"`
	// AccessKeyID: S3 only. Write-only; omit on update to keep the stored
	// credential.
	AccessKeyID *string `json:"accessKeyId,omitempty"`
	// SecretAccessKey: S3 only. Write-only, never returned.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	// URL: HTTPS destinations only. Write-only, never returned — a signed URL
	// carries its own signature, so it is treated as a bearer credential.
	URL *string `json:"url,omitempty"`
}

// CostExportObject is the `CostExportObject` schema.
type CostExportObject struct {
	// PeriodStart: The period's first day, in the export's own timezone.
	PeriodStart string `json:"periodStart"`
	From        string `json:"from"`
	To          string `json:"to"`
	// Key: `{prefix}/cost-export/{exportId}/{cadence}/{periodStart}.{format}`.
	// Deterministic, so re-exporting a restated period overwrites this object
	// instead of adding a second copy.
	Key       string `json:"key"`
	RowCount  int64  `json:"rowCount"`
	ByteCount int64  `json:"byteCount"`
}

// CostExportQuery: The rows a run selects. Reuses the same `CostFilter` and
// dimension vocabulary the dashboards, budgets and cost reports store, so a
// filter means the same thing everywhere.
type CostExportQuery struct {
	Version float64 `json:"version"`
	// Dimensions: Row-identity columns kept in the output. Dropping one
	// aggregates over it — an export grouped to provider + service is orders of
	// magnitude smaller than a per-resource one.
	Dimensions []string `json:"dimensions"`
	// TagKeys: Tag keys emitted as their own `tag_<key>` columns.
	TagKeys     []string           `json:"tagKeys"`
	Filters     []CostExportFilter `json:"filters"`
	ChargeTypes []string           `json:"chargeTypes,omitempty"`
	// CostBasis: One of "cash", "amortized".
	CostBasis *string `json:"costBasis,omitempty"`
}

// CostExportRunResult is the `CostExportRunResult` schema.
type CostExportRunResult struct {
	ExportID string `json:"exportId"`
	// Status: One of "pending", "succeeded", "failed".
	Status   string             `json:"status"`
	Objects  []CostExportObject `json:"objects"`
	RowCount int64              `json:"rowCount"`
	// CollectionWatermark: The newest day every cost-reporting account in the
	// org had data for when the run started. Stamped into every row as
	// `collection_watermark`; rows dated after it are still arriving.
	CollectionWatermark *string `json:"collectionWatermark"`
	Error               *string `json:"error"`
}

// CostFilter is the `CostFilter` schema.
type CostFilter struct {
	Dimension CostDimension `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// CostGraphConfig: The saved graph. Identical to the config an ad-hoc
// `cost_graph` dashboard widget stores inline — a report is that config given a
// name and an id.
type CostGraphConfig struct {
	Version float64 `json:"version"`
	// ChartType: One of "stacked_bar", "multi_bar", "line", "area", "pie".
	ChartType string `json:"chartType"`
	// Binning: One of "daily", "weekly", "monthly", "cumulative".
	Binning   string        `json:"binning"`
	DateRange CostDateRange `json:"dateRange"`
	// GroupBy: One of "none", "provider", "account", "service", "region",
	// "resource", "tag", "charge_type", "commitment".
	GroupBy       string             `json:"groupBy"`
	GroupByTagKey *string            `json:"groupByTagKey,omitempty"`
	Filters       []CostReportFilter `json:"filters,omitempty"`
	// SavedFilterID: A saved cost filter (see /saved-cost-filters) applied by
	// reference and AND-composed with `filters` at query time, server-side.
	// Editing the saved filter changes every graph, report and budget
	// referencing it; a reference that fails to resolve makes the query error
	// rather than silently run unfiltered.
	SavedFilterID         *string `json:"savedFilterId,omitempty"`
	TopN                  *int64  `json:"topN,omitempty"`
	ComparePreviousPeriod *bool   `json:"comparePreviousPeriod,omitempty"`
	ShowForecast          *bool   `json:"showForecast,omitempty"`
	// ScenarioModelID: A scenario model (see /cost-scenarios) overlaid on the
	// forecast — known future cost the trend cannot see, drawn as a second
	// dashed line beside the trend rather than instead of it. Only meaningful
	// alongside `showForecast`.
	ScenarioModelID *string `json:"scenarioModelId,omitempty"`
	// CostBasis: One of "cash", "amortized".
	CostBasis *string `json:"costBasis,omitempty"`
}

// CostPushRequest is the `CostPushRequest` schema.
type CostPushRequest struct {
	// Source: Stable slug naming the system that owns these rows: letters,
	// digits, `.`, `_` and `-`. It groups the rows under an `External` provider
	// and an `external:<source>` account, and re-pushing the same source over
	// the same days restates only its own rows.
	Source string          `json:"source"`
	Rows   []PushedCostRow `json:"rows"`
}

// CostPushResponse is the `CostPushResponse` schema.
type CostPushResponse struct {
	Written int64 `json:"written"`
}

// CostQueryRequest is the `CostQueryRequest` schema.
type CostQueryRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	// Binning: One of "daily", "weekly", "monthly", "cumulative".
	Binning string `json:"binning"`
	// GroupBy: One of "none", "provider", "account", "service", "region",
	// "resource", "tag", "charge_type", "commitment".
	GroupBy       string       `json:"groupBy"`
	GroupByTagKey *string      `json:"groupByTagKey,omitempty"`
	Filters       []CostFilter `json:"filters,omitempty"`
	// Query: The same filter written as text, in the cost query language — an
	// alternative to `filters`, compiled server-side into exactly that
	// structure.
	//
	// Grammar: a conjunction of equality terms joined by `AND`. A term is
	// `dimension = 'value'`, `dimension != 'value'`, `dimension IN ('a','b')` or
	// `dimension NOT IN ('a','b')`; the tag dimension takes its key in brackets,
	// `tag['owner'] = 'platform'`. Keywords are case-insensitive, strings may be
	// single- or double-quoted, and a quote inside a value is escaped by
	// doubling it (`'it''s'`) or with a backslash (`'it\'s'`).
	//
	// `OR` is deliberately not supported: the stored filter is a conjunction, so
	// several values of one dimension go in an `IN` list and unrelated
	// alternatives need separate queries. Anything the structured filter cannot
	// express is a parse error rather than a second execution path.
	//
	// Sending both `query` and a non-empty `filters` is a 400, not a precedence
	// rule. A parse failure is a 400 whose body carries `queryError` with the
	// character `offset`, the `length` of the offending span, and the `expected`
	// alternatives there.
	Query *string `json:"query,omitempty"`
	// SavedFilterID: A saved cost filter (see /saved-cost-filters) applied by
	// reference. Resolved server-side at query time and AND-composed with
	// whichever of `filters`/`query` is present — unlike those two it is a
	// composition, not an alternative. An id that does not resolve to a live
	// filter is a 400; the query is never silently run unfiltered.
	SavedFilterID         *string `json:"savedFilterId,omitempty"`
	TopN                  *int64  `json:"topN,omitempty"`
	ComparePreviousPeriod *bool   `json:"comparePreviousPeriod,omitempty"`
	Forecast              *bool   `json:"forecast,omitempty"`
	// ScenarioModelID: Apply a scenario model (see /cost-scenarios) to the
	// projection: known future cost the trend cannot see. Requires `forecast:
	// true` — sending it without one is a 400, not a no-op, because a caller who
	// asked for assumptions and silently got none back is the failure this
	// feature exists to prevent. The adjusted projection comes back as
	// `scenario`, **alongside** the untouched `forecast`, never instead of it.
	// An id that does not resolve is a 400.
	ScenarioModelID *string    `json:"scenarioModelId,omitempty"`
	CostBasis       *CostBasis `json:"costBasis,omitempty"`
	// ChargeTypes: Restrict to these kinds of charge. Omitted is all of them,
	// which is what makes an unfiltered total net rather than gross — credits,
	// refunds and commitment discounts are included. Rows collected before
	// charge types existed, and rows from providers that cannot distinguish
	// them, are `usage`.
	ChargeTypes []CostChargeType `json:"chargeTypes,omitempty"`
	// Adjusted: Apply the organization's billing rules (see /billing-rules) —
	// markups, discounts, reallocations. Omitted (the default, and what every
	// unattended reader sends) is raw collected spend. Present, the response
	// carries `adjustment` with the collected totals beside the adjusted ones
	// and the rules that moved them; it is set even for an organization with no
	// rules, because the absence of that field is the only signal that a figure
	// is unadjusted.
	Adjusted *bool `json:"adjusted,omitempty"`
}

// CostQueryResponse is the `CostQueryResponse` schema.
type CostQueryResponse struct {
	Series     []CostQuerySeries `json:"series"`
	Comparison []CostQuerySeries `json:"comparison,omitempty"`
	// Forecast: The **unadjusted trend** projection. Stays the trend even when a
	// scenario is applied, so a reader can always see what the fit said before
	// anybody's assumptions touched it.
	Forecast   []CostSeriesPoint   `json:"forecast,omitempty"`
	Scenario   *CostScenarioResult `json:"scenario,omitempty"`
	Currencies []string            `json:"currencies"`
	// Totals: Period total per currency, and always exactly the sum of `series`.
	// Fixed-amount billing-rule charges are deliberately **not** folded in here
	// — they have no series behind them and are reported in
	// `adjustment.fixedTotals` instead.
	Totals         map[string]float64     `json:"totals"`
	PreviousTotals map[string]float64     `json:"previousTotals,omitempty"`
	Adjustment     *CostAdjustmentSummary `json:"adjustment,omitempty"`
}

// CostQuerySeries is the `CostQuerySeries` schema.
type CostQuerySeries struct {
	Key      string            `json:"key"`
	Label    string            `json:"label"`
	Currency string            `json:"currency"`
	Points   []CostSeriesPoint `json:"points"`
}

// CostReport is the `CostReport` schema.
type CostReport struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Config      CostGraphConfig `json:"config"`
	// FolderID: Folder the report is filed under (see /cost-report-folders);
	// null is the top level of the Reports list. Moving a report is this same
	// PUT with a different folderId; an id from another org is a 400. Deleting a
	// folder never deletes its reports — they fall back to the top level.
	FolderID        *string `json:"folderId"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
	// Placements: The dashboards carrying a `cost_report` card for this report.
	// Empty is normal — a report exists, and can be run, whether or not any
	// dashboard shows it. Deleting the report removes these cards; removing a
	// card leaves the report alone.
	Placements []CostReportPlacement `json:"placements"`
}

// CostReportFilter is the `CostReportFilter` schema.
type CostReportFilter struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// CostReportFolder is the `CostReportFolder` schema.
type CostReportFolder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// ParentFolderID: Parent folder for nesting; null is a top-level folder.
	// Nesting is capped at 3 levels, and moving a folder inside itself or one of
	// its own subfolders is rejected — both are 400s.
	ParentFolderID *string `json:"parentFolderId"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
}

// CostReportFolderInput is the `CostReportFolderInput` schema.
type CostReportFolderInput struct {
	Name string `json:"name"`
	// ParentFolderID: Parent folder for nesting; null is a top-level folder.
	// Nesting is capped at 3 levels, and moving a folder inside itself or one of
	// its own subfolders is rejected — both are 400s.
	ParentFolderID *string `json:"parentFolderId,omitempty"`
}

// CostReportInput is the `CostReportInput` schema.
type CostReportInput struct {
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Config      CostGraphConfig `json:"config"`
	// FolderID: Folder the report is filed under (see /cost-report-folders);
	// null is the top level of the Reports list. Moving a report is this same
	// PUT with a different folderId; an id from another org is a 400. Deleting a
	// folder never deletes its reports — they fall back to the top level.
	FolderID *string `json:"folderId,omitempty"`
}

// CostReportPlacement is the `CostReportPlacement` schema.
type CostReportPlacement struct {
	WidgetID      string `json:"widgetId"`
	DashboardID   string `json:"dashboardId"`
	DashboardName string `json:"dashboardName"`
}

// CostReportRunResult is the `CostReportRunResult` schema.
type CostReportRunResult struct {
	ReportID string            `json:"reportId"`
	Name     string            `json:"name"`
	From     string            `json:"from"`
	To       string            `json:"to"`
	Result   CostQueryResponse `json:"result"`
}

// CostScenarioAdjustment is the `CostScenarioAdjustment` schema.
type CostScenarioAdjustment struct {
	// ID: Stable within the model; also the key of its per-adjustment total.
	ID string `json:"id"`
	// Label: What this adjustment is. Named on the chart whenever the scenario
	// moves a number.
	Label string `json:"label"`
	// Kind: `one_off` is a single amount on a single day; `recurring` is an
	// amount every period from a date; `rate_change` is ±X% of the trend from a
	// date. The split between an amount and a percentage of the trend is what
	// fixes the composition order — see the `scenario` field on the cost query
	// response.
	//
	// One of "one_off", "recurring", "rate_change".
	Kind      string `json:"kind"`
	StartDate string `json:"startDate"`
	// EndDate: Inclusive last day, or null for indefinitely. Refused for
	// `one_off`, which is one day.
	EndDate *string `json:"endDate"`
	// AmountCents: Minor units of the model's currency, for the amount kinds;
	// null for `rate_change`. May be negative — turning off an old cluster is as
	// real a known future cost as buying a new one.
	AmountCents *int64 `json:"amountCents"`
	// Currency: Always the model's own currency; a model that held two would sum
	// two kinds of money.
	Currency *string `json:"currency"`
	// Period: How often a `recurring` amount charges. A monthly amount is spread
	// evenly across each calendar month it covers rather than landing as a spike
	// on the 1st, so a month the scenario only partly covers costs
	// proportionally less.
	//
	// One of "daily", "monthly".
	Period *string `json:"period"`
	// Percent: Percent change to the trend, for `rate_change`. -20 is a fifth
	// cheaper.
	Percent *float64 `json:"percent"`
	// Scope: Which spend this adjustment describes; empty is the whole
	// organization. For a rate change the scope is what the percentage is *of*.
	// For an amount it decides whether the adjustment applies to a given chart
	// at all — a GCP commitment does not belong on a chart filtered to AWS, and
	// one that is excluded is named in `scenario.outOfScope`.
	Scope []CostScenarioScopeTerm `json:"scope"`
}

// CostScenarioModel is the `CostScenarioModel` schema.
type CostScenarioModel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	// Currency: The one currency every amount in this model is denominated in.
	Currency        string                   `json:"currency"`
	Adjustments     []CostScenarioAdjustment `json:"adjustments"`
	CreatedByUserID *string                  `json:"createdByUserId"`
	CreatedAt       string                   `json:"createdAt"`
	UpdatedAt       string                   `json:"updatedAt"`
}

// CostScenarioModelInput is the `CostScenarioModelInput` schema.
type CostScenarioModelInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Currency: Three-letter code. Every amount in the model must be in it — a
	// model that mixed two would produce a projection that is the sum of two
	// kinds of money, so this is refused rather than converted behind the
	// caller's back.
	Currency    string                   `json:"currency"`
	Adjustments []CostScenarioAdjustment `json:"adjustments"`
}

// CostScenarioReferent is the `CostScenarioReferent` schema.
type CostScenarioReferent struct {
	// Kind: One of "budget", "cost_report", "cost_graph_widget".
	Kind string `json:"kind"`
	// ID: Budget id, report id, or dashboard-widget id.
	ID   string `json:"id"`
	Name string `json:"name"`
	// DashboardID: Set for `cost_graph_widget` referents.
	DashboardID   *string `json:"dashboardId,omitempty"`
	DashboardName *string `json:"dashboardName,omitempty"`
}

// CostScenarioResult is the `CostScenarioResult` schema.
type CostScenarioResult struct {
	ModelID   string `json:"modelId"`
	ModelName string `json:"modelName"`
	Currency  string `json:"currency"`
	// Points: The adjusted projection — exactly the same days as `forecast`,
	// never one more or fewer. A scenario modifies the projected region; it does
	// not extend it, and it can never touch a day that already has recorded
	// spend behind it.
	Points []CostSeriesPoint `json:"points"`
	// Contributions: Signed total each adjustment added across the horizon, in
	// model order.
	Contributions []CostScenarioResultContributions `json:"contributions"`
	// TotalDelta: Signed difference from the baseline across the horizon.
	TotalDelta float64 `json:"totalDelta"`
	// ConvertedFrom: Set when the model's amounts were converted at the org's
	// stated rates.
	ConvertedFrom *string `json:"convertedFrom,omitempty"`
	// OutOfScope: Adjustments this chart's own filters exclude, by label — a GCP
	// commitment on an AWS-filtered chart is correctly left out, and saying so
	// is what makes the number trustworthy rather than quietly assumed broken.
	OutOfScope []string `json:"outOfScope"`
}

// CostScenarioScopeTerm is the `CostScenarioScopeTerm` schema.
type CostScenarioScopeTerm struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// CostSeriesPoint is the `CostSeriesPoint` schema.
type CostSeriesPoint struct {
	Bucket string  `json:"bucket"`
	Amount float64 `json:"amount"`
}

// CreateAccountRequest is the `CreateAccountRequest` schema.
type CreateAccountRequest struct {
	PluginID    string            `json:"pluginId"`
	DisplayName string            `json:"displayName"`
	Credentials map[string]string `json:"credentials"`
	// BastionID: Optional bastion id to route this account's cloud API traffic
	// through.
	BastionID *string `json:"bastionId,omitempty"`
}

// CreateAccountResponse is the `CreateAccountResponse` schema.
type CreateAccountResponse struct {
	ID        string                          `json:"id"`
	SyncError *CreateAccountResponseSyncError `json:"syncError,omitempty"`
}

// CreateAgentSession is the `CreateAgentSession` schema.
type CreateAgentSession struct {
	Repo          *string        `json:"repo,omitempty"`
	ProjectName   *string        `json:"projectName,omitempty"`
	WorkspaceName *string        `json:"workspaceName,omitempty"`
	Settings      *AgentSettings `json:"settings"`
}

// CreateAPIKeyRequest is the `CreateApiKeyRequest` schema.
//
// Spec schema: `CreateApiKeyRequest`.
type CreateAPIKeyRequest struct {
	Name      string       `json:"name"`
	Scopes    []Permission `json:"scopes"`
	ExpiresAt *string      `json:"expiresAt,omitempty"`
}

// CreateBastionRequest is the `CreateBastionRequest` schema.
type CreateBastionRequest struct {
	Name string `json:"name"`
}

// CreateBastionResponse is the `CreateBastionResponse` schema.
type CreateBastionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TokenPrefix string `json:"tokenPrefix"`
	// Token: Enrollment token in the form `iwb_<random>`. Pass to the agent
	// container as `BASTION_TOKEN`. Returned once — not recoverable later.
	Token string `json:"token"`
}

// CreateConfigRequest is the `CreateConfigRequest` schema.
type CreateConfigRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceTypeID   string      `json:"resourceTypeId"`
	PluginID         *string     `json:"pluginId,omitempty"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// CreateJiraIssueInput is the `CreateJiraIssueInput` schema.
type CreateJiraIssueInput struct {
	SourceKind JiraSourceKind `json:"sourceKind"`
	// SourceID: The finding's own id, as the detector reports it.
	SourceID    string `json:"sourceId"`
	ProjectKey  string `json:"projectKey"`
	IssueTypeID string `json:"issueTypeId"`
	Summary     string `json:"summary"`
	// Description: Plain text. Converted server-side to Atlassian Document
	// Format, which is what the Jira REST v3 description field requires; blank
	// lines become paragraphs.
	Description *string `json:"description,omitempty"`
	// Labels: Whitespace inside a label is replaced with '-', since Jira rejects
	// it.
	Labels []string `json:"labels,omitempty"`
}

// CreateJiraIssueResult is the `CreateJiraIssueResult` schema.
type CreateJiraIssueResult struct {
	Issue CreateJiraIssueResultIssue `json:"issue"`
	Link  JiraIssueLink              `json:"link"`
}

// CreateLinearIssueInput is the `CreateLinearIssueInput` schema.
type CreateLinearIssueInput struct {
	SourceKind LinearSourceKind `json:"sourceKind"`
	// SourceID: The finding's own id, as the detector reports it.
	SourceID string `json:"sourceId"`
	// TeamID: Team to file into. Every Linear issue belongs to exactly one team.
	TeamID string `json:"teamId"`
	Title  string `json:"title"`
	// Description: Markdown, passed to Linear as-is — unlike Jira, where the
	// server converts plain text to Atlassian Document Format.
	Description *string `json:"description,omitempty"`
	// LabelIDs: Ids of existing labels in the workspace. Linear cannot create
	// labels here.
	LabelIDs []string `json:"labelIds,omitempty"`
	// ProjectID: Optional project to attach the issue to.
	ProjectID *string `json:"projectId,omitempty"`
}

// CreateLinearIssueResult is the `CreateLinearIssueResult` schema.
type CreateLinearIssueResult struct {
	Issue CreateLinearIssueResultIssue `json:"issue"`
	Link  LinearIssueLink              `json:"link"`
}

// CreateOrgRequest is the `CreateOrgRequest` schema.
type CreateOrgRequest struct {
	DisplayName string `json:"displayName"`
}

// CreatePricingRequest is the `CreatePricingRequest` schema.
type CreatePricingRequest struct {
	AccountID        string                      `json:"accountId"`
	ResourceTypeID   string                      `json:"resourceTypeId"`
	RegionID         *string                     `json:"regionId,omitempty"`
	Sizes            []CreatePricingRequestSizes `json:"sizes"`
	PluginID         *string                     `json:"pluginId,omitempty"`
	ParentResourceID *ResourceID                 `json:"parentResourceId,omitempty"`
}

// CreateResourceRequest is the `CreateResourceRequest` schema.
type CreateResourceRequest struct {
	AccountID        string            `json:"accountId"`
	PluginID         string            `json:"pluginId"`
	ResourceTypeID   string            `json:"resourceTypeId"`
	Fields           map[string]string `json:"fields"`
	ParentResourceID *ResourceID       `json:"parentResourceId,omitempty"`
}

// CreateResourceResponse is the `CreateResourceResponse` schema.
type CreateResourceResponse struct {
	ID          ResourceID `json:"id"`
	DisplayName string     `json:"displayName"`
	Warnings    []string   `json:"warnings,omitempty"`
}

// CreateSharedConsole is the `CreateSharedConsole` schema.
type CreateSharedConsole struct {
	// LiveConsoleID: The pty to share, as the terminal's WebSocket reported it
	// in its `ssh:connected` frame. Everything else about the session — host,
	// account, recording — is read from the proxy's own registration rather than
	// from this body.
	LiveConsoleID string `json:"liveConsoleId"`
	RoutingKey    string `json:"routingKey"`
	// AllowHandover: Defaults to true.
	AllowHandover *bool `json:"allowHandover,omitempty"`
	// InviteTTLMinutes: Defaults to 15.
	InviteTTLMinutes *int64 `json:"inviteTtlMinutes,omitempty"`
}

// CreateWidgetRequest is the `CreateWidgetRequest` schema.
type CreateWidgetRequest struct {
	DashboardID string              `json:"dashboardId"`
	Kind        DashboardWidgetKind `json:"kind"`
	Title       *string             `json:"title,omitempty"`
	Config      JSONObject          `json:"config"`
}

// CreatedAPIKey is the `CreatedApiKey` schema.
//
// Spec schema: `CreatedApiKey`.
type CreatedAPIKey struct {
	ID string `json:"id"`
	// Key: Plaintext key. Returned once. Format: `iwk_<base64url>`.
	Key string `json:"key"`
}

// CredentialExport is the `CredentialExport` schema.
type CredentialExport struct {
	Content  string                   `json:"content"`
	Filename string                   `json:"filename"`
	MimeType string                   `json:"mimeType"`
	Fields   []CredentialExportFields `json:"fields,omitempty"`
	Warning  *string                  `json:"warning,omitempty"`
}

// CredentialField is the `CredentialField` schema.
type CredentialField struct {
	Key          string                   `json:"key"`
	Label        string                   `json:"label"`
	Description  *string                  `json:"description,omitempty"`
	Placeholder  *string                  `json:"placeholder,omitempty"`
	Sensitive    *bool                    `json:"sensitive,omitempty"`
	Multiline    *bool                    `json:"multiline,omitempty"`
	DefaultValue *string                  `json:"defaultValue,omitempty"`
	Regions      []CredentialFieldRegion  `json:"regions,omitempty"`
	HelpLink     *CredentialFieldHelpLink `json:"helpLink,omitempty"`
}

// CredentialFieldRegion is the `CredentialFieldRegion` schema.
type CredentialFieldRegion struct {
	ID       string  `json:"id"`
	Label    string  `json:"label"`
	Location *string `json:"location,omitempty"`
	Flag     *string `json:"flag,omitempty"`
}

// CredentialFormat is the `CredentialFormat` schema.
type CredentialFormat struct {
	// ID: Passed back as `formatId` on export.
	ID          string  `json:"id"`
	Label       string  `json:"label"`
	Description *string `json:"description,omitempty"`
	// MediaType: How the credential body should be presented. `binary-base64`
	// means `content` is base64.
	//
	// One of "json", "text", "ini", "binary-base64".
	MediaType string `json:"mediaType"`
	// FilenameTemplate: Suggested filename; `{resource}` is replaced with the
	// resource's external id.
	FilenameTemplate *string `json:"filenameTemplate,omitempty"`
}

// CreditBurndown is the `CreditBurndown` schema.
type CreditBurndown struct {
	Pots     []CreditPot         `json:"pots"`
	Failures []CreditPollFailure `json:"failures"`
	// PendingAccountIDs: Credit-capable accounts never yet collected — named
	// rather than omitted.
	PendingAccountIDs []string `json:"pendingAccountIds"`
	BurnWindowDays    int64    `json:"burnWindowDays"`
}

// CreditPollFailure is the `CreditPollFailure` schema.
type CreditPollFailure struct {
	AccountID   string   `json:"accountId"`
	AccountName string   `json:"accountName"`
	PluginID    PluginID `json:"pluginId"`
	Error       string   `json:"error"`
	HelpLabel   *string  `json:"helpLabel"`
	// HelpURL: Set when the plugin reported a permission gap rather than an
	// outage.
	HelpURL      *string `json:"helpUrl"`
	FailureCount int64   `json:"failureCount"`
}

// CreditPot is the `CreditPot` schema.
type CreditPot struct {
	AccountID   string   `json:"accountId"`
	AccountName string   `json:"accountName"`
	PluginID    PluginID `json:"pluginId"`
	// CapabilityLabel: The provider's own word for this pot — "Credits",
	// "Balance".
	CapabilityLabel string  `json:"capabilityLabel"`
	TopUpURL        *string `json:"topUpUrl"`
	// PotKey: Stable identity for this pot within the account — a currency code,
	// a project id — so successive readings line up into a series.
	PotKey    string  `json:"potKey"`
	Label     string  `json:"label"`
	Remaining float64 `json:"remaining"`
	Currency  string  `json:"currency"`
	// Granted: What was granted, when the provider reports it.
	Granted *float64 `json:"granted"`
	// CreditExpiresAt: Hard expiry on the credit itself, independent of burn.
	CreditExpiresAt *string `json:"creditExpiresAt"`
	ObservedAt      string  `json:"observedAt"`
	// BurnPerDay: Spend per day over the observed span. **Null means there is
	// not enough history to say** — never 0, which would read as 'nothing is
	// being spent'.
	BurnPerDay   *float64 `json:"burnPerDay"`
	BurnSpanDays float64  `json:"burnSpanDays"`
	Observations int64    `json:"observations"`
	// TopUps: Increases seen between consecutive readings. A top-up is recorded,
	// never netted off the burn — subtracting the endpoints of a window
	// containing one reports a negative burn and an infinite runway.
	TopUps      int64    `json:"topUps"`
	RunwayDays  *float64 `json:"runwayDays"`
	ExhaustedAt *string  `json:"exhaustedAt"`
	// NeverEmpties: Nothing has been spent over the observed span.
	NeverEmpties bool `json:"neverEmpties"`
	// LimitedByExpiry: The credit's own expiry, not the burn rate, is the
	// binding deadline.
	LimitedByExpiry bool `json:"limitedByExpiry"`
	// Urgency: One of "critical", "warning", "ok", "unknown".
	Urgency string `json:"urgency"`
}

// CurrencyConfig is the `CurrencyConfig` schema.
type CurrencyConfig struct {
	// DisplayCurrency: ISO 4217 code, upper-case.
	DisplayCurrency *string        `json:"displayCurrency"`
	Rates           []ExchangeRate `json:"rates"`
}

// CurrencySettings is the `CurrencySettings` schema.
type CurrencySettings struct {
	// DisplayCurrency: The currency converted amounts are expressed in, or
	// `null` for no conversion at all. `null` is the default and the state of
	// every organization that has not opted in: cost data is stored per currency
	// and never merged unless you ask.
	DisplayCurrency *string `json:"displayCurrency"`
}

// CustomGraphCheckRequest is the `CustomGraphCheckRequest` schema.
type CustomGraphCheckRequest struct {
	Source string `json:"source"`
}

// CustomGraphCheckResult is the `CustomGraphCheckResult` schema.
type CustomGraphCheckResult struct {
	Diagnostics []CustomGraphCheckResultDiagnostics `json:"diagnostics"`
	HasErrors   bool                                `json:"hasErrors"`
	Degraded    bool                                `json:"degraded"`
}

// CustomGraphFull is the `CustomGraphFull` schema.
type CustomGraphFull struct {
	ID                 string  `json:"id"`
	OrganizationID     string  `json:"organizationId"`
	Name               string  `json:"name"`
	Description        *string `json:"description"`
	Source             string  `json:"source"`
	CreatedByUserID    *string `json:"createdByUserId"`
	SourceAuthorUserID *string `json:"sourceAuthorUserId"`
	DeletedAt          *string `json:"deletedAt"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
}

// CustomGraphInput is the `CustomGraphInput` schema.
type CustomGraphInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Source      *string `json:"source,omitempty"`
}

// CustomGraphRenderRequest is the `CustomGraphRenderRequest` schema.
type CustomGraphRenderRequest struct {
	Controls map[string]any `json:"controls,omitempty"`
	Button   *string        `json:"button,omitempty"`
	// Trigger: One of "manual", "refresh", "interaction".
	Trigger *string `json:"trigger,omitempty"`
}

// CustomGraphRenderResult is the `CustomGraphRenderResult` schema.
type CustomGraphRenderResult struct {
	OK         bool                          `json:"ok"`
	Spec       JSONObject                    `json:"spec"`
	Error      *string                       `json:"error"`
	Logs       []CustomGraphRenderResultLogs `json:"logs"`
	RenderedAt string                        `json:"renderedAt"`
	DurationMs int64                         `json:"durationMs"`
}

// CustomGraphSummary is the `CustomGraphSummary` schema.
type CustomGraphSummary struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// CustomGraphUpdate is the `CustomGraphUpdate` schema.
type CustomGraphUpdate struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Source      *string `json:"source,omitempty"`
}

// Dashboard is the `Dashboard` schema.
type Dashboard struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
}

// DashboardFull is the `DashboardFull` schema.
type DashboardFull struct {
	ID             string  `json:"id"`
	OrganizationID string  `json:"organizationId"`
	Name           string  `json:"name"`
	IsDefault      bool    `json:"isDefault"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
	DeletedAt      *string `json:"deletedAt"`
	SyncVersion    int64   `json:"syncVersion"`
}

// DashboardPin is the `DashboardPin` schema.
type DashboardPin struct {
	PinID      string     `json:"pinId"`
	ResourceID ResourceID `json:"resourceId"`
	GridX      int64      `json:"gridX"`
	GridY      int64      `json:"gridY"`
	GridW      int64      `json:"gridW"`
	GridH      int64      `json:"gridH"`
}

// DashboardWidget is the `DashboardWidget` schema.
type DashboardWidget struct {
	ID          string              `json:"id"`
	DashboardID string              `json:"dashboardId"`
	Kind        DashboardWidgetKind `json:"kind"`
	Title       string              `json:"title"`
	Config      JSONObject          `json:"config"`
	GridX       int64               `json:"gridX"`
	GridY       int64               `json:"gridY"`
	GridW       int64               `json:"gridW"`
	GridH       int64               `json:"gridH"`
}

// DashboardWidgetFull is the `DashboardWidgetFull` schema.
type DashboardWidgetFull struct {
	ID             string              `json:"id"`
	OrganizationID string              `json:"organizationId"`
	DashboardID    string              `json:"dashboardId"`
	Kind           DashboardWidgetKind `json:"kind"`
	Title          string              `json:"title"`
	Config         JSONObject          `json:"config"`
	GridX          int64               `json:"gridX"`
	GridY          int64               `json:"gridY"`
	GridW          int64               `json:"gridW"`
	GridH          int64               `json:"gridH"`
	SyncVersion    int64               `json:"syncVersion"`
	DeletedAt      *string             `json:"deletedAt"`
	CreatedAt      string              `json:"createdAt"`
	UpdatedAt      string              `json:"updatedAt"`
}

// DashboardWidgetKind: `cost_graph` stores its whole config inline — a one-off
// card. `cost_report` points at a saved cost report by id, so editing the report
// updates every dashboard showing it.
type DashboardWidgetKind = string

// The values DashboardWidgetKind takes.
const (
	DashboardWidgetKindCostGraph   DashboardWidgetKind = "cost_graph"
	DashboardWidgetKindCostReport  DashboardWidgetKind = "cost_report"
	DashboardWidgetKindBudget      DashboardWidgetKind = "budget"
	DashboardWidgetKindCustomGraph DashboardWidgetKind = "custom_graph"
)

// DashboardWithPins is the `DashboardWithPins` schema.
type DashboardWithPins struct {
	Dashboard    DashboardFull          `json:"dashboard"`
	Pins         []DashboardPin         `json:"pins"`
	WorkflowPins []DashboardWorkflowPin `json:"workflowPins"`
	Widgets      []DashboardWidget      `json:"widgets"`
}

// DashboardWorkflowPin is the `DashboardWorkflowPin` schema.
type DashboardWorkflowPin struct {
	PinID      string                        `json:"pinId"`
	WorkflowID string                        `json:"workflowId"`
	GridX      int64                         `json:"gridX"`
	Name       string                        `json:"name"`
	LastRunAt  *string                       `json:"lastRunAt"`
	LastStatus *string                       `json:"lastStatus"`
	Metrics    []DashboardWorkflowPinMetrics `json:"metrics"`
}

// DependencyGraphEdge is the `DependencyGraphEdge` schema.
type DependencyGraphEdge struct {
	ConsumerResourceID ResourceID `json:"consumerResourceId"`
	// ConsumerFieldKey: The consumer field the reference fills. "parent" for
	// containment edges, where the link is the resource hierarchy itself rather
	// than a field.
	ConsumerFieldKey   string     `json:"consumerFieldKey"`
	ProviderResourceID ResourceID `json:"providerResourceId"`
	// ProviderOutputKey: The provider output or identity the reference reads —
	// an output key for output references, the matched identity ("externalId",
	// "name", "endpoint"…) for inferred edges.
	ProviderOutputKey string `json:"providerOutputKey"`
	// Kind: Where the edge came from: `output-ref` is wired by hand, `declared`
	// from the plugin's own `dependsOn` rule for the resource type,
	// `containment` from the synced parent/child link, `field-match` from a
	// field value that exactly matches another resource's identity. Absent means
	// `output-ref`.
	//
	// One of "output-ref", "declared", "containment", "field-match".
	Kind *string `json:"kind,omitempty"`
	// Label: How the plugin words the relationship ("in VPC", "guarded by"),
	// when it declared one.
	Label *string `json:"label,omitempty"`
}

// DependencyGraphNode is the `DependencyGraphNode` schema.
type DependencyGraphNode struct {
	ID                ResourceID `json:"id"`
	DisplayName       string     `json:"displayName"`
	PluginID          string     `json:"pluginId"`
	PluginDisplayName string     `json:"pluginDisplayName"`
	// PluginLogoSvg: Inline SVG markup; may be empty.
	PluginLogoSvg     string `json:"pluginLogoSvg"`
	ResourceTypeID    string `json:"resourceTypeId"`
	ResourceTypeLabel string `json:"resourceTypeLabel"`
	AccountID         string `json:"accountId"`
	AccountName       string `json:"accountName"`
}

// DependencyGraphResponse is the `DependencyGraphResponse` schema.
type DependencyGraphResponse struct {
	// Nodes: Org resources that participate in at least one edge.
	Nodes []DependencyGraphNode `json:"nodes"`
	// Edges: Directed depends-on edges (consumer → provider), deduped per
	// consumer field and provider.
	Edges []DependencyGraphEdge `json:"edges"`
	// Truncated: True when inference hit its edge cap and the returned graph is
	// a partial view of the org.
	Truncated bool `json:"truncated"`
}

// DeployCreatedResource is the `DeployCreatedResource` schema.
type DeployCreatedResource struct {
	PluginID       string                        `json:"pluginId"`
	AccountID      string                        `json:"accountId"`
	ResourceTypeID string                        `json:"resourceTypeId"`
	ResourceID     string                        `json:"resourceId"`
	ExternalID     *string                       `json:"externalId,omitempty"`
	DisplayName    string                        `json:"displayName"`
	Sidecar        *DeployCreatedResourceSidecar `json:"sidecar,omitempty"`
}

// DeployEnvs is the `DeployEnvs` schema.
type DeployEnvs struct {
	Envs   []string `json:"envs"`
	Sha    string   `json:"sha"`
	Repo   string   `json:"repo"`
	Branch string   `json:"branch"`
}

// DeployEnvsInput is the `DeployEnvsInput` schema.
type DeployEnvsInput struct {
	Repo   string  `json:"repo"`
	Branch *string `json:"branch,omitempty"`
}

// DeployPlanInput is the `DeployPlanInput` schema.
type DeployPlanInput struct {
	Repo    string            `json:"repo"`
	Branch  *string           `json:"branch,omitempty"`
	Env     *string           `json:"env,omitempty"`
	Answers map[string]string `json:"answers,omitempty"`
}

// DeployPlanResult is the `DeployPlanResult` schema.
type DeployPlanResult struct {
	RunID  string                 `json:"runId"`
	Result DeployPlanResultResult `json:"result"`
}

// DeployPlannedChange is the `DeployPlannedChange` schema.
type DeployPlannedChange struct {
	// Action: One of "create", "update", "delete".
	Action         string                      `json:"action"`
	AccountID      string                      `json:"accountId"`
	ResourceTypeID string                      `json:"resourceTypeId"`
	ResourceID     *string                     `json:"resourceId,omitempty"`
	DisplayName    string                      `json:"displayName"`
	Fields         map[string]string           `json:"fields,omitempty"`
	Sidecar        *DeployPlannedChangeSidecar `json:"sidecar,omitempty"`
}

// DeployRepo is the `DeployRepo` schema.
type DeployRepo struct {
	FullName      string `json:"fullName"`
	DefaultBranch string `json:"defaultBranch"`
}

// DeployRollbackInput is the `DeployRollbackInput` schema.
type DeployRollbackInput struct {
	DeleteCreated *bool `json:"deleteCreated,omitempty"`
}

// DeployRunLog is the `DeployRunLog` schema.
type DeployRunLog struct {
	At int64 `json:"at"`
	// Level: One of "debug", "info", "warn", "error".
	Level   string `json:"level"`
	Message string `json:"message"`
}

// DeployStage is the `DeployStage` schema.
type DeployStage = string

// The values DeployStage takes.
const (
	DeployStagePlan       DeployStage = "plan"
	DeployStageDockerfile DeployStage = "dockerfile"
	DeployStageBuild      DeployStage = "build"
	DeployStageDeploy     DeployStage = "deploy"
	DeployStageDestroy    DeployStage = "destroy"
)

// DeployStatus is the `DeployStatus` schema.
type DeployStatus = string

// The values DeployStatus takes.
const (
	DeployStatusPending  DeployStatus = "pending"
	DeployStatusRunning  DeployStatus = "running"
	DeployStatusSuccess  DeployStatus = "success"
	DeployStatusFailure  DeployStatus = "failure"
	DeployStatusCanceled DeployStatus = "canceled"
)

// DeployTrigger is the `DeployTrigger` schema.
type DeployTrigger struct {
	ID        string  `json:"id"`
	Repo      string  `json:"repo"`
	Branch    string  `json:"branch"`
	Env       string  `json:"env"`
	Enabled   bool    `json:"enabled"`
	LastSha   *string `json:"lastSha"`
	LastRunAt *string `json:"lastRunAt"`
}

// DeployTriggerInput is the `DeployTriggerInput` schema.
type DeployTriggerInput struct {
	Repo    string            `json:"repo"`
	Branch  string            `json:"branch"`
	Env     string            `json:"env"`
	Answers map[string]string `json:"answers,omitempty"`
}

// DeploymentCostImpact is the `DeploymentCostImpact` schema.
type DeploymentCostImpact struct {
	RunID      string          `json:"runId"`
	CostBasis  ChangeCostBasis `json:"costBasis"`
	WindowDays int64           `json:"windowDays"`
	// EventDay: The run's start day, UTC — what both windows hang off.
	EventDay string `json:"eventDay"`
	// Resources: One row per resource the run provisioned through
	// `infra.accounts.*.create(...)`. That is the only set attributable to a run
	// with certainty: a deploy that merely re-shipped an image links to nothing
	// and honestly reports an empty breakdown.
	Resources []DeploymentCostImpactResource `json:"resources"`
	// Total: Summed `deltaPerDay` per currency across the **measured** rows
	// only, so the breakdown always adds up to it. An unmeasurable resource
	// contributes nothing rather than zero.
	Total []DeploymentCostImpactTotal `json:"total"`
	// UnknownResources: Rows excluded from `total` because their impact could
	// not be measured.
	UnknownResources int64                      `json:"unknownResources"`
	Confidence       ChangeCostImpactConfidence `json:"confidence"`
}

// DeploymentCostImpactResource is the `DeploymentCostImpactResource` schema.
type DeploymentCostImpactResource struct {
	ResourceID     ResourceID       `json:"resourceId"`
	DisplayName    string           `json:"displayName"`
	PluginID       string           `json:"pluginId"`
	ResourceTypeID string           `json:"resourceTypeId"`
	Impact         ChangeCostImpact `json:"impact"`
}

// DeploymentRun is the `DeploymentRun` schema.
type DeploymentRun struct {
	ID     string       `json:"id"`
	Env    string       `json:"env"`
	Repo   *string      `json:"repo"`
	Branch *string      `json:"branch"`
	GitSha *string      `json:"gitSha"`
	Image  *string      `json:"image"`
	Status DeployStatus `json:"status"`
	// Origin: One of "web", "cli", "trigger".
	Origin       string       `json:"origin"`
	Stage        *DeployStage `json:"stage"`
	DurationMs   *int64       `json:"durationMs"`
	BuildSeconds *int64       `json:"buildSeconds"`
	// BuildRunner: One of "cloud-build", "ssh".
	BuildRunner *string `json:"buildRunner"`
	StartedAt   string  `json:"startedAt"`
}

// DeploymentRunInput is the `DeploymentRunInput` schema.
type DeploymentRunInput struct {
	Env              string                   `json:"env"`
	Status           DeployStatus             `json:"status"`
	Repo             *string                  `json:"repo,omitempty"`
	Branch           *string                  `json:"branch,omitempty"`
	GitSha           *string                  `json:"gitSha,omitempty"`
	Image            *string                  `json:"image,omitempty"`
	Stage            *DeployStage             `json:"stage,omitempty"`
	Notes            []string                 `json:"notes,omitempty"`
	Output           any                      `json:"output,omitempty"`
	Plan             any                      `json:"plan,omitempty"`
	CreatedResources []DeployCreatedResource  `json:"createdResources,omitempty"`
	DurationMs       *int64                   `json:"durationMs,omitempty"`
	Error            *DeploymentRunInputError `json:"error,omitempty"`
}

// DescribeRequest is the `DescribeRequest` schema.
type DescribeRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceID       ResourceID  `json:"resourceId"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// DescribeResponse is the `DescribeResponse` schema.
type DescribeResponse struct {
	Text string `json:"text"`
}

// DigestEmailRecipient is the `DigestEmailRecipient` schema.
type DigestEmailRecipient struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// DigestEmailRecipientCreate is the `DigestEmailRecipientCreate` schema.
type DigestEmailRecipientCreate struct {
	Email string `json:"email"`
}

// DigestEmailRecipientList is the `DigestEmailRecipientList` schema.
type DigestEmailRecipientList struct {
	Recipients []DigestEmailRecipient `json:"recipients"`
}

// DigestSendResult is the `DigestSendResult` schema.
type DigestSendResult struct {
	OK bool `json:"ok"`
	// Attempted: Deliveries attempted across Slack channels, Teams webhooks and
	// email recipients.
	Attempted int64                 `json:"attempted"`
	Succeeded int64                 `json:"succeeded"`
	Slack     DigestTransportResult `json:"slack"`
	Teams     DigestTransportResult `json:"teams"`
	Email     DigestTransportResult `json:"email"`
}

// DigestSettings is the `DigestSettings` schema.
type DigestSettings struct {
	// Enabled: Whether the weekly digest is enabled for this organization.
	// Delivery targets are the Slack channels and Teams webhooks whose
	// weeklyDigest trigger is on, plus the organization's digest email
	// recipients.
	Enabled bool `json:"enabled"`
	// LastSentWeekStart: Monday (ISO date, in the organization's timezone) of
	// the last week a digest covered, or null when none has been sent.
	LastSentWeekStart *string `json:"lastSentWeekStart"`
	// LastSentAt: When a digest last actually reached a destination, or null if
	// none ever has.
	LastSentAt *string `json:"lastSentAt"`
	// Timezone: IANA time zone the schedule and the Monday-to-Sunday week
	// boundary are expressed in. Defaults to UTC.
	Timezone string `json:"timezone"`
	// SendDay: ISO day of week the digest is sent on: 1 = Monday … 7 = Sunday.
	SendDay int64 `json:"sendDay"`
	// SendHour: Local hour (0–23) in `timezone` the digest is sent at.
	SendHour int64 `json:"sendHour"`
	// NarrativeEnabled: Whether an AI-written summary paragraph is placed above
	// the deterministic content. Opt-in, default off. Failures are non-fatal:
	// the digest still sends without the paragraph.
	NarrativeEnabled bool `json:"narrativeEnabled"`
	// NarrativeAvailable: Whether this deployment has an LLM API key configured.
	// False means enabling the narrative has no effect.
	NarrativeAvailable bool `json:"narrativeAvailable"`
	// EmailAvailable: Whether this deployment has a mail provider configured.
	// False means email recipients are never delivered to.
	EmailAvailable bool `json:"emailAvailable"`
	// AttemptCount: Delivery attempts made for lastSentWeekStart's window,
	// including the first.
	AttemptCount  int64   `json:"attemptCount"`
	LastAttemptAt *string `json:"lastAttemptAt"`
	// LastStatus: Outcome of the most recent delivery attempt. `partial` (some
	// destinations took it, some failed) is deliberately never retried
	// automatically — a retry would post the digest twice where it already
	// landed. `failed` (nothing landed) is retried a bounded number of times
	// with backoff, then parked until the next week.
	//
	// One of "pending", "succeeded", "partial", "failed", "no_targets".
	LastStatus *string `json:"lastStatus"`
	// LastError: Why the last attempt was not a clean success, for display in
	// the settings UI.
	LastError *string `json:"lastError"`
	// NextAttemptAt: When the next automatic retry is due, or null when none is
	// scheduled.
	NextAttemptAt *string `json:"nextAttemptAt"`
}

// DigestSettingsUpdate is the `DigestSettingsUpdate` schema.
type DigestSettingsUpdate struct {
	Enabled *bool `json:"enabled,omitempty"`
	// Timezone: IANA time zone name. Rejected with 400 if the server does not
	// know the zone.
	Timezone         *string `json:"timezone,omitempty"`
	SendDay          *int64  `json:"sendDay,omitempty"`
	SendHour         *int64  `json:"sendHour,omitempty"`
	NarrativeEnabled *bool   `json:"narrativeEnabled,omitempty"`
}

// DigestTransportResult is the `DigestTransportResult` schema.
type DigestTransportResult struct {
	Attempted int64 `json:"attempted"`
	Succeeded int64 `json:"succeeded"`
}

// DismissedAccessFinding is the `DismissedAccessFinding` schema.
type DismissedAccessFinding struct {
	// ResourceID: Infrawrench resource id the finding is on.
	ResourceID string `json:"resourceId"`
	// RuleID: Which rule was raised. Half of a dismissal's key, alongside the
	// resource id. The `access-review:` prefix is reserved so these can share
	// the posture dismissal store without colliding with plugin-declared posture
	// rule ids.
	//
	// One of "access-review:stale-principal", "access-review:admin-principal",
	// "access-review:key-past-rotation", "access-review:no-recorded-owner",
	// "access-review:no-mfa".
	RuleID string `json:"ruleId"`
	Title  string `json:"title"`
	// Severity: How bad the finding is. `critical` and `high` findings ride the
	// posture alert window; `medium` and `low` are review work surfaced on the
	// access review screen and in the weekly digest only.
	//
	// One of "critical", "high", "medium", "low".
	Severity string `json:"severity"`
	// Reason: Why this principal is flagged, in a sentence.
	Reason    string                `json:"reason"`
	Principal AccessPrincipal       `json:"principal"`
	Dismissal AccessReviewDismissal `json:"dismissal"`
}

// DismissedPostureFinding is the `DismissedPostureFinding` schema.
type DismissedPostureFinding struct {
	// ResourceID: Infrawrench resource id.
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	DisplayName      string   `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
	// RuleID: The matched rule's stable id, unique within the plugin.
	RuleID string `json:"ruleId"`
	// Title: Short rule title.
	Title string `json:"title"`
	// Severity: How bad the finding is. `critical` and `high` findings feed the
	// posture alerts; `medium` and `low` are hygiene work surfaced on the
	// posture screen only.
	//
	// One of "critical", "high", "medium", "low".
	Severity string `json:"severity"`
	// Category: Grouping bucket for what kind of exposure the finding describes.
	//
	// One of "public-exposure", "encryption", "credential-age",
	// "data-protection", "other".
	Category string `json:"category"`
	// Reason: Plugin-authored explanation of why this is a finding.
	Reason    string           `json:"reason"`
	Dismissal PostureDismissal `json:"dismissal"`
}

// DNSInventoryCounts: Record counts per status; zones counted separately.
//
// Spec schema: `DnsInventoryCounts`.
type DNSInventoryCounts struct {
	Zones       int64 `json:"zones"`
	Records     int64 `json:"records"`
	Owned       int64 `json:"owned"`
	Dangling    int64 `json:"dangling"`
	External    int64 `json:"external"`
	NotAnalysed int64 `json:"notAnalysed"`
}

// DNSInventoryResponse is the `DnsInventoryResponse` schema.
//
// Spec schema: `DnsInventoryResponse`.
type DNSInventoryResponse struct {
	// Zones: Sorted by domain, then account name.
	Zones []DNSZone `json:"zones"`
	// Records: Sorted worst status first, then by name.
	Records []DNSRecord        `json:"records"`
	Counts  DNSInventoryCounts `json:"counts"`
	// SkippedNamespaces: Provider namespaces that were declared but not
	// evaluated, and why — either no account for the plugin is connected, or no
	// claimant resource has synced. Both are missing data rather than a clean
	// bill of health, so they are reported rather than hidden.
	SkippedNamespaces []DNSSkippedNamespace `json:"skippedNamespaces"`
	GeneratedAt       string                `json:"generatedAt"`
}

// DNSRecord is the `DnsRecord` schema.
//
// Spec schema: `DnsRecord`.
type DNSRecord struct {
	// ResourceID: Infrawrench resource id of the record itself.
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	// ZoneResourceID: Owning zone's resource id, or null when the record could
	// not be attributed.
	ZoneResourceID *string `json:"zoneResourceId"`
	ZoneDomain     *string `json:"zoneDomain"`
	// Name: Fully qualified, lowercased, no trailing dot.
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	TTL      *float64 `json:"ttl"`
	Priority *float64 `json:"priority"`
	// Proxied: Whether the provider proxies the record (Cloudflare's orange
	// cloud).
	Proxied bool              `json:"proxied"`
	Targets []DNSRecordTarget `json:"targets"`
	// Status: Worst classification across `targets`.
	//
	// One of "owned", "dangling", "external", "not-analysed".
	Status string `json:"status"`
}

// DNSRecordTarget is the `DnsRecordTarget` schema.
//
// Spec schema: `DnsRecordTarget`.
type DNSRecordTarget struct {
	// Value: The target as stored, lowercased with any trailing dot removed.
	Value string `json:"value"`
	// Classification: What can be said about a record target from synced state
	// alone. `owned` — the value is an identity of a synced resource. `dangling`
	// — the value falls inside a provider namespace this workspace manages (an
	// S3 endpoint, a `*.vercel.app` alias) and no synced resource claims it,
	// which is the subdomain-takeover signature. `external` — the value points
	// somewhere there is no declaration for; not a finding. `not-analysed` — the
	// record type carries no host target that is reasoned about (TXT, MX, SOA,
	// CAA, SRV).
	//
	// One of "owned", "dangling", "external", "not-analysed".
	Classification string             `json:"classification"`
	Resource       *DNSTargetResource `json:"resource"`
	Service        *DNSTargetService  `json:"service"`
}

// DNSSkippedNamespace is the `DnsSkippedNamespace` schema.
//
// Spec schema: `DnsSkippedNamespace`.
type DNSSkippedNamespace struct {
	PluginID   PluginID `json:"pluginId"`
	PluginName string   `json:"pluginName"`
	Label      string   `json:"label"`
	Reason     string   `json:"reason"`
}

// DNSTargetResource: Set only when classification is "owned".
//
// Spec schema: `DnsTargetResource`.
//
// The API may send null in its place.
type DNSTargetResource struct {
	ResourceID       string   `json:"resourceId"`
	DisplayName      string   `json:"displayName"`
	PluginID         PluginID `json:"pluginId"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
}

// DNSTargetService: Set only when classification is "dangling".
//
// Spec schema: `DnsTargetService`.
//
// The API may send null in its place.
type DNSTargetService struct {
	PluginID       PluginID `json:"pluginId"`
	PluginName     string   `json:"pluginName"`
	ResourceTypeID string   `json:"resourceTypeId"`
	RuleID         string   `json:"ruleId"`
	Label          string   `json:"label"`
	// Severity: One of "critical", "high", "medium", "low".
	Severity string `json:"severity"`
	// Reason: Plugin-authored note on what claiming the name gets an attacker.
	Reason string `json:"reason"`
	// ClaimLabel: The instance-identifying part of the hostname, e.g. the bucket
	// or app name.
	ClaimLabel string `json:"claimLabel"`
}

// DNSZone is the `DnsZone` schema.
//
// Spec schema: `DnsZone`.
type DNSZone struct {
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	Domain           string   `json:"domain"`
	Status           *string  `json:"status"`
	// IsPrivate: Split-horizon/internal zone; listed but never analysed for
	// takeover.
	IsPrivate bool `json:"isPrivate"`
	// RecordCount: Records synced into this zone.
	RecordCount int64 `json:"recordCount"`
	// ProviderRecordCount: The provider's own record count, when reported. May
	// exceed `recordCount` — several plugins list zones without listing their
	// records.
	ProviderRecordCount *int64 `json:"providerRecordCount"`
	DanglingCount       int64  `json:"danglingCount"`
}

// DockerCommandRequest is the `DockerCommandRequest` schema.
type DockerCommandRequest struct {
	AccountID string     `json:"accountId"`
	Op        string     `json:"op"`
	Params    JSONObject `json:"params,omitempty"`
}

// DockerCommandResponse is the `DockerCommandResponse` schema.
type DockerCommandResponse struct {
	Result any `json:"result,omitempty"`
}

// DriftAlertSettings is the `DriftAlertSettings` schema.
type DriftAlertSettings struct {
	// NotifyCreated: Alert on resources that appeared.
	NotifyCreated bool `json:"notifyCreated"`
	// NotifyUpdated: Alert on field-level updates. Defaults to false — updates
	// are the bulk of the volume and are usually a provider restating a value.
	NotifyUpdated bool `json:"notifyUpdated"`
	// NotifyDeleted: Alert on resources that disappeared.
	NotifyDeleted bool `json:"notifyDeleted"`
	// CooldownMinutes: Least time between drift notifications for this
	// organization. One notification per window, no matter how many changes or
	// accounts it covers.
	CooldownMinutes int64 `json:"cooldownMinutes"`
	// MinChanges: Fewest matching changes in a window worth notifying about.
	MinChanges int64 `json:"minChanges"`
	// AccountIDs: Accounts to alert on. An empty array means every account.
	AccountIDs []string `json:"accountIds"`
	// LastNotifiedAt: When this organization last had a drift digest delivered.
	LastNotifiedAt *string `json:"lastNotifiedAt"`
}

// DriftAlertSettingsUpdate is the `DriftAlertSettingsUpdate` schema.
type DriftAlertSettingsUpdate struct {
	NotifyCreated   *bool    `json:"notifyCreated,omitempty"`
	NotifyUpdated   *bool    `json:"notifyUpdated,omitempty"`
	NotifyDeleted   *bool    `json:"notifyDeleted,omitempty"`
	CooldownMinutes *int64   `json:"cooldownMinutes,omitempty"`
	MinChanges      *int64   `json:"minChanges,omitempty"`
	AccountIDs      []string `json:"accountIds,omitempty"`
}

// DrillCoverageResponse is the `DrillCoverageResponse` schema.
type DrillCoverageResponse struct {
	Rows      []DrillCoverageRow `json:"rows"`
	Summary   DrillSummary       `json:"summary"`
	ValidDays int64              `json:"validDays"`
	// OrphanedDrills: Drills against a resource no longer in the inventory.
	// Reported rather than dropped: 'we tested this and then removed it' is a
	// fact an auditor asks about.
	OrphanedDrills []RestoreDrill `json:"orphanedDrills"`
	GeneratedAt    string         `json:"generatedAt"`
}

// DrillCoverageRow is the `DrillCoverageRow` schema.
type DrillCoverageRow struct {
	ResourceID     string  `json:"resourceId"`
	ResourceName   *string `json:"resourceName"`
	AccountID      *string `json:"accountId"`
	AccountName    *string `json:"accountName"`
	ResourceTypeID *string `json:"resourceTypeId"`
	// Standing: `never` and `stale` are kept apart because they call for
	// different conversations: one is 'nobody has ever tried', the other is 'it
	// worked in March'.
	//
	// One of "verified", "stale", "failed", "never".
	Standing    string  `json:"standing"`
	LastDrillAt *string `json:"lastDrillAt"`
	// LastOutcome: How the drill ended. Only `verified` counts as evidence the
	// backup works: a restore that produced a running system nobody looked
	// inside is exactly how a team discovers, mid-incident, that the dump had
	// been empty for months. `restored-unverified` is recorded because doing the
	// restore is worth recording, but it does not reset the clock.
	//
	// One of "verified", "restored-unverified", "failed", "blocked".
	LastOutcome        *string `json:"lastOutcome"`
	LastVerifiedAt     *string `json:"lastVerifiedAt"`
	VerifiedRtoMinutes *int64  `json:"verifiedRtoMinutes"`
	DaysUntilStale     *int64  `json:"daysUntilStale"`
}

// DrillSummary is the `DrillSummary` schema.
type DrillSummary struct {
	// EligibleCount: Resources with something to restore from. A resource with
	// no backup cannot be drilled, and listing it here would duplicate the
	// coverage page's own unprotected finding.
	EligibleCount int64 `json:"eligibleCount"`
	VerifiedCount int64 `json:"verifiedCount"`
	StaleCount    int64 `json:"staleCount"`
	FailedCount   int64 `json:"failedCount"`
	NeverCount    int64 `json:"neverCount"`
	// WorstRtoMinutes: Over currently-verified rows only; null when nothing is
	// verified, never zero.
	WorstRtoMinutes  *int64 `json:"worstRtoMinutes"`
	MedianRtoMinutes *int64 `json:"medianRtoMinutes"`
}

// EditableField is the `EditableField` schema.
type EditableField struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	// Kind: One of "string", "number", "boolean", "enum", "secret",
	// "association", "password".
	Kind        string   `json:"kind"`
	Required    bool     `json:"required"`
	Description *string  `json:"description,omitempty"`
	EnumValues  []string `json:"enumValues,omitempty"`
}

// EfficiencyAlertEvent is the `EfficiencyAlertEvent` schema.
type EfficiencyAlertEvent struct {
	ID string `json:"id"`
	// Kind: Which detector produced it.
	//
	// One of "commitment_expiry", "commitment_idle", "unit_cost_regression".
	Kind string `json:"kind"`
	// Subject: The commitment's description, or the business metric's name.
	Subject string `json:"subject"`
	// AccountID: The account, for commitment kinds; null otherwise.
	AccountID   *string `json:"accountId"`
	AccountName *string `json:"accountName"`
	// Currency: ISO 4217 of `amount`, or null when it carries none.
	Currency *string `json:"currency"`
	// Amount: The money at stake, in **units of `currency`** rather than cents —
	// commitment amounts are provider-reported in currency units. Per kind: the
	// monthly on-demand exposure for an expiry, the wasted amount for an idle
	// commitment, the current window's spend for a regression.
	Amount *float64 `json:"amount"`
	// Detail: Per-kind display facts. Free-form; nothing branches on it.
	Detail  map[string]any `json:"detail"`
	FiredAt string         `json:"firedAt"`
	// NotifiedAt: When the alert reached its routed destinations, or null when
	// nothing was routed (or the routing rule held it for quiet hours and the
	// follow-up pass has not run yet).
	NotifiedAt *string `json:"notifiedAt"`
}

// EnvironmentCaptureDraft is the `EnvironmentCaptureDraft` schema.
type EnvironmentCaptureDraft struct {
	Members             []EnvironmentCaptureDraftMember  `json:"members"`
	SuggestedParameters []EnvironmentParameter           `json:"suggestedParameters"`
	Skipped             []EnvironmentCaptureDraftSkipped `json:"skipped"`
}

// EnvironmentCaptureDraftMember is the `EnvironmentCaptureDraftMember` schema.
type EnvironmentCaptureDraftMember struct {
	// Key: Unique within the template; the id references are written against.
	Key              string   `json:"key"`
	PluginID         PluginID `json:"pluginId"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	AccountID        string   `json:"accountId"`
	SourceName       string   `json:"sourceName"`
	SourceResourceID *string  `json:"sourceResourceId,omitempty"`
	// NameFieldKey: The create-form field carrying the resource's name, detected
	// at capture by matching the captured value against the source's display
	// name. The instance name prefix is applied to this field and no other.
	NameFieldKey *string                                                `json:"nameFieldKey,omitempty"`
	ParentMember *string                                                `json:"parentMember,omitempty"`
	Fields       map[string]EnvironmentTemplateFieldValue               `json:"fields"`
	FieldMeta    map[string]EnvironmentCaptureDraftMemberFieldMetaValue `json:"fieldMeta"`
}

// EnvironmentCaptureRequest is the `EnvironmentCaptureRequest` schema.
type EnvironmentCaptureRequest struct {
	ResourceIDs []string `json:"resourceIds,omitempty"`
	AccountID   *string  `json:"accountId,omitempty"`
	TagKey      *string  `json:"tagKey,omitempty"`
	TagValue    *string  `json:"tagValue,omitempty"`
}

// EnvironmentCostEstimate is the `EnvironmentCostEstimate` schema.
type EnvironmentCostEstimate struct {
	// MonthlyAmount: Null means 'could not be priced', which is not the same as
	// zero.
	MonthlyAmount *float64 `json:"monthlyAmount"`
	Currency      *string  `json:"currency"`
	// Partial: True when at least one member is unpriced — read as 'at least'.
	Partial       bool                             `json:"partial"`
	UnpricedCount int64                            `json:"unpricedCount"`
	Members       []EnvironmentCostEstimateMembers `json:"members"`
}

// EnvironmentDiffEntry is the `EnvironmentDiffEntry` schema.
type EnvironmentDiffEntry struct {
	// Key: The pairing key both sides matched on — the resource type plus the
	// resource name with environment words removed. Stable across runs.
	Key              string `json:"key"`
	ResourceTypeID   string `json:"resourceTypeId"`
	ResourceTypeName string `json:"resourceTypeName"`
	// Status: Whether the slot exists on side A only, side B only, or on both
	// with a field divergence. Matched pairs that agree are counted in the type
	// summary rather than listed.
	//
	// One of "only-in-a", "only-in-b", "changed".
	Status string                      `json:"status"`
	A      *EnvironmentDiffResourceRef `json:"a"`
	B      *EnvironmentDiffResourceRef `json:"b"`
	// Changes: Field divergences. Empty unless `status` is `changed`.
	Changes []EnvironmentDiffFieldChange `json:"changes"`
	// SuppressedCount: Divergences hidden by the identity filter (ids, links,
	// addresses, timestamps). Always 0 when `includeIdentityFields` was
	// requested.
	SuppressedCount int64 `json:"suppressedCount"`
}

// EnvironmentDiffFieldChange is the `EnvironmentDiffFieldChange` schema.
type EnvironmentDiffFieldChange struct {
	// Field: Field key; resolved-output keys are prefixed `outputs.`.
	Field string `json:"field"`
	// A: Value on side A; null when the key is absent there.
	A any `json:"a,omitempty"`
	// B: Value on side B.
	B any `json:"b,omitempty"`
}

// EnvironmentDiffResourceRef: Null when the resource exists only on B.
//
// The API may send null in its place.
type EnvironmentDiffResourceRef struct {
	// ResourceID: Infrawrench resource id.
	ResourceID  string `json:"resourceId"`
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
}

// EnvironmentDiffResponse is the `EnvironmentDiffResponse` schema.
type EnvironmentDiffResponse struct {
	A          EnvironmentDiffSideSummary `json:"a"`
	B          EnvironmentDiffSideSummary `json:"b"`
	PluginID   PluginID                   `json:"pluginId"`
	PluginName string                     `json:"pluginName"`
	// Types: Every resource type present on either side, most-divergent first.
	Types []EnvironmentDiffTypeSummary `json:"types"`
	// Entries: Only the slots that differ; identical pairs are counted, not
	// listed.
	Entries []EnvironmentDiffEntry `json:"entries"`
	Totals  EnvironmentDiffTotals  `json:"totals"`
	// UnavailableTypes: Resource types excluded because they could not be
	// listed. Always empty over this API — it reads already-synced rows, which
	// cannot half-fail — and populated only by the desktop and CLI local modes,
	// which list live.
	UnavailableTypes      []EnvironmentDiffUnavailableType `json:"unavailableTypes"`
	IncludeIdentityFields bool                             `json:"includeIdentityFields"`
	GeneratedAt           string                           `json:"generatedAt"`
}

// EnvironmentDiffSideSummary is the `EnvironmentDiffSideSummary` schema.
type EnvironmentDiffSideSummary struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	// ResourceCount: Resources compared on this side.
	ResourceCount int64 `json:"resourceCount"`
}

// EnvironmentDiffTotals is the `EnvironmentDiffTotals` schema.
type EnvironmentDiffTotals struct {
	OnlyInA      int64 `json:"onlyInA"`
	OnlyInB      int64 `json:"onlyInB"`
	Changed      int64 `json:"changed"`
	Identical    int64 `json:"identical"`
	TypesOnlyInA int64 `json:"typesOnlyInA"`
	TypesOnlyInB int64 `json:"typesOnlyInB"`
	// SuppressedFieldChanges: Field divergences the identity filter hid across
	// every pair.
	SuppressedFieldChanges int64 `json:"suppressedFieldChanges"`
}

// EnvironmentDiffTypeSummary is the `EnvironmentDiffTypeSummary` schema.
type EnvironmentDiffTypeSummary struct {
	ResourceTypeID   string `json:"resourceTypeId"`
	ResourceTypeName string `json:"resourceTypeName"`
	CountA           int64  `json:"countA"`
	CountB           int64  `json:"countB"`
	// Delta: `countB - countA`.
	Delta   int64 `json:"delta"`
	OnlyInA int64 `json:"onlyInA"`
	OnlyInB int64 `json:"onlyInB"`
	// Changed: Matched pairs that disagree on at least one field.
	Changed int64 `json:"changed"`
	// Identical: Matched pairs with no visible divergence.
	Identical int64 `json:"identical"`
	// MissingFrom: Set when the resource type is absent from that side entirely.
	//
	// One of "a", "b".
	MissingFrom *string `json:"missingFrom"`
}

// EnvironmentDiffUnavailableType is the `EnvironmentDiffUnavailableType` schema.
type EnvironmentDiffUnavailableType struct {
	ResourceTypeID   string `json:"resourceTypeId"`
	ResourceTypeName string `json:"resourceTypeName"`
	// Message: The provider's complaint, as the lister reported it.
	Message string `json:"message"`
}

// EnvironmentEstimateRequest is the `EnvironmentEstimateRequest` schema.
type EnvironmentEstimateRequest struct {
	Parameters       map[string]string `json:"parameters,omitempty"`
	AccountOverrides map[string]string `json:"accountOverrides,omitempty"`
}

// EnvironmentInstance is the `EnvironmentInstance` schema.
type EnvironmentInstance struct {
	ID           string            `json:"id"`
	TemplateID   *string           `json:"templateId"`
	TemplateName string            `json:"templateName"`
	Name         string            `json:"name"`
	NamePrefix   string            `json:"namePrefix"`
	Parameters   map[string]string `json:"parameters"`
	// Status: `partial` means a create failed part-way: the members that were
	// created are recorded and can still be torn down, which is what stops a
	// half-finished run leaving cloud resources with no row pointing at them.
	//
	// One of "creating", "active", "partial", "tearing-down", "deleted",
	// "failed".
	Status      string                      `json:"status"`
	ExpiresAt   string                      `json:"expiresAt"`
	Error       *string                     `json:"error"`
	Members     []EnvironmentInstanceMember `json:"members"`
	CreatedAt   string                      `json:"createdAt"`
	UpdatedAt   string                      `json:"updatedAt"`
	CompletedAt *string                     `json:"completedAt"`
}

// EnvironmentInstanceConflict is the `EnvironmentInstanceConflict` schema.
type EnvironmentInstanceConflict struct {
	Error string `json:"error"`
}

// EnvironmentInstanceList is the `EnvironmentInstanceList` schema.
type EnvironmentInstanceList struct {
	Instances []EnvironmentInstance `json:"instances"`
}

// EnvironmentInstanceMember is the `EnvironmentInstanceMember` schema.
type EnvironmentInstanceMember struct {
	ID             string   `json:"id"`
	MemberKey      string   `json:"memberKey"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	AccountID      string   `json:"accountId"`
	ResourceID     *string  `json:"resourceId"`
	ExternalID     *string  `json:"externalId"`
	DisplayName    string   `json:"displayName"`
	// Status: One of "pending", "created", "failed", "deleted".
	Status string  `json:"status"`
	Error  *string `json:"error"`
	// LeaseID: The lease that auto-deletes this member at the TTL.
	LeaseID  *string `json:"leaseId"`
	Position int64   `json:"position"`
}

// EnvironmentInstantiateRequest is the `EnvironmentInstantiateRequest` schema.
type EnvironmentInstantiateRequest struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters,omitempty"`
	// TTLHours: Required. Capped by the org's `maxTtlHours` setting and by a
	// 720-hour ceiling.
	TTLHours         float64           `json:"ttlHours"`
	AccountOverrides map[string]string `json:"accountOverrides,omitempty"`
	Note             *string           `json:"note,omitempty"`
}

// EnvironmentParameter is the `EnvironmentParameter` schema.
type EnvironmentParameter struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	// Type: One of "string", "number", "select".
	Type         string                        `json:"type"`
	Required     bool                          `json:"required"`
	DefaultValue *string                       `json:"defaultValue,omitempty"`
	Options      []EnvironmentParameterOptions `json:"options,omitempty"`
	Description  *string                       `json:"description,omitempty"`
}

// EnvironmentSettings is the `EnvironmentSettings` schema.
type EnvironmentSettings struct {
	MaxTTLHours     int64 `json:"maxTtlHours"`
	DefaultTTLHours int64 `json:"defaultTtlHours"`
}

// EnvironmentStillLive is the `EnvironmentStillLive` schema.
type EnvironmentStillLive struct {
	Error string `json:"error"`
}

// EnvironmentTemplate is the `EnvironmentTemplate` schema.
type EnvironmentTemplate struct {
	ID                  string                      `json:"id"`
	Name                string                      `json:"name"`
	Description         *string                     `json:"description"`
	Parameters          []EnvironmentParameter      `json:"parameters"`
	Members             []EnvironmentTemplateMember `json:"members"`
	CreatedAt           string                      `json:"createdAt"`
	UpdatedAt           string                      `json:"updatedAt"`
	ActiveInstanceCount *int64                      `json:"activeInstanceCount,omitempty"`
}

// EnvironmentTemplateConflict is the `EnvironmentTemplateConflict` schema.
type EnvironmentTemplateConflict struct {
	Error string `json:"error"`
}

// EnvironmentTemplateFieldValue: What a captured create-form field is filled
// with at instantiation. `literal` is the captured value; `parameter` is a field
// the user chose to vary; `output` is another member's resolved output (a
// connection string, an IP — the captured half of an output reference);
// `member-id` is another member's provider-side id.
type EnvironmentTemplateFieldValue = any

// EnvironmentTemplateInput is the `EnvironmentTemplateInput` schema.
type EnvironmentTemplateInput struct {
	Name        string                      `json:"name"`
	Description *string                     `json:"description,omitempty"`
	Parameters  []EnvironmentParameter      `json:"parameters"`
	Members     []EnvironmentTemplateMember `json:"members"`
}

// EnvironmentTemplateList is the `EnvironmentTemplateList` schema.
type EnvironmentTemplateList struct {
	Templates []EnvironmentTemplate `json:"templates"`
}

// EnvironmentTemplateMember is the `EnvironmentTemplateMember` schema.
type EnvironmentTemplateMember struct {
	// Key: Unique within the template; the id references are written against.
	Key              string   `json:"key"`
	PluginID         PluginID `json:"pluginId"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	AccountID        string   `json:"accountId"`
	SourceName       string   `json:"sourceName"`
	SourceResourceID *string  `json:"sourceResourceId,omitempty"`
	// NameFieldKey: The create-form field carrying the resource's name, detected
	// at capture by matching the captured value against the source's display
	// name. The instance name prefix is applied to this field and no other.
	NameFieldKey *string                                  `json:"nameFieldKey,omitempty"`
	ParentMember *string                                  `json:"parentMember,omitempty"`
	Fields       map[string]EnvironmentTemplateFieldValue `json:"fields"`
}

// Error is the `Error` schema.
type Error struct {
	// Error: Human-readable error message
	Error string `json:"error"`
}

// EscalationPolicy: Notify these destinations too if nobody acknowledges within
// afterMinutes. Acknowledgement comes from the button on the Slack message, so
// an alert routed only to Teams or push will always escalate.
//
// The API may send null in its place.
type EscalationPolicy struct {
	AfterMinutes int64              `json:"afterMinutes"`
	Destinations []AlertDestination `json:"destinations"`
}

// ExchangeRate is the `ExchangeRate` schema.
type ExchangeRate struct {
	ID string `json:"id"`
	// FromCurrency: ISO 4217 code, upper-case.
	FromCurrency string `json:"fromCurrency"`
	// ToCurrency: ISO 4217 code, upper-case.
	ToCurrency string `json:"toCurrency"`
	// Rate: Multiply an amount in `fromCurrency` by this to get `toCurrency`. A
	// decimal **string**, not a number: it is stored in a `numeric(20, 10)`
	// column so the digits your finance system used survive the round trip
	// exactly, and a JSON number could not promise that.
	Rate string `json:"rate"`
	// EffectiveFrom: Inclusive day this rate starts applying. A given day
	// converts at the rate with the greatest `effectiveFrom` on or before it, so
	// historical periods keep the rate that applied then. A day earlier than
	// every stated rate has no rate.
	EffectiveFrom string  `json:"effectiveFrom"`
	CreatedBy     *string `json:"createdBy"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
}

// ExchangeRateInput is the `ExchangeRateInput` schema.
type ExchangeRateInput struct {
	// FromCurrency: ISO 4217 code, upper-case.
	FromCurrency string `json:"fromCurrency"`
	// ToCurrency: ISO 4217 code, upper-case.
	ToCurrency string `json:"toCurrency"`
	// Rate: Multiply an amount in `fromCurrency` by this to get `toCurrency`. A
	// decimal **string**, not a number: it is stored in a `numeric(20, 10)`
	// column so the digits your finance system used survive the round trip
	// exactly, and a JSON number could not promise that.
	Rate          string `json:"rate"`
	EffectiveFrom string `json:"effectiveFrom"`
}

// ExpiryAlertSettings is the `ExpiryAlertSettings` schema.
type ExpiryAlertSettings struct {
	// Enabled: Whether the poller sends expiry alerts for this organization at
	// all.
	Enabled bool `json:"enabled"`
	// LeadDays: Days of lead time before a deadline counts as `upcoming` and
	// alertable. Default 60.
	LeadDays int64 `json:"leadDays"`
	// LastNotifiedAt: When the organization's expiry alert scan last completed,
	// or null before the first. Owned by the poller's cooldown claim; not
	// writable through this API.
	LastNotifiedAt *string `json:"lastNotifiedAt"`
}

// ExpiryAlertSettingsUpdate is the `ExpiryAlertSettingsUpdate` schema.
type ExpiryAlertSettingsUpdate struct {
	Enabled  *bool  `json:"enabled,omitempty"`
	LeadDays *int64 `json:"leadDays,omitempty"`
}

// ExpiryItem is the `ExpiryItem` schema.
type ExpiryItem struct {
	// ResourceID: Infrawrench resource id.
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	DisplayName      string   `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
	// FieldKey: The declared field the deadline came from.
	FieldKey string `json:"fieldKey"`
	// Kind: Grouping bucket for the kind of deadline.
	//
	// One of "tls-cert", "domain", "api-token", "access-key", "k8s-cert",
	// "ssh-key", "secret-version", "other".
	Kind string `json:"kind"`
	// Label: Plugin-authored caption for the deadline.
	Label string `json:"label"`
	// Basis: `expiry` — the field held the deadline itself; `age` — the deadline
	// was derived from a creation/rotation date plus an age budget.
	//
	// One of "expiry", "age".
	Basis string `json:"basis"`
	// DueAt: The deadline.
	DueAt string `json:"dueAt"`
	// DaysRemaining: Whole days until dueAt (floor); negative once expired.
	DaysRemaining int64 `json:"daysRemaining"`
	// Severity: How close the deadline is: `expired` (in the past), `critical`
	// (due within 7 days), `warning` (within 30 days), `upcoming` (within the
	// organization's lead time), or `ok` (tracked, but further out than the lead
	// time).
	//
	// One of "expired", "critical", "warning", "upcoming", "ok".
	Severity string `json:"severity"`
}

// ExpiryListResponse is the `ExpiryListResponse` schema.
type ExpiryListResponse struct {
	// Items: All tracked deadlines, soonest first (`ok` items included).
	Items      []ExpiryItem         `json:"items"`
	TotalCount int64                `json:"totalCount"`
	Counts     ExpirySeverityCounts `json:"counts"`
	// LeadDays: The lead time the `upcoming` bucket was computed against.
	LeadDays    int64  `json:"leadDays"`
	GeneratedAt string `json:"generatedAt"`
}

// ExpirySeverityCounts: Item count per severity; every bucket present, zeros
// included.
type ExpirySeverityCounts struct {
	Expired  int64 `json:"expired"`
	Critical int64 `json:"critical"`
	Warning  int64 `json:"warning"`
	Upcoming int64 `json:"upcoming"`
	OK       int64 `json:"ok"`
}

// ExportCredentialRequest is the `ExportCredentialRequest` schema.
type ExportCredentialRequest struct {
	ResourceID       ResourceID  `json:"resourceId"`
	AccountID        string      `json:"accountId"`
	FormatID         string      `json:"formatId"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// ExportTerraformRequest is the `ExportTerraformRequest` schema.
type ExportTerraformRequest struct {
	ResourceID ResourceID `json:"resourceId"`
	AccountID  string     `json:"accountId"`
}

// FieldActionRequest is the `FieldActionRequest` schema.
type FieldActionRequest struct {
	AccountID        string            `json:"accountId"`
	ResourceTypeID   string            `json:"resourceTypeId"`
	FieldKey         string            `json:"fieldKey"`
	ActionID         string            `json:"actionId"`
	Fields           map[string]string `json:"fields"`
	ActionFields     map[string]string `json:"actionFields,omitempty"`
	PluginID         *string           `json:"pluginId,omitempty"`
	ParentResourceID *ResourceID       `json:"parentResourceId,omitempty"`
}

// FieldActionResponse is the `FieldActionResponse` schema.
type FieldActionResponse struct {
	Value  string                     `json:"value"`
	Option *FieldActionResponseOption `json:"option,omitempty"`
}

// GenerateSSHKeyRequest is the `GenerateSshKeyRequest` schema.
//
// Spec schema: `GenerateSshKeyRequest`.
type GenerateSSHKeyRequest struct {
	Name string `json:"name"`
}

// GeneratedSSHKey is the `GeneratedSshKey` schema.
//
// Spec schema: `GeneratedSshKey`.
type GeneratedSSHKey struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	KeyType     SSHKeyType `json:"keyType"`
	Fingerprint string     `json:"fingerprint"`
	PublicKey   string     `json:"publicKey"`
	// PrivateKey: Returned once. Not persisted in plaintext.
	PrivateKey string `json:"privateKey"`
}

// HygieneFinding is the `HygieneFinding` schema.
type HygieneFinding struct {
	// ID: Stable across runs, so a client can remember what has been reviewed.
	ID string `json:"id"`
	// Kind: One of "api_key_never_used", "api_key_idle",
	// "api_key_expired_not_revoked", "api_key_wildcard_scope",
	// "api_key_unused_scopes", "ssh_key_never_used", "ssh_key_idle",
	// "member_unused_permissions".
	Kind string `json:"kind"`
	// Severity: One of "high", "medium", "low".
	Severity string `json:"severity"`
	Title    string `json:"title"`
	// Detail: The evidence behind the finding.
	Detail         string `json:"detail"`
	Recommendation string `json:"recommendation"`
	// EntityType: One of "api-key", "ssh-key", "member".
	EntityType string `json:"entityType"`
	EntityID   string `json:"entityId"`
	EntityName string `json:"entityName"`
	// Facts: Structured detail for table columns and reports.
	Facts map[string]any `json:"facts"`
}

// HygieneReport is the `HygieneReport` schema.
type HygieneReport struct {
	GeneratedAt string `json:"generatedAt"`
	WindowDays  int64  `json:"windowDays"`
	// AuditHistoryDays: How much audit history the organization actually has;
	// null when it has none.
	AuditHistoryDays *int64 `json:"auditHistoryDays"`
	// PermissionFindingsWithheld: True when there was not enough audit history
	// for the unused-permission finding to mean anything, so it was withheld
	// rather than guessed at.
	PermissionFindingsWithheld bool                `json:"permissionFindingsWithheld"`
	Findings                   []HygieneFinding    `json:"findings"`
	Counts                     HygieneReportCounts `json:"counts"`
}

// IacFieldChange is the `IacFieldChange` schema.
type IacFieldChange struct {
	Field string `json:"field"`
	// From: The value Terraform state carries.
	From any `json:"from,omitempty"`
	// To: The value actually running.
	To any `json:"to,omitempty"`
}

// IacImportPlanRequest is the `IacImportPlanRequest` schema.
type IacImportPlanRequest struct {
	ResourceIDs []string `json:"resourceIds"`
}

// IacImportPlanResponse is the `IacImportPlanResponse` schema.
type IacImportPlanResponse struct {
	// Hcl: `import` blocks followed by the generated resource stanzas.
	Hcl         string                             `json:"hcl"`
	Exported    []IacImportPlanResponseExported    `json:"exported"`
	Unsupported []IacImportPlanResponseUnsupported `json:"unsupported"`
}

// IacReconciledResource is the `IacReconciledResource` schema.
type IacReconciledResource struct {
	ResourceID     string   `json:"resourceId"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	AccountID      string   `json:"accountId"`
	DisplayName    string   `json:"displayName"`
	ExternalID     *string  `json:"externalId"`
	// Status: `managed`: matched a state entry and agrees with it. `drifted`:
	// matched, but live fields differ. `unmanaged`: in inventory, absent from
	// state — somebody made it by hand.
	//
	// One of "managed", "drifted", "unmanaged".
	Status           string  `json:"status"`
	TerraformType    *string `json:"terraformType"`
	TerraformAddress *string `json:"terraformAddress"`
	// MatchedBy: How the match was made, so it can be argued with.
	//
	// One of "import-id", "external-id", "identifier".
	MatchedBy *string          `json:"matchedBy"`
	Drift     []IacFieldChange `json:"drift"`
	// UnmappableReason: Set when no Terraform block could be produced for this
	// resource, which makes its drift unknowable. Never reported as "no drift".
	UnmappableReason *string `json:"unmappableReason"`
	// Owner: Resource owner annotation, populated for unmanaged resources.
	Owner map[string]any `json:"owner"`
	// FirstSeenAt: When the change timeline first recorded this resource
	// appearing.
	FirstSeenAt *string `json:"firstSeenAt"`
}

// IacReconciliationResponse is the `IacReconciliationResponse` schema.
type IacReconciliationResponse struct {
	State     IacState                `json:"state"`
	Resources []IacReconciledResource `json:"resources"`
	// StateOnly: State entries with no inventory match — their own category.
	StateOnly []IacStateOnlyResource           `json:"stateOnly"`
	Summary   IacReconciliationResponseSummary `json:"summary"`
	// Underivable: Plugin resource types whose Terraform type could not be
	// derived from the plugin's own export mapper. Reported rather than guessed.
	Underivable []IacReconciliationResponseUnderivable `json:"underivable"`
}

// IacResourceStatusResponse is the `IacResourceStatusResponse` schema.
type IacResourceStatusResponse struct {
	// Status: One of "managed", "drifted", "unmanaged".
	Status           *string `json:"status"`
	StateID          *string `json:"stateId"`
	StateLabel       *string `json:"stateLabel"`
	TerraformAddress *string `json:"terraformAddress"`
	DriftFieldCount  int64   `json:"driftFieldCount"`
}

// IacState is the `IacState` schema.
type IacState struct {
	ID string `json:"id"`
	// Label: User-supplied name for this state, e.g. "prod / us-east-1".
	Label string `json:"label"`
	// AccountID: The account this state covers, or null when it covers the whole
	// organization.
	AccountID   *string `json:"accountId"`
	AccountName *string `json:"accountName"`
	// Format: Which document shape was uploaded: a raw state file, or `terraform
	// show -json`.
	//
	// One of "tfstate", "show-json".
	Format string `json:"format"`
	// FormatVersion: The document's own version — "4" for a state file,
	// "1.0"-style otherwise.
	FormatVersion    string  `json:"formatVersion"`
	TerraformVersion *string `json:"terraformVersion"`
	// Serial: State file serial; null for show output.
	Serial *int64 `json:"serial"`
	// Lineage: State file lineage; null for show output.
	Lineage *string `json:"lineage"`
	// ResourceCount: Managed resource instances recorded.
	ResourceCount int64 `json:"resourceCount"`
	// DataSourceCount: Data-source entries, recorded but never matched against
	// inventory.
	DataSourceCount int64 `json:"dataSourceCount"`
	// RedactedAttributeCount: Attribute values dropped because the state marked
	// them sensitive. Redaction happens at parse time — no sensitive value is
	// ever stored.
	RedactedAttributeCount int64    `json:"redactedAttributeCount"`
	ParseWarnings          []string `json:"parseWarnings"`
	UploadedByUserID       *string  `json:"uploadedByUserId"`
	UploadedByName         *string  `json:"uploadedByName"`
	CreatedAt              string   `json:"createdAt"`
}

// IacStateListResponse is the `IacStateListResponse` schema.
type IacStateListResponse struct {
	States []IacState `json:"states"`
}

// IacStateOnlyResource is the `IacStateOnlyResource` schema.
type IacStateOnlyResource struct {
	Address       string                           `json:"address"`
	TerraformType string                           `json:"terraformType"`
	Identifiers   []string                         `json:"identifiers"`
	Candidates    []IacStateOnlyResourceCandidates `json:"candidates"`
	// Reason: One of "no-inventory-match", "unknown-terraform-type".
	Reason string `json:"reason"`
}

// IacStateUploadRequest is the `IacStateUploadRequest` schema.
type IacStateUploadRequest struct {
	Label     string  `json:"label"`
	AccountID *string `json:"accountId,omitempty"`
	// Document: The state document, as text: a raw `.tfstate` (format version 4)
	// or the output of `terraform show -json` (format_version 1.x). Limited to 8
	// MiB.
	Document string `json:"document"`
}

// ImportSSHKeyRequest is the `ImportSshKeyRequest` schema.
//
// Spec schema: `ImportSshKeyRequest`.
type ImportSSHKeyRequest struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

// ImportYAMLRequest is the `ImportYamlRequest` schema.
//
// Spec schema: `ImportYamlRequest`.
type ImportYAMLRequest struct {
	AccountID        string      `json:"accountId"`
	YAML             string      `json:"yaml"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// ImportedSSHKey is the `ImportedSshKey` schema.
//
// Spec schema: `ImportedSshKey`.
type ImportedSSHKey struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	KeyType     SSHKeyType `json:"keyType"`
	Fingerprint string     `json:"fingerprint"`
	PublicKey   string     `json:"publicKey"`
	IsImported  bool       `json:"isImported"`
}

// Incident is the `Incident` schema.
type Incident struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	// Severity: Severity in the ordinary sev1..sev4 register. `sev1` is a
	// complete outage; `sev4` is cosmetic and tracked rather than paged.
	//
	// One of "sev1", "sev2", "sev3", "sev4".
	Severity string `json:"severity"`
	// Status: `mitigated` is a real state, not a synonym for resolved: impact
	// has stopped but the incident is still open for follow-up. Keeping it
	// separate is what makes time-to-mitigate a measurement rather than a guess.
	// Resolving runs the resolve path — the change freeze this incident opened
	// is lifted, and the status-page update it posted is closed.
	//
	// One of "open", "mitigated", "resolved".
	Status  string  `json:"status"`
	Summary *string `json:"summary"`
	// StartedAt: Backdatable — people declare after they start firefighting.
	StartedAt        string  `json:"startedAt"`
	MitigatedAt      *string `json:"mitigatedAt"`
	ResolvedAt       *string `json:"resolvedAt"`
	DeclaredByUserID *string `json:"declaredByUserId"`
	DeclaredByName   *string `json:"declaredByName"`
	ResolvedByUserID *string `json:"resolvedByUserId"`
	// AffectedResourceIDs: Advisory. Not foreign keys — the claim must survive
	// the resource being deleted.
	AffectedResourceIDs []string `json:"affectedResourceIds"`
	AffectedAccountIDs  []string `json:"affectedAccountIds"`
	// IssueURL: Where the write-up was filed, once anyone filed it.
	IssueURL  *string            `json:"issueUrl"`
	CreatedAt string             `json:"createdAt"`
	UpdatedAt string             `json:"updatedAt"`
	Artifacts []IncidentArtifact `json:"artifacts"`
	NoteCount int64              `json:"noteCount"`
}

// IncidentActions is the `IncidentActions` schema.
type IncidentActions struct {
	// OpenFreeze: Open an org change freeze for the duration, lifted when the
	// incident resolves. Defaults to false — freezing has blast radius beyond
	// the incident. Needs `freezes:write`; without it the freeze is recorded as
	// a failed artefact naming the permission, and the incident still stands.
	OpenFreeze *bool `json:"openFreeze,omitempty"`
	// PinMoment: Pin the moment (a timestamp and a window) so `GET /moment` is
	// one click away. Defaults to true — it cannot fail, and the investigation
	// always wants it.
	PinMoment *bool `json:"pinMoment,omitempty"`
	// PostSlack: Announce through the org's alert routing rules under the
	// `incidentAlerts` trigger, so channels, quiet hours, escalation and the
	// acknowledge button all apply unchanged. Defaults to true. If no rule
	// matches, the artefact fails and says so.
	PostSlack *bool `json:"postSlack,omitempty"`
	// StatusPageID: Post a public update on this status page. Omitted means no
	// public update.
	StatusPageID *string `json:"statusPageId,omitempty"`
	// StatusPageComponentIDs: Components on that page to mark affected. Empty
	// means the page as a whole.
	StatusPageComponentIDs []string `json:"statusPageComponentIds,omitempty"`
}

// IncidentArtifact is the `IncidentArtifact` schema.
type IncidentArtifact struct {
	ID string `json:"id"`
	// Kind: Which side effect of declaring this artefact records.
	//
	// One of "freeze", "moment", "slack", "status-page".
	Kind string `json:"kind"`
	// Status: `failed` is a stored state, not an error: declaring writes the
	// incident first and attempts each opted-in side effect afterwards, so a
	// Slack outage costs the announcement and never the incident. A failed
	// artefact carries its error and can be retried.
	//
	// `close_failed` is the other half and is deliberately distinct: the
	// artefact **was** created and resolving could not put it away, so the
	// change freeze is still in force or the public notice still reports an
	// outage. Retrying a `failed` artefact re-creates it; retrying a
	// `close_failed` one re-closes it. Collapsing the two would either strand
	// the incident with a live freeze nothing can lift, or open a second freeze.
	//
	// One of "created", "failed", "closed", "close_failed".
	Status string `json:"status"`
	// Label: Human label — the freeze name, the destination count.
	Label *string `json:"label"`
	// RefID: Freeze id, notice id, Slack channel id…
	RefID *string `json:"refId"`
	// RefSecondary: Second half of a compound reference — a Slack message ts, a
	// window width.
	RefSecondary *string `json:"refSecondary"`
	// Error: Why it failed. Null unless `status` is `failed` or `close_failed`.
	Error     *string                  `json:"error"`
	Request   *IncidentArtifactRequest `json:"request"`
	CreatedAt string                   `json:"createdAt"`
	UpdatedAt string                   `json:"updatedAt"`
}

// IncidentArtifactRequest: What the declaration asked for, recorded so a retry
// asks for the same thing. Present on the status-page artefact, where a retry
// that forgot the operator's chosen components would publish the outage against
// the whole page.
//
// The API may send null in its place.
type IncidentArtifactRequest struct {
	StatusPageID *string  `json:"statusPageId,omitempty"`
	ComponentIDs []string `json:"componentIds,omitempty"`
}

// IncidentDeclare is the `IncidentDeclare` schema.
type IncidentDeclare struct {
	Title string `json:"title"`
	// Severity: Severity in the ordinary sev1..sev4 register. `sev1` is a
	// complete outage; `sev4` is cosmetic and tracked rather than paged.
	//
	// One of "sev1", "sev2", "sev3", "sev4".
	Severity *string `json:"severity,omitempty"`
	Summary  *string `json:"summary,omitempty"`
	// StartedAt: Defaults to now.
	StartedAt           *string          `json:"startedAt,omitempty"`
	AffectedResourceIDs []string         `json:"affectedResourceIds,omitempty"`
	AffectedAccountIDs  []string         `json:"affectedAccountIds,omitempty"`
	Actions             *IncidentActions `json:"actions,omitempty"`
}

// IncidentDetail is the `IncidentDetail` schema.
type IncidentDetail struct {
	Incident Incident       `json:"incident"`
	Notes    []IncidentNote `json:"notes"`
}

// IncidentList is the `IncidentList` schema.
type IncidentList struct {
	Incidents []Incident `json:"incidents"`
}

// IncidentNote is the `IncidentNote` schema.
type IncidentNote struct {
	ID           string  `json:"id"`
	Body         string  `json:"body"`
	AuthorUserID *string `json:"authorUserId"`
	AuthorName   *string `json:"authorName"`
	// OccurredAt: When the note is *about*, which may precede when it was
	// written — a note typed at 04:00 can be dated to 03:14 and lands there on
	// the timeline.
	OccurredAt string `json:"occurredAt"`
	CreatedAt  string `json:"createdAt"`
}

// IncidentNoteCreate is the `IncidentNoteCreate` schema.
type IncidentNoteCreate struct {
	Body string `json:"body"`
	// OccurredAt: Defaults to now; backdate to place the note.
	OccurredAt *string `json:"occurredAt,omitempty"`
}

// IncidentPatch is the `IncidentPatch` schema.
type IncidentPatch struct {
	Title *string `json:"title,omitempty"`
	// Severity: Severity in the ordinary sev1..sev4 register. `sev1` is a
	// complete outage; `sev4` is cosmetic and tracked rather than paged.
	//
	// One of "sev1", "sev2", "sev3", "sev4".
	Severity *string `json:"severity,omitempty"`
	// Status: `mitigated` is a real state, not a synonym for resolved: impact
	// has stopped but the incident is still open for follow-up. Keeping it
	// separate is what makes time-to-mitigate a measurement rather than a guess.
	// Resolving runs the resolve path — the change freeze this incident opened
	// is lifted, and the status-page update it posted is closed.
	//
	// One of "open", "mitigated", "resolved".
	Status              *string  `json:"status,omitempty"`
	Summary             *string  `json:"summary,omitempty"`
	AffectedResourceIDs []string `json:"affectedResourceIds,omitempty"`
	AffectedAccountIDs  []string `json:"affectedAccountIds,omitempty"`
	IssueURL            *string  `json:"issueUrl,omitempty"`
}

// IncidentPostmortem is the `IncidentPostmortem` schema.
type IncidentPostmortem struct {
	Markdown string `json:"markdown"`
	Filename string `json:"filename"`
}

// IncidentTimeline is the `IncidentTimeline` schema.
type IncidentTimeline struct {
	IncidentID string `json:"incidentId"`
	From       string `json:"from"`
	// To: `resolvedAt`, or the server's clock while the incident is open.
	To          string                  `json:"to"`
	GeneratedAt string                  `json:"generatedAt"`
	Entries     []IncidentTimelineEntry `json:"entries"`
	// Feeds: Per-feed health, passed through from the moment union: `omitted`
	// means the caller lacks that feed's read permission, `error` means it
	// failed and the rest is still good.
	Feeds     []IncidentTimelineFeeds `json:"feeds"`
	Truncated bool                    `json:"truncated"`
}

// IncidentTimelineEntry is the `IncidentTimelineEntry` schema.
type IncidentTimelineEntry struct {
	ID string `json:"id"`
	// Source: `moment` covers everything the moment union already indexes —
	// resource changes, deployments, cost anomalies, provider status incidents,
	// audit entries, change freezes and workflow runs. Nothing is copied into
	// the incident's own tables; the timeline is a join, so re-reading it
	// reflects the record as it stands today.
	//
	// One of "incident", "note", "artifact", "moment", "probe", "metric-alert".
	Source string `json:"source"`
	// Kind: `<noun>.<verb>`. Open set — render unknown kinds generically.
	Kind   string  `json:"kind"`
	At     string  `json:"at"`
	Title  string  `json:"title"`
	Detail *string `json:"detail,omitempty"`
	// Severity: One of "info", "warning", "critical".
	Severity     string                     `json:"severity"`
	AuthorName   *string                    `json:"authorName,omitempty"`
	ResourceID   *string                    `json:"resourceId,omitempty"`
	ResourceName *string                    `json:"resourceName,omitempty"`
	PluginID     *string                    `json:"pluginId,omitempty"`
	AccountID    *string                    `json:"accountId,omitempty"`
	Link         *IncidentTimelineEntryLink `json:"link,omitempty"`
}

// Invitation is the `Invitation` schema.
type Invitation struct {
	ID         string           `json:"id"`
	Email      string           `json:"email"`
	Role       OrganizationRole `json:"role"`
	RoleID     *string          `json:"roleId"`
	RoleName   *string          `json:"roleName"`
	AcceptedAt *string          `json:"acceptedAt"`
	ExpiresAt  string           `json:"expiresAt"`
	CreatedAt  string           `json:"createdAt"`
}

// InvitationDetail is the `InvitationDetail` schema.
type InvitationDetail struct {
	ID               string           `json:"id"`
	Email            string           `json:"email"`
	Role             OrganizationRole `json:"role"`
	ExpiresAt        string           `json:"expiresAt"`
	AcceptedAt       *string          `json:"acceptedAt"`
	OrganizationID   string           `json:"organizationId"`
	OrganizationName string           `json:"organizationName"`
}

// InviteRequest is the `InviteRequest` schema.
type InviteRequest struct {
	Email  string            `json:"email"`
	Role   *OrganizationRole `json:"role,omitempty"`
	RoleID *string           `json:"roleId,omitempty"`
	// AddSeat: When the paid plan is full (409 seat_limit_reached), retry with
	// this set to buy one more monthly seat and send the invitation. Requires
	// billing:write. Only works when the 409 reported `canAddSeat: true` — an
	// org whose capacity is entirely prepaid capacity slots has no monthly seat
	// to add.
	AddSeat *bool `json:"addSeat,omitempty"`
}

// InviteResponse is the `InviteResponse` schema.
type InviteResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

// Invoice is the `Invoice` schema.
type Invoice struct {
	ID                 string `json:"id"`
	ManagedAccountID   string `json:"managedAccountId"`
	ManagedAccountName string `json:"managedAccountName"`
	// Number: `INV-2026-0001`. Null while draft — numbers are assigned at
	// approval so a deleted draft cannot leave a gap in the sequence.
	Number                *string           `json:"number"`
	Status                InvoiceStatus     `json:"status"`
	PeriodFrom            string            `json:"periodFrom"`
	PeriodTo              string            `json:"periodTo"`
	Currency              string            `json:"currency"`
	Totals                *InvoiceTotals    `json:"totals,omitempty"`
	IssuedAt              *string           `json:"issuedAt"`
	SentAt                *string           `json:"sentAt"`
	Delivery              *InvoiceDelivery  `json:"delivery"`
	VoidedAt              *string           `json:"voidedAt"`
	VoidReason            *string           `json:"voidReason"`
	SupersedesInvoiceID   *string           `json:"supersedesInvoiceId"`
	SupersededByInvoiceID *string           `json:"supersededByInvoiceId"`
	CreatedAt             string            `json:"createdAt"`
	UpdatedAt             string            `json:"updatedAt"`
	Notes                 *string           `json:"notes"`
	Lines                 []InvoiceLine     `json:"lines"`
	Derivation            InvoiceDerivation `json:"derivation"`
	// Live: True when the figures in this response were recomputed for it — true
	// for a draft, false for everything else. Say so: “these numbers will move”
	// and “these numbers are what we sent” are different claims about the same
	// fields.
	Live             bool    `json:"live"`
	ComputedAt       string  `json:"computedAt"`
	ApprovedByUserID *string `json:"approvedByUserId"`
	SentByUserID     *string `json:"sentByUserId"`
	VoidedByUserID   *string `json:"voidedByUserId"`
	CreatedByUserID  *string `json:"createdByUserId"`
}

// InvoiceDelivery: The last delivery attempt, or null when none has been made —
// including on an invoice marked sent by a deployment with no mail provider. “A
// person released this” and “we delivered it” are different claims, and this
// field is only ever the second.
//
// The API may send null in its place.
type InvoiceDelivery struct {
	// Status: `pending` means an attempt was claimed and its outcome never
	// recorded — the process died mid-send, so whether the customer received it
	// is unknown. It is not a failure and is never retried automatically.
	//
	// One of "pending", "succeeded", "partial", "failed", "no_targets".
	Status string `json:"status"`
	// Recipients: The addresses this attempt was made to, as the customer record
	// had them then.
	Recipients []string `json:"recipients"`
	// Delivered: How many the mail provider accepted.
	Delivered   int64  `json:"delivered"`
	AttemptedAt string `json:"attemptedAt"`
	// DeliveredAt: The last attempt that reached at least one address, or null
	// when none ever has. Never cleared by a later failure — it is a fact about
	// the past, and it is what decides whether sending again is a retry or a
	// second copy.
	DeliveredAt *string `json:"deliveredAt"`
	Attempts    int64   `json:"attempts"`
	Error       *string `json:"error"`
}

// InvoiceDerivation: Everything needed to re-derive the invoice by hand. Not
// decoration: an invoice a customer cannot reconcile is an invoice a customer
// does not pay.
type InvoiceDerivation struct {
	// CostBasis: One of "cash", "amortized".
	CostBasis         string `json:"costBasis"`
	ApplyBillingRules bool   `json:"applyBillingRules"`
	// RateDate: The day the exchange rates were read — always the period's last
	// day. One rate for the period rather than a per-day blend: “January, at the
	// 31 January rate” is a sentence a finance team can reproduce.
	RateDate string                   `json:"rateDate"`
	Rates    []InvoiceDerivationRates `json:"rates"`
	// Unconverted: Currencies the organisation had stated no usable rate for. A
	// non-empty list blocks approval: an invoice that cannot be expressed as one
	// number in the customer's currency must not be frozen.
	Unconverted []string                 `json:"unconverted"`
	Rules       []InvoiceDerivationRules `json:"rules"`
	Scope       InvoiceDerivationScope   `json:"scope"`
	// MissingScope: Scope entries that no longer exist. Recorded rather than
	// silently skipped — an invoice that is quietly short is worse than one that
	// says why.
	MissingScope []string `json:"missingScope"`
}

// InvoiceInput: A new invoice is always a draft. There is no status field and no
// scope field: generating and issuing are two acts, and the scope comes from the
// customer.
type InvoiceInput struct {
	ManagedAccountID string  `json:"managedAccountId"`
	PeriodFrom       string  `json:"periodFrom"`
	PeriodTo         string  `json:"periodTo"`
	Notes            *string `json:"notes,omitempty"`
	// SupersedesInvoiceID: The void invoice this one corrects. The original must
	// already be void — a correction that leaves the original standing means the
	// customer holds two live invoices for one period.
	SupersedesInvoiceID *string `json:"supersedesInvoiceId,omitempty"`
}

// InvoiceLine: One scope entry in one collected currency. Two currencies for one
// cost centre are two lines, not one blended line, because the conversion is a
// separately reconcilable step.
type InvoiceLine struct {
	// Kind: One of "cost_centre", "account", "fixed".
	Kind string `json:"kind"`
	// RefID: Cost-centre id, account id, or null for an org-level fixed charge.
	RefID *string `json:"refId"`
	// Label: The name at issue time, frozen with the numbers — renaming a cost
	// centre in March must not retitle a line on January's invoice.
	Label string `json:"label"`
	// Currency: The currency the providers billed in.
	Currency string `json:"currency"`
	// Collected: What the providers charged for this scope, before any billing
	// rule.
	Collected float64 `json:"collected"`
	// Adjustment: What the organisation's billing rules added or removed.
	Adjustment float64 `json:"adjustment"`
	// Adjusted: `collected + adjustment`.
	Adjusted float64 `json:"adjusted"`
	// Rate: The rate applied to reach `billed`. 1 when the line is already in
	// the invoice currency; null when the organisation has stated no rate for
	// this currency, in which case the amount is carried in its own currency
	// rather than dropped or invented.
	Rate *float64 `json:"rate"`
	// Billed: `adjusted × rate`, in the invoice currency.
	Billed *float64 `json:"billed"`
}

// InvoiceSendRequest is the `InvoiceSendRequest` schema.
type InvoiceSendRequest struct {
	// Resend: Send another copy of an invoice that has already reached somebody.
	// Required only in that case: retrying a delivery that reached nobody
	// (`failed`, `no_targets`) needs no flag, because there is no inbox to
	// duplicate into. Refused with 409 without it when the last attempt landed,
	// or when its outcome is unknown (`pending`).
	Resend *bool `json:"resend,omitempty"`
}

// InvoiceStatus: `draft` → `approved` → `sent`, plus `void` from either issued
// state.
//
// **A draft recomputes its figures from live spend on every read; an approved,
// sent or void invoice never does.** Approval is the freeze: the lines, the
// totals, the exchange rates and the day they were read, the billing rules in
// force and the names of everything in scope are written onto the invoice, and
// no later restatement of spend, change of rate, edit of a rule or rename can
// alter what the document says.
//
// An issued invoice is never edited and never deleted. A wrong one is voided
// with a reason and superseded by a corrective invoice; both survive. The server
// enforces this, not just the UI.
type InvoiceStatus = string

// The values InvoiceStatus takes.
const (
	InvoiceStatusDraft    InvoiceStatus = "draft"
	InvoiceStatusApproved InvoiceStatus = "approved"
	InvoiceStatusSent     InvoiceStatus = "sent"
	InvoiceStatusVoid     InvoiceStatus = "void"
)

// InvoiceSummary is the `InvoiceSummary` schema.
type InvoiceSummary struct {
	ID                 string `json:"id"`
	ManagedAccountID   string `json:"managedAccountId"`
	ManagedAccountName string `json:"managedAccountName"`
	// Number: `INV-2026-0001`. Null while draft — numbers are assigned at
	// approval so a deleted draft cannot leave a gap in the sequence.
	Number                *string          `json:"number"`
	Status                InvoiceStatus    `json:"status"`
	PeriodFrom            string           `json:"periodFrom"`
	PeriodTo              string           `json:"periodTo"`
	Currency              string           `json:"currency"`
	Totals                *InvoiceTotals   `json:"totals"`
	IssuedAt              *string          `json:"issuedAt"`
	SentAt                *string          `json:"sentAt"`
	Delivery              *InvoiceDelivery `json:"delivery"`
	VoidedAt              *string          `json:"voidedAt"`
	VoidReason            *string          `json:"voidReason"`
	SupersedesInvoiceID   *string          `json:"supersedesInvoiceId"`
	SupersededByInvoiceID *string          `json:"supersededByInvoiceId"`
	CreatedAt             string           `json:"createdAt"`
	UpdatedAt             string           `json:"updatedAt"`
}

// InvoiceTotals: **Null for a draft** — null, not zero. A draft's figures are
// recomputed on read and the list does not recompute; fetch the invoice by id
// for a draft's current numbers.
//
// The API may send null in its place.
type InvoiceTotals struct {
	// Collected: Currency code → amount in the currency's major unit.
	Collected map[string]float64 `json:"collected"`
	// Adjustment: Currency code → amount in the currency's major unit.
	Adjustment map[string]float64 `json:"adjustment"`
	// Adjusted: Currency code → amount in the currency's major unit.
	Adjusted map[string]float64 `json:"adjusted"`
	// Billed: Keyed by the invoice currency, plus any currency that could not be
	// converted — which keeps its own key so the total is never quietly short.
	Billed map[string]float64 `json:"billed"`
}

// InvoiceUpdate is the `InvoiceUpdate` schema.
type InvoiceUpdate struct {
	PeriodFrom string  `json:"periodFrom"`
	PeriodTo   string  `json:"periodTo"`
	Notes      *string `json:"notes,omitempty"`
}

// InvoiceVoidRequest is the `InvoiceVoidRequest` schema.
type InvoiceVoidRequest struct {
	// Reason: Required. The only record of why a customer was sent an invoice
	// that was then withdrawn.
	Reason string `json:"reason"`
	// Supersede: Raise the corrective draft in the same call, linked both ways
	// to the original. Doing it in one call is what keeps the pair from being
	// left half-made by a failed second request.
	Supersede *bool `json:"supersede,omitempty"`
}

// InvoiceVoidResponse is the `InvoiceVoidResponse` schema.
type InvoiceVoidResponse struct {
	Invoice     Invoice                        `json:"invoice"`
	Replacement InvoiceVoidResponseReplacement `json:"replacement"`
}

// InvokeActionRequest is the `InvokeActionRequest` schema.
type InvokeActionRequest struct {
	PluginID         string      `json:"pluginId"`
	AccountID        string      `json:"accountId"`
	ResourceTypeID   string      `json:"resourceTypeId"`
	ResourceID       ResourceID  `json:"resourceId"`
	ActionID         string      `json:"actionId"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// JiraIntegration is the `JiraIntegration` schema.
//
// The API may send null in its place.
type JiraIntegration struct {
	SiteURL      string `json:"siteUrl"`
	AccountEmail string `json:"accountEmail"`
	// TokenHint: Redacted marker for the stored API token, e.g. `…a7f2`. The
	// token itself is never returned.
	TokenHint          string  `json:"tokenHint"`
	DefaultProjectKey  *string `json:"defaultProjectKey"`
	DefaultIssueTypeID *string `json:"defaultIssueTypeId"`
	UpdatedAt          string  `json:"updatedAt"`
}

// JiraIntegrationInput is the `JiraIntegrationInput` schema.
type JiraIntegrationInput struct {
	// SiteURL: Jira Cloud site address. Must resolve to a .atlassian.net (or
	// legacy .jira.com) host; a bare hostname and a pasted board or issue URL
	// are both accepted and normalized.
	SiteURL string `json:"siteUrl"`
	// AccountEmail: Atlassian account email — the username half of the
	// basic-auth pair.
	AccountEmail string `json:"accountEmail"`
	// APIToken: API token from id.atlassian.com. Omit to keep the stored token;
	// required on first connect.
	APIToken           *string `json:"apiToken,omitempty"`
	DefaultProjectKey  *string `json:"defaultProjectKey,omitempty"`
	DefaultIssueTypeID *string `json:"defaultIssueTypeId,omitempty"`
}

// JiraIssueLink is the `JiraIssueLink` schema.
type JiraIssueLink struct {
	ID              string         `json:"id"`
	SourceKind      JiraSourceKind `json:"sourceKind"`
	SourceID        string         `json:"sourceId"`
	IssueKey        string         `json:"issueKey"`
	IssueURL        string         `json:"issueUrl"`
	CreatedByUserID *string        `json:"createdByUserId"`
	CreatedAt       string         `json:"createdAt"`
}

// JiraIssueType is the `JiraIssueType` schema.
type JiraIssueType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Subtask: Always false — subtasks need a parent issue, so they are filtered
	// out.
	Subtask     bool    `json:"subtask"`
	Description *string `json:"description"`
}

// JiraProject is the `JiraProject` schema.
type JiraProject struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

// JiraSourceKind: Which detector produced the finding the issue was filed from.
type JiraSourceKind = string

// The values JiraSourceKind takes.
const (
	JiraSourceKindCostAnomaly    JiraSourceKind = "cost_anomaly"
	JiraSourceKindOrphan         JiraSourceKind = "orphan"
	JiraSourceKindOversized      JiraSourceKind = "oversized"
	JiraSourceKindPostureFinding JiraSourceKind = "posture_finding"
	JiraSourceKindExpiring       JiraSourceKind = "expiring"
	JiraSourceKindProbe          JiraSourceKind = "probe"
)

// JiraVerifyInput: Supply all three to test credentials that have not been saved
// yet; send an empty object to re-test the stored ones.
type JiraVerifyInput struct {
	SiteURL      *string `json:"siteUrl,omitempty"`
	AccountEmail *string `json:"accountEmail,omitempty"`
	APIToken     *string `json:"apiToken,omitempty"`
}

// JiraVerifyResult is the `JiraVerifyResult` schema.
type JiraVerifyResult struct {
	OK           bool    `json:"ok"`
	AccountID    string  `json:"accountId"`
	DisplayName  string  `json:"displayName"`
	EmailAddress *string `json:"emailAddress"`
}

// JSONObject: Free-form JSON object whose shape depends on the plugin.
//
// Spec schema: `JsonObject`.
type JSONObject = map[string]any

// KVCommandRequest is the `KvCommandRequest` schema.
//
// Spec schema: `KvCommandRequest`.
type KVCommandRequest struct {
	AccountID        string      `json:"accountId"`
	Command          string      `json:"command"`
	Args             []any       `json:"args"`
	PluginID         *string     `json:"pluginId,omitempty"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// KVCommandResponse is the `KvCommandResponse` schema.
//
// Spec schema: `KvCommandResponse`.
type KVCommandResponse struct {
	Result any `json:"result,omitempty"`
}

// LeaseConflict is the `LeaseConflict` schema.
type LeaseConflict struct {
	Error string `json:"error"`
}

// LinearIntegration is the `LinearIntegration` schema.
//
// The API may send null in its place.
type LinearIntegration struct {
	// KeyHint: Redacted marker for the stored personal API key, e.g. `…a7f2`.
	// The key itself is never returned.
	KeyHint string `json:"keyHint"`
	// DefaultTeamID: Team the file-issue window preselects. A Linear team id,
	// not a team key.
	DefaultTeamID *string `json:"defaultTeamId"`
	UpdatedAt     string  `json:"updatedAt"`
}

// LinearIntegrationInput is the `LinearIntegrationInput` schema.
type LinearIntegrationInput struct {
	// APIKey: Personal API key from Linear → Settings → Security & access. Omit
	// to keep the stored key; required on first connect.
	APIKey        *string `json:"apiKey,omitempty"`
	DefaultTeamID *string `json:"defaultTeamId,omitempty"`
}

// LinearIssueLink is the `LinearIssueLink` schema.
type LinearIssueLink struct {
	ID              string           `json:"id"`
	SourceKind      LinearSourceKind `json:"sourceKind"`
	SourceID        string           `json:"sourceId"`
	IssueIdentifier string           `json:"issueIdentifier"`
	IssueURL        string           `json:"issueUrl"`
	CreatedByUserID *string          `json:"createdByUserId"`
	CreatedAt       string           `json:"createdAt"`
}

// LinearSourceKind: Which detector produced the finding the issue was filed
// from.
type LinearSourceKind = string

// The values LinearSourceKind takes.
const (
	LinearSourceKindCostAnomaly    LinearSourceKind = "cost_anomaly"
	LinearSourceKindOrphan         LinearSourceKind = "orphan"
	LinearSourceKindOversized      LinearSourceKind = "oversized"
	LinearSourceKindPostureFinding LinearSourceKind = "posture_finding"
	LinearSourceKindExpiring       LinearSourceKind = "expiring"
	LinearSourceKindProbe          LinearSourceKind = "probe"
)

// LinearTeam is the `LinearTeam` schema.
type LinearTeam struct {
	// ID: Team id (UUID) — what issueCreate wants.
	ID string `json:"id"`
	// Key: Short prefix issue identifiers are built from.
	Key  string `json:"key"`
	Name string `json:"name"`
}

// LinearVerifyInput: Supply a key to test one that has not been saved yet; send
// an empty object to re-test the stored one.
type LinearVerifyInput struct {
	APIKey *string `json:"apiKey,omitempty"`
}

// LinearVerifyResult: The Linear user behind the API key, from the `viewer`
// query.
type LinearVerifyResult struct {
	OK    bool    `json:"ok"`
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email"`
}

// LinuxAppHostCheck is the `LinuxAppHostCheck` schema.
type LinuxAppHostCheck struct {
	Preflight LinuxAppHostPreflight `json:"preflight"`
	Plan      *LinuxAppInstallPlan  `json:"plan"`
}

// LinuxAppHostPreflight is the `LinuxAppHostPreflight` schema.
type LinuxAppHostPreflight struct {
	Arch   string `json:"arch"`
	OsID   string `json:"osId"`
	OsName string `json:"osName"`
	// PackageManager: One of "apt-get", "dnf", "yum", "apk", "pacman", "zypper".
	PackageManager *string `json:"packageManager"`
	// Privilege: One of "root", "sudo", "sudo-password", "none".
	Privilege    string                `json:"privilege"`
	Requirements []LinuxAppRequirement `json:"requirements"`
	// Staging: A writable, exec-capable directory was found to stage the app
	// server in. False means every candidate is missing, unwritable, or mounted
	// noexec — which no package fixes.
	Staging  bool  `json:"staging"`
	AppCount int64 `json:"appCount"`
	Ready    bool  `json:"ready"`
}

// LinuxAppHostTarget is the `LinuxAppHostTarget` schema.
type LinuxAppHostTarget struct {
	AccountID  string `json:"accountId"`
	ResourceID string `json:"resourceId"`
	SSHKeyID   string `json:"sshKeyId"`
	Host       string `json:"host"`
	Username   string `json:"username"`
	Port       *int64 `json:"port,omitempty"`
}

// LinuxAppInstallOutcome is the `LinuxAppInstallOutcome` schema.
type LinuxAppInstallOutcome struct {
	Log       []string              `json:"log"`
	Failed    []string              `json:"failed"`
	Preflight LinuxAppHostPreflight `json:"preflight"`
}

// LinuxAppInstallPlan is the `LinuxAppInstallPlan` schema.
//
// The API may send null in its place.
type LinuxAppInstallPlan struct {
	// PackageManager: One of "apt-get", "dnf", "yum", "apk", "pacman", "zypper".
	PackageManager string `json:"packageManager"`
	// Privilege: One of "root", "sudo", "sudo-password", "none".
	Privilege    string                  `json:"privilege"`
	Requirements []LinuxAppRequirementID `json:"requirements"`
	Packages     []string                `json:"packages"`
	// Commands: Exactly what would run on the host, privilege prefix included.
	Commands      []string `json:"commands"`
	CanInstall    bool     `json:"canInstall"`
	BlockedReason *string  `json:"blockedReason,omitempty"`
}

// LinuxAppRequirement is the `LinuxAppRequirement` schema.
type LinuxAppRequirement struct {
	ID LinuxAppRequirementID `json:"id"`
	// Severity: One of "required", "recommended".
	Severity string `json:"severity"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	OK       bool   `json:"ok"`
}

// LinuxAppRequirementID: gzip unpacks the uploaded app server; xkb is the
// keyboard layout data xkbcommon compiles a keymap from; dbus is the session bus
// GTK applications wait for; fonts, mesa and icons decide what an application
// then looks like.
//
// Spec schema: `LinuxAppRequirementId`.
type LinuxAppRequirementID = string

// The values LinuxAppRequirementID takes.
const (
	LinuxAppRequirementIDGzip  LinuxAppRequirementID = "gzip"
	LinuxAppRequirementIDXkb   LinuxAppRequirementID = "xkb"
	LinuxAppRequirementIDDbus  LinuxAppRequirementID = "dbus"
	LinuxAppRequirementIDFonts LinuxAppRequirementID = "fonts"
	LinuxAppRequirementIDMesa  LinuxAppRequirementID = "mesa"
	LinuxAppRequirementIDIcons LinuxAppRequirementID = "icons"
)

// LinuxAppSetupEvent is the `LinuxAppSetupEvent` schema.
type LinuxAppSetupEvent struct {
	Line    *string                 `json:"line,omitempty"`
	Outcome *LinuxAppInstallOutcome `json:"outcome,omitempty"`
	Error   *string                 `json:"error,omitempty"`
}

// LinuxAppSetupRequest is the `LinuxAppSetupRequest` schema.
type LinuxAppSetupRequest struct {
	AccountID    string                  `json:"accountId"`
	ResourceID   string                  `json:"resourceId"`
	SSHKeyID     string                  `json:"sshKeyId"`
	Host         string                  `json:"host"`
	Username     string                  `json:"username"`
	Port         *int64                  `json:"port,omitempty"`
	Requirements []LinuxAppRequirementID `json:"requirements,omitempty"`
}

// LiteralAssociationRequest is the `LiteralAssociationRequest` schema.
type LiteralAssociationRequest struct {
	ResourceID     ResourceID `json:"resourceId"`
	FieldKey       string     `json:"fieldKey"`
	PlaintextValue string     `json:"plaintextValue"`
}

// LogCapableResource is the `LogCapableResource` schema.
type LogCapableResource struct {
	ResourceID     string   `json:"resourceId"`
	AccountID      string   `json:"accountId"`
	AccountName    string   `json:"accountName"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	DisplayName    string   `json:"displayName"`
	// ParentResourceID: Set for sidecar streams: the stored parent resource the
	// peer client is built through.
	ParentResourceID  *string `json:"parentResourceId,omitempty"`
	ParentDisplayName *string `json:"parentDisplayName,omitempty"`
}

// LogCapableResourceList is the `LogCapableResourceList` schema.
type LogCapableResourceList struct {
	Resources []LogCapableResource `json:"resources"`
}

// LogStreamSelector is the `LogStreamSelector` schema.
type LogStreamSelector struct {
	// ResourceID: Infrawrench resource id of the stream to tail — or, for a
	// sidecar stream, the peer plugin's own resource id (not a stored row).
	ResourceID     string   `json:"resourceId"`
	AccountID      string   `json:"accountId"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	// ParentResourceID: Set for sidecar streams (e.g. a pod inside a managed
	// cluster): the stored parent resource whose outputs mint the peer plugin's
	// credentials. The logs endpoint routes through the peer client when
	// present.
	ParentResourceID *string `json:"parentResourceId,omitempty"`
	// Container: Container to fetch when the resource has more than one; omit
	// for the default.
	Container *string `json:"container,omitempty"`
}

// LogWorkspaceQuery is the `LogWorkspaceQuery` schema.
type LogWorkspaceQuery struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	Resources []LogStreamSelector `json:"resources"`
	// Search: The search expression. Empty matches every line; `/pattern/`
	// (optionally `/pattern/i`) is a regular expression; otherwise
	// whitespace-separated terms that must ALL appear in a line
	// (case-insensitive), with `"quoted phrases"` and `-term` negation.
	Search string `json:"search"`
	// AlertEnabled: When true the poller periodically evaluates the query and
	// alerts on match.
	AlertEnabled bool `json:"alertEnabled"`
	// LastEvalAt: Last time the alert pass evaluated this query; null until it
	// has run.
	LastEvalAt *string `json:"lastEvalAt"`
	// LastMatchAt: Last evaluation that found at least one matching line.
	LastMatchAt *string `json:"lastMatchAt"`
	// LastAlertedAt: Last dispatched notification — the cooldown anchor.
	LastAlertedAt *string `json:"lastAlertedAt"`
	// LastEvalError: Failure detail from the last evaluation.
	LastEvalError *string `json:"lastEvalError"`
	// LastMatchSample: Truncated sample of the most recent matching line.
	LastMatchSample *string `json:"lastMatchSample"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// LogWorkspaceQueryConflict is the `LogWorkspaceQueryConflict` schema.
type LogWorkspaceQueryConflict struct {
	Error string `json:"error"`
}

// LogWorkspaceQueryCreate is the `LogWorkspaceQueryCreate` schema.
type LogWorkspaceQueryCreate struct {
	Name      string              `json:"name"`
	Resources []LogStreamSelector `json:"resources"`
	// Search: The search expression. Empty matches every line; `/pattern/`
	// (optionally `/pattern/i`) is a regular expression; otherwise
	// whitespace-separated terms that must ALL appear in a line
	// (case-insensitive), with `"quoted phrases"` and `-term` negation.
	Search       string `json:"search"`
	AlertEnabled *bool  `json:"alertEnabled,omitempty"`
}

// LogWorkspaceQueryList is the `LogWorkspaceQueryList` schema.
type LogWorkspaceQueryList struct {
	Queries []LogWorkspaceQuery `json:"queries"`
}

// LogWorkspaceQueryUpdate is the `LogWorkspaceQueryUpdate` schema.
type LogWorkspaceQueryUpdate struct {
	Name      *string             `json:"name,omitempty"`
	Resources []LogStreamSelector `json:"resources,omitempty"`
	// Search: The search expression. Empty matches every line; `/pattern/`
	// (optionally `/pattern/i`) is a regular expression; otherwise
	// whitespace-separated terms that must ALL appear in a line
	// (case-insensitive), with `"quoted phrases"` and `-term` negation.
	Search       *string `json:"search,omitempty"`
	AlertEnabled *bool   `json:"alertEnabled,omitempty"`
}

// LogWorkspaceQueryUpdateConflict is the `LogWorkspaceQueryUpdateConflict`
// schema.
type LogWorkspaceQueryUpdateConflict struct {
	Error string `json:"error"`
}

// LogsRequest is the `LogsRequest` schema.
type LogsRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceID       ResourceID  `json:"resourceId"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
	TailLines        *int64      `json:"tailLines,omitempty"`
	Container        *string     `json:"container,omitempty"`
	Previous         *bool       `json:"previous,omitempty"`
}

// LogsResponse is the `LogsResponse` schema.
type LogsResponse struct {
	// Text: Raw log text; each entry keeps its trailing newline.
	Text string `json:"text"`
	// Containers: Container names available for this resource — drives the
	// container picker.
	Containers []string `json:"containers"`
	// ActiveContainer: Container `text` was read from.
	ActiveContainer string `json:"activeContainer"`
}

// ManagedAccount: A customer a managed service provider bills. A cost centre or
// cloud account belongs to at most one managed account — billing the same money
// to two customers is refused at write time with a 409 naming the other
// customer.
type ManagedAccount struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	ContactName     *string `json:"contactName"`
	ContactEmail    *string `json:"contactEmail"`
	BillingAddress  *string `json:"billingAddress"`
	BillingCurrency string  `json:"billingCurrency"`
	// CostBasis: One of "cash", "amortized".
	CostBasis         string   `json:"costBasis"`
	ApplyBillingRules bool     `json:"applyBillingRules"`
	Notes             *string  `json:"notes"`
	CostCentreIDs     []string `json:"costCentreIds"`
	AccountIDs        []string `json:"accountIds"`
	InvoiceCount      int64    `json:"invoiceCount"`
	CreatedByUserID   *string  `json:"createdByUserId"`
	CreatedAt         string   `json:"createdAt"`
	UpdatedAt         string   `json:"updatedAt"`
}

// ManagedAccountInput is the `ManagedAccountInput` schema.
type ManagedAccountInput struct {
	Name           string  `json:"name"`
	ContactName    *string `json:"contactName,omitempty"`
	ContactEmail   *string `json:"contactEmail,omitempty"`
	BillingAddress *string `json:"billingAddress,omitempty"`
	// BillingCurrency: ISO 4217 code the customer is invoiced in. Spend
	// collected in another currency is converted through the organisation's own
	// stated exchange rates, and the rate used is frozen onto every invoice — so
	// restating a rate later cannot restate history.
	BillingCurrency string `json:"billingCurrency"`
	// CostBasis: Defaults to `amortized`. Charging a customer the whole cash
	// value of a three-year commitment in the month it was signed is not a bill
	// anyone can budget against.
	//
	// One of "cash", "amortized".
	CostBasis *string `json:"costBasis,omitempty"`
	// ApplyBillingRules: Defaults to true. False is a pass-through contract: the
	// customer is billed exactly what the providers charged, with no markup,
	// discount or fixed fee applied.
	ApplyBillingRules *bool   `json:"applyBillingRules,omitempty"`
	Notes             *string `json:"notes,omitempty"`
	// CostCentreIDs: Cost centres whose spend belongs to this customer.
	// **Subtrees are included** — naming a parent bills every descendant, and
	// naming both a parent and its child bills the child once, not twice.
	//
	// This is deliberately a list of existing cost centres rather than a rule of
	// its own. Which spend lands in which centre is already decided by the
	// organisation's allocation rules, and a second vocabulary over the same
	// data would eventually disagree with the first — at which point an invoice
	// would stop matching the showback report the customer was shown.
	CostCentreIDs []string `json:"costCentreIds,omitempty"`
	// AccountIDs: Cloud accounts whose spend belongs to this customer. Evaluated
	// **after** every allocation rule, so an account in scope claims only the
	// spend no cost centre already claimed. Every cost row therefore resolves
	// exactly once: nothing is billed twice and nothing goes missing.
	AccountIDs []string `json:"accountIds,omitempty"`
}

// Manifest is the `Manifest` schema.
type Manifest struct {
	Manifest string `json:"manifest"`
}

// MeResponse is the `MeResponse` schema.
type MeResponse struct {
	UserID      string       `json:"userId"`
	Email       string       `json:"email"`
	Role        *RoleSummary `json:"role"`
	Permissions []Permission `json:"permissions"`
}

// MetricAlertEvent is the `MetricAlertEvent` schema.
type MetricAlertEvent struct {
	ID           string `json:"id"`
	RuleID       string `json:"ruleId"`
	RuleName     string `json:"ruleName"`
	ResourceID   string `json:"resourceId"`
	ResourceName string `json:"resourceName"`
	// Status: One of "firing", "resolved".
	Status string `json:"status"`
	// ObservedValue: Worst sample observed in the breaching window, in the
	// metric's unit.
	ObservedValue float64 `json:"observedValue"`
	FiredAt       string  `json:"firedAt"`
	ResolvedAt    *string `json:"resolvedAt"`
}

// MetricAlertRule is the `MetricAlertRule` schema.
type MetricAlertRule struct {
	Name string `json:"name"`
	// PluginID: Selector: plugin the resource must belong to. Null matches any
	// plugin.
	PluginID *string `json:"pluginId"`
	// ResourceTypeID: Selector: resource type within the plugin. Null matches
	// any type.
	ResourceTypeID *string `json:"resourceTypeId"`
	// TagKey: Selector: tag key the resource must carry (matched
	// case-insensitively). Null applies no tag filter. Resources are always
	// selected by this query, never by id, so rules cover resources created
	// later.
	TagKey *string `json:"tagKey"`
	// TagValue: Selector: exact value tagKey must have. Null matches any value.
	TagValue *string `json:"tagValue"`
	// MetricKey: The metric series label as the resource's charts report it (see
	// /metric-alerts/metric-keys).
	MetricKey string `json:"metricKey"`
	// Comparator: One of ">", ">=", "<", "<=".
	Comparator string  `json:"comparator"`
	Threshold  float64 `json:"threshold"`
	// ForMinutes: Trailing window (minutes) the condition must hold for before
	// firing.
	ForMinutes int64 `json:"forMinutes"`
	// CooldownMinutes: Least minutes between notified firings for one (rule,
	// resource).
	CooldownMinutes int64   `json:"cooldownMinutes"`
	Enabled         bool    `json:"enabled"`
	ID              string  `json:"id"`
	LastEvalAt      *string `json:"lastEvalAt"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// MetricAlertRuleInput is the `MetricAlertRuleInput` schema.
type MetricAlertRuleInput struct {
	Name string `json:"name"`
	// PluginID: Selector: plugin the resource must belong to. Null matches any
	// plugin.
	PluginID *string `json:"pluginId"`
	// ResourceTypeID: Selector: resource type within the plugin. Null matches
	// any type.
	ResourceTypeID *string `json:"resourceTypeId"`
	// TagKey: Selector: tag key the resource must carry (matched
	// case-insensitively). Null applies no tag filter. Resources are always
	// selected by this query, never by id, so rules cover resources created
	// later.
	TagKey *string `json:"tagKey"`
	// TagValue: Selector: exact value tagKey must have. Null matches any value.
	TagValue *string `json:"tagValue"`
	// MetricKey: The metric series label as the resource's charts report it (see
	// /metric-alerts/metric-keys).
	MetricKey string `json:"metricKey"`
	// Comparator: One of ">", ">=", "<", "<=".
	Comparator string  `json:"comparator"`
	Threshold  float64 `json:"threshold"`
	// ForMinutes: Trailing window (minutes) the condition must hold for before
	// firing.
	ForMinutes int64 `json:"forMinutes"`
	// CooldownMinutes: Least minutes between notified firings for one (rule,
	// resource).
	CooldownMinutes int64 `json:"cooldownMinutes"`
	Enabled         bool  `json:"enabled"`
}

// MetricAlertRuleWithStatus is the `MetricAlertRuleWithStatus` schema.
type MetricAlertRuleWithStatus struct {
	Name string `json:"name"`
	// PluginID: Selector: plugin the resource must belong to. Null matches any
	// plugin.
	PluginID *string `json:"pluginId"`
	// ResourceTypeID: Selector: resource type within the plugin. Null matches
	// any type.
	ResourceTypeID *string `json:"resourceTypeId"`
	// TagKey: Selector: tag key the resource must carry (matched
	// case-insensitively). Null applies no tag filter. Resources are always
	// selected by this query, never by id, so rules cover resources created
	// later.
	TagKey *string `json:"tagKey"`
	// TagValue: Selector: exact value tagKey must have. Null matches any value.
	TagValue *string `json:"tagValue"`
	// MetricKey: The metric series label as the resource's charts report it (see
	// /metric-alerts/metric-keys).
	MetricKey string `json:"metricKey"`
	// Comparator: One of ">", ">=", "<", "<=".
	Comparator string  `json:"comparator"`
	Threshold  float64 `json:"threshold"`
	// ForMinutes: Trailing window (minutes) the condition must hold for before
	// firing.
	ForMinutes int64 `json:"forMinutes"`
	// CooldownMinutes: Least minutes between notified firings for one (rule,
	// resource).
	CooldownMinutes int64   `json:"cooldownMinutes"`
	Enabled         bool    `json:"enabled"`
	ID              string  `json:"id"`
	LastEvalAt      *string `json:"lastEvalAt"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
	// FiringCount: Resources currently in breach of this rule.
	FiringCount int64 `json:"firingCount"`
	// MatchingResourceCount: Resources the selector matches right now.
	MatchingResourceCount int64 `json:"matchingResourceCount"`
}

// MetricAlertSelectorOptions is the `MetricAlertSelectorOptions` schema.
type MetricAlertSelectorOptions struct {
	Plugins []MetricAlertSelectorOptionsPlugins `json:"plugins"`
	TagKeys []string                            `json:"tagKeys"`
}

// MetricAlertSelectorPreview is the `MetricAlertSelectorPreview` schema.
type MetricAlertSelectorPreview struct {
	MatchingResourceCount int64 `json:"matchingResourceCount"`
	// SampleResourceNames: Up to 10 matching display names, for a live form
	// preview.
	SampleResourceNames []string `json:"sampleResourceNames"`
}

// MetricSeries is the `MetricSeries` schema.
type MetricSeries struct {
	Label  string               `json:"label"`
	Unit   *string              `json:"unit,omitempty"`
	Points []MetricSeriesPoints `json:"points"`
}

// MetricSeriesKey is the `MetricSeriesKey` schema.
type MetricSeriesKey struct {
	Label string `json:"label"`
	Unit  string `json:"unit"`
	// ResourceCount: Distinct resources that reported this series in the last 7
	// days.
	ResourceCount int64 `json:"resourceCount"`
}

// MetricsRequest is the `MetricsRequest` schema.
type MetricsRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceID       ResourceID  `json:"resourceId"`
	StartMs          *int64      `json:"startMs,omitempty"`
	EndMs            *int64      `json:"endMs,omitempty"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// MetricsResponse is the `MetricsResponse` schema.
type MetricsResponse struct {
	Series []MetricSeries `json:"series"`
}

// MomentEvent is the `MomentEvent` schema.
type MomentEvent struct {
	// ID: Stable synthetic id, unique within a response (`feed:rowId[:phase]`).
	ID   string       `json:"id"`
	Feed MomentFeedID `json:"feed"`
	// Kind: Fine-grained `<noun>.<verb>` kind, e.g. `change.created`,
	// `incident.started`, `workflow-run.failed`, `deployment.finished`,
	// `freeze.started`, `drift-alert.sent`. Open set — render unknown kinds
	// generically.
	Kind      string `json:"kind"`
	Timestamp string `json:"timestamp"`
	// Title: One-line headline.
	Title string `json:"title"`
	// Detail: Optional second line — diff summary, actor, error text.
	Detail         *string          `json:"detail,omitempty"`
	Severity       MomentSeverity   `json:"severity"`
	PluginID       *string          `json:"pluginId,omitempty"`
	AccountID      *string          `json:"accountId,omitempty"`
	AccountName    *string          `json:"accountName,omitempty"`
	ResourceID     *string          `json:"resourceId,omitempty"`
	ResourceTypeID *string          `json:"resourceTypeId,omitempty"`
	ResourceName   *string          `json:"resourceName,omitempty"`
	Link           *MomentEventLink `json:"link,omitempty"`
}

// MomentEventLink is the `MomentEventLink` schema.
//
// The API may send null in its place.
type MomentEventLink struct {
	// Kind: Which native screen the event deep-links to.
	//
	// One of "resource", "changes", "incident", "costs", "workflow-run",
	// "deployment", "audit", "freeze", "expiring".
	Kind string `json:"kind"`
	// ID: Target id where the kind needs one (resource id, run id, freeze id…).
	ID *string `json:"id,omitempty"`
	// ParentID: Parent id where the target needs one (workflow id for a run).
	ParentID *string `json:"parentId,omitempty"`
	// URL: Absolute external URL — a provider's incident page. Wins when
	// present.
	URL *string `json:"url,omitempty"`
}

// MomentFeedID: One of the indexed feeds the moment union draws from.
//
// Spec schema: `MomentFeedId`.
type MomentFeedID = string

// The values MomentFeedID takes.
const (
	MomentFeedIDChanges         MomentFeedID = "changes"
	MomentFeedIDStatusIncidents MomentFeedID = "statusIncidents"
	MomentFeedIDCostAnomalies   MomentFeedID = "costAnomalies"
	MomentFeedIDWorkflowRuns    MomentFeedID = "workflowRuns"
	MomentFeedIDDeployments     MomentFeedID = "deployments"
	MomentFeedIDAudit           MomentFeedID = "audit"
	MomentFeedIDFreezes         MomentFeedID = "freezes"
	MomentFeedIDDriftAlerts     MomentFeedID = "driftAlerts"
	MomentFeedIDExpiryAlerts    MomentFeedID = "expiryAlerts"
)

// MomentFeedStatus is the `MomentFeedStatus` schema.
type MomentFeedStatus struct {
	Feed MomentFeedID `json:"feed"`
	// Status: `omitted` = the caller lacks the feed's read permission; `error` =
	// the feed's query failed but the rest of the response is still valid
	// (partial-failure tolerance).
	//
	// One of "ok", "omitted", "error".
	Status string `json:"status"`
	// Error: Short failure reason when `status` is `error`.
	Error *string `json:"error,omitempty"`
	// Truncated: True when the feed hit its row cap and events were dropped.
	Truncated *bool `json:"truncated,omitempty"`
}

// MomentIncidentSpan: A provider incident whose span overlaps the window —
// returned alongside the events so clients can badge events that fall inside it
// ("during DigitalOcean incident").
type MomentIncidentSpan struct {
	ID         string `json:"id"`
	PluginID   string `json:"pluginId"`
	PluginName string `json:"pluginName"`
	Title      string `json:"title"`
	// Impact: One of "maintenance", "minor", "major", "critical".
	Impact     string  `json:"impact"`
	StartedAt  string  `json:"startedAt"`
	ResolvedAt *string `json:"resolvedAt,omitempty"`
	URL        *string `json:"url,omitempty"`
}

// MomentResponse is the `MomentResponse` schema.
type MomentResponse struct {
	// At: The centre timestamp, normalized to ISO.
	At   string `json:"at"`
	From string `json:"from"`
	To   string `json:"to"`
	// WindowMinutes: The half-window actually applied, after clamping to 1–4320
	// minutes.
	WindowMinutes int64  `json:"windowMinutes"`
	GeneratedAt   string `json:"generatedAt"`
	// Feeds: One entry per feed, in canonical order — including omitted and
	// errored feeds.
	Feeds []MomentFeedStatus `json:"feeds"`
	// Events: Chronological, oldest first.
	Events    []MomentEvent        `json:"events"`
	Incidents []MomentIncidentSpan `json:"incidents"`
}

// MomentSeverity is the `MomentSeverity` schema.
type MomentSeverity = string

// The values MomentSeverity takes.
const (
	MomentSeverityInfo     MomentSeverity = "info"
	MomentSeverityWarning  MomentSeverity = "warning"
	MomentSeverityCritical MomentSeverity = "critical"
)

// MsTeamsStatus is the `MsTeamsStatus` schema.
type MsTeamsStatus struct {
	Webhooks []MsTeamsWebhook `json:"webhooks"`
}

// MsTeamsWebhook is the `MsTeamsWebhook` schema.
type MsTeamsWebhook struct {
	ID string `json:"id"`
	// Label: Display name for the channel, e.g. #alerts
	Label string `json:"label"`
	// URLHint: Non-secret hint at the stored webhook URL (host and last four
	// characters). The URL itself is never returned.
	URLHint string `json:"urlHint"`
}

// MsTeamsWebhookCreate is the `MsTeamsWebhookCreate` schema.
type MsTeamsWebhookCreate struct {
	Label string `json:"label"`
	// URL: The webhook URL from a Teams 'Workflows' automation. Must be https
	// and on a Microsoft-operated host (*.api.powerautomate.com,
	// *.api.powerplatform.com, *.logic.azure.com, *.flow.microsoft.com, or a
	// legacy *.webhook.office.com connector).
	URL string `json:"url"`
}

// MsTeamsWebhookUpdate is the `MsTeamsWebhookUpdate` schema.
type MsTeamsWebhookUpdate struct {
	Label string `json:"label"`
}

// NetworkFlowAccountStatus is the `NetworkFlowAccountStatus` schema.
type NetworkFlowAccountStatus struct {
	AccountID   string `json:"accountId"`
	PluginID    string `json:"pluginId"`
	DisplayName string `json:"displayName"`
	// SupportsFlows: False when the account's provider has no flow source we can
	// read. Such accounts are listed and excluded from the totals rather than
	// contributing zero bytes — zero would be a claim about their network, this
	// is a statement about our coverage.
	SupportsFlows    bool                `json:"supportsFlows"`
	CollectedThrough *string             `json:"collectedThrough"`
	LastPolledAt     *string             `json:"lastPolledAt"`
	FailureCount     int64               `json:"failureCount"`
	LastError        *string             `json:"lastError"`
	LastErrorHelpURL *string             `json:"lastErrorHelpUrl"`
	Sources          []NetworkFlowSource `json:"sources"`
	// LastQueryBytesScanned: Log data the provider billed this account for the
	// last collection's queries.
	LastQueryBytesScanned *float64 `json:"lastQueryBytesScanned"`
}

// NetworkFlowEndpoint is the `NetworkFlowEndpoint` schema.
type NetworkFlowEndpoint struct {
	// Ref: Stable endpoint identity — a provider resource id where one could be
	// resolved, otherwise a class token (`internet`, `aws:s3`,
	// `infrawrench:unattributed`). Never a raw IP address: addresses churn, so
	// the same workload would be a different row every day.
	Ref     string `json:"ref"`
	Label   string `json:"label"`
	Zone    string `json:"zone"`
	Region  string `json:"region"`
	Service string `json:"service"`
	// ResourceTypeID: Set when `ref` is a resource this organization syncs, so
	// the row can link out.
	ResourceTypeID string `json:"resourceTypeId"`
}

// NetworkFlowFeed is the `NetworkFlowFeed` schema.
type NetworkFlowFeed struct {
	Enabled             bool  `json:"enabled"`
	InitialLookbackDays int64 `json:"initialLookbackDays"`
	// Estimated: Always true. Flow bytes come from logs that sample or drop
	// under load and are priced at published list rates with no free tier, no
	// volume tier and no negotiated discount modelled — the ranking is sound,
	// the absolute figure will not reconcile to the invoice.
	Estimated bool                       `json:"estimated"`
	Range     NetworkFlowFeedRange       `json:"range"`
	Scopes    []NetworkFlowScopeSummary  `json:"scopes"`
	TopFlows  []NetworkFlowPair          `json:"topFlows"`
	Accounts  []NetworkFlowAccountStatus `json:"accounts"`
	RateCards []NetworkFlowRateCard      `json:"rateCards"`
	Totals    NetworkFlowFeedTotals      `json:"totals"`
}

// NetworkFlowPair is the `NetworkFlowPair` schema.
type NetworkFlowPair struct {
	Source      NetworkFlowEndpoint `json:"source"`
	Destination NetworkFlowEndpoint `json:"destination"`
	// Scope: Which billing boundary the traffic crossed. `unknown` means the
	// provider's record did not determine one — it is priced at zero and
	// labelled rather than folded into a neighbouring boundary.
	//
	// One of "intra_zone", "cross_zone", "cross_region", "internet_egress",
	// "internet_ingress", "provider_service", "nat_gateway",
	// "private_interconnect", "unknown".
	Scope string `json:"scope"`
	// Direction: One of "egress", "ingress".
	Direction string `json:"direction"`
	// Attribution: One of "resolved", "unattributed".
	Attribution   string  `json:"attribution"`
	Bytes         float64 `json:"bytes"`
	Packets       float64 `json:"packets"`
	EstimatedCost float64 `json:"estimatedCost"`
	Currency      string  `json:"currency"`
	AccountID     string  `json:"accountId"`
	PluginID      string  `json:"pluginId"`
	// Days: Days in the range this pair appeared on.
	Days int64 `json:"days"`
}

// NetworkFlowRateCard is the `NetworkFlowRateCard` schema.
type NetworkFlowRateCard struct {
	PluginID string `json:"pluginId"`
	Currency string `json:"currency"`
	// AsOf: Date the rates were last checked against the provider's pricing
	// page.
	AsOf  string             `json:"asOf"`
	PerGb map[string]float64 `json:"perGb"`
	// QueriesBillable: True when collecting flows runs queries the provider
	// bills to your cloud account.
	QueriesBillable bool `json:"queriesBillable"`
	// Sampled: True when the flow source samples rather than recording all
	// flows.
	Sampled bool `json:"sampled"`
}

// NetworkFlowScopeSummary is the `NetworkFlowScopeSummary` schema.
type NetworkFlowScopeSummary struct {
	// Scope: Which billing boundary the traffic crossed. `unknown` means the
	// provider's record did not determine one — it is priced at zero and
	// labelled rather than folded into a neighbouring boundary.
	//
	// One of "intra_zone", "cross_zone", "cross_region", "internet_egress",
	// "internet_ingress", "provider_service", "nat_gateway",
	// "private_interconnect", "unknown".
	Scope string `json:"scope"`
	// Direction: One of "egress", "ingress".
	Direction     string  `json:"direction"`
	Bytes         float64 `json:"bytes"`
	EstimatedCost float64 `json:"estimatedCost"`
	Currency      string  `json:"currency"`
	CrossedZone   bool    `json:"crossedZone"`
	CrossedRegion bool    `json:"crossedRegion"`
	LeftCloud     bool    `json:"leftCloud"`
	// UnattributedBytes: Bytes inside `bytes` whose endpoints could not be tied
	// to a workload. A subset, not an addition — nothing here has been
	// apportioned across the attributed rows.
	UnattributedBytes float64 `json:"unattributedBytes"`
	// TruncatedBytes: Bytes inside `bytes` that fell below the stored top-N pair
	// cap, computed by subtraction against the provider's exact totals rather
	// than estimated.
	TruncatedBytes float64 `json:"truncatedBytes"`
}

// NetworkFlowSettings is the `NetworkFlowSettings` schema.
type NetworkFlowSettings struct {
	Enabled             bool  `json:"enabled"`
	InitialLookbackDays int64 `json:"initialLookbackDays"`
}

// NetworkFlowSource is the `NetworkFlowSource` schema.
type NetworkFlowSource struct {
	ID string `json:"id"`
	// Target: What the flow log is attached to — a VPC id, a network.
	Target          string  `json:"target"`
	Region          *string `json:"region"`
	DestinationType string  `json:"destinationType"`
	Usable          bool    `json:"usable"`
	// UnusableReason: Why the source cannot be read, in terms that name the fix.
	UnusableReason *string `json:"unusableReason"`
	HelpURL        *string `json:"helpUrl"`
}

// NoSQLCommandRequest is the `NoSqlCommandRequest` schema.
//
// Spec schema: `NoSqlCommandRequest`.
type NoSQLCommandRequest struct {
	PluginID         string      `json:"pluginId"`
	AccountID        string      `json:"accountId"`
	ResourceTypeID   string      `json:"resourceTypeId"`
	ResourceID       ResourceID  `json:"resourceId"`
	Command          string      `json:"command"`
	Args             []any       `json:"args"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// OK is the `Ok` schema.
//
// Spec schema: `Ok`.
type OK struct {
	OK bool `json:"ok"`
}

// OnCallNowEntry is the `OnCallNowEntry` schema.
type OnCallNowEntry struct {
	ScheduleID   string             `json:"scheduleId"`
	ScheduleName string             `json:"scheduleName"`
	Enabled      bool               `json:"enabled"`
	Shift        *OnCallShift       `json:"shift"`
	Next         *OnCallParticipant `json:"next"`
}

// OnCallNowResponse is the `OnCallNowResponse` schema.
type OnCallNowResponse struct {
	OnCall      []OnCallNowEntry `json:"onCall"`
	GeneratedAt string           `json:"generatedAt"`
}

// OnCallOverride is the `OnCallOverride` schema.
type OnCallOverride struct {
	ID              string  `json:"id"`
	ScheduleID      string  `json:"scheduleId"`
	UserID          string  `json:"userId"`
	UserName        *string `json:"userName"`
	StartsAt        string  `json:"startsAt"`
	EndsAt          string  `json:"endsAt"`
	Reason          *string `json:"reason"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedAt       string  `json:"createdAt"`
}

// OnCallOverrideCreate is the `OnCallOverrideCreate` schema.
type OnCallOverrideCreate struct {
	ScheduleID string  `json:"scheduleId"`
	UserID     string  `json:"userId"`
	StartsAt   string  `json:"startsAt"`
	EndsAt     string  `json:"endsAt"`
	Reason     *string `json:"reason,omitempty"`
}

// OnCallParticipant: The next person in the rotation — where an escalation goes.
// Resolved from the rotation and never from a cover: a cover is somebody
// standing in for one shift.
//
// The API may send null in its place.
type OnCallParticipant struct {
	UserID string  `json:"userId"`
	Name   *string `json:"name"`
	Email  *string `json:"email"`
}

// OnCallSchedule is the `OnCallSchedule` schema.
type OnCallSchedule struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
	// RotationDays: Days per shift. 7 is the common case; 1 gives a daily
	// rotation.
	RotationDays int64 `json:"rotationDays"`
	// HandoffTime: Wall-clock time in `timezone` at which the shift changes
	// hands.
	HandoffTime string `json:"handoffTime"`
	// StartDate: The calendar date in `timezone` the first shift begins on.
	// Every later boundary is derived from it, so moving this re-anchors the
	// whole rotation.
	StartDate string `json:"startDate"`
	// Participants: Rotation order. Reordering re-plans the future,
	// deliberately.
	Participants []*OnCallParticipant `json:"participants"`
	// Enabled: Off resolves to nobody. A routing destination pointing at a
	// disabled rotation contributes nobody and the rule's other destinations
	// still deliver.
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// OnCallScheduleCreate is the `OnCallScheduleCreate` schema.
type OnCallScheduleCreate struct {
	Name               string   `json:"name"`
	Timezone           string   `json:"timezone"`
	RotationDays       int64    `json:"rotationDays"`
	HandoffTime        string   `json:"handoffTime"`
	StartDate          string   `json:"startDate"`
	ParticipantUserIDs []string `json:"participantUserIds"`
	Enabled            *bool    `json:"enabled,omitempty"`
}

// OnCallScheduleList is the `OnCallScheduleList` schema.
type OnCallScheduleList struct {
	Schedules []OnCallSchedule `json:"schedules"`
}

// OnCallScheduleUpdate is the `OnCallScheduleUpdate` schema.
type OnCallScheduleUpdate struct {
	Name               *string  `json:"name,omitempty"`
	Timezone           *string  `json:"timezone,omitempty"`
	RotationDays       *int64   `json:"rotationDays,omitempty"`
	HandoffTime        *string  `json:"handoffTime,omitempty"`
	StartDate          *string  `json:"startDate,omitempty"`
	ParticipantUserIDs []string `json:"participantUserIds,omitempty"`
	Enabled            *bool    `json:"enabled,omitempty"`
}

// OnCallShift is the `OnCallShift` schema.
//
// The API may send null in its place.
type OnCallShift struct {
	StartsAt string  `json:"startsAt"`
	EndsAt   string  `json:"endsAt"`
	UserID   string  `json:"userId"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	// Source: One of "rotation", "override".
	Source        string `json:"source"`
	RotationIndex *int64 `json:"rotationIndex"`
}

// OnCallShiftsResponse is the `OnCallShiftsResponse` schema.
type OnCallShiftsResponse struct {
	Shifts []*OnCallShift `json:"shifts"`
	// Overrides: Covers overlapping the previewed window, returned
	// **separately** rather than merged into the shifts: a preview that folded
	// them in would make it impossible to see what the rotation itself does,
	// which is the thing being edited.
	Overrides []OnCallOverride `json:"overrides"`
}

// OrgConfigAlertSettings: Org-wide notification tuning. Cooldown claims
// (`lastNotifiedAt`, `lastSentWeekStart`) are deliberately absent: they are
// poller state, and resetting one from an apply would re-open a quiet period and
// page people twice.
type OrgConfigAlertSettings struct {
	CostAnomaly *OrgConfigAlertSettingsCostAnomaly `json:"costAnomaly,omitempty"`
	Drift       *OrgConfigAlertSettingsDrift       `json:"drift,omitempty"`
	Expiry      *OrgConfigAlertSettingsExpiry      `json:"expiry,omitempty"`
	Posture     *OrgConfigAlertSettingsPosture     `json:"posture,omitempty"`
	Digest      *OrgConfigAlertSettingsDigest      `json:"digest,omitempty"`
}

// OrgConfigApplyResult is the `OrgConfigApplyResult` schema.
type OrgConfigApplyResult struct {
	// Mode: One of "merge", "replace".
	Mode       string                     `json:"mode"`
	Changes    []OrgConfigChange          `json:"changes"`
	Unresolved []OrgConfigUnresolved      `json:"unresolved"`
	Counts     OrgConfigApplyResultCounts `json:"counts"`
	Applied    bool                       `json:"applied"`
}

// OrgConfigBudget is the `OrgConfigBudget` schema.
type OrgConfigBudget struct {
	// Key: Stable slug identifying this entity across organizations. Derived
	// from the name on export; it is what an apply matches on, so renaming an
	// entity while keeping its key is a rename rather than a delete-and-create.
	Key         string                      `json:"key"`
	Name        string                      `json:"name"`
	AmountCents int64                       `json:"amountCents"`
	Currency    *string                     `json:"currency,omitempty"`
	Filters     []OrgConfigCostFilter       `json:"filters,omitempty"`
	Thresholds  []OrgConfigBudgetThresholds `json:"thresholds"`
}

// OrgConfigChange is the `OrgConfigChange` schema.
type OrgConfigChange struct {
	Section OrgConfigSection `json:"section"`
	Key     string           `json:"key"`
	Name    string           `json:"name"`
	// Action: One of "create", "update", "delete", "unchanged".
	Action string `json:"action"`
	// Fields: Fields that differ, on an update.
	Fields []string `json:"fields,omitempty"`
}

// OrgConfigCostCentre is the `OrgConfigCostCentre` schema.
type OrgConfigCostCentre struct {
	// Key: Stable slug identifying this entity across organizations. Derived
	// from the name on export; it is what an apply matches on, so renaming an
	// entity while keeping its key is a rename rather than a delete-and-create.
	Key         string                     `json:"key"`
	Name        string                     `json:"name"`
	Description *string                    `json:"description,omitempty"`
	Rules       []OrgConfigCostCentreRules `json:"rules,omitempty"`
}

// OrgConfigCostFilter is the `OrgConfigCostFilter` schema.
type OrgConfigCostFilter struct {
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// OrgConfigCustomGraph is the `OrgConfigCustomGraph` schema.
type OrgConfigCustomGraph struct {
	// Key: Stable slug identifying this entity across organizations. Derived
	// from the name on export; it is what an apply matches on, so renaming an
	// entity while keeping its key is a rename rather than a delete-and-create.
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Source: The graph's TypeScript source.
	Source string `json:"source"`
}

// OrgConfigDashboard is the `OrgConfigDashboard` schema.
type OrgConfigDashboard struct {
	// Key: Stable slug identifying this entity across organizations. Derived
	// from the name on export; it is what an apply matches on, so renaming an
	// entity while keeping its key is a rename rather than a delete-and-create.
	Key       string                   `json:"key"`
	Name      string                   `json:"name"`
	IsDefault *bool                    `json:"isDefault,omitempty"`
	Cards     []OrgConfigDashboardCard `json:"cards,omitempty"`
}

// OrgConfigDashboardCard: One card. Position is the index in the dashboard's
// `cards` array — the grid order all three card kinds share.
type OrgConfigDashboardCard = any

// OrgConfigDocument: An organization's configuration. Every section is optional
// — a document that omits one leaves it entirely alone, in both apply modes.
type OrgConfigDocument struct {
	Version       *int64                         `json:"version,omitempty"`
	ExportedAt    *string                        `json:"exportedAt,omitempty"`
	ExportedFrom  *OrgConfigDocumentExportedFrom `json:"exportedFrom,omitempty"`
	Budgets       []OrgConfigBudget              `json:"budgets,omitempty"`
	CustomGraphs  []OrgConfigCustomGraph         `json:"customGraphs,omitempty"`
	Workflows     []OrgConfigWorkflow            `json:"workflows,omitempty"`
	Dashboards    []OrgConfigDashboard           `json:"dashboards,omitempty"`
	MetricAlerts  []OrgConfigMetricAlert         `json:"metricAlerts,omitempty"`
	Probes        []OrgConfigProbe               `json:"probes,omitempty"`
	CostCentres   []OrgConfigCostCentre          `json:"costCentres,omitempty"`
	TagPolicy     *OrgConfigDocumentTagPolicy    `json:"tagPolicy,omitempty"`
	AlertSettings *OrgConfigAlertSettings        `json:"alertSettings,omitempty"`
}

// OrgConfigMetricAlert is the `OrgConfigMetricAlert` schema.
type OrgConfigMetricAlert struct {
	// Key: Stable slug identifying this entity across organizations. Derived
	// from the name on export; it is what an apply matches on, so renaming an
	// entity while keeping its key is a rename rather than a delete-and-create.
	Key            string  `json:"key"`
	Name           string  `json:"name"`
	PluginID       *string `json:"pluginId,omitempty"`
	ResourceTypeID *string `json:"resourceTypeId,omitempty"`
	TagKey         *string `json:"tagKey,omitempty"`
	TagValue       *string `json:"tagValue,omitempty"`
	MetricKey      string  `json:"metricKey"`
	// Comparator: One of ">", ">=", "<", "<=".
	Comparator      string  `json:"comparator"`
	Threshold       float64 `json:"threshold"`
	ForMinutes      *int64  `json:"forMinutes,omitempty"`
	CooldownMinutes *int64  `json:"cooldownMinutes,omitempty"`
	Enabled         *bool   `json:"enabled,omitempty"`
}

// OrgConfigPlan is the `OrgConfigPlan` schema.
type OrgConfigPlan struct {
	// Mode: One of "merge", "replace".
	Mode       string                `json:"mode"`
	Changes    []OrgConfigChange     `json:"changes"`
	Unresolved []OrgConfigUnresolved `json:"unresolved"`
	Counts     OrgConfigPlanCounts   `json:"counts"`
}

// OrgConfigProbe is the `OrgConfigProbe` schema.
type OrgConfigProbe struct {
	// Key: Stable slug identifying this entity across organizations. Derived
	// from the name on export; it is what an apply matches on, so renaming an
	// entity while keeping its key is a rename rather than a delete-and-create.
	Key              string  `json:"key"`
	Name             string  `json:"name"`
	URL              string  `json:"url"`
	Method           *string `json:"method,omitempty"`
	IntervalSeconds  *int64  `json:"intervalSeconds,omitempty"`
	TimeoutMs        *int64  `json:"timeoutMs,omitempty"`
	FailureThreshold *int64  `json:"failureThreshold,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty"`
}

// OrgConfigRequest is the `OrgConfigRequest` schema.
type OrgConfigRequest struct {
	Document OrgConfigDocument `json:"document"`
	// Mode: `merge` creates and updates what the document names and leaves
	// everything else alone. `replace` additionally deletes entities the
	// document does not name, within the sections it carries.
	//
	// One of "merge", "replace".
	Mode *string `json:"mode,omitempty"`
}

// OrgConfigSection is the `OrgConfigSection` schema.
type OrgConfigSection = string

// The values OrgConfigSection takes.
const (
	OrgConfigSectionBudgets       OrgConfigSection = "budgets"
	OrgConfigSectionCustomGraphs  OrgConfigSection = "customGraphs"
	OrgConfigSectionWorkflows     OrgConfigSection = "workflows"
	OrgConfigSectionDashboards    OrgConfigSection = "dashboards"
	OrgConfigSectionMetricAlerts  OrgConfigSection = "metricAlerts"
	OrgConfigSectionProbes        OrgConfigSection = "probes"
	OrgConfigSectionCostCentres   OrgConfigSection = "costCentres"
	OrgConfigSectionTagPolicy     OrgConfigSection = "tagPolicy"
	OrgConfigSectionAlertSettings OrgConfigSection = "alertSettings"
)

// OrgConfigUnresolved: Something the document asked for that this organization
// could not satisfy — a pin for a resource nobody has synced, an account name
// that does not exist here. Not fatal: the affected card, clause or deletion is
// dropped and the rest of the document still applies.
type OrgConfigUnresolved struct {
	Section OrgConfigSection `json:"section"`
	Key     string           `json:"key"`
	Detail  string           `json:"detail"`
}

// OrgConfigWorkflow: A workflow. The git-webhook signing secret is deliberately
// absent — it is write-only, so a document can neither leak nor set one.
type OrgConfigWorkflow struct {
	// Key: Stable slug identifying this entity across organizations. Derived
	// from the name on export; it is what an apply matches on, so renaming an
	// entity while keeping its key is a rename rather than a delete-and-create.
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Source: The workflow's TypeScript source.
	Source  string                     `json:"source"`
	Trigger OrgConfigWorkflowTrigger   `json:"trigger,omitempty"`
	Metrics []OrgConfigWorkflowMetrics `json:"metrics,omitempty"`
	Enabled *bool                      `json:"enabled,omitempty"`
}

// OrgConfigWorkflowTrigger is the `OrgConfigWorkflowTrigger` schema.
type OrgConfigWorkflowTrigger = any

// OrgMember is the `OrgMember` schema.
type OrgMember struct {
	ID            string           `json:"id"`
	Email         string           `json:"email"`
	DisplayName   *string          `json:"displayName"`
	Role          OrganizationRole `json:"role"`
	RoleID        *string          `json:"roleId"`
	RoleName      *string          `json:"roleName"`
	RoleSystemKey *string          `json:"roleSystemKey"`
	CreatedAt     string           `json:"createdAt"`
}

// OrgMembership is the `OrgMembership` schema.
type OrgMembership struct {
	ID          string           `json:"id"`
	DisplayName string           `json:"displayName"`
	Role        OrganizationRole `json:"role"`
}

// OrgStatusIncident is the `OrgStatusIncident` schema.
type OrgStatusIncident struct {
	// ID: Cached incident row id.
	ID       string `json:"id"`
	PluginID string `json:"pluginId"`
	// PluginName: Provider display name, e.g. "DigitalOcean".
	PluginName string                 `json:"pluginName"`
	Title      string                 `json:"title"`
	State      ProviderIncidentState  `json:"state"`
	Impact     ProviderIncidentImpact `json:"impact"`
	// URL: Deep link to the provider's incident page or status page.
	URL          *string `json:"url"`
	StartedAt    string  `json:"startedAt"`
	ResolvedAt   *string `json:"resolvedAt"`
	LastUpdateAt *string `json:"lastUpdateAt"`
	// LastUpdateText: Plain-text body of the provider's most recent update.
	LastUpdateText *string `json:"lastUpdateText"`
	// Regions: Plugin-native region ids the provider reports as affected.
	Regions []string `json:"regions"`
	// Services: Human-readable affected provider services/products.
	Services []string `json:"services"`
	// ProviderWide: True when the incident affects the provider as a whole.
	ProviderWide bool `json:"providerWide"`
	// AffectedResourceCount: How many of the organization's resources the
	// incident overlaps.
	AffectedResourceCount int64 `json:"affectedResourceCount"`
	// AffectedRegions: The subset of `regions` where the organization actually
	// holds resources.
	AffectedRegions []string `json:"affectedRegions"`
	// SampleResources: Up to five of the overlapped resources, for display.
	SampleResources []ProviderIncidentResourceSample `json:"sampleResources"`
	// OverlappingChangeCount: Change-timeline events recorded on this provider
	// during the incident window — "these N changes happened during an
	// incident".
	OverlappingChangeCount int64 `json:"overlappingChangeCount"`
}

// OrgStatusIncidentsResponse is the `OrgStatusIncidentsResponse` schema.
type OrgStatusIncidentsResponse struct {
	Incidents []OrgStatusIncident `json:"incidents"`
}

// Organization is the `Organization` schema.
type Organization struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

// OrganizationRef is the `OrganizationRef` schema.
type OrganizationRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrganizationRole is the `OrganizationRole` schema.
type OrganizationRole = string

// The values OrganizationRole takes.
const (
	OrganizationRoleOwner  OrganizationRole = "owner"
	OrganizationRoleAdmin  OrganizationRole = "admin"
	OrganizationRoleMember OrganizationRole = "member"
)

// OrphanAccountGroup is the `OrphanAccountGroup` schema.
type OrphanAccountGroup struct {
	AccountID   string             `json:"accountId"`
	AccountName string             `json:"accountName"`
	PluginID    PluginID           `json:"pluginId"`
	PluginName  string             `json:"pluginName"`
	Resources   []OrphanedResource `json:"resources"`
}

// OrphanCostAnnotation: Best-effort trailing spend matched from collected
// per-resource cost rows; null when the provider reports no per-resource cost.
// The flag itself never depends on billing data.
//
// The API may send null in its place.
type OrphanCostAnnotation struct {
	// Amount: Spend over the trailing cost window.
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// OrphanListResponse is the `OrphanListResponse` schema.
type OrphanListResponse struct {
	// Accounts: Groups sorted by account name.
	Accounts   []OrphanAccountGroup `json:"accounts"`
	TotalCount int64                `json:"totalCount"`
	// UnownedCount: Flagged resources with no recorded owner — the 'nobody to
	// ask' count.
	UnownedCount int64 `json:"unownedCount"`
	// CostWindowDays: Days of trailing spend the annotations cover.
	CostWindowDays int64  `json:"costWindowDays"`
	GeneratedAt    string `json:"generatedAt"`
}

// OrphanedResource is the `OrphanedResource` schema.
type OrphanedResource struct {
	// ID: Infrawrench resource id.
	ID               string   `json:"id"`
	PluginID         PluginID `json:"pluginId"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	DisplayName      string   `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
	// Reason: Plugin-authored explanation of why this resource looks wasted.
	Reason       string                   `json:"reason"`
	Cost         *OrphanCostAnnotation    `json:"cost"`
	Owner        *ResourceOwnerAnnotation `json:"owner"`
	LastSyncedAt *string                  `json:"lastSyncedAt"`
}

// OversizedAccountGroup is the `OversizedAccountGroup` schema.
type OversizedAccountGroup struct {
	AccountID   string              `json:"accountId"`
	AccountName string              `json:"accountName"`
	PluginID    PluginID            `json:"pluginId"`
	PluginName  string              `json:"pluginName"`
	Resources   []OversizedResource `json:"resources"`
}

// OversizedResource is the `OversizedResource` schema.
type OversizedResource struct {
	// ID: Infrawrench resource id.
	ID               string   `json:"id"`
	PluginID         PluginID `json:"pluginId"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	DisplayName      string   `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
	// SizeFieldKey: Field to submit through the resource-update endpoint to
	// apply the recommended size.
	SizeFieldKey string `json:"sizeFieldKey"`
	// Region: Provider region/zone/location the resource lives in.
	Region          *string              `json:"region"`
	CurrentSize     OversizedSizeSummary `json:"currentSize"`
	RecommendedSize OversizedSizeSummary `json:"recommendedSize"`
	// CPUP95: p95 CPU utilisation over the window, percent of the current size.
	CPUP95 float64 `json:"cpuP95"`
	// MemoryP95: p95 memory utilisation, percent of the current size; null when
	// unmeasured.
	MemoryP95 *float64 `json:"memoryP95"`
	// MemoryMeasured: False when the provider stores no memory series for this
	// resource.
	MemoryMeasured bool `json:"memoryMeasured"`
	// ProjectedCPUP95: Projected p95 CPU on the recommended size, for the
	// confirm dialog.
	ProjectedCPUP95 float64 `json:"projectedCpuP95"`
	// Currency: ISO 4217 code the size prices are quoted in.
	Currency string `json:"currency"`
	// MonthlySaving: Current minus recommended monthly price; null when either
	// side is unpriced.
	MonthlySaving *float64 `json:"monthlySaving"`
	// ResizeNote: Plugin-authored caveat (e.g. the provider requires the machine
	// stopped).
	ResizeNote   *string `json:"resizeNote"`
	LastSyncedAt *string `json:"lastSyncedAt"`
}

// OversizedSizeSummary is the `OversizedSizeSummary` schema.
type OversizedSizeSummary struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Vcpus    int64  `json:"vcpus"`
	MemoryMb int64  `json:"memoryMb"`
	// PriceMonthly: Monthly catalog price in `currency`; null when unpriced.
	PriceMonthly *float64 `json:"priceMonthly"`
}

// OwnerCandidate is the `OwnerCandidate` schema.
type OwnerCandidate struct {
	UserID string `json:"userId"`
	// Name: Display name, falling back to the email.
	Name  string `json:"name"`
	Email string `json:"email"`
}

// OwnerCandidateListResponse is the `OwnerCandidateListResponse` schema.
type OwnerCandidateListResponse struct {
	Members []OwnerCandidate `json:"members"`
}

// OwnershipBlocker is the `OwnershipBlocker` schema.
type OwnershipBlocker struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// MemberCount: People in the organization
	MemberCount int64 `json:"memberCount"`
}

// OwnershipTransferRequired is the `OwnershipTransferRequired` schema.
type OwnershipTransferRequired struct {
	Error string `json:"error"`
	// Code: One of "transfer_ownership_required".
	Code          string             `json:"code"`
	Organizations []OwnershipBlocker `json:"organizations"`
}

// PageClearResponse is the `PageClearResponse` schema.
type PageClearResponse struct {
	// Cleared: False when the key had no cooldown to clear.
	Cleared bool `json:"cleared"`
}

// PageRequest is the `PageRequest` schema.
type PageRequest struct {
	// Source: Stable name for the system raising the page: letters, digits, `.`,
	// `_` and `-`. It is the notification's sender, and it scopes the cooldown —
	// two services paging under the same key never throttle each other.
	Source string `json:"source"`
	// Message: The alert text. Becomes the SMS and notification body.
	Message string `json:"message"`
	// Title: Short headline for the notification. Defaults to `source`.
	Title *string `json:"title,omitempty"`
	// Key: Throttle key, `default` when unset. Pages sharing a key are
	// suppressed while that key is in cooldown, so a per-object key (a host, a
	// cluster id) alerts per object while the default key alerts once for the
	// whole source.
	Key *string `json:"key,omitempty"`
	// CooldownMinutes: Minutes to suppress repeat pages under the same key.
	// Defaults to 60; `0` sends every time.
	CooldownMinutes *int64 `json:"cooldownMinutes,omitempty"`
	// Voice: Also place a voice call to recipients who opted into voice. Off by
	// default — reserve it for things worth waking someone up for.
	Voice *bool `json:"voice,omitempty"`
}

// PageResponse is the `PageResponse` schema.
type PageResponse struct {
	// Delivered: True when at least one recipient was reached on any transport.
	Delivered bool `json:"delivered"`
	// Suppressed: True when the key was still in cooldown, so nothing was sent.
	Suppressed bool `json:"suppressed"`
	// Sms: Twilio deliveries (SMS + voice) that Twilio accepted.
	Sms int64 `json:"sms"`
	// Push: Push notifications accepted by Expo.
	Push int64 `json:"push"`
	// Slack: Slack channel posts Slack accepted.
	Slack int64 `json:"slack"`
	// MsTeams: Microsoft Teams webhook posts Teams accepted.
	MsTeams int64 `json:"msTeams"`
	// RetryAt: When suppressed, the time at which this key can page again.
	RetryAt *string `json:"retryAt,omitempty"`
}

// PeerPane is the `PeerPane` schema.
type PeerPane struct {
	TabLabel      string     `json:"tabLabel"`
	PluginLogoSvg string     `json:"pluginLogoSvg"`
	PeerPluginID  string     `json:"peerPluginId"`
	Schema        JSONObject `json:"schema"`
}

// PeerPaneStub is the `PeerPaneStub` schema.
type PeerPaneStub struct {
	TabLabel      string `json:"tabLabel"`
	PluginLogoSvg string `json:"pluginLogoSvg"`
	PeerPluginID  string `json:"peerPluginId"`
}

// PeerPanesRequest is the `PeerPanesRequest` schema.
type PeerPanesRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceID       ResourceID  `json:"resourceId"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// Permission: A permission string. Roles may grant exact permissions like the
// entries in this enum, or wildcards (e.g. `resources:*:read`, `*`).
type Permission = string

// The values Permission takes.
const (
	PermissionAccountsRead           Permission = "accounts:read"
	PermissionAccountsWrite          Permission = "accounts:write"
	PermissionAccountsDelete         Permission = "accounts:delete"
	PermissionResourcesRead          Permission = "resources:read"
	PermissionResourcesWrite         Permission = "resources:write"
	PermissionResourcesDelete        Permission = "resources:delete"
	PermissionResourcesExecute       Permission = "resources:execute"
	PermissionSecretsRead            Permission = "secrets:read"
	PermissionSecretsWrite           Permission = "secrets:write"
	PermissionStorageRead            Permission = "storage:read"
	PermissionStorageWrite           Permission = "storage:write"
	PermissionDashboardsRead         Permission = "dashboards:read"
	PermissionDashboardsWrite        Permission = "dashboards:write"
	PermissionWorkflowsRead          Permission = "workflows:read"
	PermissionWorkflowsWrite         Permission = "workflows:write"
	PermissionWorkflowsApprove       Permission = "workflows:approve"
	PermissionDeploymentsRead        Permission = "deployments:read"
	PermissionDeploymentsPlan        Permission = "deployments:plan"
	PermissionDeploymentsWrite       Permission = "deployments:write"
	PermissionCostsRead              Permission = "costs:read"
	PermissionCostsWrite             Permission = "costs:write"
	PermissionBudgetsRead            Permission = "budgets:read"
	PermissionBudgetsWrite           Permission = "budgets:write"
	PermissionMetricAlertsRead       Permission = "metric-alerts:read"
	PermissionMetricAlertsWrite      Permission = "metric-alerts:write"
	PermissionFreezesRead            Permission = "freezes:read"
	PermissionFreezesWrite           Permission = "freezes:write"
	PermissionFreezesOverride        Permission = "freezes:override"
	PermissionIncidentsRead          Permission = "incidents:read"
	PermissionIncidentsWrite         Permission = "incidents:write"
	PermissionTagPolicyOverride      Permission = "tag-policy:override"
	PermissionConfigRead             Permission = "config:read"
	PermissionConfigWrite            Permission = "config:write"
	PermissionIacRead                Permission = "iac:read"
	PermissionIacWrite               Permission = "iac:write"
	PermissionAuditRead              Permission = "audit:read"
	PermissionAccessRead             Permission = "access:read"
	PermissionAccessRequest          Permission = "access:request"
	PermissionAccessApprove          Permission = "access:approve"
	PermissionTeamRead               Permission = "team:read"
	PermissionTeamInvite             Permission = "team:invite"
	PermissionTeamRoleWrite          Permission = "team:role:write"
	PermissionTeamRemove             Permission = "team:remove"
	PermissionApikeysRead            Permission = "apikeys:read"
	PermissionApikeysWrite           Permission = "apikeys:write"
	PermissionBillingRead            Permission = "billing:read"
	PermissionBillingWrite           Permission = "billing:write"
	PermissionSSHKeysRead            Permission = "ssh-keys:read"
	PermissionSSHKeysWrite           Permission = "ssh-keys:write"
	PermissionSessionRecordingsRead  Permission = "session-recordings:read"
	PermissionSessionRecordingsWrite Permission = "session-recordings:write"
	PermissionBastionsRead           Permission = "bastions:read"
	PermissionBastionsWrite          Permission = "bastions:write"
	PermissionChatRead               Permission = "chat:read"
	PermissionChatWrite              Permission = "chat:write"
	PermissionJiraRead               Permission = "jira:read"
	PermissionJiraWrite              Permission = "jira:write"
	PermissionLinearRead             Permission = "linear:read"
	PermissionLinearWrite            Permission = "linear:write"
	PermissionInvoicesRead           Permission = "invoices:read"
	PermissionInvoicesWrite          Permission = "invoices:write"
	PermissionInvoicesIssue          Permission = "invoices:issue"
	PermissionPagesWrite             Permission = "pages:write"
	PermissionOrgSettingsWrite       Permission = "org:settings:write"
)

// PermissionCatalog is the `PermissionCatalog` schema.
type PermissionCatalog struct {
	Permissions []Permission `json:"permissions"`
}

// PickerResource is the `PickerResource` schema.
type PickerResource struct {
	ID             ResourceID `json:"id"`
	Label          string     `json:"label"`
	PluginID       string     `json:"pluginId"`
	ResourceTypeID string     `json:"resourceTypeId"`
	AccountID      string     `json:"accountId"`
	OutputKey      string     `json:"outputKey"`
	OutputValue    string     `json:"outputValue"`
}

// PickerResourcesRequest is the `PickerResourcesRequest` schema.
type PickerResourcesRequest struct {
	Sources      []PickerResourcesRequestSources `json:"sources"`
	AccountID    string                          `json:"accountId"`
	RegionHint   *string                         `json:"regionHint,omitempty"`
	CrossAccount *bool                           `json:"crossAccount,omitempty"`
}

// PinFull is the `PinFull` schema.
type PinFull struct {
	PinID             string      `json:"pinId"`
	ResourceID        ResourceID  `json:"resourceId"`
	GridX             int64       `json:"gridX"`
	GridY             int64       `json:"gridY"`
	GridW             int64       `json:"gridW"`
	GridH             int64       `json:"gridH"`
	DisplayName       string      `json:"displayName"`
	PluginID          string      `json:"pluginId"`
	ResourceTypeID    string      `json:"resourceTypeId"`
	AccountID         string      `json:"accountId"`
	FieldsJSON        JSONObject  `json:"fieldsJson"`
	OutputsJSON       JSONObject  `json:"outputsJson"`
	PluginLogoSvg     string      `json:"pluginLogoSvg"`
	PluginDisplayName string      `json:"pluginDisplayName"`
	Status            ProbeStatus `json:"status"`
}

// PinRangeMetricSeries is the `PinRangeMetricSeries` schema.
type PinRangeMetricSeries struct {
	Label  string                       `json:"label"`
	Unit   *string                      `json:"unit,omitempty"`
	Points []PinRangeMetricSeriesPoints `json:"points"`
}

// PinRangeResponse is the `PinRangeResponse` schema.
type PinRangeResponse struct {
	Series []PinRangeMetricSeries `json:"series"`
}

// PinRequest is the `PinRequest` schema.
type PinRequest struct {
	DashboardID string     `json:"dashboardId"`
	ResourceID  ResourceID `json:"resourceId"`
	GridX       *int64     `json:"gridX,omitempty"`
	GridY       *int64     `json:"gridY,omitempty"`
}

// PluginID: Manifest id of an installed plugin.
//
// Spec schema: `PluginId`.
type PluginID = string

// The values PluginID takes.
const (
	PluginIDAnthropic    PluginID = "anthropic"
	PluginIDAssemblyai   PluginID = "assemblyai"
	PluginIDAWS          PluginID = "aws"
	PluginIDAzure        PluginID = "azure"
	PluginIDCartesia     PluginID = "cartesia"
	PluginIDClickhouse   PluginID = "clickhouse"
	PluginIDCloudflare   PluginID = "cloudflare"
	PluginIDCloudinary   PluginID = "cloudinary"
	PluginIDCohere       PluginID = "cohere"
	PluginIDDatabricks   PluginID = "databricks"
	PluginIDDeepgram     PluginID = "deepgram"
	PluginIDDeepseek     PluginID = "deepseek"
	PluginIDDigitalocean PluginID = "digitalocean"
	PluginIDDocker       PluginID = "docker"
	PluginIDElevenlabs   PluginID = "elevenlabs"
	PluginIDFireworks    PluginID = "fireworks"
	PluginIDFly          PluginID = "fly"
	PluginIDGCP          PluginID = "gcp"
	PluginIDGemini       PluginID = "gemini"
	PluginIDGladia       PluginID = "gladia"
	PluginIDGroq         PluginID = "groq"
	PluginIDHetzner      PluginID = "hetzner"
	PluginIDKafka        PluginID = "kafka"
	PluginIDKubernetes   PluginID = "kubernetes"
	PluginIDMemcached    PluginID = "memcached"
	PluginIDMistral      PluginID = "mistral"
	PluginIDMongodb      PluginID = "mongodb"
	PluginIDMssql        PluginID = "mssql"
	PluginIDMysql        PluginID = "mysql"
	PluginIDNeon         PluginID = "neon"
	PluginIDNetlify      PluginID = "netlify"
	PluginIDOpenai       PluginID = "openai"
	PluginIDOpenrouter   PluginID = "openrouter"
	PluginIDOpensearch   PluginID = "opensearch"
	PluginIDOVH          PluginID = "ovh"
	PluginIDPlanetscale  PluginID = "planetscale"
	PluginIDPostgres     PluginID = "postgres"
	PluginIDRedis        PluginID = "redis"
	PluginIDReplicate    PluginID = "replicate"
	PluginIDRevai        PluginID = "revai"
	PluginIDScaleway     PluginID = "scaleway"
	PluginIDSpeechmatics PluginID = "speechmatics"
	PluginIDSSH          PluginID = "ssh"
	PluginIDTogether     PluginID = "together"
	PluginIDTurso        PluginID = "turso"
	PluginIDUploadthing  PluginID = "uploadthing"
	PluginIDVercel       PluginID = "vercel"
	PluginIDWorkos       PluginID = "workos"
	PluginIDXai          PluginID = "xai"
)

// PluginSummary is the `PluginSummary` schema.
type PluginSummary struct {
	ID               string                `json:"id"`
	DisplayName      string                `json:"displayName"`
	LogoSvg          string                `json:"logoSvg"`
	CredentialFields []CredentialField     `json:"credentialFields"`
	Preflight        *PreflightDeclaration `json:"preflight"`
}

// PolicyTemplate is the `PolicyTemplate` schema.
type PolicyTemplate struct {
	FormatLabel string `json:"formatLabel"`
	// Language: One of "json", "yaml", "text".
	Language     string                  `json:"language"`
	Document     string                  `json:"document"`
	Instructions *string                 `json:"instructions,omitempty"`
	HelpLink     *PolicyTemplateHelpLink `json:"helpLink,omitempty"`
}

// PolicyTemplateResponse is the `PolicyTemplateResponse` schema.
type PolicyTemplateResponse struct {
	Template PolicyTemplate `json:"template"`
}

// PostureAlertSettings is the `PostureAlertSettings` schema.
type PostureAlertSettings struct {
	// Enabled: Whether the poller sends posture alerts for this organization at
	// all.
	Enabled bool `json:"enabled"`
	// LastNotifiedAt: When the organization's posture alert scan last completed,
	// or null before the first. Owned by the poller's cooldown claim; not
	// writable through this API.
	LastNotifiedAt *string `json:"lastNotifiedAt"`
}

// PostureAlertSettingsUpdate is the `PostureAlertSettingsUpdate` schema.
type PostureAlertSettingsUpdate struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// PostureDismissal is the `PostureDismissal` schema.
type PostureDismissal struct {
	ResourceID string `json:"resourceId"`
	RuleID     string `json:"ruleId"`
	// DismissedAt: When the finding was accepted.
	DismissedAt string `json:"dismissedAt"`
	// DismissedBy: Display name or email of whoever accepted it; null when
	// unknown.
	DismissedBy *string `json:"dismissedBy"`
	// Reason: The operator's note, when they left one.
	Reason *string `json:"reason"`
}

// PostureDismissalCreate is the `PostureDismissalCreate` schema.
type PostureDismissalCreate struct {
	// ResourceID: Infrawrench resource id the finding is on.
	ResourceID string `json:"resourceId"`
	// RuleID: The matched rule's id.
	RuleID string `json:"ruleId"`
	// Reason: Why this finding is acceptable. Trimmed; an empty note is stored
	// as none.
	Reason *string `json:"reason,omitempty"`
}

// PostureFinding is the `PostureFinding` schema.
type PostureFinding struct {
	// ResourceID: Infrawrench resource id.
	ResourceID       string   `json:"resourceId"`
	PluginID         PluginID `json:"pluginId"`
	PluginName       string   `json:"pluginName"`
	ResourceTypeID   string   `json:"resourceTypeId"`
	ResourceTypeName string   `json:"resourceTypeName"`
	AccountID        string   `json:"accountId"`
	AccountName      string   `json:"accountName"`
	DisplayName      string   `json:"displayName"`
	// ExternalID: Provider-native id, when known.
	ExternalID *string `json:"externalId"`
	// RuleID: The matched rule's stable id, unique within the plugin.
	RuleID string `json:"ruleId"`
	// Title: Short rule title.
	Title string `json:"title"`
	// Severity: How bad the finding is. `critical` and `high` findings feed the
	// posture alerts; `medium` and `low` are hygiene work surfaced on the
	// posture screen only.
	//
	// One of "critical", "high", "medium", "low".
	Severity string `json:"severity"`
	// Category: Grouping bucket for what kind of exposure the finding describes.
	//
	// One of "public-exposure", "encryption", "credential-age",
	// "data-protection", "other".
	Category string `json:"category"`
	// Reason: Plugin-authored explanation of why this is a finding.
	Reason string `json:"reason"`
}

// PostureListResponse is the `PostureListResponse` schema.
type PostureListResponse struct {
	// Findings: Live findings, worst severity first. Dismissed findings are not
	// included.
	Findings []PostureFinding `json:"findings"`
	// TotalCount: Live finding count; dismissals excluded.
	TotalCount int64                 `json:"totalCount"`
	Counts     PostureSeverityCounts `json:"counts"`
	// Dismissed: Findings a dismissal is currently suppressing, most recently
	// dismissed first. Only dismissals whose rule still matches appear, so a
	// finding that has since been fixed simply drops out.
	Dismissed      []DismissedPostureFinding `json:"dismissed"`
	DismissedCount int64                     `json:"dismissedCount"`
	GeneratedAt    string                    `json:"generatedAt"`
}

// PostureSeverityCounts: Live finding count per severity; every bucket present,
// zeros included.
type PostureSeverityCounts struct {
	Critical int64 `json:"critical"`
	High     int64 `json:"high"`
	Medium   int64 `json:"medium"`
	Low      int64 `json:"low"`
}

// PreflightCapability is the `PreflightCapability` schema.
type PreflightCapability struct {
	ID                  string                `json:"id"`
	Label               string                `json:"label"`
	Description         *string               `json:"description,omitempty"`
	RequiredPermissions []PreflightPermission `json:"requiredPermissions"`
	Essential           *bool                 `json:"essential,omitempty"`
}

// PreflightCheck is the `PreflightCheck` schema.
type PreflightCheck struct {
	CapabilityID string `json:"capabilityId"`
	// Status: One of "ok", "missing", "unknown".
	Status             string                  `json:"status"`
	MissingPermissions []PreflightPermission   `json:"missingPermissions"`
	Message            *string                 `json:"message"`
	HelpLink           *PreflightCheckHelpLink `json:"helpLink"`
}

// PreflightDeclaration: Declared when the plugin supports credential preflight
// (per-capability permission checks). `null` for plugins without it.
//
// The API may send null in its place.
type PreflightDeclaration struct {
	Capabilities   []PreflightCapability               `json:"capabilities"`
	TemplateFormat *PreflightDeclarationTemplateFormat `json:"templateFormat,omitempty"`
}

// PreflightPermission is the `PreflightPermission` schema.
type PreflightPermission struct {
	// ID: Provider-native permission string, e.g. `ce:GetCostAndUsage`.
	ID    string `json:"id"`
	Label string `json:"label"`
}

// PreflightReport is the `PreflightReport` schema.
type PreflightReport struct {
	PluginID  string `json:"pluginId"`
	Supported bool   `json:"supported"`
	// Identity: Provider-side identity the credential resolved to (ARN, service
	// account…).
	Identity *string          `json:"identity"`
	Checks   []PreflightCheck `json:"checks"`
}

// PreflightRequest is the `PreflightRequest` schema.
type PreflightRequest struct {
	PluginID    string            `json:"pluginId"`
	Credentials map[string]string `json:"credentials"`
	// BastionID: Probe through this bastion, matching how the account will
	// egress once created.
	BastionID *string `json:"bastionId,omitempty"`
}

// ProbeMetricSeries is the `ProbeMetricSeries` schema.
type ProbeMetricSeries struct {
	// Label: "Latency" (ms) or "Up" (1/0).
	Label  string                    `json:"label"`
	Unit   *string                   `json:"unit,omitempty"`
	Points []ProbeMetricSeriesPoints `json:"points"`
}

// ProbeMetrics is the `ProbeMetrics` schema.
type ProbeMetrics struct {
	Series []ProbeMetricSeries `json:"series"`
}

// ProbeRequest is the `ProbeRequest` schema.
type ProbeRequest struct {
	Items []ProbeRequestItems `json:"items"`
}

// ProbeStatus is the `ProbeStatus` schema.
type ProbeStatus struct {
	// Phase: One of "ok", "error".
	Phase          string                      `json:"phase"`
	Error          *string                     `json:"error,omitempty"`
	Stats          []JSONObject                `json:"stats,omitempty"`
	Sparkline      []ProbeStatusSparkline      `json:"sparkline,omitempty"`
	SparklineLabel *string                     `json:"sparklineLabel,omitempty"`
	ResourceCounts []ProbeStatusResourceCounts `json:"resourceCounts,omitempty"`
}

// ProbeSuggestion is the `ProbeSuggestion` schema.
type ProbeSuggestion struct {
	// URL: Normalized to an absolute URL — bare hosts get https://.
	URL            string   `json:"url"`
	ResourceID     string   `json:"resourceId"`
	DisplayName    string   `json:"displayName"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	AccountID      string   `json:"accountId"`
	// OutputKey: The output/field key the URL was mined from.
	OutputKey string `json:"outputKey"`
}

// ProbeSuggestions is the `ProbeSuggestions` schema.
type ProbeSuggestions struct {
	Suggestions []ProbeSuggestion `json:"suggestions"`
}

// Profile is the `Profile` schema.
type Profile struct {
	ID                string  `json:"id"`
	Email             string  `json:"email"`
	EmailVerified     bool    `json:"emailVerified"`
	FirstName         *string `json:"firstName"`
	LastName          *string `json:"lastName"`
	ProfilePictureURL *string `json:"profilePictureUrl"`
	LastSignInAt      *string `json:"lastSignInAt"`
	CreatedAt         string  `json:"createdAt"`
	// Identities: Connected OAuth accounts, if any
	Identities []ProfileIdentities `json:"identities"`
}

// ProfileSummary is the `ProfileSummary` schema.
type ProfileSummary struct {
	ID                string  `json:"id"`
	Email             string  `json:"email"`
	EmailVerified     bool    `json:"emailVerified"`
	FirstName         *string `json:"firstName"`
	LastName          *string `json:"lastName"`
	ProfilePictureURL *string `json:"profilePictureUrl"`
	LastSignInAt      *string `json:"lastSignInAt"`
	CreatedAt         string  `json:"createdAt"`
}

// ProviderIncidentImpact: Normalized incident severity, least to most severe.
type ProviderIncidentImpact = string

// The values ProviderIncidentImpact takes.
const (
	ProviderIncidentImpactMaintenance ProviderIncidentImpact = "maintenance"
	ProviderIncidentImpactMinor       ProviderIncidentImpact = "minor"
	ProviderIncidentImpactMajor       ProviderIncidentImpact = "major"
	ProviderIncidentImpactCritical    ProviderIncidentImpact = "critical"
)

// ProviderIncidentResourceSample is the `ProviderIncidentResourceSample` schema.
type ProviderIncidentResourceSample struct {
	// ID: Resource id.
	ID             string `json:"id"`
	DisplayName    string `json:"displayName"`
	ResourceTypeID string `json:"resourceTypeId"`
	// Region: The resource's region field, when it has one.
	Region *string `json:"region,omitempty"`
}

// ProviderIncidentState: Normalized incident lifecycle state as the provider
// reports it.
type ProviderIncidentState = string

// The values ProviderIncidentState takes.
const (
	ProviderIncidentStateInvestigating ProviderIncidentState = "investigating"
	ProviderIncidentStateIdentified    ProviderIncidentState = "identified"
	ProviderIncidentStateMonitoring    ProviderIncidentState = "monitoring"
	ProviderIncidentStateResolved      ProviderIncidentState = "resolved"
)

// PublicStatusComponent is the `PublicStatusComponent` schema.
type PublicStatusComponent struct {
	// ID: Stable per page. Deliberately not the probe id.
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	GroupName *string `json:"groupName"`
	// State: A component's public state. A paused probe reads `unknown`
	// regardless of its last result — the page is a claim about what is being
	// checked now.
	//
	// One of "operational", "degraded", "down", "unknown".
	State     string   `json:"state"`
	Uptime24h *float64 `json:"uptime24h"`
	// History: Oldest first; empty when history is hidden.
	History []StatusHistoryDay `json:"history"`
}

// PublicStatusPage is the `PublicStatusPage` schema.
type PublicStatusPage struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	// State: Rollup over the components. `degraded` means some but not all are
	// down; components with no data are ignored rather than dragging the page to
	// unknown.
	//
	// One of "operational", "degraded", "major_outage", "unknown".
	State string `json:"state"`
	// Summary: One sentence describing `state`.
	Summary     string                  `json:"summary"`
	Components  []PublicStatusComponent `json:"components"`
	SupportURL  *string                 `json:"supportUrl"`
	ShowHistory bool                    `json:"showHistory"`
	ShowUptime  bool                    `json:"showUptime"`
	HistoryDays int64                   `json:"historyDays"`
	GeneratedAt string                  `json:"generatedAt"`
}

// PushedCostRow is the `PushedCostRow` schema.
type PushedCostRow struct {
	// Date: UTC day the spend belongs to.
	Date     string `json:"date"`
	Currency string `json:"currency"`
	// Amount: Money for this day/dimension combination. Negative for credits.
	Amount float64 `json:"amount"`
	// Service: Becomes a group/filter value.
	Service *string `json:"service,omitempty"`
	Region  *string `json:"region,omitempty"`
	// ResourceID: Opaque id of the thing being billed; groups the `resource`
	// dimension.
	ResourceID *string `json:"resourceId,omitempty"`
	// Tags: Cost-allocation tags, at most 32. Keys starting with `infrawrench:`
	// are reserved and rejected.
	Tags        map[string]string `json:"tags,omitempty"`
	UsageAmount *float64          `json:"usageAmount,omitempty"`
	UsageUnit   *string           `json:"usageUnit,omitempty"`
	// AccountID: Attribute this row to a connected account. Must belong to the
	// calling organization. Omit to attribute it to the source itself.
	AccountID *string `json:"accountId,omitempty"`
}

// QueryMonitor is the `QueryMonitor` schema.
type QueryMonitor struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	AccountID      string  `json:"accountId"`
	AccountName    *string `json:"accountName"`
	ResourceID     *string `json:"resourceId"`
	ResourceTypeID *string `json:"resourceTypeId"`
	ResourceName   *string `json:"resourceName"`
	SQL            string  `json:"sql"`
	// Mode: How the result is reduced to one number. `scalar` reads the first
	// column of the first row; `rowCount` counts the rows, which is what lets
	// `SELECT … WHERE broken` be a monitor.
	//
	// One of "scalar", "rowCount".
	Mode string `json:"mode"`
	// Operator: One of "gt", "gte", "lt", "lte", "eq", "neq".
	Operator        string  `json:"operator"`
	Threshold       float64 `json:"threshold"`
	IntervalMinutes int64   `json:"intervalMinutes"`
	// ConsecutiveBreaches: Consecutive breaching runs before the alert fires. A
	// query against a live table is a sample: a count that dips while a batch
	// job is mid-write is not an incident, and a monitor that pages on it gets
	// muted within a week.
	ConsecutiveBreaches int64 `json:"consecutiveBreaches"`
	Enabled             bool  `json:"enabled"`
	// State: `unknown` is a first-class state, not an absence: a monitor whose
	// query failed has not told you the data is fine, and rendering that as `ok`
	// is how a broken monitor becomes indistinguishable from a healthy one.
	//
	// One of "ok", "breaching", "unknown".
	State     string   `json:"state"`
	LastValue *float64 `json:"lastValue"`
	LastRunAt *string  `json:"lastRunAt"`
	// LastError: Why the last run said nothing. Kept apart from the state
	// because 'the monitor is broken' and 'the data is bad' need different
	// people.
	LastError       *string `json:"lastError"`
	BreachStreak    int64   `json:"breachStreak"`
	LastAlertedAt   *string `json:"lastAlertedAt"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// QueryMonitorCreate is the `QueryMonitorCreate` schema.
type QueryMonitorCreate struct {
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
	AccountID      string  `json:"accountId"`
	ResourceID     *string `json:"resourceId,omitempty"`
	ResourceTypeID *string `json:"resourceTypeId,omitempty"`
	SQL            string  `json:"sql"`
	// Mode: How the result is reduced to one number. `scalar` reads the first
	// column of the first row; `rowCount` counts the rows, which is what lets
	// `SELECT … WHERE broken` be a monitor.
	//
	// One of "scalar", "rowCount".
	Mode string `json:"mode"`
	// Operator: One of "gt", "gte", "lt", "lte", "eq", "neq".
	Operator            string  `json:"operator"`
	Threshold           float64 `json:"threshold"`
	IntervalMinutes     int64   `json:"intervalMinutes"`
	ConsecutiveBreaches *int64  `json:"consecutiveBreaches,omitempty"`
	Enabled             *bool   `json:"enabled,omitempty"`
}

// QueryMonitorList is the `QueryMonitorList` schema.
type QueryMonitorList struct {
	Monitors []QueryMonitor `json:"monitors"`
}

// QueryMonitorTargetAccount is the `QueryMonitorTargetAccount` schema.
type QueryMonitorTargetAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// AccountSQL: The account itself has a SQL driver, so it is a valid target
	// on its own.
	AccountSQL bool                         `json:"accountSql"`
	Resources  []QueryMonitorTargetResource `json:"resources"`
}

// QueryMonitorTargetResource is the `QueryMonitorTargetResource` schema.
type QueryMonitorTargetResource struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ResourceTypeID string `json:"resourceTypeId"`
	// TypeName: The resource type's display name, e.g. 'D1 Database'.
	TypeName string `json:"typeName"`
}

// QueryMonitorTargets is the `QueryMonitorTargets` schema.
type QueryMonitorTargets struct {
	Accounts []QueryMonitorTargetAccount `json:"accounts"`
}

// QueryMonitorTestResult is the `QueryMonitorTestResult` schema.
type QueryMonitorTestResult struct {
	Value *float64 `json:"value"`
	// State: `unknown` is a first-class state, not an absence: a monitor whose
	// query failed has not told you the data is fine, and rendering that as `ok`
	// is how a broken monitor becomes indistinguishable from a healthy one.
	//
	// One of "ok", "breaching", "unknown".
	State      string  `json:"state"`
	Error      *string `json:"error"`
	DurationMs int64   `json:"durationMs"`
	// Rows: Up to 20 rows, for the preview.
	Rows []map[string]any `json:"rows"`
}

// QueryMonitorUpdate is the `QueryMonitorUpdate` schema.
type QueryMonitorUpdate struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	AccountID      *string `json:"accountId,omitempty"`
	ResourceID     *string `json:"resourceId,omitempty"`
	ResourceTypeID *string `json:"resourceTypeId,omitempty"`
	SQL            *string `json:"sql,omitempty"`
	// Mode: How the result is reduced to one number. `scalar` reads the first
	// column of the first row; `rowCount` counts the rows, which is what lets
	// `SELECT … WHERE broken` be a monitor.
	//
	// One of "scalar", "rowCount".
	Mode *string `json:"mode,omitempty"`
	// Operator: One of "gt", "gte", "lt", "lte", "eq", "neq".
	Operator            *string  `json:"operator,omitempty"`
	Threshold           *float64 `json:"threshold,omitempty"`
	IntervalMinutes     *int64   `json:"intervalMinutes,omitempty"`
	ConsecutiveBreaches *int64   `json:"consecutiveBreaches,omitempty"`
	Enabled             *bool    `json:"enabled,omitempty"`
}

// QuietHours: A recurring local-time window during which the rule holds its
// alerts. Held, not dropped — a held alert is queued and delivered when the
// window closes.
//
// The API may send null in its place.
type QuietHours struct {
	// Timezone: IANA zone, e.g. Europe/Berlin
	Timezone    string `json:"timezone"`
	StartMinute int64  `json:"startMinute"`
	// EndMinute: May be less than startMinute for an overnight window. Equal
	// means empty.
	EndMinute int64 `json:"endMinute"`
	// Days: ISO weekdays the window applies on, matched against the day the
	// window opened. Empty means every day.
	Days           []int64        `json:"days"`
	UrgentOverride *AlertSeverity `json:"urgentOverride"`
}

// QuotaAccountStatus is the `QuotaAccountStatus` schema.
type QuotaAccountStatus struct {
	AccountID   string   `json:"accountId"`
	AccountName string   `json:"accountName"`
	PluginID    PluginID `json:"pluginId"`
	// QuotaCount: Quota rows currently stored for this account.
	QuotaCount int64 `json:"quotaCount"`
	// LastPolledAt: Last successful collection; null if never.
	LastPolledAt *string `json:"lastPolledAt"`
	// LastError: Last collection failure, or null when the last pass succeeded.
	LastError          *string `json:"lastError"`
	LastErrorHelpLabel *string `json:"lastErrorHelpLabel"`
	// LastErrorHelpURL: Set when the failure was a fixable permission gap rather
	// than an outage.
	LastErrorHelpURL *string `json:"lastErrorHelpUrl"`
	// Partial: The plugin reports a representative subset of the provider's
	// quotas, not all of them. True for AWS and DigitalOcean.
	Partial bool `json:"partial"`
}

// QuotaAlertSettings is the `QuotaAlertSettings` schema.
type QuotaAlertSettings struct {
	// Enabled: Whether the poller sends quota alerts for this organization at
	// all.
	Enabled bool `json:"enabled"`
	// Threshold: Utilisation fraction at or above which a quota alerts. Default
	// 0.8. Bounded below at 0.5 (a lower threshold makes every quota critical)
	// and above at 0.99 (at 1.0 the provider is already refusing requests, so
	// the alert reports an outage rather than warning about one). Values outside
	// the range are rejected, not clamped.
	Threshold float64 `json:"threshold"`
	// LastNotifiedAt: When the organization's quota alert scan last completed,
	// or null before the first. Owned by the poller's cooldown claim; not
	// writable through this API.
	LastNotifiedAt *string `json:"lastNotifiedAt"`
}

// QuotaAlertSettingsUpdate is the `QuotaAlertSettingsUpdate` schema.
type QuotaAlertSettingsUpdate struct {
	Enabled   *bool    `json:"enabled,omitempty"`
	Threshold *float64 `json:"threshold,omitempty"`
}

// QuotaListResponse is the `QuotaListResponse` schema.
type QuotaListResponse struct {
	// Rows: Every quota with a reading, worst first.
	Rows []QuotaRow `json:"rows"`
	// Accounts: Per-account collection status for every account on a
	// quota-capable plugin. Present even when the account has rows: an empty
	// `rows` alone cannot distinguish 'nothing is near a limit' from 'every
	// collection is failing'.
	Accounts []QuotaAccountStatus `json:"accounts"`
	// Threshold: The organization's alert threshold as a fraction, so the page's
	// marker and the alert agree.
	Threshold float64 `json:"threshold"`
	// UnsupportedPluginIDs: Plugins the organization holds accounts with that
	// cannot report quotas at all. Named rather than counted, because the
	// absence is the finding.
	UnsupportedPluginIDs []PluginID `json:"unsupportedPluginIds"`
}

// QuotaRow is the `QuotaRow` schema.
type QuotaRow struct {
	// Key: Plugin-chosen stable id for this quota within the account.
	Key         string   `json:"key"`
	AccountID   string   `json:"accountId"`
	AccountName string   `json:"accountName"`
	PluginID    PluginID `json:"pluginId"`
	// Service: Provider service in the provider's own vocabulary.
	Service string `json:"service"`
	Name    string `json:"name"`
	// Region: Provider region, or null for an account-wide quota. Never the
	// string 'global'.
	Region *string `json:"region"`
	// Limit: The ceiling the provider will enforce, in `unit`.
	Limit float64 `json:"limit"`
	// Used: How much of `limit` is consumed, in the same unit.
	Used float64 `json:"used"`
	// Utilization: used / limit. Not clamped at 1 — an over-quota reading is a
	// real state.
	Utilization float64 `json:"utilization"`
	// Unit: What is being counted, in the provider's own word.
	Unit *string `json:"unit"`
	// Adjustable: Whether the provider lets the customer request an increase.
	// Null means the plugin does not know, which is not the same as `false`.
	Adjustable *bool `json:"adjustable"`
	// DocsURL: Provider page explaining or raising this quota.
	DocsURL *string `json:"docsUrl"`
	// ObservedAt: When this reading was collected.
	ObservedAt string `json:"observedAt"`
	// Severity: Where the quota sits: `exhausted` (used >= limit — the provider
	// is already refusing requests), `critical` (at or over the organization's
	// threshold), `trending` (under the threshold, but the fitted trend reaches
	// the limit within 30 days), or `ok`. Ordered: an exhausted quota is also
	// over threshold and also trending, and reports as `exhausted`.
	//
	// One of "exhausted", "critical", "trending", "ok".
	Severity string     `json:"severity"`
	Trend    QuotaTrend `json:"trend"`
}

// QuotaTrend is the `QuotaTrend` schema.
type QuotaTrend struct {
	// PerDay: Least-squares change in utilisation fraction per day over the last
	// 14 days of snapshots. Null when fewer than 3 readings exist, or when every
	// reading shares an instant. Null means 'not enough history', never 'no
	// risk'.
	PerDay *float64 `json:"perDay"`
	// DaysToExhaustion: Days until used reaches limit at the fitted rate. Null
	// when the trend is flat or falling, when the quota is already at its limit,
	// or when exhaustion lands beyond the 30-day horizon.
	DaysToExhaustion *float64 `json:"daysToExhaustion"`
	// Points: Snapshots the fit used.
	Points int64 `json:"points"`
}

// ReauthenticationRequired is the `ReauthenticationRequired` schema.
type ReauthenticationRequired struct {
	// Error: Human-readable error message
	Error string `json:"error"`
	// Code: One of "reauthentication_required".
	Code string `json:"code"`
}

// RegisteredAgent is the `RegisteredAgent` schema.
type RegisteredAgent struct {
	RegistrationID string `json:"registration_id"`
	// Credential: Bearer credential for this registration. Format
	// `iwa_<base64url>`. Returned once and never recoverable — there is no route
	// that can show it again.
	Credential     string `json:"credential"`
	OrganizationID string `json:"organization_id"`
	// TrialExpiresAt: When the trial workspace is deleted unless a person claims
	// it.
	TrialExpiresAt string `json:"trial_expires_at"`
	ClaimURL       string `json:"claim_url"`
	// Notice: Human-readable summary of the trial terms, meant to be relayed to
	// the user.
	Notice string `json:"notice"`
}

// ReorderRequest is the `ReorderRequest` schema.
type ReorderRequest struct {
	Cards       []ReorderRequestCards `json:"cards,omitempty"`
	ResourceIDs []ResourceID          `json:"resourceIds,omitempty"`
}

// ReportDeliveryTargetOption is the `ReportDeliveryTargetOption` schema.
type ReportDeliveryTargetOption struct {
	// ID: The stored row id — what the schedule input carries.
	ID string `json:"id"`
	// Label: Display label: `#channel` for Slack, the saved label for Teams.
	Label string `json:"label"`
}

// ReportDeliveryTargets is the `ReportDeliveryTargets` schema.
type ReportDeliveryTargets struct {
	SlackChannels []ReportDeliveryTargetOption `json:"slackChannels"`
	TeamsWebhooks []ReportDeliveryTargetOption `json:"teamsWebhooks"`
	// EmailAvailable: Whether this deployment can send mail at all. Addresses
	// can be saved regardless, but they deliver nowhere until a mail provider is
	// configured.
	EmailAvailable bool `json:"emailAvailable"`
}

// ReportNotification is the `ReportNotification` schema.
type ReportNotification struct {
	ID           string `json:"id"`
	CostReportID string `json:"costReportId"`
	// Cadence: How often the schedule fires. The report itself decides what
	// window it charts.
	//
	// One of "daily", "weekly", "monthly".
	Cadence         string   `json:"cadence"`
	SendDay         int64    `json:"sendDay"`
	SendDayOfMonth  int64    `json:"sendDayOfMonth"`
	Hour            int64    `json:"hour"`
	Timezone        string   `json:"timezone"`
	SlackChannelIDs []string `json:"slackChannelIds"`
	TeamsWebhookIDs []string `json:"teamsWebhookIds"`
	EmailRecipients []string `json:"emailRecipients"`
	Enabled         bool     `json:"enabled"`
	// NextSendAt: When the next scheduled send is due; null while disabled.
	NextSendAt *string `json:"nextSendAt"`
	// LastSentAt: When a delivery last actually reached at least one
	// destination.
	LastSentAt *string `json:"lastSentAt"`
	// LastStatus: What the last attempt did. `partial` means some destinations
	// took it and some failed — never retried automatically, because a retry
	// would double-post where it landed.
	//
	// One of "pending", "succeeded", "partial", "failed", "no_targets".
	LastStatus      *string `json:"lastStatus"`
	LastError       *string `json:"lastError"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// ReportNotificationInput: A full replace, like a report's own PUT. At least one
// destination is required — a schedule with nowhere to deliver would only ever
// record failures.
type ReportNotificationInput struct {
	// Cadence: How often the schedule fires. The report itself decides what
	// window it charts.
	//
	// One of "daily", "weekly", "monthly".
	Cadence string `json:"cadence"`
	// SendDay: ISO day of week (1 = Monday … 7 = Sunday); read only when cadence
	// is weekly.
	SendDay *int64 `json:"sendDay,omitempty"`
	// SendDayOfMonth: Day of month; read only when cadence is monthly. A day the
	// month doesn't have clamps to its last day, so 31 means month end
	// everywhere.
	SendDayOfMonth *int64 `json:"sendDayOfMonth,omitempty"`
	// Hour: Local hour in `timezone` the delivery fires at.
	Hour int64 `json:"hour"`
	// Timezone: IANA zone, e.g. `Europe/Berlin`. Validated server-side.
	Timezone string `json:"timezone"`
	// SlackChannelIDs: Stored Slack channel row ids (from the targets endpoint)
	// to post to.
	SlackChannelIDs []string `json:"slackChannelIds"`
	// TeamsWebhookIDs: Stored Teams webhook row ids (from the targets endpoint)
	// to post to.
	TeamsWebhookIDs []string `json:"teamsWebhookIds"`
	// EmailRecipients: Email addresses; normalized (lowercased) server-side. At
	// most 20.
	EmailRecipients []string `json:"emailRecipients"`
	Enabled         bool     `json:"enabled"`
}

// ReportNotificationSendResult is the `ReportNotificationSendResult` schema.
type ReportNotificationSendResult struct {
	Attempted int64                             `json:"attempted"`
	Succeeded int64                             `json:"succeeded"`
	Slack     ReportNotificationSendResultSlack `json:"slack"`
	Teams     ReportNotificationSendResultTeams `json:"teams"`
	Email     ReportNotificationSendResultEmail `json:"email"`
}

// RequiredTag is the `RequiredTag` schema.
type RequiredTag struct {
	Key string `json:"key"`
	// AllowedValues: When set, the tag's value must be one of these (compared
	// exactly).
	AllowedValues []string `json:"allowedValues,omitempty"`
}

// Resource is the `Resource` schema.
type Resource struct {
	ID               ResourceID  `json:"id"`
	PluginID         string      `json:"pluginId"`
	ResourceTypeID   string      `json:"resourceTypeId"`
	AccountID        string      `json:"accountId"`
	DisplayName      string      `json:"displayName"`
	ExternalID       *string     `json:"externalId"`
	FieldsJSON       JSONObject  `json:"fieldsJson"`
	OutputsJSON      JSONObject  `json:"outputsJson"`
	ParentResourceID *ResourceID `json:"parentResourceId"`
}

// ResourceChangeEntry is the `ResourceChangeEntry` schema.
type ResourceChangeEntry struct {
	ID             string     `json:"id"`
	ResourceID     ResourceID `json:"resourceId"`
	AccountID      string     `json:"accountId"`
	PluginID       string     `json:"pluginId"`
	ResourceTypeID string     `json:"resourceTypeId"`
	// DisplayName: Resource display name at the time of the change — survives
	// deletion.
	DisplayName string             `json:"displayName"`
	ChangeKind  ResourceChangeKind `json:"changeKind"`
	// Diff: Changed fields for `updated` events; empty for `created` and
	// `deleted`.
	Diff []ResourceFieldChange `json:"diff"`
	// Origin: Who caused the change when a non-sync writer knows: `schedule` for
	// sleep/wake schedule transitions. Absent/null = observed by sync.
	//
	// One of "schedule".
	Origin    *string `json:"origin,omitempty"`
	CreatedAt string  `json:"createdAt"`
	// RevertedAt: When this event was reverted, or null if it never was.
	// Reverting is a one-shot: an event carrying a timestamp here cannot be
	// reverted again.
	RevertedAt *string `json:"revertedAt,omitempty"`
}

// ResourceChangeFeedEntry is the `ResourceChangeFeedEntry` schema.
type ResourceChangeFeedEntry struct {
	ID             string     `json:"id"`
	ResourceID     ResourceID `json:"resourceId"`
	AccountID      string     `json:"accountId"`
	PluginID       string     `json:"pluginId"`
	ResourceTypeID string     `json:"resourceTypeId"`
	// DisplayName: Resource display name at the time of the change — survives
	// deletion.
	DisplayName string             `json:"displayName"`
	ChangeKind  ResourceChangeKind `json:"changeKind"`
	// Diff: Changed fields for `updated` events; empty for `created` and
	// `deleted`.
	Diff []ResourceFieldChange `json:"diff"`
	// Origin: Who caused the change when a non-sync writer knows: `schedule` for
	// sleep/wake schedule transitions. Absent/null = observed by sync.
	//
	// One of "schedule".
	Origin    *string `json:"origin,omitempty"`
	CreatedAt string  `json:"createdAt"`
	// RevertedAt: When this event was reverted, or null if it never was.
	// Reverting is a one-shot: an event carrying a timestamp here cannot be
	// reverted again.
	RevertedAt  *string `json:"revertedAt,omitempty"`
	AccountName *string `json:"accountName"`
}

// ResourceChangeFeedResponse is the `ResourceChangeFeedResponse` schema.
type ResourceChangeFeedResponse struct {
	Entries []ResourceChangeFeedEntry `json:"entries"`
	Total   int64                     `json:"total"`
}

// ResourceChangeKind: What happened between two consecutive syncs: the resource
// appeared, a stored field changed, or the resource disappeared upstream.
type ResourceChangeKind = string

// The values ResourceChangeKind takes.
const (
	ResourceChangeKindCreated ResourceChangeKind = "created"
	ResourceChangeKindUpdated ResourceChangeKind = "updated"
	ResourceChangeKindDeleted ResourceChangeKind = "deleted"
)

// ResourceChangeListResponse is the `ResourceChangeListResponse` schema.
type ResourceChangeListResponse struct {
	Entries []ResourceChangeEntry `json:"entries"`
}

// ResourceDetail is the `ResourceDetail` schema.
type ResourceDetail struct {
	DetailSchema            JSONObject         `json:"detailSchema"`
	ChildResources          []ChildResourceRef `json:"childResources"`
	ChildTypes              []ChildTypeRef     `json:"childTypes"`
	PluginID                string             `json:"pluginId"`
	PluginLogoSvg           string             `json:"pluginLogoSvg"`
	ResourceID              ResourceID         `json:"resourceId"`
	AccountID               string             `json:"accountId"`
	ResourceTypeID          string             `json:"resourceTypeId"`
	PeerPanes               []PeerPane         `json:"peerPanes"`
	PeerIntegrationStubs    []PeerPaneStub     `json:"peerIntegrationStubs"`
	CanDelete               bool               `json:"canDelete"`
	CanEdit                 bool               `json:"canEdit"`
	EditableFields          []EditableField    `json:"editableFields"`
	CredentialFormats       []CredentialFormat `json:"credentialFormats"`
	SupportsTerraformExport bool               `json:"supportsTerraformExport"`
	HasManifestEditor       bool               `json:"hasManifestEditor"`
	HasSecretVersions       bool               `json:"hasSecretVersions"`
	ResourceDisplayName     string             `json:"resourceDisplayName"`
	ResourceTypeLabel       string             `json:"resourceTypeLabel"`
	ResourceFields          JSONObject         `json:"resourceFields"`
	HasSQLEditor            bool               `json:"hasSqlEditor"`
	HasStorageBrowser       bool               `json:"hasStorageBrowser"`
	HasArtifactRegistry     bool               `json:"hasArtifactRegistry"`
	HasKVBrowser            bool               `json:"hasKvBrowser"`
	HasKVConsole            bool               `json:"hasKvConsole"`
	KVDriverName            *string            `json:"kvDriverName,omitempty"`
	IsMongoDB               bool               `json:"isMongoDb"`
	HasDockerActions        bool               `json:"hasDockerActions"`
	HasSSHTerminal          bool               `json:"hasSshTerminal"`
	HasSFTPBrowser          bool               `json:"hasSftpBrowser"`
	SSHHost                 *string            `json:"sshHost"`
	SSHPrivateHost          *string            `json:"sshPrivateHost,omitempty"`
	DefaultSSHUsername      *string            `json:"defaultSshUsername"`
	ContainerID             string             `json:"containerId"`
	DatabaseName            string             `json:"databaseName"`
	StorageBucketName       string             `json:"storageBucketName"`
	SupportsMetrics         bool               `json:"supportsMetrics"`
	// Schedulable: The type declares lifecycle start/stop actions, so this
	// resource can carry a sleep/wake schedule.
	Schedulable bool `json:"schedulable"`
}

// ResourceFieldChange is the `ResourceFieldChange` schema.
type ResourceFieldChange struct {
	// Field: Top-level field key that changed. Resolved-output keys are prefixed
	// `outputs.`.
	Field string `json:"field"`
	// From: Previous value (null when the field was absent).
	From any `json:"from,omitempty"`
	// To: New value.
	To any `json:"to,omitempty"`
}

// ResourceID: Composite id `pluginId:accountId:externalId`.
//
// Spec schema: `ResourceId`.
type ResourceID = string

// ResourceLease is the `ResourceLease` schema.
type ResourceLease struct {
	ID string `json:"id"`
	// ResourceID: Infrawrench resource id the lease is attached to.
	ResourceID     string   `json:"resourceId"`
	AccountID      string   `json:"accountId"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	// ResourceName: Resource display name (denormalized at lease time, so it
	// survives deletion).
	ResourceName string `json:"resourceName"`
	AccountName  string `json:"accountName"`
	// ExpiresAt: The lease deadline.
	ExpiresAt string `json:"expiresAt"`
	// AutoDelete: Whether the resource is deleted at expiry. Auto-delete is
	// announced twice before it fires and deferred while an org change freeze is
	// in effect.
	AutoDelete bool `json:"autoDelete"`
	// Note: Why/who-for; shown on the expiry radar.
	Note *string `json:"note"`
	// Status: Lease lifecycle: `active` (counting down), `deleted` (auto-delete
	// completed), `failed` (auto-delete was retried and given up on — see
	// `lastError`), or `canceled` (called off; the resource stays).
	//
	// One of "active", "deleted", "failed", "canceled".
	Status string `json:"status"`
	// FirstWarningAt: When the first auto-delete announcement went out; null
	// until sent.
	FirstWarningAt *string `json:"firstWarningAt"`
	// FinalWarningAt: When the final auto-delete announcement went out; null
	// until sent.
	FinalWarningAt *string `json:"finalWarningAt"`
	DeleteAttempts int64   `json:"deleteAttempts"`
	// LastError: Last auto-delete failure or freeze-deferral detail; never
	// silent.
	LastError *string `json:"lastError"`
	// CompletedAt: When the lease reached a terminal status.
	CompletedAt *string `json:"completedAt"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// ResourceLeaseCreate is the `ResourceLeaseCreate` schema.
type ResourceLeaseCreate struct {
	ResourceID string `json:"resourceId"`
	AccountID  string `json:"accountId"`
	// ExpiresAt: Must be in the future, at most 365 days out.
	ExpiresAt string `json:"expiresAt"`
	// AutoDelete: Requires the `resources:delete` permission when true.
	AutoDelete *bool   `json:"autoDelete,omitempty"`
	Note       *string `json:"note,omitempty"`
}

// ResourceLeaseList is the `ResourceLeaseList` schema.
type ResourceLeaseList struct {
	Leases []ResourceLease `json:"leases"`
}

// ResourceLeaseLookup is the `ResourceLeaseLookup` schema.
type ResourceLeaseLookup struct {
	Lease any `json:"lease"`
}

// ResourceLeaseUpdate is the `ResourceLeaseUpdate` schema.
type ResourceLeaseUpdate struct {
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// AutoDelete: Requires the `resources:delete` permission when set to true.
	AutoDelete *bool `json:"autoDelete,omitempty"`
	// Note: `null` clears the note.
	Note *string `json:"note,omitempty"`
}

// ResourceOwnerAnnotation: Who owns this resource, or null when nobody has
// claimed it. Present only when the owner can be named: a resource carrying a
// purpose but no owner reads as null, because the question this answers is who
// to tell.
//
// The API may send null in its place.
type ResourceOwnerAnnotation struct {
	// UserID: Set when a routable org member owns it.
	UserID *string `json:"userId"`
	// DisplayName: The member's name, or the free-text owner.
	DisplayName string `json:"displayName"`
	// IsLabel: True when the owner is free text — nothing can be routed to it.
	IsLabel   bool    `json:"isLabel"`
	TicketURL *string `json:"ticketUrl"`
	Purpose   *string `json:"purpose"`
}

// ResourceOwnership is the `ResourceOwnership` schema.
type ResourceOwnership struct {
	ID             string   `json:"id"`
	ResourceID     string   `json:"resourceId"`
	AccountID      string   `json:"accountId"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	// ResourceName: Resource display name, denormalized so a report can name a
	// deleted resource.
	ResourceName string `json:"resourceName"`
	// OwnerUserID: The routable owner — an org member. Alerts about this
	// resource reach them.
	OwnerUserID *string `json:"ownerUserId"`
	// OwnerName: Resolved server-side; null when unset or removed.
	OwnerName  *string `json:"ownerName"`
	OwnerEmail *string `json:"ownerEmail"`
	// OwnerLabel: Free-text owner (a team, a rota, a contractor). Display-only,
	// never routed.
	OwnerLabel *string `json:"ownerLabel"`
	// Purpose: What this resource is for.
	Purpose *string `json:"purpose"`
	// TicketURL: Link to the ticket that authorized it.
	TicketURL *string `json:"ticketUrl"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

// ResourceOwnershipEnvelope is the `ResourceOwnershipEnvelope` schema.
type ResourceOwnershipEnvelope struct {
	Ownership any `json:"ownership"`
}

// ResourceOwnershipListResponse is the `ResourceOwnershipListResponse` schema.
type ResourceOwnershipListResponse struct {
	Ownership []ResourceOwnership `json:"ownership"`
}

// ResourceOwnershipPatch is the `ResourceOwnershipPatch` schema.
type ResourceOwnershipPatch struct {
	ResourceID string `json:"resourceId"`
	// OwnerUserID: Omit to keep, null to clear.
	OwnerUserID *string `json:"ownerUserId,omitempty"`
	// OwnerLabel: Omit to keep, null to clear.
	OwnerLabel *string `json:"ownerLabel,omitempty"`
	// Purpose: Omit to keep, null to clear.
	Purpose *string `json:"purpose,omitempty"`
	// TicketURL: Omit to keep, null to clear.
	TicketURL *string `json:"ticketUrl,omitempty"`
}

// ResourceStatus: Normalized status reported by a plugin's
// renderSidebarItem/renderDetail.
type ResourceStatus = string

// The values ResourceStatus takes.
const (
	ResourceStatusHealthy      ResourceStatus = "healthy"
	ResourceStatusDegraded     ResourceStatus = "degraded"
	ResourceStatusError        ResourceStatus = "error"
	ResourceStatusUnknown      ResourceStatus = "unknown"
	ResourceStatusProvisioning ResourceStatus = "provisioning"
	ResourceStatusInfo         ResourceStatus = "info"
)

// ResourceTypeID: Resource type id. Note: not every plugin exposes every type —
// see the plugin's `resourceTypes` for the valid (pluginId, typeId) pairs.
//
// Spec schema: `ResourceTypeId`.
type ResourceTypeID = string

// The values ResourceTypeID takes.
const (
	ResourceTypeIDAccessApplication              ResourceTypeID = "access-application"
	ResourceTypeIDAccessPolicy                   ResourceTypeID = "access-policy"
	ResourceTypeIDAccount                        ResourceTypeID = "account"
	ResourceTypeIDAcmCertificate                 ResourceTypeID = "acm-certificate"
	ResourceTypeIDAgentAPIKey                    ResourceTypeID = "agent-api-key"
	ResourceTypeIDAiGateway                      ResourceTypeID = "ai-gateway"
	ResourceTypeIDAiSearch                       ResourceTypeID = "ai-search"
	ResourceTypeIDAlb                            ResourceTypeID = "alb"
	ResourceTypeIDAlertPolicy                    ResourceTypeID = "alert-policy"
	ResourceTypeIDAlloydbCluster                 ResourceTypeID = "alloydb-cluster"
	ResourceTypeIDAlloydbInstance                ResourceTypeID = "alloydb-instance"
	ResourceTypeIDAPIGateway                     ResourceTypeID = "api-gateway"
	ResourceTypeIDAPIKey                         ResourceTypeID = "api-key"
	ResourceTypeIDApp                            ResourceTypeID = "app"
	ResourceTypeIDAppEngineService               ResourceTypeID = "app-engine-service"
	ResourceTypeIDApprunnerService               ResourceTypeID = "apprunner-service"
	ResourceTypeIDArtifactRegistryRepo           ResourceTypeID = "artifact-registry-repo"
	ResourceTypeIDAuditEvent                     ResourceTypeID = "audit-event"
	ResourceTypeIDAutoScalingGroup               ResourceTypeID = "auto-scaling-group"
	ResourceTypeIDAzureAksCluster                ResourceTypeID = "azure-aks-cluster"
	ResourceTypeIDAzureAppGateway                ResourceTypeID = "azure-app-gateway"
	ResourceTypeIDAzureAppRegistration           ResourceTypeID = "azure-app-registration"
	ResourceTypeIDAzureAppService                ResourceTypeID = "azure-app-service"
	ResourceTypeIDAzureAppServicePlan            ResourceTypeID = "azure-app-service-plan"
	ResourceTypeIDAzureContainerInstance         ResourceTypeID = "azure-container-instance"
	ResourceTypeIDAzureContainerRegistry         ResourceTypeID = "azure-container-registry"
	ResourceTypeIDAzureCosmosDB                  ResourceTypeID = "azure-cosmos-db"
	ResourceTypeIDAzureDisk                      ResourceTypeID = "azure-disk"
	ResourceTypeIDAzureDNSZone                   ResourceTypeID = "azure-dns-zone"
	ResourceTypeIDAzureEventHub                  ResourceTypeID = "azure-event-hub"
	ResourceTypeIDAzureFirewall                  ResourceTypeID = "azure-firewall"
	ResourceTypeIDAzureFunctionApp               ResourceTypeID = "azure-function-app"
	ResourceTypeIDAzureKeyVault                  ResourceTypeID = "azure-key-vault"
	ResourceTypeIDAzureLoadBalancer              ResourceTypeID = "azure-load-balancer"
	ResourceTypeIDAzureLogAnalytics              ResourceTypeID = "azure-log-analytics"
	ResourceTypeIDAzureManagedIdentity           ResourceTypeID = "azure-managed-identity"
	ResourceTypeIDAzureMysqlFlexible             ResourceTypeID = "azure-mysql-flexible"
	ResourceTypeIDAzureNatGateway                ResourceTypeID = "azure-nat-gateway"
	ResourceTypeIDAzureNsg                       ResourceTypeID = "azure-nsg"
	ResourceTypeIDAzurePostgresFlexible          ResourceTypeID = "azure-postgres-flexible"
	ResourceTypeIDAzurePrivateDNSZone            ResourceTypeID = "azure-private-dns-zone"
	ResourceTypeIDAzurePublicIP                  ResourceTypeID = "azure-public-ip"
	ResourceTypeIDAzureRedisCache                ResourceTypeID = "azure-redis-cache"
	ResourceTypeIDAzureResourceGroup             ResourceTypeID = "azure-resource-group"
	ResourceTypeIDAzureRouteTable                ResourceTypeID = "azure-route-table"
	ResourceTypeIDAzureServiceBus                ResourceTypeID = "azure-service-bus"
	ResourceTypeIDAzureSQLDatabase               ResourceTypeID = "azure-sql-database"
	ResourceTypeIDAzureStorageAccount            ResourceTypeID = "azure-storage-account"
	ResourceTypeIDAzureSubnet                    ResourceTypeID = "azure-subnet"
	ResourceTypeIDAzureVM                        ResourceTypeID = "azure-vm"
	ResourceTypeIDAzureVnet                      ResourceTypeID = "azure-vnet"
	ResourceTypeIDBackendService                 ResourceTypeID = "backend-service"
	ResourceTypeIDBackupVault                    ResourceTypeID = "backup-vault"
	ResourceTypeIDBalance                        ResourceTypeID = "balance"
	ResourceTypeIDBatch                          ResourceTypeID = "batch"
	ResourceTypeIDBatchInferenceJob              ResourceTypeID = "batch-inference-job"
	ResourceTypeIDBatchJobQueue                  ResourceTypeID = "batch-job-queue"
	ResourceTypeIDBedrockModel                   ResourceTypeID = "bedrock-model"
	ResourceTypeIDBigqueryDataset                ResourceTypeID = "bigquery-dataset"
	ResourceTypeIDBigqueryTable                  ResourceTypeID = "bigquery-table"
	ResourceTypeIDBigtableInstance               ResourceTypeID = "bigtable-instance"
	ResourceTypeIDBlockVolume                    ResourceTypeID = "block-volume"
	ResourceTypeIDCacheRule                      ResourceTypeID = "cache-rule"
	ResourceTypeIDCachedContent                  ResourceTypeID = "cached-content"
	ResourceTypeIDCertificate                    ResourceTypeID = "certificate"
	ResourceTypeIDChDatabase                     ResourceTypeID = "ch-database"
	ResourceTypeIDChService                      ResourceTypeID = "ch-service"
	ResourceTypeIDCloudArmorPolicy               ResourceTypeID = "cloud-armor-policy"
	ResourceTypeIDCloudBuildTrigger              ResourceTypeID = "cloud-build-trigger"
	ResourceTypeIDCloudDeployPipeline            ResourceTypeID = "cloud-deploy-pipeline"
	ResourceTypeIDCloudDNSRecordSet              ResourceTypeID = "cloud-dns-record-set"
	ResourceTypeIDCloudDNSZone                   ResourceTypeID = "cloud-dns-zone"
	ResourceTypeIDCloudFunction                  ResourceTypeID = "cloud-function"
	ResourceTypeIDCloudNat                       ResourceTypeID = "cloud-nat"
	ResourceTypeIDCloudRouter                    ResourceTypeID = "cloud-router"
	ResourceTypeIDCloudRunService                ResourceTypeID = "cloud-run-service"
	ResourceTypeIDCloudSchedulerJob              ResourceTypeID = "cloud-scheduler-job"
	ResourceTypeIDCloudTasksQueue                ResourceTypeID = "cloud-tasks-queue"
	ResourceTypeIDCloudformationStack            ResourceTypeID = "cloudformation-stack"
	ResourceTypeIDCloudfrontDistribution         ResourceTypeID = "cloudfront-distribution"
	ResourceTypeIDCloudsqlInstance               ResourceTypeID = "cloudsql-instance"
	ResourceTypeIDCloudtrailTrail                ResourceTypeID = "cloudtrail-trail"
	ResourceTypeIDCloudwatchAlarm                ResourceTypeID = "cloudwatch-alarm"
	ResourceTypeIDCloudwatchLogGroup             ResourceTypeID = "cloudwatch-log-group"
	ResourceTypeIDCodebuildProject               ResourceTypeID = "codebuild-project"
	ResourceTypeIDCodepipelinePipeline           ResourceTypeID = "codepipeline-pipeline"
	ResourceTypeIDCognitoUserPool                ResourceTypeID = "cognito-user-pool"
	ResourceTypeIDCollection                     ResourceTypeID = "collection"
	ResourceTypeIDComposerEnvironment            ResourceTypeID = "composer-environment"
	ResourceTypeIDConnection                     ResourceTypeID = "connection"
	ResourceTypeIDContainer                      ResourceTypeID = "container"
	ResourceTypeIDContainerRegistry              ResourceTypeID = "container-registry"
	ResourceTypeIDCustomHostname                 ResourceTypeID = "custom-hostname"
	ResourceTypeIDCustomVoice                    ResourceTypeID = "custom-voice"
	ResourceTypeIDD1Database                     ResourceTypeID = "d1-database"
	ResourceTypeIDDatabricksApp                  ResourceTypeID = "databricks-app"
	ResourceTypeIDDatabricksCatalog              ResourceTypeID = "databricks-catalog"
	ResourceTypeIDDatabricksCluster              ResourceTypeID = "databricks-cluster"
	ResourceTypeIDDatabricksClusterPolicy        ResourceTypeID = "databricks-cluster-policy"
	ResourceTypeIDDatabricksDashboard            ResourceTypeID = "databricks-dashboard"
	ResourceTypeIDDatabricksFunction             ResourceTypeID = "databricks-function"
	ResourceTypeIDDatabricksJob                  ResourceTypeID = "databricks-job"
	ResourceTypeIDDatabricksModelVersion         ResourceTypeID = "databricks-model-version"
	ResourceTypeIDDatabricksNodeType             ResourceTypeID = "databricks-node-type"
	ResourceTypeIDDatabricksPipeline             ResourceTypeID = "databricks-pipeline"
	ResourceTypeIDDatabricksRegisteredModel      ResourceTypeID = "databricks-registered-model"
	ResourceTypeIDDatabricksRepo                 ResourceTypeID = "databricks-repo"
	ResourceTypeIDDatabricksSchema               ResourceTypeID = "databricks-schema"
	ResourceTypeIDDatabricksSecretScope          ResourceTypeID = "databricks-secret-scope"
	ResourceTypeIDDatabricksServingEndpoint      ResourceTypeID = "databricks-serving-endpoint"
	ResourceTypeIDDatabricksSQLQuery             ResourceTypeID = "databricks-sql-query"
	ResourceTypeIDDatabricksSQLWarehouse         ResourceTypeID = "databricks-sql-warehouse"
	ResourceTypeIDDatabricksTable                ResourceTypeID = "databricks-table"
	ResourceTypeIDDatabricksVectorSearchEndpoint ResourceTypeID = "databricks-vector-search-endpoint"
	ResourceTypeIDDatabricksVectorSearchIndex    ResourceTypeID = "databricks-vector-search-index"
	ResourceTypeIDDatabricksVolume               ResourceTypeID = "databricks-volume"
	ResourceTypeIDDatabricksWorkspaceObject      ResourceTypeID = "databricks-workspace-object"
	ResourceTypeIDDataflowJob                    ResourceTypeID = "dataflow-job"
	ResourceTypeIDDataset                        ResourceTypeID = "dataset"
	ResourceTypeIDDBSubnetGroup                  ResourceTypeID = "db-subnet-group"
	ResourceTypeIDDBUser                         ResourceTypeID = "db-user"
	ResourceTypeIDDedicatedInference             ResourceTypeID = "dedicated-inference"
	ResourceTypeIDDeployedModel                  ResourceTypeID = "deployed-model"
	ResourceTypeIDDeployment                     ResourceTypeID = "deployment"
	ResourceTypeIDDirectory                      ResourceTypeID = "directory"
	ResourceTypeIDDirectoryGroup                 ResourceTypeID = "directory-group"
	ResourceTypeIDDirectoryUser                  ResourceTypeID = "directory-user"
	ResourceTypeIDDNSRecord                      ResourceTypeID = "dns-record"
	ResourceTypeIDDockerContainer                ResourceTypeID = "docker-container"
	ResourceTypeIDDockerImage                    ResourceTypeID = "docker-image"
	ResourceTypeIDDockerNetwork                  ResourceTypeID = "docker-network"
	ResourceTypeIDDockerVolume                   ResourceTypeID = "docker-volume"
	ResourceTypeIDDocumentdbCluster              ResourceTypeID = "documentdb-cluster"
	ResourceTypeIDDoksCluster                    ResourceTypeID = "doks-cluster"
	ResourceTypeIDDomain                         ResourceTypeID = "domain"
	ResourceTypeIDDroplet                        ResourceTypeID = "droplet"
	ResourceTypeIDDurableObjectNamespace         ResourceTypeID = "durable-object-namespace"
	ResourceTypeIDDynamodbTable                  ResourceTypeID = "dynamodb-table"
	ResourceTypeIDEbsVolume                      ResourceTypeID = "ebs-volume"
	ResourceTypeIDEc2Instance                    ResourceTypeID = "ec2-instance"
	ResourceTypeIDEcrRepository                  ResourceTypeID = "ecr-repository"
	ResourceTypeIDEcsService                     ResourceTypeID = "ecs-service"
	ResourceTypeIDEfsFileSystem                  ResourceTypeID = "efs-file-system"
	ResourceTypeIDEksCluster                     ResourceTypeID = "eks-cluster"
	ResourceTypeIDElasticIP                      ResourceTypeID = "elastic-ip"
	ResourceTypeIDElasticacheCluster             ResourceTypeID = "elasticache-cluster"
	ResourceTypeIDEmailRoutingRule               ResourceTypeID = "email-routing-rule"
	ResourceTypeIDEmbedJob                       ResourceTypeID = "embed-job"
	ResourceTypeIDEndpoint                       ResourceTypeID = "endpoint"
	ResourceTypeIDEval                           ResourceTypeID = "eval"
	ResourceTypeIDEvaluation                     ResourceTypeID = "evaluation"
	ResourceTypeIDEventbridgeRule                ResourceTypeID = "eventbridge-rule"
	ResourceTypeIDFile                           ResourceTypeID = "file"
	ResourceTypeIDFileSearchDocument             ResourceTypeID = "file-search-document"
	ResourceTypeIDFileSearchStore                ResourceTypeID = "file-search-store"
	ResourceTypeIDFineTune                       ResourceTypeID = "fine-tune"
	ResourceTypeIDFineTuningJob                  ResourceTypeID = "fine-tuning-job"
	ResourceTypeIDFinetunedModel                 ResourceTypeID = "finetuned-model"
	ResourceTypeIDFirestoreDatabase              ResourceTypeID = "firestore-database"
	ResourceTypeIDFirewall                       ResourceTypeID = "firewall"
	ResourceTypeIDFirewallRule                   ResourceTypeID = "firewall-rule"
	ResourceTypeIDFloatingIP                     ResourceTypeID = "floating-ip"
	ResourceTypeIDFolder                         ResourceTypeID = "folder"
	ResourceTypeIDForwardingRule                 ResourceTypeID = "forwarding-rule"
	ResourceTypeIDGateway                        ResourceTypeID = "gateway"
	ResourceTypeIDGceDisk                        ResourceTypeID = "gce-disk"
	ResourceTypeIDGceInstance                    ResourceTypeID = "gce-instance"
	ResourceTypeIDGCPProject                     ResourceTypeID = "gcp-project"
	ResourceTypeIDGCPServiceAccount              ResourceTypeID = "gcp-service-account"
	ResourceTypeIDGcsBucket                      ResourceTypeID = "gcs-bucket"
	ResourceTypeIDGenAiAgent                     ResourceTypeID = "gen-ai-agent"
	ResourceTypeIDGenAiKnowledgeBase             ResourceTypeID = "gen-ai-knowledge-base"
	ResourceTypeIDGenAiModelRouter               ResourceTypeID = "gen-ai-model-router"
	ResourceTypeIDGkeCluster                     ResourceTypeID = "gke-cluster"
	ResourceTypeIDGlueDatabase                   ResourceTypeID = "glue-database"
	ResourceTypeIDGroqBatch                      ResourceTypeID = "groq-batch"
	ResourceTypeIDGroqFile                       ResourceTypeID = "groq-file"
	ResourceTypeIDGroqFineTuning                 ResourceTypeID = "groq-fine-tuning"
	ResourceTypeIDGroqModel                      ResourceTypeID = "groq-model"
	ResourceTypeIDHardware                       ResourceTypeID = "hardware"
	ResourceTypeIDHealthCheck                    ResourceTypeID = "health-check"
	ResourceTypeIDHealthcheck                    ResourceTypeID = "healthcheck"
	ResourceTypeIDHistoryItem                    ResourceTypeID = "history-item"
	ResourceTypeIDHyperdrive                     ResourceTypeID = "hyperdrive"
	ResourceTypeIDIamRole                        ResourceTypeID = "iam-role"
	ResourceTypeIDIamUser                        ResourceTypeID = "iam-user"
	ResourceTypeIDImage                          ResourceTypeID = "image"
	ResourceTypeIDInferenceBatch                 ResourceTypeID = "inference-batch"
	ResourceTypeIDInstance                       ResourceTypeID = "instance"
	ResourceTypeIDInstanceGroup                  ResourceTypeID = "instance-group"
	ResourceTypeIDInstanceTemplate               ResourceTypeID = "instance-template"
	ResourceTypeIDInternetGateway                ResourceTypeID = "internet-gateway"
	ResourceTypeIDInvitation                     ResourceTypeID = "invitation"
	ResourceTypeIDInvite                         ResourceTypeID = "invite"
	ResourceTypeIDIPAccessRule                   ResourceTypeID = "ip-access-rule"
	ResourceTypeIDIPAllocation                   ResourceTypeID = "ip-allocation"
	ResourceTypeIDJob                            ResourceTypeID = "job"
	ResourceTypeIDK8sCluster                     ResourceTypeID = "k8s-cluster"
	ResourceTypeIDK8sConfigmap                   ResourceTypeID = "k8s-configmap"
	ResourceTypeIDK8sCronjob                     ResourceTypeID = "k8s-cronjob"
	ResourceTypeIDK8sDaemonset                   ResourceTypeID = "k8s-daemonset"
	ResourceTypeIDK8sDeployment                  ResourceTypeID = "k8s-deployment"
	ResourceTypeIDK8sIngress                     ResourceTypeID = "k8s-ingress"
	ResourceTypeIDK8sJob                         ResourceTypeID = "k8s-job"
	ResourceTypeIDK8sNamespace                   ResourceTypeID = "k8s-namespace"
	ResourceTypeIDK8sNode                        ResourceTypeID = "k8s-node"
	ResourceTypeIDK8sPod                         ResourceTypeID = "k8s-pod"
	ResourceTypeIDK8sSecret                      ResourceTypeID = "k8s-secret"
	ResourceTypeIDK8sService                     ResourceTypeID = "k8s-service"
	ResourceTypeIDK8sStatefulset                 ResourceTypeID = "k8s-statefulset"
	ResourceTypeIDKafkaCluster                   ResourceTypeID = "kafka-cluster"
	ResourceTypeIDKafkaConsumerGroup             ResourceTypeID = "kafka-consumer-group"
	ResourceTypeIDKafkaTopic                     ResourceTypeID = "kafka-topic"
	ResourceTypeIDKapsuleCluster                 ResourceTypeID = "kapsule-cluster"
	ResourceTypeIDKinesisStream                  ResourceTypeID = "kinesis-stream"
	ResourceTypeIDKmsKey                         ResourceTypeID = "kms-key"
	ResourceTypeIDKmsKeyRing                     ResourceTypeID = "kms-key-ring"
	ResourceTypeIDKVNamespace                    ResourceTypeID = "kv-namespace"
	ResourceTypeIDLambdaFunction                 ResourceTypeID = "lambda-function"
	ResourceTypeIDLoadBalancer                   ResourceTypeID = "load-balancer"
	ResourceTypeIDLogSink                        ResourceTypeID = "log-sink"
	ResourceTypeIDLogpushJob                     ResourceTypeID = "logpush-job"
	ResourceTypeIDMachine                        ResourceTypeID = "machine"
	ResourceTypeIDManagedDatabase                ResourceTypeID = "managed-database"
	ResourceTypeIDManagedDB                      ResourceTypeID = "managed-db"
	ResourceTypeIDManagedEndpoint                ResourceTypeID = "managed-endpoint"
	ResourceTypeIDManagedKube                    ResourceTypeID = "managed-kube"
	ResourceTypeIDMediaAsset                     ResourceTypeID = "media-asset"
	ResourceTypeIDMember                         ResourceTypeID = "member"
	ResourceTypeIDMemcachedInstance              ResourceTypeID = "memcached-instance"
	ResourceTypeIDMemorystoreMemcached           ResourceTypeID = "memorystore-memcached"
	ResourceTypeIDMemorystoreRedis               ResourceTypeID = "memorystore-redis"
	ResourceTypeIDMessageBatch                   ResourceTypeID = "message-batch"
	ResourceTypeIDMistralAPIKey                  ResourceTypeID = "mistral-api-key"
	ResourceTypeIDMistralBatchJob                ResourceTypeID = "mistral-batch-job"
	ResourceTypeIDMistralFile                    ResourceTypeID = "mistral-file"
	ResourceTypeIDMistralFineTuningJob           ResourceTypeID = "mistral-fine-tuning-job"
	ResourceTypeIDMistralModel                   ResourceTypeID = "mistral-model"
	ResourceTypeIDMistralVoice                   ResourceTypeID = "mistral-voice"
	ResourceTypeIDModel                          ResourceTypeID = "model"
	ResourceTypeIDModelAPIKey                    ResourceTypeID = "model-api-key"
	ResourceTypeIDModelEndpoint                  ResourceTypeID = "model-endpoint"
	ResourceTypeIDMongodbDatabase                ResourceTypeID = "mongodb-database"
	ResourceTypeIDMqBroker                       ResourceTypeID = "mq-broker"
	ResourceTypeIDMskCluster                     ResourceTypeID = "msk-cluster"
	ResourceTypeIDMssqlDatabase                  ResourceTypeID = "mssql-database"
	ResourceTypeIDMysqlDatabase                  ResourceTypeID = "mysql-database"
	ResourceTypeIDNatGateway                     ResourceTypeID = "nat-gateway"
	ResourceTypeIDNeonAiGateway                  ResourceTypeID = "neon-ai-gateway"
	ResourceTypeIDNeonAuth                       ResourceTypeID = "neon-auth"
	ResourceTypeIDNeonAuthDomain                 ResourceTypeID = "neon-auth-domain"
	ResourceTypeIDNeonAuthOAuthProvider          ResourceTypeID = "neon-auth-oauth-provider"
	ResourceTypeIDNeonBranch                     ResourceTypeID = "neon-branch"
	ResourceTypeIDNeonBucket                     ResourceTypeID = "neon-bucket"
	ResourceTypeIDNeonCredential                 ResourceTypeID = "neon-credential"
	ResourceTypeIDNeonDataAPI                    ResourceTypeID = "neon-data-api"
	ResourceTypeIDNeonDatabase                   ResourceTypeID = "neon-database"
	ResourceTypeIDNeonEndpoint                   ResourceTypeID = "neon-endpoint"
	ResourceTypeIDNeonFunction                   ResourceTypeID = "neon-function"
	ResourceTypeIDNeonProject                    ResourceTypeID = "neon-project"
	ResourceTypeIDNeonRole                       ResourceTypeID = "neon-role"
	ResourceTypeIDNeonSnapshot                   ResourceTypeID = "neon-snapshot"
	ResourceTypeIDNeptuneCluster                 ResourceTypeID = "neptune-cluster"
	ResourceTypeIDNetlifyBuildHook               ResourceTypeID = "netlify-build-hook"
	ResourceTypeIDNetlifyDeploy                  ResourceTypeID = "netlify-deploy"
	ResourceTypeIDNetlifyDNSRecord               ResourceTypeID = "netlify-dns-record"
	ResourceTypeIDNetlifyDNSZone                 ResourceTypeID = "netlify-dns-zone"
	ResourceTypeIDNetlifyEnvVar                  ResourceTypeID = "netlify-env-var"
	ResourceTypeIDNetlifyForm                    ResourceTypeID = "netlify-form"
	ResourceTypeIDNetlifySite                    ResourceTypeID = "netlify-site"
	ResourceTypeIDNetwork                        ResourceTypeID = "network"
	ResourceTypeIDNfsShare                       ResourceTypeID = "nfs-share"
	ResourceTypeIDNotificationPolicy             ResourceTypeID = "notification-policy"
	ResourceTypeIDObjectStorageBucket            ResourceTypeID = "object-storage-bucket"
	ResourceTypeIDOpensearchCluster              ResourceTypeID = "opensearch-cluster"
	ResourceTypeIDOpensearchDomain               ResourceTypeID = "opensearch-domain"
	ResourceTypeIDOrganization                   ResourceTypeID = "organization"
	ResourceTypeIDOrganizationMembership         ResourceTypeID = "organization-membership"
	ResourceTypeIDOrganizationUser               ResourceTypeID = "organization-user"
	ResourceTypeIDPageRule                       ResourceTypeID = "page-rule"
	ResourceTypeIDPgDatabase                     ResourceTypeID = "pg-database"
	ResourceTypeIDPgSchema                       ResourceTypeID = "pg-schema"
	ResourceTypeIDPlacementGroup                 ResourceTypeID = "placement-group"
	ResourceTypeIDPrediction                     ResourceTypeID = "prediction"
	ResourceTypeIDPrimaryIP                      ResourceTypeID = "primary-ip"
	ResourceTypeIDPrivateNetwork                 ResourceTypeID = "private-network"
	ResourceTypeIDProject                        ResourceTypeID = "project"
	ResourceTypeIDProjectAPIKey                  ResourceTypeID = "project-api-key"
	ResourceTypeIDPronunciationDict              ResourceTypeID = "pronunciation-dict"
	ResourceTypeIDPronunciationDictionary        ResourceTypeID = "pronunciation-dictionary"
	ResourceTypeIDProvider                       ResourceTypeID = "provider"
	ResourceTypeIDPsBackup                       ResourceTypeID = "ps-backup"
	ResourceTypeIDPsBranch                       ResourceTypeID = "ps-branch"
	ResourceTypeIDPsDatabase                     ResourceTypeID = "ps-database"
	ResourceTypeIDPsDeployRequest                ResourceTypeID = "ps-deploy-request"
	ResourceTypeIDPsPassword                     ResourceTypeID = "ps-password"
	ResourceTypeIDPubsubSubscription             ResourceTypeID = "pubsub-subscription"
	ResourceTypeIDPubsubTopic                    ResourceTypeID = "pubsub-topic"
	ResourceTypeIDQueue                          ResourceTypeID = "queue"
	ResourceTypeIDQuota                          ResourceTypeID = "quota"
	ResourceTypeIDR2Bucket                       ResourceTypeID = "r2-bucket"
	ResourceTypeIDRateLimitRule                  ResourceTypeID = "rate-limit-rule"
	ResourceTypeIDRdbInstance                    ResourceTypeID = "rdb-instance"
	ResourceTypeIDRdsCluster                     ResourceTypeID = "rds-cluster"
	ResourceTypeIDRdsInstance                    ResourceTypeID = "rds-instance"
	ResourceTypeIDRedirectRule                   ResourceTypeID = "redirect-rule"
	ResourceTypeIDRedisInstance                  ResourceTypeID = "redis-instance"
	ResourceTypeIDRedshiftCluster                ResourceTypeID = "redshift-cluster"
	ResourceTypeIDReservedIP                     ResourceTypeID = "reserved-ip"
	ResourceTypeIDRole                           ResourceTypeID = "role"
	ResourceTypeIDRouteTable                     ResourceTypeID = "route-table"
	ResourceTypeIDRoute53HealthCheck             ResourceTypeID = "route53-health-check"
	ResourceTypeIDRoute53HostedZone              ResourceTypeID = "route53-hosted-zone"
	ResourceTypeIDRoute53RecordSet               ResourceTypeID = "route53-record-set"
	ResourceTypeIDS3Bucket                       ResourceTypeID = "s3-bucket"
	ResourceTypeIDSagemakerEndpoint              ResourceTypeID = "sagemaker-endpoint"
	ResourceTypeIDSecret                         ResourceTypeID = "secret"
	ResourceTypeIDSecretManagerSecret            ResourceTypeID = "secret-manager-secret"
	ResourceTypeIDSecretsManagerSecret           ResourceTypeID = "secrets-manager-secret"
	ResourceTypeIDSecurityGroup                  ResourceTypeID = "security-group"
	ResourceTypeIDServer                         ResourceTypeID = "server"
	ResourceTypeIDSnapshot                       ResourceTypeID = "snapshot"
	ResourceTypeIDSnsTopic                       ResourceTypeID = "sns-topic"
	ResourceTypeIDSpacesBucket                   ResourceTypeID = "spaces-bucket"
	ResourceTypeIDSpannerBackup                  ResourceTypeID = "spanner-backup"
	ResourceTypeIDSpannerDatabase                ResourceTypeID = "spanner-database"
	ResourceTypeIDSpannerInstance                ResourceTypeID = "spanner-instance"
	ResourceTypeIDSpectrumApplication            ResourceTypeID = "spectrum-application"
	ResourceTypeIDSqsQueue                       ResourceTypeID = "sqs-queue"
	ResourceTypeIDSSHKey                         ResourceTypeID = "ssh-key"
	ResourceTypeIDSSHTarget                      ResourceTypeID = "ssh-target"
	ResourceTypeIDSSLCertificate                 ResourceTypeID = "ssl-certificate"
	ResourceTypeIDSsmParameter                   ResourceTypeID = "ssm-parameter"
	ResourceTypeIDStaticIP                       ResourceTypeID = "static-ip"
	ResourceTypeIDStepFunction                   ResourceTypeID = "step-function"
	ResourceTypeIDSubnet                         ResourceTypeID = "subnet"
	ResourceTypeIDSupervisedFineTuningJob        ResourceTypeID = "supervised-fine-tuning-job"
	ResourceTypeIDTargetGroup                    ResourceTypeID = "target-group"
	ResourceTypeIDTraining                       ResourceTypeID = "training"
	ResourceTypeIDTranscript                     ResourceTypeID = "transcript"
	ResourceTypeIDTranscription                  ResourceTypeID = "transcription"
	ResourceTypeIDTransformation                 ResourceTypeID = "transformation"
	ResourceTypeIDTunedModel                     ResourceTypeID = "tuned-model"
	ResourceTypeIDTunnel                         ResourceTypeID = "tunnel"
	ResourceTypeIDTurnstileWidget                ResourceTypeID = "turnstile-widget"
	ResourceTypeIDTursoAPIToken                  ResourceTypeID = "turso-api-token"
	ResourceTypeIDTursoDatabase                  ResourceTypeID = "turso-database"
	ResourceTypeIDTursoDatabaseInstance          ResourceTypeID = "turso-database-instance"
	ResourceTypeIDTursoGroup                     ResourceTypeID = "turso-group"
	ResourceTypeIDTursoLocation                  ResourceTypeID = "turso-location"
	ResourceTypeIDTursoOrganizationInvite        ResourceTypeID = "turso-organization-invite"
	ResourceTypeIDTursoOrganizationMember        ResourceTypeID = "turso-organization-member"
	ResourceTypeIDUploadPreset                   ResourceTypeID = "upload-preset"
	ResourceTypeIDUser                           ResourceTypeID = "user"
	ResourceTypeIDUtApp                          ResourceTypeID = "ut-app"
	ResourceTypeIDUtFile                         ResourceTypeID = "ut-file"
	ResourceTypeIDVectorStore                    ResourceTypeID = "vector-store"
	ResourceTypeIDVectorizeIndex                 ResourceTypeID = "vectorize-index"
	ResourceTypeIDVercelDeployment               ResourceTypeID = "vercel-deployment"
	ResourceTypeIDVercelDomain                   ResourceTypeID = "vercel-domain"
	ResourceTypeIDVercelEnvVar                   ResourceTypeID = "vercel-env-var"
	ResourceTypeIDVercelProject                  ResourceTypeID = "vercel-project"
	ResourceTypeIDVercelTeam                     ResourceTypeID = "vercel-team"
	ResourceTypeIDVertexAiEndpoint               ResourceTypeID = "vertex-ai-endpoint"
	ResourceTypeIDVertexGeminiModel              ResourceTypeID = "vertex-gemini-model"
	ResourceTypeIDVocabulary                     ResourceTypeID = "vocabulary"
	ResourceTypeIDVoice                          ResourceTypeID = "voice"
	ResourceTypeIDVolume                         ResourceTypeID = "volume"
	ResourceTypeIDVpc                            ResourceTypeID = "vpc"
	ResourceTypeIDVpcNetwork                     ResourceTypeID = "vpc-network"
	ResourceTypeIDWafWebACL                      ResourceTypeID = "waf-web-acl"
	ResourceTypeIDWaitingRoom                    ResourceTypeID = "waiting-room"
	ResourceTypeIDWebhookEndpoint                ResourceTypeID = "webhook-endpoint"
	ResourceTypeIDWorker                         ResourceTypeID = "worker"
	ResourceTypeIDWorkerRoute                    ResourceTypeID = "worker-route"
	ResourceTypeIDWorkersAiModel                 ResourceTypeID = "workers-ai-model"
	ResourceTypeIDWorkflow                       ResourceTypeID = "workflow"
	ResourceTypeIDWorkspace                      ResourceTypeID = "workspace"
	ResourceTypeIDZone                           ResourceTypeID = "zone"
)

// ResourceTypeSummary is the `ResourceTypeSummary` schema.
type ResourceTypeSummary struct {
	ID                    string                             `json:"id"`
	DisplayName           string                             `json:"displayName"`
	PluralDisplayName     *string                            `json:"pluralDisplayName,omitempty"`
	ParentTypeID          *string                            `json:"parentTypeId,omitempty"`
	SupportsCreate        bool                               `json:"supportsCreate"`
	AttachTargets         []ResourceTypeSummaryAttachTargets `json:"attachTargets,omitempty"`
	IsSSHHost             *bool                              `json:"isSshHost,omitempty"`
	SSHTunnelAttachSource *bool                              `json:"sshTunnelAttachSource,omitempty"`
	ShowInSidebar         *bool                              `json:"showInSidebar,omitempty"`
	AccountRoot           *bool                              `json:"accountRoot,omitempty"`
	// Schedulable: The type declares lifecycle start/stop actions, so its
	// resources can carry a sleep/wake schedule.
	Schedulable *bool `json:"schedulable,omitempty"`
}

// RestoreDrill is the `RestoreDrill` schema.
type RestoreDrill struct {
	ID           string  `json:"id"`
	ResourceID   string  `json:"resourceId"`
	ResourceName *string `json:"resourceName"`
	AccountID    *string `json:"accountId"`
	AccountName  *string `json:"accountName"`
	// PerformedAt: When the drill was performed, which is **not** when it was
	// recorded — people write these up on Monday for a drill they ran on
	// Saturday, and every staleness computation uses this.
	PerformedAt string `json:"performedAt"`
	// Outcome: How the drill ended. Only `verified` counts as evidence the
	// backup works: a restore that produced a running system nobody looked
	// inside is exactly how a team discovers, mid-incident, that the dump had
	// been empty for months. `restored-unverified` is recorded because doing the
	// restore is worth recording, but it does not reset the clock.
	//
	// One of "verified", "restored-unverified", "failed", "blocked".
	Outcome string `json:"outcome"`
	// RtoMinutes: Measured wall-clock minutes. Null when the drill never got
	// that far; a blocked drill has no RTO, and an invented one would be the
	// most dangerous number on the page.
	RtoMinutes *int64 `json:"rtoMinutes"`
	// RestoredFrom: Snapshot id, S3 key, a date — free text.
	RestoredFrom      *string `json:"restoredFrom"`
	Notes             *string `json:"notes"`
	PerformedByUserID *string `json:"performedByUserId"`
	PerformedByName   *string `json:"performedByName"`
	CreatedAt         string  `json:"createdAt"`
}

// RestoreDrillCreate is the `RestoreDrillCreate` schema.
type RestoreDrillCreate struct {
	ResourceID  string `json:"resourceId"`
	PerformedAt string `json:"performedAt"`
	// Outcome: How the drill ended. Only `verified` counts as evidence the
	// backup works: a restore that produced a running system nobody looked
	// inside is exactly how a team discovers, mid-incident, that the dump had
	// been empty for months. `restored-unverified` is recorded because doing the
	// restore is worth recording, but it does not reset the clock.
	//
	// One of "verified", "restored-unverified", "failed", "blocked".
	Outcome      string  `json:"outcome"`
	RtoMinutes   *int64  `json:"rtoMinutes,omitempty"`
	RestoredFrom *string `json:"restoredFrom,omitempty"`
	Notes        *string `json:"notes,omitempty"`
}

// RevertApplyResponse is the `RevertApplyResponse` schema.
type RevertApplyResponse struct {
	ChangeID   string     `json:"changeId"`
	ResourceID ResourceID `json:"resourceId"`
	// AppliedFields: The fields written, in plan order. Empty on a
	// reconciliation.
	AppliedFields []string   `json:"appliedFields"`
	Plan          RevertPlan `json:"plan"`
	RevertedAt    string     `json:"revertedAt"`
	// Reconciled: True when this request wrote nothing and instead recorded an
	// *earlier* interrupted attempt's write — the resource was already back, and
	// the event is now marked reverted. Nothing was sent to the provider by this
	// request.
	Reconciled *bool `json:"reconciled,omitempty"`
	// AuditRecorded: Present and `false` only when the audit entry could not be
	// written. The provider change still happened; its attribution did not reach
	// the audit table and was written to the server log instead. Attribution is
	// best-effort — nothing transactional spans a third-party cloud API and
	// Infrawrench's database.
	AuditRecorded *bool `json:"auditRecorded,omitempty"`
}

// RevertFieldPlan is the `RevertFieldPlan` schema.
type RevertFieldPlan struct {
	Field string `json:"field"`
	// RevertTo: The value a revert would write.
	RevertTo any `json:"revertTo,omitempty"`
	// ChangedTo: The value the recorded change set.
	ChangedTo any `json:"changedTo,omitempty"`
	// Current: The value the resource holds right now, read live.
	Current any               `json:"current,omitempty"`
	Status  RevertFieldStatus `json:"status"`
	// Reason: One sentence explaining the status.
	Reason string `json:"reason"`
}

// RevertFieldStatus: What a revert would do to one field. `revertible` — the
// field still holds the value the change set, and the plugin's edit form can
// write the old one back. `already-reverted` — it is already at the old value;
// nothing to do. `conflict` — it changed again since, so reverting would discard
// the newer value. `not-writable` — outside the plugin's editable surface, or
// the old value is not something the edit form can submit. `provider-derived` —
// an `outputs.*` entry, which the provider computes rather than accepts.
type RevertFieldStatus = string

// The values RevertFieldStatus takes.
const (
	RevertFieldStatusRevertible      RevertFieldStatus = "revertible"
	RevertFieldStatusAlreadyReverted RevertFieldStatus = "already-reverted"
	RevertFieldStatusConflict        RevertFieldStatus = "conflict"
	RevertFieldStatusNotWritable     RevertFieldStatus = "not-writable"
	RevertFieldStatusProviderDerived RevertFieldStatus = "provider-derived"
)

// RevertPlan is the `RevertPlan` schema.
type RevertPlan struct {
	// Fields: Every field of the recorded diff, in the order the event recorded
	// them.
	Fields []RevertFieldPlan `json:"fields"`
	// RevertibleFields: The keys that would actually be written.
	RevertibleFields []string `json:"revertibleFields"`
	Revertible       bool     `json:"revertible"`
	// BlockedReason: Why nothing would be written, or null when something would.
	BlockedReason *string `json:"blockedReason"`
}

// RevertPreviewResponse is the `RevertPreviewResponse` schema.
type RevertPreviewResponse struct {
	ChangeID       string     `json:"changeId"`
	ResourceID     ResourceID `json:"resourceId"`
	DisplayName    string     `json:"displayName"`
	PluginID       string     `json:"pluginId"`
	ResourceTypeID string     `json:"resourceTypeId"`
	AccountID      string     `json:"accountId"`
	Plan           RevertPlan `json:"plan"`
	RevertedAt     *string    `json:"revertedAt"`
}

// RightsizingListResponse is the `RightsizingListResponse` schema.
type RightsizingListResponse struct {
	// Accounts: Groups sorted by account name.
	Accounts   []OversizedAccountGroup `json:"accounts"`
	TotalCount int64                   `json:"totalCount"`
	// WindowDays: Days of stored metrics the percentiles cover.
	WindowDays  int64  `json:"windowDays"`
	GeneratedAt string `json:"generatedAt"`
}

// Role is the `Role` schema.
type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	IsSystem    bool         `json:"isSystem"`
	SystemKey   *string      `json:"systemKey"`
	Permissions []Permission `json:"permissions"`
}

// RoleChangeRequest is the `RoleChangeRequest` schema.
type RoleChangeRequest struct {
	Role   *OrganizationRole `json:"role,omitempty"`
	RoleID *string           `json:"roleId,omitempty"`
}

// RoleCreateRequest is the `RoleCreateRequest` schema.
type RoleCreateRequest struct {
	Name        string       `json:"name"`
	Description *string      `json:"description,omitempty"`
	Permissions []Permission `json:"permissions"`
}

// RoleSummary is the `RoleSummary` schema.
//
// The API may send null in its place.
type RoleSummary struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IsSystem    bool    `json:"isSystem"`
	SystemKey   *string `json:"systemKey"`
}

// RoleUpdateRequest is the `RoleUpdateRequest` schema.
type RoleUpdateRequest struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

// Runbook is the `Runbook` schema.
type Runbook struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	Steps       []RunbookStep `json:"steps"`
	// ResourceTypeIDs: Resource types this runbook is about; empty means it is
	// not scoped to a type. Used to answer 'which runbooks apply here',
	// **never** to restrict who may open it — a runbook nobody can find is the
	// failure this feature exists to fix.
	ResourceTypeIDs []string `json:"resourceTypeIds"`
	// TagKey: Optional tag narrowing. Matched case-insensitively.
	TagKey *string `json:"tagKey"`
	// TagValue: Required value of `tagKey`, matched exactly.
	TagValue *string `json:"tagValue"`
	// Enabled: Off keeps the row and hides it from the 'what applies here'
	// lookup. Retiring a runbook must not cost you the history of the runs
	// performed against it.
	Enabled         bool    `json:"enabled"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedByName   *string `json:"createdByName"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
	RunCount        int64   `json:"runCount"`
	LastRunAt       *string `json:"lastRunAt"`
}

// RunbookCreate is the `RunbookCreate` schema.
type RunbookCreate struct {
	Name            string             `json:"name"`
	Description     *string            `json:"description,omitempty"`
	Steps           []RunbookStepInput `json:"steps,omitempty"`
	ResourceTypeIDs []string           `json:"resourceTypeIds,omitempty"`
	TagKey          *string            `json:"tagKey,omitempty"`
	TagValue        *string            `json:"tagValue,omitempty"`
	Enabled         *bool              `json:"enabled,omitempty"`
}

// RunbookList is the `RunbookList` schema.
type RunbookList struct {
	Runbooks []Runbook `json:"runbooks"`
}

// RunbookRun is the `RunbookRun` schema.
type RunbookRun struct {
	ID        string `json:"id"`
	RunbookID string `json:"runbookId"`
	// RunbookName: The runbook's name when the run started.
	RunbookName string `json:"runbookName"`
	// Status: One of "running", "completed", "abandoned".
	Status string `json:"status"`
	// IncidentID: The incident this was performed under. Not a cascading
	// reference: deleting the incident must not delete the record that somebody
	// followed the failover procedure at 03:14.
	IncidentID      *string          `json:"incidentId"`
	StartedByUserID *string          `json:"startedByUserId"`
	StartedByName   *string          `json:"startedByName"`
	StartedAt       string           `json:"startedAt"`
	CompletedAt     *string          `json:"completedAt"`
	Summary         *string          `json:"summary"`
	Steps           []RunbookRunStep `json:"steps"`
}

// RunbookRunClose is the `RunbookRunClose` schema.
type RunbookRunClose struct {
	// Status: One of "completed", "abandoned".
	Status  string  `json:"status"`
	Summary *string `json:"summary,omitempty"`
}

// RunbookRunList is the `RunbookRunList` schema.
type RunbookRunList struct {
	Runs []RunbookRun `json:"runs"`
}

// RunbookRunStart is the `RunbookRunStart` schema.
type RunbookRunStart struct {
	IncidentID *string `json:"incidentId,omitempty"`
}

// RunbookRunStep is the `RunbookRunStep` schema.
type RunbookRunStep struct {
	StepID string `json:"stepId"`
	// Title: The step's title **when the run started**. Copied rather than
	// joined: a runbook is edited between incidents, and a postmortem showing
	// today's wording against last month's run is not stale, it is quietly
	// wrong.
	Title string `json:"title"`
	// Kind: What the step does. Three kinds and not a scripting language: a
	// runbook is written by whoever is on call for whoever is on call next, and
	// the moment it needs a language it stops being written. `workflow` is the
	// escape hatch — anything genuinely automated belongs in a workflow, which
	// already has a sandbox, approvals, secrets and a history.
	//
	// One of "manual", "workflow", "link".
	Kind string `json:"kind"`
	// Status: One of "pending", "done", "skipped", "failed".
	Status string `json:"status"`
	// Note: What the responder typed — output, or why it was skipped.
	Note *string `json:"note"`
	// WorkflowRunID: The workflow run this step kicked off. Recorded here; the
	// run itself goes through the workflow routes with their own permission,
	// approvals and secrets.
	WorkflowRunID *string `json:"workflowRunId"`
	ActorUserID   *string `json:"actorUserId"`
	ActorName     *string `json:"actorName"`
	UpdatedAt     *string `json:"updatedAt"`
}

// RunbookStep is the `RunbookStep` schema.
type RunbookStep struct {
	// ID: Stable across edits, because a run's per-step records reference it.
	// Reordering or retitling keeps the same step; deleting one orphans its
	// history, which is why runs keep the title they saw.
	ID string `json:"id"`
	// Kind: What the step does. Three kinds and not a scripting language: a
	// runbook is written by whoever is on call for whoever is on call next, and
	// the moment it needs a language it stops being written. `workflow` is the
	// escape hatch — anything genuinely automated belongs in a workflow, which
	// already has a sandbox, approvals, secrets and a history.
	//
	// One of "manual", "workflow", "link".
	Kind  string `json:"kind"`
	Title string `json:"title"`
	// Body: Markdown — the detail nobody remembers at 03:00.
	Body string `json:"body"`
	// WorkflowID: For `workflow` steps: which workflow the button runs.
	WorkflowID *string `json:"workflowId,omitempty"`
	// URL: For `link` steps. `https:` only.
	URL *string `json:"url,omitempty"`
}

// RunbookStepInput is the `RunbookStepInput` schema.
type RunbookStepInput struct {
	// ID: Omitted for a new step; the server assigns one.
	ID *string `json:"id,omitempty"`
	// Kind: What the step does. Three kinds and not a scripting language: a
	// runbook is written by whoever is on call for whoever is on call next, and
	// the moment it needs a language it stops being written. `workflow` is the
	// escape hatch — anything genuinely automated belongs in a workflow, which
	// already has a sandbox, approvals, secrets and a history.
	//
	// One of "manual", "workflow", "link".
	Kind       string  `json:"kind"`
	Title      string  `json:"title"`
	Body       *string `json:"body,omitempty"`
	WorkflowID *string `json:"workflowId,omitempty"`
	URL        *string `json:"url,omitempty"`
}

// RunbookStepUpdate is the `RunbookStepUpdate` schema.
type RunbookStepUpdate struct {
	// Status: One of "pending", "done", "skipped", "failed".
	Status string `json:"status"`
	// Note: Omitted leaves the note alone; `null` clears it.
	Note          *string `json:"note,omitempty"`
	WorkflowRunID *string `json:"workflowRunId,omitempty"`
}

// RunbookUpdate is the `RunbookUpdate` schema.
type RunbookUpdate struct {
	Name            *string            `json:"name,omitempty"`
	Description     *string            `json:"description,omitempty"`
	Steps           []RunbookStepInput `json:"steps,omitempty"`
	ResourceTypeIDs []string           `json:"resourceTypeIds,omitempty"`
	TagKey          *string            `json:"tagKey,omitempty"`
	TagValue        *string            `json:"tagValue,omitempty"`
	Enabled         *bool              `json:"enabled,omitempty"`
}

// SavedCostFilter is the `SavedCostFilter` schema.
type SavedCostFilter struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	Filters     []SavedCostFilterTerm `json:"filters"`
	// Query: The canonical cost-query-language rendering of `filters`, derived
	// server-side.
	Query           string  `json:"query"`
	CreatedByUserID *string `json:"createdByUserId"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// SavedCostFilterInput is the `SavedCostFilterInput` schema.
type SavedCostFilterInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Filters: The structured filter. May be omitted only when `query` is sent
	// instead.
	Filters []SavedCostFilterTerm `json:"filters,omitempty"`
	// Query: The same filter written in the cost query language — an alternative
	// spelling of `filters`, compiled server-side into exactly that structure.
	// Sending both a query and a non-empty `filters` is a 400, not a precedence
	// rule. Whichever spelling is used, the result must be non-empty (an empty
	// saved filter matches everything, which is the same as no filter wearing a
	// name) and every tag term must carry its key.
	Query *string `json:"query,omitempty"`
}

// SavedCostFilterReferent is the `SavedCostFilterReferent` schema.
type SavedCostFilterReferent struct {
	// Kind: One of "budget", "cost_report", "cost_graph_widget".
	Kind string `json:"kind"`
	// ID: Budget id, report id, or dashboard-widget id.
	ID string `json:"id"`
	// Name: Budget name, report name, or the widget's title.
	Name string `json:"name"`
	// DashboardID: Set for `cost_graph_widget` referents.
	DashboardID   *string `json:"dashboardId,omitempty"`
	DashboardName *string `json:"dashboardName,omitempty"`
}

// SavedCostFilterTerm is the `SavedCostFilterTerm` schema.
type SavedCostFilterTerm struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag", "charge_type", "commitment".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// ScheduleConflict is the `ScheduleConflict` schema.
type ScheduleConflict struct {
	Error string `json:"error"`
}

// ScheduleTransition is the `ScheduleTransition` schema.
type ScheduleTransition struct {
	At string `json:"at"`
	// Action: A schedule transition: `stop` powers the resource off, `start`
	// powers it on.
	//
	// One of "stop", "start".
	Action string `json:"action"`
}

// SearchHit is the `SearchHit` schema.
type SearchHit struct {
	ID                ResourceID `json:"id"`
	PluginID          string     `json:"pluginId"`
	PluginDisplayName string     `json:"pluginDisplayName"`
	PluginLogoSvg     string     `json:"pluginLogoSvg"`
	ResourceTypeID    string     `json:"resourceTypeId"`
	ResourceTypeLabel string     `json:"resourceTypeLabel"`
	AccountID         string     `json:"accountId"`
	AccountName       string     `json:"accountName"`
	DisplayName       string     `json:"displayName"`
	Subtitle          *string    `json:"subtitle,omitempty"`
}

// SeatLimitResponse is the `SeatLimitResponse` schema.
type SeatLimitResponse struct {
	Error string `json:"error"`
	// Code: One of "seat_limit_reached".
	Code string `json:"code"`
	// SeatCount: Total capacity: monthly subscription seats plus prepaid
	// capacity-slot seats
	SeatCount int64 `json:"seatCount"`
	// SeatsUsed: Members plus pending unexpired invitations
	SeatsUsed int64 `json:"seatsUsed"`
	// CanAddSeat: Whether retrying with `addSeat: true` can succeed. False when
	// the org's capacity is entirely prepaid capacity slots: there is no monthly
	// seat to buy, so the only remedy is another capacity slot.
	CanAddSeat bool `json:"canAddSeat"`
}

// SecretAccessRequest is the `SecretAccessRequest` schema.
type SecretAccessRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceID       ResourceID  `json:"resourceId"`
	VersionID        string      `json:"versionId"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// SecretAccessResponse is the `SecretAccessResponse` schema.
type SecretAccessResponse struct {
	Value string `json:"value"`
}

// SecretAddRequest is the `SecretAddRequest` schema.
type SecretAddRequest struct {
	AccountID        string      `json:"accountId"`
	ResourceID       ResourceID  `json:"resourceId"`
	Value            string      `json:"value"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// SecretExportTemplate is the `SecretExportTemplate` schema.
type SecretExportTemplate struct {
	ID          string                        `json:"id"`
	Label       string                        `json:"label"`
	Description *string                       `json:"description,omitempty"`
	Entries     []SecretExportTemplateEntries `json:"entries"`
}

// SecretModifyRequest is the `SecretModifyRequest` schema.
type SecretModifyRequest struct {
	AccountID  string     `json:"accountId"`
	ResourceID ResourceID `json:"resourceId"`
	VersionID  string     `json:"versionId"`
	// Action: One of "enable", "disable", "destroy".
	Action           string      `json:"action"`
	ParentResourceID *ResourceID `json:"parentResourceId,omitempty"`
}

// SecretVersion is the `SecretVersion` schema.
type SecretVersion struct {
	ID string `json:"id"`
	// State: One of "enabled", "disabled", "destroyed".
	State string `json:"state"`
	// CreatedAt: ISO-8601.
	CreatedAt string `json:"createdAt"`
	// DestroyedAt: Set only when destroyed.
	DestroyedAt *string `json:"destroyedAt,omitempty"`
	IsLatest    *bool   `json:"isLatest,omitempty"`
}

// SecretVersionResponse is the `SecretVersionResponse` schema.
type SecretVersionResponse struct {
	Version SecretVersion `json:"version"`
}

// SecretVersionsResponse is the `SecretVersionsResponse` schema.
type SecretVersionsResponse struct {
	Versions []SecretVersion `json:"versions"`
}

// Session is the `Session` schema.
type Session struct {
	UserID          string  `json:"userId"`
	Email           *string `json:"email"`
	NeedsOnboarding bool    `json:"needsOnboarding"`
}

// SessionRecording is the `SessionRecording` schema.
type SessionRecording struct {
	ID string `json:"id"`
	// UserID: Who opened the session; null when the socket authenticated with an
	// API key.
	UserID *string `json:"userId"`
	// UserName: Display-name snapshot taken at record time, so a departed member
	// still reads as one.
	UserName   *string `json:"userName"`
	AccountID  *string `json:"accountId"`
	ResourceID *string `json:"resourceId"`
	// Host: Final hop, as dialled.
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Username string `json:"username"`
	// HopCount: 1 for a direct session; higher when it jumped through bastions.
	HopCount int64 `json:"hopCount"`
	Cols     int64 `json:"cols"`
	Rows     int64 `json:"rows"`
	// HasInput: True when the cast also contains keystrokes (the org opted into
	// input capture).
	HasInput bool `json:"hasInput"`
	// SharedConsoleID: Set when this session was shared with colleagues while it
	// ran.
	SharedConsoleID *string `json:"sharedConsoleId,omitempty"`
	// Participants: Everyone who was attached to this session and in what role —
	// the **highest** role they held, not their role at the end. Null or empty
	// for an ordinary solo session. Once a session can be shared, `userId` alone
	// stops answering 'whose hands were on this box'; this does. The cast
	// carries the same facts in-band as asciicast `"m"` marker events, so a
	// viewer sees *when* the keyboard moved.
	Participants []SessionRecordingParticipants `json:"participants,omitempty"`
	// Status: `recording` (live), `complete` (closed cleanly), `truncated` (hit
	// the per-session capture ceiling — the tape is a genuine partial and says
	// so), or `abandoned` (the server handling the session went away before it
	// could close the row).
	//
	// One of "recording", "complete", "truncated", "abandoned".
	Status string `json:"status"`
	// OutputBytes: Terminal bytes captured, before compression.
	OutputBytes int64   `json:"outputBytes"`
	EventCount  int64   `json:"eventCount"`
	StartedAt   string  `json:"startedAt"`
	EndedAt     *string `json:"endedAt"`
	DurationMs  *int64  `json:"durationMs"`
}

// SessionRecordingSettings is the `SessionRecordingSettings` schema.
type SessionRecordingSettings struct {
	Enabled bool `json:"enabled"`
	// CaptureInput: Also record keystrokes. Separate from `enabled` because it
	// captures input at prompts the remote host chose not to echo — a sudo
	// password, a pasted token — which is a materially different promise to the
	// people being recorded.
	CaptureInput  bool                  `json:"captureInput"`
	RetentionDays int64                 `json:"retentionDays"`
	Usage         SessionRecordingUsage `json:"usage"`
}

// SessionRecordingSettingsUpdate is the `SessionRecordingSettingsUpdate` schema.
type SessionRecordingSettingsUpdate struct {
	Enabled       *bool  `json:"enabled,omitempty"`
	CaptureInput  *bool  `json:"captureInput,omitempty"`
	RetentionDays *int64 `json:"retentionDays,omitempty"`
}

// SessionRecordingUsage is the `SessionRecordingUsage` schema.
type SessionRecordingUsage struct {
	RecordingCount int64 `json:"recordingCount"`
	// StoredBytes: Compressed size actually stored.
	StoredBytes     int64   `json:"storedBytes"`
	CapturedBytes   int64   `json:"capturedBytes"`
	OldestStartedAt *string `json:"oldestStartedAt"`
}

// SFTPDeleteRequest is the `SftpDeleteRequest` schema.
//
// Spec schema: `SftpDeleteRequest`.
type SFTPDeleteRequest struct {
	AccountID   string  `json:"accountId"`
	Path        string  `json:"path"`
	SSHKeyID    *string `json:"sshKeyId,omitempty"`
	SSHHost     *string `json:"sshHost,omitempty"`
	SSHUsername *string `json:"sshUsername,omitempty"`
	IsDir       bool    `json:"isDir"`
}

// SFTPEntry is the `SftpEntry` schema.
//
// Spec schema: `SftpEntry`.
type SFTPEntry struct {
	// Key: Absolute remote path.
	Key          string  `json:"key"`
	Name         string  `json:"name"`
	Size         float64 `json:"size"`
	LastModified string  `json:"lastModified"`
	IsDirectory  bool    `json:"isDirectory"`
	ContentType  *string `json:"contentType,omitempty"`
}

// SFTPListRequest is the `SftpListRequest` schema.
//
// Spec schema: `SftpListRequest`.
type SFTPListRequest struct {
	AccountID   string  `json:"accountId"`
	Path        string  `json:"path"`
	SSHKeyID    *string `json:"sshKeyId,omitempty"`
	SSHHost     *string `json:"sshHost,omitempty"`
	SSHUsername *string `json:"sshUsername,omitempty"`
}

// SFTPPathRequest is the `SftpPathRequest` schema.
//
// Spec schema: `SftpPathRequest`.
type SFTPPathRequest struct {
	AccountID   string  `json:"accountId"`
	Path        string  `json:"path"`
	SSHKeyID    *string `json:"sshKeyId,omitempty"`
	SSHHost     *string `json:"sshHost,omitempty"`
	SSHUsername *string `json:"sshUsername,omitempty"`
}

// SFTPUploadForm is the `SftpUploadForm` schema.
//
// Spec schema: `SftpUploadForm`.
type SFTPUploadForm struct {
	AccountID   string    `json:"accountId"`
	RemotePath  string    `json:"remotePath"`
	File        io.Reader `json:"file"`
	SSHKeyID    *string   `json:"sshKeyId,omitempty"`
	SSHHost     *string   `json:"sshHost,omitempty"`
	SSHUsername *string   `json:"sshUsername,omitempty"`
}

// SharedConsole is the `SharedConsole` schema.
type SharedConsole struct {
	ID string `json:"id"`
	// RoutingKey: Load-balancer affinity hint. A guest's WebSocket must carry it
	// as `?sid=` so the upgrade lands on the replica holding the pty. Not a
	// secret and not authorisation.
	RoutingKey  string  `json:"routingKey"`
	OwnerUserID *string `json:"ownerUserId"`
	OwnerName   *string `json:"ownerName"`
	AccountID   *string `json:"accountId"`
	ResourceID  *string `json:"resourceId"`
	// Host: Final hop, as the proxy dialled it — never as a client asserted it.
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Username string `json:"username"`
	// AllowHandover: False makes the share strictly read-only: nobody but the
	// sharer can ever type. This is the one hard safety property the feature
	// offers, as opposed to inferring intent from command text.
	AllowHandover bool `json:"allowHandover"`
	// Status: `revoked` — somebody ended the share; `ended` — the underlying SSH
	// session closed. Either way the fan-out stops and attached guests are
	// disconnected.
	//
	// One of "active", "revoked", "ended".
	Status            string  `json:"status"`
	InviteTokenPrefix *string `json:"inviteTokenPrefix"`
	InviteExpiresAt   *string `json:"inviteExpiresAt"`
	// InviteConsumedAt: Set once an invite admitted somebody new. The link stops
	// working for anyone else at that moment; the sharer mints a replacement for
	// the next guest.
	InviteConsumedAt *string `json:"inviteConsumedAt"`
	// RecordingID: The session recording this console is being taped into, when
	// the org records. Participants are attributed in that recording's own
	// metadata and as asciicast markers on its timeline.
	RecordingID *string `json:"recordingId"`
	PtyCols     int64   `json:"ptyCols"`
	// PtyRows: The pty's geometry, which is the **driver's** geometry. One pty
	// has one size, so everyone else letterboxes rather than reflowing.
	PtyRows   int64  `json:"ptyRows"`
	CreatedAt string `json:"createdAt"`
}

// SharedConsoleCreated is the `SharedConsoleCreated` schema.
type SharedConsoleCreated struct {
	Share        SharedConsole              `json:"share"`
	Participants []SharedConsoleParticipant `json:"participants"`
	// InviteToken: The invite, returned exactly once. Only its sha256 is stored,
	// so it cannot be shown again — mint a replacement instead.
	InviteToken string `json:"inviteToken"`
}

// SharedConsoleInvitePreview is the `SharedConsoleInvitePreview` schema.
type SharedConsoleInvitePreview struct {
	Share    SharedConsole `json:"share"`
	Joinable bool          `json:"joinable"`
	// Rejoin: You are already on this console and would resume.
	Rejoin *bool   `json:"rejoin,omitempty"`
	Error  *string `json:"error,omitempty"`
	Code   *string `json:"code,omitempty"`
}

// SharedConsoleJoined is the `SharedConsoleJoined` schema.
type SharedConsoleJoined struct {
	Share        SharedConsole              `json:"share"`
	Participants []SharedConsoleParticipant `json:"participants"`
	You          SharedConsoleParticipant   `json:"you"`
	RoutingKey   string                     `json:"routingKey"`
}

// SharedConsoleParticipant is the `SharedConsoleParticipant` schema.
type SharedConsoleParticipant struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	// UserName: Display-name snapshot taken when they joined.
	UserName *string `json:"userName"`
	// Role: `driver` holds the keyboard; `observer` sees the terminal and cannot
	// type into it. Exactly one participant per console is a driver at any
	// moment, enforced by a partial unique index rather than by the application
	// — two simultaneous handovers cannot both win.
	//
	// One of "observer", "driver".
	Role string `json:"role"`
	// Status: `left` walked away and may resume on the same row without a new
	// invite; `removed` was ejected or lost the permission mid-session and needs
	// a fresh one.
	//
	// One of "joined", "left", "removed".
	Status string `json:"status"`
	// DriverRequestedAt: Set when this participant has asked for the keyboard
	// and nobody has answered yet. Asking grants nothing — only the current
	// driver or the sharer can move it.
	DriverRequestedAt *string `json:"driverRequestedAt"`
	JoinedAt          string  `json:"joinedAt"`
}

// SharedConsoleState is the `SharedConsoleState` schema.
type SharedConsoleState struct {
	Share        SharedConsole              `json:"share"`
	Participants []SharedConsoleParticipant `json:"participants"`
}

// SharedConsoleSummary is the `SharedConsoleSummary` schema.
type SharedConsoleSummary struct {
	ID string `json:"id"`
	// RoutingKey: Load-balancer affinity hint. A guest's WebSocket must carry it
	// as `?sid=` so the upgrade lands on the replica holding the pty. Not a
	// secret and not authorisation.
	RoutingKey  string  `json:"routingKey"`
	OwnerUserID *string `json:"ownerUserId"`
	OwnerName   *string `json:"ownerName"`
	AccountID   *string `json:"accountId"`
	ResourceID  *string `json:"resourceId"`
	// Host: Final hop, as the proxy dialled it — never as a client asserted it.
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Username string `json:"username"`
	// AllowHandover: False makes the share strictly read-only: nobody but the
	// sharer can ever type. This is the one hard safety property the feature
	// offers, as opposed to inferring intent from command text.
	AllowHandover bool `json:"allowHandover"`
	// Status: `revoked` — somebody ended the share; `ended` — the underlying SSH
	// session closed. Either way the fan-out stops and attached guests are
	// disconnected.
	//
	// One of "active", "revoked", "ended".
	Status            string  `json:"status"`
	InviteTokenPrefix *string `json:"inviteTokenPrefix"`
	InviteExpiresAt   *string `json:"inviteExpiresAt"`
	// InviteConsumedAt: Set once an invite admitted somebody new. The link stops
	// working for anyone else at that moment; the sharer mints a replacement for
	// the next guest.
	InviteConsumedAt *string `json:"inviteConsumedAt"`
	// RecordingID: The session recording this console is being taped into, when
	// the org records. Participants are attributed in that recording's own
	// metadata and as asciicast markers on its timeline.
	RecordingID *string `json:"recordingId"`
	PtyCols     int64   `json:"ptyCols"`
	// PtyRows: The pty's geometry, which is the **driver's** geometry. One pty
	// has one size, so everyone else letterboxes rather than reflowing.
	PtyRows      int64                      `json:"ptyRows"`
	CreatedAt    string                     `json:"createdAt"`
	Participants []SharedConsoleParticipant `json:"participants"`
}

// ShowbackReport is the `ShowbackReport` schema.
type ShowbackReport struct {
	From       string                 `json:"from"`
	To         string                 `json:"to"`
	Currencies []string               `json:"currencies"`
	Adjustment *CostAdjustmentSummary `json:"adjustment,omitempty"`
	// Centres: Depth-first: each centre immediately followed by its children,
	// siblings name-sorted, with the "Unallocated" bucket last.
	Centres []ShowbackReportCentres `json:"centres"`
}

// SignSSHKeyRequest is the `SignSshKeyRequest` schema.
//
// Spec schema: `SignSshKeyRequest`.
type SignSSHKeyRequest struct {
	// Data: The exact bytes SSH wants signed (a publickey-auth challenge),
	// base64-encoded.
	Data      string           `json:"data"`
	Algorithm SSHSignAlgorithm `json:"algorithm"`
	// Context: Recorded in the audit log entry for this signature.
	Context *SignSshkeyRequestContext `json:"context,omitempty"`
}

// SignSSHKeyResponse is the `SignSshKeyResponse` schema.
//
// Spec schema: `SignSshKeyResponse`.
type SignSSHKeyResponse struct {
	// Signature: Raw signature bytes, base64-encoded — Ed25519/RSA as-is, ECDSA
	// in DER as node produces it.
	Signature string           `json:"signature"`
	Algorithm SSHSignAlgorithm `json:"algorithm"`
}

// SlackAvailableChannel is the `SlackAvailableChannel` schema.
type SlackAvailableChannel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"isPrivate"`
}

// SlackChannel is the `SlackChannel` schema.
type SlackChannel struct {
	ID             string `json:"id"`
	InstallationID string `json:"installationId"`
	// ChannelID: Slack channel id (C…/G…)
	ChannelID string `json:"channelId"`
	// ChannelName: Channel name without the leading #
	ChannelName string `json:"channelName"`
	IsPrivate   bool   `json:"isPrivate"`
}

// SlackChannelCreate is the `SlackChannelCreate` schema.
type SlackChannelCreate struct {
	InstallationID string `json:"installationId"`
	ChannelID      string `json:"channelId"`
	ChannelName    string `json:"channelName"`
	IsPrivate      *bool  `json:"isPrivate,omitempty"`
}

// SlackChannelUpdate is the `SlackChannelUpdate` schema.
type SlackChannelUpdate struct {
	ChannelName string `json:"channelName"`
}

// SlackInstallation is the `SlackInstallation` schema.
type SlackInstallation struct {
	// ID: Infrawrench id for this workspace connection
	ID string `json:"id"`
	// TeamID: Slack workspace id (T…)
	TeamID   string  `json:"teamId"`
	TeamName *string `json:"teamName"`
}

// SlackStatus is the `SlackStatus` schema.
type SlackStatus struct {
	// Configured: True when this deployment has a Slack app registered
	Configured    bool                `json:"configured"`
	Installations []SlackInstallation `json:"installations"`
	Channels      []SlackChannel      `json:"channels"`
}

// SleepSchedule is the `SleepSchedule` schema.
type SleepSchedule struct {
	ID string `json:"id"`
	// ResourceID: Infrawrench resource id the schedule powers on and off.
	ResourceID     string   `json:"resourceId"`
	AccountID      string   `json:"accountId"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	// ResourceName: Resource display name at read time.
	ResourceName string `json:"resourceName"`
	AccountName  string `json:"accountName"`
	// DaysOfWeek: ISO weekdays the resource is worked on: 1 = Monday … 7 =
	// Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek"`
	// StopTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StopTime string `json:"stopTime"`
	// StartTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StartTime string `json:"startTime"`
	// Timezone: IANA timezone the wall-clock times are computed in (DST-safe).
	Timezone string `json:"timezone"`
	// Paused: Paused schedules keep their timing but never fire.
	Paused bool `json:"paused"`
	// NextTransitionAt: Next due transition; null while paused.
	NextTransitionAt *string `json:"nextTransitionAt"`
	// NextTransitionAction: A schedule transition: `stop` powers the resource
	// off, `start` powers it on.
	//
	// One of "stop", "start".
	NextTransitionAction *string `json:"nextTransitionAction"`
	LastRunAt            *string `json:"lastRunAt"`
	// LastRunAction: A schedule transition: `stop` powers the resource off,
	// `start` powers it on.
	//
	// One of "stop", "start".
	LastRunAction *string `json:"lastRunAction"`
	// LastRunStatus: Outcome of the last executed transition: `ok`, `failed`
	// (see `lastRunError`), or `skipped_freeze` (an org change freeze was in
	// effect, so the transition was skipped).
	//
	// One of "ok", "failed", "skipped_freeze".
	LastRunStatus *string `json:"lastRunStatus"`
	// LastRunError: Failure detail for a failed run.
	LastRunError *string `json:"lastRunError"`
	// ProjectedMonthlySaving: Projected monthly saving from trailing
	// per-resource spend × the weekly off-hours fraction; null when billing
	// holds no rows for the resource.
	ProjectedMonthlySaving *float64 `json:"projectedMonthlySaving"`
	// Currency: Currency of the projection, when present.
	Currency  *string `json:"currency"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

// SleepScheduleCreate is the `SleepScheduleCreate` schema.
type SleepScheduleCreate struct {
	ResourceID string `json:"resourceId"`
	AccountID  string `json:"accountId"`
	// DaysOfWeek: ISO weekdays the resource is worked on: 1 = Monday … 7 =
	// Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek"`
	// StopTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StopTime string `json:"stopTime"`
	// StartTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StartTime string `json:"startTime"`
	// Timezone: IANA timezone the wall-clock times are computed in (DST-safe).
	Timezone string `json:"timezone"`
}

// SleepScheduleList is the `SleepScheduleList` schema.
type SleepScheduleList struct {
	Schedules []SleepSchedule `json:"schedules"`
}

// SleepSchedulePreview is the `SleepSchedulePreview` schema.
type SleepSchedulePreview struct {
	// OffFraction: Fraction of the week (0–1) the schedule keeps the resource
	// stopped.
	OffFraction float64 `json:"offFraction"`
	// MonthlyCost: Trailing spend normalized to a month; null when billing holds
	// no rows.
	MonthlyCost            *float64 `json:"monthlyCost"`
	ProjectedMonthlySaving *float64 `json:"projectedMonthlySaving"`
	Currency               *string  `json:"currency"`
	// CostWindowDays: Days of billing data the estimate was computed over (0 =
	// none found).
	CostWindowDays int64 `json:"costWindowDays"`
	// NextTransitions: The next few transitions, soonest first — a timezone
	// sanity check.
	NextTransitions []ScheduleTransition `json:"nextTransitions"`
}

// SleepSchedulePreviewRequest is the `SleepSchedulePreviewRequest` schema.
type SleepSchedulePreviewRequest struct {
	ResourceID string `json:"resourceId"`
	AccountID  string `json:"accountId"`
	// DaysOfWeek: ISO weekdays the resource is worked on: 1 = Monday … 7 =
	// Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek"`
	// StopTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StopTime string `json:"stopTime"`
	// StartTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StartTime string `json:"startTime"`
	// Timezone: IANA timezone the wall-clock times are computed in (DST-safe).
	Timezone string `json:"timezone"`
}

// SleepScheduleUpdate is the `SleepScheduleUpdate` schema.
type SleepScheduleUpdate struct {
	// DaysOfWeek: ISO weekdays the resource is worked on: 1 = Monday … 7 =
	// Sunday.
	DaysOfWeek []int64 `json:"daysOfWeek,omitempty"`
	// StopTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StopTime *string `json:"stopTime,omitempty"`
	// StartTime: Wall-clock time of day, 24-hour `"HH:MM"`, in the schedule's
	// timezone.
	StartTime *string `json:"startTime,omitempty"`
	// Timezone: IANA timezone the wall-clock times are computed in (DST-safe).
	Timezone *string `json:"timezone,omitempty"`
	Paused   *bool   `json:"paused,omitempty"`
}

// SQLEstimateRequest is the `SqlEstimateRequest` schema.
//
// Spec schema: `SqlEstimateRequest`.
type SQLEstimateRequest struct {
	AccountID  string     `json:"accountId"`
	ResourceID ResourceID `json:"resourceId"`
	SQL        string     `json:"sql"`
}

// SQLExecuteRequest is the `SqlExecuteRequest` schema.
//
// Spec schema: `SqlExecuteRequest`.
type SQLExecuteRequest struct {
	AccountID      string      `json:"accountId"`
	ResourceID     *ResourceID `json:"resourceId,omitempty"`
	ResourceTypeID *string     `json:"resourceTypeId,omitempty"`
	SQL            string      `json:"sql"`
	Params         []any       `json:"params,omitempty"`
}

// SQLExecuteResponse is the `SqlExecuteResponse` schema.
//
// Spec schema: `SqlExecuteResponse`.
type SQLExecuteResponse struct {
	AffectedRows int64 `json:"affectedRows"`
}

// SQLQueryRequest is the `SqlQueryRequest` schema.
//
// Spec schema: `SqlQueryRequest`.
type SQLQueryRequest struct {
	AccountID      string      `json:"accountId"`
	ResourceID     *ResourceID `json:"resourceId,omitempty"`
	ResourceTypeID *string     `json:"resourceTypeId,omitempty"`
	SQL            string      `json:"sql"`
}

// SQLQueryResponse is the `SqlQueryResponse` schema.
//
// Spec schema: `SqlQueryResponse`.
type SQLQueryResponse struct {
	Rows       []JSONObject `json:"rows"`
	DurationMs *int64       `json:"durationMs,omitempty"`
}

// SSHExecRequest is the `SshExecRequest` schema.
//
// Spec schema: `SshExecRequest`.
type SSHExecRequest struct {
	SSHHost  string `json:"sshHost"`
	SSHPort  int64  `json:"sshPort"`
	SSHUser  string `json:"sshUser"`
	SSHKeyID string `json:"sshKeyId"`
	Command  string `json:"command"`
}

// SSHExecResponse is the `SshExecResponse` schema.
//
// Spec schema: `SshExecResponse`.
type SSHExecResponse struct {
	Stdout string  `json:"stdout"`
	Stderr *string `json:"stderr,omitempty"`
	Code   int64   `json:"code"`
}

// SSHFanoutHostResult is the `SshFanoutHostResult` schema.
//
// Spec schema: `SshFanoutHostResult`.
type SSHFanoutHostResult struct {
	// Kind: One of "account", "resource".
	Kind     string `json:"kind"`
	TargetID string `json:"targetId"`
	Label    string `json:"label"`
	// Status: One of "done", "error", "blocked".
	Status       string                           `json:"status"`
	ExitCode     *int64                           `json:"exitCode"`
	Stdout       string                           `json:"stdout"`
	Stderr       string                           `json:"stderr"`
	Error        *string                          `json:"error,omitempty"`
	DurationMs   float64                          `json:"durationMs"`
	HostKeyTrust *SshfanoutHostResultHostKeyTrust `json:"hostKeyTrust,omitempty"`
}

// SSHFanoutRunRequest is the `SshFanoutRunRequest` schema.
//
// Spec schema: `SshFanoutRunRequest`.
type SSHFanoutRunRequest struct {
	Command     string                       `json:"command"`
	Targets     []SshfanoutRunRequestTargets `json:"targets"`
	SSHKeyID    *string                      `json:"sshKeyId,omitempty"`
	Username    *string                      `json:"username,omitempty"`
	Concurrency *int64                       `json:"concurrency,omitempty"`
}

// SSHFanoutRunResponse is the `SshFanoutRunResponse` schema.
//
// Spec schema: `SshFanoutRunResponse`.
type SSHFanoutRunResponse struct {
	Results []SSHFanoutHostResult `json:"results"`
}

// SSHFanoutTarget is the `SshFanoutTarget` schema.
//
// Spec schema: `SshFanoutTarget`.
type SSHFanoutTarget struct {
	// Kind: One of "account", "resource".
	Kind            string   `json:"kind"`
	ID              string   `json:"id"`
	AccountID       string   `json:"accountId"`
	Label           string   `json:"label"`
	PluginID        string   `json:"pluginId"`
	ResourceTypeID  *string  `json:"resourceTypeId,omitempty"`
	Host            *string  `json:"host,omitempty"`
	DefaultUsername *string  `json:"defaultUsername,omitempty"`
	Running         bool     `json:"running"`
	NeedsKey        bool     `json:"needsKey"`
	Tags            []string `json:"tags"`
}

// SSHFanoutTargetsResponse is the `SshFanoutTargetsResponse` schema.
//
// Spec schema: `SshFanoutTargetsResponse`.
type SSHFanoutTargetsResponse struct {
	Targets []SSHFanoutTarget `json:"targets"`
}

// SSHKey is the `SshKey` schema.
//
// Spec schema: `SshKey`.
type SSHKey struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	KeyType     SSHKeyType `json:"keyType"`
	IsImported  bool       `json:"isImported"`
	Fingerprint *string    `json:"fingerprint"`
	PublicKey   string     `json:"publicKey"`
	UserID      string     `json:"userId"`
	OwnerEmail  string     `json:"ownerEmail"`
	OwnerName   string     `json:"ownerName"`
	CreatedAt   string     `json:"createdAt"`
}

// SSHKeyType is the `SshKeyType` schema.
//
// Spec schema: `SshKeyType`.
type SSHKeyType = string

// The values SSHKeyType takes.
const (
	SSHKeyTypeSSHRsa                        SSHKeyType = "ssh-rsa"
	SSHKeyTypeSSHEd25519                    SSHKeyType = "ssh-ed25519"
	SSHKeyTypeSSHDss                        SSHKeyType = "ssh-dss"
	SSHKeyTypeEcdsaSha2Nistp256             SSHKeyType = "ecdsa-sha2-nistp256"
	SSHKeyTypeEcdsaSha2Nistp384             SSHKeyType = "ecdsa-sha2-nistp384"
	SSHKeyTypeEcdsaSha2Nistp521             SSHKeyType = "ecdsa-sha2-nistp521"
	SSHKeyTypeSkSSHEd25519OpensshCom        SSHKeyType = "sk-ssh-ed25519@openssh.com"
	SSHKeyTypeSkEcdsaSha2Nistp256OpensshCom SSHKeyType = "sk-ecdsa-sha2-nistp256@openssh.com"
)

// SSHSignAlgorithm is the `SshSignAlgorithm` schema.
//
// Spec schema: `SshSignAlgorithm`.
type SSHSignAlgorithm = string

// The values SSHSignAlgorithm takes.
const (
	SSHSignAlgorithmSSHEd25519        SSHSignAlgorithm = "ssh-ed25519"
	SSHSignAlgorithmSSHRsa            SSHSignAlgorithm = "ssh-rsa"
	SSHSignAlgorithmRsaSha2256        SSHSignAlgorithm = "rsa-sha2-256"
	SSHSignAlgorithmRsaSha2512        SSHSignAlgorithm = "rsa-sha2-512"
	SSHSignAlgorithmEcdsaSha2Nistp256 SSHSignAlgorithm = "ecdsa-sha2-nistp256"
	SSHSignAlgorithmEcdsaSha2Nistp384 SSHSignAlgorithm = "ecdsa-sha2-nistp384"
	SSHSignAlgorithmEcdsaSha2Nistp521 SSHSignAlgorithm = "ecdsa-sha2-nistp521"
)

// SSHSnippet is the `SshSnippet` schema.
//
// Spec schema: `SshSnippet`.
type SSHSnippet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Command     string  `json:"command"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// SSHSnippetInput is the `SshSnippetInput` schema.
//
// Spec schema: `SshSnippetInput`.
type SSHSnippetInput struct {
	Name        string  `json:"name"`
	Command     string  `json:"command"`
	Description *string `json:"description,omitempty"`
}

// SSHTunnelCreateAccountRequest is the `SshTunnelCreateAccountRequest` schema.
//
// Spec schema: `SshTunnelCreateAccountRequest`.
type SSHTunnelCreateAccountRequest struct {
	SSHHost     string            `json:"sshHost"`
	SSHPort     int64             `json:"sshPort"`
	SSHUser     string            `json:"sshUser"`
	SSHKeyID    string            `json:"sshKeyId"`
	RemoteHost  string            `json:"remoteHost"`
	RemotePort  int64             `json:"remotePort"`
	PluginID    string            `json:"pluginId"`
	DisplayName string            `json:"displayName"`
	Credentials map[string]string `json:"credentials"`
}

// SSHTunnelCreateAccountResponse is the `SshTunnelCreateAccountResponse` schema.
//
// Spec schema: `SshTunnelCreateAccountResponse`.
type SSHTunnelCreateAccountResponse struct {
	AccountID string `json:"accountId"`
}

// StatusDot is the `StatusDot` schema.
type StatusDot struct {
	// Kind: One of "status-dot".
	Kind   string         `json:"kind"`
	Status ResourceStatus `json:"status"`
	Label  *string        `json:"label,omitempty"`
}

// StatusHistoryDay is the `StatusHistoryDay` schema.
type StatusHistoryDay struct {
	// Day: `YYYY-MM-DD`, UTC.
	Day string `json:"day"`
	// Uptime: Fraction of the day the endpoint was up (0–1), or null when
	// nothing was recorded.
	Uptime *float64 `json:"uptime"`
}

// StatusPage is the `StatusPage` schema.
type StatusPage struct {
	ID string `json:"id"`
	// Slug: The public URL segment, and the page's only access credential.
	// Generated with real entropy rather than derived from the title.
	Slug        string  `json:"slug"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	// Published: False until deliberately published; a fresh page is never
	// reachable.
	Published   bool                  `json:"published"`
	ShowHistory bool                  `json:"showHistory"`
	ShowUptime  bool                  `json:"showUptime"`
	SupportURL  *string               `json:"supportUrl"`
	Components  []StatusPageComponent `json:"components"`
	CreatedAt   string                `json:"createdAt"`
	UpdatedAt   string                `json:"updatedAt"`
}

// StatusPageComponent is the `StatusPageComponent` schema.
type StatusPageComponent struct {
	ID      string `json:"id"`
	ProbeID string `json:"probeId"`
	// Label: Public name; null falls back to the probe's own name.
	Label     *string `json:"label"`
	GroupName *string `json:"groupName"`
	// Position: Ascending display order.
	Position int64 `json:"position"`
	// ProbeName: The probe's internal name — editor-only.
	ProbeName string `json:"probeName"`
	// ProbeStatus: One of "up", "down", "unknown".
	ProbeStatus string `json:"probeStatus"`
	// ProbeEnabled: False when the probe is paused.
	ProbeEnabled bool `json:"probeEnabled"`
}

// StatusPageComponentInput is the `StatusPageComponentInput` schema.
type StatusPageComponentInput struct {
	ProbeID   string  `json:"probeId"`
	Label     *string `json:"label,omitempty"`
	GroupName *string `json:"groupName,omitempty"`
}

// StatusPageCreate is the `StatusPageCreate` schema.
type StatusPageCreate struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	// Published: Defaults to false.
	Published   *bool   `json:"published,omitempty"`
	ShowHistory *bool   `json:"showHistory,omitempty"`
	ShowUptime  *bool   `json:"showUptime,omitempty"`
	SupportURL  *string `json:"supportUrl,omitempty"`
	// Components: Order is significant — it is the public render order.
	Components []StatusPageComponentInput `json:"components,omitempty"`
}

// StatusPageListResponse is the `StatusPageListResponse` schema.
type StatusPageListResponse struct {
	Pages []StatusPage `json:"pages"`
}

// StatusPagePatch is the `StatusPagePatch` schema.
type StatusPagePatch struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Published   *bool   `json:"published,omitempty"`
	ShowHistory *bool   `json:"showHistory,omitempty"`
	ShowUptime  *bool   `json:"showUptime,omitempty"`
	SupportURL  *string `json:"supportUrl,omitempty"`
	// Components: When present, replaces the whole set.
	Components []StatusPageComponentInput `json:"components,omitempty"`
}

// StorageListRequest is the `StorageListRequest` schema.
type StorageListRequest struct {
	AccountID string `json:"accountId"`
	Bucket    string `json:"bucket"`
	Prefix    string `json:"prefix"`
}

// StorageObject is the `StorageObject` schema.
type StorageObject struct {
	// Key: Full path within the bucket.
	Key string `json:"key"`
	// Name: Last path segment — what the browser renders.
	Name         string  `json:"name"`
	Size         float64 `json:"size"`
	LastModified string  `json:"lastModified"`
	IsDirectory  bool    `json:"isDirectory"`
	ContentType  *string `json:"contentType,omitempty"`
}

// StoragePathRequest is the `StoragePathRequest` schema.
type StoragePathRequest struct {
	AccountID string `json:"accountId"`
	Bucket    string `json:"bucket"`
	Key       string `json:"key"`
}

// StorageUploadForm is the `StorageUploadForm` schema.
type StorageUploadForm struct {
	AccountID string `json:"accountId"`
	Bucket    string `json:"bucket"`
	Key       string `json:"key"`
	// File: Raw file bytes
	File io.Reader `json:"file"`
}

// StripeRedirectURL is the `StripeRedirectUrl` schema.
//
// Spec schema: `StripeRedirectUrl`.
type StripeRedirectURL struct {
	URL string `json:"url"`
}

// Subscription is the `Subscription` schema.
//
// The API may send null in its place.
type Subscription struct {
	// Status: One of "trialing", "active", "past_due", "canceled", "unpaid".
	Status           string  `json:"status"`
	SeatCount        int64   `json:"seatCount"`
	CurrentPeriodEnd *string `json:"currentPeriodEnd"`
	StripeCustomerID string  `json:"stripeCustomerId"`
}

// SwapAllocationRulesBody: Two allocation rule ids in the same org whose
// priorities should be swapped.
type SwapAllocationRulesBody struct {
	AID string `json:"aId"`
	BID string `json:"bId"`
}

// SyncResponse is the `SyncResponse` schema.
type SyncResponse struct {
	Synced int64 `json:"synced"`
}

// SyncedResource is the `SyncedResource` schema.
type SyncedResource struct {
	ID               ResourceID  `json:"id"`
	PluginID         string      `json:"pluginId"`
	ResourceTypeID   string      `json:"resourceTypeId"`
	DisplayName      string      `json:"displayName"`
	ExternalID       *string     `json:"externalId"`
	FieldsJSON       JSONObject  `json:"fieldsJson"`
	OutputsJSON      JSONObject  `json:"outputsJson"`
	ParentResourceID *ResourceID `json:"parentResourceId"`
}

// SyntheticProbe is the `SyntheticProbe` schema.
type SyntheticProbe struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// URL: Absolute http(s) URL the check hits from the edge proxy.
	URL string `json:"url"`
	// Method: HTTP method the probe uses — GET, HEAD or OPTIONS. Unknown values
	// become GET.
	Method string `json:"method"`
	// IntervalSeconds: Seconds between checks. Clamped server-side to 60–86400.
	IntervalSeconds int64 `json:"intervalSeconds"`
	// TimeoutMs: Per-check timeout in milliseconds. Clamped server-side to
	// 1000–60000.
	TimeoutMs int64 `json:"timeoutMs"`
	// FailureThreshold: Consecutive failures before the probe flips to `down`
	// and notifies. Clamped 1–20.
	FailureThreshold int64 `json:"failureThreshold"`
	Enabled          bool  `json:"enabled"`
	// AccountID: Account of the linked resource, when the URL came from one.
	AccountID *string `json:"accountId"`
	// ResourceID: Linked resource id; advisory, not a foreign key.
	ResourceID     *string   `json:"resourceId"`
	PluginID       *PluginID `json:"pluginId"`
	ResourceTypeID *string   `json:"resourceTypeId"`
	// OutputKey: The resource output/field key the URL was suggested from.
	OutputKey *string `json:"outputKey"`
	// Status: The probe's state machine: `unknown` until the first result,
	// `down` after `failureThreshold` consecutive failures, `up` on any success.
	//
	// One of "up", "down", "unknown".
	Status              string  `json:"status"`
	ConsecutiveFailures int64   `json:"consecutiveFailures"`
	LastProbeAt         *string `json:"lastProbeAt"`
	LastStatusCode      *int64  `json:"lastStatusCode"`
	LastLatencyMs       *int64  `json:"lastLatencyMs"`
	// LastError: Failure detail; null after a success.
	LastError *string `json:"lastError"`
	// LastStateChangeAt: When status last flipped up/down.
	LastStateChangeAt *string `json:"lastStateChangeAt"`
	// Uptime24h: Fraction (0–1) of the trailing 24h the endpoint was up, from
	// the recorded series; null before the first result lands in the metric
	// store.
	Uptime24h *float64 `json:"uptime24h"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

// SyntheticProbeCreate is the `SyntheticProbeCreate` schema.
type SyntheticProbeCreate struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	// Method: HTTP method the probe uses — GET, HEAD or OPTIONS. Unknown values
	// become GET.
	Method *string `json:"method,omitempty"`
	// IntervalSeconds: Seconds between checks. Clamped server-side to 60–86400.
	IntervalSeconds *int64 `json:"intervalSeconds,omitempty"`
	// TimeoutMs: Per-check timeout in milliseconds. Clamped server-side to
	// 1000–60000.
	TimeoutMs *int64 `json:"timeoutMs,omitempty"`
	// FailureThreshold: Consecutive failures before the probe flips to `down`
	// and notifies. Clamped 1–20.
	FailureThreshold *int64 `json:"failureThreshold,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
	// ResourceID: Link the probe to the resource whose output suggested the URL.
	ResourceID *string `json:"resourceId,omitempty"`
	OutputKey  *string `json:"outputKey,omitempty"`
}

// SyntheticProbeList is the `SyntheticProbeList` schema.
type SyntheticProbeList struct {
	Probes []SyntheticProbe `json:"probes"`
}

// SyntheticProbeUpdate is the `SyntheticProbeUpdate` schema.
type SyntheticProbeUpdate struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
	// Method: HTTP method the probe uses — GET, HEAD or OPTIONS. Unknown values
	// become GET.
	Method *string `json:"method,omitempty"`
	// IntervalSeconds: Seconds between checks. Clamped server-side to 60–86400.
	IntervalSeconds *int64 `json:"intervalSeconds,omitempty"`
	// TimeoutMs: Per-check timeout in milliseconds. Clamped server-side to
	// 1000–60000.
	TimeoutMs *int64 `json:"timeoutMs,omitempty"`
	// FailureThreshold: Consecutive failures before the probe flips to `down`
	// and notifies. Clamped 1–20.
	FailureThreshold *int64 `json:"failureThreshold,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
}

// TabTarget is the `TabTarget` schema.
type TabTarget struct {
	// Kind: One of "dashboard", "account", "resource", "agents", "costs",
	// "savings", "cost-reports", "invoices", "graph", "logs", "changes",
	// "expiring", "posture", "access-review", "backups", "wallboard",
	// "calendar", "runbooks", "query-monitors", "dns", "iac",
	// "environment-diff", "environments", "ssh-fanout", "metric-alerts",
	// "probes", "status-pages", "quotas", "incidents", "workflows",
	// "deployments", "settings", "chat", "linux-app".
	Kind           string      `json:"kind"`
	DashboardID    *string     `json:"dashboardId,omitempty"`
	AccountID      *string     `json:"accountId,omitempty"`
	ResourceID     *ResourceID `json:"resourceId,omitempty"`
	ConversationID *string     `json:"conversationId,omitempty"`
	ReportID       *string     `json:"reportId,omitempty"`
	InvoiceID      *string     `json:"invoiceId,omitempty"`
	SessionID      *string     `json:"sessionId,omitempty"`
	WindowID       *int64      `json:"windowId,omitempty"`
	AppID          *string     `json:"appId,omitempty"`
}

// TagComplianceReport is the `TagComplianceReport` schema.
type TagComplianceReport struct {
	Policy   TagPolicy              `json:"policy"`
	Accounts []AccountTagCompliance `json:"accounts"`
}

// TagPolicy is the `TagPolicy` schema.
type TagPolicy struct {
	RequiredTags []RequiredTag `json:"requiredTags"`
	// EnforceOnCreate: When true, resource creation is rejected with a 422
	// (`tag_policy_unmet`) if the submitted fields carry a tag map missing a
	// required tag. Types whose create form has no `tags`/`labels` field are
	// exempt.
	EnforceOnCreate bool `json:"enforceOnCreate"`
}

// TagPolicyBlocked is the `TagPolicyBlocked` schema.
type TagPolicyBlocked struct {
	Error string `json:"error"`
	// Code: One of "tag_policy_unmet".
	Code         string               `json:"code"`
	Violations   []TagPolicyViolation `json:"violations"`
	RequiredTags []RequiredTag        `json:"requiredTags"`
}

// TagPolicyViolation is the `TagPolicyViolation` schema.
type TagPolicyViolation struct {
	Key string `json:"key"`
	// Reason: One of "missing", "value_not_allowed".
	Reason        string   `json:"reason"`
	Value         *string  `json:"value,omitempty"`
	AllowedValues []string `json:"allowedValues,omitempty"`
}

// TerraformExport is the `TerraformExport` schema.
type TerraformExport struct {
	Hcl         string                       `json:"hcl"`
	Exported    []TerraformExportExported    `json:"exported"`
	Unsupported []TerraformExportUnsupported `json:"unsupported"`
}

// TOTPEnrollment is the `TotpEnrollment` schema.
//
// Spec schema: `TotpEnrollment`.
type TOTPEnrollment struct {
	FactorID    string `json:"factorId"`
	ChallengeID string `json:"challengeId"`
	// QrCode: Data-URI image of the enrolment QR code
	QrCode *string `json:"qrCode"`
	// Secret: Base32 secret, for manual entry
	Secret *string `json:"secret"`
	// URI: `otpauth://` URI
	URI *string `json:"uri"`
}

// UnitCostPoint is the `UnitCostPoint` schema.
type UnitCostPoint struct {
	// Bucket: Bucket start date, YYYY-MM-DD.
	Bucket string `json:"bucket"`
	// Value: The ratio, or **null for a gap**. Never 0 and never infinite: a
	// bucket with no reported metric value is unknown, not free, and rendering
	// it as 0 would say the opposite of the truth. A zero numerator over a
	// positive denominator is a real 0 and is returned as one.
	Value *float64 `json:"value"`
	// Cost: Spend summed over the bucket, in the series' currency.
	Cost float64 `json:"cost"`
	// MetricValue: Metric value summed over the bucket, or null when nothing was
	// reported.
	MetricValue *float64 `json:"metricValue"`
	// Gap: Set exactly when `value` is null.
	//
	// One of "no_metric_value", "non_positive_metric_value",
	// "unconvertible_currency".
	Gap *string `json:"gap,omitempty"`
	// ReportedDays: Days in the bucket carrying a reported value, out of
	// `bucketDays`. When it is smaller, the denominator covers only part of the
	// bucket and the ratio there reads high.
	ReportedDays int64 `json:"reportedDays"`
	BucketDays   int64 `json:"bucketDays"`
}

// UnitCostQueryRequest is the `UnitCostQueryRequest` schema.
type UnitCostQueryRequest struct {
	// From: Inclusive, YYYY-MM-DD.
	From string `json:"from"`
	To   string `json:"to"`
	// Binning: One of "daily", "weekly", "monthly", "cumulative".
	Binning string `json:"binning"`
	// Mode: Absent is `unit_cost` (spend ÷ metric value). `margin` is `(revenue
	// − spend) ÷ revenue` as a fraction, and is a 400 for a metric whose `kind`
	// is not `currency`.
	//
	// One of "unit_cost", "margin".
	Mode *string `json:"mode,omitempty"`
	// Filters: Narrowing on top of the metric's own `costScope` — AND-composed,
	// never a replacement.
	Filters []BusinessMetricScopeTerm `json:"filters,omitempty"`
	// Query: The same narrowing as cost-query-language text.
	Query         *string `json:"query,omitempty"`
	SavedFilterID *string `json:"savedFilterId,omitempty"`
	// CostBasis: One of "cash", "amortized".
	CostBasis   *string  `json:"costBasis,omitempty"`
	ChargeTypes []string `json:"chargeTypes,omitempty"`
	// DisplayCurrency: Fold spend currencies the organization holds a rate for
	// into this one before dividing. Ignored for `margin`, which always converts
	// to the metric's own currency.
	DisplayCurrency *string `json:"displayCurrency,omitempty"`
}

// UnitCostQueryResponse is the `UnitCostQueryResponse` schema.
type UnitCostQueryResponse struct {
	Metric UnitCostQueryResponseMetric `json:"metric"`
	// Mode: One of "unit_cost", "margin".
	Mode string `json:"mode"`
	// Binning: One of "daily", "weekly", "monthly", "cumulative".
	Binning string `json:"binning"`
	// Series: One series per currency the numerator ended up in — usually one.
	// More than one means the organization has spend in a currency it holds no
	// rate for; rather than dropping that spend (understating every unit cost)
	// or adding it to another currency (inventing a number), each currency
	// divides the same denominator on its own.
	Series []UnitCostSeries `json:"series"`
	// Conversion: Set only when spend currencies were folded together; absent
	// means untouched.
	Conversion *UnitCostQueryResponseConversion `json:"conversion,omitempty"`
	// GapBuckets: Buckets on the axis that produced no ratio at all.
	GapBuckets int64 `json:"gapBuckets"`
	// PartialBuckets: Buckets whose denominator covers only part of the bucket.
	PartialBuckets int64 `json:"partialBuckets"`
}

// UnitCostSeries is the `UnitCostSeries` schema.
type UnitCostSeries struct {
	Currency string          `json:"currency"`
	Points   []UnitCostPoint `json:"points"`
	// OverallValue: The period ratio: **summed numerator ÷ summed denominator**,
	// not the mean of the per-bucket ratios — the mean weights a quiet Sunday
	// exactly as heavily as a peak Monday. Only buckets that produced a ratio
	// contribute, on both sides.
	OverallValue       *float64 `json:"overallValue"`
	OverallCost        float64  `json:"overallCost"`
	OverallMetricValue *float64 `json:"overallMetricValue"`
}

// UnpinRequest is the `UnpinRequest` schema.
type UnpinRequest struct {
	DashboardID string     `json:"dashboardId"`
	ResourceID  ResourceID `json:"resourceId"`
}

// UntaggedSpendReport is the `UntaggedSpendReport` schema.
type UntaggedSpendReport struct {
	From         string   `json:"from"`
	To           string   `json:"to"`
	RequiredKeys []string `json:"requiredKeys"`
	Currencies   []string `json:"currencies"`
	// Totals: Currency code → amount in the currency's major unit.
	Totals map[string]float64 `json:"totals"`
	// UntaggedTotals: Spend on rows missing at least one required tag key, per
	// currency.
	UntaggedTotals map[string]float64               `json:"untaggedTotals"`
	ByKey          []UntaggedSpendReportByKey       `json:"byKey"`
	TopUntagged    []UntaggedSpendReportTopUntagged `json:"topUntagged"`
}

// UpdateAccountRequest is the `UpdateAccountRequest` schema.
type UpdateAccountRequest struct {
	DisplayName *string `json:"displayName,omitempty"`
	// BastionID: Pass `null` to unbind, a uuid to bind, or omit the field to
	// leave the binding unchanged.
	BastionID *string `json:"bastionId,omitempty"`
}

// UpdateResourceRequest is the `UpdateResourceRequest` schema.
type UpdateResourceRequest struct {
	AccountID        string            `json:"accountId"`
	PluginID         string            `json:"pluginId"`
	ResourceTypeID   string            `json:"resourceTypeId"`
	ResourceID       ResourceID        `json:"resourceId"`
	Fields           map[string]string `json:"fields"`
	ParentResourceID *ResourceID       `json:"parentResourceId,omitempty"`
}

// UpdateResourceResponse is the `UpdateResourceResponse` schema.
type UpdateResourceResponse struct {
	ID          ResourceID        `json:"id"`
	DisplayName string            `json:"displayName"`
	Fields      map[string]string `json:"fields"`
}

// UpdateWidgetRequest is the `UpdateWidgetRequest` schema.
type UpdateWidgetRequest struct {
	Title  *string    `json:"title,omitempty"`
	Config JSONObject `json:"config,omitempty"`
	GridX  *int64     `json:"gridX,omitempty"`
	GridY  *int64     `json:"gridY,omitempty"`
	GridW  *int64     `json:"gridW,omitempty"`
	GridH  *int64     `json:"gridH,omitempty"`
}

// UpdatedAccount is the `UpdatedAccount` schema.
type UpdatedAccount struct {
	ID          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	BastionID   *string `json:"bastionId"`
}

// UserSession is the `UserSession` schema.
type UserSession struct {
	ID         string  `json:"id"`
	IPAddress  *string `json:"ipAddress"`
	UserAgent  *string `json:"userAgent"`
	AuthMethod string  `json:"authMethod"`
	Status     string  `json:"status"`
	ExpiresAt  string  `json:"expiresAt"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	// Current: True for the session making this request
	Current bool `json:"current"`
}

// ValidateTabsRequest is the `ValidateTabsRequest` schema.
type ValidateTabsRequest struct {
	Tabs []ValidateTabsRequestTabs `json:"tabs"`
}

// ValidateTabsResponse is the `ValidateTabsResponse` schema.
type ValidateTabsResponse struct {
	ValidTabIDs []string `json:"validTabIds"`
}

// WallboardFailureLine is the `WallboardFailureLine` schema.
type WallboardFailureLine struct {
	ID     string  `json:"id"`
	Label  string  `json:"label"`
	Detail string  `json:"detail"`
	Since  *string `json:"since"`
}

// WallboardIncidentLine is the `WallboardIncidentLine` schema.
type WallboardIncidentLine struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Severity  string `json:"severity"`
	StartedAt string `json:"startedAt"`
	Status    string `json:"status"`
}

// WallboardResponse is the `WallboardResponse` schema.
type WallboardResponse struct {
	// Status: Three states rather than five, because at four metres a person
	// distinguishes three colours reliably and nothing more. `down` is reserved
	// for the two things that mean customers are affected now — a sev1 incident
	// or a probe that is down; everything else that is wrong is `degraded`. A
	// source that could not be read is `degraded` and never `ok`.
	//
	// One of "ok", "degraded", "down".
	Status string          `json:"status"`
	Tiles  []WallboardTile `json:"tiles"`
	// Incidents: Unresolved incidents, newest first.
	Incidents []WallboardIncidentLine `json:"incidents"`
	// Failures: Probes that are down, query monitors breaching or unable to run,
	// accounts that stopped syncing.
	Failures []WallboardFailureLine `json:"failures"`
	// FailedSources: Sources that could not be read, **named on the screen**. A
	// wallboard showing green because a query failed is worse than a blank one —
	// it is actively telling the room the wrong thing.
	FailedSources []string `json:"failedSources"`
	GeneratedAt   string   `json:"generatedAt"`
}

// WallboardTile is the `WallboardTile` schema.
type WallboardTile struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	// Value: The number or short phrase, rendered in large type.
	Value  string  `json:"value"`
	Detail *string `json:"detail"`
	// Status: Three states rather than five, because at four metres a person
	// distinguishes three colours reliably and nothing more. `down` is reserved
	// for the two things that mean customers are affected now — a sev1 incident
	// or a probe that is down; everything else that is wrong is `degraded`. A
	// source that could not be read is `degraded` and never `ok`.
	//
	// One of "ok", "degraded", "down".
	Status string `json:"status"`
}

// WorkflowApproval is the `WorkflowApproval` schema.
type WorkflowApproval struct {
	ID            string                 `json:"id"`
	WorkflowID    string                 `json:"workflowId"`
	WorkflowName  *string                `json:"workflowName"`
	RunID         string                 `json:"runId"`
	Title         string                 `json:"title"`
	Message       string                 `json:"message"`
	Status        WorkflowApprovalStatus `json:"status"`
	ExpiresAt     string                 `json:"expiresAt"`
	DecidedAt     *string                `json:"decidedAt"`
	DecidedByName *string                `json:"decidedByName"`
	CreatedAt     string                 `json:"createdAt"`
}

// WorkflowApprovalStatus is the `WorkflowApprovalStatus` schema.
type WorkflowApprovalStatus = string

// The values WorkflowApprovalStatus takes.
const (
	WorkflowApprovalStatusPending  WorkflowApprovalStatus = "pending"
	WorkflowApprovalStatusApproved WorkflowApprovalStatus = "approved"
	WorkflowApprovalStatusDenied   WorkflowApprovalStatus = "denied"
	WorkflowApprovalStatusExpired  WorkflowApprovalStatus = "expired"
)

// WorkflowPinRequest is the `WorkflowPinRequest` schema.
type WorkflowPinRequest struct {
	DashboardID string `json:"dashboardId"`
	WorkflowID  string `json:"workflowId"`
}

// WorkflowSchedule is the `WorkflowSchedule` schema.
type WorkflowSchedule struct {
	// Expression: Standard 5-field cron expression (minute hour day-of-month
	// month day-of-week). Supports `*`, lists, ranges, and steps; 3-letter
	// month/weekday names; `7` as Sunday. When both day fields are restricted, a
	// date matches if either does (POSIX).
	Expression string `json:"expression"`
	// Timezone: IANA timezone the expression's wall times are evaluated in. Omit
	// or null for UTC.
	Timezone *string `json:"timezone"`
	// Enabled: Mirrors the workflow's enabled flag — a disabled workflow's
	// schedule never fires.
	Enabled bool `json:"enabled"`
	// LastRunAt: When the workflow last finished a run (any trigger source).
	LastRunAt *string `json:"lastRunAt"`
	// NextRunAt: The persisted next fire time the scheduler will claim. Null
	// while disabled, or when the expression never matches.
	NextRunAt *string `json:"nextRunAt"`
	// NextRuns: Preview of the next few fire times, computed at read time.
	NextRuns []string `json:"nextRuns"`
}

// WorkflowScheduleInput is the `WorkflowScheduleInput` schema.
type WorkflowScheduleInput struct {
	// Expression: Standard 5-field cron expression (minute hour day-of-month
	// month day-of-week). Supports `*`, lists, ranges, and steps; 3-letter
	// month/weekday names; `7` as Sunday. When both day fields are restricted, a
	// date matches if either does (POSIX).
	Expression string `json:"expression"`
	// Timezone: IANA timezone the expression's wall times are evaluated in. Omit
	// or null for UTC.
	Timezone *string `json:"timezone,omitempty"`
	// Enabled: Also set the workflow's enabled flag. Omit to leave it unchanged.
	Enabled *bool `json:"enabled,omitempty"`
}

// WorkflowScheduleResponse is the `WorkflowScheduleResponse` schema.
type WorkflowScheduleResponse struct {
	// Schedule: Null when the workflow's trigger is not cron.
	Schedule *WorkflowSchedule `json:"schedule"`
}

// WorkflowSecret is the `WorkflowSecret` schema.
type WorkflowSecret struct {
	ID string `json:"id"`
	// Name: JavaScript dot identifier used to expose the value to workflow code,
	// for example `API_TOKEN` or `stripe.apiKey`.
	Name        string  `json:"name"`
	Description *string `json:"description"`
	// HasValue: Whether an encrypted value is stored. The value is never
	// returned.
	HasValue  bool   `json:"hasValue"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// WorkflowSecretAssignment is the `WorkflowSecretAssignment` schema.
type WorkflowSecretAssignment struct {
	SecretIDs []string         `json:"secretIds"`
	Secrets   []WorkflowSecret `json:"secrets"`
}

// WorkflowSecretAssignmentInput is the `WorkflowSecretAssignmentInput` schema.
type WorkflowSecretAssignmentInput struct {
	SecretIDs []string `json:"secretIds"`
}

// WorkflowSecretCreate is the `WorkflowSecretCreate` schema.
type WorkflowSecretCreate struct {
	// Name: JavaScript dot identifier used to expose the value to workflow code,
	// for example `API_TOKEN` or `stripe.apiKey`.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// WorkflowSecretUpdate is the `WorkflowSecretUpdate` schema.
type WorkflowSecretUpdate struct {
	// Name: JavaScript dot identifier used to expose the value to workflow code,
	// for example `API_TOKEN` or `stripe.apiKey`.
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// WorkflowSecretValueWrite is the `WorkflowSecretValueWrite` schema.
type WorkflowSecretValueWrite struct {
	// Value: Write-only plaintext. It is AES-256-GCM encrypted before storage
	// and is never returned.
	Value string `json:"value"`
}

// WorkflowTypingsResponse is the `WorkflowTypingsResponse` schema.
type WorkflowTypingsResponse struct {
	// Dts: Ambient TypeScript declarations for this workflow's `infra` API — the
	// same file the Monaco editor and `check` endpoint type against.
	Dts string `json:"dts"`
}

// AcceptInvitationResponseOrganization is an object the spec declares inline.
type AcceptInvitationResponseOrganization struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

// AccountDetailAccount is an object the spec declares inline.
type AccountDetailAccount struct {
	ID          string `json:"id"`
	PluginID    string `json:"pluginId"`
	DisplayName string `json:"displayName"`
}

// AlertRulesResponseSlackChannels is an object the spec declares inline.
type AlertRulesResponseSlackChannels struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"isPrivate"`
}

// AlertRulesResponseMsTeamsWebhooks is an object the spec declares inline.
type AlertRulesResponseMsTeamsWebhooks struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// AlertRulesResponseAccounts is an object the spec declares inline.
type AlertRulesResponseAccounts struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	PluginID    string `json:"pluginId"`
}

// AlertRulesResponseOnCallSchedules is an object the spec declares inline.
type AlertRulesResponseOnCallSchedules struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BlastRadiusDependantVia is an object the spec declares inline.
type BlastRadiusDependantVia struct {
	// FieldKey: The dependant's field holding the reference.
	FieldKey string `json:"fieldKey"`
	// OutputKey: The output or identity the reference reads.
	OutputKey string `json:"outputKey"`
	// Kind: Where the edge came from. Absent means `output-ref` — a reference
	// wired by hand.
	//
	// One of "output-ref", "declared", "containment", "field-match".
	Kind *string `json:"kind,omitempty"`
	// Label: How the plugin words the relationship ("in VPC"), when it declared
	// one.
	Label *string `json:"label,omitempty"`
}

// BlastRadiusReportFlowTotals is an object the spec declares inline.
type BlastRadiusReportFlowTotals struct {
	Bytes         float64 `json:"bytes"`
	EstimatedCost float64 `json:"estimatedCost"`
	Currency      string  `json:"currency"`
}

// BudgetWithStatusCurrentMonthEvents is an object the spec declares inline.
type BudgetWithStatusCurrentMonthEvents struct {
	ID string `json:"id"`
	// ThresholdType: One of "actual", "forecast".
	ThresholdType    string `json:"thresholdType"`
	ThresholdPercent int64  `json:"thresholdPercent"`
	TriggeredAt      string `json:"triggeredAt"`
}

// BudgetWithStatusPlacements is an object the spec declares inline.
type BudgetWithStatusPlacements struct {
	WidgetID      string `json:"widgetId"`
	DashboardID   string `json:"dashboardId"`
	DashboardName string `json:"dashboardName"`
}

// BusinessMetricValuesInputValues is an object the spec declares inline.
type BusinessMetricValuesInputValues struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

// ChangeFreezeBlockedFreeze is an object the spec declares inline.
type ChangeFreezeBlockedFreeze struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Reason   *string `json:"reason"`
	StartsAt string  `json:"startsAt"`
	EndsAt   *string `json:"endsAt"`
}

// CostAccountStatusCostPollError is an object the spec declares inline.
type CostAccountStatusCostPollError struct {
	Message  string                                  `json:"message"`
	HelpLink *CostAccountStatusCostPollErrorHelpLink `json:"helpLink"`
}

// CostAccountStatusCoverage is an object the spec declares inline.
type CostAccountStatusCoverage struct {
	FirstDay string `json:"firstDay"`
	LastDay  string `json:"lastDay"`
}

// CostAdjustmentSummaryRules is an object the spec declares inline.
type CostAdjustmentSummaryRules struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Kind: One of "percentage", "fixed", "reallocation".
	Kind    string `json:"kind"`
	Summary string `json:"summary"`
}

// CostAnomalyAcknowledgement is an object the spec declares inline.
type CostAnomalyAcknowledgement struct {
	// Explanation: What somebody established this finding was. Also the
	// annotation's text.
	Explanation string `json:"explanation"`
	// AcknowledgedAt: When the current explanation was recorded — restamped by a
	// correction.
	AcknowledgedAt       string  `json:"acknowledgedAt"`
	AcknowledgedByUserID *string `json:"acknowledgedByUserId"`
	// AnnotationID: The cost annotation this created, drawn on every chart
	// covering the anomalous day. Null once that note has been deleted — which
	// removes the marker, never the acknowledgement: the finding stays
	// explained.
	AnnotationID *string `json:"annotationId"`
}

// CostScenarioResultContributions is an object the spec declares inline.
type CostScenarioResultContributions struct {
	AdjustmentID string `json:"adjustmentId"`
	Label        string `json:"label"`
	// Kind: One of "one_off", "recurring", "rate_change".
	Kind   string  `json:"kind"`
	Amount float64 `json:"amount"`
}

// CreateAccountResponseSyncError is an object the spec declares inline.
type CreateAccountResponseSyncError struct {
	Message string `json:"message"`
}

// CreateJiraIssueResultIssue is an object the spec declares inline.
type CreateJiraIssueResultIssue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
	URL string `json:"url"`
}

// CreateLinearIssueResultIssue is an object the spec declares inline.
type CreateLinearIssueResultIssue struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	URL        string `json:"url"`
}

// CreatePricingRequestSizes is an object the spec declares inline.
type CreatePricingRequestSizes struct {
	ID       string  `json:"id"`
	Vcpus    float64 `json:"vcpus"`
	MemoryMb float64 `json:"memoryMb"`
}

// CredentialExportFields is an object the spec declares inline.
type CredentialExportFields struct {
	Label     string  `json:"label"`
	Value     string  `json:"value"`
	Sensitive *bool   `json:"sensitive,omitempty"`
	Hint      *string `json:"hint,omitempty"`
}

// CredentialFieldHelpLink is an object the spec declares inline.
type CredentialFieldHelpLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// CustomGraphCheckResultDiagnostics is an object the spec declares inline.
type CustomGraphCheckResultDiagnostics struct {
	Line     int64  `json:"line"`
	Column   int64  `json:"column"`
	Code     int64  `json:"code"`
	Category string `json:"category"`
	Message  string `json:"message"`
}

// CustomGraphRenderResultLogs is an object the spec declares inline.
type CustomGraphRenderResultLogs struct {
	// Level: One of "info", "warn", "error".
	Level   string `json:"level"`
	Message string `json:"message"`
}

// DashboardWorkflowPinMetrics is an object the spec declares inline.
type DashboardWorkflowPinMetrics struct {
	Key   string  `json:"key"`
	Label string  `json:"label"`
	Unit  *string `json:"unit"`
	Value any     `json:"value,omitempty"`
}

// DeployCreatedResourceSidecar is an object the spec declares inline.
type DeployCreatedResourceSidecar struct {
	PluginID         string `json:"pluginId"`
	ParentResourceID string `json:"parentResourceId"`
}

// DeployPlanResultResult is an object the spec declares inline.
type DeployPlanResultResult struct {
	Status           DeployStatus                 `json:"status"`
	Env              string                       `json:"env"`
	Plan             any                          `json:"plan,omitempty"`
	Dockerfile       *string                      `json:"dockerfile,omitempty"`
	Image            *string                      `json:"image,omitempty"`
	Notes            []string                     `json:"notes"`
	CreatedResources []DeployCreatedResource      `json:"createdResources"`
	PlannedChanges   []DeployPlannedChange        `json:"plannedChanges"`
	Logs             []DeployRunLog               `json:"logs"`
	ReachedStage     *DeployStage                 `json:"reachedStage,omitempty"`
	Error            *DeployPlanResultResultError `json:"error,omitempty"`
	DurationMs       int64                        `json:"durationMs"`
}

// DeployPlannedChangeSidecar is an object the spec declares inline.
type DeployPlannedChangeSidecar struct {
	PluginID         string `json:"pluginId"`
	ParentResourceID string `json:"parentResourceId"`
}

// DeploymentCostImpactTotal is an object the spec declares inline.
type DeploymentCostImpactTotal struct {
	Currency    string  `json:"currency"`
	DeltaPerDay float64 `json:"deltaPerDay"`
}

// DeploymentRunInputError is an object the spec declares inline.
type DeploymentRunInputError struct {
	Message string `json:"message"`
}

// EnvironmentCaptureDraftSkipped is an object the spec declares inline.
type EnvironmentCaptureDraftSkipped struct {
	ResourceID  string `json:"resourceId"`
	DisplayName string `json:"displayName"`
	Reason      string `json:"reason"`
}

// EnvironmentCaptureDraftMemberFieldMetaValue is an object the spec declares
// inline.
type EnvironmentCaptureDraftMemberFieldMetaValue struct {
	Label           string                                               `json:"label"`
	Kind            string                                               `json:"kind"`
	Required        bool                                                 `json:"required"`
	Options         []EnvironmentCaptureDraftMemberFieldMetaValueOptions `json:"options,omitempty"`
	Parameterisable bool                                                 `json:"parameterisable"`
}

// EnvironmentCostEstimateMembers is an object the spec declares inline.
type EnvironmentCostEstimateMembers struct {
	MemberKey     string   `json:"memberKey"`
	DisplayName   string   `json:"displayName"`
	MonthlyAmount *float64 `json:"monthlyAmount"`
	Currency      *string  `json:"currency"`
}

// EnvironmentParameterOptions is an object the spec declares inline.
type EnvironmentParameterOptions struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// FieldActionResponseOption is an object the spec declares inline.
type FieldActionResponseOption struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// HygieneReportCounts is an object the spec declares inline.
type HygieneReportCounts struct {
	High   int64 `json:"high"`
	Medium int64 `json:"medium"`
	Low    int64 `json:"low"`
	Total  int64 `json:"total"`
}

// IacImportPlanResponseExported is an object the spec declares inline.
type IacImportPlanResponseExported struct {
	ResourceID string  `json:"resourceId"`
	Address    string  `json:"address"`
	ImportID   *string `json:"importId"`
}

// IacImportPlanResponseUnsupported is an object the spec declares inline.
type IacImportPlanResponseUnsupported struct {
	ResourceID  string `json:"resourceId"`
	DisplayName string `json:"displayName"`
	Reason      string `json:"reason"`
}

// IacReconciliationResponseSummary is an object the spec declares inline.
type IacReconciliationResponseSummary struct {
	InventoryTotal     int64 `json:"inventoryTotal"`
	Managed            int64 `json:"managed"`
	Drifted            int64 `json:"drifted"`
	Unmanaged          int64 `json:"unmanaged"`
	StateOnly          int64 `json:"stateOnly"`
	Undiffable         int64 `json:"undiffable"`
	StateResources     int64 `json:"stateResources"`
	DataSourcesIgnored int64 `json:"dataSourcesIgnored"`
}

// IacReconciliationResponseUnderivable is an object the spec declares inline.
type IacReconciliationResponseUnderivable struct {
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
	Reason         string   `json:"reason"`
}

// IacStateOnlyResourceCandidates is an object the spec declares inline.
type IacStateOnlyResourceCandidates struct {
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
}

// IncidentTimelineFeeds is an object the spec declares inline.
type IncidentTimelineFeeds struct {
	Feed string `json:"feed"`
	// Status: One of "ok", "omitted", "error".
	Status string  `json:"status"`
	Error  *string `json:"error,omitempty"`
}

// IncidentTimelineEntryLink is an object the spec declares inline.
type IncidentTimelineEntryLink struct {
	// Kind: One of "resource", "changes", "provider-incident", "costs",
	// "workflow-run", "deployment", "audit", "freeze", "expiring", "probe",
	// "metric-alert", "incident".
	Kind     string  `json:"kind"`
	ID       *string `json:"id,omitempty"`
	ParentID *string `json:"parentId,omitempty"`
	URL      *string `json:"url,omitempty"`
}

// InvoiceDerivationRates is an object the spec declares inline.
type InvoiceDerivationRates struct {
	Currency      string  `json:"currency"`
	Rate          float64 `json:"rate"`
	EffectiveFrom string  `json:"effectiveFrom"`
}

// InvoiceDerivationRules is an object the spec declares inline.
type InvoiceDerivationRules struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Kind: One of "percentage", "fixed", "reallocation".
	Kind    string `json:"kind"`
	Summary string `json:"summary"`
}

// InvoiceDerivationScope is an object the spec declares inline.
type InvoiceDerivationScope struct {
	CostCentres []InvoiceDerivationScopeCostCentres `json:"costCentres"`
	Accounts    []InvoiceDerivationScopeAccounts    `json:"accounts"`
}

// InvoiceVoidResponseReplacement is an object the spec declares inline.
type InvoiceVoidResponseReplacement struct {
	ID                 string `json:"id"`
	ManagedAccountID   string `json:"managedAccountId"`
	ManagedAccountName string `json:"managedAccountName"`
	// Number: `INV-2026-0001`. Null while draft — numbers are assigned at
	// approval so a deleted draft cannot leave a gap in the sequence.
	Number                *string           `json:"number"`
	Status                InvoiceStatus     `json:"status"`
	PeriodFrom            string            `json:"periodFrom"`
	PeriodTo              string            `json:"periodTo"`
	Currency              string            `json:"currency"`
	Totals                *InvoiceTotals    `json:"totals,omitempty"`
	IssuedAt              *string           `json:"issuedAt"`
	SentAt                *string           `json:"sentAt"`
	Delivery              *InvoiceDelivery  `json:"delivery"`
	VoidedAt              *string           `json:"voidedAt"`
	VoidReason            *string           `json:"voidReason"`
	SupersedesInvoiceID   *string           `json:"supersedesInvoiceId"`
	SupersededByInvoiceID *string           `json:"supersededByInvoiceId"`
	CreatedAt             string            `json:"createdAt"`
	UpdatedAt             string            `json:"updatedAt"`
	Notes                 *string           `json:"notes"`
	Lines                 []InvoiceLine     `json:"lines"`
	Derivation            InvoiceDerivation `json:"derivation"`
	// Live: True when the figures in this response were recomputed for it — true
	// for a draft, false for everything else. Say so: “these numbers will move”
	// and “these numbers are what we sent” are different claims about the same
	// fields.
	Live             bool    `json:"live"`
	ComputedAt       string  `json:"computedAt"`
	ApprovedByUserID *string `json:"approvedByUserId"`
	SentByUserID     *string `json:"sentByUserId"`
	VoidedByUserID   *string `json:"voidedByUserId"`
	CreatedByUserID  *string `json:"createdByUserId"`
}

// MetricAlertSelectorOptionsPlugins is an object the spec declares inline.
type MetricAlertSelectorOptionsPlugins struct {
	PluginID        string   `json:"pluginId"`
	ResourceTypeIDs []string `json:"resourceTypeIds"`
}

// MetricSeriesPoints is an object the spec declares inline.
type MetricSeriesPoints struct {
	// Timestamp: Unix epoch milliseconds.
	Timestamp float64 `json:"timestamp"`
	Value     float64 `json:"value"`
}

// NetworkFlowFeedRange is an object the spec declares inline.
type NetworkFlowFeedRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// NetworkFlowFeedTotals is an object the spec declares inline.
type NetworkFlowFeedTotals struct {
	Bytes             float64 `json:"bytes"`
	EstimatedCost     float64 `json:"estimatedCost"`
	Currency          string  `json:"currency"`
	UnattributedBytes float64 `json:"unattributedBytes"`
	TruncatedBytes    float64 `json:"truncatedBytes"`
}

// OrgConfigAlertSettingsCostAnomaly is an object the spec declares inline.
type OrgConfigAlertSettingsCostAnomaly struct {
	Sigmas            float64 `json:"sigmas"`
	MinDeltaCents     int64   `json:"minDeltaCents"`
	NewSourceMinCents int64   `json:"newSourceMinCents"`
	// SmsAlerts: One of "off", "new_source", "all".
	SmsAlerts string `json:"smsAlerts"`
}

// OrgConfigAlertSettingsDrift is an object the spec declares inline.
type OrgConfigAlertSettingsDrift struct {
	NotifyCreated   bool  `json:"notifyCreated"`
	NotifyUpdated   bool  `json:"notifyUpdated"`
	NotifyDeleted   bool  `json:"notifyDeleted"`
	CooldownMinutes int64 `json:"cooldownMinutes"`
	MinChanges      int64 `json:"minChanges"`
	// Accounts: Account display names; empty means every account.
	Accounts []string `json:"accounts"`
}

// OrgConfigAlertSettingsExpiry is an object the spec declares inline.
type OrgConfigAlertSettingsExpiry struct {
	Enabled  bool  `json:"enabled"`
	LeadDays int64 `json:"leadDays"`
}

// OrgConfigAlertSettingsPosture is an object the spec declares inline.
type OrgConfigAlertSettingsPosture struct {
	Enabled bool `json:"enabled"`
}

// OrgConfigAlertSettingsDigest is an object the spec declares inline.
type OrgConfigAlertSettingsDigest struct {
	Enabled          bool     `json:"enabled"`
	Timezone         string   `json:"timezone"`
	SendDay          int64    `json:"sendDay"`
	SendHour         int64    `json:"sendHour"`
	NarrativeEnabled bool     `json:"narrativeEnabled"`
	Recipients       []string `json:"recipients"`
}

// OrgConfigApplyResultCounts is an object the spec declares inline.
type OrgConfigApplyResultCounts struct {
	Create    int64 `json:"create"`
	Update    int64 `json:"update"`
	Delete    int64 `json:"delete"`
	Unchanged int64 `json:"unchanged"`
}

// OrgConfigBudgetThresholds is an object the spec declares inline.
type OrgConfigBudgetThresholds struct {
	// Type: One of "actual", "forecast".
	Type    string `json:"type"`
	Percent int64  `json:"percent"`
}

// OrgConfigCostCentreRules is an object the spec declares inline.
type OrgConfigCostCentreRules struct {
	Priority int64                          `json:"priority"`
	Match    *OrgConfigCostCentreRulesMatch `json:"match,omitempty"`
}

// OrgConfigDocumentExportedFrom is an object the spec declares inline.
type OrgConfigDocumentExportedFrom struct {
	OrganizationID   string `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
}

// OrgConfigDocumentTagPolicy is an object the spec declares inline.
type OrgConfigDocumentTagPolicy struct {
	RequiredTags    []OrgConfigDocumentTagPolicyRequiredTags `json:"requiredTags"`
	EnforceOnCreate bool                                     `json:"enforceOnCreate"`
}

// OrgConfigPlanCounts is an object the spec declares inline.
type OrgConfigPlanCounts struct {
	Create    int64 `json:"create"`
	Update    int64 `json:"update"`
	Delete    int64 `json:"delete"`
	Unchanged int64 `json:"unchanged"`
}

// OrgConfigWorkflowMetrics is an object the spec declares inline.
type OrgConfigWorkflowMetrics struct {
	Key   string  `json:"key"`
	Label string  `json:"label"`
	Unit  *string `json:"unit,omitempty"`
	Type  *string `json:"type,omitempty"`
}

// PickerResourcesRequestSources is an object the spec declares inline.
type PickerResourcesRequestSources struct {
	PluginID       string `json:"pluginId"`
	ResourceTypeID string `json:"resourceTypeId"`
	OutputKey      string `json:"outputKey"`
}

// PinRangeMetricSeriesPoints is an object the spec declares inline.
type PinRangeMetricSeriesPoints struct {
	Timestamp float64 `json:"timestamp"`
	Value     float64 `json:"value"`
}

// PolicyTemplateHelpLink is an object the spec declares inline.
type PolicyTemplateHelpLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// PreflightCheckHelpLink is an object the spec declares inline.
type PreflightCheckHelpLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// PreflightDeclarationTemplateFormat is an object the spec declares inline.
type PreflightDeclarationTemplateFormat struct {
	Label string `json:"label"`
	// Language: One of "json", "yaml", "text".
	Language string `json:"language"`
}

// ProbeMetricSeriesPoints is an object the spec declares inline.
type ProbeMetricSeriesPoints struct {
	// Timestamp: Unix epoch milliseconds.
	Timestamp float64 `json:"timestamp"`
	Value     float64 `json:"value"`
}

// ProbeRequestItems is an object the spec declares inline.
type ProbeRequestItems struct {
	ResourceID     ResourceID `json:"resourceId"`
	AccountID      string     `json:"accountId"`
	PluginID       string     `json:"pluginId"`
	ResourceTypeID string     `json:"resourceTypeId"`
}

// ProbeStatusSparkline is an object the spec declares inline.
type ProbeStatusSparkline struct {
	// Timestamp: Unix epoch milliseconds.
	Timestamp float64 `json:"timestamp"`
	Value     float64 `json:"value"`
}

// ProbeStatusResourceCounts is an object the spec declares inline.
type ProbeStatusResourceCounts struct {
	TypeLabel string `json:"typeLabel"`
	Count     int64  `json:"count"`
}

// ProfileIdentities is an object the spec declares inline.
type ProfileIdentities struct {
	// Provider: WorkOS OAuth provider id
	Provider string `json:"provider"`
}

// ReorderRequestCards is an object the spec declares inline.
type ReorderRequestCards struct {
	// Kind: One of "resource", "workflow", "widget".
	Kind string `json:"kind"`
	ID   string `json:"id"`
}

// ReportNotificationSendResultSlack is an object the spec declares inline.
type ReportNotificationSendResultSlack struct {
	Attempted int64 `json:"attempted"`
	Succeeded int64 `json:"succeeded"`
}

// ReportNotificationSendResultTeams is an object the spec declares inline.
type ReportNotificationSendResultTeams struct {
	Attempted int64 `json:"attempted"`
	Succeeded int64 `json:"succeeded"`
}

// ReportNotificationSendResultEmail is an object the spec declares inline.
type ReportNotificationSendResultEmail struct {
	Attempted int64 `json:"attempted"`
	Succeeded int64 `json:"succeeded"`
}

// ResourceTypeSummaryAttachTargets is an object the spec declares inline.
type ResourceTypeSummaryAttachTargets struct {
	PluginID       string  `json:"pluginId"`
	ResourceTypeID string  `json:"resourceTypeId"`
	MatchField     *string `json:"matchField,omitempty"`
	Verb           *string `json:"verb,omitempty"`
}

// SecretExportTemplateEntries is an object the spec declares inline.
type SecretExportTemplateEntries struct {
	OutputKey string `json:"outputKey"`
	EnvKey    string `json:"envKey"`
}

// SessionRecordingParticipants is an object the spec declares inline.
type SessionRecordingParticipants struct {
	UserID   *string `json:"userId"`
	UserName *string `json:"userName"`
	// Role: One of "observer", "driver".
	Role     string  `json:"role"`
	JoinedAt string  `json:"joinedAt"`
	LeftAt   *string `json:"leftAt"`
}

// ShowbackReportCentres is an object the spec declares inline.
type ShowbackReportCentres struct {
	// CostCentreID: Null for the synthetic "Unallocated" bucket.
	CostCentreID *string `json:"costCentreId"`
	Name         string  `json:"name"`
	// Totals: Spend allocated directly to this centre. A cost row is allocated
	// exactly once, so summing this across every entry equals the organization's
	// spend for the period.
	Totals map[string]float64 `json:"totals"`
	// SubtreeTotals: This centre's own spend plus every descendant's. Equal to
	// `totals` for a leaf and for every centre in an organization that does not
	// nest. Do not sum this across entries — parents already contain their
	// children.
	SubtreeTotals map[string]float64 `json:"subtreeTotals"`
	// ParentID: The centre this one sits under; null for a root and for
	// Unallocated.
	ParentID *string `json:"parentId"`
	// Depth: 0 for a root; the indentation level.
	Depth int64 `json:"depth"`
}

// SignSshkeyRequestContext is an object the spec declares inline.
type SignSshkeyRequestContext struct {
	Host     *string `json:"host,omitempty"`
	Username *string `json:"username,omitempty"`
}

// SshfanoutHostResultHostKeyTrust is an object the spec declares inline.
type SshfanoutHostResultHostKeyTrust struct {
	// Kind: One of "unknown", "mismatch".
	Kind                 string  `json:"kind"`
	Host                 string  `json:"host"`
	Port                 int64   `json:"port"`
	PresentedFingerprint string  `json:"presentedFingerprint"`
	StoredFingerprint    *string `json:"storedFingerprint"`
}

// SshfanoutRunRequestTargets is an object the spec declares inline.
type SshfanoutRunRequestTargets struct {
	// Kind: One of "account", "resource".
	Kind string `json:"kind"`
	ID   string `json:"id"`
}

// TerraformExportExported is an object the spec declares inline.
type TerraformExportExported struct {
	ID             ResourceID `json:"id"`
	DisplayName    string     `json:"displayName"`
	PluginID       string     `json:"pluginId"`
	ResourceTypeID string     `json:"resourceTypeId"`
	Address        string     `json:"address"`
	ImportID       *string    `json:"importId,omitempty"`
}

// TerraformExportUnsupported is an object the spec declares inline.
type TerraformExportUnsupported struct {
	ID             ResourceID `json:"id"`
	DisplayName    string     `json:"displayName"`
	PluginID       string     `json:"pluginId"`
	ResourceTypeID string     `json:"resourceTypeId"`
	Reason         string     `json:"reason"`
}

// UnitCostQueryResponseMetric is an object the spec declares inline.
type UnitCostQueryResponseMetric struct {
	ID       string             `json:"id"`
	Key      string             `json:"key"`
	Name     string             `json:"name"`
	Unit     string             `json:"unit"`
	Kind     BusinessMetricKind `json:"kind"`
	Currency *string            `json:"currency"`
}

// UnitCostQueryResponseConversion is an object the spec declares inline.
type UnitCostQueryResponseConversion struct {
	DisplayCurrency string                                     `json:"displayCurrency"`
	Converted       []UnitCostQueryResponseConversionConverted `json:"converted"`
	Unconverted     []string                                   `json:"unconverted"`
}

// UntaggedSpendReportByKey is an object the spec declares inline.
type UntaggedSpendReportByKey struct {
	Key string `json:"key"`
	// Untagged: Currency code → amount in the currency's major unit.
	Untagged map[string]float64 `json:"untagged"`
}

// UntaggedSpendReportTopUntagged is an object the spec declares inline.
type UntaggedSpendReportTopUntagged struct {
	AccountID    string  `json:"accountId"`
	AccountLabel string  `json:"accountLabel"`
	Service      string  `json:"service"`
	Currency     string  `json:"currency"`
	Amount       float64 `json:"amount"`
}

// ValidateTabsRequestTabs is an object the spec declares inline.
type ValidateTabsRequestTabs struct {
	ID     string    `json:"id"`
	Target TabTarget `json:"target"`
}

// CostAccountStatusCostPollErrorHelpLink is an object the spec declares inline.
type CostAccountStatusCostPollErrorHelpLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// DeployPlanResultResultError is an object the spec declares inline.
type DeployPlanResultResultError struct {
	Message string  `json:"message"`
	Stack   *string `json:"stack,omitempty"`
}

// EnvironmentCaptureDraftMemberFieldMetaValueOptions is an object the spec
// declares inline.
type EnvironmentCaptureDraftMemberFieldMetaValueOptions struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// InvoiceDerivationScopeCostCentres is an object the spec declares inline.
type InvoiceDerivationScopeCostCentres struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// InvoiceDerivationScopeAccounts is an object the spec declares inline.
type InvoiceDerivationScopeAccounts struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// OrgConfigCostCentreRulesMatch is an object the spec declares inline.
type OrgConfigCostCentreRulesMatch struct {
	TagKey   *string `json:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty"`
	// Account: Account display name.
	Account  *string `json:"account,omitempty"`
	PluginID *string `json:"pluginId,omitempty"`
	Service  *string `json:"service,omitempty"`
}

// OrgConfigDocumentTagPolicyRequiredTags is an object the spec declares inline.
type OrgConfigDocumentTagPolicyRequiredTags struct {
	Key           string   `json:"key"`
	AllowedValues []string `json:"allowedValues,omitempty"`
}

// UnitCostQueryResponseConversionConverted is an object the spec declares
// inline.
type UnitCostQueryResponseConversionConverted struct {
	Currency string                                          `json:"currency"`
	Rates    []UnitCostQueryResponseConversionConvertedRates `json:"rates"`
}

// UnitCostQueryResponseConversionConvertedRates is an object the spec declares
// inline.
type UnitCostQueryResponseConversionConvertedRates struct {
	EffectiveFrom string  `json:"effectiveFrom"`
	Rate          float64 `json:"rate"`
}

// AccountsCreateRequest is an object the spec declares inline.
type AccountsCreateRequest struct {
	PluginID    *PluginID         `json:"pluginId,omitempty"`
	DisplayName string            `json:"displayName"`
	Credentials map[string]string `json:"credentials"`
	// BastionID: Optional bastion id to route this account's cloud API traffic
	// through.
	BastionID *string `json:"bastionId,omitempty"`
}

// AccountsCredentialsUpdateRequest is an object the spec declares inline.
type AccountsCredentialsUpdateRequest struct {
	// Credentials: Complete credentials map. Sensitive fields the caller doesn't
	// want to change should be re-sent with their previous value (the server
	// doesn't merge with the existing blob).
	Credentials map[string]string `json:"credentials"`
}

// AccountsCredentialsUpdateResponse is an object the spec declares inline.
type AccountsCredentialsUpdateResponse struct {
	OK bool `json:"ok"`
}

// AgentsSessionsDeleteResponse is an object the spec declares inline.
type AgentsSessionsDeleteResponse struct {
	OK bool `json:"ok"`
}

// AgentsSessionsOpenResponse is an object the spec declares inline.
type AgentsSessionsOpenResponse struct {
	Command    string  `json:"command"`
	Cwd        string  `json:"cwd"`
	SSHKeyID   *string `json:"sshKeyId,omitempty"`
	SSHKeyName *string `json:"sshKeyName,omitempty"`
}

// AgentsSessionsReconcileResponse is an object the spec declares inline.
type AgentsSessionsReconcileResponse struct {
	BranchName string `json:"branchName"`
	Message    string `json:"message"`
}

// AlertRulesAdoptDefaultsResponse is an object the spec declares inline.
type AlertRulesAdoptDefaultsResponse struct {
	Rules   []AlertRule `json:"rules"`
	Adopted bool        `json:"adopted"`
}

// AlertRulesUpdateRequest is an object the spec declares inline.
type AlertRulesUpdateRequest struct {
	Rules []AlertRuleInput `json:"rules"`
}

// AlertRulesUpdateResponse is an object the spec declares inline.
type AlertRulesUpdateResponse struct {
	Rules []AlertRule `json:"rules"`
}

// AlertRulesDeliveriesAckResponse is an object the spec declares inline.
type AlertRulesDeliveriesAckResponse struct {
	Acknowledged          bool    `json:"acknowledged"`
	AlreadyAcknowledgedBy *string `json:"alreadyAcknowledgedBy,omitempty"`
	// Reason: Why the acknowledgement did not take. `not_pending` means the
	// delivery exists but was never awaiting one — still held, already sent, or
	// expired.
	//
	// One of "not_pending", "already_escalated", "already_acknowledged".
	Reason *string `json:"reason,omitempty"`
	Title  *string `json:"title,omitempty"`
}

// AlertRulesDeliveriesCancelRequest is an object the spec declares inline.
type AlertRulesDeliveriesCancelRequest struct {
	IDs []string `json:"ids"`
}

// AlertRulesDeliveriesCancelResponse is an object the spec declares inline.
type AlertRulesDeliveriesCancelResponse struct {
	Cancelled int64 `json:"cancelled"`
}

// BackupsDrillsLogResponse is an object the spec declares inline.
type BackupsDrillsLogResponse struct {
	Drills []RestoreDrill `json:"drills"`
}

// BusinessMetricsGetGetResponse is an object the spec declares inline.
type BusinessMetricsGetGetResponse struct {
	Metrics []BusinessMetric `json:"metrics"`
}

// BusinessMetricsValuesCreateResponse is an object the spec declares inline.
type BusinessMetricsValuesCreateResponse struct {
	// Written: Days written, counting restatements.
	Written int64 `json:"written"`
}

// BusinessMetricsValuesGetResponse is an object the spec declares inline.
type BusinessMetricsValuesGetResponse struct {
	Values []BusinessMetricValue `json:"values"`
}

// CostAlertsEventsResponse is an object the spec declares inline.
type CostAlertsEventsResponse struct {
	Events []CostAlertEvent `json:"events"`
}

// CostAlertsGetGetResponse is an object the spec declares inline.
type CostAlertsGetGetResponse struct {
	Alerts []CostAlert `json:"alerts"`
}

// CostAnnotationsGetResponse is an object the spec declares inline.
type CostAnnotationsGetResponse struct {
	Annotations []CostAnnotation `json:"annotations"`
}

// CostScenariosReferentsResponse is an object the spec declares inline.
type CostScenariosReferentsResponse struct {
	Referents []CostScenarioReferent `json:"referents"`
}

// CostScenariosGetGetResponse is an object the spec declares inline.
type CostScenariosGetGetResponse struct {
	Models []CostScenarioModel `json:"models"`
}

// CostsEfficiencyAlertsResponse is an object the spec declares inline.
type CostsEfficiencyAlertsResponse struct {
	Events []EfficiencyAlertEvent `json:"events"`
}

// CostsStatusResponse is an object the spec declares inline.
type CostsStatusResponse struct {
	Accounts []CostAccountStatus `json:"accounts"`
}

// CostsAnomaliesAcknowledgeRequest is an object the spec declares inline.
type CostsAnomaliesAcknowledgeRequest struct {
	// Explanation: One sentence on what caused the spend. Becomes the
	// annotation's text, so the annotation's 500-character ceiling applies.
	Explanation string `json:"explanation"`
}

// CostsAnomaliesGetResponse is an object the spec declares inline.
type CostsAnomaliesGetResponse struct {
	Anomalies []CostAnomaly `json:"anomalies"`
}

// CurrencyRatesDeleteResponse is an object the spec declares inline.
type CurrencyRatesDeleteResponse struct {
	OK bool `json:"ok"`
}

// DashboardsCreateRequest is an object the spec declares inline.
type DashboardsCreateRequest struct {
	Name string `json:"name"`
}

// DashboardsRenameRequest is an object the spec declares inline.
type DashboardsRenameRequest struct {
	Name string `json:"name"`
}

// DeploymentsRunsCreateResponse is an object the spec declares inline.
type DeploymentsRunsCreateResponse struct {
	ID string `json:"id"`
}

// DeploymentsTriggersUpdateRequest is an object the spec declares inline.
type DeploymentsTriggersUpdateRequest struct {
	Enabled bool `json:"enabled"`
}

// DigestRecipientsDeleteResponse is an object the spec declares inline.
type DigestRecipientsDeleteResponse struct {
	OK bool `json:"ok"`
}

// IacStatesCreateResponse is an object the spec declares inline.
type IacStatesCreateResponse struct {
	State IacState `json:"state"`
}

// JiraGetResponse is an object the spec declares inline.
type JiraGetResponse struct {
	Integration *JiraIntegration `json:"integration"`
}

// LinearGetResponse is an object the spec declares inline.
type LinearGetResponse struct {
	Integration *LinearIntegration `json:"integration"`
}

// MsteamsTestResponse is an object the spec declares inline.
type MsteamsTestResponse struct {
	OK           bool  `json:"ok"`
	WebhookCount int64 `json:"webhookCount"`
	Attempted    int64 `json:"attempted"`
	Succeeded    int64 `json:"succeeded"`
}

// OnCallOverridesGetResponse is an object the spec declares inline.
type OnCallOverridesGetResponse struct {
	Overrides []OnCallOverride `json:"overrides"`
}

// ProfilePasswordResetResponse is an object the spec declares inline.
type ProfilePasswordResetResponse struct {
	PasswordResetURL string `json:"passwordResetUrl"`
	ExpiresAt        string `json:"expiresAt"`
}

// ProfileUpdateRequest is an object the spec declares inline.
type ProfileUpdateRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// ProfileEmailChangeConfirmRequest is an object the spec declares inline.
type ProfileEmailChangeConfirmRequest struct {
	Code string `json:"code"`
}

// ProfileEmailChangeConfirmResponse is an object the spec declares inline.
type ProfileEmailChangeConfirmResponse struct {
	Email string `json:"email"`
}

// ProfileEmailChangeCreateRequest is an object the spec declares inline.
type ProfileEmailChangeCreateRequest struct {
	NewEmail string `json:"newEmail"`
}

// ProfileEmailChangeCreateResponse is an object the spec declares inline.
type ProfileEmailChangeCreateResponse struct {
	NewEmail  string `json:"newEmail"`
	ExpiresAt string `json:"expiresAt"`
}

// ProfileMfachallengeResponse is an object the spec declares inline.
type ProfileMfachallengeResponse struct {
	ChallengeID string `json:"challengeId"`
}

// ProfileMfaverifyRequest is an object the spec declares inline.
type ProfileMfaverifyRequest struct {
	ChallengeID string `json:"challengeId"`
	Code        string `json:"code"`
}

// ProfileMfaverifyResponse is an object the spec declares inline.
type ProfileMfaverifyResponse struct {
	Verified bool `json:"verified"`
}

// ProfileSessionsRevokeOthersResponse is an object the spec declares inline.
type ProfileSessionsRevokeOthersResponse struct {
	Revoked int64 `json:"revoked"`
}

// ResourcesCostEstimateResponse is an object the spec declares inline.
type ResourcesCostEstimateResponse struct {
	Estimate *CostEstimate `json:"estimate"`
}

// ResourcesNoSqlcommandResponse is an object the spec declares inline.
type ResourcesNoSqlcommandResponse struct {
	Result JSONObject `json:"result"`
}

// SavedCostFiltersReferentsResponse is an object the spec declares inline.
type SavedCostFiltersReferentsResponse struct {
	Referents []SavedCostFilterReferent `json:"referents"`
}

// SharedConsolesHandoverRequest is an object the spec declares inline.
type SharedConsolesHandoverRequest struct {
	ParticipantID string `json:"participantId"`
}

// SharedConsolesJoinRequest is an object the spec declares inline.
type SharedConsolesJoinRequest struct {
	Token string `json:"token"`
}

// SharedConsolesInvitesCreateRequest is an object the spec declares inline.
type SharedConsolesInvitesCreateRequest struct {
	InviteTTLMinutes *int64 `json:"inviteTtlMinutes,omitempty"`
}

// SlackInstallUrlresponse is an object the spec declares inline.
type SlackInstallUrlresponse struct {
	URL string `json:"url"`
}

// SlackTestResponse is an object the spec declares inline.
type SlackTestResponse struct {
	OK           bool  `json:"ok"`
	ChannelCount int64 `json:"channelCount"`
	Attempted    int64 `json:"attempted"`
	Succeeded    int64 `json:"succeeded"`
}

// SlackInstallationsAvailableChannelsResponse is an object the spec declares
// inline.
type SlackInstallationsAvailableChannelsResponse struct {
	Channels []SlackAvailableChannel `json:"channels"`
}

// SshfanoutSnippetsCreateResponse is an object the spec declares inline.
type SshfanoutSnippetsCreateResponse struct {
	ID string `json:"id"`
}

// SshfanoutSnippetsGetResponse is an object the spec declares inline.
type SshfanoutSnippetsGetResponse struct {
	Snippets []SSHSnippet `json:"snippets"`
}

// SshtunnelsCloseRequest is an object the spec declares inline.
type SshtunnelsCloseRequest struct {
	TunnelID string `json:"tunnelId"`
}

// SshtunnelsOpenRequest is an object the spec declares inline.
type SshtunnelsOpenRequest struct {
	AccountID string `json:"accountId"`
}

// SshtunnelsOpenResponse is an object the spec declares inline.
type SshtunnelsOpenResponse struct {
	TunnelID  string `json:"tunnelId"`
	LocalPort int64  `json:"localPort"`
}
