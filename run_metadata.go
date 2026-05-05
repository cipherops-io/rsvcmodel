package rsvcmodel

import "time"

// RunMetadata captures top-level information about a discovery execution.
type RunMetadata struct {
	TenantID        string    `json:"tenant_id"`
	ProjectID       string    `json:"project_id"`
	DiscoveryType   string    `json:"discovery_type"`
	Mode            string    `json:"mode,omitempty"`
	Region          string    `json:"region,omitempty"`
	StartedAt       time.Time `json:"started_at"`
	CompletedAt     time.Time `json:"completed_at,omitempty"`
	DurationSeconds int64     `json:"duration_seconds,omitempty"`
	CompletenessPct int       `json:"completeness_pct,omitempty"`
}

// StartRunMetadata initializes a new RunMetadata instance, recording the start time.
func StartRunMetadata(tenantID, projectID, discoveryType string, now time.Time) RunMetadata {
	return RunMetadata{
		TenantID:      tenantID,
		ProjectID:     projectID,
		DiscoveryType: discoveryType,
		StartedAt:     now.UTC(),
	}
}

// Complete finalizes the execution metadata, setting the completion time and calculating the duration.
func (m *RunMetadata) Complete(now time.Time) {
	finishedAt := now.UTC()
	m.CompletedAt = finishedAt
	m.DurationSeconds = int64(finishedAt.Sub(m.StartedAt).Seconds())
	if m.DurationSeconds < 0 {
		m.DurationSeconds = 0
	}
}
