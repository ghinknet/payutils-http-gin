package httpgin

import "go.gh.ink/payutils/v3/driver"

func init() {
	driver.RegisterHttp(Name, Driver{})
}
