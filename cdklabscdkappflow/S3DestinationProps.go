package cdklabscdkappflow


// Experimental.
type S3DestinationProps struct {
	// The S3 location of the files with the retrieved data.
	// Experimental.
	Location *S3Location `field:"required" json:"location" yaml:"location"`
	// The AWS Glue cataloging options.
	// Experimental.
	Catalog *S3Catalog `field:"optional" json:"catalog" yaml:"catalog"`
	// The formatting options for the output files.
	// Experimental.
	Formatting *S3OutputFormatting `field:"optional" json:"formatting" yaml:"formatting"`
}

