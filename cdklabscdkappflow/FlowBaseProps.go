package cdklabscdkappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Experimental.
type FlowBaseProps struct {
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
	// Experimental.
	Type FlowType `field:"required" json:"type" yaml:"type"`
	// Experimental.
	Status FlowStatus `field:"optional" json:"status" yaml:"status"`
	// Experimental.
	TriggerConfig *TriggerConfig `field:"optional" json:"triggerConfig" yaml:"triggerConfig"`
}

