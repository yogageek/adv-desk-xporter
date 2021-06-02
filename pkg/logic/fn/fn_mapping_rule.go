package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/gochan"
	. "porter/pkg/logic/gql"
	. "porter/pkg/logic/vars"
)

//匯入mapping rule總共只需要這些
/*
{
  "input": {
    "name": "yoga4",
    "pType": "Number",
    "codes": [ //可放多個
		{
		"code": "1",
		"message": "yoogatest",
		"translations": {
			"lang": "",
			"message": ""
		},
		"statusId": "TVN0YXR1cw.X56UsK_4oAAGcJKU"
			},
		{
		"code": "2",
		"message": "yoogatest",
		"translations": {
			"lang": "",
			"message": ""
		},
		"statusId": ""
		}
    ]
  }
}
*/

//所以查詢只需要拿
/*
最外層的
parameterMappings {
	name
	pType
	codes { #parameterMappingCode
		code
		message
		statusId
		messages {
			lang
			text
		}
	}
}
*/

var DefaultLang string

func SetDefaultLang() {
	DefaultLang = GetDefaultLangFromEnv()
}

func ImportMappingRule(jsonData *JsonData) {
	//set DefaultLang for ImportMappingRule
	SetDefaultLang()

	mappingRuleDatas := jsonData.MappingRuleData

	c := 0
	for _, v := range mappingRuleDatas {

		//channel寫法
		c++
		ChannelIn(MappingRule, c)

		codes := []model.AddParameterMappingCodesEntry{}
		for _, v := range v.Detail {
			code := model.AddParameterMappingCodesEntry{
				Code:     v.Code,
				Message:  v.Message,
				StatusId: v.StatusId,
				Translations: model.ParameterMappingCodeTranslationEntry{
					Lang:    DefaultLang, //#這裡暫時先放要導入的環境的預設語言(因為如果預設語言不存在會錯)
					Message: v.Text,      //fix to v.message?
				},
			}
			codes = append(codes, code)
		}

		input := model.AddParameterMappingRuleInput{
			Name:  NamePrefix + v.Name,
			PType: v.PType,
			Codes: codes,
		}
		//set new id
		parameterMappings := AddParameterMappingRule(input)
		v.NewId = parameterMappings.Id
	}
}
