package fmt

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("trace.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	Logger = log.New(io.MultiWriter(os.Stdout, file), "Log: ", log.LstdFlags)
}
