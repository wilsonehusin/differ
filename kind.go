package differ

type Kind int

const (
	UNKNOWN = iota
	EQUAL
	ADDED
	DELETED
	CHANGED
)

func (k Kind) String() string {
	var str string
	switch k {
	case EQUAL:
		str = "e"
	case ADDED:
		str = "a"
	case DELETED:
		str = "d"
	case CHANGED:
		str = "c"
	default:
		str = "x"
	}
	return str
}
