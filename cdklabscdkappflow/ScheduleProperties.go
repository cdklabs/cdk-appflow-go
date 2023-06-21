package cdklabscdkappflow

import (
	"time"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type ScheduleProperties struct {
	// Experimental.
	EndTime *time.Time `field:"optional" json:"endTime" yaml:"endTime"`
	// Timestamp for the records to import from the connector in the first flow run.
	// Experimental.
	FirstExecutionFrom *time.Time `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Experimental.
	Offset awscdk.Duration `field:"optional" json:"offset" yaml:"offset"`
	// Experimental.
	StartTime *time.Time `field:"optional" json:"startTime" yaml:"startTime"`
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

