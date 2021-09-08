package linker

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
	"testing"

	bhlent "github.com/gnames/bhlnames/domain/entity"
	"github.com/gnames/gnfmt"
)

type MockReferencer struct{}

func (mr MockReferencer) Refs(name string) (*bhlent.NameRefs, error) {
	mocks := loadOutputMocks()
	if res, ok := mocks[name]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("Unknown name '%s'", name)
}

func (mr MockReferencer) RefsStream(chIn <-chan string, chOut chan<- *bhlent.NameRefs) {
	mocks := loadOutputMocks()
	mocksStream := make(map[string]*bhlent.NameRefs)
	for k, v := range mocks {
		mocksStream[k] = v
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for name := range chIn {
			if res, ok := mocksStream[name]; ok {
				chOut <- res
			} else {
				chOut <- &bhlent.NameRefs{NameString: name}
			}
		}
	}()
	wg.Wait()
	close(chOut)
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

func loadOutputMocks() map[string]*bhlent.NameRefs {
	enc := gnfmt.GNjson{}
	var res map[string]*bhlent.NameRefs
	path := filepath.Join("..", "testdata", "referencer-mock.json")
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

func TestRefsStream(t *testing.T) {
	chIn := make(chan string)
	chOut := make(chan *bhlent.NameRefs)
	var wg sync.WaitGroup
	wg.Add(1)
	mr := MockReferencer{}

	go func() {
		for _, name := range loadNamesMock() {
			chIn <- name
		}
		close(chIn)
	}()
	go func() {
		defer wg.Done()

		for res := range chOut {
			if res == nil {
				t.Error("Refs stream result is empty")
			}
			if len(res.References) == 0 {
				t.Errorf("No references for %s", res.NameString)
			}
		}
	}()
	mr.RefsStream(chIn, chOut)
	wg.Wait()
}
