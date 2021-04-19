package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/client"
	. "porter/pkg/logic/gql"
)

func (o groups) GetSource() []model.Groups {
	return QueryGroups()
}

func (o machineStatus) GetSource() (mm []map[string]interface{}) {
	// mm := []map[string]interface{}{}
	res := QueryMachineStatuses()
	for _, v := range res {
		m := map[string]interface{}{
			"id":    v.Id,
			"name":  v.Name,
			"index": v.Index,
			"color": v.Color,
			"depth": v.Depth,
		}
		mm = append(mm, m)
		for _, v := range v.Children {
			m := map[string]interface{}{
				"id":          v.Id,
				"parentId":    v.ParentId,
				"parentIndex": v.Parent.Index,
				"name":        v.Name,
				"index":       v.Index,
				"color":       v.Color,
				"depth":       v.Depth,
			}
			mm = append(mm, m)
			for _, v := range v.Children {
				m := map[string]interface{}{
					"id":          v.Id,
					"parentId":    v.ParentId,
					"parentIndex": v.Parent.Index,
					"name":        v.Name,
					"index":       v.Index,
					"color":       v.Color,
					"depth":       v.Depth,
				}
				mm = append(mm, m)
			}
		}
	}
	// debugging
	// util.PrintJson(mm)
	return
}

func (o machines) GetSource() []model.Machines {
	return QueryMachines()
}

func (o mappingRule) GetSource() (results []map[string]interface{}) {
	// mm := []map[string]interface{}{}

	res := QueryParameterMappings()
	for _, v := range res {
		M := map[string]interface{}{
			"id":    v.Id,
			"name":  v.Name,
			"pType": v.PType,
		}

		//array
		MM := []map[string]interface{}{}
		for _, v := range v.Codes {
			m := map[string]interface{}{
				"code":     v.Code,
				"message":  v.Message,
				"statusId": v.StatusId,
			}
			for _, v := range v.Messages { //目前messages只有一組
				mm := map[string]interface{}{
					"lang": v.Lang,
					"text": v.Text,
				}
				for k, v := range mm {
					m[k] = v
				}
			}
			MM = append(MM, m)
		}
		M["detail"] = MM

		results = append(results, M)
	}
	return
}

func (o parameters) GetSource(machineIds []string) (objects []model.QueryParametersOb) {

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

func (o profileMachine) GetSource() (results []model.ProfileMachine) {
	// mm := []map[string]interface{}{}
	res := QueryProfileMachines()
	return res
}

func GetSourceTranslations() []model.TranslationLangs {
	return QueryTranslationLangs(GclientQ)
}
