package bus_stop

import (
	"log"
	"os"
)

var (
	busLogger  *log.Logger
	stopLogger *log.Logger
)

func init() {
	busLogger = log.New(os.Stdout,
		"BUS: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	stopLogger = log.New(os.Stdout,
		"STOP: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
