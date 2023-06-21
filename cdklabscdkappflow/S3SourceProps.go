package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type S3SourceProps struct {
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Experimental.
	Format *S3InputFormat `field:"optional" json:"format" yaml:"format"`
}

