import NonFungibleToken from 0x01

//
// NFT for Cogito ergo sum!
//
pub contract Cogito: NonFungibleToken {

    // Events
    //
    pub event ContractInitialized()
    pub event Withdraw(id: UInt64, from: Address?)
    pub event Deposit(id: UInt64, to: Address?)
    pub event Minted(id: UInt64, metadata: String)
    pub event Burned(id: UInt64)

    // Named Paths
    //
    pub let CollectionStoragePath: StoragePath
    pub let CollectionPublicPath: PublicPath

    // totalSupply
    // The total number of Cogito that have been minted
    //
    pub var totalSupply: UInt64

    // NFT
    // A Kitty Item as an NFT
    //
    pub resource NFT: NonFungibleToken.INFT {

        pub let id: UInt64

        pub let metadata: String

        // initializer
        //
        init(id: UInt64, metadata: String) {
            self.id = id
            self.metadata = metadata
        }
    }

    // This is the interface that users can cast their Cogito Collection as
    // to allow others to deposit Cogito into their Collection. It also allows for reading
    // the details of Cogito in the Collection.
    pub resource interface CogitoCollectionPublic {
        pub fun deposit(token: @NonFungibleToken.NFT)
        pub fun getIDs(): [UInt64]
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT
        pub fun borrowCogito(id: UInt64): &Cogito.NFT? {
            // If the result isn't nil, the id of the returned reference
            // should be the same as the argument to the function
            post {
                (result == nil) || (result?.id == id):
                    "Cannot borrow Cogito reference: The ID of the returned reference is incorrect"
            }
        }
        pub fun mintNFT(metadata: String)
        pub fun burnNFT(id: UInt64)
    }

    // Collection
    // A collection of Cogito NFTs owned by an account
    //
    pub resource Collection: CogitoCollectionPublic, NonFungibleToken.Provider, NonFungibleToken.Receiver, NonFungibleToken.CollectionPublic {
        // dictionary of NFT conforming tokens
        // NFT is a resource type with an `UInt64` ID field
        //
        pub var ownedNFTs: @{UInt64: NonFungibleToken.NFT}

        // withdraw
        // Removes an NFT from the collection and moves it to the caller
        //
        pub fun withdraw(withdrawID: UInt64): @NonFungibleToken.NFT {
            let token <- self.ownedNFTs.remove(key: withdrawID) ?? panic("missing NFT")

            emit Withdraw(id: token.id, from: self.owner?.address)

            return <-token
        }

        // deposit
        // Takes a NFT and adds it to the collections dictionary
        // and adds the ID to the id array
        //
        pub fun deposit(token: @NonFungibleToken.NFT) {
            let token <- token as! @Cogito.NFT

            let id: UInt64 = token.id

            // add the new token to the dictionary which removes the old one
            let oldToken <- self.ownedNFTs[id] <- token

            emit Deposit(id: id, to: self.owner?.address)

            destroy oldToken
        }

        // getIDs
        // Returns an array of the IDs that are in the collection
        //
        pub fun getIDs(): [UInt64] {
            return self.ownedNFTs.keys
        }

        // borrowNFT
        // Gets a reference to an NFT in the collection
        // so that the caller can read its metadata and call its methods
        //
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT {
            return &self.ownedNFTs[id] as &NonFungibleToken.NFT
        }

        // borrowCogito
        // Gets a reference to an NFT in the collection as a Cogito,
        // exposing all of its fields (including the typeID).
        // This is safe as there are no functions that can be called on the Cogito.
        //
        pub fun borrowCogito(id: UInt64): &Cogito.NFT? {
            if self.ownedNFTs[id] != nil {
                let ref = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT
                return ref as! &Cogito.NFT
            } else {
                return nil
            }
        }

        pub fun mintNFT(metadata: String) {
            emit Minted(id: Cogito.totalSupply, metadata: metadata)

            self.deposit(token: <-create Cogito.NFT(id: Cogito.totalSupply, metadata: metadata))

            Cogito.totalSupply = Cogito.totalSupply + (1 as UInt64)
        }

        pub fun burnNFT(id: UInt64) {
            emit Burned(id: id)

            let token <- self.withdraw(withdrawID: id)

            destroy token
        }

        // destructor
        destroy() {
            destroy self.ownedNFTs
        }

        // initializer
        //
        init () {
            self.ownedNFTs <- {}
        }
    }

    // createEmptyCollection
    // public function that anyone can call to create a new empty collection
    //
    pub fun createEmptyCollection(): @NonFungibleToken.Collection {
        return <- create Collection()
    }

    // fetch
    // Get a reference to a Cogito from an account's Collection, if available.
    // If an account does not have a Cogito.Collection, panic.
    // If it has a collection but does not contain the itemID, return nil.
    // If it has a collection and that collection contains the itemID, return a reference to that.
    //
    pub fun fetch(_ from: Address, itemID: UInt64): &Cogito.NFT? {
        let collection = getAccount(from)
            .getCapability(Cogito.CollectionPublicPath)
            .borrow<&Cogito.Collection{Cogito.CogitoCollectionPublic}>()
            ?? panic("Couldn't get collection")
        // We trust Cogito.Collection.borowCogito to get the correct itemID
        // (it checks it before returning it).
        return collection.borrowCogito(id: itemID)
    }

    // initializer
    //
	init() {
        // Set our named paths
        self.CollectionStoragePath = /storage/cogitoCollection
        self.CollectionPublicPath = /public/cogitoCollection

        // Initialize the total supply
        self.totalSupply = 0

        emit ContractInitialized()
	}
}
