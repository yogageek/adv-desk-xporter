package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/client"
	. "porter/pkg/logic/gql"
)

func (o groups) GetSource() interface{} {
	return QueryGroups()
}

func (o machineStatus) GetSource() interface{} {
	mm := []map[string]interface{}{}

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
	return mm
}

func (o machines) GetSource() interface{} {
	machines := QueryMachines()
	SetMachineIds(machines)
	return machines
}

func (o mappingRule) GetSource() interface{} {
	results := []map[string]interface{}{}

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
	return results
}

func (o parameters) GetSource() interface{} {
	machineIds := GetMachineIds()

	var objects []model.QueryParametersOb
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

func (o profileMachine) GetSource() interface{} {
	// mm := []map[string]interface{}{}
	res := QueryProfileMachines()
	return res
}

func GetSourceTranslations() interface{} {
	return QueryTranslationLangs(GclientQ)
}

//------------------------------------------------

var machineIds []string

func SetMachineIds(machines []model.Machines) (ids []string) {
	for _, v := range machines {
		ids = append(ids, v.Id)
	}
	machineIds = ids
	return
}

func GetMachineIds() (ids []string) {
	return machineIds
}
