package logic

import (
	"io/ioutil"
	"porter/util"
)

func Export() {
	ss := []string{"url", "machineStatusData", "mappingRuleData", "profileData", "groupData", "machineData", "parameterData"}
	ii := []interface{}{}

	ii = append(ii, IFP_URL)
	ii = append(ii, getSourceMachineStatus())
	// goto debugging
	ii = append(ii, getSourceMappingRule())
	// goto debugging
	ii = append(ii, getSourceProfileMachines())
	// goto debugging

	ii = append(ii, getSourceGroups())
	// goto debugging

	machineData := getSourceMachines()
	ii = append(ii, machineData)
	// goto debugging

	parameterData := getSourceParameters(getMachineIds(machineData))
	ii = append(ii, parameterData)
	// goto debugging

	// debugging:
	// b, _ := json.MarshalIndent(ii, "", " ")
	// fmt.Printf("%s", b)

	b := appendJson(ss, ii)
	util.PrintCyan(b)

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
	// file, _ := os.OpenFile("exportingData_encode.json", os.O_CREATE, os.ModePerm)
	// defer file.Close()
	// encoder := json.NewEncoder(file)
	// err := encoder.Encode(b)
	// if err != nil {
	// 	glog.Error(err)
	// }
}
