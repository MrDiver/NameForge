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

package common

import (
	"math/rand/v2" // https://go.dev/blog/randv2
	"strings"
)

func Lines(content string) []string {
	return strings.Split(content, "\n")
}

func SelectRandom[T any](l []T) T {
	return l[rand.N(len(l))]
}
