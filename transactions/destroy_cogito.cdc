import NonFungibleToken from 0xNFTADDRESS
import Cogito from 0xCOGITOADDRESS

// This script uses the Cogito Collection resource to destroy a Cogito NFT
// It must be run with the account that has the collection resource
// stored in /storage/CogitoCollection

transaction(id: UInt64) {

    // local variable for storing the minter reference
    let collection: &Cogito.Collection

    prepare(signer: AuthAccount) {

        // borrow a reference to the NFTMinter resource in storage
        self.collection = signer.borrow<&Cogito.Collection>(from: /storage/CogitoCollection)
            ?? panic("Could not borrow a reference to the NFT collection")
    }

    execute {
        // Destroy the NFT
        self.collection.destroyNFT(id: id)
    }
}
