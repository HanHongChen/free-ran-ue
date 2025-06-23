package logger

type LoggerLevel int

const (
	ERROR = "error"
	WARN  = "warn"
	INFO  = "info"
	DEBUG = "debug"

	LOGGER_ERROR LoggerLevel = 0
	LOGGER_WARN  LoggerLevel = 1
	LOGGER_INFO  LoggerLevel = 2
	LOGGER_DEBUG LoggerLevel = 3
)
