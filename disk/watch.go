package disk

import (
	"Reader/storage"
	"fmt"
	"path"
	"strings"
	"time"
	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	handle *fsnotify.Watcher
	model  *storage.Model
}

func NewWatch(path string, model *storage.Model) *Watch {
	handle, new_e := fsnotify.NewWatcher()
	if new_e != nil {
		panic(new_e)
	}
	add_e := handle.Add(path)
	if add_e != nil {
		panic(add_e)
	}
	watch := Watch{handle, model}
	return &watch
}

func (w *Watch) Run() {
    for {
        w.poll()
    }
}

func (w *Watch) poll() {
    select {
case event, ok := <-w.handle.Events:
    if !ok {
        return
    }
    if event.Op == fsnotify.Create {
        w.create(event.Name)
    } else
    if event.Op == fsnotify.Remove {
        w.remove(event.Name)
    }
case e, ok := <-w.handle.Errors:
    if !ok {
        return
    }
    if e != nil {
        fmt.Println(e)
    }
    }
}

func (w *Watch) remove(p string) {
    e := w.model.Delete(p)
    if e != nil {
        fmt.Println(e)
    }
}

func (w *Watch) create(p string) {
    time.Sleep(time.Second)
    e := w.parse(p)
    if e != nil {
        fmt.Println(e)
    }
}

func (w *Watch) parse(p string) error {
	book, e := Parse(p)
	if e != nil {
		return e
	}
	name := path.Base(strings.ReplaceAll(p, "\\", "/"))
	w_e := w.model.Write(name, book)
	if w_e != nil {
		return w_e
	}
	return nil
}
