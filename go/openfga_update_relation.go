// this is used to delete tuple that has been added previously.

package main
import (
    fgaSdk "github.com/openfga/go-sdk"
	"context"
	"fmt"
)

func main() {
    // Initialize the SDK with no auth - see "How to setup SDK client" for more options
    configuration, err := fgaSdk.NewConfiguration(fgaSdk.Configuration{
        ApiScheme:      "http", //os.Getenv("FGA_SCHEME"), // Either "http" or "https", defaults to "https"
		ApiHost:        "0.0.0.0:8080", //os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        "01GX46JJ81KW8NY8K6TXAB3EGD", //os.Getenv("FGA_STORE_ID"), // optional, not needed for `CreateStore` and `ListStores`, required before calling for all other methods
    })

    if err != nil {
    // .. Handle error
    }

    fgaClient := fgaSdk.NewAPIClient(configuration)

	body := fgaSdk.WriteRequest{
		Deletes: &fgaSdk.TupleKeys{
			TupleKeys: []fgaSdk.TupleKey {
				{
					User: fgaSdk.PtrString("user:anne"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
			},
		},
		AuthorizationModelId: fgaSdk.PtrString("01GX472S23YCQ7MFF2PPA9PABB")} //01GX472S23YCQ7MFF2PPA9PABB

	_, response, err := fgaClient.OpenFgaApi.Write(context.Background()).Body(body).Execute()
	fmt.Println(response)
	if err != nil {
		// .. Handle error
	}
}
