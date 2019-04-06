package collections

// NewList creates and returns a new List. If an array of elements is
// passed as a parameter, the List is popuated with those elements.
func NewList(v ...[]int64) List {
	l := newlist()
	for _, a := range v {
		for _, e := range a {
			l.PushBack(e)
		}
	}
	return l
}
