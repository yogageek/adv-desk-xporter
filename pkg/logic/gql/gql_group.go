package logic

import (
	"context"
	"encoding/json"
	"fmt"
	model "porter/model/gqlclient"
	. "porter/pkg/logic/client"
	gochan "porter/pkg/logic/gochan"
	. "porter/util"

	"github.com/golang/glog"
)

func QueryGroups() []model.Groups {
	gqlQuery := model.QueryGroups

	variables := map[string]interface{}{} //if no variables

	err := GclientQ.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// util.PrintGreen(b)

	return gqlQuery.Groups
}

func AddGroup(input model.AddGroupInput) (id string) {
	gqlQuery := model.AddGroup

	variables := map[string]interface{}{
		"input": input,
	}

	// //debugging
	// v, _ := json.MarshalIndent(variables, "", " ")
	// fmt.Printf("%s", v)

	err := GclientM.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		// c := gochan.GetChan()
		// c.SendToChan(err)
		gochan.GetChan().SendToChan(err)

		glog.Error(err)
		// glog.Fatal(err)
	}

	// //debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	PrintYellow(b)

	id = gqlQuery.AddGroup.Group.Id
	return
}

func AddGroupSample(parentId string) {
	input := model.AddGroupInput{
		Groups: model.Groups{
			ParentId:    parentId, //如果無也可以直接放""
			Name:        "yogaGroup3",
			Description: "test",
			TimeZone:    "Asia/Taipei",
			Coordinate:  nil,
		},
	}
	id := AddGroup(input)
	fmt.Println("id:", id)
}
