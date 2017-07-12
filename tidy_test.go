package tidy_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/knzm/tidy"
	"github.com/knzm/tidy/sample"
)

//go:generate go-bindata -pkg sample -o sample/bindata.go sample/

var (
	SampleInput  = string(sample.MustAsset("sample/input.txt"))
	SampleOutput = string(sample.MustAsset("sample/output.txt"))
)

func IsDiffsEmpty(diffs []diffmatchpatch.Diff) bool {
	for _, diff := range diffs {
		if diff.Type != diffmatchpatch.DiffEqual {
			return false
		}
	}
	return true
}

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
	if !IsDiffsEmpty(diffs) {
		t.Error("The expected and actual data did not match.")
		t.Log(dmp.DiffPrettyText(diffs))
	}
}
