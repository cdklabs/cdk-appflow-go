//go:build no_runtime_type_checking

package cdklabscdkappflow

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleAdsSource) validateBindParameters(scope IFlow) error {
	return nil
}

func validateNewGoogleAdsSourceParameters(props *GoogleAdsSourceProps) error {
	return nil
}

