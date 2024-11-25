/*
Copyright (C) 2024  Tristan S.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
