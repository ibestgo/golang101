package main

import (
	"fmt"
	//"sample/logger"			// if sources at "sample/logger" defined "package logger", I can just use logger
	mylogger2 "sample/logger"	// Note that "sample/logger" has "mylogger" package defined but I can import as any name like mylogger2.
)

func main() {
	fmt.Println("Started")
	
	// this is from helper.go (main package) hence it is at the same folder (/basic1) with main.go.
	// No need to import within the same folder (package)
	Haha("Baechul")

	// However this is from helper2.go (logger package) hence it has to be imported from another folder that happens
	// to be /basic1/logger/helper2.go. All source codes of logger package should be under /basic1/logger
	// How to import other package? 
	// import path always starts with module path (in go.mod) of that package. As logger is in the same model and 
	// its module path in go.mod is "sample". Hence import path would be sample/logger
	mylogger2.Log("car received")
}
