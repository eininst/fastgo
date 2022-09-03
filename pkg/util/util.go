package util

import (
	"github.com/valyala/fasttemplate"
)

func Sprintf(format string, parms map[string]interface{}) string {
	t := fasttemplate.New(format, "{{", "}}")
	return t.ExecuteString(parms)
}
