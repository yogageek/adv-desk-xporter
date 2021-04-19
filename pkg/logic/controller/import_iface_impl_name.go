package logic

import (
	. "porter/pkg/logic/vars"
)

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
