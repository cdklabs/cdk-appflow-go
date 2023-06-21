package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2"
)

// Experimental.
type S3Catalog struct {
	// Experimental.
	Database awscdkgluealpha.Database `field:"required" json:"database" yaml:"database"`
	// Experimental.
	TablePrefix *string `field:"required" json:"tablePrefix" yaml:"tablePrefix"`
}

