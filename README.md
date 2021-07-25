## 项目简介

1、本项目主要实现了聊天服务之客户端界面与功能。界面采用fyne设计，界面风格使用时下流行的简约风，功能完整。通信协议使用websocket协议实现客户端和服务端的通信，通信数据采用了protobuf。



## 代码逻辑分层



| 层      | 文件夹            | 功能介绍                           | 调用关系               |
| ------- | ----------------- | ---------------------------------- | ---------------------- |
| ws      | /ws               | 管理websocket连接，断开，发送消息  | 被ctrl调用             |
| ctrl    | /Internal/ctrl    | 参数校验、按钮事件的响应           | 被view层调用           |
| view    | /view             | 客户端界面启动                     | 调用ctrl层处理点击事件 |
| model   | /internal/model   | 数据模型、界面信息交互             | 被其他层调用           |
| message | /internal/message | 通信数据protobuf文件和编译后go文件 | 被其他层调用           |
| service | /internal/service | 通用业务处理                       | 被ctrl层调用           |
| msgUtil | /internal/msgUtil | 封装待发送的信息和解析接收的信息   | 被其他层调用           |

## 通信数据设计

本项目使用protobuf、以二进制数据传输通信信息。

| 变量名     | 变量类型 | 注释             |
| ---------- | -------- | ---------------- |
| msgType    | string   | 消息类型         |
| msgContent | string   | 消息内容         |
| sendName   | string   | 发送者名字       |
| userList   | []string | 当前在线用户列表 |

## 界面设计

未命名文件 (8).jpg![未命名文件 (8)]

连接

![image-20210723174651889](/Users/yangzhenghai/Library/Application Support/typora-user-images/image-20210723174651889.png)



断开连接

![image-20210723174733889](/Users/yangzhenghai/Library/Application Support/typora-user-images/image-20210723174733889.png)

## 第三方库

### gin

```
go语言的web框架
https://github.com/gin-gonic/gin
```

### Gorilla WebSocket

```
go语言对websocket协议的实现
http://github.com/gorilla/websocket
```

### protobuf

```
包含go语言处理proto数据的函数
http://github.com/golang/protobuf/proto
```

### fyne

```go
fyne.io/fyne/v2
```

## 如何编译执行

进入app目录编译

```
go build
```

运行可执行文件

```
./main
```

## todo

将代码进一步分层

## 流程图

(https://user-images.githubusercontent.com/86946999/126900421-8c7c165e-e528-451a-92a5-1f16a70f44f6.jpg)


