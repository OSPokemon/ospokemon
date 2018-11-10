package ospokemonjs

import (
	"github.com/fsnotify/fsnotify"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

func Watch() {
	env := env.Global()
	watcher, err := fsnotify.NewWatcher()
	path := getPath(env)
	if err != nil {
		log.Error(err)
		return
	} else if err = watcher.Add(path); err != nil {
		log.Error(err)
		return
	}

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
