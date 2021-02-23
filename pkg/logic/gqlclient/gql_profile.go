package logic

import (
	"context"
	"encoding/json"
	"fmt"
	model "porter/model/gqlclient"

	"github.com/golang/glog"
)

func QueryProfileMachines() []model.ProfileMachine {
	gqlQuery := model.QueryProfileMachines

	variables := map[string]interface{}{} //if no variables

	err := gclient.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// fmt.Printf("%s", b)

	return gqlQuery.ProfileMachines
}

func AddProfileMachine(input model.AddProfileMachineInput) (id string) {
	gqlQuery := model.AddProfileMachine

	variables := map[string]interface{}{
		"input": input,
	}

	// //debugging
	// v, _ := json.MarshalIndent(variables, "", " ")
	// fmt.Printf("%s", v)

	err := gclient.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	// //debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// fmt.Printf("%s", b)

	id = gqlQuery.AddProfileMachine.ProfileMachine.Id
	return
}

func AddProfileParameter(input model.AddProfileParameterInput) {
	gqlQuery := model.AddProfileParameter

	variables := map[string]interface{}{
		"input": input,
	}

	// //debugging
	// // v, _ := json.MarshalIndent(variables, "", " ")
	// // fmt.Printf("%s", v)

	err := gclient.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	fmt.Printf("%s", b)
}
