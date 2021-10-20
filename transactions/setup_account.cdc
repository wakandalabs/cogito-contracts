import NonFungibleToken from 0xNFTADDRESS
import Cogito from 0xCOGITOADDRESS

// This transaction configures an account to hold and mint Cogito.

transaction {
    prepare(signer: AuthAccount) {
        // if the account doesn't already have a collection
        if signer.borrow<&Cogito.Collection>(from: Cogito.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- Cogito.createEmptyCollection()

            // save it to the account
            signer.save(<-collection, to: Cogito.CollectionStoragePath)

            // create a public capability for the collection
            signer.link<&Cogito.Collection{NonFungibleToken.CollectionPublic, Cogito.CogitoCollectionPublic}>(Cogito.CollectionPublicPath, target: Cogito.CollectionStoragePath)
        }
    }
}
