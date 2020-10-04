package sxlogger

import (
	"io"
	"log"
)

type SxLogger struct {
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
}

func New(out io.Writer) *SxLogger {
	return &SxLogger{
		Info: log.New(out, "[INFO] ", log.Ldate|log.Ltime|log.LUTC),
		Warning: log.New(out, "[WARNING] ", log.Ldate|log.Ltime|log.LUTC),
		Error: log.New(out, "[ERROR] ", log.Ldate|log.Ltime|log.LUTC),
	}
}