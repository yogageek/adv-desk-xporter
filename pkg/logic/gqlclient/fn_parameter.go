package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/var"
)

// . "porter/util"

func getSourceParameters(machineIds []string) (objects []model.QueryParametersOb) {

	for _, id := range machineIds {
		cursor := ""
	again:
		// fmt.Println(cursor)
		res := QueryParameters(id, cursor)
		objects = append(objects, res)
		if res.Machine.Parameters.PageInfo.HasNextPage {
			cursor = res.Machine.Parameters.PageInfo.EndCursor
			goto again
		}
	}
	// return objects[0:1] //debug用
	return objects
}

var loadedParameters int

func ImportParameters(jsonData *JsonData) {
	parameters := jsonData.ParameterData

	c := 0
	for _, v := range parameters {

		//channel寫法
		c++
		ChannelCount(Parameter, c)

		for _, v := range v.Machine.Parameters.Nodes {
			input := model.AddParameterInput{
				MachineId:   v.MachineId,
				Name:        v.Name,
				Description: v.Description,
				ValueType:   v.ValueType,
				MappingId:   v.MappingId,
				Kind:        v.Kind,
				// ScadaId:     v.ScadaId,
				// TagId:       v.TagId,
			}
			// util.Lg.PrintJson(input)
			// fmt.Println(input.Name)
			// fmt.Println(v.Name)
			// fmt.Println(input.Description)
			// fmt.Println(v.Description)

			AddParameter(input)
			loadedParameters++
		}
	}
}
