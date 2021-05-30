package model

import "github.com/shurcooL/graphql"

//for query, no need json tag------
type ParameterMappings struct { //這裡一定要大寫
	Id    string
	Name  graphql.String //裡面的欄位名稱一定要大寫開頭, 而且型態要正確!
	PType graphql.String
	Codes []Code
}

type Code struct {
	Code     graphql.String
	Message  graphql.String
	StatusId graphql.String
	Messages []Message //多語言資料
}

type Message struct {
	Lang graphql.String
	Text graphql.String
}

var QueryParameterMappings struct { //這行可隨便定義
	ParameterMappings []ParameterMappings //為graphql規格定義
}

//add後返回的查詢 注意和input是兩回事
var AddParameterMappingRule struct { //這行可隨便定義 但盡量和下面同名
	AddParameterMappingRule struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		MappingRule struct { //為graphql規格定義
			ParameterMappings
		}
	} `graphql:"addParameterMappingRule(input: $input)"`
}

//for input,  need json tag------

type AddParameterMappingRuleInput struct { //參數
	//要塞的參數 大小寫要注意
	Name  interface{}                     `json:"name"` //tag重要! 攸關轉gql後的大小寫
	PType interface{}                     `json:"pType"`
	Codes []AddParameterMappingCodesEntry `json:"codes"`

	// use extend style fail
	// ParameterMappings
}

type AddParameterMappingCodesEntry struct {
	Code         interface{}                          `json:"code"`
	Message      interface{}                          `json:"message"`
	StatusId     interface{}                          `json:"statusId"` //用來綁定machine status
	Translations ParameterMappingCodeTranslationEntry `json:"translations"`
}

type ParameterMappingCodeTranslationEntry struct {
	Lang    interface{} `json:"lang"`
	Message interface{} `json:"message"`
}
