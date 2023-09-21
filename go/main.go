package main

import (
	"os"
	"time"
	"fmt"
	"context"
	"reflect"
//	"github.com/oklog/ulid/v2"
	"github.com/openfga/openfga/cmd"
	"github.com/openfga/openfga/cmd/run"
	"github.com/openfga/openfga/pkg/storage/postgres"
	"github.com/openfga/openfga/pkg/storage/common"
	"github.com/openfga/openfga/pkg/tuple"
	openfgapb "go.buf.build/openfga/go/openfga/api/openfga/v1"
	"github.com/openfga/openfga/pkg/encoder"
	"github.com/openfga/openfga/pkg/server"
	"github.com/openfga/openfga/internal/gateway"
	"github.com/openfga/openfga/pkg/logger"
)


func check() {

	uri := "postgres://postgres:postgres@127.0.0.1:5433/postgres"
	datastore, err := postgres.New(uri, common.NewConfig())
	if err != nil {
		fmt.Println(err)
	}
	config := run.DefaultConfig()
	fmt.Println(config.Log.Format)
	logger := logger.MustNewLogger(config.Log.Format, config.Log.Level)

	var experimentals []server.ExperimentalFeatureFlag
	for _, feature := range config.Experimentals {
		experimentals = append(experimentals, server.ExperimentalFeatureFlag(feature))
	}

	fmt.Println(experimentals)


	svr := server.New(&server.Dependencies{
				Datastore:    datastore,
				Logger:       logger,
				TokenEncoder: encoder.NewBase64Encoder(),
				Transport:    gateway.NewRPCTransport(logger),
			}, &server.Config{
				ResolveNodeLimit:         config.ResolveNodeLimit,
				ChangelogHorizonOffset:   config.ChangelogHorizonOffset,
				ListObjectsDeadline:      config.ListObjectsDeadline,
				ListObjectsMaxResults:    config.ListObjectsMaxResults,
				Experimentals:            experimentals,
				AllowEvaluating1_0Models: config.AllowEvaluating1_0Models,
				AllowWriting1_0Models:    config.AllowWriting1_0Models,
			})
	fmt.Println(svr)

	storeID := "01GXSZYRPJSH2WYVMKF60T9Y2B"
	authorizationModelId := "01GXT0W64VE5KVZA5G20DFTA00"
	resp, err1 := svr.Check(context.Background(), &openfgapb.CheckRequest{
        StoreId:              storeID,
        TupleKey:             tuple.NewTupleKey("document:Z", "writer", "user:anne"),
        AuthorizationModelId: authorizationModelId,
    })


	if resp == nil {
		fmt.Println("no relation exists")
	}else{
		fmt.Println("resp: ", resp)
	}
	fmt.Println("err1: ", err1)
	fmt.Println(experimentals)

	fmt.Println("type of check response = ", reflect.TypeOf(resp))
}

//func create(){
//
//	uri := "postgres://postgres:postgres@127.0.0.1:5433/postgres"
//	postgres, err := postgres.New(uri, common.NewConfig())
//	store, err1 := postgres.CreateStore(context.Background(), &openfgapb.Store{
//						Id:   ulid.Make().String(),
//						Name: "rajutesting",
//					})
//	if err != nil {
//		fmt.Println(err)
//	}
//}



func main() {

	t1 := time.Now()
	check()
	t2 := time.Now()
    fmt.Println(t2.Sub(t1))

	return

	if false {
	rootCmd := cmd.NewRootCommand()
	runCmd := run.NewRunCommand()
	rootCmd.AddCommand(runCmd)
	migrateCmd := cmd.NewMigrateCommand()
	rootCmd.AddCommand(migrateCmd)

	versionCmd := cmd.NewVersionCommand()
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
}
