package v1

import (
	"fmt"
	"log"
	"net/http"

	gochan "porter/pkg/logic/gochan"
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
	Event_FAIL      = "event-fail"
	Event_DONE      = "event-done"
)

func ProcessWs(ws *websocket.Conn) {

	if err := ws.WriteJSON(
		wsResponse{
			Event: Event_CONNECTED,
		},
	); err != nil {
		log.Println("write err:", err)
	}

	//setp1 檢查是否正在做
	if vars.StateIsAvailable() {
		if err := ws.WriteJSON(
			wsResponse{
				Event: Event_FAIL,
				Error: "in process",
			},
		); err != nil {
			log.Println("write err:", err)
		}
		return
	}

	//step temp 檢查後台response(包含total數)是否已建好
	for {
		if len(vars.Res.Details) > 0 {
			break
		}
		time.Sleep(time.Second * 1)
	}
	log.Println("Already Init Response...")

	//setp2 寫入新事件
	//#這裡用channel來做 後端只要有發 這裡就一直取出來 直到取完為指

	var closeMaxCount int
	for {
		b := gochan.ChannelOut() //如果匯入資料送完 這裡取完 會停在這行  only mStatus

		sendWs := func() {
			// msg, _ := json.MarshalIndent(logic.Res, "", " ")
			if err := ws.WriteJSON(
				wsResponse{
					Event:    Event_PROCESS,
					Response: vars.Res,
				},
			); err != nil {
				log.Println("write err:", err)
			}
		}

		// if ChannelOut()==true 代表有從channel取出值
		if b {
			sendWs()
		} else {
			closeMaxCount++
			fmt.Println(closeMaxCount)
			if closeMaxCount >= 1 {
				if err := ws.WriteJSON(
					wsResponse{
						Event: Event_DONE,
					},
				); err != nil {
					log.Println("write:", err)
				}
			}
			vars.Res.State = vars.StateDone
			return //return 等於斷開連結(不可!)
		}
		// testing write ws
		// s := strconv.Itoa(i)
		// ws.WriteMessage(1, []byte(s))

	}
}

/*
修改socekt邏輯

基本上用推送的方式就是事件進度有更新的時候發送，而不是要前端主動定期去要，這樣本質上又變成週期輪詢了
而且前端可能會是多個
後端也不一定要固定周期發送，是「進度有更新發送」

{
  type: 'event-name',
  data: {
    ...
  },
}

啟動連線->返回連上成功/連上失敗
{
	type:"eventStart"
}
過程->後台會一直更新狀態值,有更新則寫回ws
{
	type:"eventProcess"
	value...
}
沒更新就不用再繼續回復
*/
