package dockerw

import "testing"

var testTable = []struct {
	name   string
	input  []string
	output Params
	err    error
}{
	{
		"metadata",
		[]string{metadataArg},
		Params{Cmd: "metadata"},
		nil,
	},
	{
		"noop",
		[]string{"binpath", "watch"},
		Params{Cmd: "noop"},
		nil,
	},
	{
		"version",
		[]string{"binpath", "watch", "version"},
		Params{Cmd: "version"},
		nil,
	},
	{
		"start",
		[]string{"binpath", "watch", "start", "dir1:app1", "dir2:app2"},
		Params{"start", []string{"dir1:app1", "dir2:app2"}},
		nil,
	},

	{
		"argsErr",
		[]string{"binpath", "watch", "foo"},
		Params{},
		argsErr,
	},
}

func TestArgs(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Args(tt.input)
			if err != tt.err {
				t.Errorf("want %s, got %s", tt.err, err)
			}
			if !equal(p, tt.output) {
				t.Errorf("want %v, got %v", tt.output, p)
			}
		})
	}
}

// equal compares two Params (I couldn't get reflect.DeepEqual to work)
func equal(a, b Params) bool {
	if a.Cmd != b.Cmd {
		return false
	}
	if len(a.Pairs) != len(b.Pairs) {
		return false
	}
	for i, v := range a.Pairs {
		if v != b.Pairs[i] {
			return false
		}
	}
	return true
}
