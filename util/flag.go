package util

import (
	"flag"
)

var FLAG_LogLevel = "info"
var FLAG_LogPath = "/var/log/ospokemon/ospokemon.log"
var FLAG_ServerPort = "80"
var FLAG_ServerPath = "/srv/http/ospokemon.io/"
var FLAG_MailServer = "mail.ospokemon.io"
var FLAG_MailPort = 25
var FLAG_MailUser = "auto"
var FLAG_MailPassword = "auto"
var FLAG_MailPath = "/srv/smtp/mail.ospokemon.io/"
var FLAG_DatabasePath = "/var/lib/ospokemon/ospokemon.db"
var FLAG_LaunchDaemon = false
var FLAG_KillDaemon = false
var FLAG_LaunchInteract = false

func init() {
	flag.StringVar(&FLAG_LogLevel, "log", FLAG_LogLevel, "One of [debug,info,warn,error,fatal,panic]")
	flag.StringVar(&FLAG_LogPath, "logpath", FLAG_LogPath, "Log file path")
	flag.StringVar(&FLAG_ServerPort, "port", FLAG_ServerPort, "Server port to open")
	flag.StringVar(&FLAG_ServerPath, "webpath", FLAG_ServerPath, "Server homepage path")
	flag.StringVar(&FLAG_MailServer, "mailserver", FLAG_MailServer, "Email host for outbound messages")
	flag.IntVar(&FLAG_MailPort, "mailport", FLAG_MailPort, "Outbound email port")
	flag.StringVar(&FLAG_MailUser, "mailuser", FLAG_MailUser, "Outbound email username")
	flag.StringVar(&FLAG_MailPassword, "mailpass", FLAG_MailPassword, "Outbound email password")
	flag.StringVar(&FLAG_MailPath, "mailpath", FLAG_MailPath, "Outbound email template path")
	flag.StringVar(&FLAG_DatabasePath, "dbpath", FLAG_DatabasePath, "Database sqlite3 path")
	flag.BoolVar(&FLAG_LaunchDaemon, "daemon", FLAG_LaunchDaemon, "Launch daemon mode")
	flag.BoolVar(&FLAG_KillDaemon, "kill", FLAG_KillDaemon, "Kill existing daemon")
	flag.BoolVar(&FLAG_LaunchInteract, "interact", FLAG_LaunchInteract, "Launch interactive mode")
	flag.Parse()

	loginit()
}

func LogFlags() {
	Log.Debug("FLAG_LogLevel: ", FLAG_LogLevel)
	Log.Debug("FLAG_LogPath: ", FLAG_LogPath)
	Log.Debug("FLAG_ServerPort: ", FLAG_ServerPort)
	Log.Debug("FLAG_ServerPath: ", FLAG_ServerPath)
	Log.Debug("FLAG_MailServer: ", FLAG_MailServer)
	Log.Debug("FLAG_MailPort: ", FLAG_MailPort)
	Log.Debug("FLAG_MailUser: ", FLAG_MailUser)
	Log.Debug("FLAG_MailPassword: ", FLAG_MailPassword)
	Log.Debug("FLAG_MailPath: ", FLAG_MailPath)
	Log.Debug("FLAG_DatabasePath: ", FLAG_DatabasePath)
	Log.Debug("FLAG_LaunchDaemon: ", FLAG_LaunchDaemon)
	Log.Debug("FLAG_KillDaemon: ", FLAG_KillDaemon)
	Log.Debug("FLAG_LaunchInteract: ", FLAG_LaunchInteract)
}
