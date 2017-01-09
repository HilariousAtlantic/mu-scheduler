package main

import (
	"log"
	"runtime"
)

func handleError(err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
}
