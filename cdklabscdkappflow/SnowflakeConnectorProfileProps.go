package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a Snowflake connectorprofile.
// Experimental.
type SnowflakeConnectorProfileProps struct {
	// TODO: think if this should be here as not all connector profiles have that.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Experimental.
	BasicAuth *SnowflakeBasicAuthSettings `field:"required" json:"basicAuth" yaml:"basicAuth"`
	// The name of the Snowflake database.
	// Experimental.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Experimental.
	Location *S3Location `field:"required" json:"location" yaml:"location"`
	// The name of the Snowflake stage.
	// Experimental.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// The name of the Snowflake warehouse.
	// Experimental.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// Details of the Snowflake Storage Integration.
	//
	// When provided, this construct will automatically create an IAM Role allowing access to the S3 Bucket which will be available as a [integrationROle property]{@link SnowflakeConnectorProfile#integrationRole}
	//
	// For details of the integration see {@link https://docs.snowflake.com/en/user-guide/data-load-s3-config-storage-integration}
	// Experimental.
	Integration *SnowflakeStorageIntegration `field:"optional" json:"integration" yaml:"integration"`
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The name of the Snowflake schema.
	// Default: PUBLIC.
	//
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

