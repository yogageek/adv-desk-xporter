package logic

import (
	"context"
	"encoding/json"
	model "porter/model/gqlclient"
	. "porter/pkg/logic/client"
	. "porter/util"

	"github.com/golang/glog"
)

func QueryProfileMachines() []model.ProfileMachine {
	gqlQuery := model.QueryProfileMachines

	variables := map[string]interface{}{} //if no variables

	err := GclientQ.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// util.PrintGreen(b)

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

	err := GclientM.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
		// glog.Fatal(err)
	}

	// //debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// util.PrintYellow(b)

	id = gqlQuery.AddProfileMachine.ProfileMachine.Id
	return
}

func AddProfileParameter(input model.AddParameterInput) {
	gqlQuery := model.AddProfileParameter

	variables := map[string]interface{}{
		"input": input,
	}

	// //debugging
	// // v, _ := json.MarshalIndent(variables, "", " ")
	// // fmt.Printf("%s", v)

	err := GclientM.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
		// glog.Fatal(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	PrintYellow(b)
}
