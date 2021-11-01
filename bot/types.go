package bot

type Quote struct {
	Last float64
	Bid  float64
	Ask  float64
}

type Security struct {
	Tick string
	Src int
	quote Quote
}