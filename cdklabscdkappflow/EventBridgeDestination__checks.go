//go:build !no_runtime_type_checking

package cdklabscdkappflow

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (e *jsiiProxy_EventBridgeDestination) validateBindParameters(flow IFlow) error {
	if flow == nil {
		return fmt.Errorf("parameter flow is required, but nil was provided")
	}

	return nil
}

func validateNewEventBridgeDestinationParameters(props *EventBridgeDestinationProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

