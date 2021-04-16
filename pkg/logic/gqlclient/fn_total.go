package logic

// . "porter/util"

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
