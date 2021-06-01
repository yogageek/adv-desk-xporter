package model

import "github.com/shurcooL/graphql"

//for query, no need json tag------
type ProfileMachine struct { //這裡一定要大寫
	//---import profileMachines need
	Id          graphql.String
	Name        graphql.String //裡面的欄位名稱一定要大寫開頭, 而且型態要正確! #import profile1 need
	Description graphql.String
	ImageUrl    graphql.String
	//---import profileParameters need
	Parameters []Parameters
	//新增多語言資料
	Names        []Name        `json:"names"`
	Descriptions []Description `json:"descriptions"`
	ImageUrls    []ImageUrl    `json:"imageUrls"`
}

type Parameters struct {
	Id           graphql.String
	MachineId    graphql.String
	Name         graphql.String
	Description  graphql.String
	Descriptions []Description `json:"descriptions"`
	ValueType    graphql.String
	Mapping      Mapping
}

//綁到mapping rule
type Mapping struct {
	Name graphql.String
	Id   graphql.String
}

var QueryProfileMachines struct { //這行可隨便定義
	ProfileMachines []ProfileMachine //為graphql規格定義
}

//add後返回的查詢 注意和input是兩回事
var AddProfileMachine struct { //這行可隨便定義 但盡量和下面同名
	AddProfileMachine struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		ProfileMachine struct { //為graphql規格定義
			Id   string //-->需要給AddProfileParameter的machineId用
			Name string
		}
	} `graphql:"addProfileMachine(input: $input)"`
}

type AddProfileMachineInput struct {
	Name        interface{} `json:"name"` //tag重要! 攸關轉gql後的大小寫
	Description interface{} `json:"description"`
	ImageUrl    interface{} `json:"imageUrl"`
}

//add後返回的查詢 注意和input是兩回事
var AddProfileParameter struct { //這行可隨便定義 但盡量和下面同名
	AddProfileParameter struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		ProfileParameter struct {
			Name string
		}
	} `graphql:"addProfileParameter(input: $input)"`
}

// #M 被parameter.go中定義的取代
// type AddParameterInput struct { //這裡struct名字要和graphql($input: AddParameterInput!)一樣
// 	MachineId   interface{} `json:"machineId"`
// 	Name        interface{} `json:"name"` //tag重要! 攸關轉gql後的大小寫
// 	Description interface{} `json:"description"`
// 	ValueType   interface{} `json:"valueType"`
// 	MappingId   interface{} `json:"mappingId"`
// }
