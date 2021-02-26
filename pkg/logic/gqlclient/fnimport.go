package logic

import (
	"encoding/json"
	"io/ioutil"

	"github.com/golang/glog"
)

func Import() {
	data := ReadFile()

	//business logic
	ImportMachineStatus(data.MachineStatusData) //ok
	ImportMappingRule(data.MappingRuleData)     //ok
	ImportProfileMachine(data.ProfileData)      //ok

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
