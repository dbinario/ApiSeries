package interfaces

type EstadoSeries struct {
	NumeroSeries    int `json:"numero_series"`
	EpisodiosVistos int `json:"episodios_vistos"`
	EpisodiosPorVer int `json:"episodios_por_ver"`
	Series          []struct {
		IDserie         int    `json:"id_serie"`
		NombreSerie     string `json:"nombre_serie"`
		NumeroEpisodios int    `json:"numero_episodios"`
		EpisodiosVistos int    `json:"episodios_vistos"`
		EpisodiosPorVer int    `json:"episodios_por_ver"`
	} `json:"series"`
}
