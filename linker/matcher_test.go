package linker

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gdower/bhlinker/domain/entity"
	"github.com/gnames/gnames/lib/encode"
)

func loadInputMock() (map[string]entity.Input, error) {
	enc := encode.GNjson{}
	var res map[string]entity.Input
	path := filepath.Join("..", "testdata", "input-mock.json")
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
	enc := encode.GNjson{}
	var res map[string]entity.Output
	path := filepath.Join("..", "testdata", "output-mock.json")
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

func TestGetLink(t *testing.T) {
	inputs, err := loadInputMock()
	if err != nil {
		t.Errorf("cannot load mock inputs: %s", err)
	}
	outputs, err := loadOutputMock()
	if err != nil {
		t.Errorf("cannot load mock outputs: %s", err)
	}

	mr := MockReferencer{}
	l := NewLinker(mr)
	for k, v := range inputs {
		out, err := l.GetLink(v)
		if err != nil {
			t.Errorf("cannot get link for '%s': %s", k, err)
		}
		if out.Score.Overall != outputs[k].Score.Overall {
			t.Errorf("scores do not match for %s: %0.2f vs %0.2f",
				k, out.Score.Overall, outputs[k].Score.Overall)
		}
		if out.BHLlink.Link != outputs[k].BHLlink.Link {
			t.Errorf("scores do not match for %s: %s vs %s",
				k, out.BHLlink.Link, outputs[k].BHLlink.Link)
		}
	}
}
