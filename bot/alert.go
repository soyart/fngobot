package bot

import "github.com/artnoi43/fngobot/enums"

// Alert struct stores info about price alerts
type Alert struct {
	Security
	Condition int
	QuoteType int
	Target    float64
}

// Match sends a truthy value into matched channel
// if the specified market condition is matched.
func (a *Alert) Match(matched chan<- bool) {
	q, err := a.Security.Quote()
	if err != nil {
		matched <- false
	}

	var p float64
	switch a.QuoteType {
	case enums.Bid:
		p = q.Bid
	case enums.Ask:
		p = q.Ask
	case enums.Last:
		p = q.Last
	}

	switch a.Condition {
	case enums.Lt:
		if p <= a.Target {
			matched <- true
		}
	case enums.Gt:
		if p >= a.Target {
			matched <- true
		}
	}
}
