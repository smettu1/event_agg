package models

// example data {"id":121509,"market":5773,"price":1.234,"volume":1234.56,"is_buy":true}
type Input struct {
	Id     int     `json:"id,omitempty"`
	Market int     `json:"market,omitempty"`
	Price  float64 `json:"price,omitempty"`
	Volume float64 `json:"volume,omitempty"`
	IsBuy  bool    `json:"is_buy,omitempty"`
}

//Per market raw data and derived metrics
// Lets use float64 for most of agg. metrics.
//https://stackoverflow.com/questions/63358564/why-cant-go-floats-overflow-but-integers-can
type Market struct {
	NumOfElements uint64
	NumOfTrue     uint64
	SumOfVolume   float64
	SumOfPrice    float64
	TotalVol      float64
	MeanPrice     float64
	MeanVol       float64
	VolWeighted   float64
	PercentBuy    float64
}

//Set mean price
func (m *Market) SetMeanPrice() {
	m.MeanPrice = m.SumOfPrice / float64(m.NumOfElements)
}

//Set mean volume
func (m *Market) SetMeanVol() {
	m.MeanVol = m.SumOfVolume / float64(m.NumOfElements)
}

//Set percentage buy
func (m *Market) SetPercent() {
	m.PercentBuy = (float64(m.NumOfTrue) / float64(m.NumOfElements)) * 100
}
