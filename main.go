package main

import (
	"errors"
	"invokeurls/getdata"
	"invokeurls/validatedata"
	"log"
	"os"
)

func main() {
	defer func() {
		if panicRun := recover(); panicRun != nil {
			logError := errors.New("ERROR: runtime error | PANIC: index out of range")
			logFail := "FAIL: Necessary parameters | EXAMPLE: ./invokeurls <your-urls-list> <your-browser>"
			panicInstruction :=
				`
			Necessary parameters:
			./invokeurls <your-urls-list> <your-browser>
			Example: ./invokeurls urls.txt google-chrome
				`
			log.Println(panicRun, panicInstruction)
			validatedata.WriteErrorLogFile(logFail, logError)
		}
	}()
	validatedata.WriteLogFile("START: invokeurls" + " | LIST: " + (os.Args[1]) + " | BROWSER: " + (os.Args[2]))
	getdata.GetAndFormatUrlsFromFile()
}
