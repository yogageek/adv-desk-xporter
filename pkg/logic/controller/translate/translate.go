package translate

import (
	"encoding/json"
	"io/ioutil"
	fn "porter/pkg/logic/fn"
	gql "porter/pkg/logic/gql"

	"github.com/golang/glog"
)

func TestTranslate() {
	//step1 Read json file...
	b, err := ioutil.ReadFile("./exportingData.json")
	if err != nil {
		glog.Fatal(err)
	}

	var data fn.JsonData
	err = json.Unmarshal(b, &data)
	if err != nil {
		glog.Fatal(err)
	}
	Translate(&data)
}

func Translate(data *fn.JsonData) {
	//從export檔取得langs種類
	// var Langs []string
	// for _, v := range data.TranslationLangs {
	// 	Langs = append(Langs, v.Lang)
	// }
	// fmt.Println("langs:", Langs)

	machineStatusData := data.MachineStatusData
	for _, v := range machineStatusData {
		for _, name := range v.Names {
			gql.TranslateMachineStatus(v.Id, name.Text, name.Lang)
		}
	}

	groupData := data.GroupData
	for _, v := range groupData {
		for i, name := range v.Names {
			gql.TranslateGroup(v.Id, name.Text, name.Lang, v.Descriptions[i].Text)
		}
	}

	machineData := data.MachineData
	for _, v := range machineData {
		for i, name := range v.Names {
			gql.TranslateMachine(v.Id, name.Text, name.Lang, v.Descriptions[i].Text, v.ImageUrls[i].Text)
		}
	}

	parameterData := data.ParameterData
	for _, v := range parameterData {
		for _, vv := range v.Machine.Parameters.Nodes {
			for _, description := range vv.Descriptions {
				gql.TranslateParameter(vv.Id.(string), description.Text, description.Lang)
			}
		}
	}

	profileData := data.ProfileData
	for _, v := range profileData {
		for i, name := range v.Names {
			gql.TranslateProfileMachine(string(v.Id), name.Text, name.Text, v.Descriptions[i].Text, v.ImageUrls[i].Text)
		}
	}

	mappingRuleData := data.MappingRuleData
	for _, v := range mappingRuleData {
		for _, vv := range v.Detail {
			for _, message := range vv.Messages {
				gql.TranslateParameterMappingCode(vv.Id, message.Lang, message.Text)
			}
		}
	}

}
