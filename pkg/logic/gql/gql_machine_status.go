package logic

import (
	"context"
	"encoding/json"
	"fmt"
	model "porter/model/gqlclient"
	. "porter/pkg/logic/client"
	. "porter/util"

	"github.com/golang/glog"
	"github.com/prometheus/common/log"
	"github.com/shurcooL/graphql"
)

//MachineStatusHierarchy(可以查到所有machine的status)
func QueryMachineStatuses() []model.MachineStatuses {
	gqlQuery := model.QueryMachineStatuses

	variables := map[string]interface{}{
		"layer1Only": graphql.Boolean(true),
	}

	err := GclientQ.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// util.PrintGreen(b)

	return gqlQuery.MachineStatuses
}

func AddMachineStatus(input model.AddMachineStatusInput) model.MachineStatus {
	gqlQuery := model.AddMachineStatus

	variables := map[string]interface{}{
		"input": input,
	}

	//debugging
	// v, _ := json.MarshalIndent(variables, "", " ")
	// fmt.Printf("%s", v)

	err := GclientM.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
		// glog.Fatal(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	PrintYellow(b)

	return gqlQuery.AddMachineStatus.MachineStatus
}

func AddMachineStatusSample(parentId string) {
	input := model.AddMachineStatusInput{
		ParentId: parentId, //如果無也可以直接放""
		Name:     "yoga",
		Index:    "1000",
		Color:    "#96E796",
	}
	AddMachineStatus(input)
}

func UpdateMachineStatus(id, name, color interface{}) {
	gqlQuery := model.UpdateMachineStatus
	variables := map[string]interface{}{
		"input": model.UpdateMachineStatusInput{
			Id:    id,
			Name:  name,
			Color: color,
		},
	}

	//debugging
	// v, _ := json.MarshalIndent(variables, "", " ")
	// fmt.Printf("%s", v)

	err := GclientM.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		log.Error(err)
		glog.Fatal(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	fmt.Printf("%s", b)
}
