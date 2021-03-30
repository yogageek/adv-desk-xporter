package logic

func BeforeProcess(mode Mode) {
	func() {
		Res = Response{}
		Res.State = StateDoing
		Res.Mode = mode //modeImport,modeExport
	}()
}

func AfterProcess() {
	func() {
		Res.State = StateDone
	}()
}

func ToDoProcess() {
	var a machineStatus
	var b mappingRule
	var c profileMachine
	var d groups
	var e machines
	var f parameters

	processes := []Processer{a, b, c, d, e, f}

	//read data
	data := ReadFile()

	//init response info
	PrepareDetailTotal(&data, processes)

	//set DefaultLang for ImportMappingRule
	SetDefaultLang()

	//import data
	ProcessData(&data, processes)
}

func PrepareDetailTotal(data *jsonData, processes []Processer) {
	//處理detail total分母
	for i := 0; i < len(processes); i++ {
		details := detail{
			Name:    processes[i].GetName(),
			counter: NewCounter(),
		}

		//for debugging
		if processes[i].GetName() == "parameters" {
			details = detail{
				Name:    processes[i].GetName(),
				counter: GetCounter(data),
			}
		}

		Res.Details = append(Res.Details, details)
	}
	// util.PrintJson(Res)
}

func ProcessData(data *jsonData, processes []Processer) {
	for i := 0; i < len(processes); i++ {
		//for debugging
		if processes[i].GetName() == "parameters" {
			continue
		}
		processes[i].Process(data)
		Res.Details[i].Count = 1 //把一大類當做一
	}
	// util.PrintJson(Res)
}
