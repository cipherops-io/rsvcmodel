// Package rsvcmodel provides a normalized data contract for cloud findings.
// It defines the core structures used across various cloud providers to represent
// security, finops, and observability issues in a unified way.
package rsvcmodel

// Segment categorizes the high-level domain of a finding (e.g., security vs. finops).

type Segment string

const (
	// SegmentObservability indicates findings related to performance, reliability, and tracing.
	SegmentObservability Segment = "observability"
	// SegmentSecurity indicates findings related to vulnerabilities, access control, and compliance.
	SegmentSecurity Segment = "security"
	// SegmentFinOps indicates findings related to cost optimization and resource waste.
	SegmentFinOps Segment = "finops"
)

// Severity indicates the urgency and critical nature of a finding.
type Severity string

const (
	// SeverityInfo indicates informational items that require no immediate action.
	SeverityInfo Severity = "info"
	// SeverityLow indicates low-priority issues that can be addressed when time permits.
	SeverityLow Severity = "low"
	// SeverityMedium indicates moderate issues that should be addressed in standard planning cycles.
	SeverityMedium Severity = "medium"
	// SeverityHigh indicates significant issues that require prompt attention.
	SeverityHigh Severity = "high"
	// SeverityCritical indicates severe issues that pose immediate risks and require emergency remediation.
	SeverityCritical Severity = "critical"
)

// Finding is the normalized output unit across all clouds and segments.
// It aggregates flagged resources and their metadata for a specific check or rule.
type Finding struct {
	// ID is a unique identifier for this specific finding output.
	ID string `json:"id,omitempty"`
	// Title is the human-readable name of the check or rule that triggered this finding.
	Title string `json:"title"`
	// Category is the provider-specific sub-category of the finding.
	Category string `json:"category"`
	// Segment maps the finding to a high-level domain (e.g., finops, security).
	Segment Segment `json:"segment"`
	// Severity is the standardized threat or urgency level of the finding.
	Severity Severity `json:"severity,omitempty"`
	// Provider indicates the source of the finding (e.g., aws, gcp, azure).
	Provider string `json:"provider,omitempty"`
	// Summary provides a brief description or context of what the finding means.
	Summary string `json:"summary,omitempty"`
	// Count represents the total number of resources flagged in this finding.
	Count int `json:"count,omitempty"`
	// Items contains human-readable names or pretty identifiers for the flagged resources.
	Items []string `json:"items,omitempty"`
	// ResourceIDs contains the exact native cloud IDs (e.g., ARNs, UUIDs) for the flagged resources.
	ResourceIDs []string `json:"resourceIds,omitempty"`
	// Metadata holds arbitrary key-value pairs for additional context, such as deep links or native properties.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// NewFinding creates a new Finding with the basic required fields initialized.
// It automatically sets the Count field based on the length of the provided items slice.
func NewFinding(segment Segment, title, category string, items []string) *Finding {
	f := &Finding{
		Title:    title,
		Category: category,
		Segment:  segment,
		Items:    items,
	}

	if len(items) > 0 {
		f.Count = len(items)
	}

	return f
}
