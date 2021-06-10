package logic

import (
	. "porter/pkg/logic/gochan"

	model "porter/model/gqlclient"

	. "porter/pkg/logic/gql"
	. "porter/pkg/logic/vars"
)

// . "porter/util"

func ImportGroups(jsonData *JsonData) map[string]string {
	groups := jsonData.GroupData

	//oldId and newId mapping relations
	M := map[string]string{}

	c := 0
	for _, v := range groups {

		//channel寫法
		c++
		ChannelIn(Group, c)

		//先匯入上層
		if v.ParentId == "" {
			input := model.AddGroupInput{
				// ParentId:    M[v.ParentId], //groups graphql response will be list like parent-child, parent, parent, parent-child (outter first)
				Name:        v.Name,
				Description: v.Description,
				TimeZone:    v.TimeZone,
				Coordinate:  v.Coordinate,
			}
			newId := AddGroup(input)
			if newId != "" {
				M[v.Id] = newId //save new id
			}
		}
	}

	for _, v := range groups {
		//再匯入下層
		if v.ParentId != "" {
			input := model.AddGroupInput{
				ParentId:    M[v.ParentId], //groups graphql response will be list like parent-child, parent, parent, parent-child (outter first)
				Name:        v.Name,
				Description: v.Description,
				TimeZone:    v.TimeZone,
				Coordinate:  v.Coordinate,
			}
			newId := AddGroup(input)
			M[v.Id] = newId //save new id
		}
	}

	//先不處理Group三層以上 照理說graphql返回來的資料是由上而下照層級順序

	return M
}
