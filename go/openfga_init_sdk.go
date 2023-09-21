package main
// ApiTokenIssuer, ApiAudience, ClientId and ClientSecret are optional.
import (
    fgaSdk "github.com/openfga/go-sdk"
    "os"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
    // Initialize the SDK with no auth - see "How to setup SDK client" for more options
    configuration, err := fgaSdk.NewConfiguration(fgaSdk.UserConfiguration{
        ApiScheme:      os.Getenv("FGA_SCHEME"), // Either "http" or "https", defaults to "https"
        ApiHost:        os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        os.Getenv("FGA_STORE_ID"), // optional, not needed for `CreateStore` and `ListStores`, required before calling for all other methods
    })

    if err != nil {
    // .. Handle error
    }

    fgaClient := fgaSdk.NewAPIClient(configuration)
}
