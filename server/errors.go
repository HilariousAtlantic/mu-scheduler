package main

import (
	"fmt"
	"log"
	"runtime"
)

var isDebug bool = true

func handleError(v interface{}) bool {
	if !isDebug {
		log.Fatal(v)
	}

	if v != nil {
		pc, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, v)
		return false
	} else {
		return true
	}
}

func debug(v interface{}) {
	if isDebug {
		fmt.Println(v)
	}
}
