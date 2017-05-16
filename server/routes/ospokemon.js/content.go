package ospokemonjs

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"io/ioutil"
	"ospokemon.com/log"
	"ospokemon.com/option"
)

var Content string
var path = option.String("webpath") + "/ospokemon.js/"
var minifier = minify.New()

func init() {
	minifier.AddFunc("text/javascript", js.Minify)
}

func CreateContent() {
	Content = "$(function(){\n"

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		file, _ := ioutil.ReadFile(path + f.Name())
		Content += string(file)
	}
	Content += "})"

	if option.Bool("js-minify") {
		Content, _ = minifier.String("text/javascript", Content)
	}

	log.Add("js-minify", option.Bool("js-minify")).Debug("ospokemon.js: compile")
}
