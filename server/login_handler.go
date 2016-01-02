package server

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
	}

	// redirect if already logged in
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil && Sessions[int(sessionId)] != nil {
			log.WithFields(log.Fields{
				"SessionID": sessionId,
			}).Debug("Auth request redirected")
			http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
			return
		}
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	// hash := md5.Sum([]byte(password))
	// password = string(hash[:])

	if _, err := LoginAccount(username, password); err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Warn("Authorization rejected")
		http.Redirect(w, r, "/?invalid", http.StatusMovedPermanently)
		return
	}

	session := CreateSession(username)
	Sessions[session.SessionId] = session

	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(session.SessionId))
	http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
})
