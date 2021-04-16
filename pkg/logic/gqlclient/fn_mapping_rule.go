package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/var"
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

func getSourceMappingRule() (results []map[string]interface{}) {
	// mm := []map[string]interface{}{}

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
	return
}

func ImportMappingRule(jsonData *JsonData) {
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
					Message: v.Text,
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
