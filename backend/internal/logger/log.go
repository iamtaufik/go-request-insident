package logger

import (
	"io"
	"log"
	"os"
)

func SetupLogger() {
	os.MkdirAll("logs", 0755)
	file, err := os.OpenFile(
		"./logs/app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
}
