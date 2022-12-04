package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	logrus "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var (
	logPath          = "./findurl.log"
	logLevel         = logrus.WarnLevel
	logStderr        = true
	logTSFormat      = "2006-01-02 15:04:05"
	logComponentName = "cli"
	logger           = logrus.New()
	defaultSepSize   = 80
)

func InitLog() {
	logger.SetFormatter(
		&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%][" + logComponentName + "] %time% - %msg%\n",
		})

	logger.SetLevel(logLevel)
	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		// Cannot open log file. Logging to stderr
		os.Exit(1)
	}

	if logStderr {
		mw := io.MultiWriter(os.Stderr, f)
		logger.SetOutput(mw)
	} else {
		logger.SetOutput(f)
	}

}

func HandleLogMessage(msg string) string {
	return msg
}

func Trace(msg string) {
	logger.Trace(HandleLogMessage(msg))
}

func Debug(msg string) {
	logger.Debug(HandleLogMessage(msg))
}

func Info(msg string) {
	logger.Info(HandleLogMessage(msg))
}

func Warn(msg string) {
	logger.Warn(HandleLogMessage(msg))
}

func Error(msg string) {
	logger.Error(HandleLogMessage(msg))
}

func Panic(msg string) {
	logger.Panic(HandleLogMessage(msg))
}

func Fatal(msg string) {
	logger.Fatal(HandleLogMessage(msg))
}

func Ssep(sep string, nb int) {
	var res = ""
	size := nb
	if size == 0 {
		size = defaultSepSize
	}
	for i := 0; i < size; i = i + 1 {
		res = res + sep
	}
	Info(res)
}

func Sep(sep string) {
	Ssep(sep, 0)
}
func Dsep() {
	Ssep("#", 0)
}

func Usep() {
	Ssep("_", 0)
}
func Dasep() {
	Ssep("-", 0)
}
func Banner(title string) {
	Title1("BEGIN: " + title)
}
func Title1(msg string) {
	Dsep()
	Info(msg)
	Dsep()
}
func FooterRC(title string, rc int) {
	returnCode := rc
	Dsep()

	if returnCode != 0 {
		Error("RETURN CODE FOR '" + title + "' (" + strconv.Itoa(returnCode) + ")")
		Dsep()
		Fatal("'" + title + "' ENDED WITH ERROR(S): " + strconv.Itoa(returnCode))
	}
	Info("'" + title + "' ENDED SUCCESSFULLY")
	Dsep()
}
func Footer(title string) {
	FooterRC(title, 0)
}

func Println(msg string) {
	fmt.Println(msg)
}

// func Sprintf(format string, a ...interface{}) string {
// 	return fmt.Sprintf(string, a)
// }

func Title2(msg string) {
	Dasep()
	Info(msg)
	Dasep()
}

func HandleError(err error, msg string) {
	if err != nil {
		Error(msg + " [FAILED]")
		Panic(err.Error())
	}
	Info(msg + " [OK]")
}
