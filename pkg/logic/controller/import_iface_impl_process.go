package logic

import (
	"bytes"
	"encoding/json"
	. "porter/pkg/logic/fn"
)

func (o machineStatus) Process(data *JsonData) {
	ImportMachineStatus(data)

	b, _ := json.MarshalIndent(data, "", " ")
	m := map[string]string{} //儲存新舊id對應關係
	func() {
		for _, v := range data.MachineStatusData {
			if m[v.Id] == "" && v.NewId != "" {
				m[v.Id] = v.NewId
			}
		}
	}()
	for k, v := range m { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	//把舊&data取代掉
	json.Unmarshal(b, &data)
}

func (o mappingRule) Process(data *JsonData) {
	ImportMappingRule(data)

	b, _ := json.MarshalIndent(data, "", " ")

	m := map[string]string{} //儲存新舊id對應關係
	func() {
		for _, v := range data.MappingRuleData {
			if m[v.Id] == "" && v.NewId != "" {
				m[v.Id] = v.NewId
			}
			for _, v := range v.Detail {
				if m[v.Id] == "" && v.NewId != "" {
					m[v.Id] = v.NewId
				}
			}
		}
	}()

	for k, v := range m { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}

	json.Unmarshal(b, data)
}

func (o profileMachine) Process(data *JsonData) {
	ImportProfileMachine(data)
}

func (o groups) Process(data *JsonData) {
	idMap := ImportGroups(data)

	b, _ := json.MarshalIndent(data, "", " ")
	for k, v := range idMap { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	json.Unmarshal(b, &data) //將替換後的資料b 賦予給&data
}

func (o machines) Process(data *JsonData) {
	idMap := ImportMachines(data)

	b, _ := json.MarshalIndent(data, "", " ")
	for k, v := range idMap { //用新id取代整個file中的舊id
		b = bytes.ReplaceAll(b, []byte(k), []byte(v))
	}
	json.Unmarshal(b, data) //將替換後的資料b 賦予給&data
}

func (o parameters) Process(data *JsonData) {
	ImportParameters(data)
}
