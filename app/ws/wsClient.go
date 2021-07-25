package ws

import (
	model2 "chatClient/internal/model"
	utils2 "chatClient/utils"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Connecting(header http.Header){
	url := "ws://"+ model2.ServerEntry.Text+"/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Println(err.Error())
	}
	model2.CreateConn(conn, model2.NameEntry.Text)
	model2.StatusLabel.SetText("Connected")
	//开启监听客户端的动作和服务端的信息
	client := model2.GetConn()
	go client.Accept()
	go client.Send()
}

func DisConnect()  {
	client := model2.GetConn()
	client.Disconnect()
	model2.DisConnect()
}

func SendMsg() {
	client := model2.GetConn()
	sendMsg := utils2.TalkMessage(client.GetName(), model2.TalkEntry.Text)
	client.SendMsg <- sendMsg
	model2.TalkEntry.SetText("")
}
