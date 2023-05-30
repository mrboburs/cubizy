package util

import (
	"reflect"
	"strconv"
	"strings"
)

var uintType = reflect.TypeOf(uint(0))
var int64Type = reflect.TypeOf(int64(0))
var intType = reflect.TypeOf(int(0))
var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")

//GetFloat will give flote from interface
func GetFloat(unk interface{}) float64 {
	var value float64
	var err error
	switch i := unk.(type) {
	case bool:
		valueb := bool(i)
		if valueb {
			return 1
		}
		return 0
	case float64:
		value, err = i, nil
	case float32:
		value, err = float64(i), nil
	case int64:
		value, err = float64(i), nil
	case int32:
		value, err = float64(i), nil
	case int:
		value, err = float64(i), nil
	case uint64:
		value, err = float64(i), nil
	case uint32:
		value, err = float64(i), nil
	case uint:
		value, err = float64(i), nil
	case string:
		value, err = strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			value, err = fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			value, err = strconv.ParseFloat(s, 64)
		} else {
			value = 0
		}
	}
	if err != nil {
		value = 0
	}
	return value
}

func GetString(unk interface{}) string {
	var value string = ""
	switch i := unk.(type) {
	case nil:
		value = ""
	case string:
		value = string(i)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			value = sv.String()
		} else {
			value = ""
		}
	}
	return strings.TrimSpace(value)
}

//GetBool will give bool from interface
func GetBool(unk interface{}) bool {

	switch i := unk.(type) {
	case bool:
		return i
	default:
		value := GetInt(unk)
		if value == 0 {
			return false
		}
		return true
	}
}

func GetInt(unk interface{}) int {
	var value int
	var err error
	switch i := unk.(type) {
	case float64:
		value, err = int(i), nil
	case float32:
		value, err = int(i), nil
	case int64:
		value, err = int(i), nil
	case int32:
		value, err = int(i), nil
	case int:
		value, err = i, nil
	case uint64:
		value, err = int(i), nil
	case uint32:
		value, err = int(i), nil
	case uint:
		value, err = int(i), nil
	case string:
		value, err = strconv.Atoi(i)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if unk == nil {
			return 0
		} else if v.Type().ConvertibleTo(intType) {
			fv := v.Convert(intType)
			value, err = fv.NumField(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			value, err = strconv.Atoi(s)
		} else {
			value = 0
		}
	}
	if err != nil {
		value = 0
	}
	return value
}

func GetInt64(unk interface{}) int64 {
	var value int64
	var err error
	switch i := unk.(type) {
	case float64:
		value, err = int64(i), nil
	case float32:
		value, err = int64(i), nil
	case int64:
		value, err = i, nil
	case int32:
		value, err = int64(i), nil
	case int:
		value, err = int64(i), nil
	case uint64:
		value, err = int64(i), nil
	case uint32:
		value, err = int64(i), nil
	case uint:
		value, err = int64(i), nil
	case string:
		value, err = strconv.ParseInt(i, 10, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if unk == nil {
			return 0
		} else if v.Type().ConvertibleTo(int64Type) {
			fv := v.Convert(int64Type)
			value, err = fv.Int(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			value, err = strconv.ParseInt(s, 10, 64)
		} else {
			value = 0
		}
	}
	if err != nil {
		value = 0
	}
	return value
}

func GetUint(unk interface{}) uint {
	var value64 uint64
	var value uint
	var err error
	switch i := unk.(type) {
	case float64:
		value, err = uint(i), nil
	case float32:
		value, err = uint(i), nil
	case int64:
		value, err = uint(i), nil
	case int32:
		value, err = uint(i), nil
	case int:
		value, err = uint(i), nil
	case uint64:
		value, err = uint(i), nil
	case uint32:
		value, err = uint(i), nil
	case uint:
		value, err = i, nil
	case string:
		value64, err = strconv.ParseUint(i, 10, 64)
		value = uint(value64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if unk == nil {
			return 0
		} else if v.Type().ConvertibleTo(uintType) {
			fv := v.Convert(uintType)
			value = uint(fv.Uint())
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			value64, err = strconv.ParseUint(s, 10, 64)
			value = uint(value64)
		} else {
			value = 0
		}
	}
	if err != nil {
		value = 0
	}
	return value
}

func GetID(postItem map[string]interface{}) uint {
	var id uint = 0
	if _, ok := postItem["ID"]; ok {
		id = GetUint(postItem["ID"])
	}
	return id
}

func CantorFunction(k1, k2 uint) uint {
	return ((k1 + k2) * (k1 + k2 + 1) / 2) + k2
}
