package server

type WebsocketMessage struct {
	Username string
	Event    string
	Message  string
}
