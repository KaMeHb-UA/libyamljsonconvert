package main

import "C"

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func toCString(str string) *C.char {
	cstr := C.CString(str)
	go preserveCPointer(cstr)
	return cstr
}

func convertToProperType(v interface{}) interface{} {
	switch v.(type) {
	case map[interface{}]interface{}:
		r := map[string]interface{}{}
		for key, val := range v.(map[interface{}]interface{}) {
			r[fmt.Sprint(key)] = convertToProperType(val)
		}
		return r
	case []interface{}:
		r := make([]interface{}, len(v.([]interface{})))
		for i, val := range v.([]interface{}) {
			r[i] = convertToProperType(val)
		}
		return r
	default:
		return v
	}
}

//export decode
func decode(data *C.char) *C.char {
	var parsed interface{}
	err := yaml.Unmarshal([]byte(C.GoString(data)), &parsed)
	if err != nil {
		return toCString("1" + err.Error())
	}
	converted := convertToProperType(parsed)
	res, err := json.Marshal(&converted)
	if err != nil {
		return toCString("1" + err.Error())
	}
	return toCString("0" + string(res))
}

//export encode
func encode(data *C.char) *C.char {
	var parsed interface{}
	err := json.Unmarshal([]byte(C.GoString(data)), &parsed)
	if err != nil {
		return toCString("1" + err.Error())
	}
	res, err := yaml.Marshal(&parsed)
	if err != nil {
		return toCString("1" + err.Error())
	}
	return toCString("0" + string(res))
}
