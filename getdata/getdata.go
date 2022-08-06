package getdata

import (
	"invokeurls/returndata"
	"invokeurls/validatedata"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func GetAndFormatUrlsFromFile() {
	urlsList, browser := getParametersFromArguments()
	urls := strings.Split(string(urlsList), "\n")
	for _, url := range urls {
		returndata.PrintInteraction(browser, url)
		time.Sleep(time.Second)
		returndata.OpenBrowserWithPassedUrl(browser, url)
	}
}

func getParametersFromArguments() ([]byte, string) {
	urlsList, logError := ioutil.ReadFile(os.Args[1])
	browser := os.Args[2]
	defer validatedata.WriteErrorLogFile("", logError)
	return urlsList, browser
}
