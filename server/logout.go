package server

import (
	"net/http"
	"ospokemon.com"
	"ospokemon.com/log"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := readsession(r); s != nil {
		account, err := ospokemon.GetAccount(s.Username)

		if err != nil {
			log.Add("Error", err).Error("logout: get account")
			return
		}

		ospokemon.Accounts.Delete(account)
		ospokemon.Accounts.Insert(account)

		w.Header().Set("Set-Cookie", "SessionId=0; Path=/;")
		http.Redirect(w, r, "/login/#"+s.Username, 307)
	} else {
		log.Add("RemoteAddr", r.RemoteAddr).Warn("logout: no session")
		http.Redirect(w, r, "/login/", 307)
	}
})
