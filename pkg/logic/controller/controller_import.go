package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "porter/pkg/logic/fn"
	gochan "porter/pkg/logic/gochan"
	vars "porter/pkg/logic/vars"
	. "porter/util"
	"sync"

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
	mutex := sync.Mutex{}
	mutex.Lock() //似乎不一定需要
	go func() {
		rs := []vars.Response{}
		for {
			if gochan.ChannelOut() {
				details := []vars.ResponseDetail{}
				publicDetails := vars.GetResponse().Details
				for _, v := range publicDetails {
					detail := vars.ResponseDetail{
						Name:    v.Name,
						Counter: v.Counter,
					}
					details = append(details, detail)
				}
				r := vars.Response{
					Details: details,
				}
				rs = append(rs, r)
			} else {
				break
			}
		}
		PrintJson(rs)
		fmt.Println(len(rs))
	}()

	importController()
	mutex.Unlock()
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

		details := vars.ResponseDetail{
			Name: name,
			Counter: vars.Counter{
				Total: total,
				Count: 0,
			},
		}
		vars.Res.Details = append(vars.Res.Details, details)
	}
	PrintJson(vars.Res)
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
