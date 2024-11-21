package formatter

import "strings"

type FormatFunc func(string) string
type JoinFunc func([]string) string
type SplitFunc func(string) []string

var (
	DEFAULTFORMAT = FDEFAULT
	DEFAULTJOIN   = JDASHED
	DEFAULTSPLIT  = SSPACE
)

// Splitter
func SSPACE(s string) []string {
	return strings.Split(s, " ")
}

func SDASH(s string) []string {
	return strings.Split(s, "-")
}

func SCOMMA(s string) []string {
	return strings.Split(s, ",")
}

func SUNDER(s string) []string {
	return strings.Split(s, "_")
}

// Formatter
func FDEFAULT(s string) string {
	return DEFAULTJOIN(DEFAULTSPLIT(s))
}

// Joiner
func JDASHED(ss []string) string {
	return strings.Join(ss, "-")
}

func JUNDER(ss []string) string {
	return strings.Join(ss, "_")
}

func JCOMMA(ss []string) string {
	return strings.Join(ss, ",")
}
