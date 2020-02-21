package log

import log "github.com/jeanphorn/log4go"

func GetLogger() *log.Filter {
	log.LoadConfiguration("./conf/log4go.json", "json")
	return log.LOGGER("logs")
}