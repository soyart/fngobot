package bot

// Quote is a generic struct to store quote infos from different sources
type Quote struct {
	Last float64
	Bid  float64
	Ask  float64
}

// Security is a struct storing info about how to get the quote. It also has Quote embedded
type Security struct {
	Tick  string
	Src   int
	quote Quote
}
