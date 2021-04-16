package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/var"
)

// . "porter/util"

func getSourceMachines() []model.Machines {
	return QueryMachines()
}

func getMachineIds(machines []model.Machines) (ids []string) {
	for _, v := range machines {
		ids = append(ids, v.Id)
	}
	return
}

func ImportMachines(jsonData *JsonData) map[string]string {
	machines := jsonData.MachineData

	//oldId and newId mapping relations
	M := map[string]string{}

	c := 0
	for _, v := range machines {

		//channel寫法
		c++
		ChannelIn(Machine, c)

		input := model.AddMachineInput{
			Machines: model.Machines{ //有就放parentId 找不到就放""
				GroupId:     v.GroupId,
				Name:        v.Name,
				Description: v.Description,
				ImageUrl:    v.ImageUrl,
				IsStation:   v.IsStation,
			},
		}
		newId := AddMachine(input)
		M[v.Id] = newId //save new id
	}
	return M
}
