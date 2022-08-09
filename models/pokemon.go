package models

//Pokemon datos del pokemon
type Pokemon struct {
	Orden   int     `json:"order"`
	Name    string  `json:"name"`
	Weight  int     `json:"weight"`
	Height  int     `json:"height"`
	Sprites Sprite  `json:"sprites"`
	Types   []Tipos `json:"types"`
}

//Pokemons Lista de los pokemones
type Pokemons []*Pokemon

//Sprite url de la imagen por defecto del pokemon
type Sprite struct {
	Front_default string `json:"front_default"`
}
type Tipos struct {
	Tipo Tipo `json:"type"`
}
type Tipo struct {
	Name string `json:"name"`
}
