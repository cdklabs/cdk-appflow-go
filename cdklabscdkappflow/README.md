# Amazon AppFlow Construct Library

*Note:* this library is currently in technical preview.

## Introduction

Amazon AppFlow is a service that enables creating managed, bi-directional data transfer integrations between various SaaS applications and AWS services.

For more information, see the [Amazon AppFlow User Guide](https://docs.aws.amazon.com/appflow/latest/userguide/what-is-appflow.html).

## Example

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow"

var clientSecret iSecret
var accessToken string
var refreshToken string
var instanceUrl string


profile := cdklabscdkappflow.NewSalesforceConnectorProfile(this, jsii.String("MyConnectorProfile"), &SalesforceConnectorProfileProps{
	OAuth: &SalesforceOAuthSettings{
		AccessToken: accessToken,
		Flow: &SalesforceOAuthFlow{
			RefreshTokenGrant: &SalesforceOAuthRefreshTokenGrantFlow{
				RefreshToken: refreshToken,
				Client: clientSecret,
			},
		},
	},
	InstanceUrl: instanceUrl,
	IsSandbox: jsii.Boolean(false),
})

source := cdklabscdkappflow.NewSalesforceSource(&SalesforceSourceProps{
	Profile: profile,
	Object: jsii.String("Account"),
})

bucket := awscdk.NewBucket(this, jsii.String("DestinationBucket"))

destination := cdklabscdkappflow.NewS3Destination(&S3DestinationProps{
	Location: &S3Location{
		Bucket: *Bucket,
	},
})

cdklabscdkappflow.NewOnDemandFlow(this, jsii.String("SfAccountToS3"), &OnDemandFlowProps{
	Source: source,
	Destination: destination,
	Mappings: []iMapping{
		*cdklabscdkappflow.Mapping_MapAll(),
	},
	Transforms: []iTransform{
		*cdklabscdkappflow.Transform_Mask(&Field{
			Name: jsii.String("Name"),
		}, jsii.String("*")),
	},
	Validations: []iValidation{
		*cdklabscdkappflow.Validation_When(*cdklabscdkappflow.ValidationCondition_IsNull(jsii.String("Name")), *cdklabscdkappflow.ValidationAction_IgnoreRecord()),
	},
	Filters: []iFilter{
		*cdklabscdkappflow.Filter_When(*cdklabscdkappflow.FilterCondition_TimestampLessThanEquals(&Field{
			Name: jsii.String("LastModifiedDate"),
			DataType: jsii.String("datetime"),
		}, NewDate(date.parse(jsii.String("2022-02-02"))))),
	},
})
```

# Concepts

Amazon AppFlow introduces several concepts that abstract away the technicalities of setting up and managing data integrations.

An `Application` is any SaaS data integration component that can be either a *source* or a *destination* for Amazon AppFlow. A source is an application from which Amazon AppFlow will retrieve data, whereas a destination is an application to which Amazon AppFlow will send data.

A `Flow` is Amazon AppFlow's integration between a source and a destination.

A `ConnectorProfile` is Amazon AppFlow's abstraction over authentication/authorization with a particular SaaS application. The per-SaaS application permissions given to a particular `ConnectorProfile` will determine whether the connector profile can support the application as a source or as a destination (see whether a particular application is supported as either a source or a destination in [the documentation](https://docs.aws.amazon.com/appflow/latest/userguide/app-specific.html)).

## Types of Flows

The library introduces three, separate types of flows:

* `OnDemandFlow` - a construct representing a flow that can be triggered programmatically with the use of a [StartFlow API call](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_StartFlow.html).
* `OnEventFlow` - a construct representing a flow that is triggered by a SaaS application event published to AppFlow. At the time of writing only a Salesforce source is able to publish events that can be consumed by AppFlow flows.
* `OnScheduleFlow` - a construct representing a flow that is triggered on a [`Schedule`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_events.Schedule.html)

## Tasks

Tasks are steps that can be taken upon fields. Tasks compose higher level objects that in this library are named `Operations`. There are four operations identified:

* Transforms - 1-1 transforms on source fields, like truncation or masking
* Mappings - 1-1 or many-to-1 operations from source fields to a destination field
* Filters - operations that limit the source data on a particular conditions
* Validations - operations that work on a per-record level and can have either a record-level consequence (i.e. dropping the record) or a global one (terminating the flow).

Each flow exposes dedicated properties to each of the operation types that one can use like in the example below:

```go
import "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow"

var stack stack
var source iSource
var destination iDestination


flow := cdklabscdkappflow.NewOnDemandFlow(stack, jsii.String("OnDemandFlow"), &OnDemandFlowProps{
	Source: source,
	Destination: destination,
	Transforms: []iTransform{
		*cdklabscdkappflow.Transform_Mask(&Field{
			Name: jsii.String("Name"),
		}, jsii.String("*")),
	},
	Mappings: []iMapping{
		*cdklabscdkappflow.Mapping_Map(&Field{
			Name: jsii.String("Name"),
			DataType: jsii.String("String"),
		}, &Field{
			Name: jsii.String("Name"),
			DataType: jsii.String("string"),
		}),
	},
	Filters: []iFilter{
		*cdklabscdkappflow.Filter_When(*cdklabscdkappflow.FilterCondition_TimestampLessThanEquals(&Field{
			Name: jsii.String("LastModifiedDate"),
			DataType: jsii.String("datetime"),
		}, NewDate(date.parse(jsii.String("2022-02-02"))))),
	},
	Validations: []iValidation{
		*cdklabscdkappflow.Validation_When(*cdklabscdkappflow.ValidationCondition_IsNull(jsii.String("Name")), *cdklabscdkappflow.ValidationAction_IgnoreRecord()),
	},
})
```

## EventBridge notifications

Each flow publishes events to the default EventBridge bus:

* `onRunStarted`
* `onRunCompleted`
* `onDeactivated` (only for the `OnEventFlow` and the `OnScheduleFlow`)
* `onStatus` (only for the `OnEventFlow` )

This way one can consume the notifications as in the example below:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow"

var flow iFlow
var myTopic iTopic


flow.OnRunCompleted(jsii.String("OnRunCompleted"), &OnEventOptions{
	Target: awscdk.NewSnsTopic(myTopic),
})
```

# Notable distinctions from CloudFormation specification

## `OnScheduleFlow` and `incrementalPullConfig`

In CloudFormation the definition of the `incrementalPullConfig` (which effectively gives a name of the field used for tracking the last pulled timestamp) is on the [`SourceFlowConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-incrementalpullconfig) property. In the library this has been moved to the `OnScheduleFlow` constructor properties.

## `S3Destination` and Glue Catalog

Although in CloudFormation the Glue Catalog configuration is settable on the flow level - it works only when the destination is S3. That is why the library shifts the Glue Catalog properties definition to the `S3Destination`, which in turn requires using Lazy for populating `metadataCatalogConfig` in the flow.

# Security considerations

It is *recommended* to follow [data protection mechanisms for Amazon AppFlow](https://docs.aws.amazon.com/appflow/latest/userguide/data-protection.html).

## Confidential information

Amazon AppFlow application integration is done using `ConnectionProfiles`. A `ConnectionProfile` requires providing sensitive information in the form of e.g. access and refresh tokens. It is *recommended* that such information is stored securely and passed to AWS CDK securely. All sensitive fields are effectively `IResolvable` and this means they can be resolved at deploy time. With that one should follow the [best practices for credentials with CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/security-best-practices.html#creds).

An example of using a predefined AWS Secrets Manager secret for storing sensitive information can be found below:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-appflow-go/cdklabscdkappflow"

var stack stack


secret := awscdk.Secret_FromSecretNameV2(stack, jsii.String("GA4Secret"), jsii.String("appflow/ga4"))

profile := cdklabscdkappflow.NewGoogleAnalytics4ConnectorProfile(stack, jsii.String("GA4Connector"), &GoogleAnalytics4ConnectorProfileProps{
	OAuth: &GoogleAnalytics4OAuthSettings{
		Flow: &GoogleAnalytics4OAuthFlow{
			RefreshTokenGrant: &GoogleAnalytics4RefreshTokenGrantFlow{
				RefreshToken: secret.SecretValueFromJson(jsii.String("refreshToken")).ToString(),
				ClientId: secret.*SecretValueFromJson(jsii.String("clientId")).*ToString(),
				ClientSecret: secret.*SecretValueFromJson(jsii.String("clientSecret")).*ToString(),
			},
		},
	},
})
```

## An approach to managing permissions

This library relies on an internal `AppFlowPermissionsManager` class to automatically infer and apply appropriate resource policy statements to the S3 Bucket, KMS Key, and Secrets Manager Secret resources. `AppFlowPermissionsManager` places the statements exactly once for the `appflow.amazonaws.com` principal no matter how many times a resource is reused in the code.

### Confused Deputy Problem

Amazon AppFlow is an account-bound and a regional service. With this it is invurlnerable to the confused deputy problem (see, e.g. [here](https://docs.aws.amazon.com/IAM/latest/UserGuide/confused-deputy.html)). However, `AppFlowPermissionsManager` still introduces the `aws:SourceAccount` condtition to the resource policies as a *best practice*.
