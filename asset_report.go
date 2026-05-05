package rsvcmodel

// Domain is the UI column grouping for a finding.
type Domain string

const (
	DomainFinOps        Domain = "finops"
	DomainObservability Domain = "observability"
	DomainCompliance    Domain = "compliance"
)

// AssetKey is the canonical key for an asset row in the UI.
type AssetKey string

const (
	AssetEC2            AssetKey = "ec2"
	AssetS3             AssetKey = "s3"
	AssetVPC            AssetKey = "vpc"
	AssetSecurityGroups AssetKey = "securitygroups"
	AssetIAM            AssetKey = "iam"
	AssetKMS            AssetKey = "kms"
	AssetCloudTrail     AssetKey = "cloudtrail"
	AssetGuardDuty      AssetKey = "guardduty"
	AssetRoute53        AssetKey = "route53"
	AssetEKS            AssetKey = "eks"
	AssetELB            AssetKey = "elb"
	AssetRDS            AssetKey = "rds"
	AssetLambda         AssetKey = "lambda"
	AssetECR            AssetKey = "ecr"
	AssetSecretsManager AssetKey = "secretsmanager"
	AssetAccount        AssetKey = "account"
)

type DomainStatus string

const (
	StatusAct  DomainStatus = "Act"
	StatusGood DomainStatus = "Good"
	StatusNA   DomainStatus = "-NA-"
)

type CostSummary struct {
	WindowDays     int                `json:"windowDays"`
	TotalUSD       float64            `json:"totalUsd"`
	AnnualizedUSD  float64            `json:"annualizedUsd,omitempty"`
	ByServiceLabel map[string]float64 `json:"byServiceLabel,omitempty"`
}

// AssetRow is the primary unit the UI renders in “Impacted Assets”.
type AssetRow struct {
	Key         AssetKey         `json:"key"`
	DisplayName string           `json:"displayName"`
	Subtitle    string           `json:"subtitle,omitempty"`
	EntityCounts map[string]int  `json:"entityCounts,omitempty"`
	Cost        CostSummary      `json:"cost"`
	DomainStatus map[Domain]DomainStatus `json:"domainStatus,omitempty"`
	Findings     []*Finding      `json:"findings,omitempty"`
	FindingsCount int            `json:"findingsCount"`
}

