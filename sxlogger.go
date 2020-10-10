package sxlogger

import (
	"io"
	"log"
	"strings"
)

type SxLogger struct {
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
}

func New(out io.Writer, hostname string) *SxLogger {
	return &SxLogger{
		Info: log.New(out, strings.ToUpper(hostname)+" [INFO] ", log.Ldate|log.Ltime|log.LUTC),
		Warning: log.New(out, strings.ToUpper(hostname)+" [WARNING] ", log.Ldate|log.Ltime|log.LUTC),
		Error: log.New(out, strings.ToUpper(hostname)+" [ERROR] ", log.Ldate|log.Ltime|log.LUTC),
	}
}