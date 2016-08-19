package mail

import (
	"github.com/go-gomail/gomail"
	"github.com/ospokemon/ospokemon/util"
)

var dialer = gomail.NewDialer(util.Opt("mailserver"), util.OptInt("mailport"), util.Opt("mailuser"), util.Opt("mailpass"))
