package cdklabscdkappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow"
)

// A class that represents a Salesforce Marketing Cloud Source.
// Experimental.
type SalesforceMarketingCloudSource interface {
	ISource
	// The AppFlow type of the connector that this source is implemented for.
	// Experimental.
	ConnectorType() ConnectorType
	// Experimental.
	Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty
}

// The jsii proxy struct for SalesforceMarketingCloudSource
type jsiiProxy_SalesforceMarketingCloudSource struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_SalesforceMarketingCloudSource) ConnectorType() ConnectorType {
	var returns ConnectorType
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSalesforceMarketingCloudSource(props *SalesforceMarketingCloudSourceProps) SalesforceMarketingCloudSource {
	_init_.Initialize()

	if err := validateNewSalesforceMarketingCloudSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SalesforceMarketingCloudSource{}

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSalesforceMarketingCloudSource_Override(s SalesforceMarketingCloudSource, props *SalesforceMarketingCloudSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SalesforceMarketingCloudSource) Bind(scope IFlow) *awsappflow.CfnFlow_SourceFlowConfigProperty {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsappflow.CfnFlow_SourceFlowConfigProperty

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

