package main

import (
	"leaf_chat/cmd/front/center"
	"leaf_chat/cmd/front/gate"
	"leaf_chat/conf"
	"leaf_chat/leaf"
	"leaf_chat/tools/gob"
)

func main() {
	conf.Init()

	gob.Init()

	leaf.Run(
		gate.Module,
		center.Module,
	)
}
