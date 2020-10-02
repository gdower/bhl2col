package usecase

import (
	"github.com/gdower/bhlinker/domain/entity"
	"github.com/gnames/bhlnames/refs"
)

// Plugger provides API to the module.
type Plugger interface {
	// GetLink takes an input with a name-string and optionally the nomenclatural reference data
	// for the name-string and returns back BHL references filtered by scoring algorithms.
	// The references are the best attempt to find first nomenclatural descriptions
	// for a names in BHL.
	GetLink(input entity.Input) entity.Output
	// GetLinks takes a stream of name-strings and returns a stream of references in BHL.
	GetLinks(chIn <-chan entity.Input, chOut chan<- entity.Output)
}

// Referencer allows us to inverse dependency to BHLnames. It provides signatures
// to BHLnames methods needed for functionality of BHLinker.
type Referencer interface {
	// RefsStream takes a stream of name-strings, returns back a stream of
	// references found in BHL for the names-strigns.
	RefsStream(chIn <-chan string, chOut chan<- *refs.RefsResult)
	// Refs takes a name-string and returns an BHLnames' output that contains
	// found BHL references for the name-string.
	Refs(string) (*refs.Output, error)
}
