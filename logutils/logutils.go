package logutils

import (
	"log"
	"os"
)

// New wraps log.New.
func New() *log.Logger {
	const flg = log.LstdFlags | log.Lmicroseconds | log.Lshortfile
	return log.New(os.Stdout, "", flg)
}
