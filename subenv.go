package subenv

import (
	"os"
	"strconv"
)

/**
Get an integer equivalent of environmental variable. Return default value of the value is not found or is empty
 */
func EnvI(key string, defaultIfNull int) int {
	var val string
	var ok bool

	if val, ok = os.LookupEnv(key); ok {
		if len(val) == 0 {
			return defaultIfNull
		}
	}

	if eint, err := strconv.Atoi(val); err != nil {
		return defaultIfNull
	} else {
		return eint
	}
}

/**
Get the string equivalent of environmental variable. Return default value of the value is not found or is empty
 */
func Env(key string, defaultIfNull string) string {
	var val string
	var ok bool

	if val, ok = os.LookupEnv(key); ok {
		if len(val) == 0 {
			return defaultIfNull
		} else {
			return val
		}
	}

	return defaultIfNull
}

/**
Get the boolean equivalent of environmental variable. Return default value of the value is not found or is empty
 */
func EnvB(key string, defaultIfNull bool) bool {
	var val string
	var ok bool

	if val, ok = os.LookupEnv(key); ok {
		if len(val) == 0 {
			return defaultIfNull
		} else {
			if v, e := strconv.ParseBool(val); e != nil {
				return defaultIfNull
			} else {
				return v
			}
		}
	}

	return defaultIfNull
}
