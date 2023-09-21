package main
import (
	"context"
	"fmt"
    openfga "github.com/openfga/go-sdk"
//    "os"
)

func main() {
    configuration, err := openfga.NewConfiguration(openfga.Configuration{
		ApiScheme:      "http",
		ApiHost:        "0.0.0.0:8080",
    })

    if err != nil {
        // .. Handle error
		fmt.Println("testing raju")
    }
	fmt.Println("raju testing", configuration, err)

    apiClient := openfga.NewAPIClient(configuration)
	fmt.Println(apiClient, configuration)

    resp, _, err := apiClient.OpenFgaApi.CreateStore(context.Background()).Body(openfga.CreateStoreRequest{
        Name: "create_demo_store",
    }).Execute()
	fmt.Println("raju response")
	fmt.Println(resp, err)
    if err != nil {
        // .. Handle error
    }
}
