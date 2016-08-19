package mail

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/ospokemon/ospokemon/util"
	"io/ioutil"
)

func Signup(username string, email string) {
	m := gomail.NewMessage()
	m.SetHeader("From", util.Opt("mailuser")+"@"+util.Opt("mailserver"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "OSPokemon: Confirm Registration")

	template, _ := ioutil.ReadFile(util.Opt("mailpath") + "signup.html")
	m.SetBody("text/html", fmt.Sprintf(string(template), username))

	if err := dialer.DialAndSend(m); err != nil {
		panic(err)
	}
}
