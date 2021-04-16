package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/var"
)

// . "porter/util"

func GetSourceGroups() []model.Groups {
	return QueryGroups()
}

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
				Groups: model.Groups{ //有就放parentId 找不到就放""
					// ParentId:    M[v.ParentId], //groups graphql response will be list like parent-child, parent, parent, parent-child (outter first)
					Name:        v.Name,
					Description: v.Description,
					TimeZone:    v.TimeZone,
					Coordinate:  v.Coordinate,
				},
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
				Groups: model.Groups{ //有就放parentId 找不到就放""
					ParentId:    M[v.ParentId], //groups graphql response will be list like parent-child, parent, parent, parent-child (outter first)
					Name:        v.Name,
					Description: v.Description,
					TimeZone:    v.TimeZone,
					Coordinate:  v.Coordinate,
				},
			}
			newId := AddGroup(input)
			M[v.Id] = newId //save new id
		}
	}

	//先不處理Group三層以上 照理說graphql返回來的資料是由上而下照層級順序

	return M
}
