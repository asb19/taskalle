package utils

import (
	"net/url"
	"strconv"
)

func ParseQueryParamInt(q url.Values, key string, defaultValue int) int {
	valStr := q.Get(key)
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return val
}
