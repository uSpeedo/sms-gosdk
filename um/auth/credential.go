/*
Package auth is the credential utilities of sdk
*/
package auth

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

// Credential is the information of credential keys
type Credential struct {
	// Access Key Secret
	AccessKeySecret string

	// Access Key Id
	AccessKeyId string

	// Time the credentials will expire.
	CanExpire bool
	Expires   time.Time
}

// NewCredential will return credential config with default values
func NewCredential() Credential {
	return Credential{}
}

// CreateSign will encode query string to credential signature.
func (c *Credential) CreateSign(params map[string]interface{}) string {
	query := mapToQuery(params)
	return c.CalculateSignature(query)
}

func (c *Credential) MapToQuery(params map[string]interface{}) string {
	query := mapToQuery(params)
	return query
}

func (c *Credential) CalculateSignature(query string) string {
	var buf bytes.Buffer
	buf.WriteString(query)
	buf.WriteString(c.AccessKeySecret)
	origin := buf.String()
	hashed := sha1.Sum([]byte(origin))

	sign := hex.EncodeToString(hashed[:])

	return sign
}

func mapToQuery(params map[string]interface{}) string {
	sortedKeys := extractKeys(params)
	var buf bytes.Buffer
	for _, k := range sortedKeys {
		buf.WriteString(k)
		buf.WriteString(any2String(params[k]))
	}

	origin := buf.String()
	return origin
}
