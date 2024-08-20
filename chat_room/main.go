package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Message struct {
	Event   string `json:"event"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TimeStamp    time.Time `json:"timestamp"`
}

func NewMessage(event, name, content string) *Message {
	return &Message{
		Event:   event,
		Name:    name,
		Content: content,
		TimeStamp: time.Now(),
	}
}

func (m *Message) GetByteMessage() []byte { //將Message物件轉換成 JSON 格式的 byte陣列
	result, _ := json.Marshal(m)
	return result
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/html/*") //載入HTML檔案
	r.Static("/assets", "./template/assets") // 設定靜態檔案路徑，讓HTML檔案可以引用到CSS、JS檔案
	
	// 設定首頁路由，回傳index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// ----------------- WebSocket -----------------
	// 其他 API 使用 melody 套件來處理 WebSocket
	m := melody.New()
	
	r.GET("/ws", func(c *gin.Context) { //為gin server 建立 WebSocket 連線
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) { //使用 melody 套件的 HandleMessage 事件處理 WebSocket 訊息
		// 任一連線的client發送訊息，就會廣播給所有已連線的client
		m.Broadcast(msg)
	})

	m.HandleConnect(func(session *melody.Session) { //使用 melody 套件的 HandleConnect 事件處理 WebSocket 連線session
		id := session.Request.URL.Query().Get("id") // 透過此function 的參數 session 從連線的client的URL中取出id
		m.Broadcast(NewMessage("other", id, "加入聊天室").GetByteMessage()) //廣播給所有已連線的client
	})

	m.HandleClose(func(session *melody.Session, i int, s string) error { //使用 melody 套件的 HandleConnect 事件處理 WebSocket session 離線
		id := session.Request.URL.Query().Get("id")
		m.Broadcast(NewMessage("other", id, "離開聊天室").GetByteMessage())
		return nil
	})
	r.Run(":5000")
}