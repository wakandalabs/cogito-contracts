import NonFungibleToken from 0xNFTADDRESS
import Cogito from 0xCOGITOADDRESS

// This transaction uses the Collection resource to mint a new NFT.
//
// It must be run with the account that has the minter resource
// stored at path Cogito.CollectionStoragePath

transaction(metadata: String) {

    // local variable for storing the minter reference
    let minter: &Cogito.Collection

    prepare(signer: AuthAccount) {

        // borrow a reference to the NFTMinter resource in storage
        self.minter = signer.borrow<&Cogito.Collection>(from: Cogito.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }

    execute {
        // mint the NFT and deposit it to minter's collection
        self.minter.mintNFT(metadata: metadata)
    }
}
