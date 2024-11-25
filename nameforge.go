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

package nameforge

import (
	_ "embed"
	"github.com/MrDiver/nameforge/common"
	formatter "github.com/MrDiver/nameforge/format"
	"github.com/MrDiver/nameforge/globals"
	"github.com/MrDiver/nameforge/language"
	"github.com/MrDiver/nameforge/universe"
	"strings"
)

var (
	DEFAULTUNIVERSE globals.Universe = &universe.STARWARS
	DEFAULTLANGUAGE globals.Language = &language.ENGLISH
)

func SetFormat(f formatter.FormatFunc) {
	formatter.DEFAULTFORMAT = f
}

func SetJoin(f formatter.JoinFunc) {
	formatter.DEFAULTJOIN = f
}

func SetSplit(f formatter.SplitFunc) {
	formatter.DEFAULTSPLIT = f
}

func SetUniverse(u globals.Universe) {
	DEFAULTUNIVERSE = u
}

func SetLanguage(l globals.Language) {
	DEFAULTLANGUAGE = l
}

func AvailableUniverses() map[string]globals.Universe {
	return globals.AvailableUniverses
}

func AvailableLanguages() map[string]globals.Language {
	return globals.AvailableLanguages
}

func NameU(u globals.Universe) string {
	name := strings.TrimSpace(common.SelectRandom(u.Names()))
	return formatter.DEFAULTFORMAT(name)
}

func Name() string {
	return NameU(DEFAULTUNIVERSE)
}

func SentenceC(format string) string {
	parts := formatter.SCOMMA(format)
	words := make([]string, 0)
	for _, part := range parts {
		var word string
		switch strings.TrimSpace(strings.ToUpper(part)) {
		case "N":
			word = common.SelectRandom(DEFAULTLANGUAGE.Noun())
		case "V":
			word = common.SelectRandom(DEFAULTLANGUAGE.Verb())
		case "VING":
			word = common.SelectRandom(DEFAULTLANGUAGE.Verb()) + "ing"
		case "ADJ":
			word = common.SelectRandom(DEFAULTLANGUAGE.Advective())
		case "ADV":
			word = common.SelectRandom(DEFAULTLANGUAGE.Adverb())
		case "NAME":
			word = Name()
		default:
			continue
		}
		words = append(words, word)
	}
	return formatter.DEFAULTJOIN(words)
}

func DescribedName() string {
	return SentenceC("adv,adj,name")
}

func DescribedNoun() string {
	return SentenceC("adv,adj,n")
}
