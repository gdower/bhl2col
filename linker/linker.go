package linker

import (
	"strconv"

	"github.com/gdower/bhlinker/datamatcher"
	entity "github.com/gdower/bhlinker/ent"
	bhln "github.com/gnames/bhlnames/domain/entity"
)

func BestMatchBHL(input entity.Input, nameRefs *bhln.NameRefs) entity.Output {
	bhlRefs := nameRefs.References
	year := input.Reference.Year
	if year == "" {
		year = input.Name.Year
	}
	refBest, score := bestBHLReference(bhlRefs, year)
	return output(input, refBest, score)
}

func bestBHLReference(bhlRefs []*bhln.Reference, year string) (*bhln.Reference, entity.Score) {
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

func output(input entity.Input, refBest *bhln.Reference, score entity.Score) entity.Output {
	if refBest == nil {
		return entity.Output{InputID: input.ID, InputName: input.Name}
	}

	res := entity.Output{
		InputID:      input.ID,
		InputName:    input.Name,
		InputRef:     input.Reference,
		BHLref:       refBest,
		Score:        score,
		EditDistance: refBest.EditDistance,
	}
	return res
}

func matchYear(refYear string, refs []*bhln.Reference) (*bhln.Reference, float32) {
	yr, err := strconv.Atoi(refYear)
	if err != nil {
		yr = 0
	}
	var refBest *bhln.Reference
	var score, scoreBest float32
	for _, r := range refs {
		score = datamatcher.YearScore(yr, r)
		if score > scoreBest {
			refBest = r
			scoreBest = score
		}
	}
	return refBest, scoreBest
}

func matchAnnot(refYear string, refs []*bhln.Reference) (*bhln.Reference, float32, float32) {
	var refBest *bhln.Reference
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
