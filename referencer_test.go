package bhlinker

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/gnames/bhlnames/config"
	bhln "github.com/gnames/bhlnames/domain/entity"
	"github.com/gnames/gnfmt"
)

type MockReferencer struct{}

func (mr MockReferencer) Refs(name string, ops ...config.Option) (*bhln.NameRefs, error) {
	mocks := loadOutputMocks()
	if res, ok := mocks[name]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("Unknown name '%s'", name)
}

func loadNamesMock() []string {
	mocks := loadOutputMocks()
	res := make([]string, len(mocks))
	count := 0
	for k := range mocks {
		res[count] = k
		count++
	}
	return res
}

func loadOutputMocks() map[string]*bhln.NameRefs {
	enc := gnfmt.GNjson{}
	var res map[string]*bhln.NameRefs
	path := filepath.Join("testdata", "referencer-mock.json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = enc.Decode(data, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func TestRefs(t *testing.T) {
	mr := MockReferencer{}
	data, _ := mr.Refs("something")
	if data != nil {
		t.Error("it should not find name 'somthing'")
	}
	data, err := mr.Refs("Licaria simulans")
	if err != nil {
		t.Error("Error for 'Licaria simulans' should be nil")
	}
	if data.NameString != "Licaria simulans" {
		t.Errorf("Wrong name '%s'", data.NameString)
	}
	if data.ReferenceNumber != 5 {
		t.Errorf("Wrong number of refs '%d'", len(data.References))
	}
}
