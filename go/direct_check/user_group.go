// ApiTokenIssuer, ApiAudience, ClientId and ClientSecret are optional.
package main

import (
    fgaSdk "github.com/openfga/go-sdk"
//    "os"
	"fmt"
	"context"
	"encoding/json"
)

func main() {
    // Initialize the SDK with no auth - see "How to setup SDK client" for more options
    configuration, err := fgaSdk.NewConfiguration(fgaSdk.Configuration{
        ApiScheme:      "http", //os.Getenv("FGA_SCHEME"), // Either "http" or "https", defaults to "https"
		ApiHost:        "0.0.0.0:8080", //os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        "01GX461BHKSQB62Z7WQN8XT7B5", //os.Getenv("FGA_STORE_ID"), // optional, not needed for `CreateStore` and `ListStores`, required before calling for all other methods
    })

    if err != nil {
    // .. Handle error
    }
// write mode into store
    fgaClient := fgaSdk.NewAPIClient(configuration)
	WriteAuthorizationModelRequestString := "{\"schema_version\":\"1.1\",\"type_definitions\":[{\"type\":\"user\"},{\"type":"document\",\"relations\":{\"editor\":{\"this\":{}}},\"metadata\":{\"relations\":{\"editor\":{\"directly_related_user_types\":[{\"type\":\"team\",\"relation\":\"member\"}]}}}},{\"type\":\"team\",\"relations\":{\"member\":{\"this\":{}}},\"metadata\":{\"relations\":{\"member\":{\"directly_related_user_types\":[{\"type\":\"user\"}]}}}}]}"

	var body fgaSdk.WriteAuthorizationModelRequest
//	var body map[string]interface{}
	if err := json.Unmarshal([]byte(WriteAuthorizationModelRequestString), &body); err != nil {
      fmt.Printf("%s", err)
      // .. Handle error
      return
	}

	data, response, err := fgaClient.OpenFgaApi.WriteAuthorizationModel(context.Background()).Body(body).Execute()
	fmt.Println(data, response, err)
	if err != nil {
      // .. Handle error
	}
}
