package cdklabscdkappflow


// Experimental.
type ErrorHandlingConfiguration struct {
	// Experimental.
	ErrorLocation *S3Location `field:"optional" json:"errorLocation" yaml:"errorLocation"`
	// Experimental.
	FailOnFirstError *bool `field:"optional" json:"failOnFirstError" yaml:"failOnFirstError"`
}

