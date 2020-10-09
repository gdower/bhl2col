package linker

import (
	"fmt"
	"testing"

	"github.com/gdower/bhlinker/domain/entity"
)

func TestGetLink(t *testing.T) {
	mr := MockReferencer{}
	linker := NewLinker(mr)
	inp := entity.Input{
		ID: 1,
		Name: entity.Name{
			Name:     "Hamotus gracilicornis",
			NameType: entity.CanonicalForm,
		},
	}
	out, err := linker.GetLink(inp)
	if err != nil {
		t.Errorf("Error should be nil for '%s'", inp.Name.Name)
	}
	fmt.Printf("out: %+v\n", out)
}
