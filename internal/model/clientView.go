package model

import (
	"fyne.io/fyne/v2/widget"
	"time"
)

//NameEntry 用户名
var NameEntry *widget.Entry

var NameStatus *widget.Label

//ServerEntry 服务器地址 ip+port
var ServerEntry *widget.Entry

var ServerStatus *widget.Label

//UserList 用户列表
var UserList *widget.Label

//MsgEntry 消息列表
var MsgEntry *widget.Label

//TalkEntry 聊天框
var TalkEntry *widget.Entry

//StatusLabel 接收反馈信息
var StatusLabel *widget.Label


func OnlineChange(name string)  {
	UserList.SetText(UserList.Text+"\n"+name)
}

//ChangeUserList 接收用户列表信息
func ChangeUserList(list []string) {
	UserList.SetText("")
	for _,user := range list{
		UserList.Text += "\n"+user
	}
}

//ChangeMsg 接收聊天信息
func ChangeMsg(name string,content string) {
	if name == ""{
		name = "server"
	}
	now := time.Now().Format("15:04")
	show := now +":" + content+"\t by "+name
	MsgEntry.SetText(MsgEntry.Text+"\n"+show)
}

//DisConnect 下线
func DisConnect()  {
	UserList.SetText("")
	MsgEntry.SetText("")
	TalkEntry.SetText("")
	StatusLabel.SetText("Disconnected")
}