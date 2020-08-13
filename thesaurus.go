package thesaurus

// Thesaurus is a intetface to use thesaurus service.
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
