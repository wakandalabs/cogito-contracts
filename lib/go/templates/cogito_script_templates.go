package templates

import (
	"github.com/wakandalabs/cogito-contracts/lib/go/templates/internal/assets"
)

const (
	scriptsPath                  = "../../../transactions/scripts/"
	totalSupplyFilename          = "get_totalSupply.cdc"
	readCollectionIdsFilename    = "read_collection_ids.cdc"
	readCollectionLengthFilename = "read_collection_length.cdc"
	readNFTByIdFilename          = "read_nft_by_id.cdc"
)

// GenerateGetTotalSupplyScript generates a script to read collection ids
func GenerateGetTotalSupplyScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + totalSupplyFilename)

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

// GenerateReadNFTByIdScript generates a script to read nft by id
func GenerateReadNFTByIdScript(env Environment) []byte {
	code := assets.MustAssetString(scriptsPath + readNFTByIdFilename)

	return []byte(replaceAddresses(code, env))
}
