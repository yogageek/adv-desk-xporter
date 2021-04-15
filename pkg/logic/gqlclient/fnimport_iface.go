package logic

//Processer 的 processes包含多種要匯入的類型的資料
type Processer interface {
	Process(jsonData *JsonData)
	GetName() string
	GetTotal(jsonData *JsonData) int
	//要在fnimport_iface_impl裡為每個type新增GetTotal方法
}

type machineStatus struct {
	counter
}

type mappingRule struct {
	counter
}

type profileMachine struct {
	counter
}

type groups struct {
	counter
}

type machines struct {
	counter
}

type parameters struct {
	counter
}
