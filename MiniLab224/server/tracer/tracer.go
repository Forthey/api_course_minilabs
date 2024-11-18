package tracer

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
)

type Trace struct {
	file     string
	line     int
	function string
}

const (
	Info = uint8(iota)
	Warning
	Error
)

var GetTrace = func() *Trace {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return &Trace{"?", 0, "?"}
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return &Trace{file, line, "?"}
	}

	return &Trace{file, line, fn.Name()}
}

func getColoredNames() (string, string, string) {
	c := color.New(color.FgCyan)
	return c.Sprint("file"), c.Sprint("func"), c.Sprint("line")
}

func AppendTrace(trace *Trace, text string, textType_ ...uint8) string {
	textType := Error
	if len(textType_) == 1 {
		textType = textType_[0]
	}
	typeTag := ""
	col := color.FgCyan
	switch textType {
	case Error:
		typeTag = "ERROR"
		col = color.FgRed
	case Warning:
		typeTag = "WARNING"
		col = color.FgYellow
	case Info:
		typeTag = "INFO"
		col = color.FgGreen
	}
	typeTag = color.New(col).Sprintf(typeTag)
	file, f, line := getColoredNames()
	return fmt.Sprintf(
		"%s: %s, %s: %s, %s: %d\n\t%s\t%s",
		file, trace.file, f, trace.function, line, trace.line, typeTag, text,
	)
}
