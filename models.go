// github.com/Infrawrench/infrawrench-go v0.30.0 | MIT | Copyright (c) 2026 Infrawrench LLC
// https://github.com/Infrawrench/Infrawrench
//
// Generated from the Infrawrench API OpenAPI 3.1 spec (API version 0.30.0).
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
	Tool       string `json:"tool"`
	BranchName string `json:"branchName"`
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
	Tool   string            `json:"tool"`
	Fields map[string]string `json:"fields"`
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
	ID         string     `json:"id"`
	UserID     *string    `json:"userId"`
	APIKeyID   *string    `json:"apiKeyId"`
	Action     string     `json:"action"`
	EntityType string     `json:"entityType"`
	EntityID   string     `json:"entityId"`
	Metadata   JSONObject `json:"metadata"`
	IPAddress  *string    `json:"ipAddress"`
	CreatedAt  string     `json:"createdAt"`
	UserName   *string    `json:"userName"`
	UserEmail  *string    `json:"userEmail"`
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

// BillingStatus is the `BillingStatus` schema.
type BillingStatus struct {
	// Complimentary: Platform-granted complimentary access: all paid perks,
	// uncapped AI chat, never billed.
	Complimentary bool          `json:"complimentary"`
	Subscription  *Subscription `json:"subscription"`
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

// BudgetCostFilter is the `BudgetCostFilter` schema.
type BudgetCostFilter struct {
	// Dimension: One of "provider", "account", "service", "region", "resource",
	// "tag".
	Dimension string `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
}

// BudgetFull is the `BudgetFull` schema.
type BudgetFull struct {
	ID              string             `json:"id"`
	OrganizationID  string             `json:"organizationId"`
	Name            string             `json:"name"`
	AmountCents     int64              `json:"amountCents"`
	Currency        string             `json:"currency"`
	Filters         []BudgetCostFilter `json:"filters"`
	Thresholds      []BudgetThreshold  `json:"thresholds"`
	CreatedByUserID *string            `json:"createdByUserId"`
	DeletedAt       *string            `json:"deletedAt"`
	CreatedAt       string             `json:"createdAt"`
	UpdatedAt       string             `json:"updatedAt"`
}

// BudgetInput is the `BudgetInput` schema.
type BudgetInput struct {
	Name        string             `json:"name"`
	AmountCents int64              `json:"amountCents"`
	Currency    *string            `json:"currency,omitempty"`
	Filters     []BudgetCostFilter `json:"filters,omitempty"`
	Thresholds  []BudgetThreshold  `json:"thresholds"`
}

// BudgetThreshold is the `BudgetThreshold` schema.
type BudgetThreshold struct {
	// Type: One of "actual", "forecast".
	Type    string `json:"type"`
	Percent int64  `json:"percent"`
}

// BudgetWithStatus is the `BudgetWithStatus` schema.
type BudgetWithStatus struct {
	ID                 string                               `json:"id"`
	Name               string                               `json:"name"`
	AmountCents        int64                                `json:"amountCents"`
	Currency           string                               `json:"currency"`
	Filters            []BudgetCostFilter                   `json:"filters"`
	Thresholds         []BudgetThreshold                    `json:"thresholds"`
	Month              string                               `json:"month"`
	ActualCents        int64                                `json:"actualCents"`
	ForecastCents      *int64                               `json:"forecastCents"`
	CurrentMonthEvents []BudgetWithStatusCurrentMonthEvents `json:"currentMonthEvents"`
	Placements         []BudgetWithStatusPlacements         `json:"placements"`
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
	AccountID            string   `json:"accountId"`
	PluginID             string   `json:"pluginId"`
	DisplayName          string   `json:"displayName"`
	SupportsCosts        bool     `json:"supportsCosts"`
	PeriodNative         bool     `json:"periodNative"`
	Dimensions           []string `json:"dimensions"`
	CostLastPolledAt     *string  `json:"costLastPolledAt"`
	CostBackfilledAt     *string  `json:"costBackfilledAt"`
	CostPollFailureCount int64    `json:"costPollFailureCount"`
	// CostPollError: Last cost-collection failure for this account, cleared on
	// the next success. `helpLink` points at the provider page that fixes a
	// setup problem when the plugin can identify one (e.g. GCP's billing export
	// console).
	CostPollError *CostAccountStatusCostPollError `json:"costPollError"`
	Coverage      *CostAccountStatusCoverage      `json:"coverage"`
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

// CostCentre is the `CostCentre` schema.
type CostCentre struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// CostCentreInput is the `CostCentreInput` schema.
type CostCentreInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// CostDimension is the `CostDimension` schema.
type CostDimension = string

// The values CostDimension takes.
const (
	CostDimensionProvider CostDimension = "provider"
	CostDimensionAccount  CostDimension = "account"
	CostDimensionService  CostDimension = "service"
	CostDimensionRegion   CostDimension = "region"
	CostDimensionResource CostDimension = "resource"
	CostDimensionTag      CostDimension = "tag"
)

// CostDimensionValues is the `CostDimensionValues` schema.
type CostDimensionValues struct {
	Values []any `json:"values"`
}

// CostFilter is the `CostFilter` schema.
type CostFilter struct {
	Dimension CostDimension `json:"dimension"`
	// Op: One of "in", "not_in".
	Op     string   `json:"op"`
	Values []string `json:"values"`
	TagKey *string  `json:"tagKey,omitempty"`
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
	// "resource", "tag".
	GroupBy               string       `json:"groupBy"`
	GroupByTagKey         *string      `json:"groupByTagKey,omitempty"`
	Filters               []CostFilter `json:"filters,omitempty"`
	TopN                  *int64       `json:"topN,omitempty"`
	ComparePreviousPeriod *bool        `json:"comparePreviousPeriod,omitempty"`
	Forecast              *bool        `json:"forecast,omitempty"`
}

// CostQueryResponse is the `CostQueryResponse` schema.
type CostQueryResponse struct {
	Series         []CostQuerySeries  `json:"series"`
	Comparison     []CostQuerySeries  `json:"comparison,omitempty"`
	Forecast       []CostSeriesPoint  `json:"forecast,omitempty"`
	Currencies     []string           `json:"currencies"`
	Totals         map[string]float64 `json:"totals"`
	PreviousTotals map[string]float64 `json:"previousTotals,omitempty"`
}

// CostQuerySeries is the `CostQuerySeries` schema.
type CostQuerySeries struct {
	Key      string            `json:"key"`
	Label    string            `json:"label"`
	Currency string            `json:"currency"`
	Points   []CostSeriesPoint `json:"points"`
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
	Repo          string         `json:"repo"`
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

// CreateCostEstimateRequest is the `CreateCostEstimateRequest` schema.
type CreateCostEstimateRequest struct {
	AccountID        string            `json:"accountId"`
	ResourceTypeID   string            `json:"resourceTypeId"`
	Fields           map[string]string `json:"fields"`
	PluginID         *string           `json:"pluginId,omitempty"`
	ParentResourceID *ResourceID       `json:"parentResourceId,omitempty"`
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

// DashboardWidgetKind is the `DashboardWidgetKind` schema.
type DashboardWidgetKind = string

// The values DashboardWidgetKind takes.
const (
	DashboardWidgetKindCostGraph   DashboardWidgetKind = "cost_graph"
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

// Error is the `Error` schema.
type Error struct {
	// Error: Human-readable error message
	Error string `json:"error"`
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
	// this set to buy one more seat and send the invitation. Requires
	// billing:write.
	AddSeat *bool `json:"addSeat,omitempty"`
}

// InviteResponse is the `InviteResponse` schema.
type InviteResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
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
}

// LogCapableResourceList is the `LogCapableResourceList` schema.
type LogCapableResourceList struct {
	Resources []LogCapableResource `json:"resources"`
}

// LogStreamSelector is the `LogStreamSelector` schema.
type LogStreamSelector struct {
	// ResourceID: Infrawrench resource id of the stream to tail.
	ResourceID     string   `json:"resourceId"`
	AccountID      string   `json:"accountId"`
	PluginID       PluginID `json:"pluginId"`
	ResourceTypeID string   `json:"resourceTypeId"`
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
	URLHint       string `json:"urlHint"`
	SyncIncidents bool   `json:"syncIncidents"`
	BudgetAlerts  bool   `json:"budgetAlerts"`
	// AnomalyAlerts: Statistical spend-spike (cost anomaly) alerts
	AnomalyAlerts bool `json:"anomalyAlerts"`
	// MetricAlerts: Metric threshold rule firings and recoveries
	MetricAlerts bool `json:"metricAlerts"`
	// ResourceDrift: Batched resource-drift digests from the change timeline.
	// Defaults to false when a channel is added — drift is continuous where the
	// other triggers are exceptional.
	ResourceDrift bool `json:"resourceDrift"`
	// WorkflowPages: Pages and approval requests raised by a workflow
	// (infra.page / infra.waitForApproval) or by POST /pages
	WorkflowPages bool `json:"workflowPages"`
	// ProviderIncidents: A provider status-page incident overlaps resources you
	// hold.
	ProviderIncidents bool `json:"providerIncidents"`
	// ExpiryAlerts: Daily digests of approaching resource deadlines — expiring
	// certificates, domains, tokens and keys past their rotation budget.
	ExpiryAlerts bool `json:"expiryAlerts"`
	// LogMatchAlerts: A saved log-workspace query with alerting enabled found
	// matching log lines.
	LogMatchAlerts bool `json:"logMatchAlerts"`
	// WeeklyDigest: The Monday-morning weekly digest. Only sends when the
	// organization has enabled the digest (see /digest).
	WeeklyDigest bool `json:"weeklyDigest"`
}

// MsTeamsWebhookCreate is the `MsTeamsWebhookCreate` schema.
type MsTeamsWebhookCreate struct {
	Label string `json:"label"`
	// URL: The webhook URL from a Teams 'Workflows' automation. Must be https
	// and on a Microsoft-operated host (*.api.powerautomate.com,
	// *.api.powerplatform.com, *.logic.azure.com, *.flow.microsoft.com, or a
	// legacy *.webhook.office.com connector).
	URL               string `json:"url"`
	SyncIncidents     *bool  `json:"syncIncidents,omitempty"`
	BudgetAlerts      *bool  `json:"budgetAlerts,omitempty"`
	AnomalyAlerts     *bool  `json:"anomalyAlerts,omitempty"`
	MetricAlerts      *bool  `json:"metricAlerts,omitempty"`
	ResourceDrift     *bool  `json:"resourceDrift,omitempty"`
	WorkflowPages     *bool  `json:"workflowPages,omitempty"`
	ProviderIncidents *bool  `json:"providerIncidents,omitempty"`
	ExpiryAlerts      *bool  `json:"expiryAlerts,omitempty"`
	LogMatchAlerts    *bool  `json:"logMatchAlerts,omitempty"`
	WeeklyDigest      *bool  `json:"weeklyDigest,omitempty"`
}

// MsTeamsWebhookUpdate is the `MsTeamsWebhookUpdate` schema.
type MsTeamsWebhookUpdate struct {
	Label             *string `json:"label,omitempty"`
	SyncIncidents     *bool   `json:"syncIncidents,omitempty"`
	BudgetAlerts      *bool   `json:"budgetAlerts,omitempty"`
	AnomalyAlerts     *bool   `json:"anomalyAlerts,omitempty"`
	MetricAlerts      *bool   `json:"metricAlerts,omitempty"`
	ResourceDrift     *bool   `json:"resourceDrift,omitempty"`
	WorkflowPages     *bool   `json:"workflowPages,omitempty"`
	ProviderIncidents *bool   `json:"providerIncidents,omitempty"`
	ExpiryAlerts      *bool   `json:"expiryAlerts,omitempty"`
	LogMatchAlerts    *bool   `json:"logMatchAlerts,omitempty"`
	WeeklyDigest      *bool   `json:"weeklyDigest,omitempty"`
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
	Reason       string                `json:"reason"`
	Cost         *OrphanCostAnnotation `json:"cost"`
	LastSyncedAt *string               `json:"lastSyncedAt"`
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
	PermissionAccountsRead      Permission = "accounts:read"
	PermissionAccountsWrite     Permission = "accounts:write"
	PermissionAccountsDelete    Permission = "accounts:delete"
	PermissionResourcesRead     Permission = "resources:read"
	PermissionResourcesWrite    Permission = "resources:write"
	PermissionResourcesDelete   Permission = "resources:delete"
	PermissionResourcesExecute  Permission = "resources:execute"
	PermissionSecretsRead       Permission = "secrets:read"
	PermissionSecretsWrite      Permission = "secrets:write"
	PermissionStorageRead       Permission = "storage:read"
	PermissionStorageWrite      Permission = "storage:write"
	PermissionDashboardsRead    Permission = "dashboards:read"
	PermissionDashboardsWrite   Permission = "dashboards:write"
	PermissionWorkflowsRead     Permission = "workflows:read"
	PermissionWorkflowsWrite    Permission = "workflows:write"
	PermissionWorkflowsApprove  Permission = "workflows:approve"
	PermissionDeploymentsRead   Permission = "deployments:read"
	PermissionDeploymentsPlan   Permission = "deployments:plan"
	PermissionDeploymentsWrite  Permission = "deployments:write"
	PermissionCostsRead         Permission = "costs:read"
	PermissionCostsWrite        Permission = "costs:write"
	PermissionBudgetsRead       Permission = "budgets:read"
	PermissionBudgetsWrite      Permission = "budgets:write"
	PermissionMetricAlertsRead  Permission = "metric-alerts:read"
	PermissionMetricAlertsWrite Permission = "metric-alerts:write"
	PermissionFreezesRead       Permission = "freezes:read"
	PermissionFreezesWrite      Permission = "freezes:write"
	PermissionFreezesOverride   Permission = "freezes:override"
	PermissionTagPolicyOverride Permission = "tag-policy:override"
	PermissionAuditRead         Permission = "audit:read"
	PermissionTeamRead          Permission = "team:read"
	PermissionTeamInvite        Permission = "team:invite"
	PermissionTeamRoleWrite     Permission = "team:role:write"
	PermissionTeamRemove        Permission = "team:remove"
	PermissionApikeysRead       Permission = "apikeys:read"
	PermissionApikeysWrite      Permission = "apikeys:write"
	PermissionBillingRead       Permission = "billing:read"
	PermissionBillingWrite      Permission = "billing:write"
	PermissionSSHKeysRead       Permission = "ssh-keys:read"
	PermissionSSHKeysWrite      Permission = "ssh-keys:write"
	PermissionBastionsRead      Permission = "bastions:read"
	PermissionBastionsWrite     Permission = "bastions:write"
	PermissionChatRead          Permission = "chat:read"
	PermissionChatWrite         Permission = "chat:write"
	PermissionPagesWrite        Permission = "pages:write"
	PermissionOrgSettingsWrite  Permission = "org:settings:write"
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
	PluginIDVercel       PluginID = "vercel"
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

// ReauthenticationRequired is the `ReauthenticationRequired` schema.
type ReauthenticationRequired struct {
	// Error: Human-readable error message
	Error string `json:"error"`
	// Code: One of "reauthentication_required".
	Code string `json:"code"`
}

// ReorderRequest is the `ReorderRequest` schema.
type ReorderRequest struct {
	Cards       []ReorderRequestCards `json:"cards,omitempty"`
	ResourceIDs []ResourceID          `json:"resourceIds,omitempty"`
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
	Origin      *string `json:"origin,omitempty"`
	CreatedAt   string  `json:"createdAt"`
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
	DetailSchema         JSONObject         `json:"detailSchema"`
	ChildResources       []ChildResourceRef `json:"childResources"`
	ChildTypes           []ChildTypeRef     `json:"childTypes"`
	PluginID             string             `json:"pluginId"`
	PluginLogoSvg        string             `json:"pluginLogoSvg"`
	ResourceID           ResourceID         `json:"resourceId"`
	AccountID            string             `json:"accountId"`
	ResourceTypeID       string             `json:"resourceTypeId"`
	PeerPanes            []PeerPane         `json:"peerPanes"`
	PeerIntegrationStubs []PeerPaneStub     `json:"peerIntegrationStubs"`
	CanDelete            bool               `json:"canDelete"`
	CanEdit              bool               `json:"canEdit"`
	EditableFields       []EditableField    `json:"editableFields"`
	CredentialFormats    []CredentialFormat `json:"credentialFormats"`
	HasManifestEditor    bool               `json:"hasManifestEditor"`
	HasSecretVersions    bool               `json:"hasSecretVersions"`
	ResourceDisplayName  string             `json:"resourceDisplayName"`
	ResourceTypeLabel    string             `json:"resourceTypeLabel"`
	ResourceFields       JSONObject         `json:"resourceFields"`
	HasSQLEditor         bool               `json:"hasSqlEditor"`
	HasStorageBrowser    bool               `json:"hasStorageBrowser"`
	HasArtifactRegistry  bool               `json:"hasArtifactRegistry"`
	HasKVBrowser         bool               `json:"hasKvBrowser"`
	HasKVConsole         bool               `json:"hasKvConsole"`
	KVDriverName         *string            `json:"kvDriverName,omitempty"`
	IsMongoDB            bool               `json:"isMongoDb"`
	HasDockerActions     bool               `json:"hasDockerActions"`
	HasSSHTerminal       bool               `json:"hasSshTerminal"`
	HasSFTPBrowser       bool               `json:"hasSftpBrowser"`
	SSHHost              *string            `json:"sshHost"`
	SSHPrivateHost       *string            `json:"sshPrivateHost,omitempty"`
	DefaultSSHUsername   *string            `json:"defaultSshUsername"`
	ContainerID          string             `json:"containerId"`
	DatabaseName         string             `json:"databaseName"`
	StorageBucketName    string             `json:"storageBucketName"`
	SupportsMetrics      bool               `json:"supportsMetrics"`
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
	// Schedulable: The type declares lifecycle start/stop actions, so its
	// resources can carry a sleep/wake schedule.
	Schedulable *bool `json:"schedulable,omitempty"`
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
	// SeatCount: Seats on the plan
	SeatCount int64 `json:"seatCount"`
	// SeatsUsed: Members plus pending unexpired invitations
	SeatsUsed int64 `json:"seatsUsed"`
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

// ShowbackReport is the `ShowbackReport` schema.
type ShowbackReport struct {
	From       string                  `json:"from"`
	To         string                  `json:"to"`
	Currencies []string                `json:"currencies"`
	Centres    []ShowbackReportCentres `json:"centres"`
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
	ChannelName   string `json:"channelName"`
	IsPrivate     bool   `json:"isPrivate"`
	SyncIncidents bool   `json:"syncIncidents"`
	BudgetAlerts  bool   `json:"budgetAlerts"`
	// AnomalyAlerts: Statistical spend-spike (cost anomaly) alerts
	AnomalyAlerts bool `json:"anomalyAlerts"`
	// MetricAlerts: Metric threshold rule firings and recoveries
	MetricAlerts bool `json:"metricAlerts"`
	// ResourceDrift: Batched resource-drift digests from the change timeline.
	// Defaults to false when a channel is added — drift is continuous where the
	// other triggers are exceptional.
	ResourceDrift bool `json:"resourceDrift"`
	// WorkflowPages: Pages and approval requests raised by a workflow
	// (infra.page / infra.waitForApproval) or by POST /pages
	WorkflowPages bool `json:"workflowPages"`
	// ProviderIncidents: A provider status-page incident overlaps resources you
	// hold.
	ProviderIncidents bool `json:"providerIncidents"`
	// ExpiryAlerts: Daily digests of approaching resource deadlines — expiring
	// certificates, domains, tokens and keys past their rotation budget.
	ExpiryAlerts bool `json:"expiryAlerts"`
	// LogMatchAlerts: A saved log-workspace query with alerting enabled found
	// matching log lines.
	LogMatchAlerts bool `json:"logMatchAlerts"`
	// WeeklyDigest: The Monday-morning weekly digest. Only sends when the
	// organization has enabled the digest (see /digest).
	WeeklyDigest bool `json:"weeklyDigest"`
}

// SlackChannelCreate is the `SlackChannelCreate` schema.
type SlackChannelCreate struct {
	InstallationID    string `json:"installationId"`
	ChannelID         string `json:"channelId"`
	ChannelName       string `json:"channelName"`
	IsPrivate         *bool  `json:"isPrivate,omitempty"`
	SyncIncidents     *bool  `json:"syncIncidents,omitempty"`
	BudgetAlerts      *bool  `json:"budgetAlerts,omitempty"`
	AnomalyAlerts     *bool  `json:"anomalyAlerts,omitempty"`
	MetricAlerts      *bool  `json:"metricAlerts,omitempty"`
	ResourceDrift     *bool  `json:"resourceDrift,omitempty"`
	WorkflowPages     *bool  `json:"workflowPages,omitempty"`
	ProviderIncidents *bool  `json:"providerIncidents,omitempty"`
	ExpiryAlerts      *bool  `json:"expiryAlerts,omitempty"`
	LogMatchAlerts    *bool  `json:"logMatchAlerts,omitempty"`
	WeeklyDigest      *bool  `json:"weeklyDigest,omitempty"`
}

// SlackChannelUpdate is the `SlackChannelUpdate` schema.
type SlackChannelUpdate struct {
	SyncIncidents     *bool `json:"syncIncidents,omitempty"`
	BudgetAlerts      *bool `json:"budgetAlerts,omitempty"`
	AnomalyAlerts     *bool `json:"anomalyAlerts,omitempty"`
	MetricAlerts      *bool `json:"metricAlerts,omitempty"`
	ResourceDrift     *bool `json:"resourceDrift,omitempty"`
	WorkflowPages     *bool `json:"workflowPages,omitempty"`
	ProviderIncidents *bool `json:"providerIncidents,omitempty"`
	ExpiryAlerts      *bool `json:"expiryAlerts,omitempty"`
	LogMatchAlerts    *bool `json:"logMatchAlerts,omitempty"`
	WeeklyDigest      *bool `json:"weeklyDigest,omitempty"`
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

// TabTarget is the `TabTarget` schema.
type TabTarget struct {
	// Kind: One of "dashboard", "account", "resource", "agents", "costs",
	// "savings", "graph", "workflows", "deployments", "chat".
	Kind           string      `json:"kind"`
	DashboardID    *string     `json:"dashboardId,omitempty"`
	AccountID      *string     `json:"accountId,omitempty"`
	ResourceID     *ResourceID `json:"resourceId,omitempty"`
	ConversationID *string     `json:"conversationId,omitempty"`
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

// CreateAccountResponseSyncError is an object the spec declares inline.
type CreateAccountResponseSyncError struct {
	Message string `json:"message"`
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

// DeploymentRunInputError is an object the spec declares inline.
type DeploymentRunInputError struct {
	Message string `json:"message"`
}

// FieldActionResponseOption is an object the spec declares inline.
type FieldActionResponseOption struct {
	ID    string `json:"id"`
	Label string `json:"label"`
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

// ShowbackReportCentres is an object the spec declares inline.
type ShowbackReportCentres struct {
	// CostCentreID: Null for the synthetic "Unallocated" bucket.
	CostCentreID *string `json:"costCentreId"`
	Name         string  `json:"name"`
	// Totals: Currency code → amount in the currency's major unit.
	Totals map[string]float64 `json:"totals"`
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

// CostsAnomaliesResponse is an object the spec declares inline.
type CostsAnomaliesResponse struct {
	Anomalies []CostAnomaly `json:"anomalies"`
}

// CostsStatusResponse is an object the spec declares inline.
type CostsStatusResponse struct {
	Accounts []CostAccountStatus `json:"accounts"`
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

// MsteamsTestResponse is an object the spec declares inline.
type MsteamsTestResponse struct {
	OK           bool  `json:"ok"`
	WebhookCount int64 `json:"webhookCount"`
	Attempted    int64 `json:"attempted"`
	Succeeded    int64 `json:"succeeded"`
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

// ResourcesCreateCostEstimateResponse is an object the spec declares inline.
type ResourcesCreateCostEstimateResponse struct {
	Estimate JSONObject `json:"estimate"`
}

// ResourcesNoSqlcommandResponse is an object the spec declares inline.
type ResourcesNoSqlcommandResponse struct {
	Result JSONObject `json:"result"`
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
