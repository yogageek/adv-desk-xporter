package main

import (
	logic "porter/pkg/logic/gqlclient"
)

func init() {

	logic.RefreshToken()
	logic.InitGqlClientAndToken()
}

func main() {
	// logic.AddMachineStatus("", "yoga", 8111, "#008cd6")

	// r := logic.QueryParameterMappings()
	// var i []interface{}
	// for _, v := range r {
	// 	for _, v := range v.Codes {
	// 		i = append(i, v.Status.Index)
	// 	}
	// }
	// fmt.Println(i)

	// logic.AddParameterMappingRuleSample()

	logic.Export()

	// logic.GetSource_mapping_rule()

	// func() {
	// 	input := model.AddProfileMachineInput{
	// 		Name:        "yoga1111122",
	// 		Description: "ttestsetset",
	// 		ImageUrl:    "",
	// 	}
	// 	logic.AddProfileMachine(input)
	// }()
}
