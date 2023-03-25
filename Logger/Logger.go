package Logger

import (
	"log"
	"os"
)

func ErrorLogger() *log.Logger {
	return log.New(os.Stdout, "ERROR\t", log.Lshortfile|log.Ldate|log.Ltime)
}
