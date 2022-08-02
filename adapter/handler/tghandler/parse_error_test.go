package tghandler

import (
	"fmt"
	"testing"

	"github.com/artnoi43/fngobot/adapter/parse"
)

func TestFormString(t *testing.T) {
	expected := make(map[parse.ParseError]string)
	// Construct expected values
	for parseError, err := range parse.ErrMsgs {
		errMsg := fmt.Sprintf(
			"failed to parse command:\n%s",
			err.Error(),
		)

		expected[parseError] = errMsg
		actual := formString(parseError)
		if actual != expected[parseError] {
			t.Fatal("strings not matched")
		}
	}
}
