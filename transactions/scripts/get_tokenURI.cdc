import NonFungibleToken from 0x01
import Cogito from 0x02

pub fun main(address: Address, id: UInt64): String? {
  if let collection = getAccount(address).getCapability<&Cogito.Collection{Cogito.CollectionPublic}>(/public/CogitoCollection).borrow() {
    if let item = collection.borrowCogito(id: id) {
      return item.metadata
    }
  }

  return nil
}
