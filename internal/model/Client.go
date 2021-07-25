package model

import (
	message2 "chatClient/message"
	utils2 "chatClient/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"time"
)
const (
	writeWait = 10 * time.Second
	pongWait = 10 * time.Second
	pingPeriod = (pongWait * 9) / 10
)


var Client = &client{
	socket: nil,
	name: "",
	SendMsg: make(chan []byte),
	AcceptMsg: make(chan []byte),
}

type client struct {
	socket	*websocket.Conn
	name	string
	SendMsg	chan []byte
	AcceptMsg	chan []byte
}

func (conn *client) GetName() string{
	return conn.name
}

func CreateConn(conn *websocket.Conn,name string) {
	if Client.socket == nil {
		Client.name = name
		Client.socket = conn
	}
}

func GetConn() *client {
	return Client
}

func (conn *client) Disconnect()  {
	conn.SendMsg <- utils2.ExitMessage(conn.name)
}


func (conn *client) Accept()  {
	conn.socket.SetPongHandler(func(string) error {
		fmt.Println("心跳检测"+time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})
	for {
		if conn.socket == nil {
			break
		}
		_, msg, err := conn.socket.ReadMessage()
		//如果出错就让该用户下线
		if err != nil{
			break
		}
		message := utils2.UnMarshal(msg)
		switch message.MsgType {
		//服务器群发信息
		case "talk":
			ChangeMsg(message.SendName,message.MsgContent)
		if len(message.SendName) == 0{
			conn.SendMsg<- utils2.UserListMessage(conn.name)
		}
		//用户列表信息
		case "userList":
			ChangeUserList(message.UserList)
		}
	}
}

func (conn *client) Send()  {
	ticker := time.NewTicker(pingPeriod)
	for {
		if conn.socket == nil {
			break
		}
		select {
		case msg := <- conn.SendMsg:
			conn.socket.WriteMessage(websocket.TextMessage,msg)
			message := message2.Message{}
			err := proto.Unmarshal(msg, &message)
			if err != nil {
				fmt.Println(err.Error())
			}
			if message.MsgType == "exit"{
				conn.Close()
			}
		case <- ticker.C:
			conn.socket.SetWriteDeadline(time.Now().Add(writeWait))
			err := conn.socket.WriteMessage(websocket.PingMessage, nil)
			if err != nil{
				return
			}
		}
	}
}

func (conn *client) Close()  {
	close(conn.AcceptMsg)
	close(conn.SendMsg)
	conn.socket.Close()
	conn.socket = nil
	conn.AcceptMsg = make(chan []byte)
	conn.SendMsg = make(chan []byte)
}
