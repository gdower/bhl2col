package entity

import (
	bhln "github.com/gnames/bhlnames/domain/entity"
)

type Input struct {
	ID        string `json:"id"`
	Name      `json:"name"`
	Reference `json:"reference"`
}

type Name struct {
	NameString string `json:"nameString,omitempty"`
	Canonical  string `json:"canonical,omitempty"`
	Authorship string `json:"authorship,omitempty"`
	Year       string `json:"year,omitempty"`
}

type Reference struct {
	RefString string `json:"refString,omitempty"`
	Authors   string `json:"authors,omitempty"`
	Journal   string `json:"journal,omitempty"`
	Volume    string `json:"volume,omitempty"`
	Pages     string `json:"pages,omitempty"`
	Year      string `json:"year,omitempty"`
}

type Output struct {
	InputID      string          `json:"inputId"`
	InputName    Name            `json:"inputName"`
	InputRef     Reference       `json:"inputRef,omitempty"`
	OutputName   string          `json:"outputName,omitempty"`
	EditDistance int             `json:"editDistance,omitempty"`
	Error        error           `json:"error,omitempty"`
	BHLref       *bhln.Reference `json:"referenceBHL"`
	Score        `json:"score"`
}

type Score struct {
	Overall float32 `json:"overall"`
	Annot   float32 `json:"annot"`
	Year    float32 `json:"year"`
}
