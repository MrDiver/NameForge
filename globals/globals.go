package globals

type Universe interface {
	Names() []string
	String() string
}

type Language interface {
	Noun() []string
	Advective() []string
	Adverb() []string
	Verb() []string
	String() string
}

var (
	AvailableUniverses map[string]Universe = make(map[string]Universe)
	AvailableLanguages map[string]Language = make(map[string]Language)
)
