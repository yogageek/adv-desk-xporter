package translate

import (
	"encoding/json"
	"fmt"
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
	var Langs []string
	for _, v := range data.TranslationLangsData {
		Langs = append(Langs, v.Lang)
	}
	fmt.Println("langs:", Langs)

	groupData := data.GroupData
	for _, v := range groupData {
		//--->
		names := v.Names
		desps := v.Descriptions

		for _, lang := range Langs {
			var name, desp string
			for _, v := range names {
				if v.Lang == lang {
					name = v.Text
				}
			}
			for _, v := range desps {
				if v.Lang == lang {
					desp = v.Text
				}
			}
			gql.TranslateGroup(v.Id, name, lang, desp)
		}
		//<---

		// for i, name := range v.Names {
		// 	if i > len(v.Descriptions)-1 {
		// 		break
		// 	}
		// 	gql.TranslateGroup(v.Id, name.Text, name.Lang, v.Descriptions[i].Text)
		// }
	}

	machineData := data.MachineData
	for _, v := range machineData {
		//--->
		names := v.Names
		desps := v.Descriptions
		urls := v.ImageUrls

		for _, lang := range Langs {
			var name, desp, url string
			for _, v := range names {
				if v.Lang == lang {
					name = v.Text
				}
			}
			for _, v := range desps {
				if v.Lang == lang {
					desp = v.Text
				}
			}
			for _, v := range urls {
				if v.Lang == lang {
					url = v.Text
				}
			}
			gql.TranslateMachine(v.Id, name, lang, desp, url)
		}
		//<---

		// for i, name := range v.Names {
		// 	if i > len(v.Descriptions)-1 || i > len(v.ImageUrls)-1 {
		// 		break
		// 	}
		// 	gql.TranslateMachine(v.Id, name.Text, name.Lang, v.Descriptions[i].Text, v.ImageUrls[i].Text)
		// }
	}

	machineStatusData := data.MachineStatusData
	func() { // fix 忽略預設machine status
		var newMachineStatusDatas []*fn.MachineStatusData
		for _, v := range machineStatusData {
			if v.Index >= 5000 { //1000~4000開頭為預設
				newMachineStatusDatas = append(newMachineStatusDatas, v)
			}
		}
		machineStatusData = newMachineStatusDatas
	}()
	for _, v := range machineStatusData {
		for _, name := range v.Names {
			gql.TranslateMachineStatus(v.Id, name.Text, name.Lang)
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

	parameterData := data.ParameterData
	for _, v := range parameterData {
		for _, vv := range v.Machine.Parameters.Nodes {
			for _, description := range vv.Descriptions {
				gql.TranslateParameter(*vv.Id, description.Lang, description.Text)
			}
		}
	}

	profileData := data.ProfileData
	for _, v := range profileData {
		//--->
		names := v.Names
		desps := v.Descriptions
		urls := v.ImageUrls

		for _, lang := range Langs {
			var name, desp, url string
			for _, v := range names {
				if v.Lang == lang {
					name = v.Text
				}
			}
			for _, v := range desps {
				if v.Lang == lang {
					desp = v.Text
				}
			}
			for _, v := range urls {
				if v.Lang == lang {
					url = v.Text
				}
			}
			gql.TranslateProfileMachine(string(v.Id), name, lang, desp, url)
		}
		//<---

		// for i, name := range v.Names {
		// 	if i > len(v.Descriptions)-1 || i > len(v.ImageUrls)-1 {
		// 		break
		// 	}
		// 	gql.TranslateProfileMachine(string(v.Id), name.Text, name.Lang, v.Descriptions[i].Text, v.ImageUrls[i].Text)
		// }

		for _, parameters := range v.Parameters {
			for _, description := range parameters.Descriptions {
				gql.TranslateProfileParameter(*parameters.Id, description.Lang, description.Text)
			}
		}
	}

	// defer func() {
	// 	// 可以取得 panic 的回傳值
	// 	r := recover()
	// 	if r != nil {
	// 		fmt.Println("Recovered in f", r)
	// 	}
	// }()
}
