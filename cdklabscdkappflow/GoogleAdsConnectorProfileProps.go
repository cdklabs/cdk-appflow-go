package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Experimental.
type GoogleAdsConnectorProfileProps struct {
	// TODO: think if this should be here as not all connector profiles have that.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Experimental.
	DeveloperToken awscdk.SecretValue `field:"required" json:"developerToken" yaml:"developerToken"`
	// Experimental.
	OAuth *GoogleAdsOAuthSettings `field:"required" json:"oAuth" yaml:"oAuth"`
	// Experimental.
	ManagerID awscdk.SecretValue `field:"optional" json:"managerID" yaml:"managerID"`
}

