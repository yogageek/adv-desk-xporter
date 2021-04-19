package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/gochan"
	. "porter/pkg/logic/gql"
	. "porter/pkg/logic/vars"
)

// . "porter/util"

var loadedParameters int

func ImportParameters(jsonData *JsonData) {
	parameters := jsonData.ParameterData

	c := 0
	for _, v := range parameters {
		for _, v := range v.Machine.Parameters.Nodes {

			//channel寫法
			c++
			ChannelIn(Parameter, c)

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
