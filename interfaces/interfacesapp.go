package interfaces

type Recuperar struct {
	IdSerie   int `json:"id_serie"`
	Temporada int `json:"temporada"`
}

type CapituloVisto struct {
	IdSerie   int `json:"id_serie"`
	Temporada int `json:"temporada"`
	Capitulo  int `json:"capitulo"`
}

type EliminarSerie struct {
	IdSerie int `json:"id_serie"`
}

type EliminarTemporada struct {
	IdSerie   int `json:"id_serie"`
	Temporada int `json:"temporada"`
}
