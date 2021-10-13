package templates

import (
	"github.com/onflow/flow-nft/lib/go/templates/internal/assets"
)

const (
	transactionsPath     = "../../../transactions/"
	mintNFTFilename      = "mint_nft.cdc"
	setupAccountFilename = "setup_account.cdc"
	transferNFTFilename  = "transfer_nft.cdc"
)

// GenerateMintNFTScript generates a script to mint a new nft
func GenerateMintNFTScript(env Environment) []byte {
	code := assets.MustAssetString(transactionsPath + mintNFTFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateSetupAccountScript generates a script to setup account
func GenerateSetupAccountScript(env Environment) []byte {
	code := assets.MustAssetString(transactionsPath + setupAccountFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateTransferNFTScript generates a script to transfer nft
func GenerateTransferNFTScript(env Environment) []byte {
	code := assets.MustAssetString(transactionsPath + transactionsPath)

	return []byte(replaceAddresses(code, env))
}
