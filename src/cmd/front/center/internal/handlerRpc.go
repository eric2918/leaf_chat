package internal

import (
	"fmt"
	"leaf_chat/conf"
	"leaf_chat/leaf/cluster"
	"leaf_chat/leaf/gate"
	"leaf_chat/leaf/log"
	"leaf_chat/msg"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var (
	clientCount     = 0
	accountAgentMap = map[bson.ObjectId]gate.Agent{}
	userAgentMap    = map[bson.ObjectId]gate.Agent{}
)

func handleRpc(id interface{}, f interface{}) {
	cluster.SetRoute(id, ChanRPC)
	skeleton.RegisterChanRPC(id, f)
}

func init() {
	skeleton.RegisterChanRPC("KickAccount", KickAccount)
	skeleton.RegisterChanRPC("AccountOnline", AccountOnline)
	skeleton.RegisterChanRPC("AccountOffline", AccountOffline)
	skeleton.RegisterChanRPC("UserOnline", UserOnline)
	skeleton.RegisterChanRPC("UserOffline", UserOffline)

	handleRpc("GetFrontInfo", GetFrontInfo)
	handleRpc("AddClusterClient", AddClusterClient)
	handleRpc("RemoveClusterClient", RemoveClusterClient)
	handleRpc("BroadcastChatMsg", BroadcastChatMsg)

	handleRpc("Broadcast", Broadcast)
}

func KickAccount(args []interface{}) {
	accountId := args[0].(bson.ObjectId)
	oldAgent, ok := accountAgentMap[accountId]
	if ok {
		oldAgent.Destroy()
	}
}

func AccountOnline(args []interface{}) (interface{}, error) {
	accountId := args[0].(bson.ObjectId)
	agent := args[1].(gate.Agent)
	if oldAgent, ok := accountAgentMap[accountId]; ok {
		oldAgent.Destroy()
		return false, nil
	} else {
		accountAgentMap[accountId] = agent

		clientCount += 1
		cluster.Go("world", "UpdateFrontInfo", conf.Server.ServerName, clientCount)

		log.Debug("%v account is online", accountId)
		return true, nil
	}
}

func AccountOffline(args []interface{}) {
	accountId := args[0].(bson.ObjectId)
	agent := args[1].(gate.Agent)
	oldAgent, ok := accountAgentMap[accountId]
	if ok && agent == oldAgent {
		delete(accountAgentMap, accountId)

		clientCount -= 1
		cluster.Go("world", "UpdateFrontInfo", conf.Server.ServerName, clientCount)

		log.Debug("%v account is offline", accountId)
	}
}

var onlineUserMap map[bson.ObjectId]gate.Agent

func UserOnline(args []interface{}) {
	userId := args[0].(bson.ObjectId)
	agent := args[1].(gate.Agent)
	userAgentMap[userId] = agent
	log.Debug("%v user is online", userId)

	if onlineUserMap == nil{
		onlineUserMap = make(map[bson.ObjectId]gate.Agent)
	}

	onlineUserMap[userId] = agent

	// ???????????????????????????
	msgContent := []byte(fmt.Sprintf("%s online...",userId))
	cluster.Go("world", "Broadcast", msgContent)
}

func UserOffline(args []interface{}) {
	userId := args[0].(bson.ObjectId)
	agent := args[1].(gate.Agent)
	oldAgent, ok := userAgentMap[userId]
	if !ok || agent != oldAgent {
		return
	}

	delete(userAgentMap, userId)
	log.Debug("%v user is offline", userId)


	delete(onlineUserMap, userId)

	// ???????????????????????????
	msgContent := []byte(fmt.Sprintf("%s offline...",userId))
	cluster.Go("world", "Broadcast", msgContent)
}

func Broadcast(args []interface{}) {
	msgContent := args[0].([]byte)
	chatMsg := &msg.ChatMsg{MsgTime: time.Now().Unix(), MsgContent: msgContent}
	for userId, agent := range onlineUserMap {
		chatMsg.UserId = userId
		sendMsg := &msg.F2C_MsgList{MsgList: []*msg.ChatMsg{chatMsg}}
		agent.WriteMsg(sendMsg)
	}
}


func GetFrontInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{clientCount, conf.Server.MaxConnNum, conf.Server.TCPAddr}, nil
}

func AddClusterClient(args []interface{}) {
	serverInfoMap := args[0].(map[string]string)
	for serverName, addr := range serverInfoMap {
		cluster.AddClient(serverName, addr)
	}
}

func RemoveClusterClient(args []interface{}) {
	serverName := args[0].(string)
	cluster.RemoveClient(serverName)
}

func BroadcastChatMsg(args []interface{}) {
	userIds := args[0].([]bson.ObjectId)
	chatMsg := args[1].(*msg.ChatMsg)
	for _, userId := range userIds {
		if agent, ok := userAgentMap[userId]; ok {
			sendMsg := &msg.F2C_MsgList{MsgList: []*msg.ChatMsg{chatMsg}}
			agent.WriteMsg(sendMsg)
		}
	}
}
