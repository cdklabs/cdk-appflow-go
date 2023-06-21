package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/internal"
)

// Experimental.
type IConnectorProfile interface {
	awscdk.IResource
	// Experimental.
	Arn() *string
	// Experimental.
	Credentials() awssecretsmanager.ISecret
	// Experimental.
	Name() *string
}

// The jsii proxy for IConnectorProfile
type jsiiProxy_IConnectorProfile struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IConnectorProfile) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorProfile) Credentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

