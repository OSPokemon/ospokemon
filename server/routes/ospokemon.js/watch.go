package ospokemonjs

import (
	"github.com/fsnotify/fsnotify"
	"ospokemon.com/log"
)

var watcher *fsnotify.Watcher

func init() {
	if w, err := fsnotify.NewWatcher(); err != nil {
		log.Error(err)
	} else if err = w.Add(path); err != nil {
		log.Error(err)
	} else {
		watcher = w
		go Watch()
	}
}

func Watch() {
	for {
		select {
		case event := <-watcher.Events:
			log.Add("Event", event).Debug("ospokemon.js: rebuild")
			CreateContent()
		case err := <-watcher.Errors:
			log.Add("Error", err.Error()).Error("ospokemon.js: watch error")
		}
	}
}
