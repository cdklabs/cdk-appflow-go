package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type SalesforceMarketingCloudOAuthClientSettings struct {
	// Experimental.
	ClientId awscdk.SecretValue `field:"required" json:"clientId" yaml:"clientId"`
	// Experimental.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

