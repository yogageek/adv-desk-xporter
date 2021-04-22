package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "porter/pkg/logic/fn"
	gochan "porter/pkg/logic/gochan"
	vars "porter/pkg/logic/vars"

	// . "porter/util"

	"github.com/golang/glog"
)

func Import() {
	go syncDoImport()
	importController()
	// PrintJson(vars.Get_PublicRess())
	// mutex := sync.Mutex{} //似乎不一定需要，尚未驗證
	// mutex.Lock()
	// mutex.Unlock()

}

func syncDoImport() {
	vars.ResetPublicRess()
	vars.ChanDone = false
	for {
		if gochan.ChannelOut() {
			vars.AppendResToRess()
		} else {
			vars.ChanDone = true
			fmt.Println("channel out done, break for loop, length:", len(vars.PublicRess))
			break
		}
		// log.Info("channel out done, length:", len(vars.PublicRess))
	}
}

func importController() {
	data := readFile() //read data
	processes := implIface()
	updatePuclicResTotal(&data, processes)
	doImport(&data, processes)
}

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

func updatePuclicResTotal(data *JsonData, processes []Processer) {
	//處理detail total分母
	for i := 0; i < len(processes); i++ {
		name := processes[i].GetName()
		total := processes[i].GetTotal(data)
		vars.Update_PuclicRes_Detail_Total(name, total)
	}
	// PrintJson(vars.Res)
}

func doImport(data *JsonData, processes []Processer) {
	for i := 0; i < len(processes); i++ {
		processes[i].Process(data)
	}
}

func readFile() JsonData {

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
