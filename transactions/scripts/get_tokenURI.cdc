import NonFungibleToken from 0xNFTADDRESS
import Cogito from 0xCOGITOADDRESS

pub fun main(address: Address, id: UInt64): String? {
  if let collection = getAccount(address).getCapability<&Cogito.Collection{NonFungibleToken.CollectionPublic, Cogito.CollectionPublic}>(/public/CogitoCollection).borrow() {
    if let item = collection.borrowCogito(id: id) {
      return item.metadata
    }
  }

  return nil
}
