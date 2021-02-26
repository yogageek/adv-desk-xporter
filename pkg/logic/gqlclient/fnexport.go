package logic

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

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

	b := appendJson(ss, ii)
	// fmt.Printf("%s", b)

	//output the json file
	writeFile(b)
}

func writeFile(b []byte) {
	// 查看當前目錄
	// listAll(".")

	//method1
	_ = ioutil.WriteFile("exportingData.json", b, 0644)

	//method2
	//不知道怎麼避掉\n問題
	file, _ := os.OpenFile("exportingData_encode.json", os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err := encoder.Encode(b)
	if err != nil {
		glog.Error(err)
	}
}
