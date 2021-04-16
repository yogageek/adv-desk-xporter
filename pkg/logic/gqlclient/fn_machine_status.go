package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/var"
)

// . "porter/util"

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

func getSourceMachineStatus() (mm []map[string]interface{}) {
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

//目前最多只能匯入三層
func ImportMachineStatus(jsonData *JsonData) {
	machineStatusDatas := jsonData.MachineStatusData

	//儲存index與parentId對應關係
	M1 := map[int]string{}
	M2 := map[int]string{}

	// debugging 目前只抓萬以下的測
	var newMachineStatusDatas []*machineStatusData
	for _, v := range machineStatusDatas {
		if v.Index < 10000 {
			newMachineStatusDatas = append(newMachineStatusDatas, v)
		}
	}

	machineStatusDatas = newMachineStatusDatas

	// debugging
	// util.PrintJson(machineStatusDatas)
	// fmt.Println("----------------------------")

	c := 0
	//fix: add set new machine status id after import
	for _, v := range machineStatusDatas {

		//channel寫法
		c++
		ChannelIn(MachineStatus, c)

		if v.Depth == 1 {
			input := model.AddMachineStatusInput{
				Name:  NamePrefix + v.Name,
				Index: IndexPrefix + v.Index,
				Color: v.Color,
			}
			//set new id after import
			v.NewId = AddMachineStatus(input).Id
			//set parent id map (根據index來判斷parentid是誰)
			M1[v.Index] = v.NewId
		}
	}

	for _, v := range machineStatusDatas {
		if v.Depth == 2 && M1[v.ParentIndex] != "" {
			input := model.AddMachineStatusInput{
				ParentId: M1[v.ParentIndex], //從parentIndex找出Id(上面已放入新id)
				Name:     NamePrefix + v.Name,
				Index:    IndexPrefix + v.Index,
				Color:    v.Color,
			}
			//set new id
			v.NewId = AddMachineStatus(input).Id
			//set parent id map (根據index來判斷parentid是誰)
			M2[v.Index] = v.NewId
		}
	}

	for _, v := range machineStatusDatas {
		if v.Depth == 3 && M2[v.ParentIndex] != "" {
			input := model.AddMachineStatusInput{
				ParentId: M2[v.ParentIndex], //從parentIndex找出Id(上面已放入新id)
				Name:     NamePrefix + v.Name,
				Index:    IndexPrefix + v.Index,
				Color:    v.Color,
			}
			v.NewId = AddMachineStatus(input).Id //這裡也要設
		}
	}
	// debugging
	// util.PrintJson(machineStatusDatas)
}

/*
	這裡要改寫法 ok
	當第一層導入後 返回parentId跟index 要存起來
	當導入第二層時 parentId要放剛剛返回的Id
*/

//這裡要處理parentId的問題
