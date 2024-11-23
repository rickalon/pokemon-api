package data

type Pokemon struct {
	Name   string  `json:"name"`
	Weight int     `json:"weight"`
	Sprite Sprites `json:"sprites"`
}

type Sprites struct {
	ImageURL string `json:"front_default"`
}

func emptyPokemon() *Pokemon {
	return &Pokemon{}
}

func NewPokemon() *Pokemon {
	dflt := emptyPokemon()
	return dflt
}
