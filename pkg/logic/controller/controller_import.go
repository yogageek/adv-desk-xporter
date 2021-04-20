package logic

import (
	"encoding/json"
	"io/ioutil"
	. "porter/pkg/logic/fn"
	. "porter/pkg/logic/vars"
	. "porter/util"

	"github.com/golang/glog"
)

func implIface() []Processer {
	var (
		a machineStatus
		b mappingRule
		c profileMachine
		d groups
		e machines
		f parameters
	)
	processes := []Processer{a, b, c, d, e, f}
	return processes
}

func Import() {
	importController()
}

func importController() {
	data := ReadFile() //read data
	processes := implIface()
	makeResponse(&data, processes)
	doImport(&data, processes)
}

func doImport(data *JsonData, processes []Processer) {

	for i := 0; i < len(processes); i++ {
		processes[i].Process(data)
		//testing
		// if i == 4 {
		// 	break
		// }
	}
	// util.PrintJson(Res)
}

func makeResponse(data *JsonData, processes []Processer) {
	//處理detail total分母
	for i := 0; i < len(processes); i++ {
		name := processes[i].GetName()
		total := processes[i].GetTotal(data)

		details := ResponseDetail{
			Name: name,
			Counter: Counter{
				Total: total,
				Count: 0,
			},
		}
		Res.Details = append(Res.Details, &details)
	}
	PrintJson(Res)
}

func ReadFile() JsonData {

	// checkFilePath()

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
	var data JsonData
	err = json.Unmarshal(b, &data)
	if err != nil {
		glog.Fatal(err)
	}

	// debugging
	// util.PrintJson(data)
	// fmt.Println(data)

	return data
}
