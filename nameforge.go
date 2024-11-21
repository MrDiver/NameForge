package nameforge

import (
	"CodingGhost.de/nameforge/common"
	formatter "CodingGhost.de/nameforge/format"
	"CodingGhost.de/nameforge/globals"
	"CodingGhost.de/nameforge/language"
	"CodingGhost.de/nameforge/universe"
	_ "embed"
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
