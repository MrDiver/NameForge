package language

import (
	_ "embed"

	"CodingGhost.de/nameforge/common"
	"CodingGhost.de/nameforge/globals"
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
