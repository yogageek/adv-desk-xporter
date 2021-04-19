package logic

import (
	"context"
	"encoding/json"
	model "porter/model/gqlclient"
	. "porter/pkg/logic/client"
	. "porter/util"

	"github.com/golang/glog"
)

func QueryParameterMappings() []model.ParameterMappings {
	gqlQuery := model.QueryParameterMappings

	variables := map[string]interface{}{} //if no variables

	err := GclientQ.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
		// glog.Fatal(err)
	}

	//debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// util.PrintGreen(b)

	return gqlQuery.ParameterMappings
}

func AddParameterMappingRule(input model.AddParameterMappingRuleInput) model.ParameterMappings {
	gqlQuery := model.AddParameterMappingRule

	variables := map[string]interface{}{
		"input": input,
	}

	//debugging
	// v, _ := json.MarshalIndent(variables, "", " ")
	// fmt.Printf("%s", v)

	err := GclientM.Mutate(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	PrintYellow(b)

	return gqlQuery.AddParameterMappingRule.MappingRule.ParameterMappings
}

func AddParameterMappingRuleSample() {
	input := model.AddParameterMappingRuleInput{
		Name:  "yogaxxxxx",
		PType: "Number",
		Codes: []model.AddParameterMappingCodesEntry{
			model.AddParameterMappingCodesEntry{
				Code:    "1",
				Message: "message",
			},
		},
	}
	AddParameterMappingRule(input)
}
