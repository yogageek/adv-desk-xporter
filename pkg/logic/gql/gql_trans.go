package logic

import (
	"context"
	"encoding/json"
	"fmt"
	. "porter/pkg/logic/client"

	// . "porter/util"
	util "porter/util"

	"github.com/golang/glog"
)

func Query(input interface{}, gqlQuery interface{}) interface{} {
	// gqlQuery := model.TranslateGroup

	variables := map[string]interface{}{
		"input": input,
	}

	err := GclientQ.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	util.PrintGreen(b)

	return gqlQuery
}

func Mutate(input interface{}, gqlQuery interface{}) interface{} {

	variables := map[string]interface{}{
		"input": input,
	}

	//debugging
	v, _ := json.MarshalIndent(variables, "", " ")
	fmt.Printf("%s", v)

	err := GclientM.Mutate(context.Background(), gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	util.PrintYellow(b)

	return gqlQuery
}
