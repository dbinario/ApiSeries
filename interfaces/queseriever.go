package interfaces

type QueSerieVer struct {
	IdSerie         int    `json:"id_serie"`
	Temporada       int    `json:"temporada"`
	Episodio        int    `json:"episodio"`
	NombreSerie     string `json:"nombre_serie"`
	NombreTemporada string `json:"nombre_temporada"`
	NombreEpisodio  string `json:"nombre_episodio"`
	Resumen         string `json:"resumen"`
}
