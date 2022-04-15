package config

import "testing"

func TestParsePath(t *testing.T) {
	defaultPath := "$HOME/.config/fngobot/config.yml"
	anotherPath := "$HOME/fngobot/config.json"
	rootPath := "/etc/fngobot/config.yml"
	expected := map[string]Location{
		defaultPath: {
			Dir:  "$HOME/.config/fngobot/",
			Name: "config",
			Ext:  "yml",
		},
		anotherPath: {
			Dir:  "$HOME/fngobot/",
			Name: "config",
			Ext:  "json",
		},
		rootPath: {
			Dir:  "/etc/",
			Name: "config",
			Ext:  "json",
		},
	}

	tests := []string{defaultPath, anotherPath}
	for _, test := range tests {
		actual := *ParsePath(test)
		if actual != expected[test] {
			t.Logf(
				"Expected:\n%v\nActual:\n%v\n",
				expected[test],
				actual,
			)
			t.Fatal("Not matched")
		}
	}
}
