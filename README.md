# rsvcmodel

`rsvcmodel` provides a standardized, normalized data contract for cross-cloud security, observability, and finops findings. It serves as one of the core schema library for the `rtifact` ecosystem.

## Usage

This package provides standard `Segment` and `Severity` enums alongside the `Finding` struct. It is designed to be easily serialized to JSON across multiple microservices or CLI tools.

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
