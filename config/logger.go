package config

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:  log.New(writer, color.New(color.FgBlue).Sprint("DEBUG: "), logger.Flags()),
		info:   log.New(writer, color.New(color.FgGreen).Sprint("INFO: "), logger.Flags()),
		warn:   log.New(writer, color.New(color.FgYellow).Sprint("WARNING: "), logger.Flags()),
		err:    log.New(writer, color.New(color.FgRed).Sprint("ERROR: "), logger.Flags()),
		writer: writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Err(v ...interface{}) {
	l.err.Println(v...)
}

// Create Formatted Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
