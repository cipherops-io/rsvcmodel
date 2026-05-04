package rsvcmodel

import (
	"testing"
)

func TestNewFinding(t *testing.T) {
	tests := []struct {
		name          string
		segment       Segment
		title         string
		category      string
		items         []string
		expectedCount int
	}{
		{
			name:          "With items",
			segment:       SegmentSecurity,
			title:         "S3 Bucket Public",
			category:      "security",
			items:         []string{"my-bucket", "other-bucket"},
			expectedCount: 2,
		},
		{
			name:          "Without items",
			segment:       SegmentFinOps,
			title:         "Idle EC2 Instances",
			category:      "cost_optimizing",
			items:         nil,
			expectedCount: 0,
		},
		{
			name:          "Empty items slice",
			segment:       SegmentObservability,
			title:         "Missing tags",
			category:      "tagging",
			items:         []string{},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFinding(tt.segment, tt.title, tt.category, tt.items)

			if f.Segment != tt.segment {
				t.Errorf("expected segment %q, got %q", tt.segment, f.Segment)
			}
			if f.Title != tt.title {
				t.Errorf("expected title %q, got %q", tt.title, f.Title)
			}
			if f.Category != tt.category {
				t.Errorf("expected category %q, got %q", tt.category, f.Category)
			}
			if len(f.Items) != len(tt.items) {
				t.Errorf("expected %d items, got %d", len(tt.items), len(f.Items))
			}
			if f.Count != tt.expectedCount {
				t.Errorf("expected count %d, got %d", tt.expectedCount, f.Count)
			}
		})
	}
}
