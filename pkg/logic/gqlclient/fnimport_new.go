package logic

func ImplIface() []Processer {
	var (
		a machineStatus
		b mappingRule
		c profileMachine
		d groups
		e machines
		f parameters
	)
	processes := []Processer{a, b, c, d, e, f}
	return processes
}

func ImportController() {
	data := ReadFile() //read data
	processes := ImplIface()
	makeResponse(&data, processes)
	doImport(&data, processes)
}

func doImport(data *JsonData, processes []Processer) {
	//set DefaultLang for ImportMappingRule
	SetDefaultLang()

	for i := 0; i < len(processes); i++ {
		processes[i].Process(data)
		//testing
		if i == 4 {
			break
		}
	}
	// util.PrintJson(Res)
}
