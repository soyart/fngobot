package enums

import (
	"github.com/artnoi43/gsl/gslutils"
)

type enum interface {
	comparable
	any
}

// isValid is a generic function for interface enum,
// it iterates through valids and compare value to each elem.
// It is called by IsValid(), which just wraps it.
func isValid[T enum](value T, valids []T) bool {
	for _, valid := range valids {
		if value == valid {
			return true
		}
	}
	return false
}

// IsValid returns whether the given value is present in a valid list.
// Note: InputCommand values are NOT capitalized
func (t InputCommand) IsValid() bool {
	return isValid(
		t,
		ValidInputCommands,
	)
}

// IsValid returns whether the given value is present in a valid list.
// Note: Src values are NOT capitalized
func (t Src) IsValid() bool {
	return isValid(
		t,
		validSrc,
	)
}

// IsValid returns whether the given value is present in a valid list.
// Note: BotType values are capitalized
func (t BotType) IsValid() bool {
	return isValid(
		gslutils.ToUpper(t),
		validBotTypes,
	)
}

// IsValid returns whether the given value is present in a valid list.
// Note: QuoteType values are capitalized
func (t QuoteType) IsValid() bool {
	return isValid(
		gslutils.ToUpper(t),
		validQuoteTypes,
	)
}

// IsValid returns whether the given value is present in a valid list.
// Note: Condition values are capitalized
func (t Condition) IsValid() bool {
	return isValid(
		gslutils.ToUpper(t),
		validConditions,
	)
}
