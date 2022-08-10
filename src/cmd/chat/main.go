package main

import (
	"leaf_chat/cmd/chat/center"
	"leaf_chat/cmd/chat/room"
	"leaf_chat/conf"
	"leaf_chat/leaf"
	"leaf_chat/leaf/module"
	"leaf_chat/pkg/gob"
)

func main() {
	conf.Init()

	gob.Init()

	modules := []module.Module{center.Module}
	modules = append(modules, room.CreateModules()...)
	leaf.Run(modules...)
}
