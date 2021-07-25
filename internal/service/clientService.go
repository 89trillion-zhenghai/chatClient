package service

import (
	ws2 "chatClient/app/ws"
	model2 "chatClient/internal/model"
	"net/http"
)

//Connect 连接服务器
func Connect()  {
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("name", model2.NameEntry.Text)
	ws2.Connecting(req.Header)
	model2.StatusLabel.SetText("Connected")
	return
}
