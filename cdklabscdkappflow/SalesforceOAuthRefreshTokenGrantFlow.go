package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Experimental.
type SalesforceOAuthRefreshTokenGrantFlow struct {
	// Experimental.
	Client awssecretsmanager.ISecret `field:"optional" json:"client" yaml:"client"`
	// Experimental.
	RefreshToken awscdk.SecretValue `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

