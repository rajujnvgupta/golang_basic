package main

import (
	"context"
	"fmt"
	. "github.com/openfga/go-sdk/client"
	openfga "github.com/openfga/go-sdk"
	"reflect"
	//	"strconv"
)
func main() {
	store_id := "01GYC2YN1QBEA1MHF650GCW41G"
//	model_id := openfga.PtrString("01GYQQ6Q4JWK9J2EEEAQR6WRCM")
    fgaClient, err := NewSdkClient(&ClientConfiguration{
        ApiScheme:				"http",
		ApiHost:				"0.0.0.0:8080",
        StoreId:				store_id, //openfga.PtrString("01GYC2YN1QBEA1MHF650GCW41G"),
    })

    if err != nil {
        // .. Handle error
    }
	options := ClientBatchCheckOptions{
		// You can rely on the model id set in the configuration or override it for this specific request
		AuthorizationModelId: openfga.PtrString("01GYQQ6Q4JWK9J2EEEAQR6WRCM"),
		MaxParallelRequests: openfga.PtrInt32(5), // Max number of requests to issue in parallel, defaults to 10
	}
	body := ClientBatchCheckBody{
		{
			User:     "user:kamlesh",
			Relation: "can_edit",
			Object:   "folder:service1",
			ContextualTuples: &[]ClientTupleKey{
				{
					User:     "user:kamlesh",
					Relation: "user_in_context",
					Object:   "user_groups:ug1",
				},
				{
					User:     "user:kamlesh",
					Relation: "folder_in_context",
					Object:   "folder_groups:it",
				},
			},
		},
		{
			User:     "user:ramesh",
			Relation: "can_edit",
			Object:   "folder_groups:hr",
			ContextualTuples: &[]ClientTupleKey{
				{
					User:     "user:ramesh",
					Relation: "user_in_context",
					Object:   "user_groups:ug1",
				},
				{
					User:     "user:ramesh",
					Relation: "folder_in_context",
					Object:   "folder_groups:hr",
				},
			},
		},
	}
	data, err := fgaClient.BatchCheck(context.Background()).Body(body).Options(options).Execute()

	/// check docs of CheckResponse.md in go-sdk
	fmt.Println(*data)
	for a , b := range *data{
		fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
		fmt.Println(b.HasResolution())
		fmt.Println(b.HasAllowed())
		fmt.Println(b.GetAllowed())
	}
	fmt.Println(reflect.TypeOf(data))
	//data (variable of type *[]client.ClientBatchCheckSingleResponse)
}
