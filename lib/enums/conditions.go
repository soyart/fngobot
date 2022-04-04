package enums

type Condition string

const (
	Lt Condition = "LT"
	Gt Condition = "GT"
)

var (
	validConditions = []Condition{
		Lt,
		Gt,
	}
)
