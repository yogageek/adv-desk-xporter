package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/golang/glog"
)

func Import() {
	data := ReadFile()

	//business logic
	ImportMachineStatus(&data) //ok

	// util.PrintJson(data)
	b, _ := json.MarshalIndent(data, "", " ")
	m := map[string]string{}
	func() {
		for _, v := range data.MachineStatusData {
			if m[v.Id] == "" && v.NewId != "" {
				m[v.Id] = v.NewId
			}
		}
	}()
	for k, v := range m {
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	json.Unmarshal(b, &data)

	// util.PrintJson(data)

	//ok但4000以下不支援 因為無法匯入 因此也無法拿到新id

	//--->
	//machine status id 已設為新的 ->錯誤
	//要把舊的取代為新的
	//舊得要保留才能知道新舊對應關係
	//<---

	//法一:把新舊存成一個map 最後再替換全部
	//.....比較簡單
	//法二:把新的存到newid欄位 當其他是拿id時要改拿newid(沒辦法,因為其他結構不包含新id,它是塞在前一個結構中)
	//....所以還是要把newid拿出來取代所有舊id

	ImportMappingRule(&data) //ok
	//重大bug: 如果兩邊default language不同 會無法匯入  因為匯入時強制要放lang
	// ERROR: logging before flag.Parse: E0301 02:38:36.074068    7368 gql_mapping_rule.go:43] Translation validation failed: lang: Validator failed for path `lang` with value `en`
	//接下來存新舊id讓profile能用
	//-->REPLACE old id with new id
	//待測試...
	b, _ = json.MarshalIndent(data, "", " ")
	m = map[string]string{}
	func() {
		for _, v := range data.MappingRuleData {
			if m[v.Id] == "" && v.NewId != "" {
				m[v.Id] = v.NewId
			}
		}
	}()
	for k, v := range m {
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}

	//debugging
	fmt.Println("new json before import profile-----------------------------------------")
	_ = ioutil.WriteFile("importingDataWithNewId.json", b, 0644)

	json.Unmarshal(b, &data)
	// util.PrintJson(data)

	ImportProfileMachine(&data) //ok

}

func ReadFile() jsonData {
	//step1 Read json file...
	b, err := ioutil.ReadFile("./importingData.json")
	if err != nil {
		glog.Fatal(err)
	}
	// fmt.Printf("%s", b)

	// step2 Convert []byte to struct
	// method1
	// result := gjson.GetBytes(b, "machineStatusData") //get all values which key is "id" in a array
	// machineStatusDataB := []byte(result.Raw)
	// method2
	var data jsonData
	err = json.Unmarshal(b, &data)
	if err != nil {
		glog.Fatal(err)
	}

	// debugging
	// util.PrintJson(data)
	// fmt.Println(data)

	return data
}
