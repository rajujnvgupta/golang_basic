package main
import (
    fgaSdk "github.com/openfga/go-sdk"
	"fmt"
	"context"
)

func main() {
    // Initialize the SDK with no auth - see "How to setup SDK client" for more options

    configuration, err1 := fgaSdk.NewConfiguration(fgaSdk.Configuration{
        ApiScheme:      "http", //os.Getenv("FGA_SCHEME"), // Either "http" or "https", defaults to "https"
		ApiHost:        "0.0.0.0:8080", //os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        "01GZWV3948RQ8R2JGTGJYPRYBK", //os.Getenv("FGA_STORE_ID"), // optional, not needed for `CreateStore` and `ListStores`, required before calling for all other methods
    }) 
	if err1 != nil { // .. Handle error

    }

    fgaClient := fgaSdk.NewAPIClient(configuration)


	body := fgaSdk.CheckRequest{
		AuthorizationModelId: fgaSdk.PtrString("01GZWW3HCK93Q6QADY4EWK7N9R"),
		TupleKey: fgaSdk.TupleKey{
			User: fgaSdk.PtrString("user:ramesh"),
			Relation: fgaSdk.PtrString("can_edit"),
			Object: fgaSdk.PtrString("folder:note4"),
		},
		ContextualTuples: &fgaSdk.ContextualTupleKeys{
			TupleKeys: []fgaSdk.TupleKey{
				{
					User: fgaSdk.PtrString("user:ramesh"),
					Relation: fgaSdk.PtrString("user_in_context"),
					Object: fgaSdk.PtrString("user_groups:ug1"),
				},
			},
		},
	}
	data, response, err := fgaClient.OpenFgaApi.Check(context.Background()).Body(body).Execute()
	fmt.Println(data, response)

	if err != nil {
	// .. Handle error
	}
}

