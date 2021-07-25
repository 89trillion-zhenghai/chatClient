package view

import (
	ctrl2 "chatClient/internal/ctrl"
	model2 "chatClient/internal/model"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Init() {
	//初始化窗体
	myAPP := app.New()
	myWin := myAPP.NewWindow("chat")
	myWin.Resize(fyne.NewSize(800, 700))
	nameLabel := widget.NewLabel("enter name:")
	model2.NameStatus = widget.NewLabel("")
	model2.NameEntry = widget.NewEntry()
	model2.NameEntry.SetPlaceHolder("SmallBai")
	model2.NameEntry.SetText("SmallBai")
	model2.ServerEntry = widget.NewEntry()
	serverLabel := widget.NewLabel("enter server:")
	model2.ServerStatus = widget.NewLabel("")
	model2.ServerEntry.SetPlaceHolder("127.0.0.1:8082")
	model2.ServerEntry.SetText("127.0.0.1:8082")
	statusLabel := widget.NewLabel("status:")
	model2.StatusLabel = widget.NewLabel("Disconnect")
	nameBox := container.NewHBox(nameLabel, model2.NameEntry, model2.NameStatus)
	connect := widget.NewButton("Connect", ctrl2.Connect)
	disconnect := widget.NewButton("Disconnect", ctrl2.DisConnect)
	statusBox := container.NewHBox(statusLabel, model2.StatusLabel)
	serverBox := container.NewHBox(serverLabel, model2.ServerEntry, model2.ServerStatus)
	head := container.NewVBox(container.NewHBox(nameBox, layout.NewSpacer(),statusBox), container.NewHBox(serverBox,layout.NewSpacer(), connect, disconnect))
	model2.UserList = widget.NewLabel("")
	left := widget.NewCard("UserList", "", container.NewScroll(model2.UserList))
	model2.MsgEntry = widget.NewLabel("")
	right := widget.NewCard("", "", container.NewScroll(model2.MsgEntry))
	model2.TalkEntry = widget.NewMultiLineEntry()
	msgCard := widget.NewCard("Input:", "", model2.TalkEntry)
	send := widget.NewButton("send", ctrl2.SendMsg)
	chat := container.NewBorder(head, container.NewVBox(msgCard, send), left, nil, right)
	myWin.SetContent(chat)
	myWin.ShowAndRun()
}
