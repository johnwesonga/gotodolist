package backend

import "labix.org/v2/mgo"

type MongoDBConn struct {
	session *mgo.Session
}

type ToDo struct {
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
	return m.session
}

func (m *MongoDBConn) Stop() {
	m.session.Close()
}

func (m *MongoDBConn) AddToDo(title, description string) (err error) {
	c := m.session.DB("test").C("people")
	err = c.Insert(&ToDo{title, description})
	if err != nil {
		panic(err)
	}
	return nil
}
