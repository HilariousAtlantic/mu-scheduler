package main

import (
	"log"
	"runtime"
)

var debug bool = true

func handleError(v interface{}) {
	if !debug {
		log.Fatal(v)
	}

	if v != nil {
		pc, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, v)
	}
}
