package mail

import (
	"github.com/go-gomail/gomail"
	"github.com/ospokemon/ospokemon/util"
)

var dialer = gomail.NewDialer(util.FLAG_MailServer, util.FLAG_MailPort, util.FLAG_MailUser, util.FLAG_MailPassword)
