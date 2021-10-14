module github.com/wakandalabs/cogito-contracts/libs/go/test

go 1.6

require (
	github.com/onflow/cadence v0.18.0
	github.com/onflow/flow-emulator v0.21.0
	github.com/onflow/flow-go-sdk v0.20.0
	github.com/stretchr/testify v1.7.0
	github.com/wakandalabs/cogito-contracts/lib/go/contracts v0.0.0-00010101000000-000000000000
	github.com/wakandalabs/cogito-contracts/lib/go/templates v0.0.0-00010101000000-000000000000
)

replace github.com/wakandalabs/cogito-contracts/lib/go/contracts => ../contracts

replace github.com/wakandalabs/cogito-contracts/lib/go/templates => ../templates
