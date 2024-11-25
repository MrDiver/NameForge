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
package language

import (
	_ "embed"

	"github.com/MrDiver/nameforge/common"
	"github.com/MrDiver/nameforge/globals"
)

var (
	//go:embed assets/EN/nouns.txt
	ennouns string
	//go:embed assets/EN/verbs.txt
	enverbs string
	//go:embed assets/EN/adjectives.txt
	enadjectives string
	//go:embed assets/EN/adverbs.txt
	enadverbs string
	ENGLISH   = english{}
)

func init() {
	ENGLISH.nouns = common.Lines(ennouns)
	ENGLISH.verbs = common.Lines(enverbs)
	ENGLISH.adjectives = common.Lines(enadjectives)
	ENGLISH.adverbs = common.Lines(enadverbs)
	ennouns = ""
	enverbs = ""
	enadjectives = ""
	enadverbs = ""
	globals.AvailableLanguages[ENGLISH.String()] = &ENGLISH
}

type english struct {
	nouns      []string
	verbs      []string
	adjectives []string
	adverbs    []string
}

func (en *english) String() string      { return "EN" }
func (en *english) Noun() []string      { return en.nouns }
func (en *english) Advective() []string { return en.adjectives }
func (en *english) Adverb() []string    { return en.adverbs }
func (en *english) Verb() []string      { return en.verbs }
