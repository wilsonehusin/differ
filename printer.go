package differ

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

func String[T comparable](d []Delta[T], toString func(T) string, opts ...PrinterOption) string {
	p := NewPrinter[T](toString, opts...)
	p.Add(d...)
	return p.String()
}

func Reader[T comparable](d []Delta[T], toString func(T) string, opts ...PrinterOption) io.Reader {
	p := NewPrinter[T](toString, opts...)
	p.Add(d...)
	return p.Reader()
}

type Printer[T comparable] struct {
	b   *strings.Builder
	opt *PrintOption

	toString func(T) string
}

func NewPrinter[T comparable](toString func(T) string, opts ...PrinterOption) *Printer[T] {
	p := &Printer[T]{
		b: &strings.Builder{},
		opt: &PrintOption{
			prefixLeft:  "<",
			prefixRight: ">",
			prefixEmpty: " ",
		},

		toString: toString,
	}
	p.Configure(opts...)

	return p
}

func (p *Printer[T]) Configure(opts ...PrinterOption) {
	for _, opt := range opts {
		opt(p.opt)
	}
}

func (p *Printer[T]) Add(deltas ...Delta[T]) {
	for _, delta := range deltas {
		sprintf := fmt.Sprintf
		if p.opt.color {
			c := color.New(color.Bold)
			switch delta.Kind {
			case DELETED:
				c.Add(color.FgRed)
			case CHANGED:
				c.Add(color.FgYellow)
			case ADDED:
				c.Add(color.FgGreen)
			}
			sprintf = c.Sprintf
		}

		for _, s := range p.stringers(delta) {
			//nolint:errcheck // strings.Builder.WriteString always returns nil error
			p.b.WriteString(sprintf(s.String()))
		}
	}
}

func (p *Printer[T]) stringers(d Delta[T]) []fmt.Stringer {
	result := []fmt.Stringer{}
	if p.opt.number {
		result = append(result,
			d.Left.Index,
			d.Kind,
			d.Right.Index,
			newliner,
		)
	}
	if d.Kind == EQUAL {
		if p.opt.unified {
			result = append(result, d.Left.Stringer(p.toString, p.opt.prefixEmpty))
		} else {
			result = []fmt.Stringer{}
		}
	} else {
		left := d.Left.Stringer(p.toString, p.opt.prefixLeft)
		right := d.Right.Stringer(p.toString, p.opt.prefixRight)

		if !d.Left.Empty() {
			result = append(result, left)
			if !d.Right.Empty() {
				result = append(result, separatorLiner, newliner, right)
			}
		} else if !d.Right.Empty() {
			result = append(result, right)
		}
	}
	return result
}

func (p *Printer[T]) Reader() io.Reader {
	return strings.NewReader(p.String())
}

func (p *Printer[T]) String() string {
	return p.b.String()
}

type PrinterOption func(p *PrintOption)

type PrintOption struct {
	unified bool
	color   bool
	number  bool

	prefixLeft  string
	prefixRight string
	prefixEmpty string
}

func PrintUnified(p *PrintOption) {
	p.unified = true
}

func PrintColor(p *PrintOption) {
	p.color = true
}

func PrintNumber(p *PrintOption) {
	p.number = true
}

func PrintPrefix(left, right string) PrinterOption {
	return func(p *PrintOption) {
		max := len(left)
		if max < len(right) {
			max = len(right)
		}

		p.prefixLeft = fmt.Sprintf("%-*s", max, left)
		p.prefixRight = fmt.Sprintf("%-*s", max, right)
		p.prefixEmpty = fmt.Sprintf("%-*s", max, " ")
	}
}
