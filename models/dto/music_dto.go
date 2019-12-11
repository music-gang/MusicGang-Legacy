package dto

// MusicDTO - Definisce il DTO per l'oggetto music
type MusicDTO struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Genre    string `json:"genre"`
	Notes    string `json:"notes"`
	Prompter string `json:"prompter"`
}
