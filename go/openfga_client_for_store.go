package main
import (
    openfga "github.com/openfga/go-sdk"
    "os"
)

func main() {
    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiScheme:      os.Getenv("FGA_API_SCHEME"), // Optional. Can be "http" or "https". Defaults to "https"
        ApiHost:        os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        os.Getenv("FGA_STORE_ID"),
    })

    if err != nil {
        // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)
}
