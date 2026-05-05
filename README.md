# rsvcmodel

`rsvcmodel` provides a standardized, normalized data contract for cross-cloud security, observability, and finops findings. It serves as one of the core schema libraries for the `rtifact` ecosystem.

## Features

- **Normalized Findings**: A unified structure (`Finding`) to represent issues across different cloud providers (AWS, GCP, Azure).
- **Execution Metadata**: Tracking structures (`RunMetadata`) for discovery and scanning jobs.
- **Asset Reporting**: Standardized UI models (`AssetRow`, `CostSummary`, `DomainStatus`) to render impacted assets consistently.

## Usage

This package provides standard `Segment` and `Severity` enums alongside the `Finding` struct. It is designed to be easily serialized to JSON across multiple microservices or CLI tools.

### Example: Creating a Finding

```go
package main

import (
	"encoding/json"
	"fmt"
	
	"github.com/cipherops-io/rsvcmodel"
)

func main() {
	// Create a new normalized finding
	finding := rsvcmodel.NewFinding(
		rsvcmodel.SegmentSecurity,
		"S3 Bucket Public Access",
		"security",
		[]string{"my-public-bucket"},
	)
	
	finding.Severity = rsvcmodel.SeverityCritical
	finding.Provider = "aws"
	
	// Convert to JSON
	bytes, _ := json.MarshalIndent(finding, "", "  ")
	fmt.Println(string(bytes))
}
```

## Data Types

### Core Models

- **`Finding`**: The normalized output unit across all clouds and segments.
- **`RunMetadata`**: Captures top-level information about a discovery execution (e.g., duration, tenant, region).
- **`AssetRow`**: The primary unit used to render "Impacted Assets" in the UI, combining findings, cost, and domain status.

### Segments
- `SegmentSecurity`
- `SegmentFinOps`
- `SegmentObservability`

### Severities
- `SeverityInfo`
- `SeverityLow`
- `SeverityMedium`
- `SeverityHigh`
- `SeverityCritical`

### Domains
- `DomainFinOps`
- `DomainObservability`
- `DomainCompliance`

### Asset Keys
Canonical keys for asset types: `AssetEC2`, `AssetS3`, `AssetVPC`, `AssetIAM`, etc.
