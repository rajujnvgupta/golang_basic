package main

import (
	openfgapb "go.buf.build/openfga/go/openfga/api/openfga/v1"
	"fmt"
)

func main(){
	expected := []*openfgapb.TupleKey{
		tuple.NewTupleKey("document:doc1", "viewer", "bill"),
		tuple.NewTupleKey("document:doc2", "editor", "bob"),
	}
	fmt.Println(expected)
}
