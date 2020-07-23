package ospokemon

import (
	_env "taylz.io/env"
	_log "taylz.io/log"
)

var env _env.Service
var log *_log.Service

// SetLogger sets the global logger that is required to exist by this codebase
func SetLogger(logger *_log.Service) { log = logger }

// LOG returns the global logger that is required to exist by this codebase
func LOG() *_log.Service {
	return log
}

// ENV returns the global env that is required to exist by this codebase
func ENV() _env.Service {
	if env == nil {
		env = _env.Service{
			"loglevel": "info",
			"logpath":  "log/",
		}
	}
	return env
}
