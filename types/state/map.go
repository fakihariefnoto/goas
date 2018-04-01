package state

import (
	c "github.com/fakihariefnoto/goas/types/constanta"
	"os"
	"strings"
)

var (

	// for int64 only
	envMap int64

	// for map with key int64
	sexMap    map[int64]string
	deviceMap map[int64]string

	// for string only

	// fir map with key string
	deviceIDMap map[string]int64
)

func init() {
	fillState()
}

func fillState() {
	sexMap = fillSexMap()
	deviceMap = fillDeviceMap()
	deviceIDMap = fillDeviceIDMap()
	envMap = fillEnvMap()
}

func fillSexMap() map[int64]string {
	return map[int64]string{
		0: "UNDEFINED",
		1: "PRIA",
		2: "WANITA",
	}
}

func fillDeviceMap() map[int64]string {
	return map[int64]string{
		0: "Web",
		1: "Mobile",
		2: "Android",
		3: "IOS",
	}
}

func fillDeviceIDMap() map[string]int64 {
	return map[string]int64{
		"web":     0,
		"mobile":  1,
		"android": 2,
		"ios":     3,
	}
}

func fillEnvMap() int64 {
	env := os.Getenv("APPENV")
	res := map[string]int64{
		c.Development: c.Local,
		c.Staging:     c.Stag,
		c.Production:  c.Prod,
	}
	return res[strings.ToUpper(env)]
}
