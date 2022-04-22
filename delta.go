package differ

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var (
	newliner       = bytes.NewBufferString("\n")
	separatorLiner = bytes.NewBufferString("---")
)

type Delta[T comparable] struct {
	Left  Subset[T]
	Right Subset[T]

	Kind Kind
}

type Subset[T comparable] struct {
	Index Range
	Value []T
}

func (s Subset[T]) Empty() bool {
	return s.Value == nil || len(s.Value) == 0
}

func (s Subset[T]) Stringer(toString func(T) string, prefix string) *strings.Builder {
	b := &strings.Builder{}
	if s.Value == nil {
		return b
	}

	for _, v := range s.Value {
		if prefix != "" {
			fmt.Fprintf(b, "%s ", prefix)
		}
		fmt.Fprintf(b, "%s\n", toString(v))
	}
	return b
}

type Range struct {
	Start int
	End   int
}

func (r Range) String() string {
	s := strconv.Itoa(r.Start)
	if r.Start != r.End {
		s += "," + strconv.Itoa(r.End)
	}
	return s
}
