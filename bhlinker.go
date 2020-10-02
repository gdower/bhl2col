package bhlinker

import (
	"github.com/gdower/bhlinker/domain/entity"
	"github.com/gdower/bhlinker/domain/usecase"
	"github.com/gdower/bhlinker/linker"
)

type BHLinker struct {
	usecase.Plugger
}

func NewBHLinker() BHLinker {
	return BHLinker{
		Plugger: linker.Linker{},
	}
}

func (l BHLinker) GetLink(input entity.Input) entity.Output {
	return l.Plugger.GetLink(input)
}

func (l BHLinker) GetLinks(chIn <-chan entity.Input, chOut chan<- entity.Output) {
	l.Plugger.GetLinks(chIn, chOut)
}
