import NonFungibleToken from 0xNFTADDRESS
import Cogito from 0xCOGITOADDRESS

pub fun main(address: Address): Bool {
  let collection: Bool = getAccount(address)
      .getCapability<&{NonFungibleToken.CollectionPublic}>(Cogito.CollectionPublicPath)
      .check()

  return collection
}