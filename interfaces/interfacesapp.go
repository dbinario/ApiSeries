package interfaces

type Recuperar struct {
	IdSerie   int `json:"id_serie"`
	Temporada int `json:"temporada"`
}

type EpisodioVisto struct {
	IdSerie   int `json:"id_serie"`
	Temporada int `json:"temporada"`
	Episodio  int `json:"episodio"`
}

type EliminarSerie struct {
	IdSerie int `json:"id_serie"`
}

type EliminarTemporada struct {
	IdSerie   int `json:"id_serie"`
	Temporada int `json:"temporada"`
}
