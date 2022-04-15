package tghandler

import (
	"fmt"
	"testing"

	"github.com/artnoi43/fngobot/adapter/parse"
)

func TestFormString(t *testing.T) {
	expected := make(map[parse.ParseError]string)
	// Construct expected values
	for parseError, errMsg := range parse.ErrMsg {
		errMsg = fmt.Sprintf(
			"%s\n%s",
			"failed to parse command:",
			errMsg,
		)
		expected[parseError] = errMsg

		actual := formString(parseError)
		if actual != expected[parseError] {
			t.Fatal("strings not matched")
		}
	}
}
