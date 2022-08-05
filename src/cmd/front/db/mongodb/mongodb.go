package mongodb

import (
	"leaf_chat/conf"
	"leaf_chat/leaf/db/mongodb"
	"leaf_chat/leaf/log"
)

var (
	Context *mongodb.DialContext
)

func init() {
	var err error
	Context, err = mongodb.Dial(conf.Server.MongodbAddr, conf.Server.MongodbSessionNum)
	if err != nil {
		log.Fatal("mongondb init is error(%v)", err)
	}
}
