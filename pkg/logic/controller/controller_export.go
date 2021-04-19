package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	. "porter/pkg/logic/client"

	. "porter/pkg/logic/fn"
	. "porter/pkg/logic/vars"

	"porter/util"
)

func ExportController() {

	// processes := ImplIface()
	// MakeResponse(&data, processes)
	// doImport(&data, processes)

	// for i := 0; i < len(processes); i++ {
	// 	processes[i].
	// 	//testing
	// 	// if i == 4 {
	// 	// 	break
	// 	// }
	// }
	// util.PrintJson(Res)

}

func Export() {
	//注意這裡如果少加 整個json會錯誤且不易發現
	ss := []string{Url, MachineStatus, MappingRule, Profile, Group, Machine, Parameter, Translation}
	ii := []interface{}{}

	ii = append(ii, IFP_URL)
	ii = append(ii, GetSourceMachineStatus())
	// goto debugging
	ii = append(ii, GetSourceMappingRule())
	// goto debugging`
	ii = append(ii, GetSourceProfileMachines())
	// goto debugging

	ii = append(ii, GetSourceGroups())
	// goto debugging

	machineData := GetSourceMachines()
	ii = append(ii, machineData)
	// goto debugging
	parameterData := GetSourceParameters(GetMachineIds(machineData))
	ii = append(ii, parameterData)
	// goto debugging

	translationLangs := GetSourceTranslations()
	ii = append(ii, translationLangs)

	// debugging:
	// b, _ := json.MarshalIndent(ii, "", " ")
	// fmt.Printf("%s", b)

	b := AppendJson(ss, ii)
	util.PrintCyan(b)

	//output the json file
	writeFile(b)

	//-----------
	checkFilePath()
}

func writeFile(b []byte) {

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

func checkFilePath() {
	// 查看當前目錄
	fmt.Println("----------ListAll(.)-----------")
	ListAll(".")
	fmt.Println("----------ListAll(./)-----------")
	ListAll("./")

	fmt.Println("------------os.Executable()--------------")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	fmt.Println("------------ os.Getwd()--------------")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user

	fmt.Println("------------abs--------------")
	files, _ := ioutil.ReadDir(".")
	paths, _ := filepath.Abs(".")
	for _, file := range files {
		fmt.Println(filepath.Join(paths, file.Name()))
	}
}

func AppendJson(keyName []string, data []interface{}) []byte {

	m := map[string]interface{}{}

	if len(keyName) == len(data) {
		for i := 0; i < len(keyName); i++ {
			m[keyName[i]] = data[i]
		}
	}

	b, _ := json.MarshalIndent(m, "", " ")
	return b
}

func ListAll(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, fi := range files {
		if fi.IsDir() {
			//ListAll(path + "/" + fi.Name())
			println(path + "/" + fi.Name())
		} else {
			println(path + "/" + fi.Name())
		}
	}
}
