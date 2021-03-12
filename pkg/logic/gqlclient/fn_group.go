package logic

import model "porter/model/gqlclient"

// . "porter/util"

func getSourceGroups() []model.Groups {
	return QueryGroups()
}

//#還沒處理newid問題
func importGroups(jsonData *jsonData) {
	groups := jsonData.GroupData

	//oldId and newId mapping relations
	M := map[string]string{}
	for _, v := range groups {
		input := model.AddGroupInput{
			Groups: model.Groups{
				ParentId:    M[v.ParentId], //groups graphql response will be list like parent-child, parent, parent, parent-child (outter first)
				Name:        v.Name,
				Description: v.Description,
				TimeZone:    v.TimeZone,
				Coordinate: &model.Coordinate{
					Longitude: v.Coordinate.Longitude,
					Latitude:  v.Coordinate.Latitude,
				},
			},
		}
		newId := AddGroup(input)
		M[v.Id] = newId
	}
}

func replaceIds() {

}
