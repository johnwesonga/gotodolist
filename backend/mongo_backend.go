package backend

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type MongoDBConn struct {
	session *mgo.Session
}

type ToDo struct {
  Id          bson.ObjectId    "_id,omitempty"
	Title       string
	Description string
}

func NewMongoDBConn() *MongoDBConn {
	return &MongoDBConn{}
}

func (m *MongoDBConn) Connect(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	m.session = session
	return m.session.Clone()
}

func (m *MongoDBConn) Stop() {
	m.session.Close()
}

func (m *MongoDBConn) AddToDo(title, description string) (err error) {
	c := m.session.DB("test").C("todo")
	err = c.Insert(&ToDo{"", title, description})
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func (m *MongoDBConn) ListToDo() []ToDo {
	results := []ToDo{}
	collection := m.session.DB("test").C("todo")
	iter := collection.Find(nil).Limit(100).Iter()
	err := iter.All(&results)
	if err != nil {
		panic(iter.Err())
	}
	return results
}

func (m *MongoDBConn) DeleteToDo(id string) {
	collection := m.session.DB("test").C("todo")
	err := collection.Find(bson.M{"_id": id})
	if err != nil {
		panic(err)
	}

}
