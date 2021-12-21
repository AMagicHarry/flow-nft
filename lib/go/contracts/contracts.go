package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts

import (
	"regexp"

	_ "github.com/kevinburke/go-bindata"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-nft/lib/go/contracts/internal/assets"
)

var (
	placeholderNonFungibleToken = regexp.MustCompile(`"[^"\s].*/NonFungibleToken.cdc"`)
	placeholderViews            = regexp.MustCompile(`"[^"\s].*/Views.cdc"`)
)

const (
	filenameNonFungibleToken = "NonFungibleToken.cdc"
	filenameExampleNFT       = "ExampleNFT.cdc"
	filenameViews            = "Views.cdc"
)

// NonFungibleToken returns the NonFungibleToken contract interface.
func NonFungibleToken() []byte {
	return assets.MustAsset(filenameNonFungibleToken)
}

// ExampleNFT returns the ExampleNFT contract.
//
// The returned contract will import the NonFungibleToken contract from the specified address.
func ExampleNFT(nftAddress, viewsAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameExampleNFT)

	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())
	code = placeholderViews.ReplaceAllString(code, "0x"+viewsAddress.String())

	return []byte(code)
}

func Views() []byte {
	return assets.MustAsset(filenameViews)
}
