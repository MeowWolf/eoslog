/*
Package eoslog implements a basic custom logger with standard log levels
*/
package eoslog

import (
	"log"
	"os"
)

// Standard log levels
var (
	Trace    *log.Logger // Trace - just about anything
	Debug    *log.Logger // Debug - Debug messages
	Info     *log.Logger // Info - Important information
	Warn     *log.Logger // Warn - Be concerned
	Error    *log.Logger // Error - Serious problem
	Critical *log.Logger // Critical - Critical problem
)

func init() {
	Trace = log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Debug = log.New(os.Stdout,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warn = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stdout,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Critical = log.New(os.Stdout,
		"CRITICAL: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
