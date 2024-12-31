package utils

import (
	"errors"
	"fmt"
	"time"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

// Tag definitions for logger
const (
	ErrorTAG = 1
	InfoTAG  = 2
	LogTAG   = 3
)

var tagToString = map[int]string{
	ErrorTAG: Red + "[ERROR]" + Reset,
	InfoTAG:  Blue + "[INFO]" + Reset,
	LogTAG:   Cyan + "[LOG]" + Reset,
}

func getTime() string {
	return fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
}

func out(tag int, msg string) {
	fmt.Println(getTime(), tagToString[tag]+White+": "+msg+Reset)
}

func LogError(err error) {
	if err == nil {
		return
	}
	out(ErrorTAG, err.Error())
}

func LogNewError(msg string) {
	LogError(errors.New(msg))
}

func LogInfo(msg string) {
	out(InfoTAG, msg)
}

func Log(msg string) {
	out(LogTAG, msg)
}

func LogCustom(color string, context string, msg string) {
	fmt.Println(getTime(), color+"["+context+"]"+Reset+White+": "+msg+Reset)
}
