package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2"
)

// Experimental.
type S3Catalog struct {
	// The AWS Glue database that will contain the tables created when the flow executes.
	// Experimental.
	Database awscdkgluealpha.IDatabase `field:"required" json:"database" yaml:"database"`
	// The prefix for the tables created in the AWS Glue database.
	// Experimental.
	TablePrefix *string `field:"required" json:"tablePrefix" yaml:"tablePrefix"`
	// The IAM Role that will be used for data catalog operations.
	// Default: - A new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

