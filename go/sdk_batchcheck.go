package main

import (
    openfga "github.com/openfga/go-sdk"
    . "github.com/openfga/go-sdk/client"
//    "github.com/openfga/go-sdk/credentials"
 //   "os"
	"context"
	"fmt"
)

func main() {
	store_id := "01GYC2YN1QBEA1MHF650GCW41G"
//	model_id := openfga.PtrString("01GYQQ6Q4JWK9J2EEEAQR6WRCM")
    fgaClient, err := NewSdkClient(&ClientConfiguration{
		ApiScheme:		        "http",
		ApiHost:				"0.0.0.0:8080",
        StoreId:                store_id, //openfga.PtrString("01GYC2YN1QBEA1MHF650GCW41G"),
 //       AuthorizationModelId:   model_id,
    })

    if err != nil {
        // .. Handle error
    }
	options := ClientListStoresOptions{
		PageSize:          openfga.PtrInt32(10),
	    ContinuationToken: openfga.PtrString("..."),
	}
	stores, err1 := fgaClient.ListStores(context.Background()).Options(options).Execute()
	fmt.Println(stores, err1)
    if err1 != nil {
        // .. Handle error
    }
}

