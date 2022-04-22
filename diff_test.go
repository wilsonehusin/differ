package differ_test

import (
	"testing"

	"go.husin.dev/differ"
)

func TestDiffNone(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	d := differ.Diff[int](slice, slice)

	p := differ.NewPrinter(differ.StrconvInt)
	p.Add(d...)

	if len(d) == 1 {
		if p.String() != "" {
			t.Fatalf("expected empty string diff, got:\n%s", p.String())
		}
	} else {
		t.Log("\n" + p.String())
		t.Fatalf("expected 1 delta, found %d", len(d))
	}
}

func TestDiffInts(t *testing.T) {
	left := []int{1, 2, 3, 1, 2, 2, 1}
	right := []int{3, 2, 3, 1, 1, 3}

	expectDiffString := `1c0,1
< 1
---
> 3
1,4e1,4
  2
  3
  1
5,6d4
< 2
< 2
6,7e4,5
  1
8a6
> 3
`

	d := differ.Diff[int](left, right)

	p := differ.NewPrinter(differ.StrconvInt, differ.PrintUnified, differ.PrintNumber)
	p.Add(d...)

	if len(d) == 5 {
		if received := p.String(); received != expectDiffString {
			t.Logf("expected (%d):\n%s", len(expectDiffString), expectDiffString)
			t.Logf("received (%d):\n%s", len(received), received)
			t.Fatalf("received unexpected diff content")
		}
	} else {
		t.Log(p.String())
		t.Fatalf("expected 5 deltas, received %d", len(d))
	}
}
