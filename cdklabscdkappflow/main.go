// @cdklabs/cdk-appflow
package cdklabscdkappflow

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlBasicAuthSettings",
		reflect.TypeOf((*AmazonRdsForPostgreSqlBasicAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlConnectorProfile",
		reflect.TypeOf((*AmazonRdsForPostgreSqlConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonRdsForPostgreSqlConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlConnectorProfileProps",
		reflect.TypeOf((*AmazonRdsForPostgreSqlConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlDestination",
		reflect.TypeOf((*AmazonRdsForPostgreSqlDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonRdsForPostgreSqlDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlDestinationProps",
		reflect.TypeOf((*AmazonRdsForPostgreSqlDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.AmazonRdsForPostgreSqlObject",
		reflect.TypeOf((*AmazonRdsForPostgreSqlObject)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.AsanaConnectorProfile",
		reflect.TypeOf((*AsanaConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_AsanaConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.AsanaConnectorProfileProps",
		reflect.TypeOf((*AsanaConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.AsanaSource",
		reflect.TypeOf((*AsanaSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_AsanaSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.AsanaSourceProps",
		reflect.TypeOf((*AsanaSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.ConnectionMode",
		reflect.TypeOf((*ConnectionMode)(nil)).Elem(),
		map[string]interface{}{
			"PUBLIC": ConnectionMode_PUBLIC,
			"PRIVATE": ConnectionMode_PRIVATE,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.ConnectorAuthenticationType",
		reflect.TypeOf((*ConnectorAuthenticationType)(nil)).Elem(),
		map[string]interface{}{
			"APIKEY": ConnectorAuthenticationType_APIKEY,
			"BASIC": ConnectorAuthenticationType_BASIC,
			"CUSTOM": ConnectorAuthenticationType_CUSTOM,
			"OAUTH2": ConnectorAuthenticationType_OAUTH2,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ConnectorProfileBase",
		reflect.TypeOf((*ConnectorProfileBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_ConnectorProfileBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectorProfile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ConnectorProfileProps",
		reflect.TypeOf((*ConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ConnectorType",
		reflect.TypeOf((*ConnectorType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "asProfileConnectorLabel", GoGetter: "AsProfileConnectorLabel"},
			_jsii_.MemberProperty{JsiiProperty: "asProfileConnectorType", GoGetter: "AsProfileConnectorType"},
			_jsii_.MemberProperty{JsiiProperty: "asTaskConnectorOperatorOrigin", GoGetter: "AsTaskConnectorOperatorOrigin"},
			_jsii_.MemberProperty{JsiiProperty: "isCustom", GoGetter: "IsCustom"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_ConnectorType{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.DataPullConfig",
		reflect.TypeOf((*DataPullConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.DataPullMode",
		reflect.TypeOf((*DataPullMode)(nil)).Elem(),
		map[string]interface{}{
			"COMPLETE": DataPullMode_COMPLETE,
			"INCREMENTAL": DataPullMode_INCREMENTAL,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ErrorHandlingConfiguration",
		reflect.TypeOf((*ErrorHandlingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.EventBridgeDestination",
		reflect.TypeOf((*EventBridgeDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_EventBridgeDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.EventBridgeDestinationProps",
		reflect.TypeOf((*EventBridgeDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.EventSources",
		reflect.TypeOf((*EventSources)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EventSources{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.Field",
		reflect.TypeOf((*Field)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.Filter",
		reflect.TypeOf((*Filter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
		},
		func() interface{} {
			j := jsiiProxy_Filter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OperationBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFilter)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.FilterCondition",
		reflect.TypeOf((*FilterCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "field", GoGetter: "Field"},
			_jsii_.MemberProperty{JsiiProperty: "filter", GoGetter: "Filter"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
		},
		func() interface{} {
			return &jsiiProxy_FilterCondition{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.FlowBase",
		reflect.TypeOf((*FlowBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionRecordsProcessed", GoMethod: "MetricFlowExecutionRecordsProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsFailed", GoMethod: "MetricFlowExecutionsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsStarted", GoMethod: "MetricFlowExecutionsStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsSucceeded", GoMethod: "MetricFlowExecutionsSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionTime", GoMethod: "MetricFlowExecutionTime"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onRunCompleted", GoMethod: "OnRunCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onRunStarted", GoMethod: "OnRunStarted"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_FlowBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlow)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.FlowBaseProps",
		reflect.TypeOf((*FlowBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.FlowProps",
		reflect.TypeOf((*FlowProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.FlowStatus",
		reflect.TypeOf((*FlowStatus)(nil)).Elem(),
		map[string]interface{}{
			"ACTIVE": FlowStatus_ACTIVE,
			"SUSPENDED": FlowStatus_SUSPENDED,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.FlowType",
		reflect.TypeOf((*FlowType)(nil)).Elem(),
		map[string]interface{}{
			"EVENT": FlowType_EVENT,
			"ON_DEMAND": FlowType_ON_DEMAND,
			"SCHEDULED": FlowType_SCHEDULED,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.GoogleAdsApiVersion",
		reflect.TypeOf((*GoogleAdsApiVersion)(nil)).Elem(),
		map[string]interface{}{
			"V16": GoogleAdsApiVersion_V16,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.GoogleAdsConnectorProfile",
		reflect.TypeOf((*GoogleAdsConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleAdsConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAdsConnectorProfileProps",
		reflect.TypeOf((*GoogleAdsConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAdsOAuthEndpoints",
		reflect.TypeOf((*GoogleAdsOAuthEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAdsOAuthFlow",
		reflect.TypeOf((*GoogleAdsOAuthFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAdsOAuthSettings",
		reflect.TypeOf((*GoogleAdsOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAdsRefreshTokenGrantFlow",
		reflect.TypeOf((*GoogleAdsRefreshTokenGrantFlow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.GoogleAdsSource",
		reflect.TypeOf((*GoogleAdsSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleAdsSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAdsSourceProps",
		reflect.TypeOf((*GoogleAdsSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.GoogleAnalytics4ApiVersion",
		reflect.TypeOf((*GoogleAnalytics4ApiVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1BETA": GoogleAnalytics4ApiVersion_V1BETA,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.GoogleAnalytics4ConnectorProfile",
		reflect.TypeOf((*GoogleAnalytics4ConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleAnalytics4ConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAnalytics4ConnectorProfileProps",
		reflect.TypeOf((*GoogleAnalytics4ConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAnalytics4OAuthEndpoints",
		reflect.TypeOf((*GoogleAnalytics4OAuthEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAnalytics4OAuthFlow",
		reflect.TypeOf((*GoogleAnalytics4OAuthFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAnalytics4OAuthSettings",
		reflect.TypeOf((*GoogleAnalytics4OAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAnalytics4RefreshTokenGrantFlow",
		reflect.TypeOf((*GoogleAnalytics4RefreshTokenGrantFlow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.GoogleAnalytics4Source",
		reflect.TypeOf((*GoogleAnalytics4Source)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleAnalytics4Source{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleAnalytics4SourceProps",
		reflect.TypeOf((*GoogleAnalytics4SourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.GoogleBigQueryApiVersion",
		reflect.TypeOf((*GoogleBigQueryApiVersion)(nil)).Elem(),
		map[string]interface{}{
			"V2": GoogleBigQueryApiVersion_V2,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.GoogleBigQueryConnectorProfile",
		reflect.TypeOf((*GoogleBigQueryConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleBigQueryConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleBigQueryConnectorProfileProps",
		reflect.TypeOf((*GoogleBigQueryConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleBigQueryOAuthEndpoints",
		reflect.TypeOf((*GoogleBigQueryOAuthEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleBigQueryOAuthFlow",
		reflect.TypeOf((*GoogleBigQueryOAuthFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleBigQueryOAuthSettings",
		reflect.TypeOf((*GoogleBigQueryOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleBigQueryObject",
		reflect.TypeOf((*GoogleBigQueryObject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleBigQueryRefreshTokenGrantFlow",
		reflect.TypeOf((*GoogleBigQueryRefreshTokenGrantFlow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.GoogleBigQuerySource",
		reflect.TypeOf((*GoogleBigQuerySource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleBigQuerySource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.GoogleBigQuerySourceProps",
		reflect.TypeOf((*GoogleBigQuerySourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IConnectorProfile",
		reflect.TypeOf((*IConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IDestination",
		reflect.TypeOf((*IDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_IDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVertex)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IFilter",
		reflect.TypeOf((*IFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_IFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperation)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IFlow",
		reflect.TypeOf((*IFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionRecordsProcessed", GoMethod: "MetricFlowExecutionRecordsProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsFailed", GoMethod: "MetricFlowExecutionsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsStarted", GoMethod: "MetricFlowExecutionsStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsSucceeded", GoMethod: "MetricFlowExecutionsSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionTime", GoMethod: "MetricFlowExecutionTime"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onRunCompleted", GoMethod: "OnRunCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onRunStarted", GoMethod: "OnRunStarted"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_IFlow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IMapping",
		reflect.TypeOf((*IMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_IMapping{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperation)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IOperation",
		reflect.TypeOf((*IOperation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IOperation{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.ISource",
		reflect.TypeOf((*ISource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_ISource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVertex)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.ITask",
		reflect.TypeOf((*ITask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ITask{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.ITransform",
		reflect.TypeOf((*ITransform)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ITransform{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperation)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IValidation",
		reflect.TypeOf((*IValidation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_IValidation{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperation)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-appflow.IVertex",
		reflect.TypeOf((*IVertex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			return &jsiiProxy_IVertex{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.JdbcDriver",
		reflect.TypeOf((*JdbcDriver)(nil)).Elem(),
		map[string]interface{}{
			"POSTGRES": JdbcDriver_POSTGRES,
			"MYSQL": JdbcDriver_MYSQL,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleBasicAuthSettings",
		reflect.TypeOf((*JdbcSmallDataScaleBasicAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfile",
		reflect.TypeOf((*JdbcSmallDataScaleConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_JdbcSmallDataScaleConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleConnectorProfileProps",
		reflect.TypeOf((*JdbcSmallDataScaleConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleObject",
		reflect.TypeOf((*JdbcSmallDataScaleObject)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleSource",
		reflect.TypeOf((*JdbcSmallDataScaleSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_JdbcSmallDataScaleSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.JdbcSmallDataScaleSourceProps",
		reflect.TypeOf((*JdbcSmallDataScaleSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.MailchimpApiVersion",
		reflect.TypeOf((*MailchimpApiVersion)(nil)).Elem(),
		map[string]interface{}{
			"V3": MailchimpApiVersion_V3,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MailchimpConnectorProfile",
		reflect.TypeOf((*MailchimpConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_MailchimpConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MailchimpConnectorProfileProps",
		reflect.TypeOf((*MailchimpConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MailchimpSource",
		reflect.TypeOf((*MailchimpSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_MailchimpSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MailchimpSourceProps",
		reflect.TypeOf((*MailchimpSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MapAllConfig",
		reflect.TypeOf((*MapAllConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.Mapping",
		reflect.TypeOf((*Mapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Mapping{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OperationBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMapping)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MarketoConnectorProfile",
		reflect.TypeOf((*MarketoConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_MarketoConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MarketoConnectorProfileProps",
		reflect.TypeOf((*MarketoConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MarketoInstanceUrlBuilder",
		reflect.TypeOf((*MarketoInstanceUrlBuilder)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MarketoInstanceUrlBuilder{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MarketoOAuthClientCredentialsFlow",
		reflect.TypeOf((*MarketoOAuthClientCredentialsFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MarketoOAuthFlow",
		reflect.TypeOf((*MarketoOAuthFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MarketoOAuthSettings",
		reflect.TypeOf((*MarketoOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MarketoSource",
		reflect.TypeOf((*MarketoSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_MarketoSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MarketoSourceProps",
		reflect.TypeOf((*MarketoSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365ApiUrlBuilder",
		reflect.TypeOf((*MicrosoftDynamics365ApiUrlBuilder)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MicrosoftDynamics365ApiUrlBuilder{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365ApiVersion",
		reflect.TypeOf((*MicrosoftDynamics365ApiVersion)(nil)).Elem(),
		map[string]interface{}{
			"V9_2": MicrosoftDynamics365ApiVersion_V9_2,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365ConnectorProfile",
		reflect.TypeOf((*MicrosoftDynamics365ConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_MicrosoftDynamics365ConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365ConnectorProfileProps",
		reflect.TypeOf((*MicrosoftDynamics365ConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365OAuthEndpointsSettings",
		reflect.TypeOf((*MicrosoftDynamics365OAuthEndpointsSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365OAuthFlow",
		reflect.TypeOf((*MicrosoftDynamics365OAuthFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365OAuthSettings",
		reflect.TypeOf((*MicrosoftDynamics365OAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365RefreshTokenGrantFlow",
		reflect.TypeOf((*MicrosoftDynamics365RefreshTokenGrantFlow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365Source",
		reflect.TypeOf((*MicrosoftDynamics365Source)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_MicrosoftDynamics365Source{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365SourceProps",
		reflect.TypeOf((*MicrosoftDynamics365SourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MicrosoftDynamics365TokenUrlBuilder",
		reflect.TypeOf((*MicrosoftDynamics365TokenUrlBuilder)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MicrosoftDynamics365TokenUrlBuilder{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineApiVersion",
		reflect.TypeOf((*MicrosoftSharepointOnlineApiVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1": MicrosoftSharepointOnlineApiVersion_V1,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineConnectorProfile",
		reflect.TypeOf((*MicrosoftSharepointOnlineConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_MicrosoftSharepointOnlineConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineConnectorProfileProps",
		reflect.TypeOf((*MicrosoftSharepointOnlineConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineOAuthEndpointsSettings",
		reflect.TypeOf((*MicrosoftSharepointOnlineOAuthEndpointsSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineOAuthFlow",
		reflect.TypeOf((*MicrosoftSharepointOnlineOAuthFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineOAuthSettings",
		reflect.TypeOf((*MicrosoftSharepointOnlineOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineObject",
		reflect.TypeOf((*MicrosoftSharepointOnlineObject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineRefreshTokenGrantFlow",
		reflect.TypeOf((*MicrosoftSharepointOnlineRefreshTokenGrantFlow)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineSource",
		reflect.TypeOf((*MicrosoftSharepointOnlineSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_MicrosoftSharepointOnlineSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineSourceProps",
		reflect.TypeOf((*MicrosoftSharepointOnlineSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.MicrosoftSharepointOnlineTokenUrlBuilder",
		reflect.TypeOf((*MicrosoftSharepointOnlineTokenUrlBuilder)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MicrosoftSharepointOnlineTokenUrlBuilder{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.OAuth2GrantType",
		reflect.TypeOf((*OAuth2GrantType)(nil)).Elem(),
		map[string]interface{}{
			"CLIENT_CREDENTIALS": OAuth2GrantType_CLIENT_CREDENTIALS,
			"AUTHORIZATION_CODE": OAuth2GrantType_AUTHORIZATION_CODE,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.OnDemandFlow",
		reflect.TypeOf((*OnDemandFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionRecordsProcessed", GoMethod: "MetricFlowExecutionRecordsProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsFailed", GoMethod: "MetricFlowExecutionsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsStarted", GoMethod: "MetricFlowExecutionsStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsSucceeded", GoMethod: "MetricFlowExecutionsSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionTime", GoMethod: "MetricFlowExecutionTime"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onRunCompleted", GoMethod: "OnRunCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onRunStarted", GoMethod: "OnRunStarted"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_OnDemandFlow{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FlowBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlow)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.OnDemandFlowProps",
		reflect.TypeOf((*OnDemandFlowProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.OnEventFlow",
		reflect.TypeOf((*OnEventFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionRecordsProcessed", GoMethod: "MetricFlowExecutionRecordsProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsFailed", GoMethod: "MetricFlowExecutionsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsStarted", GoMethod: "MetricFlowExecutionsStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsSucceeded", GoMethod: "MetricFlowExecutionsSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionTime", GoMethod: "MetricFlowExecutionTime"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onDeactivated", GoMethod: "OnDeactivated"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onRunCompleted", GoMethod: "OnRunCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onRunStarted", GoMethod: "OnRunStarted"},
			_jsii_.MemberMethod{JsiiMethod: "onStatus", GoMethod: "OnStatus"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_OnEventFlow{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TriggeredFlowBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlow)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.OnEventFlowProps",
		reflect.TypeOf((*OnEventFlowProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.OnScheduleFlow",
		reflect.TypeOf((*OnScheduleFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionRecordsProcessed", GoMethod: "MetricFlowExecutionRecordsProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsFailed", GoMethod: "MetricFlowExecutionsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsStarted", GoMethod: "MetricFlowExecutionsStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsSucceeded", GoMethod: "MetricFlowExecutionsSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionTime", GoMethod: "MetricFlowExecutionTime"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onDeactivated", GoMethod: "OnDeactivated"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onRunCompleted", GoMethod: "OnRunCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onRunStarted", GoMethod: "OnRunStarted"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_OnScheduleFlow{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TriggeredFlowBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlow)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.OnScheduleFlowProps",
		reflect.TypeOf((*OnScheduleFlowProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.OperationBase",
		reflect.TypeOf((*OperationBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_OperationBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOperation)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.RedshiftConnectorBasicCredentials",
		reflect.TypeOf((*RedshiftConnectorBasicCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.RedshiftConnectorProfile",
		reflect.TypeOf((*RedshiftConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_RedshiftConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.RedshiftConnectorProfileProps",
		reflect.TypeOf((*RedshiftConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.RedshiftDestination",
		reflect.TypeOf((*RedshiftDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_RedshiftDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.RedshiftDestinationObject",
		reflect.TypeOf((*RedshiftDestinationObject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.RedshiftDestinationProps",
		reflect.TypeOf((*RedshiftDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3Catalog",
		reflect.TypeOf((*S3Catalog)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.S3Destination",
		reflect.TypeOf((*S3Destination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_S3Destination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3DestinationProps",
		reflect.TypeOf((*S3DestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3FileAggregation",
		reflect.TypeOf((*S3FileAggregation)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.S3InputFileType",
		reflect.TypeOf((*S3InputFileType)(nil)).Elem(),
		map[string]interface{}{
			"CSV": S3InputFileType_CSV,
			"JSON": S3InputFileType_JSON,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3InputFormat",
		reflect.TypeOf((*S3InputFormat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3Location",
		reflect.TypeOf((*S3Location)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.S3OutputAggregationType",
		reflect.TypeOf((*S3OutputAggregationType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": S3OutputAggregationType_NONE,
			"SINGLE_FILE": S3OutputAggregationType_SINGLE_FILE,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3OutputFilePrefix",
		reflect.TypeOf((*S3OutputFilePrefix)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.S3OutputFilePrefixFormat",
		reflect.TypeOf((*S3OutputFilePrefixFormat)(nil)).Elem(),
		map[string]interface{}{
			"DAY": S3OutputFilePrefixFormat_DAY,
			"HOUR": S3OutputFilePrefixFormat_HOUR,
			"MINUTE": S3OutputFilePrefixFormat_MINUTE,
			"MONTH": S3OutputFilePrefixFormat_MONTH,
			"YEAR": S3OutputFilePrefixFormat_YEAR,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.S3OutputFilePrefixHierarchy",
		reflect.TypeOf((*S3OutputFilePrefixHierarchy)(nil)).Elem(),
		map[string]interface{}{
			"EXECUTION_ID": S3OutputFilePrefixHierarchy_EXECUTION_ID,
			"SCHEMA_VERSION": S3OutputFilePrefixHierarchy_SCHEMA_VERSION,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.S3OutputFilePrefixType",
		reflect.TypeOf((*S3OutputFilePrefixType)(nil)).Elem(),
		map[string]interface{}{
			"FILENAME": S3OutputFilePrefixType_FILENAME,
			"PATH": S3OutputFilePrefixType_PATH,
			"PATH_AND_FILE": S3OutputFilePrefixType_PATH_AND_FILE,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.S3OutputFileType",
		reflect.TypeOf((*S3OutputFileType)(nil)).Elem(),
		map[string]interface{}{
			"CSV": S3OutputFileType_CSV,
			"JSON": S3OutputFileType_JSON,
			"PARQUET": S3OutputFileType_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3OutputFormatting",
		reflect.TypeOf((*S3OutputFormatting)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.S3Source",
		reflect.TypeOf((*S3Source)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_S3Source{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.S3SourceProps",
		reflect.TypeOf((*S3SourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataBasicAuthSettings",
		reflect.TypeOf((*SAPOdataBasicAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SAPOdataConnectorProfile",
		reflect.TypeOf((*SAPOdataConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_SAPOdataConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataConnectorProfileProps",
		reflect.TypeOf((*SAPOdataConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SAPOdataDestination",
		reflect.TypeOf((*SAPOdataDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_SAPOdataDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataDestinationProps",
		reflect.TypeOf((*SAPOdataDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataOAuthEndpoints",
		reflect.TypeOf((*SAPOdataOAuthEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataOAuthFlows",
		reflect.TypeOf((*SAPOdataOAuthFlows)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataOAuthRefreshTokenGrantFlow",
		reflect.TypeOf((*SAPOdataOAuthRefreshTokenGrantFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataOAuthSettings",
		reflect.TypeOf((*SAPOdataOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SAPOdataSource",
		reflect.TypeOf((*SAPOdataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_SAPOdataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataSourceProps",
		reflect.TypeOf((*SAPOdataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SAPOdataSuccessResponseHandlingConfiguration",
		reflect.TypeOf((*SAPOdataSuccessResponseHandlingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SalesforceConnectorProfile",
		reflect.TypeOf((*SalesforceConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_SalesforceConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceConnectorProfileProps",
		reflect.TypeOf((*SalesforceConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.SalesforceDataTransferApi",
		reflect.TypeOf((*SalesforceDataTransferApi)(nil)).Elem(),
		map[string]interface{}{
			"AUTOMATIC": SalesforceDataTransferApi_AUTOMATIC,
			"BULKV2": SalesforceDataTransferApi_BULKV2,
			"REST_SYNC": SalesforceDataTransferApi_REST_SYNC,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SalesforceDestination",
		reflect.TypeOf((*SalesforceDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_SalesforceDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceDestinationProps",
		reflect.TypeOf((*SalesforceDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudApiVersions",
		reflect.TypeOf((*SalesforceMarketingCloudApiVersions)(nil)).Elem(),
		map[string]interface{}{
			"V1": SalesforceMarketingCloudApiVersions_V1,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudConnectorProfile",
		reflect.TypeOf((*SalesforceMarketingCloudConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_SalesforceMarketingCloudConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudConnectorProfileProps",
		reflect.TypeOf((*SalesforceMarketingCloudConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudFlowSettings",
		reflect.TypeOf((*SalesforceMarketingCloudFlowSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudOAuthClientSettings",
		reflect.TypeOf((*SalesforceMarketingCloudOAuthClientSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudOAuthEndpoints",
		reflect.TypeOf((*SalesforceMarketingCloudOAuthEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudOAuthSettings",
		reflect.TypeOf((*SalesforceMarketingCloudOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudSource",
		reflect.TypeOf((*SalesforceMarketingCloudSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_SalesforceMarketingCloudSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceMarketingCloudSourceProps",
		reflect.TypeOf((*SalesforceMarketingCloudSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceOAuthFlow",
		reflect.TypeOf((*SalesforceOAuthFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceOAuthRefreshTokenGrantFlow",
		reflect.TypeOf((*SalesforceOAuthRefreshTokenGrantFlow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceOAuthSettings",
		reflect.TypeOf((*SalesforceOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SalesforceSource",
		reflect.TypeOf((*SalesforceSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_SalesforceSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SalesforceSourceProps",
		reflect.TypeOf((*SalesforceSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ScheduleProperties",
		reflect.TypeOf((*ScheduleProperties)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ServiceNowBasicSettings",
		reflect.TypeOf((*ServiceNowBasicSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ServiceNowConnectorProfile",
		reflect.TypeOf((*ServiceNowConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceNowConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ServiceNowConnectorProfileProps",
		reflect.TypeOf((*ServiceNowConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ServiceNowInstanceUrlBuilder",
		reflect.TypeOf((*ServiceNowInstanceUrlBuilder)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ServiceNowInstanceUrlBuilder{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ServiceNowSource",
		reflect.TypeOf((*ServiceNowSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceNowSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ServiceNowSourceProps",
		reflect.TypeOf((*ServiceNowSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SlackConnectorProfile",
		reflect.TypeOf((*SlackConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_SlackConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SlackConnectorProfileProps",
		reflect.TypeOf((*SlackConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SlackInstanceUrlBuilder",
		reflect.TypeOf((*SlackInstanceUrlBuilder)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SlackInstanceUrlBuilder{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SlackOAuthSettings",
		reflect.TypeOf((*SlackOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SlackSource",
		reflect.TypeOf((*SlackSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_SlackSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SlackSourceProps",
		reflect.TypeOf((*SlackSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SnowflakeBasicAuthSettings",
		reflect.TypeOf((*SnowflakeBasicAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SnowflakeConnectorProfile",
		reflect.TypeOf((*SnowflakeConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "integrationRole", GoGetter: "IntegrationRole"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_SnowflakeConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SnowflakeConnectorProfileProps",
		reflect.TypeOf((*SnowflakeConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.SnowflakeDestination",
		reflect.TypeOf((*SnowflakeDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_SnowflakeDestination{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SnowflakeDestinationObject",
		reflect.TypeOf((*SnowflakeDestinationObject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SnowflakeDestinationProps",
		reflect.TypeOf((*SnowflakeDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.SnowflakeStorageIntegration",
		reflect.TypeOf((*SnowflakeStorageIntegration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.Task",
		reflect.TypeOf((*Task)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorOperator", GoGetter: "ConnectorOperator"},
			_jsii_.MemberProperty{JsiiProperty: "destinationField", GoGetter: "DestinationField"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFields", GoGetter: "SourceFields"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Task{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITask)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.TaskConnectorOperator",
		reflect.TypeOf((*TaskConnectorOperator)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.TaskProperty",
		reflect.TypeOf((*TaskProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.Transform",
		reflect.TypeOf((*Transform)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Transform{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OperationBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITransform)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.TriggerConfig",
		reflect.TypeOf((*TriggerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.TriggerProperties",
		reflect.TypeOf((*TriggerProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.TriggeredFlowBase",
		reflect.TypeOf((*TriggeredFlowBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionRecordsProcessed", GoMethod: "MetricFlowExecutionRecordsProcessed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsFailed", GoMethod: "MetricFlowExecutionsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsStarted", GoMethod: "MetricFlowExecutionsStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionsSucceeded", GoMethod: "MetricFlowExecutionsSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricFlowExecutionTime", GoMethod: "MetricFlowExecutionTime"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onDeactivated", GoMethod: "OnDeactivated"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onRunCompleted", GoMethod: "OnRunCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onRunStarted", GoMethod: "OnRunStarted"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_TriggeredFlowBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FlowBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlow)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.TriggeredFlowBaseProps",
		reflect.TypeOf((*TriggeredFlowBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.Validation",
		reflect.TypeOf((*Validation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
		},
		func() interface{} {
			j := jsiiProxy_Validation{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_OperationBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IValidation)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ValidationAction",
		reflect.TypeOf((*ValidationAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
		},
		func() interface{} {
			return &jsiiProxy_ValidationAction{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ValidationCondition",
		reflect.TypeOf((*ValidationCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "field", GoGetter: "Field"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
		},
		func() interface{} {
			return &jsiiProxy_ValidationCondition{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.WriteOperation",
		reflect.TypeOf((*WriteOperation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ids", GoGetter: "Ids"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_WriteOperation{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-appflow.WriteOperationType",
		reflect.TypeOf((*WriteOperationType)(nil)).Elem(),
		map[string]interface{}{
			"DELETE": WriteOperationType_DELETE,
			"INSERT": WriteOperationType_INSERT,
			"UPDATE": WriteOperationType_UPDATE,
			"UPSERT": WriteOperationType_UPSERT,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ZendeskConnectorProfile",
		reflect.TypeOf((*ZendeskConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileCredentials", GoMethod: "BuildConnectorProfileCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "buildConnectorProfileProperties", GoMethod: "BuildConnectorProfileProperties"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryAddNodeDependency", GoMethod: "TryAddNodeDependency"},
		},
		func() interface{} {
			j := jsiiProxy_ZendeskConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConnectorProfileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ZendeskConnectorProfileProps",
		reflect.TypeOf((*ZendeskConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ZendeskInstanceUrlBuilder",
		reflect.TypeOf((*ZendeskInstanceUrlBuilder)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ZendeskInstanceUrlBuilder{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ZendeskOAuthSettings",
		reflect.TypeOf((*ZendeskOAuthSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-appflow.ZendeskSource",
		reflect.TypeOf((*ZendeskSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
		},
		func() interface{} {
			j := jsiiProxy_ZendeskSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-appflow.ZendeskSourceProps",
		reflect.TypeOf((*ZendeskSourceProps)(nil)).Elem(),
	)
}
