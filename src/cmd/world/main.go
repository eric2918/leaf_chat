package main

import (
	"leaf_chat/cmd/world/center"
	"leaf_chat/conf"
	"leaf_chat/leaf"
	"leaf_chat/tools/gob"
)

func main() {
	conf.Init()

	gob.Init()

	leaf.Run(
		center.Module,
	)
}
