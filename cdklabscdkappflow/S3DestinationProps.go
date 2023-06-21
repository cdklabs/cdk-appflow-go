package cdklabscdkappflow


// Experimental.
type S3DestinationProps struct {
	// Experimental.
	Location *S3Location `field:"required" json:"location" yaml:"location"`
	// Experimental.
	Catalog *S3Catalog `field:"optional" json:"catalog" yaml:"catalog"`
	// Experimental.
	Formatting *S3OutputFormatting `field:"optional" json:"formatting" yaml:"formatting"`
}

