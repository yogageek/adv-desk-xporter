package logic

import (
	. "porter/pkg/logic/fn"
	. "porter/pkg/logic/vars"
)

//Processer 的 processes包含多種要匯入的類型的資料
type Processer interface {
	Process(jsonData *JsonData)
	GetName() string
	GetTotal(jsonData *JsonData) int
	GetSource() interface{}
	//要在fnimport_iface_impl裡為每個type新增GetTotal方法
}

type machineStatus struct {
	Counter
}

type mappingRule struct {
	Counter
}

type profileMachine struct {
	Counter
}

type groups struct {
	Counter
}

type machines struct {
	Counter
}

type parameters struct {
	Counter
}
