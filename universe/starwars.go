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
package universe

import (
	_ "embed"

	"github.com/MrDiver/nameforge/common"
	"github.com/MrDiver/nameforge/globals"
)

var (
	//go:embed assets/STARWARS.LST
	content  string
	STARWARS = starwars{}
)

func init() {
	STARWARS.names = common.Lines(content)
	content = ""
	globals.AvailableUniverses[STARWARS.String()] = &STARWARS
}

type starwars struct {
	names []string
}

func (sw *starwars) String() string  { return "STARWARS" }
func (sw *starwars) Names() []string { return sw.names }
