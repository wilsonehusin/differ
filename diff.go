package differ

import "log"

// Diff implements Myers Diff algorithm, slice of Delta
// which contains the merged content of the parameters.
func Diff[T comparable](left, right []T) []Delta[T] {
	// edits is to track the edit index and history
	// In the paper, Myers defined this as array / slice,
	// but here it is map to support negative indexing, i.e. edits[-1]
	edits := map[int]editTracker[T]{}

	edits[0] = editTracker[T]{
		index:   0,
		history: []Delta[T]{},
	}

	max := len(left) + len(right) + 1
	for d := 0; d <= max; d++ {
		for k := -d; k <= d+1; k += 2 {
			// x is index pointer for left
			// y is index pointer for right
			var x, y int

			var prevHistory []Delta[T]

			// downwards mean that an item was inserted on right
			// the opposite (rightwards) mean that an item was deleted on left
			downwards := k == -d || (k != d && edits[k-1].index < edits[k+1].index)

			if downwards {
				x = edits[k+1].index
				prevHistory = edits[k+1].history
			} else {
				x = edits[k-1].index + 1
				prevHistory = edits[k-1].history
			}

			y = x - k

			history := make([]Delta[T], len(prevHistory))
			_ = copy(history, prevHistory)

			prev := &Delta[T]{}
			if i := len(history) - 1; i >= 0 {
				prev = &history[i]
			}

			e := &Delta[T]{
				Left: Subset[T]{
					Index: Range{Start: x, End: x},
					Value: []T{},
				},
				Right: Subset[T]{
					Index: Range{Start: y, End: y},
					Value: []T{},
				},
			}
			if 0 < y && y <= len(right) && downwards {
				target := e
				if prev.Kind == DELETED || prev.Kind == CHANGED {
					target = prev
					target.Kind = CHANGED
				} else if prev.Kind == ADDED {
					target = prev
				} else {
					target.Kind = ADDED
				}
				target.Right.Index.End = y
				target.Right.Value = append(target.Right.Value, right[wrapIndex(y-1, len(right))])
			} else if 0 < x && x <= len(left) {
				target := e
				if prev.Kind == DELETED || prev.Kind == CHANGED {
					target = prev
				} else if prev.Kind == ADDED {
					target = prev
					target.Kind = CHANGED
				} else {
					target.Kind = DELETED
				}
				target.Left.Index.End = x
				target.Left.Value = append(target.Left.Value, left[wrapIndex(x-1, len(left))])
			}
			if e.Kind != UNKNOWN {
				history = append(history, *e)
			}

			rangeX := Range{Start: x}
			rangeY := Range{Start: y}
			diagonal := []T{}
			for x < len(left) && y < len(right) {
				realX := wrapIndex(x, len(left))
				realY := wrapIndex(y, len(right))
				if left[realX] != right[realY] {
					break
				}

				// When values of left[x] and right[y] are the same, greedily advance pointers
				diagonal = append(diagonal, left[realX])
				x = x + 1
				y = y + 1
			}
			if len(diagonal) > 0 {
				rangeX.End = x
				rangeY.End = y
				e := Delta[T]{
					Kind: EQUAL,
					Left: Subset[T]{
						Index: rangeX,
						Value: diagonal,
					},
					Right: Subset[T]{
						Index: rangeY,
						Value: diagonal,
					},
				}
				history = append(history, e)
			}

			if x > len(left) && y > len(right) {
				return history
			}

			edits[k] = editTracker[T]{
				index:   x,
				history: history,
			}
		}
	}
	log.Printf("for loop ended")
	return edits[max-1].history
}

type editTracker[T comparable] struct {
	index   int
	history []Delta[T]
}

func wrapIndex(i, n int) int {
	if i < 0 {
		return i + n
	}
	return i
}
