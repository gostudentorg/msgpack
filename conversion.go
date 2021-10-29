package msgpack

import (
	"strconv"
	"time"

	"git.gostudent.de/pkg/log"
	"git.gostudent.de/pkg/log/errors"
)

// ToInt converts a number to an integer value.
func ToInt(i interface{}) int {
	switch v := i.(type) {
	case nil:
		return 0
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		log.Debug().Err(errors.New("ToInt: lossy conversion from uint64"), "")
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		val, err := stringToInt64(v)
		if err != nil {
			log.Warn().Err(errors.Errorf("ToInt: string '%s' is not a number", v), "")
			return 0
		}
		return int(val)
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		log.Warn().Err(errors.Errorf("ToInt: type %T not implemented", i), "")
	}
	return 0
}

// ToInt64 converts a number to an int64 value.
func ToInt64(i interface{}) int64 {
	switch v := i.(type) {
	case nil:
		return 0
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		log.Debug().Err(errors.New("ToInt64: lossy conversion from uint64"), "")
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		val, err := stringToInt64(v)
		if err != nil {
			log.Warn().Err(errors.Errorf("ToInt64: string '%s' is not a number", v), "")
			return 0
		}
		return val
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		log.Warn().Err(errors.Errorf("ToInt64: type %T not implemented", i), "")
	}
	return 0
}

// ToUInt64 converts a number to an uint64 value.
func ToUInt64(i interface{}) uint64 {
	switch v := i.(type) {
	case nil:
		return 0
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		val, err := stringToUInt64(v)
		if err != nil {
			log.Warn().Err(errors.Errorf("ToInt64: string '%s' is not a number", v), "")
			return 0
		}
		return val
	default:
		log.Warn().Err(errors.Errorf("ToUInt64: type %T not implemented", i), "")
	}
	return 0
}

// ToFloat64 converts a number to float64 value.
func ToFloat64(i interface{}) float64 {
	switch v := i.(type) {
	case nil:
		return 0
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case string:
		if v == "" {
			return 0
		}
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Warn().Err(errors.Errorf("ToFloat64: string '%s' is not a number", v), "")
			return 0
		}
		return val
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		log.Warn().Err(errors.Errorf("ToFloat64: type %T not implemented", i), "")
	}
	return 0
}

// ToString converts a value to string.
func ToString(i interface{}) string {
	switch v := i.(type) {
	case nil:
		return ""
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		log.Warn().Err(errors.Errorf("ToString: type %T not implemented", i), "")
	}
	return ""
}

// ToBool converts a value to bool.
func ToBool(i interface{}) bool {
	switch v := i.(type) {
	case nil:
		return false
	case int:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from int (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case int8:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from int8 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case int16:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from int16 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case int32:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from int32 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case int64:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from int64 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case uint8:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from uint8 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case uint16:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from uint16 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case uint32:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from uint32 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case uint64:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from uint64 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case float32:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from float32 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case float64:
		log.Debug().Err(errors.Errorf("ToBool: lossy conversion from float64 (%d)", v), "")
		if v == 1 {
			return true
		}
		return false
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			log.Warn().Err(errors.Errorf("ToBool: string '%s' is not a boolean", v), "")
		}
		return b
	case bool:
		return v
	default:
		log.Warn().Err(errors.Errorf("ToBool: type %T not implemented", i), "")
	}
	return false
}

// ToTime converts a value to time.Time.
func ToTime(i interface{}) time.Time {
	switch v := i.(type) {
	case nil:
		return time.Time{}
	case float64:
		return time.Unix(0, int64(v*millisec))
	case uint32:
		return time.Unix(int64(v), 0)
	case int64:
		return time.Unix(v/1000, (v%1000)*millisec)
	case string:
		t, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			log.Warn().Err(err, "ToTime: invalid string value or not unimplemented layout",
				log.String("value", v))
		}
		return t
	case time.Time:
		return v
	case *time.Time:
		return *v
	default:
		log.Warn().Err(errors.Errorf("ToTime: type %T not implemented", i), "")
	}
	return time.Time{}
}

func stringToInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	val, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return val, nil
	}

	// try if it is a float
	val2, r := strconv.ParseFloat(s, 64)
	if r != nil {
		// return original error, not float parsing error
		return 0, err
	}
	return int64(val2), nil
}

func stringToUInt64(s string) (uint64, error) {
	if s == "" {
		return 0, nil
	}
	val, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return val, nil
	}

	// try if it is a float
	val2, r := strconv.ParseFloat(s, 64)
	if r != nil {
		// return original error, not float parsing error
		return 0, err
	}
	return uint64(val2), nil
}
