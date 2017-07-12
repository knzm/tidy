package tidy_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/knzm/tidy"
)

var SampleInput = `
4
132
1000
7
111111111111111110
`[1:]

var SampleOutput = `
Case #1: 129
Case #2: 999
Case #3: 7
Case #4: 99999999999999999
`[1:]

func TestSample(t *testing.T) {
	r := strings.NewReader(SampleInput)
	var buf bytes.Buffer

	ns, err := tidy.ParseInput(r)
	if err != nil {
		t.Fatal(err)
	}
	for i, n := range ns {
		ans := tidy.Solve(n)
		tidy.PrintOutput(&buf, i, n, ans)
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(buf.String(), SampleOutput, false)
	if len(diffs) > 0 {
		t.Error("The expected and actual data did not match.")
		t.Log(dmp.DiffPrettyText(diffs))
	}
}
