package logging

import (
	"fmt"
	"time"
)

func printWithPrefix(prefix string, msg string) {
	fmt.Printf("[%s][%s]\t| %s\n", time.Now().Format("2006/01/02"), prefix, msg)
}
func Log(msg string) {
	printWithPrefix("LOG", msg)
}

func Warn(msg string) {
	printWithPrefix("WARN", msg)
}

func Error(msg string, err error) {
	printWithPrefix("ERROR", fmt.Sprintf(msg+"\n\t- %s", err))
}
