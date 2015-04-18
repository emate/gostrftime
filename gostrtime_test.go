package gostrtime

import "testing"
import "github.com/stretchr/testify/assert"

var formatTests = []struct {
	in  string
	out string
}{
	{
		in:  "%Y-%m-%dT15:%M:%S%z",
		out: "2006-01-02T15:04:05Z0700",
	},
	{
		in:  "%Y-%m-%dT15:%M:%S%:z",
		out: "2006-01-02T15:04:05Z07:00",
	},
	{
		in:  "%Y-%m-%d %H:%M:%S",
		out: "2006-01-02 15:04:05",
	},
}

func TestFormatSample(t *testing.T) {
	for _, test := range formatTests {
		f := FormatSample(test.in)
		assert.Equal(t, test.out, f)
	}
}
