package logic

import (
	model "porter/model/gqlclient"
	logic "porter/pkg/logic/client"
	"porter/util"
)

//gql query全部添加上多語言欄位  ok
func TestQuery() {
	//machine status OK
	func() {
		res := QueryMachineStatuses()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//group ok
	func() {
		res := QueryGroups()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//machines ok
	func() {
		res := QueryMachines()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//ParameterMappingCode ok
	func() {
		res := QueryParameterMappings()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	//ProfileMachine ok
	func() {
		res := QueryProfileMachines()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()

	// Parameter ok
	func() {
		res := QueryParameters("TWFjaGluZQ.YD86DsGunQAGpEaI", "")
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出
	}()
}

//取export出來的data, parse後，打 gql translate多語言欄位資料
func TestAdd() {

}

//這裡測試增加group的多語言(做到一半)
func TestMutation() {
	//init
	func() {
		// logic.PrepareGQLClientByAppSecret()
		logic.PrepareGQLCLient()
	}()

	//machine status
	machineStatus := func() {
		//query ok
		res := QueryMachineStatuses()
		util.PrintJson(res[0])
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出

		//mutation ok
		input := model.TranslateMachineStatusInput{
			Id:   "TVN0YXR1cw.YDXtHHkI5QAHox3q",
			Name: "Production_multilang",
			Lang: "en-US",
		}
		gqlQuery := model.TranslateMachineStatus
		Mutate(input, &gqlQuery)
	}

	//group
	group := func() {
		//query ok
		res := QueryGroups()
		util.PrintJson(res[3])
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出

		//mutation ok
		input := model.TranslateGroupInput{
			Id:          "R3JvdXA.YIFPI-3BRgAGL3vX",
			Name:        "測試群" + "_multilang",
			Lang:        "en-US",
			Description: "test description...",
		}
		gqlQuery := model.TranslateGroup
		Mutate(input, &gqlQuery)
	}

	//machines
	machines := func() {
		// query ok
		res := QueryMachines()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出

		// mutation ok
		input := model.TranslateMachineInput{
			Id:          "TWFjaGluZQ.YIFPJO3BRgAGL3vc",
			Name:        "aaa" + "_multilang",
			Lang:        "en-US",
			Description: "test description...",
			ImageUrl:    "",
		}
		gqlQuery := model.TranslateMachine
		Mutate(input, &gqlQuery)
	}

	//ParameterMappingCode ok
	parameterMappingCode := func() {
		// query ok
		res := QueryParameterMappings()
		util.PrintJson(res)
		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出

		// mutation ok
		input := model.TranslateParameterMappingCodeInput{
			Id:      "UE1hcENvZGU.YIFPI-3BRgAGL3vR",
			Lang:    "en-US",
			Message: "測試msg2" + "_multilang",
		}
		gqlQuery := model.TranslateParameterMappingCode
		Mutate(input, &gqlQuery)

	}

	//ProfileMachine ok
	profileMachine := func() {
		// query ok
		res := QueryProfileMachines()
		util.PrintJson(res)

		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出

		// mutation ok
		input := model.TranslateMachineInput{
			Id:          "UGZNYWNoaW5l.YIFPI-3BRgAGL3vT",
			Name:        "測試profile" + "_multilang",
			Lang:        "en-US",
			Description: "test description...",
			ImageUrl:    "",
		}
		gqlQuery := model.TranslateProfileMachine
		Mutate(input, &gqlQuery)
	}

	// Parameter ok
	profileParameter := func() {
		// query ok
		res := QueryProfileMachines() //ProfileParameter的資料要去查ProfileMachines!!!
		util.PrintJson(res)

		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出

		// mutation ok
		input := model.TranslateParameterInput{
			Id:          "UGZQYXJhbWV0ZXI.YIFPI-3BRgAGL3vU",
			Lang:        "en-US",
			Description: "test description...",
		}
		gqlQuery := model.TranslateProfileParameter
		Mutate(input, &gqlQuery)

		//profileParameter只有description支援多語言, 不支援name!!!
	}

	parameter := func() {
		// query ok
		res := QueryParameters("TWFjaGluZQ.YMQh1u3BRgAGL4Uv", "") //放query machine的Id
		util.PrintJson(res)                                       //如果查不出代表machine底下沒綁parameters

		//確認有查到多語言資料後
		//再去改GetSource() interface{} 把查到的資料拿出

		// mutation ok
		input := model.TranslateParameterInput{
			Id: "UGFyYW1ldGVy.YIFPM-3BRgAGL3xA",
			/*"Machine": {
			  "Parameters": {
			    "Nodes": [
			      {
			        "id": "UGFyYW1ldGVy.YIFPJe3BRgAGL3vd", //用這個id
			*/
			Lang:        "en-US",
			Description: "test description...",
		}
		gqlQuery := model.TranslateParameter
		Mutate(input, &gqlQuery)
	}

	machineStatus()        //1
	group()                //2
	machines()             //3
	parameter()            //4 (machine的parameter
	parameterMappingCode() //5
	profileMachine()       //6
	profileParameter()     //7

	//return
}

func TranslateMachineStatus(id, name, lang string) {
	input := model.TranslateMachineStatusInput{
		Id: id, Name: name, Lang: lang,
	}
	gqlQuery := model.TranslateMachineStatus
	Mutate(input, &gqlQuery)
}

func TranslateGroup(id, name, lang, desc string) {
	input := model.TranslateGroupInput{
		Id:          id,
		Name:        name,
		Lang:        lang,
		Description: desc,
	}
	gqlQuery := model.TranslateGroup
	Mutate(input, &gqlQuery)
}

func TranslateMachine(id, name, lang, desc, imageUrl string) {
	input := model.TranslateMachineInput{
		Id:          id,
		Name:        name,
		Lang:        lang,
		Description: desc,
		ImageUrl:    imageUrl,
	}
	gqlQuery := model.TranslateMachine
	Mutate(input, &gqlQuery)
}

func TranslateParameterMappingCode(id, lang, message string) {
	input := model.TranslateParameterMappingCodeInput{
		Id:      id,
		Lang:    lang,
		Message: message,
	}
	gqlQuery := model.TranslateParameterMappingCode
	Mutate(input, &gqlQuery)
}

func TranslateParameter(id, lang, desc string) {
	input := model.TranslateParameterInput{
		Id:          id, //UGFyYW1ldGVy.YMJ2pu3BRgAGL4Lc
		Lang:        lang,
		Description: desc,
	}
	// gqlQuery := model.TranslateProfileParameter
	//# fix bug
	gqlQuery := model.TranslateParameter
	Mutate(input, &gqlQuery)
}

func TranslateProfileMachine(id, name, lang, desc, imageUrl string) {
	input := model.TranslateMachineInput{
		Id:          id,
		Name:        name,
		Lang:        lang,
		Description: desc,
		ImageUrl:    imageUrl,
	}
	gqlQuery := model.TranslateProfileMachine
	Mutate(input, &gqlQuery)
}

func TranslateProfileParameter(id, lang, desc string) {
	input := model.TranslateParameterInput{
		Id:          id,
		Lang:        lang,
		Description: desc,
	}
	gqlQuery := model.TranslateProfileParameter
	Mutate(input, &gqlQuery)
}
