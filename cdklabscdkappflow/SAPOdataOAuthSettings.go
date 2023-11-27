package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type SAPOdataOAuthSettings struct {
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Experimental.
	Endpoints *SAPOdataOAuthEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Experimental.
	Flow *SAPOdataOAuthFlows `field:"optional" json:"flow" yaml:"flow"`
}

