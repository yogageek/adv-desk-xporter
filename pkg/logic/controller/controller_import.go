package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"porter/pkg/logic/controller/translate"
	. "porter/pkg/logic/fn"
	gochan "porter/pkg/logic/gochan"
	vars "porter/pkg/logic/vars"

	// . "porter/util"
	"porter/util"

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

//偵測匯入進度
func syncDoImport() {
	vars.ResetPublicRess() //重設ws response
	vars.ChanDone = false  //重設是否完成狀態
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

func CheckBeforeImport() error {
	data, err := readFile() //read data
	if err != nil {
		return fmt.Errorf("file error: %s", err.Error())
	}
	//檢查目標預設語言是否在匯出的清單內
	var targetDefaultLang string //system default lang
	targetTranslation := GetSourceTranslations()
	b := util.IfaceToJson(targetTranslation)
	m := util.JsonAryToMap(b)
	for _, v := range m {
		if v["isDefault"].(bool) == true {
			targetDefaultLang = v["lang"].(string)
		}
	}

	fileTranslationLangsData := data.TranslationLangsData
	for _, v := range fileTranslationLangsData {
		if v.Lang == targetDefaultLang {
			//-------暫時增加檢查json檔預設語言是否等於目標預設語言
			if !v.IsDefault {
				return fmt.Errorf("system default lang %s not equal to default lang in file", targetDefaultLang)
			}
			//-------
			return nil
		}
	}
	return fmt.Errorf("default lang %s not included in file", targetDefaultLang)
}

func importController() {
	data, _ := readFile() //read data

	processes := implIface()
	updatePuclicResTotal(&data, processes)
	doImport(&data, processes)
	//new
	translate.Translate(&data)
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

func readFile() (data JsonData, err error) {
	// checkFilePath()

	//step1 Read json file...
	b, err := ioutil.ReadFile("./importingData.json")
	if err != nil {
		glog.Errorln(err)
		return data, err
	}
	// fmt.Printf("%s", b)

	// step2 Convert []byte to struct
	// method1
	// result := gjson.GetBytes(b, "machineStatusData") //get all values which key is "id" in a array
	// machineStatusDataB := []byte(result.Raw)
	// method2
	err = json.Unmarshal(b, &data)
	if err != nil {
		//handle bom format
		cb := Clean(b)
		err = json.Unmarshal(cb, &data)
		if err != nil {
			glog.Errorln(err)
			return data, err
		}
	}

	// debugging
	// util.PrintJson(data)
	// fmt.Println(data)

	return data, nil
}
