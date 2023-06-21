package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type S3Location struct {
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

