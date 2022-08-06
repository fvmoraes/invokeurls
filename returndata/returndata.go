package returndata

import (
	"fmt"
	"invokeurls/validatedata"
	"os/exec"
	"runtime"
	"strings"
)

func PrintInteraction(browser, url string) {
	isUrl := strings.Contains(url, "http")
	if isUrl {
		fmt.Println("BROWSER:", browser)
		fmt.Println("URL:", url)
		validatedata.UrlConnectionStatus(url)
	}
}

func OpenBrowserWithPassedUrl(browser, url string) {
	var logError error
	switch runtime.GOOS {
	case "linux":
		logError = exec.Command(browser, url).Start()
	default:
		logError = fmt.Errorf("S.O.: 'Unsupported platform'")
	}
	defer validatedata.WriteErrorLogFile("", logError)
}
