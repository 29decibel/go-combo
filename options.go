package gocombo

import (
  "os"
  "strings"
)

func OptionValue(name string) string {
	for _, arg := range os.Args {
		options := strings.Split(arg, "=")
		if len(options) > 1 && options[0] == name {
      return options[1]
		}
	}
  return ""
}
