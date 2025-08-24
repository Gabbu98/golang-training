package models

type PokemonDetails struct {
	Abilities      []Abilities `json:"abilities"`
	BaseExperience int         `json:"base_experience"`
	Weight         int         `json:"weight"`
}
