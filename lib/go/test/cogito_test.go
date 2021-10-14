package test

import (
	"testing"

	//"github.com/onflow/cadence"
	//jsoncdc "github.com/onflow/cadence/encoding/json"
	//"github.com/onflow/flow-go-sdk/crypto"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	//"github.com/onflow/flow-go-sdk/test"

	"github.com/stretchr/testify/assert"

	//"github.com/onflow/flow-go-sdk"
	"github.com/wakandalabs/cogito-contracts/lib/go/contracts"
	//"github.com/wakandalabs/cogito-contracts/lib/go/templates"
)

const (
	emulatorFTAddress        = "ee82856bf20e2aa6"
	emulatorFlowTokenAddress = "0ae53cb6e3f42a79"
)

// This test is for testing the deployment the cogito smart contracts
func TestCogitoDeployment(t *testing.T) {
	b := newBlockchain()

	// Should be able to deploy the NFT contract
	// as a new account with no keys.
	nftCode := contracts.GenerateNFTContract()
	nftAddr, err := b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "NonFungibleToken",
			Source: string(nftCode),
		},
	})
	if !assert.NoError(t, err) {
		t.Log(err.Error())
	}
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Should be able to deploy the cogito contract
	// as a new account with no keys.
	cogitoCode := contracts.GenerateCogitoContract(nftAddr.String())
	_, err = b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "Cogito",
			Source: string(cogitoCode),
		},
	})
	if !assert.NoError(t, err) {
		t.Log(err.Error())
	}
	_, err = b.CommitBlock()
	assert.NoError(t, err)
}
