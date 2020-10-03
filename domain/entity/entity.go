package entity

type NameType int

const (
	FullString NameType = iota
	CanonicalForm
)

func (st NameType) String() string {
	if int(st) == 0 {
		return "FullString"
	} else {
		return "CanonicalForm"
	}
}

type Input struct {
	ID        int `json:"id"`
	Name      `json:"name"`
	Reference `json:"reference"`
}

type Name struct {
	Name     string `json:"name"`
	Authors  string `json:"authors"`
	Year     string `json:"year"`
	NameType `json:"nameType"`
}

type Reference struct {
	Authors string `json:"authors"`
	Journal string `json:"journal"`
	Volume  string `json:"volume"`
	Pages   string `json:"pages"`
	Year    string `json:"year"`
}

type Output struct {
	InputID      int    `json:"id"`
	AnnotNomen   string `json:"annotNomen"`
	EditDistance int    `json:"editDistance"`
	Error        error  `json:"error"`
	Name         `json:"name"`
	BHLlink      `json:"linkBHL"`
	Score        `json:"score"`
}

type BHLlink struct {
	Link          string `json:"link"`
	PageImageLink string `json:"pageImageLink"`
}

type Score struct {
	Overall float32 `json:"overall"`
	Annot   float32 `json:"annot"`
	Year    float32 `json:"year"`
}
