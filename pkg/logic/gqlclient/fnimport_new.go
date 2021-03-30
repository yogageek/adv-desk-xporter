package logic

func ToDoProcess() {
	var a machineStatus
	var b mappingRule
	var c profileMachine
	var d groups
	var e machines
	var f parameters

	processes := []Processer{a, b, c, d, e, f}

	func() {
		Res = Response{}
		Res.State = StateDoing
		Res.Mode = modeImport
	}()

	//init response info
	PrepareDetailTotal(processes)
	//read data
	data := ReadFile()

	//set DefaultLang for ImportMappingRule
	SetDefaultLang()

	//import data
	ProcessData(&data, processes)

	defer func() {
		Res.State = StateDone
	}()
}

func PrepareDetailTotal(processes []Processer) {
	//處理detail total分母
	for i := 0; i < len(processes); i++ {
		detail := detail{
			Name:    processes[i].GetName(),
			counter: NewCounter(),
		}
		Res.Details = append(Res.Details, detail)
	}
	// util.PrintJson(Res)
}

func ProcessData(data *jsonData, processes []Processer) {
	for i := 0; i < len(processes); i++ {
		processes[i].Process(data)
		Res.Details[i].Count = 1 //把一大類當做一
	}
	// util.PrintJson(Res)
}
