package bhlinker

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
	"testing"

	entity "github.com/gdower/bhlinker/ent"
	"github.com/gnames/gnfmt"
)

func loadInputMock() (map[string]entity.Input, error) {
	enc := gnfmt.GNjson{}
	var res map[string]entity.Input
	path := filepath.Join("testdata", "input-mock.json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return res, err
	}
	err = enc.Decode(data, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func loadOutputMock() (map[string]entity.Output, error) {
	enc := gnfmt.GNjson{}
	var res map[string]entity.Output
	path := filepath.Join("testdata", "output-mock.json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return res, err
	}
	err = enc.Decode(data, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func data() (BHLinker, map[string]entity.Input, map[string]entity.Output) {
	inputs, err := loadInputMock()
	if err != nil {
		log.Fatalf("cannot load mock inputs: %s", err)
	}
	outputs, err := loadOutputMock()
	if err != nil {
		log.Fatalf("cannot load mock outputs: %s", err)
	}

	mr := MockReferencer{}
	linker := NewBHLinker(mr, 4)
	return linker, inputs, outputs
}

func TestGetLink(t *testing.T) {
	l, inputs, outputs := data()
	for k, v := range inputs {
		out, err := l.GetLink(v)
		if err != nil {
			t.Errorf("cannot get link for '%s': %s", k, err)
		}
		if out.Score.Overall != outputs[k].Score.Overall {
			t.Errorf("scores do not match for %s: %0.2f vs %0.2f",
				k, out.Score.Overall, outputs[k].Score.Overall)
		}
		if out.BHLref.URL != outputs[k].BHLref.URL {
			t.Errorf("BHL links do not match for %s: %s vs %s",
				k, out.BHLref.URL, outputs[k].BHLref.URL)
		}
	}
}

func TestGetLinks(t *testing.T) {
	l, inputs, outputs := data()
	chIn := make(chan entity.Input)
	chOut := make(chan entity.Output)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for _, v := range inputs {
			chIn <- v
		}
		close(chIn)
	}()

	go func() {
		defer wg.Done()
		for output := range chOut {
			name := output.InputName.Canonical
			if output.Error != nil {
				t.Errorf("cannot get link for '%s': %s", name, output.Error)
			}
			if output.Score.Overall != outputs[name].Score.Overall {
				t.Errorf("scores do not match for %s: %0.2f vs %0.2f",
					name, output.Score.Overall, outputs[name].Score.Overall)
			}
			if output.BHLref.URL != outputs[name].BHLref.URL {
				t.Errorf("BHL links do not match for %s: %s vs %s",
					name, output.BHLref.URL, outputs[name].BHLref.URL)
			}
		}
	}()
	l.GetLinks(chIn, chOut)
	wg.Wait()
}
