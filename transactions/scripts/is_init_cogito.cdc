import NonFungibleToken from 0xNFTADDRESS

pub fun main(address: Address): Bool {
  let collection: Bool = getAccount(address)
      .getCapability<&{NonFungibleToken.CollectionPublic}>(/public/CogitoCollection)
      .check()

  return collection
}