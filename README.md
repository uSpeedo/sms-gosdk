<h1 align="center">USMS Go SDK</h1>

- [Website](https://uspeedo.com)
- [Documentation]([Hello from uSpeedo | uSpeedo](https://docs.uspeedo.com))

## Installation

### Requirements

- Go 1.10+

### Use `go get`

```bash
go get github.com/uSpeedo/usms-sdk-go
```

**Note** if meet network problem, you can use go proxy to speed up the downloaded, eg: use GOPROXY environment variable

```go
export GOPROXY=https://goproxy.io
```

Replay the command to retry installation.

### Use `go mod`

Add the following snippet to any code.

```go
import _ "github.com/uSpeedo/usms-sdk-go"
```

And execute this commands：

```bash
go mod init
go mod tidy
```

**Note**：If using `go mod` and `Goland IDE` together, please search `vgo` on `Settings`, and enable `vgo` support.

**Note**：If using `go mod` 和 `GOPATH`, notice the `go mod init/tidy` can not run with `GOPATH`, please move out current project from `GOPATH`.

### Use `dep`

```bash
dep ensure -add github.com/uSpeedo/usms-sdk-go
```

## First Using

Currently, Go SDK use `AccessKeyId/AccessKeySecret` as authentication method, the key can be found from：

- [UAPI Key Generation](https://console.uspeedo.com/dashboard)

Here is a simple example：

```go
package main

import (
	"fmt"

	"github.com/uSpeedo/usms-sdk-go/services/usms"
	"github.com/uSpeedo/usms-sdk-go/um"
	"github.com/uSpeedo/usms-sdk-go/um/auth"
	"github.com/uSpeedo/usms-sdk-go/um/config"
	"github.com/uSpeedo/usms-sdk-go/um/log"
)

func main() {
	cfg := config.NewConfig()
	cfg.LogLevel = log.DebugLevel

	credential := auth.NewCredential()
	credential.AccessKeySecret = "..."

	client := usms.NewClient(&cfg, &credential)

	// send request
	req := client.NewSendUSMSMessageRequest()
	req.SigContent = um.String("...")
	req.TemplateId = um.String("UTA2***50501BD")
	req.PhoneNumbers = []string{
		"...",
		"...",
	}
	req.TemplateParams = []string{
		"424242",
	}
	// add header
	req.SetNonce("hz3xevqz")
	req.SetAccessKeyId("314d47318c25a38f5c24df03f6a2a255")
	req.SetSignature("314d47318c25a38f5c24df03f6a2a255")
	req.SetTimestamp(1669370992)
	resp, err := client.SendUSMSMessage(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}


```

## Feedback & Contribution

- [Issue](https://github.com/uSpeedo/usms-gosdk/issues)
- [Pull Request](https://github.com/uSpeedo/usms-gosdk/pulls)
