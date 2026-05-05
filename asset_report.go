package rsvcmodel

// Domain is the UI column grouping for a finding.
type Domain string

const (
	// DomainFinOps represents the financial operations and cost optimization domain.
	DomainFinOps Domain = "finops"
	// DomainObservability represents the observability, logging, and monitoring domain.
	DomainObservability Domain = "observability"
	// DomainCompliance represents the compliance, regulatory, and security domain.
	DomainCompliance Domain = "compliance"
)

// AssetKey is the canonical key for an asset row in the UI.
type AssetKey string

const (
	// AssetEC2 represents an Amazon EC2 instance.
	AssetEC2 AssetKey = "ec2"
	// AssetS3 represents an Amazon S3 bucket.
	AssetS3 AssetKey = "s3"
	// AssetVPC represents an Amazon Virtual Private Cloud.
	AssetVPC AssetKey = "vpc"
	// AssetSecurityGroups represents AWS Security Groups.
	AssetSecurityGroups AssetKey = "securitygroups"
	// AssetIAM represents AWS Identity and Access Management resources.
	AssetIAM AssetKey = "iam"
	// AssetKMS represents AWS Key Management Service resources.
	AssetKMS AssetKey = "kms"
	// AssetCloudTrail represents AWS CloudTrail resources.
	AssetCloudTrail AssetKey = "cloudtrail"
	// AssetGuardDuty represents AWS GuardDuty findings or detectors.
	AssetGuardDuty AssetKey = "guardduty"
	// AssetRoute53 represents Amazon Route 53 resources.
	AssetRoute53 AssetKey = "route53"
	// AssetEKS represents Amazon Elastic Kubernetes Service clusters.
	AssetEKS AssetKey = "eks"
	// AssetELB represents Elastic Load Balancers.
	AssetELB AssetKey = "elb"
	// AssetRDS represents Amazon Relational Database Service instances.
	AssetRDS AssetKey = "rds"
	// AssetLambda represents AWS Lambda functions.
	AssetLambda AssetKey = "lambda"
	// AssetECR represents Amazon Elastic Container Registry repositories.
	AssetECR AssetKey = "ecr"
	// AssetSecretsManager represents AWS Secrets Manager secrets.
	AssetSecretsManager AssetKey = "secretsmanager"
	// AssetAccount represents a cloud account (e.g., AWS Account).
	AssetAccount AssetKey = "account"
)

// DomainStatus indicates the high-level health or actionability of a domain for an asset.
type DomainStatus string

const (
	// StatusAct indicates that there are active issues requiring attention.
	StatusAct DomainStatus = "Act"
	// StatusGood indicates that the domain is healthy with no significant issues.
	StatusGood DomainStatus = "Good"
	// StatusNA indicates that the domain is not applicable to the asset.
	StatusNA DomainStatus = "-NA-"
)

// CostSummary aggregates the cost information for an asset over a specific time window.
type CostSummary struct {
	// WindowDays is the number of days over which the cost is calculated.
	WindowDays int `json:"windowDays"`
	// TotalUSD is the total cost incurred during the window in USD.
	TotalUSD float64 `json:"totalUsd"`
	// AnnualizedUSD is the projected annual cost based on the window in USD.
	AnnualizedUSD float64 `json:"annualizedUsd,omitempty"`
	// ByServiceLabel breaks down the total cost by individual service labels.
	ByServiceLabel map[string]float64 `json:"byServiceLabel,omitempty"`
}

// AssetRow is the primary unit the UI renders in “Impacted Assets”.
// It provides a consolidated view of an asset's findings, cost, and domain status.
type AssetRow struct {
	// Key is the unique identifier for the type of asset (e.g., ec2, s3).
	Key AssetKey `json:"key"`
	// DisplayName is the human-readable name of the asset.
	DisplayName string `json:"displayName"`
	// Subtitle provides additional context or identification for the asset.
	Subtitle string `json:"subtitle,omitempty"`
	// EntityCounts holds counts of specific entities associated with the asset.
	EntityCounts map[string]int `json:"entityCounts,omitempty"`
	// Cost contains the cost summary for the asset.
	Cost CostSummary `json:"cost"`
	// DomainStatus maps each domain to its current health or action status.
	DomainStatus map[Domain]DomainStatus `json:"domainStatus,omitempty"`
	// Findings is a list of specific issues or observations associated with the asset.
	Findings []*Finding `json:"findings,omitempty"`
	// FindingsCount is the total number of findings for the asset.
	FindingsCount int `json:"findingsCount"`
}
