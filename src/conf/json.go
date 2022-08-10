package conf

import (
	"encoding/json"
	"io/ioutil"
	lconf "leaf_chat/leaf/conf"
	"leaf_chat/leaf/log"
	"os"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ConsoleStdin bool
	ProfilePath string

	MongodbAddr       string
	MongodbSessionNum int

	ServerName      string
	ListenAddr      string
	ConnAddrs       map[string]string
	PendingWriteNum int

	RedisAddr     string
	RedisPassword string
	RedisDb       int

	RoomModuleCount int
	LoginAddr       string
}

func init() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatal("os args of len(%v) less than 2", argsLen)
	}

	filePath := os.Args[1]
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}

func Init() {
	lconf.LogLevel = Server.LogLevel
	lconf.LogPath = Server.LogPath
	lconf.LogFlag = LogFlag
	lconf.ConsolePort = Server.ConsolePort
	lconf.ConsoleStdin = Server.ConsoleStdin
	lconf.ProfilePath = Server.ProfilePath
	lconf.ServerName = Server.ServerName
	lconf.ListenAddr = Server.ListenAddr
	lconf.ConnAddrs = Server.ConnAddrs
	lconf.PendingWriteNum = Server.PendingWriteNum
	lconf.HeartBeatInterval = HeartBeatInterval
}
