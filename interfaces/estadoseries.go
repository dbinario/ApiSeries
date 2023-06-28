package interfaces

type EstadoSeries struct {
	NumeroSeries    int `json:"numero_series"`
	CapitulosVistos int `json:"capitulos_vistos"`
	CapitulosPorVer int `json:"capitulos_por_ver"`
	Series          []struct {
		IDserie         int    `json:"id_serie"`
		NombreSerie     string `json:"nombre_serie"`
		NumeroCapitulos int    `json:"numero_capitulos"`
		CapitulosVistos int    `json:"capitulos_vistos"`
		CapitulosPorVer int    `json:"capitulos_por_ver"`
	} `json:"series"`
}
