package utils

import (
	message2 "chatClient/message"
	"google.golang.org/protobuf/proto"
	"log"
)

//ExitMessage 断开连接信息
func ExitMessage(name string) []byte{
	exitMsg := &message2.Message{
		MsgType: "exit",
		SendName: name,
	}
	msg, err := proto.Marshal(exitMsg)
	if err != nil{
		log.Println("发送失败，请重新发送！")
	}
	return msg
}

//TalkMessage 发送聊天信息
func TalkMessage(name string,content string) []byte {
	talkMsg := &message2.Message{
		MsgType: "talk",
		MsgContent: content,
		SendName: name,
		UserList: nil,
	}
	msg, err := proto.Marshal(talkMsg)
	if err != nil{
		log.Println("发送失败，请重新发送！")
	}
	return msg
}

//UnMarshal	解析聊天信息
func UnMarshal(msg [] byte) message2.Message {
	Msg := &message2.Message{}
	proto.Unmarshal(msg,Msg)
	return *Msg
}

func UserListMessage(name string) []byte {
	userListMsg := &message2.Message{
		SendName: name,
		MsgType: "userList",
	}
	msg, err := proto.Marshal(userListMsg)
	if err != nil{
		log.Println("发送失败，请重新发送！")
	}
	return msg
}