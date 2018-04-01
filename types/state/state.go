package state

import (
	"strings"
)

func Sex(p int64) string {
	return sexMap[p]
}

func DeviceID(p string) int64 {
	p = strings.ToLower(p)
	return deviceIDMap[p]
}

func Device(p int64) string {
	return deviceMap[p]
}

func Environtment() int64 {
	return envMap
}
