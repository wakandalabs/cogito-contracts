package data

// CogitoMetadata Cadence requires a mapping of string->string, which can be handled through json tags when marshalling.
// It also does not allow for null values, so we will be omitting them if empty
// Reference: https://docs.google.com/spreadsheets/d/1muUZowii0pqoyi6OPK1VPNJi7keSc8_5zI9vu_QvfOY/edit#gid=375111836
type CogitoMetadata struct {
	original  string // Original cid of metadata, which may be encrypted
	temporary string // Storage decoding progress
}

// GenerateEmptyCogitoMetadata generates a play with all its fields
// empty except for FullName for testing
func GenerateEmptyCogitoMetadata(cid string) CogitoMetadata {
	return CogitoMetadata{
		original:  cid,
		temporary: cid,
	}
}
