package validatedata

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const myTimeFormat = "02/01/2006 15:04:05"

func WriteErrorLogFile(logFail string, logError error) {
	if logFail != "" {
		WriteLogFile(logFail)
	}
	if logError != nil {
		fmt.Println("ERROR: " + logError.Error())
		WriteLogFile("ERROR: " + logError.Error())
		os.Exit(0)
	}
}

func WriteLogFile(information string) {
	logFile, _ := os.OpenFile("invokeurls.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logFile.WriteString(time.Now().Format(myTimeFormat) + " - " + information + "\n")
	defer logFile.Close()
}

func UrlConnectionStatus(url string) {
	httpStatus, logError := http.Get(url)
	if httpStatus.StatusCode != 200 {
		WriteErrorLogFile("STATUS: "+httpStatus.Status+" | URL: "+url, logError)
	}
	println("STATUS:", httpStatus.Status)
	println("---")
	defer httpStatus.Body.Close()
}
