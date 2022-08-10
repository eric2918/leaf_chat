package main

import (
	client2 "leaf_chat/cmd/client/client"
	"leaf_chat/conf"
	"leaf_chat/leaf"
	"leaf_chat/msg"
	"leaf_chat/pkg/gob"
)

func main() {
	conf.Init()

	gob.Init()

	client2.Init(msg.Processor)

	leaf.Run(
		client2.Module,
	)
}
