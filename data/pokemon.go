package data

type Pokemon struct {
	Name   string  `json:"name"`
	Weight int     `json:"weight"`
	Sprite Sprites `json:"sprites"`
}

type Sprites struct {
	ImageURL string `json:"front_default"`
}

func NewPokemon() *Pokemon {
	return &Pokemon{}
}
