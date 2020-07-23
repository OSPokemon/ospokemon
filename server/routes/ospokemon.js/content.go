package ospokemonjs

import (
	"io/ioutil"

	"github.com/ospokemon/ospokemon"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"taylz.io/env"
	"taylz.io/types"
)

var Content string
var minifier = minify.New()

func init() {
	minifier.AddFunc("text/javascript", js.Minify)
}

func CreateContent() {
	env := ospokemon.ENV()
	path := getPath(env)
	Content = "$(function(){\n"

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		file, _ := ioutil.ReadFile(path + f.Name())
		Content += string(file)
	}
	Content += "})"

	useminify := types.Bool(env["js-minify"])

	if useminify {
		Content, _ = minifier.String("text/javascript", Content)
	}

	ospokemon.LOG().Add("js-minify", useminify).Debug("ospokemon.js: compile")
}

func getPath(env env.Service) string {
	return env["webpath"] + "/ospokemon.js/"
}
