package routes

import (
	"net/http"

	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/server/security"
	"github.com/ospokemon/ospokemon/server/sessionman"
)

func Login(dbsalt string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(404)
			return
		}

		if session := sessionman.FromRequestCookie(r); session != nil {
			http.Redirect(w, r, "/", 307)
			return
		}

		username := r.FormValue("username")
		password := security.HashPassword(dbsalt, r.FormValue("password"))

		account, _ := ospokemon.GetAccount(username)
		if account == nil {
			ospokemon.LOG().Add("Username", username).Warn("login: account not found")
			http.Redirect(w, r, "/login/?account", 307)
			return
		}

		if account.Password != password {
			ospokemon.LOG().Add("Username", username).Warn("login: incorrect password")
			delete(ospokemon.Accounts.Cache, username)
			delete(ospokemon.Players.Cache, username)
			http.Redirect(w, r, "/login/?password#"+username, 307)
			return
		}

		if session := sessionman.Get(account); session == nil {
			session = sessionman.Add(account)
			ospokemon.LOG().Add("Username", username).Add("SessionId", session.SessionId).Info("login: create session")
		}

		sessionman.Get(account).WriteSessionId(w)
		http.Redirect(w, r, "/", 307)
	})
}
