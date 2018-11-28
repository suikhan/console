package console

import (
	"fmt"
)

const (
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
)

var fc = map[string]string{
	"black":        "0;30",
	"dark_gray":    "1;30",
	"blue":         "0;34",
	"light_blue":   "1;34",
	"green":        "0;32",
	"light_green":  "1;32",
	"cyan":         "0;36",
	"light_cyan":   "1;36",
	"red":          "0;31",
	"light_red":    "1;31",
	"purple":       "0;35",
	"light_purple": "1;35",
	"brown":        "0;33",
	"yellow":       "1;33",
	"light_gray":   "0;37",
	"white":        "1;37",
}

func ColorSF(str, foreground string) string {
	code := fc[foreground]
	return fmt.Sprintf(FullColorTpl, code, str)
}
