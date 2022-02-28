package data

type Data struct {
	Schemes  []Scheme  `json:"schemes"`
	Modulars []Modular `json:"modulars"`
	Heroes   []Hero    `json:"heroes"`
	Aspects  []string  `json:"aspects"`
}

type Scheme struct {
	SchemeId    int      `json:"schemeId"`
	SchemeName  string   `json:"schemeName"`
	Villain     string   `json:"villain"`
	VillainSet  string   `json:"villainSet"`
	Recommended []string `json:"recommended"`
	Required    []string `json:"required"`
}

type Modular struct {
	ModularId   int    `json:"modularId"`
	ModularName string `json:"modularName"`
}

type Hero struct {
	HeroId     int    `json:"heroId"`
	HeroName   string `json:"heroName"`
	HeroAspect int    `json:"aspect"`
}
