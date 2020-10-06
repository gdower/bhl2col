package datamatcher

import (
	"math"
	"testing"

	bhln "github.com/gnames/bhlnames/domain/entity"
)

func TestAnnotScore(t *testing.T) {

	type data struct {
		name       string
		matchName  string
		annotation annotation
		score      float32
	}

	dataArray := []data{
		//  Name
		// 1 MatchName
		// 2 Annotation
		// 3
		{"Aus bus", "Aus bus", spNov, 1.0},
		{"Aus bus cus", "Aus bus cus", spNov, 0.0},
		{"Aus bus cus", "Aus bus cus", subsNov, 1.0},
		{"Aus bus", "Aus bus cus", spNov, 0.1},
		{"Aus bus cus", "Aus bus", spNov, 0.6},
		{"Aus bus Ower", "Bus cus Mozzherin", spNov, 1.0},
		{"Aus (Bus) cus", "Aus cus", spNov, 1.0},
		{"Aus bus", "Aus bus", subsNov, 0},
		{"Aus bus", "Aus bus cus", subsNov, 0.4},
		{"Aus bus cus", "Aus bus", subsNov, 0.1},
		{"Aus bus", "Aus bus", combNov, 1.0},
		{"Aus bus cus", "Aus bus cus", combNov, 1.0},
		{"Aus bus", "Aus bus cus", combNov, 0.6},
		{"Aus", "Aus", combNov, 0},
		{"Aus bus cus", "Aus", noAnnot, 0.0},
		{"Aus bus", "Aus bus", noAnnot, 0.0},
		{"Aus virus", "Bus cus", spNov, 0.0},
	}

	for _, d := range dataArray {
		testRef := bhln.Reference{
			Name:       d.name,
			MatchName:  d.matchName,
			AnnotNomen: d.annotation.String(),
		}
		result := AnnotScore(&testRef)
		if math.Abs(float64(result)-float64(d.score)) > 0.0001 {
			t.Errorf("Wrong score for AnnotScore(%#v) %f, %f", testRef, result, d.score)
		}

	}
}
