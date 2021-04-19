package logic

import (
	. "porter/pkg/logic/fn"
)

// . "porter/util"

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

//------------------------------------------------

func GetTotalMachineStatus(jsonData *JsonData) int {
	d := jsonData.MachineStatusData
	return len(d)
}

func GetTotalMappineRule(jsonData *JsonData) int {
	d := jsonData.MappingRuleData
	return len(d)
}

func GetTotalProfile(jsonData *JsonData) int {
	d := jsonData.ProfileData
	return len(d)
}

func GetTotalGroup(jsonData *JsonData) int {
	d := jsonData.GroupData
	return len(d)
}

func GetTotalMachine(jsonData *JsonData) int {
	d := jsonData.MachineData
	return len(d)
}

func GetTotalParameter(jsonData *JsonData) int {
	var sum int
	d := jsonData.ParameterData
	for _, v := range d {
		sum = sum + len(v.Machine.Parameters.Nodes)
	}
	return sum
}
