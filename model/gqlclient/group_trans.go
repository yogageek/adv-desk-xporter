package model

/*
mutation translateGroup($input: TranslateGroupInput!) {
  translateGroup(input: $input) {
    group {
      parentId
      id
      name
    }
  }
}

{
  "input": {
    "id":  "TWFjaGluZQ.YFCF6HkI5QAHoyd9",
    "lang": "testMachine4",
    "name":  "testing",
    "description": "Number"
  }
}
*/

//add後返回的查詢 注意和input是兩回事
var TranslateGroup struct { //這行可隨便定義 但盡量和下面同名
	TranslateGroup struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		Group struct {
			Name string
		} //為graphql規格定義
	} `graphql:"translateGroup(input: $input)"`
}

type TranslateGroupInput struct { //參數
	Id          string `json:"id"`
	Name        string `json:"name"`
	Lang        string `json:"lang"`
	Description string `json:"description"`
}
