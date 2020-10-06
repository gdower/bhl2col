package datamatcher

import (
	bhln "github.com/gnames/bhlnames/domain/entity"
	"gitlab.com/gogna/gnparser"
)

type annotation int

const (
	noAnnot annotation = iota
	spNov
	subsNov
	combNov
)

func NewAnnot(annot string) annotation {
	annotations := map[string]annotation{
		"NO_ANNOT":  noAnnot,
		"SP_NOV":    spNov,
		"SUBSP_NOV": subsNov,
		"COMB_NOV":  combNov,
	}
	if a, ok := annotations[annot]; ok {
		return a
	}
	return noAnnot
}

func (a annotation) String() string {
	switch int(a) {
	case 1:
		return "SP_NOV"
	case 2:
		return "SUBSP_NOV"
	case 3:
		return "COMB_NOV"
	default:
		return "NO_ANNOT"
	}

}

// NO_ANNOT = 1
// SP_NOV
// f:sp v:sp = 1
// f:sp v:gen = 0
// f:sp v:ssp = 0.1
// f:ssp v:sp = 0.6
// f:ssp v:gen = 0
// f:gen v:gen = 0
// SUBSP_NOV
// f:ssp v:ssp = 1
// f:ssp v:sp = 0
// f:ssp v:gen = 0
// f:sp v:ssp = 0.6
// f:gen v:gen = 0
// COMB_NOV
// f:sp v:sp = 1
// f:ssp v:ssp = 1
// f:ssp v:sp = 0.4
// f:sp v:ssp = 0.1
// f:gen v:gen = 0

func AnnotScore(ref *bhln.Reference) float32 {
	annot := NewAnnot(ref.AnnotNomen)
	cardName, cardMatchName := cardinality(ref)
	if cardName == 0 || cardMatchName == 0 {
		return 0.0
	}
	switch annot {
	case spNov:
		switch {
		case cardName == 2 && cardMatchName == 2:
			return 1.0
		case cardName == 2 && cardMatchName == 3:
			return 0.1
		case cardName == 3 && cardMatchName == 2:
			return 0.6
		default:
			return 0
		}
	case subsNov:
		switch {
		case cardName == 3 && cardMatchName == 3:
			return 1.0
		case cardName == 3 && cardMatchName == 2:
			return 0.1
		case cardName == 2 && cardMatchName == 3:
			return 0.4
		default:
			return 0
		}
	case combNov:
		switch {
		case cardName == 2 && cardMatchName == 2:
			return 1.0
		case cardName == 3 && cardMatchName == 3:
			return 1.0
		case cardName == 2 && cardMatchName == 3:
			return 0.6
		}
	case noAnnot:
		return 0.0
	}
	return 0
}

func cardinality(ref *bhln.Reference) (int32, int32) {
	gnp := gnparser.NewGNparser()
	n := gnp.ParseToObject(ref.Name)
	mn := gnp.ParseToObject(ref.MatchName)
	return n.Cardinality, mn.Cardinality
}
