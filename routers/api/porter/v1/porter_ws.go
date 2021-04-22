package v1

import (
	"log"
	"net/http"

	// . "porter/util"

	vars "porter/pkg/logic/vars"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsEvent(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	ProcessWs(ws)
}

type wsResponse struct {
	Event string `json:"event"`
	Error string `json:"error,omitempty"`
	vars.Response
}

// TextMessage denotes a text data message. The text message payload is
// interpreted as UTF-8 encoded text data.
const textMsg = 1

const (
	Event_CONNECTED = "event-connected"
	Event_START     = "event-start"
	Event_PROCESS   = "event-process"
	Event_LOCKED    = "event-locked"
	Event_UNLOCKED  = "event-unlocked"
	Event_DONE      = "event-done"
)

func ProcessWs(ws *websocket.Conn) {
	defer ws.Close()
	//注意 return 等於斷開連結(server不應主動斷，client關閉網頁就等於斷線)

	if err := ws.WriteJSON(
		wsResponse{
			Event: Event_CONNECTED,
		},
	); err != nil {
		log.Println("write err:", err)
	}

	//setp1 檢查是否正在做

	sendEventLocked := func() {
		if err := ws.WriteJSON(
			wsResponse{
				Event: Event_LOCKED,
				Error: "work is in process",
			},
		); err != nil {
			log.Println("write err:", err)
		}
	}
	sendEventUnlocked := func() {
		if err := ws.WriteJSON(
			wsResponse{
				Event: Event_UNLOCKED,
			},
		); err != nil {
			log.Println("write err:", err)
		}
	}
	sendEventStart := func() {
		if err := ws.WriteJSON(
			wsResponse{
				Event: Event_START,
				// Response: vars.GetResponse(),
			},
		); err != nil {
			log.Println("write err:", err)
		}
	}
	snedEventProcess := func(r vars.Response) {
		// msg, _ := json.MarshalIndent(logic.Res, "", " ")
		if err := ws.WriteJSON(
			wsResponse{
				Event:    Event_PROCESS,
				Response: r,
			},
		); err != nil {
			log.Println("write err:", err)
		}
	}
	sendEventDone := func() {
		if err := ws.WriteJSON(
			wsResponse{
				Event: Event_DONE,
			},
		); err != nil {
			log.Println("write:", err)
		}
	}

	for {
		log.Println("check if STATE is available...")
		if !vars.Get_PublicRes_State() { //if false, means not available now
			sendEventLocked() //send event locked
			for {
				if !vars.Get_PublicRes_State() {
					time.Sleep(time.Second * 1)
				} else {
					sendEventUnlocked()
					break
				}
			}
		} else {
			break
		}
	}

	//檢查後台response(包含total數)是否已建好
	for {
		time.Sleep(time.Second * 1)
		log.Println("check if DETAIL is available...")
		if vars.Get_PubliceRes_Detail_Prepared() {
			sendEventStart() //發送事件開始前的通知
			log.Println("check if DETAIL is available... -> ok")
			break
		}
	}

	//setp2 寫入新事件
	//#這裡用channel來做 後端只要有發 這裡就一直取出來 直到取完為指

	var oldLoaded int
	// PrintJson(vars.PublicRess)
	for {
		time.Sleep(time.Second * 1)
		l := len(vars.PublicRess)
		// fmt.Println("PublicRess length so far:", l)
		v := vars.PublicRess[l-1]

		realtimeLoaded := vars.SumLoaded(v)

		if oldLoaded < realtimeLoaded {
			snedEventProcess(v)
			oldLoaded = realtimeLoaded
		} else {
			if vars.ChanDone {
				sendEventDone()
				return
			}
		}
	}
}
