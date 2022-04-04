package enums

import (
	"github.com/artnoi43/fngobot/lib/utils"
)

type enum interface {
	comparable
	~string
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

func (t InputCommand) IsValid() bool {
	return isValid(
		t,
		ValidInputCommands,
	)
}

func (t BotType) IsValid() bool {
	return isValid(
		utils.ToUpper(t),
		validBotTypes,
	)
}

func (t Src) IsValid() bool {
	return isValid(
		t,
		validSrc,
	)
}

func (t QuoteType) IsValid() bool {
	return isValid(
		utils.ToUpper(t),
		validQuoteTypes,
	)
}

func (t Condition) IsValid() bool {
	return isValid(
		utils.ToUpper(t),
		validConditions,
	)
}
