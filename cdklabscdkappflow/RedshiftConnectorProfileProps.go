package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2"
)

// Experimental.
type RedshiftConnectorProfileProps struct {
	// TODO: think if this should be here as not all connector profiles have that.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	BasicAuth *RedshiftConnectorBasicCredentials `field:"required" json:"basicAuth" yaml:"basicAuth"`
	// The Redshift cluster to use this connector profile with.
	// Experimental.
	Cluster awscdkredshiftalpha.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database which the RedshiftConnectorProfile will be working with.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// An intermediate location for the data retrieved from the flow source that will be further transferred to the Redshfit database.
	// Experimental.
	IntermediateLocation *S3Location `field:"required" json:"intermediateLocation" yaml:"intermediateLocation"`
	// An IAM Role that the Redshift cluster will assume to get data from the intermiediate S3 Bucket.
	// Experimental.
	BucketAccessRole awsiam.IRole `field:"optional" json:"bucketAccessRole" yaml:"bucketAccessRole"`
	// An IAM Role that AppFlow will assume to interact with the Redshift cluster's Data API.
	// Experimental.
	DataApiRole awsiam.IRole `field:"optional" json:"dataApiRole" yaml:"dataApiRole"`
}

