package logic

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

//之後要改成存檔並匯出一個檔案流
func setJson(keyName []string, data []interface{}) []byte {

	m := map[string]interface{}{}

	if len(keyName) == len(data) {
		for i := 0; i < len(keyName); i++ {
			m[keyName[i]] = data[i]
		}
	}

	b, _ := json.MarshalIndent(m, "", " ")
	return b
}

func Export() {
	ss := []string{"machineStatusData", "mappingRuleData", "profileData"}
	ii := []interface{}{}

	ii = append(ii, getSourceMachineStatus())
	// goto debugging
	ii = append(ii, getSourceMappingRule())
	// goto debugging
	ii = append(ii, getSourceProfileMachines())

	// debugging:
	// b, _ := json.MarshalIndent(ii, "", " ")
	// fmt.Printf("%s", b)

	b := setJson(ss, ii)
	fmt.Printf("%s", b)

	//...

	//output json file
}

//之後要改成讀取一個檔案流
func GetJson(b []byte) {
	// method1
	// result := gjson.GetBytes(b, "machineStatusData") //get all values which key is "id" in a array
	// machineStatusDataB := []byte(result.Raw)

	// method2
	var j jsonFile
	err := json.Unmarshal(b, j)
	if err != nil {
		glog.Error(err)
	}

	ImportMachineStatus(j.machineStatusDatas)
	ImportMappingRule(j.mappingRuleData)
	ImportProfileMachine(j.profileData)
}

//接著測試import功能

// 匯入profile第二步時候 需要帶profileMachineId 但這個id會在匯入profile第一步的時候才會產生 因此要在第一步時抓下來
// 所以我們可以直接忽略export導出的machineId
