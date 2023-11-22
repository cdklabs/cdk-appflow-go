package cdklabscdkappflow


// Experimental.
type RedshiftDestinationObject struct {
	// Experimental.
	Table interface{} `field:"required" json:"table" yaml:"table"`
	// Default: public.
	//
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

