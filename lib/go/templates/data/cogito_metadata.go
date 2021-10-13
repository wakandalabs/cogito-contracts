package data

// IpfsFile is stored on IPFS
type IpfsFile struct {
	FileName string `json:",omitempty"`
	Cid      string `json:",omitempty"`
}

type Property struct {
	Attachment       IpfsFile `json:",omitempty"`
	Copyright        string   `json:",omitempty"`
	CreatedTimestamp *int32   `json:",omitempty,string"`
}

type Properties []Property

// CogitoMetadata Cadence requires a mapping of string->string, which can be handled through json tags when marshalling.
// It also does not allow for null values, so we will be omitting them if empty
// Reference: https://docs.google.com/spreadsheets/d/1muUZowii0pqoyi6OPK1VPNJi7keSc8_5zI9vu_QvfOY/edit#gid=375111836
type CogitoMetadata struct {
	Name        string     `json:",omitempty"`
	Description string     `json:",omitempty"`
	Image       IpfsFile   `json:",omitempty"`
	Properties  Properties `json:",omitempty"`
}

// GenerateEmptyCogitoMetadata generates a play with all its fields
// empty except for FullName for testing
func GenerateEmptyCogitoMetadata(name string) CogitoMetadata {
	var properties Properties
	timeStamp := int32(123)
	property1 := Property{
		Attachment: IpfsFile{
			FileName: "attachment1",
			Cid:      "cid1",
		},
		Copyright:        "Test",
		CreatedTimestamp: &timeStamp,
	}
	properties = append(properties, property1)

	return CogitoMetadata{
		Name:        "Name of Cogito",
		Description: "Description of Cogito",
		Image: IpfsFile{
			FileName: "Image Cover",
			Cid:      "Cid",
		},
		Properties: properties,
	}
}
