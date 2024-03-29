package contracts

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts

import (
	"regexp"

	_ "github.com/kevinburke/go-bindata"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-nft/lib/go/contracts/internal/assets"
)

var (
	placeholderNonFungibleToken = regexp.MustCompile(`"NonFungibleToken"`)
	placeholderMetadataViews    = regexp.MustCompile(`"MetadataViews"`)
	placeholderFungibleToken    = regexp.MustCompile(`"FungibleToken"`)
	placeholderResolverToken    = regexp.MustCompile(`"ViewResolver"`)
)

const (
	filenameNonFungibleToken    = "NonFungibleToken.cdc"
	filenameOldNonFungibleToken = "utility/NonFungibleToken_old.cdc"
	filenameExampleNFT          = "ExampleNFT.cdc"
	filenameMetadataViews       = "MetadataViews.cdc"
	filenameResolver            = "ViewResolver.cdc"
	filenameFungibleToken       = "utility/FungibleToken.cdc"
)

// NonFungibleToken returns the NonFungibleToken contract interface.
func NonFungibleToken() []byte {
	return assets.MustAsset(filenameNonFungibleToken)
}

// OldNonFungibleToken returns the old NonFungibleToken contract interface
// without default implementations
func OldNonFungibleToken() []byte {
	return assets.MustAsset(filenameOldNonFungibleToken)
}

// ExampleNFT returns the ExampleNFT contract.
//
// The returned contract will import the NonFungibleToken contract from the specified address.
func ExampleNFT(nftAddress, metadataAddress, resolverAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameExampleNFT)

	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())
	code = placeholderMetadataViews.ReplaceAllString(code, "0x"+metadataAddress.String())
	code = placeholderResolverToken.ReplaceAllString(code, "0x"+resolverAddress.String())

	return []byte(code)
}

func MetadataViews(ftAddress flow.Address, nftAddress flow.Address) []byte {
	code := assets.MustAssetString(filenameMetadataViews)

	code = placeholderFungibleToken.ReplaceAllString(code, "0x"+ftAddress.String())
	code = placeholderNonFungibleToken.ReplaceAllString(code, "0x"+nftAddress.String())

	return []byte(code)
}

func Resolver() []byte {
	code := assets.MustAssetString(filenameResolver)
	return []byte(code)
}

// FungibleToken returns the FungibleToken contract interface.
func FungibleToken() []byte {
	return assets.MustAsset(filenameFungibleToken)
}
