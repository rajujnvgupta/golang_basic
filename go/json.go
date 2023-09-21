package main

import (
	"encoding/json"
	"fmt"
)


func main() {

	WriteAuthorizationModelRequestString := "{\"schema_version\":\"1.1\",\"type_definitions\":[{\"type\":\"user\"},{\"type\":\"document\",\"relations\":{\"reader\":{\"this\":{}},\"writer\":{\"this\":{}},\"owner\":{\"this\":{}}},\"metadata\":{\"relations\":{\"reader\":{\"directly_related_user_types\":[{\"type\":\"user\"}]},\"writer\":{\"directly_related_user_types\":[{\"type\":\"user\"}]},\"owner\":{\"directly_related_user_types\":[{\"type\":\"user\"}]}}}}]}"
	xyz, _ := json.Marshal(WriteAuthorizationModelRequestString)

	var model interface{}
	json.Unmarshal(xyz, &model);
	fmt.Println(model)
}
