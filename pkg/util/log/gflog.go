package log

import (
	"fmt"
	gfconfig "gofirst/pkg/config/webserver"
	"log"
	"os"
)

type logDetail struct {
	file_desc  *os.File
	lg_info    *log.Logger
	lg_debug   *log.Logger
	lg_warning *log.Logger
	lg_error   *log.Logger
}

var file_logger *logDetail = nil

// TODO: Add support to mask passwords
func getString(msg *string) *string {
	return msg
}

func GetErrorLogger() *log.Logger {
	return file_logger.lg_error
}

func Info(msg string) {
	msg = *getString(&msg)
	file_logger.lg_info.Printf("%s\n", msg)
}

func Debug(msg string) {
	msg = *getString(&msg)
	file_logger.lg_debug.Printf("%s\n", msg)
}

func Warning(msg string) {
	msg = *getString(&msg)
	file_logger.lg_warning.Printf("%s\n", msg)
}

func Error(msg string) {
	msg = *getString(&msg)
	file_logger.lg_error.Printf("%s\n", msg)
}

func Panic(msg string) {
	msg = *getString(&msg)
	file_logger.lg_error.Panic(msg)
}

func init() {
	fmt.Printf("Log: %s\n", gfconfig.LOGFILE)
	if file_logger == nil {
		file_desc, err := os.OpenFile(gfconfig.LOGFILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
		file_logger = &logDetail{
			file_desc: file_desc,
		}
		file_logger.lg_info = log.New(file_desc, fmt.Sprintf("%-8s:", "INFO"), log.Ldate|log.Ltime|log.Lshortfile)
		file_logger.lg_debug = log.New(file_desc, fmt.Sprintf("%-8s:", "DEBUG"), log.Ldate|log.Ltime|log.Lshortfile)
		file_logger.lg_error = log.New(file_desc, fmt.Sprintf("%-8s:", "ERROR"), log.Ldate|log.Ltime|log.Lshortfile)
		file_logger.lg_warning = log.New(file_desc, fmt.Sprintf("%-8s:", "WARNING"), log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func Close() error {
	fmt.Printf("*** Closing Application ***\n")
	file_logger.file_desc.Close()
	return nil
}
