package interfaces

type ListaSeries struct {
	NumeroSeries    int `json:"numero_series"`
	EpisodiosPorVer int `json:"episodios_por_ver"`
	Episodios       []struct {
		IDserie         int    `json:"id_serie"`
		NombreSerie     string `json:"nombre_serie"`
		NumeroTemporada int    `json:"numero_temporada"`
		NombreTemporada string `json:"nombre_temporada"`
		NumeroEpisodio  int    `json:"numero_episodio"`
		NombreEpisodio  string `json:"nombre_episodio"`
	} `json:"episodios"`
}
