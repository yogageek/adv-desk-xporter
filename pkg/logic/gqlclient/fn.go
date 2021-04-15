package logic

import (
	"encoding/json"
	"io/ioutil"
	model "porter/model/gqlclient"
)

//匯出的結構

type JsonData struct {
	MachineStatusData []*machineStatusData
	MappingRuleData   []*mappingRuleData
	ProfileData       []*profileData
	GroupData         []*groupData
	MachineData       []*machineData
	ParameterData     []*parameterData
	TranslationLangs  []*translationLangs
}

//未來改成從model拿原始input
type machineStatusData struct {
	Id          string
	NewId       string //改直接放Id就好(錯)
	ParentId    string
	ParentIndex int //注意這裡形態要對否則會拿到空值
	Name        string
	Index       int
	Color       string
	Depth       int
}

type mappingRuleData struct {
	Id     string
	NewId  string
	Name   string
	PType  string
	Detail []Detail
}

type Detail struct { //這裡的大小寫會影響到json大小寫
	Code     string
	Message  string
	StatusId string
	Lang     string
	Text     string
}

// 匯入profile第二步時候 需要帶profileMachineId 但這個id會在匯入profile第一步的時候才會產生 因此要在第一步時抓下來
// 所以我們可以直接忽略export導出的machineId
type profileData struct {
	model.ProfileMachine
}

type groupData struct {
	model.Groups
}

type machineData struct {
	model.Machines
}

type parameterData struct {
	model.QueryParametersOb
}

type translationLangs struct {
	model.TranslationLangs
}

func appendJson(keyName []string, data []interface{}) []byte {

	m := map[string]interface{}{}

	if len(keyName) == len(data) {
		for i := 0; i < len(keyName); i++ {
			m[keyName[i]] = data[i]
		}
	}

	b, _ := json.MarshalIndent(m, "", " ")
	return b
}

func listAll(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, fi := range files {
		if fi.IsDir() {
			//listAll(path + "/" + fi.Name())
			println(path + "/" + fi.Name())
		} else {
			println(path + "/" + fi.Name())
		}
	}
}
