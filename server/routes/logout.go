package routes

import (
	"net/http"

	"ospokemon.com/server/routes/logout"
	"ospokemon.com/server/sessionman"
	"ztaylor.me/log"
)

var Logout = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if session := sessionman.FromRequestCookie(r); session != nil {
		log.Add("Username", session.Username).Add("SessionId", session.SessionId).Info("api/logout")
		logout.LogoutPlayer(session.Username)

		w.Header().Set("Set-Cookie", "SessionId=0; Path=/;")
		http.Redirect(w, r, "/login/#"+session.Username, 307)
	} else {
		log.Add("RemoteAddr", r.RemoteAddr).Warn("logout: no session")
		http.Redirect(w, r, "/login/", 307)
	}
})
