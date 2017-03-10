package game


type Scripter struct {
	Script string
	Data   map[string]interface{}
}

func MakeScripter() Scripter {
	return Scripter{
		Data: make(map[string]interface{}),
	}
}

var Scripts = make(map[string]func(*Entity, map[string]interface{}) error)

