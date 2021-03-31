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
	loadedParameters = 0 //testing put here
	ProcessData(&data, processes)
}

func PrepareDetailTotal(data *jsonData, processes []Processer) {
	//處理detail total分母
	for i := 0; i < len(processes); i++ {
		details := detail{
			Name:    processes[i].GetName(),
			counter: NewCounter(),
		}

		// parameter count 問題暫時沒解
		//for debugging
		// if processes[i].GetName() == "parameters" {
		// 	details = detail{
		// 		Name:    processes[i].GetName(),
		// 		counter: GetCounter(data),
		// 	}
		// }

		Res.Details = append(Res.Details, details)
	}
	// util.PrintJson(Res)
}

func ProcessData(data *jsonData, processes []Processer) {

	for i := 0; i < len(processes); i++ {

		// 會有bug
		// for debugging
		// if processes[i].GetName() == "parameters" {
		// 	go func() {
		// 		for {
		// 			Res.Details[i].Count = loadedParameters
		// 			// fmt.Println("ok")
		// 			time.Sleep(time.Second * 2)
		// 		}
		// 	}()
		// }

		processes[i].Process(data)

		//上面做完這裡才會加一(邏輯不對)
		Res.Details[i].Count = 1 //把一大類當做一

	}
	// util.PrintJson(Res)
}
