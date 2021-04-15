package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	logic "porter/pkg/logic/gqlclient"
	"porter/util"
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

type wsError struct {
	Event string `json:"event"`
	Error string `json:"error"`
	State int    `json:"state"`
}

// TextMessage denotes a text data message. The text message payload is
// interpreted as UTF-8 encoded text data.
const textMsg = 1

func ProcessWs(ws *websocket.Conn) {

	ws.WriteMessage(textMsg, []byte("connected"))

	//setp1 檢查是否正在做
	if logic.StateIsAvailable() {
		r := wsError{
			Event: "event-fail",
			Error: "still in progress",
			State: int(logic.StateDoing),
		}
		err := ws.WriteJSON(r)
		if err != nil {
			log.Println("write:", err)
		}
		return
	}

	//setp2 寫入新事件
	//#這裡用channel來做 後端只要有發 這裡就一直取出來 直到取完為指
	var (
		err error
		res *logic.Response
	)

	sendWs := func() {
		r := logic.Res
		msg, _ := json.MarshalIndent(r, "", " ")
		err = ws.WriteMessage(textMsg, msg)
		if err != nil {
			log.Println("write:", err)
		}
	}

	func() {
		for {
			res := logic.Res
			if len(res.Details) > 0 {
				res = logic.Res
				break
			}
			time.Sleep(time.Second * 1)
		}
		fmt.Println("ws--------------------")
		util.PrintJson(res)
	}()

	for {
		logic.ChannelGetCount2() //如果匯入資料送完 這裡取完 會停在這行  only mStatus
		sendWs()
		// testing write ws
		// s := strconv.Itoa(i)
		// ws.WriteMessage(1, []byte(s))

		// msg, _ := json.MarshalIndent(res, "", " ")
		// err = ws.WriteMessage(textMsg, msg)
		// if err != nil {
		// 	log.Println("write:", err)
		// }
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
