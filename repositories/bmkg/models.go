package bmkg

// Wilayah model for wilayah
type Wilayah struct {
	ID   string `json:"id"`
	City string `json:"kota"`
}

// Weather model for cuaca
type Weather struct {
	Cuaca        string `json:"cuaca"`
	TemperatureC string `json:"tempC"`
	JamCuaca     string `json:"jamCuaca"`
}
