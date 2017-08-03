package common

import (
	"time"
	"bytes"
	"strconv"
	"strings"
	"encoding/json"

	"unichain-go/log"

	"github.com/google/uuid"
)

func GenTimestamp() string {
	t := time.Now()
	nanos := t.UnixNano()
	millis := nanos / 1000000 //ms len=13
	return strconv.FormatInt(millis, 10)
}

func GenDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func GenerateUUID() string {
	return uuid.New().String()
}

func Serialize(obj interface{}, escapeHTML ...bool) string { //FIXME
	setEscapeHTML := false
	if len(escapeHTML) >= 1 {
		setEscapeHTML = escapeHTML[0]
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	// disabled the HTMLEscape for &, <, and > to \u0026, \u003c, and \u003e in json string
	enc.SetEscapeHTML(setEscapeHTML)
	err := enc.Encode(obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return strings.TrimSpace(buf.String())
	//return strings.Replace(strings.TrimSpace(buf.String()), "\n", "", -1)
}