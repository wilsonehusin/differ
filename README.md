# differ

This package is an importable implementation of [Myers Diff Algorithm][mda-paper] in Go.

> "The easiest way to compare and print two slices in diff edit script format!"
>
> — Someone, hopefully... well at least me.

## Usage

```go
// import "go.husin.dev/differ"

left := []int{1, 2, 3, 1, 2, 2, 1}
right := []int{3, 2, 3, 1, 1, 3}

d := differ.Diff(left, right)

p := differ.NewPrinter(differ.StrconvInt, differ.PrintUnified, differ.PrintNumber)
p.Add(d...)

fmt.Printf("diff: \n%s", p.String())
```

```go
// import "go.husin.dev/differ"

left := strings.Split("the quick brown fox jumps over the lazy dog", " ") 
right := strings.Split("the slow yellow fox jumps over the dog", " ") 

d := differ.Diff(left, right)

_, _ = io.Copy(os.Stdout, differ.Reader(d, differ.StrconvString, differ.PrintNumber, differ.PrintColor))
```

## Not for you?

This library was created because I was not successful in finding the one I need.
More generally, "Here are 2 slices, diff it for me".

However, if you find this library not for you, here are others that I found in my research which might suit your needs:

- [`pkg/diff`](https://github.com/pkg/diff) (I wish I had seen this one before I decided to implement myself, although the ergonomics is a bit awkward like `container/` standard library)
- [`sergi/go-diff`](https://github.com/sergi/go-diff)
- [`r3labs/diff`](https://github.com/r3labs/diff)

## References

- [_An O(ND) Difference Algorithm and Its Variations_ by Eugene Myers][mda-paper]

[mda-paper]: https://neil.fraser.name/writing/diff/myers.pdf
