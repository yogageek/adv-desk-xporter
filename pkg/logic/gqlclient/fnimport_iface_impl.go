package logic

import (
	"bytes"
	"encoding/json"
	. "porter/pkg/logic/var"
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

//------------------------------------------------

func (o machineStatus) GetName() string {
	return MachineStatus
}

func (o mappingRule) GetName() string {
	return MappingRule
}

func (o profileMachine) GetName() string {
	return Profile
}

func (o groups) GetName() string {
	return Group
}

func (o machines) GetName() string {
	return Machine
}

func (o parameters) GetName() string {
	return Parameter
}

//------------------------------------------------
func (o machineStatus) GetTotal(jsonData *JsonData) int {
	return GetTotalMachineStatus(jsonData)
}

func (o mappingRule) GetTotal(jsonData *JsonData) int {
	return GetTotalMappineRule(jsonData)
}

func (o profileMachine) GetTotal(jsonData *JsonData) int {
	return GetTotalProfile(jsonData)
}

func (o groups) GetTotal(jsonData *JsonData) int {
	return GetTotalGroup(jsonData)
}

func (o machines) GetTotal(jsonData *JsonData) int {
	return GetTotalMachine(jsonData)
}

func (o parameters) GetTotal(jsonData *JsonData) int {
	return GetTotalParameter(jsonData)
}

// func GetCounter(data *JsonData) counter {
// 	var total int
// 	for _, v := range data.ParameterData {
// 		m := v.Machine.Parameters.Nodes
// 		total = total + len(m)
// 	}

// 	c := 0
// 	return counter{
// 		Count: &c,
// 		Total: total,
// 	}
// }
