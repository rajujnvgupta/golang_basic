package main
import (
    fgaSdk "github.com/openfga/go-sdk"
	"fmt"
	"context"
)

func main() {
    // Initialize the SDK with no auth - see "How to setup SDK client" for more options

    configuration, err := fgaSdk.NewConfiguration(fgaSdk.Configuration{
        ApiScheme:      "http", //os.Getenv("FGA_SCHEME"), // Either "http" or "https", defaults to "https"
		ApiHost:        "0.0.0.0:8080", //os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        "", //os.Getenv("FGA_STORE_ID"), // optional, not needed for `CreateStore` and `ListStores`, required before calling for all other methods
    })

    if err != nil {
    // .. Handle error
    }

    fgaClient := fgaSdk.NewAPIClient(configuration)

	body := fgaSdk.WriteRequest{
		Writes:  &fgaSdk.TupleKeys{
			TupleKeys: []fgaSdk.TupleKey {
				{
					User: fgaSdk.PtrString("user:anne1"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne2"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne3"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne4"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne5"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne6"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne7"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne8"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne9"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
				{
					User: fgaSdk.PtrString("user:anne10"),
					Relation: fgaSdk.PtrString("reader"),
					Object: fgaSdk.PtrString("document:Z"),
				},
			},
		},
		AuthorizationModelId: fgaSdk.PtrString("01GX6QBJHX2PDW2MDB6E6F2EHE")}

	_, response, err := fgaClient.OpenFgaApi.Write(context.Background()).Body(body).Execute()
	if err != nil {
		// .. Handle error
	}
	fmt.Println(response)
}
