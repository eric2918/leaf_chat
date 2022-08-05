package accountDB

import (
	"leaf_chat/cmd/login/db/mongodb"
	lmongodb "leaf_chat/leaf/db/mongodb"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Data struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Password string
}

func init() {
	session := mongodb.Context.Ref()
	defer mongodb.Context.UnRef(session)

	getCollection(session).EnsureIndex(mgo.Index{
		Key:    []string{"name"},
		Unique: true,
		Sparse: true,
	})
}

func getCollection(session *lmongodb.Session) *mgo.Collection {
	return session.DB("login").C("account")
}

func Get(name string) (*Data, error) {
	session := mongodb.Context.Ref()
	defer mongodb.Context.UnRef(session)

	result := &Data{}
	err := getCollection(session).Find(bson.M{"name": name}).One(result)
	return result, err
}

func Create(account *Data) error {
	session := mongodb.Context.Ref()
	defer mongodb.Context.UnRef(session)

	return getCollection(session).Insert(account)
}
