package gate

import (
	"net"
	"leaf_chat/leaf/module"
	"leaf_chat/leaf/chanrpc"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	Skeleton() *module.Skeleton
	ChanRPC()  *chanrpc.Server
}
