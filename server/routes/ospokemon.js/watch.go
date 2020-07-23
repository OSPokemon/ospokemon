package ospokemonjs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/ospokemon/ospokemon"
)

func Watch() {
	env := ospokemon.ENV()
	watcher, err := fsnotify.NewWatcher()
	path := getPath(env)
	if err != nil {
		ospokemon.LOG().Error(err)
		return
	} else if err = watcher.Add(path); err != nil {
		ospokemon.LOG().Error(err)
		return
	}

	for {
		select {
		case event := <-watcher.Events:
			ospokemon.LOG().Add("Event", event).Debug("ospokemon.js: rebuild")
			CreateContent()
		case err := <-watcher.Errors:
			ospokemon.LOG().Add("Error", err.Error()).Error("ospokemon.js: watch error")
		}
	}
}
