package logging

import (
	"fmt"
	"log"
	"os"
)

type LoggerWrapper struct {
	logger *log.Logger
}

var (
	Error = &LoggerWrapper{
		logger: log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	Info = &LoggerWrapper{
		logger: log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	Trace = &LoggerWrapper{
		logger: log.New(os.Stdout, "TRACE ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	Warning = &LoggerWrapper{
		logger: log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile),
	}
)

func (lw *LoggerWrapper) Println(v ...interface{}) {
	lw.write(3, fmt.Sprintf("%v", v...))
}

func (lw *LoggerWrapper) Printf(format string, v ...interface{}) {
	lw.write(3, fmt.Sprintf(format, v...))
}

func (lw *LoggerWrapper) Fatal(v ...interface{}) {
	lw.write(3, fmt.Sprintf("%v", v...))
	os.Exit(1)
}

func (lw *LoggerWrapper) Fatalf(format string, v ...interface{}) {
	lw.write(3, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (lw *LoggerWrapper) Panic(v ...interface{}) {
	s := fmt.Sprintf("%v", v...)
	lw.write(3, s)
	panic(s)
}

func (lw *LoggerWrapper) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	lw.write(3, s)
	panic(s)
}

func (lw *LoggerWrapper) Output(calldepth int, s string) error {
	return lw.write(calldepth+1, s)
}

func (lw *LoggerWrapper) write(calldepth int, line string) error {
	return lw.logger.Output(calldepth, line)
}
