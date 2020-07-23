package routes

import (
	"net/http"

	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/server/routes/logout"
	"github.com/ospokemon/ospokemon/server/sessionman"
)

var Logout = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if session := sessionman.FromRequestCookie(r); session != nil {
		ospokemon.LOG().Add("Username", session.Username).Add("SessionId", session.SessionId).Info("api/logout")
		logout.LogoutPlayer(session.Username)

		w.Header().Set("Set-Cookie", "SessionId=0; Path=/;")
		http.Redirect(w, r, "/login/#"+session.Username, 307)
	} else {
		ospokemon.LOG().Add("RemoteAddr", r.RemoteAddr).Warn("logout: no session")
		http.Redirect(w, r, "/login/", 307)
	}
})
