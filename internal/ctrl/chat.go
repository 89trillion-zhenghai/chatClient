package ctrl

import (
	ws2 "chatClient/app/ws"
	model2 "chatClient/internal/model"
	"net/http"
)

func Connect() {
	name := model2.NameEntry.Text
	//如果是已连接状态点击连接按钮，不响应
	if model2.StatusLabel.Text == "Connected" {
		return
	}
	//用户名为空
	if len(name) == 0 {
		model2.NameStatus.SetText("name can't null")
		return
	}
	//服务器地址为空
	if len(model2.ServerEntry.Text) == 0 {
		model2.ServerStatus.SetText("server can't null")
		return
	}
	//开启连接
	header := http.Header{}
	header.Add("name",name)
	ws2.Connecting(header)
}

func DisConnect() {
	if model2.StatusLabel.Text == "DisConnected" {
		return
	}
	ws2.DisConnect()
}

func SendMsg() {
	sendMsg := model2.TalkEntry.Text
	if len(sendMsg) == 0 {
		return
	}
	ws2.SendMsg()
}