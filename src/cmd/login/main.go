package main

import (
	"leaf_chat/cmd/login/gate"
	"leaf_chat/cmd/login/login"
	"leaf_chat/conf"
	"leaf_chat/leaf"
	"leaf_chat/pkg/gob"
)

func main() {
	conf.Init()

	gob.Init()

	leaf.Run(
		gate.Module,
		login.Module,
	)
}
