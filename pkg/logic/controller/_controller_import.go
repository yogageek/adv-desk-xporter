package logic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	. "porter/pkg/logic/fn"
)

//先做完group->machine->parameter再做
//增加防呆功能
//先去取得系統預設語言
//檢查要匯入的內容是否滿足
//ok之後再匯入

func Import() {
	data := ReadFile()

	//goroutine status
	// chTotal := make(chan int)
	// chCount := make(chan int)

	// chTotal <- 1
	// chCount <- 1

	//business logic
	ImportMachineStatus(&data) //ok

	// util.PrintJson(data)
	b, _ := json.MarshalIndent(data, "", " ")
	m := map[string]string{} //儲存新舊id對應關係
	func() {
		for _, v := range data.MachineStatusData {
			if m[v.Id] == "" && v.NewId != "" {
				m[v.Id] = v.NewId
			}
		}
	}()
	for k, v := range m { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	//把舊&data取代掉
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
	m = map[string]string{} //儲存新舊id對應關係
	func() {
		for _, v := range data.MappingRuleData {
			if m[v.Id] == "" && v.NewId != "" {
				m[v.Id] = v.NewId
			}
		}
	}()
	for k, v := range m { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}

	ImportProfileMachine(&data) //ok

	//groups---
	idMap := ImportGroups(&data)
	for k, v := range idMap { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	json.Unmarshal(b, &data) //將替換後的資料b 賦予給&data
	//machines---
	idMap = ImportMachines(&data)
	for k, v := range idMap { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	json.Unmarshal(b, &data) //將替換後的資料b 賦予給&data
	//parameters---
	ImportParameters(&data)

	//debugging------>
	_ = ioutil.WriteFile("importingData.json", b, 0644)
	// util.PrintJson(data)
	//<--------------

}
