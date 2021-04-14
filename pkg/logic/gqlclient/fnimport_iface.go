package logic

type State int
type Mode string

const (
	StateDoing State = 1
	StateDone  State = 0
	ModeImport Mode  = "import"
	ModeExport Mode  = "export"
)

//返回目前是否為可執行狀態
func StateIsAvailable() bool {
	return Res.State == StateDoing
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
		Total: 1, //#暫時寫死0
	}
}

//Processer 的 processes包含多種要匯入的類型的資料
type Processer interface {
	Process(jsonData *jsonData)
	GetName() string
	//GetTotal() int
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
