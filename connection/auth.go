package connection

import (
	// "crypto/md5"
	"fmt"
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/data"
	"net/http"
	"strconv"
)

var sessionGen, _ = mathutil.NewFC32(0, 999999, true)
var Sessions = make(map[int]string)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
	}

	// redirect if already logged in
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0)

		if err == nil && Sessions[int(sessionId)] != "" {
			http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
			return
		}
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	// hash := md5.Sum([]byte(password))
	// password = string(hash[:])

	if data.Players[username] != nil {
		http.Redirect(w, r, "/?duplicate", http.StatusMovedPermanently)
		return
	}

	realPassword := data.PlayerStore.FetchPassword(username)

	if realPassword == "" {
		http.Redirect(w, r, "/?invalid", http.StatusMovedPermanently)
		return
	}

	if realPassword != password {
		http.Redirect(w, r, "/?invalid", http.StatusMovedPermanently)
		return
	}

	sessionId := sessionGen.Next()
	Sessions[sessionId] = username
	w.Header().Set("Set-Cookie", fmt.Sprintf("SessionId=%d", sessionId))

	http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
})
