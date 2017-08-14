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


/*
The json package always orders keys when marshalling. Specifically:

Maps have their keys sorted lexicographically.
Structs keys are marshalled in the order defined in the struct

*/

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	var mapObj map[string]interface{}
	objBytes, err := json.Marshal(obj)
	if err != nil {
		log.Error(err.Error())
		return mapObj, err
	}
	json.Unmarshal(objBytes, &mapObj)
	return mapObj, err
}

func MapToStruct(mapObj map[string]interface{}) (interface{}, error) {
	var obj interface{}
	mapObjBytes, err := json.Marshal(mapObj)
	if err != nil {
		log.Error(err.Error())
		return obj, err
	}
	json.Unmarshal(mapObjBytes, &obj)
	return obj, err
}

/*------------------------------ struct serialize must use this -----------------------------*/
/*------------------------------ Hash and Sign use this -----------------------------*/
func Serialize(obj interface{}, escapeHTML ...bool) string {
	objMap, err := StructToMap(obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	if len(escapeHTML) >= 1 {
		return _Serialize(objMap, escapeHTML[0])
	}
	return _Serialize(objMap)
}


/*------------- Structs keys are marshalled in the order defined in the struct ------------------*/
func _Serialize(obj interface{}, escapeHTML ...bool) string {
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

//only for selfTest, format json output
func StructSerializePretty(obj interface{}, escapeHTML ...bool) string {
	objMap, err := StructToMap(obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	if len(escapeHTML) >= 1 {
		return SerializePretty(objMap, escapeHTML[0])
	}
	return SerializePretty(objMap)
}

//only for selfTest, format json output
func SerializePretty(obj interface{}, escapeHTML ...bool) string {
	setEscapeHTML := false
	if len(escapeHTML) >= 1 {
		setEscapeHTML = escapeHTML[0]
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	// disabled the HTMLEscape for &, <, and > to \u0026, \u003c, and \u003e in json string
	enc.SetEscapeHTML(setEscapeHTML)
	enc.SetIndent("", "\t")
	err := enc.Encode(obj)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return strings.TrimSpace(buf.String())
	//return strings.Replace(strings.TrimSpace(buf.String()), "\n", "", -1)
}

func Deserialize(jsonStr string) interface{} {
	var dat interface{}
	err := json.Unmarshal([]byte(jsonStr), &dat)
	if err != nil {
		log.Error(err.Error())
	}
	return dat
}
