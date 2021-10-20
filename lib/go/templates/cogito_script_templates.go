package templates

import (
	"github.com/wakandalabs/cogito-contracts/lib/go/templates/internal/assets"
)

const (
	scriptsPath                  = "../../../transactions/scripts/"
	readTotalSupplyFilename      = "read_totalSupply.cdc"
	readCollectionIdsFilename    = "read_collection_ids.cdc"
	readCollectionLengthFilename = "read_collection_length.cdc"
	readTokenURIFilename         = "read_tokenURI.cdc"
	isSetupCogitoFilename        = "is_setup_cogito.cdc"
)

// GenerateGetTotalSupplyScript generates a script to read collection ids
func GenerateGetTotalSupplyScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + readTotalSupplyFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateReadCollectionIdsScript generates a script to read collection ids
func GenerateReadCollectionIdsScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + readCollectionIdsFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateReadCollectionLengthScript generates a script to read collection length of nft
func GenerateReadCollectionLengthScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + readCollectionLengthFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateReadTokenURIScript generates a script to read nft tokenURI
func GenerateReadTokenURIScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + readTokenURIFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateIsSetupCogitoScript generates a script is setup Cogito
func GenerateIsSetupCogitoScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + isSetupCogitoFilename)

	return []byte(replaceAddresses(code, env))
}
