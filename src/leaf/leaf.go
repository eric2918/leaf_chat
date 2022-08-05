package leaf

import (
	"leaf_chat/leaf/cluster"
	"leaf_chat/leaf/conf"
	"leaf_chat/leaf/console"
	"leaf_chat/leaf/log"
	"leaf_chat/leaf/module"
	"os"
	"os/signal"
)

var (
	OnDestroy func()
)

func Run(mods ...module.Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Leaf %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Leaf closing down (signal: %v)", sig)

	if OnDestroy != nil {
		OnDestroy()
	}
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
