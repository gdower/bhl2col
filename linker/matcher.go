package linker

import (
	"fmt"
	"strconv"

	"github.com/gdower/bhlinker/datamatcher"
	"github.com/gdower/bhlinker/domain/entity"
	rfs "github.com/gnames/bhlnames/refs"
)

func bestMatchBHL(input entity.Input, bhlRefs []*rfs.Reference) entity.Output {
	year := input.Reference.Year
	refBest, score := bestBHLReference(bhlRefs, year)
	return output(input, refBest, score)
}

func bestBHLReference(bhlRefs []*rfs.Reference, year string) (*rfs.Reference, entity.Score) {
	refYear, scoreYear := matchYear(year, bhlRefs)
	refAnnot, scoreAnnot, scoreYearComposite := matchAnnot(year, bhlRefs)
	scoreComposite := scoreAnnot + scoreYearComposite
	if scoreYear+scoreComposite == 0 {
		return nil, entity.Score{}
	}
	refBest := refYear
	scoreBest := scoreYear
	if refBest == nil {
		refBest = refAnnot
		scoreBest = scoreComposite
	} else if refAnnot != nil {
		if scoreComposite > 0 && refBest.PageID != refAnnot.PageID {
			if scoreYear < scoreComposite {
				refBest = refAnnot
				scoreBest = scoreComposite
				scoreYear = scoreYearComposite
			}
		}
	}
	score := entity.Score{Overall: scoreBest, Annot: scoreAnnot, Year: scoreYear}
	return refBest, score
}

func output(input entity.Input, refBest *rfs.Reference, score entity.Score) entity.Output {
	if refBest == nil {
		return entity.Output{}
	}

	link := entity.BHLlink{
		Link: refBest.URL,
	}
	res := entity.Output{
		InputID:      input.ID,
		Name:         input.Name,
		BHLlink:      link,
		Score:        score,
		AnnotNomen:   refBest.AnnotNomen,
		EditDistance: refBest.EditDistance,
	}
	return res
}

func matchYear(refYear string, refs []*rfs.Reference) (*rfs.Reference, float32) {
	var refBest *rfs.Reference
	var score, scoreBest float32
	for _, r := range refs {
		yr, err := strconv.Atoi(refYear)
		if err != nil {
			fmt.Printf("Weird year: %s", refYear)
			continue
		}
		score = datamatcher.YearScore(yr, r)
		if score > scoreBest {
			refBest = r
			scoreBest = score
		}
	}
	return refBest, scoreBest
}

func matchAnnot(refYear string, refs []*rfs.Reference) (*rfs.Reference, float32, float32) {
	var refBest *rfs.Reference
	var scoreAnnot, score float32
	for _, r := range refs {
		score = datamatcher.AnnotScore(r)
		if score > scoreAnnot {
			refBest = r
			scoreAnnot = score
		}
	}
	var scoreYear float32 = 0
	if scoreAnnot > 0 {
		yr, err := strconv.Atoi(refYear)
		if err == nil {
			scoreYear = datamatcher.YearScore(yr, refBest)
		}
	}
	return refBest, scoreAnnot, scoreYear
}
