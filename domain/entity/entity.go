package entity

type NameType int

const (
	FullString NameType = iota
	CanonicalForm
)

type AnnotNomen int

const (
	NoAnnot AnnotNomen = iota
	SpNov
	CombNov
	SubspNov
)


func (st NameType) String() string {
	if int(st) == 0 {
		return "FullString"
	} else {
		return "CanonicalForm"
	}
}

type Name struct {
	Name string
	Authors string
	Year int
	NameType
}

type Reference struct {
	Authors string
	Journal string
	Volume string
	Pages string
	Year int
}


type BHLink stuct {
	Link string
	PageImageLink string
}

type Score struct {
	Overall float32
	Annot float32
	Year float32
}

type Input struct {
	ID int
	Name
	Reference
}

type Output struct {
	InputID int
	Name
	BHLink
	Score
	AnnotNomen
	EditDistance uint
	Error error
}

