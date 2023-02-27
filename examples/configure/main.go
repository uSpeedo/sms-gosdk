package main

import (
	"os"
	"time"

	"github.com/uSpeedo/usms-sdk-go/um"
	"github.com/uSpeedo/usms-sdk-go/um/auth"
	"github.com/uSpeedo/usms-sdk-go/um/log"
)

func main() {
	cfg := um.NewConfig()

	// enable auto-retry for any request, default is disabled(max retries is 0)
	// it will auto retry for
	//   * network error
	//   * server error with http status 429, 502, 503 and 504
	cfg.MaxRetries = 3

	// disable sdk log, default is log.InfoLevel
	cfg.LogLevel = log.PanicLevel

	// enable sdk debug log level
	// if debug log is enable, it will print all of request params and response body
	cfg.LogLevel = log.DebugLevel

	// set timeout for any request, default is 30s
	cfg.Timeout = 30 * time.Second

	// the followed is used for private user and partner
	// should not be set in general usage

	// custom User-Agent for any request
	cfg.UserAgent = "UCloud-CLI/0.1.0"

	cred := auth.NewCredential()

	// set credential info for any request
	// it is required
	cred.AccessKeySecret = os.Getenv("UCLOUD_ACCESS_KEY_SECRET")
}
