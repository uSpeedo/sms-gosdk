package request

import (
	"encoding/json"
	"fmt"

	"github.com/uSpeedo/usms-sdk-go/private/protocol/http"
	"github.com/uSpeedo/usms-sdk-go/um/auth"
	"github.com/uSpeedo/usms-sdk-go/um/config"
)

type JSONEncoder struct {
	cfg  *config.Config
	cred *auth.Credential
}

// Encode will encode request struct instance as a map for json encoded
func (e *JSONEncoder) Encode(req Common) (*http.HttpRequest, error) {
	if req == nil {
		return nil, fmt.Errorf("invalid request, got nil")
	}

	httpReq := http.NewHttpRequest()
	_ = httpReq.SetURL(e.cfg.BaseUrl)
	_ = httpReq.SetTimeout(req.GetTimeout())
	_ = httpReq.SetMethod("POST")
	_ = httpReq.SetQuery("Action", req.GetAction()) // workaround for http log handler
	_ = httpReq.SetHeader(http.HeaderNameContentType, http.MimeJSON)
	_ = httpReq.SetHeader(http.HeaderNameContentType, http.MimeJSON)
	_ = httpReq.SetHeader("AccessKeyId", req.GetAccessKeyId())
	_ = httpReq.SetHeader("Nonce", req.GetNonce())
	_ = httpReq.SetHeader("Signature", req.GetSignature())
	_ = httpReq.SetHeader("Timestamp", req.GetTimestamp())
	// encode struct to map
	payload, err := EncodeJSON(req)
	if err != nil {
		return nil, err
	}
	payload = e.cred.Apply(payload)

	// marshal payload as request body
	bs, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	_ = httpReq.SetRequestBody(bs)
	return httpReq, nil
}

func NewJSONEncoder(cfg *config.Config, cred *auth.Credential) Encoder {
	return &JSONEncoder{cfg: cfg, cred: cred}
}

func EncodeJSON(req Common) (map[string]interface{}, error) {
	return structToMap(req)
}
