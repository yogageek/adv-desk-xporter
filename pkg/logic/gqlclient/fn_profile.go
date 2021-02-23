package logic

import model "porter/model/gqlclient"

//import profile有兩個步驟
/*
//1.
// 新增一個profile名稱
mutation addProfileMachine($input: AddProfileMachineInput!) {
	addProfileMachine(input: $input) {
	  profileMachine{
		name
	  }
	}
}

{
  "input": {
    "name": "gplaytest",
    "description": "testestset",
    "imageUrl": ""
  }
}

//2.
// 新增profile裡的項目
mutation addProfileParameter($input: AddParameterInput!) {
  addProfileParameter(input: $input) {
    profileParameter{
      name
      # mapping{
      #   name
      # }
      description
      valueType
    }
  }
}

{
  "input": {
    "machineId":  "UGZNYWNoaW5l.YDMcl-gKbgAG-RWX",
    "name": "graqltest2",
    "description": "zzzzz",
    "valueType": "String",
  }
}
*/

//查詢profile
/*
query profileMachineList {
  profileMachines {
    name #import profile1 need
    description #import profile1 need
    imageUrl #import profile1 need
    #---
    id #import profile2 need(=machineId)
    parameters {
      mapping{
        name
        id
      }
      name
      description
      valueType
    }
  }
}
*/

func getSourceProfileMachines() (results []model.ProfileMachine) {
	// mm := []map[string]interface{}{}
	res := QueryProfileMachines()
	return res
}

//處理新增的machineId

//先
func ImportProfileMachine(profileDatas []profileData) {
	for _, v := range profileDatas {
		input := model.AddProfileMachineInput{
			Name:        v.Name,
			Description: v.Description,
			ImageUrl:    v.ImageUrl,
		}
		id := AddProfileMachine(input)
		ImportProfileParameter(id, v)
	}
}

//後
func ImportProfileParameter(machineId string, profileData profileData) {
	for _, v := range profileData.Parameters {
		input := model.AddProfileParameterInput{
			MachineId:   v.MachineId,
			Name:        v.Name,
			Description: v.Description,
			ValueType:   v.ValueType,
			MappingId:   v.Mapping.Id,
		}
		AddProfileParameter(input)
	}

}
