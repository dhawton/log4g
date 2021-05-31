package log4g

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	CRITICAL
	FATAL
)

var (
	levelStrings = [...]string{"DEBUG", "INFO", "WARN", "ERROR", "CRITICAL", "FATAL"}
)

func (l Level) String() string {
	return levelStrings[l]
}

func (l Level) Index() int {
	return int(l)
}
