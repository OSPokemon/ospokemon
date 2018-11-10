package ospokemonjs

import (
	"io/ioutil"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"ztaylor.me/cast"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

var Content string
var minifier = minify.New()

func init() {
	minifier.AddFunc("text/javascript", js.Minify)
}

func CreateContent() {
	env := env.Global()
	path := getPath(env)
	Content = "$(function(){\n"

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		file, _ := ioutil.ReadFile(path + f.Name())
		Content += string(file)
	}
	Content += "})"

	if cast.Bool(env.Get("js-minify")) {
		Content, _ = minifier.String("text/javascript", Content)
	}

	log.Add("js-minify", cast.Bool(env.Get("js-minify"))).Debug("ospokemon.js: compile")
}

func getPath(env env.Provider) string {
	return env.Get("webpath") + "/ospokemon.js/"
}
