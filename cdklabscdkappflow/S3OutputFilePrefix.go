package cdklabscdkappflow


// Experimental.
type S3OutputFilePrefix struct {
	// Experimental.
	Format S3OutputFilePrefixFormat `field:"optional" json:"format" yaml:"format"`
	// Experimental.
	Hierarchy *[]S3OutputFilePrefixHierarchy `field:"optional" json:"hierarchy" yaml:"hierarchy"`
	// Experimental.
	Type S3OutputFilePrefixType `field:"optional" json:"type" yaml:"type"`
}

