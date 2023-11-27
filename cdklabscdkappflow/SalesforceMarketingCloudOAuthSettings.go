package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type SalesforceMarketingCloudOAuthSettings struct {
	// Experimental.
	Endpoints *SalesforceMarketingCloudOAuthEndpoints `field:"required" json:"endpoints" yaml:"endpoints"`
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Flow *SalesforceMarketingCloudFlowSettings `field:"optional" json:"flow" yaml:"flow"`
}

