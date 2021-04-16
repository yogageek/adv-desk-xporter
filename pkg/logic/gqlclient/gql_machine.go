package logic

import (
	"context"
	"encoding/json"
	model "porter/model/gqlclient"
	"porter/util"

	"github.com/golang/glog"
)

func QueryMachines() []model.Machines {
	gqlQuery := model.QueryMachines

	variables := map[string]interface{}{} //if no variables

	err := gclientQ.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	util.PrintGreen(b)

	return gqlQuery.Machines
}

func AddMachine(input model.AddMachineInput) (id string) {
	gqlQuery := model.AddMachine

	variables := map[string]interface{}{
		"input": input,
	}

	// //debugging
	// v, _ := json.MarshalIndent(variables, "", " ")
	// fmt.Printf("%s", v)

	err := gclientM.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
		// glog.Fatal(err)
	}
	//如果返回這個error代表可能找不到groupId
	//ERROR: logging before flag.Parse: E0309 15:27:03.319725   12064 gql_machine.go:43] Machine validation failed: group: Validator failed for path `group` with value `604060f3c1ae9d0006a44699`

	// //debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	util.PrintYellow(b)

	id = gqlQuery.AddMachine.Machine.Id
	return
}

func AddMachineSample() {
	input := model.AddMachineInput{
		Machines: model.Machines{
			GroupId:     "R3JvdXA.YD9bt3kI5QAHoyNw",
			Name:        "testMachine4",
			Description: "testing",
			ImageUrl:    "",
			IsStation:   nil,
		},
	}
	AddMachine(input)
}
