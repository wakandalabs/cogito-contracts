package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts

import (
	"strings"

	_ "github.com/kevinburke/go-bindata"

	"github.com/onflow/flow-nft/lib/go/contracts/internal/assets"
)

const (
	nftFilename       = "NonFungibleToken.cdc"
	cogitoFilename    = "Cogito.cdc"
	defaultNFTAddress = "NonFungibleTokenAddr"
)

// GenerateNFTContract returns the NonFungibleToken contract interface.
func GenerateNFTContract() []byte {
	return assets.MustAsset(nftFilename)
}

// GenerateCogitoContract returns the Cogito contract.
//
// The returned contract will import the NonFungibleToken contract from the specified address.
func GenerateCogitoContract(nftAddress string) []byte {
	code := assets.MustAssetString(cogitoFilename)

	code = strings.ReplaceAll(
		code,
		defaultNFTAddress,
		nftAddress,
	)

	return []byte(code)
}
