package File

import (
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/Thenecromance/OurStories/utility/mq"
	"github.com/fsnotify/fsnotify"
)

type watchDog struct {
	watcher *fsnotify.Watcher
}

func (w *watchDog) evenHandler() {
	for {
		select {
		case event := <-w.watcher.Events:
			log.Debugf("watch dog event: %s", event.String())
			if event.Op.Has(fsnotify.Write) {

			}
			if event.Op.Has(fsnotify.Create) {

			}
			if event.Op.Has(fsnotify.Remove) {

			}
			if event.Op.Has(fsnotify.Rename) {

			}
		case err := <-w.watcher.Errors:
			log.Error(err)
		}
	}
}

func (w *watchDog) publishToSubscribers(fileName string) {
	mq.Publish(mq.FileOp, fileName)

}

func (w *watchDog) hasOp(Op, target fsnotify.Op) bool {
	return Op&target == target
}

func (w *watchDog) Close() {
	w.watcher.Close()
}
