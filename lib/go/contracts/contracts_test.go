package contracts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-nft/lib/go/contracts"
)

const addrA = "0A"

func TestNonFungibleTokenContract(t *testing.T) {
	contract := contracts.GenerateNFTContract()
	assert.NotNil(t, contract)
}

func TestCogitoContract(t *testing.T) {
	contract := contracts.GenerateCogitoContract(addrA)
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA)
}
