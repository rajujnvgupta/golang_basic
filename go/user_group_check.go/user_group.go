// ApiTokenIssuer, ApiAudience, ClientId and ClientSecret are optional.
package main

import (
    fgaSdk "github.com/openfga/go-sdk"
	"strconv"
	"fmt"
	"context"
	"time"
)

func main() {
    // Initialize the SDK with no auth - see "How to setup SDK client" for more options
	t1 := time.Now()
    configuration, err1 := fgaSdk.NewConfiguration(fgaSdk.Configuration{
        ApiScheme:      "http", //os.Getenv("FGA_SCHEME"), // Either "http" or "https", defaults to "https"
		ApiHost:        "0.0.0.0:8080", //os.Getenv("FGA_API_HOST"), // required, define without the scheme (e.g. api.openfga.example instead of https://api.openfga.example)
        StoreId:        "01GX461BHKSQB62Z7WQN8XT7B5", //os.Getenv("FGA_STORE_ID"), // optional, not needed for `CreateStore` and `ListStores`, required before calling for all other methods
    })

    if err1 != nil {
    // .. Handle error
    }
// write mode into store
    fgaClient := fgaSdk.NewAPIClient(configuration)
/////////////      add model user group to store 
//	WriteAuthorizationModelRequestString := "{\"schema_version\":\"1.1\",\"type_definitions\":[{\"type\":\"user\"},{\"type\":\"document\",\"relations\":{\"editor\":{\"this\":{}}},\"metadata\":{\"relations\":{\"editor\":{\"directly_related_user_types\":[{\"type\":\"team\",\"relation\":\"member\"}]}}}},{\"type\":\"team\",\"relations\":{\"member\":{\"this\":{}}},\"metadata\":{\"relations\":{\"member\":{\"directly_related_user_types\":[{\"type\":\"user\"}]}}}}]}"
//	var body fgaSdk.WriteAuthorizationModelRequest
////	var body map[string]interface{}
//	if err := json.Unmarshal([]byte(WriteAuthorizationModelRequestString), &body); err != nil {
//      fmt.Printf("%s", err)
//      // .. Handle error
//      return
//	} //
//	data, response, err := fgaClient.OpenFgaApi.WriteAuthorizationModel(context.Background()).Body(body).Execute()
//	fmt.Println(data, response, err)
//	if err != nil {
//      // .. Handle error
//	}
//////////////


//	user := "user:alice"
//	users := []string{user}
//	for i := 0; i < 10; i++{
//		if i != 0 {
//			user = "user:alice"
//			user = user + strconv.Itoa(i)
//			fmt.Println(user)
//			users = append(users)
//		}
//		body := fgaSdk.WriteRequest{
//			Writes:  &fgaSdk.TupleKeys{
//				TupleKeys: []fgaSdk.TupleKey {
//					{
//						User: fgaSdk.PtrString(user),
//						Relation: fgaSdk.PtrString("member"),
//						Object: fgaSdk.PtrString("team:writers"),
//					},
//				},
//			},
//			AuthorizationModelId: fgaSdk.PtrString("01GXKH147FRHYYZZ44VJVWV65P")}
//
//		_, _, err := fgaClient.OpenFgaApi.Write(context.Background()).Body(body).Execute()
//		if err != nil {
//			// .. Handle error
//		}
//	}
//	fmt.Println(users)
//
//	body := fgaSdk.WriteRequest{
//		Writes:  &fgaSdk.TupleKeys{
//			TupleKeys: []fgaSdk.TupleKey {
//				{
//					// Set of users related to 'team:writers' as 'member'
//					User: fgaSdk.PtrString("team:writers#member"),
//					Relation: fgaSdk.PtrString("editor"),
//					Object: fgaSdk.PtrString("document:meeting_notes.doc"),
//				},
//			},
//		},
//		AuthorizationModelId: fgaSdk.PtrString("01GXKH147FRHYYZZ44VJVWV65P")}
//
//	_, response, err2 := fgaClient.OpenFgaApi.Write(context.Background()).Body(body).Execute()
//	if err2 != nil {
//		// .. Handle error
//	}
//	fmt.Println(response)

	user := "user:alice"
	for i := 0; i < 10; i++{
		if i != 0 {
			user = "user:alice"
			user = user + strconv.Itoa(i) 
			fmt.Println(user)
		}

		body := fgaSdk.CheckRequest{
			AuthorizationModelId: fgaSdk.PtrString("01GXKH147FRHYYZZ44VJVWV65P"),
			TupleKey: fgaSdk.TupleKey{
				User: fgaSdk.PtrString(user),
				Relation: fgaSdk.PtrString("editor"),
				Object: fgaSdk.PtrString("document:meeting_notes.doc"),
			},
		}
		_, _, err := fgaClient.OpenFgaApi.Check(context.Background()).Body(body).Execute()
//		fmt.Println(data, response, err)

		if err != nil {
		// .. Handle error
		}
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
