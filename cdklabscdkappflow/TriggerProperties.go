package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Experimental.
type TriggerProperties struct {
	// Experimental.
	DataPullConfig *DataPullConfig `field:"required" json:"dataPullConfig" yaml:"dataPullConfig"`
	// Experimental.
	Schedule awsevents.Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// Experimental.
	FlowErrorDeactivationThreshold *float64 `field:"optional" json:"flowErrorDeactivationThreshold" yaml:"flowErrorDeactivationThreshold"`
	// Experimental.
	Properties *ScheduleProperties `field:"optional" json:"properties" yaml:"properties"`
}

