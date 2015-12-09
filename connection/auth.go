package connection

import (
	// "crypto/md5"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/registry"
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
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil && Sessions[int(sessionId)] != "" {
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

	if registry.Accounts[username] == nil {
		registry.AccountLoader(username)
	}
	account := registry.Accounts[username]

	if account == nil {
		log.WithFields(log.Fields{
			"player": username,
		}).Warn("Authorization rejected. account does not exist")
		http.Redirect(w, r, "/?invalid", http.StatusMovedPermanently)
		return
	}

	if account.Online {
		log.WithFields(log.Fields{
			"player": username,
		}).Warn("Authorization rejected. Player already signed on")
		http.Redirect(w, r, "/?duplicate", http.StatusMovedPermanently)
		return
	}

	if password != account.Password {
		log.WithFields(log.Fields{
			"player": username,
		}).Warn("Authorization password failure")
		http.Redirect(w, r, "/?invalid", http.StatusMovedPermanently)
		return
	}

	sessionId := sessionGen.Next()
	Sessions[sessionId] = username

	log.WithFields(log.Fields{
		"Player":    username,
		"SessionID": sessionId,
	}).Info("Authorization successful")

	w.Header().Set("Set-Cookie", fmt.Sprintf("SessionId=%d", sessionId))
	http.Redirect(w, r, "/play/", http.StatusMovedPermanently)
})
