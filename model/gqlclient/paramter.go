package model

/*
//for query, no need json tag------
type Parameters struct { //這裡一定要大寫
	Id          string `json:"id,omitempty"` //裡面的欄位名稱一定要大寫開頭, 而且型態要正確!
	GroupId     string `json:"groupId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	IsStation   *bool  `json:"isStation,omitempty"` //omitempty的代表參數不一定要放或不用放
}

var QueryParameters struct { //這行可隨便定義
	Machines []Machines //為graphql規格定義
}

//add後返回的查詢 注意和input是兩回事
var AddParameter struct { //這行可隨便定義 但盡量和下面同名
	AddParameter struct { //這裡只用來對應addMachine 沒其他作用
		Parameter Parameters //Machine這個名稱是對應到gql規格
	} `graphql:"addParameter(input: $input)"`
}


*/

//for input,  need json tag------
// type AddParameterInput struct { //參數
// 	Machines
// }

// 源自ui
// input:
// {machineId: "TWFjaGluZQ.YD9byHkI5QAHoyNy", name: "xxx", description: "xxx",…}
// description: "xxx"
// machineId: "TWFjaGluZQ.YD9byHkI5QAHoyNy"
// name: "xxx"
// scadaId: "scada_tcYO2SsWVWKr"
// tagId: "atmc_twm8_production_value_id"
// valueType: "Number"

type QueryParametersOb struct { //這行可隨便命名
	Machine struct { //對應規格裡第一層物件
		Parameters struct {
			Nodes    []Nodes
			PageInfo PageInfo
		} `graphql:"parameters(first: $first, after: $after)"` //注意要加在正確的struct後面
	} `graphql:"machine(id: $id)"`
}

type Nodes struct {
	ScadaId      interface{}   `json:"scadaId"`
	TagId        interface{}   `json:"tagId"`
	MachineId    interface{}   `json:"machineId"`
	Name         interface{}   `json:"name"`
	Description  interface{}   `json:"description"`
	Descriptions []Description `json:"descriptions"`
	ValueType    interface{}   `json:"valueType"`
	MappingId    interface{}   `json:"mappingId"`
	Kind         interface{}   `json:"kind"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

//新的寫法 看起來較一致
var QueryParameters struct { //這行可隨便定義
	QueryParametersOb
}

//add後返回的查詢 注意和input是兩回事
var AddParameter struct { //這行可隨便定義 但盡量和下面同名
	AddParameter struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		Parameter struct {
			Parameters
		}
	} `graphql:"addParameter(input: $input)"`
}

type AddParameterInput struct { //這裡struct名字要和graphql($input: AddParameterInput!)一樣
	MachineId   interface{} `json:"machineId"`
	Name        interface{} `json:"name"` //tag重要! 攸關轉gql後的大小寫
	Description interface{} `json:"description"`
	ValueType   interface{} `json:"valueType,omitempty"`
	MappingId   interface{} `json:"mappingId,omitempty"`
	ScadaId     interface{} `json:"scadaId,omitempty"`
	TagId       interface{} `json:"tagId,omitempty"`
	Kind        interface{} `json:"kind,omitempty"`
}
