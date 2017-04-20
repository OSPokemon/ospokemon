package server

import (
	"net/http"
	"ospokemon.com/log"
	"ospokemon.com/server/api/logout"
	"ospokemon.com/server/session"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if s := session.Find(r); s != nil {
		log.Add("Username", s.Username).Add("SessionId", s.SessionId).Info("api/logout")
		logout.LogoutPlayer(s.Username)

		w.Header().Set("Set-Cookie", "SessionId=0; Path=/;")
		http.Redirect(w, r, "/login/#"+s.Username, 307)
	} else {
		log.Add("RemoteAddr", r.RemoteAddr).Warn("logout: no session")
		http.Redirect(w, r, "/login/", 307)
	}
})
