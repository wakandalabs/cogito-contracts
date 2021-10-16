package templates

import "github.com/wakandalabs/cogito-contracts/lib/go/templates/internal/assets"

const (
	transactionsPath       = "../../../transactions/"
	mintCogitoFilename     = "mint_cogito.cdc"
	setupAccountFilename   = "setup_account.cdc"
	transferCogitoFilename = "transfer_cogito.cdc"
)

// GenerateMintNFTScript generates a script to mint a new nft
func GenerateMintNFTScript(env Environment) []byte {
	code := assets.MustAssetString(transactionsPath + mintCogitoFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateSetupAccountScript generates a script to setup account
func GenerateSetupAccountScript(env Environment) []byte {
	code := assets.MustAssetString(transactionsPath + setupAccountFilename)

	return []byte(replaceAddresses(code, env))
}

// GenerateTransferNFTScript generates a script to transfer nft
func GenerateTransferNFTScript(env Environment) []byte {
	code := assets.MustAssetString(transactionsPath + transferCogitoFilename)

	return []byte(replaceAddresses(code, env))
}
