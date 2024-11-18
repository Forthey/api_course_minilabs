package logger

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"time"
)

type Logger struct {
	wrappedHandler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.wrappedHandler.ServeHTTP(w, r)

	fmt.Printf(
		"%s %s %v %s\n",
		color.MagentaString(r.Method),
		color.CyanString(r.URL.String()),
		time.Since(start),
		time.Now().Format("02.01.2006 15:04:05"),
	)
}

func NewLogger(wrappedHandler http.Handler) *Logger {
	return &Logger{wrappedHandler}
}
