import NonFungibleToken from 0xNFTADDRESS
import Cogito from 0xCOGITOADDRESS
// This transaction is what an account would run
// to set itself up to receive NFTs

transaction {

    prepare(acct: AuthAccount) {

        // Return early if the account already has a collection
        if acct.borrow<&Cogito.Collection>(from: /storage/CogitoCollection) != nil {
            return
        }

        // Create a new empty collection
        let collection <- Cogito.createEmptyCollection()

        // save it to the account
        acct.save(<-collection, to: /storage/CogitoCollection)

        // create a public capability for the collection
        acct.link<&{NonFungibleToken.CollectionPublic}>(
            /public/CogitoCollection,
            target: /storage/CogitoCollection
        )

        // Create a new minter
        let minter <- Cogito.createNewMinter()

        acct.save(<-minter, to: /storage/CogitoMinter)
    }
}
