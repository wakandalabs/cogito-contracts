package test

import (
	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/wakandalabs/cogito-contracts/lib/go/templates"
	"testing"

	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/stretchr/testify/assert"
	"github.com/wakandalabs/cogito-contracts/lib/go/contracts"
)

const (
	emulatorFTAddress        = "ee82856bf20e2aa6"
	emulatorFlowTokenAddress = "0ae53cb6e3f42a79"
)

// TestCogitoDeployment tests the deployment the cogito smart contracts
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

// TestMintCogito tests the pure functionality of the smart contract
func TestMintCogito(t *testing.T) {
	b := newBlockchain()

	accountKeys := test.AccountKeyGenerator()

	env := templates.Environment{
		FungibleTokenAddress: emulatorFTAddress,
		FlowTokenAddress:     emulatorFlowTokenAddress,
	}

	// Should be able to deploy the NFT contract
	// as a new account with no keys.
	nftCode := contracts.GenerateNFTContract()
	nftAddr, _ := b.CreateAccount(nil, []sdktemplates.Contract{
		{
			Name:   "NonFungibleToken",
			Source: string(nftCode),
		},
	})

	env.NFTAddress = nftAddr.String()

	// Should be able to deploy the cogito contract
	// as a new account with no keys.
	cogitoCode := contracts.GenerateCogitoContract(nftAddr.String())
	cogitoAccountKey, cogitoSigner := accountKeys.NewWithSigner()
	cogitoAddr, _ := b.CreateAccount([]*flow.AccountKey{cogitoAccountKey}, []sdktemplates.Contract{
		{
			Name:   "Cogito",
			Source: string(cogitoCode),
		},
	})

	_, _ = b.CommitBlock()

	env.CogitoAddress = cogitoAddr.String()

	// Create a new user account
	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, _ := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)

	// Check that that main contract fields were initialized correctly
	result := executeScriptAndCheck(t, b, templates.GenerateGetTotalSupplyScript(env), nil)
	assert.Equal(t, cadence.NewUInt64(0), result)

	// create a new Collection for Admin
	t.Run("Should be able to setup account", func(t *testing.T) {
		tx := createTxWithTemplateAndAuthorizer(b, templates.GenerateSetupAccountScript(env), cogitoAddr)

		result = executeScriptAndCheck(t, b, templates.GenerateIsSetupCogitoScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(cogitoAddr))})
		assert.Equal(t, cadence.NewBool(false), result)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, cogitoAddr}, []crypto.Signer{b.ServiceKey().Signer(), cogitoSigner},
			false,
		)

		result = executeScriptAndCheck(t, b, templates.GenerateIsSetupCogitoScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(cogitoAddr))})
		assert.Equal(t, cadence.NewBool(true), result)
	})

	// Admin sends a transaction that mints a cogito
	t.Run("Should be able to mint a cogito", func(t *testing.T) {
		tx := createTxWithTemplateAndAuthorizer(b, templates.GenerateMintNFTScript(env), cogitoAddr)

		_ = tx.AddArgument(cadence.NewString("metadata"))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, cogitoAddr}, []crypto.Signer{b.ServiceKey().Signer(), cogitoSigner},
			false,
		)

		//make sure the cogito was minted correctly and is stored in the collection with the correct data
		result = executeScriptAndCheck(t, b, templates.GenerateReadCollectionIdsScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(cogitoAddr))})
		assert.Equal(t, cadence.NewArray([]cadence.Value{cadence.NewUInt64(0)}), result)
	})

	// create a new Collection for a user address
	t.Run("Should be able to create a cogito collection", func(t *testing.T) {
		tx := createTxWithTemplateAndAuthorizer(b, templates.GenerateSetupAccountScript(env), joshAddress)
		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress}, []crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)
	})

	// Admin sends a transaction to transfer a cogito to a user
	t.Run("Should be able to transfer a cogito", func(t *testing.T) {
		tx := createTxWithTemplateAndAuthorizer(b, templates.GenerateTransferNFTScript(env), cogitoAddr)

		_ = tx.AddArgument(cadence.NewAddress(joshAddress))
		_ = tx.AddArgument(cadence.NewUInt64(0))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, cogitoAddr}, []crypto.Signer{b.ServiceKey().Signer(), cogitoSigner},
			false,
		)

		// make sure the user received it
		result = executeScriptAndCheck(t, b, templates.GenerateReadCollectionIdsScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		assert.Equal(t, cadence.NewArray([]cadence.Value{cadence.NewUInt64(0)}), result)
	})

	t.Run("Should be able to burn a cogito", func(t *testing.T) {
		tx := createTxWithTemplateAndAuthorizer(b, templates.GenerateBurnNFTScript(env), joshAddress)

		_ = tx.AddArgument(cadence.NewUInt64(0))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{b.ServiceKey().Address, joshAddress}, []crypto.Signer{b.ServiceKey().Signer(), joshSigner},
			false,
		)

		result = executeScriptAndCheck(t, b, templates.GenerateReadCollectionLengthScript(env), [][]byte{jsoncdc.MustEncode(cadence.Address(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), result)
	})
}
