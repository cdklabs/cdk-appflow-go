package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Experimental.
type OnScheduleFlowProps struct {
	// Experimental.
	Destination IDestination `field:"required" json:"destination" yaml:"destination"`
	// Experimental.
	Mappings *[]IMapping `field:"required" json:"mappings" yaml:"mappings"`
	// Experimental.
	Source ISource `field:"required" json:"source" yaml:"source"`
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Experimental.
	Filters *[]IFilter `field:"optional" json:"filters" yaml:"filters"`
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	Transforms *[]ITransform `field:"optional" json:"transforms" yaml:"transforms"`
	// Experimental.
	Validations *[]IValidation `field:"optional" json:"validations" yaml:"validations"`
	// Deprecated: . This property is deprecated and will be removed in a future release. Use {@link status } instead
	AutoActivate *bool `field:"optional" json:"autoActivate" yaml:"autoActivate"`
	// The status to set on the flow.
	//
	// Use this over {@link autoActivate}.
	// Experimental.
	Status FlowStatus `field:"optional" json:"status" yaml:"status"`
	// Experimental.
	PullConfig *DataPullConfig `field:"required" json:"pullConfig" yaml:"pullConfig"`
	// Experimental.
	Schedule awsevents.Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// Experimental.
	ScheduleProperties *ScheduleProperties `field:"optional" json:"scheduleProperties" yaml:"scheduleProperties"`
}

