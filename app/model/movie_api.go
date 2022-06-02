package model

//MovieAPI struct
type MovieAPI struct {
	Title    string `json:"title,omitempty" example:"Avengers : Endgame"`
	Year     int    `json:"year,omitempty" example:"2019"`
	Summary  string `json:"summary,omitempty" example:"good movie"`
	Director string `json:"director,omitempty" example:"Joe Russo"`
}
