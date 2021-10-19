/*
    Description: Central Smart Contract for Cogito Ergo Sum
*/

import NonFungibleToken from 0xNFTADDRESS

pub contract Cogito: NonFungibleToken {

    // -----------------------------------------------------------------------
    // Cogito contract Events
    // -----------------------------------------------------------------------

    // Emitted when the Cogito contract is created
    pub event ContractInitialized()
    // Emitted when a cogito is withdrawn from a Collection
    pub event Withdraw(id: UInt64, from: Address?)
    // Emitted when a cogito is deposited into a Collection
    pub event Deposit(id: UInt64, to: Address?)
    // Emitted when a cogito is minted into a Collection
    pub event Minted(id: UInt64, metadata: String)
    // Emitted when a Cogito is destroyed
    pub event Destroyed(id: UInt64)

    // The total number of Cogito NFTs that have been created
    // Because NFTs can be destroyed, it doesn't necessarily mean that this
    // reflects the total number of NFTs in existence, just the number that
    // have been minted to date. Also used as global cogito IDs for minting.
    pub var totalSupply: UInt64

    pub resource interface CogitoPublic {
        pub fun getTokenURI(): String
    }

    pub resource NFT: NonFungibleToken.INFT, CogitoPublic {
        pub let id: UInt64

        pub var metadata: String

        init(initID: UInt64, metadata: String) {
            self.id = initID
            self.metadata = metadata

            emit Minted(id: self.id, metadata: self.metadata)
        }

        // If the Cogito is destroyed, emit an event to indicate
        // to outside observers that it has been destroyed
        destroy() {
            emit Destroyed(id: self.id)
        }

        pub fun getTokenURI():String {
            return self.metadata
        }
    }

    pub resource interface CollectionMinter {
        pub fun mintNFT(metadata: String)
    }

    pub resource interface CollectionDestroy {
        pub fun destroyNFT(id: UInt64)
    }

    pub resource Collection: NonFungibleToken.Provider, NonFungibleToken.Receiver, NonFungibleToken.CollectionPublic, CollectionMinter, CollectionDestroy {
        // dictionary of NFT conforming tokens
        // NFT is a resource type with an `UInt64` ID field
        pub var ownedNFTs: @{UInt64: NonFungibleToken.NFT}

        init () {
            self.ownedNFTs <- {}
        }

        // withdraw removes an NFT from the collection and moves it to the caller
        pub fun withdraw(withdrawID: UInt64): @NonFungibleToken.NFT {
            let token <- self.ownedNFTs.remove(key: withdrawID) ?? panic("missing NFT")

            emit Withdraw(id: token.id, from: self.owner?.address)

            return <-token
        }

        // deposit takes a NFT and adds it to the collections dictionary
        // and adds the ID to the id array
        pub fun deposit(token: @NonFungibleToken.NFT) {
            let token <- token as! @Cogito.NFT

            let id: UInt64 = token.id

            // add the new token to the dictionary which removes the old one
            let oldToken <- self.ownedNFTs[id] <- token

            emit Deposit(id: id, to: self.owner?.address)

            destroy oldToken
        }

        // getIDs returns an array of the IDs that are in the collection
        pub fun getIDs(): [UInt64] {
            return self.ownedNFTs.keys
        }

        // borrowNFT gets a reference to an NFT in the collection
        // so that the caller can read its metadata and call its methods
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT {
            return &self.ownedNFTs[id] as &NonFungibleToken.NFT
        }

        destroy() {
            destroy self.ownedNFTs
        }

            // mintNFT can mint cogito for self
        pub fun mintNFT(metadata: String) {

            // create a new NFT
            var newNFT <- create NFT(initID: Cogito.totalSupply, metadata: metadata)

            // deposit it in the recipient's account using their reference
            self.deposit(token: <-newNFT)

            Cogito.totalSupply = Cogito.totalSupply + (1 as UInt64)
        }

          pub fun destroyNFT(id: UInt64) {
            let token <- self.withdraw(withdrawID: id)

            destroy token
        }
    }

    // public function that anyone can call to create a new empty collection
    pub fun createEmptyCollection(): @NonFungibleToken.Collection {
        return <- create Collection()
    }

    init() {
        // Initialize the total supply
        self.totalSupply = 0

        // Create a Collection resource and save it to storage
        let collection <- create Collection()
        self.account.save(<-collection, to: /storage/CogitoCollection)

        // create a public capability for the collection
        self.account.link<&{NonFungibleToken.CollectionPublic}>(
            /public/CogitoCollection,
            target: /storage/CogitoCollection
        )

        emit ContractInitialized()
    }
}
