package fileWatcher

import (
	"github.com/Thenecromance/OurStories/base/log"
	"github.com/fsnotify/fsnotify"
)

var (
	inst *Watcher
)

func init() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
		return
	}

	inst = &Watcher{
		ptr:          watcher,
		callbackList: make(map[string]FileCallback),
	}

	go inst.watchThread()

}

func Close() {
	inst.ptr.Close()
	inst.callbackList = nil
}

type FileCallback struct {
	OnChanged func()
	OnDeleted func()
	OnCreated func()
	OnRenamed func()
}

type Watcher struct {
	ptr          *fsnotify.Watcher
	callbackList map[string]FileCallback
}

func (w *Watcher) watchThread() {
	for {
		select {
		case event := <-w.ptr.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				w.callbackList[event.Name].OnChanged()
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				w.callbackList[event.Name].OnCreated()
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				w.callbackList[event.Name].OnDeleted()
			}
			if event.Op&fsnotify.Rename == fsnotify.Rename {
				w.callbackList[event.Name].OnRenamed()
			}
		case err := <-w.ptr.Errors:
			log.Error(err)
		}
	}
}

func WatchFile(path string, callback FileCallback) {
	inst.callbackList[path] = callback
	inst.ptr.Add(path)
}
