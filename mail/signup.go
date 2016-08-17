package mail

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/ospokemon/ospokemon/util"
	"io/ioutil"
)

func Signup(username string, email string) {
	m := gomail.NewMessage()
	m.SetHeader("From", util.FLAG_MailUser+"@"+util.FLAG_MailServer)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "OSPokemon: Confirm Registration")

	template, _ := ioutil.ReadFile(util.FLAG_MailPath + "signup.html")
	m.SetBody("text/html", fmt.Sprintf(string(template), username))

	if err := dialer.DialAndSend(m); err != nil {
		panic(err)
	}
}
