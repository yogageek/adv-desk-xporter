package logic

import (
	model "porter/model/gqlclient"
)

//匯入machine status總共只需要這些
/*
{
   "input":{
      "parentId":"TVN0YXR1cw.YBpMaugKbgAG-RS4",
      "name":"gplaytest",
      "index":7100,
      "color":"#96E796"
   }
}
*/

//所以查詢只需要拿
/*
最外層的
machineStatuses{
	name
	index
	color
	以及所有children的
	parentId
	name
	index
	color
}
*/

type jsonFile struct {
	machineStatusDatas []machineStatusData
	mappingRuleData    []mappingRuleData
	profileData        []profileData
}

//未來改成從model拿原始input
type machineStatusData struct {
	ParentId string
	Name     string
	Index    string
	Color    string
}

type mappingRuleData struct {
	Name   string
	PType  string
	Detail []detail
}

type detail struct {
	Code     string
	Message  string
	StatusId string
	Lang     string
	Text     string
}

type profileData struct {
	model.ProfileMachine
}

func getSourceMachineStatus() (mm []map[string]interface{}) {
	// mm := []map[string]interface{}{}
	res := QueryMachineStatuses()
	for _, v := range res {
		m := map[string]interface{}{
			"name":  v.Name,
			"index": v.Index,
			"color": v.Color,
		}
		mm = append(mm, m)
		for _, v := range v.Children {
			m := map[string]interface{}{
				"parentId": v.ParentId,
				"name":     v.Name,
				"index":    v.Index,
				"color":    v.Color,
			}
			mm = append(mm, m)
		}
	}
	return
}

func ImportMachineStatus(machineStatusDatas []machineStatusData) {
	//要模擬打資料的話就把這裡參數都改名
	for _, v := range machineStatusDatas {
		input := model.AddMachineStatusInput{
			ParentId: v.ParentId,
			Name:     v.Name,
			Index:    v.Index,
			Color:    v.Color,
		}
		AddMachineStatus(input)
	}
}
