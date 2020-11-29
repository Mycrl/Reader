package storage

import (
	"encoding/json"
	"sync"
	"github.com/syndtr/goleveldb/leveldb"
)

type Chapter struct {
	Name   string `json:"name"`
	Offset int32  `json:"offset"`
}

type Book struct {
    Chapters []Chapter `json:"chapters"`
    Size int64 `json:"size"`
}

type Model struct {
	db    *leveldb.DB
	mutex sync.Mutex
}

func NewModel(path string) *Model {
	db, e := leveldb.OpenFile(path, nil)
	if e != nil {
		panic(e)
	}
	var mutex sync.Mutex
	model := Model{db, mutex}
	return &model
}

func (self *Model) Query(name string) (*Book, error) {
	self.mutex.Lock()
defer self.mutex.Unlock()
	key := []byte(name)
	buffer, get_e := self.db.Get(key, nil)
	if get_e != nil {
		return nil, get_e
	}
    book := Book{}
	parse_e := json.Unmarshal(buffer, &book)
	return &book, parse_e
}

func (self *Model) Write(name string, b *Book) error {
	self.mutex.Lock()
defer self.mutex.Unlock()
	buffer, e := json.Marshal(b)
	if e != nil {
		return e
	}
	key := []byte(name)
	return self.db.Put(key, buffer, nil)
}

func (self *Model) Delete(name string) error {
	self.mutex.Lock()
defer self.mutex.Unlock()
	key := []byte(name)
	return self.db.Delete(key, nil)
}

func (self *Model) List() []string {
	self.mutex.Lock()
	var books []string
	iter := self.db.NewIterator(nil, nil)
defer func() {
    iter.Release()
    self.mutex.Unlock()
}()
for iter.Next() {
    name := string(iter.Key())
    books = append(books, name)
}
	return books
}
