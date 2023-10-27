package cdklabscdkappflow


// Represents a list of Microsoft Sharepoint Online site drives from which to retrieve the documents.
// Experimental.
type MicrosoftSharepointOnlineObject struct {
	// An array of Microsoft Sharepoint Online site drives from which the documents are to be retrieved.
	//
	// Note: each drive requires full name starting with 'drives/'.
	// Experimental.
	Drives *[]*string `field:"required" json:"drives" yaml:"drives"`
	// The Microsoft Sharepoint Online site from which the documents are to be retrieved.
	//
	// Note: requires full name starting with 'sites/'.
	// Experimental.
	Site *string `field:"required" json:"site" yaml:"site"`
}

