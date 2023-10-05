package igboapi

// key= 2df2f3fa-5013-4602-93f7-02cab7be020c

// A term describes a word being searched for.
type Term struct {
	Definitions    []string  `json:"definitions"`
	Dialects       Dialect   `json:"dialects"`
	Pronounciation string    `json:"pronounciation"`
	Tenses         Tense     `json:"tenses"`
	Attributes     Attribute `json:"attributes"`
	RelatedTerms   []string  `json:"relatedTerms"`
	Word           string    `json:"word"`
	Examples       []Example `json:"examples"`
	ID             string    `json:"id"`
	WordClass      string    `json:"wordClass"`
	Nsibidi        string    `json:"nsibidi"`
}

type TermParams struct {
	Keyword string
	Page string
	WordRange string
	Strict bool
	Dialects bool 
	Examples bool
}

type SearchTermResponse struct {
	Definitions []string `json:"definitions"`
	Word string `json:"word"`
	Pronounciation string `json:"pronounciation"`
	Attributes Attribute `json:"attribute"`
	RelatedTerms []string `json:"relatedTerms"`
	ID string `json:"id"`
	WordClass string `json:"wordClass"`
	Nsibidi string `json:"nsibidi"`
	Tenses Tense
}

type Dialect struct {
	Variations     []string `json:"variations"`
	Dialects       []string `json:"dialects"`
	Pronounciation string   `json:"pronounciation"`
	ID             string   `json:"_id"`
	Word           string   `json:"word"`
}

type Tense struct {
	Infinitive        string `json:"infinitive"`
	Imperative        string `json:"imperative"`
	SimplePast        string `json:"simplePast"`
	PresentPassive    string `json:"presentPassive"`
	SimplePresent     string `json:"simplePresent"`
	PresentContinuous string `json:"presentContinuous"`
	Future            string `json:"future"`
}

type Attribute struct {
	IsStandardIgbo bool `json:"isStandardIgbo"`
	IsAccented bool `json:"isAccented"`
	IsSlang bool `json:"isSlang"`
	IsConstructedTerm bool `json:"isContructedTerm"`
	IsBorrowedTerm bool `json:"isBorrowedTerm"`
	IsStem bool `json:"isStem"`
	IsCommon bool `json:"isCommon"`
}

type Example struct {
	Igbo string `json:"igbo"`
	English string `json:"english"`
	Pronounciations []Pronounciation `json:"pronounciations"`
}

type Pronounciation struct {
	Audio string `json:"audio"`
	Speaker string `json:"speaker"`
}