package universe

import (
	_ "embed"

	"CodingGhost.de/nameforge/common"
	"CodingGhost.de/nameforge/globals"
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
