package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Experimental.
type SalesforceOAuthRefreshTokenGrantFlow struct {
	// Experimental.
	Client awssecretsmanager.ISecret `field:"optional" json:"client" yaml:"client"`
	// Experimental.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

