package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/gochan"
	. "porter/pkg/logic/gql"
	. "porter/pkg/logic/vars"
)

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

//處理新增的machineId

//先
func ImportProfileMachine(jsonData *JsonData) (oldAndnewId_profile, oldAndnewId_profileparameter map[string]string) {
	c := 0

	oldAndnewId_profile, oldAndnewId_profileparameter = map[string]string{}, map[string]string{}

	for _, v := range jsonData.ProfileData {

		//channel寫法
		c++
		ChannelIn(Profile, c)

		input := model.AddProfileMachineInput{
			Name:        NamePrefix + string(v.Name), //# wait to refac
			Description: v.Description,
			ImageUrl:    v.ImageUrl,
		}
		newId := AddProfileMachine(input)
		oldAndnewId_profile[v.Id] = newId
		oldAndnewId_profileparameter = importProfileParameter(newId, *v)
	}
	return
}

//後
func importProfileParameter(machineId string, profileData ProfileData) (oldAndnewId map[string]string) {
	oldAndnewId = map[string]string{}
	for _, v := range profileData.Parameters {
		input := model.AddParameterInput{
			MachineId:   machineId,
			Name:        v.Name,
			Description: v.Description,
			ValueType:   v.ValueType,
			MappingId:   v.Mapping.Id, //這裡是放mapping rule id
		}
		id := AddProfileParameter(input) //由於沒有其他支要用到profileid 所以這裡不需要處理
		oldAndnewId[*v.Id] = id
	}
	return
}
