package cdklabscdkappflow


// Represents a list of Microsoft Sharepoint Online site drives from which to retrieve the documents.
// Experimental.
type MicrosoftSharepointOnlineObject struct {
	// The Microsoft Sharepoint Online site from which the documents are to be retrieved.
	//
	// Note: requires full name starting with 'sites/'.
	// Experimental.
	Site *string `field:"required" json:"site" yaml:"site"`
	// An array of Microsoft Sharepoint Online site drives from which the documents are to be retrieved.
	//
	// Note: each drive requires full name starting with 'drives/'.
	// Deprecated: . This property is deprecated and will be removed in a future release. Use {@link entities} instead
	Drives *[]*string `field:"optional" json:"drives" yaml:"drives"`
	// An array of Microsoft Sharepoint Online site entities from which the documents are to be retrieved.
	//
	// Note: each entity requires full name starting with 'drives/' followed by driveID and optional '/items/' followed by itemID.
	//
	// Example:
	//   "drives/${driveID}/items/${itemID}"
	//
	// Experimental.
	Entities *[]*string `field:"optional" json:"entities" yaml:"entities"`
}

