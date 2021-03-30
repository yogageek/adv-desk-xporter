package logic

type State int
type Mode string

const (
	StateDoing State = 1
	StateDone  State = 0
	modeImport Mode  = "import"
	modeExport Mode  = "export"
)

func StateIsAvailable() bool {
	return Res.State == StateDone
}

var Res Response

type Response struct {
	Mode    Mode     `json:"mode"`  //import,export
	State   State    `json:"state"` //1,0
	Details []detail `json:"details"`
}

type detail struct {
	Name string `json:"name"`
	counter
}

type counter struct {
	Count int `json:"count"`
	Total int `json:"total"`
}

func NewCounter() counter {
	return counter{
		Count: 0,
		Total: 1,
	}
}

type Processer interface {
	Process(jsonData *jsonData)
	GetName() string
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
