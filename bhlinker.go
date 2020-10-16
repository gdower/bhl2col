package bhlinker

import (
	"sync"

	"github.com/gdower/bhlinker/domain/entity"
	"github.com/gdower/bhlinker/domain/usecase"
	"github.com/gdower/bhlinker/linker"
	"github.com/gnames/bhlnames/config"
)

type BHLinker struct {
	usecase.Referencer
	JobsNum int
}

func NewBHLinker(r usecase.Referencer, jobsNum int) BHLinker {
	return BHLinker{Referencer: r, JobsNum: jobsNum}
}

// GetLink takes name-string with its reference data. The reference data is
// expected to be a paper with original nomenclatural description of the
// the name-string. The method tries to find the best BHL match to that
// reference and sends back a BHL reference metadata as well as URL link to
// the reference.
func (l BHLinker) GetLink(input entity.Input) (entity.Output, error) {
	name := input.Name.Canonical
	if name == "" {
		name = input.Name.NameString
	}
	opts := []config.Option{config.OptNoSynonyms(true)}
	refsBHL, err := l.Refs(name, opts...)
	if err != nil {
		out := entity.Output{
			InputID:   input.ID,
			InputName: input.Name,
			InputRef:  input.Reference,
		}
		return out, err
	}
	return linker.BestMatchBHL(input, refsBHL), nil
}

// GetLinks takes a stream of name-strings with their reference data. The
// reference data for each name-string is expected to be a paper with original
// nomenclatural description of the the name-string.  The method tries to find
// the best BHL match to that reference and sends back a BHL reference metadata
// as well as URL link to the reference.
//
// The streams are implemented as channels. This approach allows to work with
// inputs of any size.
func (l BHLinker) GetLinks(chIn <-chan entity.Input, chOut chan<- entity.Output) {
	var wg sync.WaitGroup
	wg.Add(l.JobsNum)
	for i := 0; i < l.JobsNum; i++ {
		go l.worker(chIn, chOut, &wg)
	}
	wg.Wait()
	close(chOut)
}

func (l BHLinker) worker(chIn <-chan entity.Input, chOut chan<- entity.Output,
	wg *sync.WaitGroup) {
	defer wg.Done()
	for input := range chIn {
		output, err := l.GetLink(input)
		output.Error = err
		chOut <- output
	}
}
