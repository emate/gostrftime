// Author: Marcin Matlaszek <marcin@matlaszek.pl>
// See the file "LICENSE" for the full license governing this code.

package gostrtime

import (
	"strings"
)

var strtimeMap = map[string]string{
	"%B":  "January",
	"%b":  "Jan",
	"%h":  "Jan",
	"%m":  "01",
	"%_m": " 1",
	"%-m": "1",
	"%Y":  "2006",
	"%y":  "06",
	"%d":  "02",
	"%-d": "2",
	"%H":  "15",
	"%I":  "03",
	"%M":  "04",
	"%S":  "05",
	"%f":  "999999",
	"%z":  "Z0700",
	"%:z": "Z07:00",
	"%Z":  "MST",
	"%p":  "PM",
}

func FormatSample(format string) (sample string) {
	sample = format
	for f, s := range strtimeMap {
		sample = strings.Replace(sample, f, s, -1)
	}
	return sample
}
