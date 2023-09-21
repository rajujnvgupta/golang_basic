package main
import (
    fgaSdk "github.com/openfga/go-sdk"
	"fmt"
	"context"
	"strconv"
	"time"
)

func main() {
    // Initialize the SDK with no auth - see "How to setup SDK client" for more options

	t1 := time.Now()
    configuration, err := fgaSdk.NewConfiguration(fgaSdk.Configuration{
        ApiScheme:      "http", //os.Getenv("FGA_SCHEME"), // Either "http" or "https", defaults to "https"
		ApiHost:        "0.0.0.0:8080", //os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        "01GX461BHKSQB62Z7WQN8XT7B5", //os.Getenv("FGA_STORE_ID"), // optional, not needed for `CreateStore` and `ListStores`, required before calling for all other methods
    })

    if err != nil {
    // .. Handle error
    }

    fgaClient := fgaSdk.NewAPIClient(configuration)


	user := "user:anne"
	for i := 0; i < 12; i++{
		if i != 0 {
			user = "user:anne"
			user = user + strconv.Itoa(i) 
			fmt.Println(user)
		}
		body := fgaSdk.CheckRequest{
			AuthorizationModelId: fgaSdk.PtrString("01GX6QBJHX2PDW2MDB6E6F2EHE"),
			TupleKey: fgaSdk.TupleKey{
				User: fgaSdk.PtrString(user),
				Relation: fgaSdk.PtrString("reader"),
				Object: fgaSdk.PtrString("document:Z"),
			},
		}
		_, _, err := fgaClient.OpenFgaApi.Check(context.Background()).Body(body).Execute()

		if err != nil {
		// .. Handle error
		}
//		fmt.Println(data)
//		fmt.Println(response)
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}

