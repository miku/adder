package fmtutil

import "strings"

func Underline(s string) string {
	return s + "\n" + strings.Repeat("-", len(s))
}
