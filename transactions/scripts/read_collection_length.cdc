import NonFungibleToken from 0xNFTADDRESS
import Cogito from 0xCOGITOADDRESS

// This transaction gets the length of an account's nft collection

pub fun main(address: Address): Int {
    let account = getAccount(address)

    let collectionRef = account.getCapability(Cogito.CollectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")

    return collectionRef.getIDs().length
}
