package util

import (
	"fmt"
	"strconv"

	"mysvrs/errors"

	"github.com/nybuxtsui/log"
)

func GetVersion() string {
	return version
}

func AnyToUint64(v interface{}) (uint64, error) {
	if v == nil {
		log.Error("AnyToInt64|nil")
		return 0, errors.DataError
	}
	switch v := v.(type) {
	case string:
		if v == "" {
			log.Error("AnyToInt64|empty_string", v)
			return 0, errors.DataError
		} else {
			outdata, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				log.Error("AnyToInt64|%v|%v", v, err)
				return 0, errors.DataError
			}
			return outdata, nil
		}
	case []string:
		if len(v) == 0 || v[0] == "" {
			log.Error("AnyToInt64|empty_array", v)
			return 0, errors.DataError
		} else {
			outdata, err := strconv.ParseUint(v[0], 10, 64)
			if err != nil {
				log.Error("AnyToInt64|%v|%v", v, err)
				return 0, errors.DataError
			}
			return outdata, nil
		}
	case []interface{}:
		if len(v) == 0 {
			log.Error("AnyToInt64|empty_array", v)
			return 0, errors.DataError
		} else {
			return AnyToUint64(v[0])
		}
	default:
		return NumToUint64(v)
	}
}

func AnyToInt64(v interface{}) (int64, error) {
	if v == nil {
		log.Error("AnyToInt64|nil")
		return 0, errors.DataError
	}
	switch v := v.(type) {
	case string:
		if v == "" {
			log.Error("AnyToInt64|empty_string", v)
			return 0, errors.DataError
		} else {
			outdata, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Error("AnyToInt64|%v|%v", v, err)
				return 0, errors.DataError
			}
			return outdata, nil
		}
	case []string:
		if len(v) == 0 || v[0] == "" {
			log.Error("AnyToInt64|empty_array", v)
			return 0, errors.DataError
		} else {
			outdata, err := strconv.ParseInt(v[0], 10, 64)
			if err != nil {
				log.Error("AnyToInt64|%v|%v", v, err)
				return 0, errors.DataError
			}
			return outdata, nil
		}
	case []interface{}:
		if len(v) == 0 {
			log.Error("AnyToInt64|empty_array", v)
			return 0, errors.DataError
		} else {
			return AnyToInt64(v[0])
		}
	default:
		return NumToInt64(v)
	}
}

func NumToUint64(arg interface{}) (uint64, error) {
	switch v := arg.(type) {
	case uint64:
		return v, nil
	case int:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case float32:
		return uint64(v), nil
	default:
		log.Error("ToInt64|unknown_type|%T,%v\n", arg, arg)
		return 0, errors.DataError
	}
}

func NumToInt64(arg interface{}) (int64, error) {
	switch v := arg.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	default:
		log.Error("ToInt64|unknown_type|%T,%v\n", arg, arg)
		return 0, errors.DataError
	}
}

//转换数字字符串
func StrToInt64(instr string) (int64, error) {
	tmp, err := strconv.ParseInt(instr, 10, 64)
	if err != nil {
		log.Error("StrToInt64|Invalid_String|%v|%v", err, instr)
		return 0, errors.DataError
	}
	return tmp, nil
}

//转换为字符串
func ToString(arg interface{}) string {
	if arg == nil {
		return ""
	}
	switch v := arg.(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
