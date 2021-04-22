package logic

import "time"

type State int
type Mode string

const (
	StateDoing State = 1
	StateDone  State = 0
	ModeImport Mode  = "import"
	ModeExport Mode  = "export"
)

//----------------------

var PublicRes Response
var PublicRess []Response
var ChanDone bool

type Response struct {
	Mode     Mode      `json:"mode,omitempty"`  //import,export
	State    State     `json:"state,omitempty"` //1,0
	Rdetails []Rdetail `json:"details,omitempty"`
}

type Rdetail struct {
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

func Update_PublicRes_Start(mode Mode) {
	PublicRes = Response{}
	PublicRes.State = StateDoing
	PublicRes.Mode = mode //modeImport,modeExport
}

//更新response資料為結束狀態
func Update_PublicRes_Done() {
	time.Sleep(2 * time.Second)
	PublicRes.State = StateDone
	PublicRes.Rdetails = nil //這裡清掉後，ws會拿不到資料，故要sleep
}

func Get_PublicRes() Response {
	return PublicRes
}

func Get_PublicRess() []Response {
	return PublicRess
}

//是否目前為可執行狀態
func Get_PublicRes_State() bool {
	return PublicRes.State == StateDone
}

//是否detail已準備好
func Get_PubliceRes_Detail_Prepared() bool {
	return len(PublicRes.Rdetails) > 0
}

func Update_PuclicRes_Detail_Total(name string, total int) {
	d := Rdetail{
		Name: name,
		Counter: Counter{
			Total: total,
			Count: 0,
		},
	}
	PublicRes.Rdetails = append(PublicRes.Rdetails, d)
}

func Update_PuclicRes_Detail(name string, count int) {
	for i, v := range PublicRes.Rdetails {
		if v.Name == name {
			if v.Count < count { //channel順序不定
				PublicRes.Rdetails[i].Count = count
			}
		}
	}
}

func ResetPublicRess() {
	PublicRess = []Response{}
}

func ResetPublicRes() {
	PublicRes = Response{}
}

func AppendResToRess() {
	details := []Rdetail{}

	for _, v := range PublicRes.Rdetails {
		detail := Rdetail{
			Name:    v.Name,
			Counter: v.Counter,
		}
		details = append(details, detail)
	}

	r := Response{
		Rdetails: details,
	}

	PublicRess = append(PublicRess, r)
}

func SumLoaded(r Response) int {
	loaded := 0
	for _, v := range r.Rdetails {
		loaded += v.Count
	}
	return loaded
}
