package differ_test

import (
	"testing"

	"go.husin.dev/differ"
)

var (
	simpleDelta = &differ.Delta[string]{
		Left: differ.Subset[string]{
			Index: differ.Range{
				Start: 1,
				End:   2,
			},
			Value: []string{"LEFT", "KIRI"},
		},
		Right: differ.Subset[string]{
			Index: differ.Range{
				Start: 3,
				End:   4,
			},
			Value: []string{"RIGHT", "KANAN"},
		},
		Kind: differ.CHANGED,
	}

	simpleDiff = `1,2c3,4
< LEFT
< KIRI
---
> RIGHT
> KANAN
`
	simpleDiffCustomPrefix = `1,2c3,4
<<  LEFT
<<  KIRI
---
>>> RIGHT
>>> KANAN
`
)

func TestDeltaStringer(t *testing.T) {
	p := differ.NewPrinter(differ.StrconvString, differ.PrintNumber)
	p.Add(*simpleDelta)

	received := p.String()
	expected := simpleDiff

	if received != expected {
		t.Logf("expected:\n%s\n", expected)
		t.Logf("received:\n%s\n", received)
		t.Fatalf("expected %d characters, received %d characters", len(expected), len(received))
	}

	p = differ.NewPrinter(differ.StrconvString, differ.PrintNumber, differ.PrintPrefix("<<", ">>>"))
	p.Add(*simpleDelta)

	received = p.String()
	expected = simpleDiffCustomPrefix

	if received != expected {
		t.Logf("expected:\n%s\n", expected)
		t.Logf("received:\n%s\n", received)
		t.Fatalf("expected %d characters, received %d characters", len(expected), len(received))
	}
}
