package client

import (
	"leaf_chat/base"
	"leaf_chat/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	Module   = new(ClientModule)
	ChanRPC  = skeleton.ChanRPCServer
)

type ClientModule struct {
	*module.Skeleton
}

func (m *ClientModule) OnInit() {
	m.Skeleton = skeleton
}

func (m *ClientModule) OnDestroy() {

}
