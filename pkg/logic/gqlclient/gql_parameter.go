package logic

import (
	"context"
	model "porter/model/gqlclient"

	"github.com/golang/glog"
	"github.com/shurcooL/graphql"
)

func QueryParameters(machineId, afterCursor string) model.QueryParametersOb {
	gqlQuery := model.QueryParameters

	variables := map[string]interface{}{
		"id":    graphql.ID(machineId),
		"first": graphql.Int(100), //代表查100個(已經上限)
		"after": graphql.String(afterCursor),
	} //if no variables

	err := gclientQ.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	// //debugging
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// util.PrintGreen(b)

	return gqlQuery.QueryParametersOb
}

func AddParameter(input model.AddParameterInput) {
	gqlQuery := model.AddParameter

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
	// b, _ := json.MarshalIndent(gqlQuery, "", " ")
	// util.PrintYellow(b)

}

func AddParameterSample() {
	input := model.AddParameterInput{
		MachineId:   "TWFjaGluZQ.YEg1wHkI5QAHoyOt",
		Name:        "testMachinGo1",
		Description: "testing",
		ValueType:   "Number",
		ScadaId:     nil,
		TagId:       nil,
		Kind:        nil,
	}
	AddParameter(input)
}
