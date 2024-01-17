package cdklabscdkappflow


// Experimental.
type GoogleBigQueryObject struct {
	// Experimental.
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Experimental.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Experimental.
	Table *string `field:"required" json:"table" yaml:"table"`
}

