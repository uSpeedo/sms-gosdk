<h1 align="center">USMS Go SDK</h1>

- [Website](https://uspeedo.com)
- [Hello from uSpeedo | uSpeedo](https://docs.uspeedo.com)

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
	"encoding/json"
	"fmt"
	"github.com/uSpeedo/usms-sdk-go/private/utils"
	"time"

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
	credential.AccessKeyId = "..."
	credential.AccessKeySecret = "..."

	client := usms.NewClient(&cfg, &credential)
	// send request
	req := client.NewSendBatchUSMSMessageRequest()
	req.AccountId = um.Int(0)
	req.Action = um.String("SendBatchUSMSMessage")
	req.TaskContent = []usms.SendBatchInfo{
		{
			TemplateId: "...",
			SenderId:   "",
			Target: []usms.SendBatchTarget{
				{
					Phone: "130xxxx1321",
				},
				{
					Phone: "130xxxx1321",
				},
			},
		},
	}
	//add header
	req.SetNonce(utils.RandStr(10))
	req.SetAccessKeyId(credential.AccessKeyId)
	req.SetSignature(credential.CreateSign(JSONMethod(req)))
	t, _ := time.ParseDuration("-2m")
	req.SetTimestamp(time.Now().Add(t).Unix())
	resp, err := client.SendBatchUSMSMessage(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func JSONMethod(content interface{}) map[string]interface{} {
	data, _ := json.Marshal(&content)
	m := make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	return m
}


```

## Feedback & Contribution

- [Issue](https://github.com/uSpeedo/usms-gosdk/issues)
- [Pull Request](https://github.com/uSpeedo/usms-gosdk/pulls)
