package logger

import (
	"fmt"

	"github.com/fatih/color"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Notice(msg string, args ...interface{}) {
	c := color.New(color.FgHiCyan)
	c.Println(fmt.Sprintf(msg, args...))
}

func (l *Logger) Error(err error) {
	c := color.New(color.FgHiRed)
	c.Println(fmt.Sprintf("%#v", err))
}

func (l *Logger) Info(msg string, args ...interface{}) {
	white := color.New(color.FgHiWhite)
	white.Println(fmt.Sprintf(msg, args...))
}
