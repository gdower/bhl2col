package linker

import (
	"github.com/gdower/bhlinker/domain/entity"
)

type Linker struct {
}

func (l Linker) GetLink(input entity.Input) entity.Output {
	res := entity.Output{}
	return res
}

func (l Linker) GetLinks(chIn <-chan entity.Input, chOut chan<- entity.Output) {
}
