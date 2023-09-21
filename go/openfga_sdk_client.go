package main
import (
    openfga "github.com/openfga/go-sdk"
//    "os"
	"fmt"
)

func main() {
    configuration, err := openfga.NewConfiguration(openfga.Configuration{
		ApiScheme:      "http", //os.Getenv("FGA_API_SCHEME"), // Optional. Can be "http" or "https". Defaults to "https"
		ApiHost:        "0.0.0.0:8080", //os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        "01GX461BHKSQB62Z7WQN8XT7B5", //os.Getenv("FGA_STORE_ID"),
    })

    if err != nil {
        // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)
	fmt.Println(apiClient)
}
