package logic

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	. "porter/pkg/logic/var"
	"porter/util"
)

func Export() {
	//注意這裡如果少加 整個json會錯誤且不易發現
	ss := []string{Url, MachineStatus, MappingRule, Profile, Group, Machine, Parameter, Translation}
	ii := []interface{}{}

	ii = append(ii, IFP_URL)
	ii = append(ii, getSourceMachineStatus())
	// goto debugging
	ii = append(ii, getSourceMappingRule())
	// goto debugging`
	ii = append(ii, getSourceProfileMachines())
	// goto debugging

	ii = append(ii, GetSourceGroups())
	// goto debugging

	machineData := getSourceMachines()
	ii = append(ii, machineData)
	// goto debugging
	parameterData := getSourceParameters(getMachineIds(machineData))
	ii = append(ii, parameterData)
	// goto debugging

	translationLangs := getSourceTranslations()
	ii = append(ii, translationLangs)

	// debugging:
	// b, _ := json.MarshalIndent(ii, "", " ")
	// fmt.Printf("%s", b)

	b := appendJson(ss, ii)
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
	fmt.Println("----------listAll(.)-----------")
	listAll(".")
	fmt.Println("----------listAll(./)-----------")
	listAll("./")

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
