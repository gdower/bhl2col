package datamatcher

import (
	"math"
	"time"

	bhln "github.com/gnames/bhlnames/domain/entity"
)

func YearScore(yearInput int, ref *bhln.Reference) float32 {
	var score float32 = 1
	yearPart, itemYearStart, itemYearEnd, titleYearStart, titleYearEnd := getRefYears(ref)

	if yearPart > 0 {
		return score * yearNear(yearInput, yearPart)
	}
	var score1, score2 float32
	item := int(itemYearStart+itemYearEnd) > 0
	title := int(titleYearStart+titleYearEnd) > 0
	if item || (!item && !title) {
		score1 = yearBetween(yearInput, itemYearStart, itemYearEnd)
	}
	if title {
		score2 = yearBetween(yearInput, titleYearStart, titleYearEnd)
	}

	if score1 > score2 {
		return score1
	}
	return score2
}

func invalidYear(year int) bool {
	return year < 1740 || year > (time.Now().Year()+2)
}

func yearNear(year1, year2 int) float32 {
	if invalidYear(year1) {
		return 0
	}

	coef := 0.7
	dif := math.Abs(float64(year1) - float64(year2))
	if dif > 10 {
		return 0
	}
	return float32(math.Pow(float64(coef), dif))
}

func yearBetween(year, yearMin, yearMax int) float32 {
	if invalidYear(year) {
		return 0
	}
	if yearMin == 0 && yearMax == 0 {
		return 0.01
	}

	if yearMax < yearMin && yearMax != 0 {
		return 0.01
	}

	if yearMax == 0 {
		return yearNear(year, yearMin)
	}

	if !(year <= yearMax && year >= yearMin) {
		return 0
	}

	return yearNear(year, yearMax)
}

func getRefYears(ref *bhln.Reference) (int, int, int, int, int) {
	var yearPart int
	if ref.YearType == "Part" {
		yearPart = ref.YearAggr
	}
	return yearPart, ref.ItemYearStart, ref.ItemYearEnd, ref.TitleYearStart, ref.TitleYearEnd
}
