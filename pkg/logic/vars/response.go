package logic

type State int
type Mode string

const (
	StateDoing State = 1
	StateDone  State = 0
	ModeImport Mode  = "import"
	ModeExport Mode  = "export"
)

var Res Response

//返回目前是否為可執行狀態
func StateIsAvailable() bool {
	return Res.State == StateDoing
}

type Response struct {
	Mode    Mode              `json:"mode,omitempty"`  //import,export
	State   State             `json:"state,omitempty"` //1,0
	Details []*ResponseDetail `json:"details,omitempty"`
}

type ResponseDetail struct {
	Name string `json:"name"`
	Counter
}

type Counter struct {
	Count int `json:"loaded"`
	Total int `json:"total"`
}

// func NewCounter(c, t int) counter {
// 	return counter{
// 		Count: &c,
// 		Total: t, //#暫時寫死0
// 	}
// }

//初始化response資料
func SetResponseDoing(mode Mode) {
	func() {
		Res = Response{}
		Res.State = StateDoing
		Res.Mode = mode //modeImport,modeExport
	}()
}

//更新response資料為結束狀態
func SetResponseDone() {
	func() {
		Res.State = StateDone
	}()
}

func SetResponseCount(s string, count int) {
	for _, v := range Res.Details {
		if s == v.Name {
			if v.Count < count { //channel順序不定
				v.Count = count
			}
		}
	}
	// util.PrintJson(Res)
}
