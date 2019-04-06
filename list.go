package collections

// List provides functionalities of a doubly linked list
type List interface {
	PushBack(int64)
	PushFront(int64)
	PopBack()
	PopFront()
	Len() int
	Empty() bool
	Prev() List
	Next() List
	Val() int64
	Begin() List
	End() List
	NotNil() bool
	Reverse() List
	ToSlice() []int64
}

type node struct {
	val  int64
	next *node
	prev *node
}

type list struct {
	head *node
	tail *node
	curr *node
	len  int
}

func newnode() *node {
	return &node{0, nil, nil}
}

func newlist() *list {
	return &list{nil, nil, nil, 0}
}

// PushBack adds a new element to the end of the list
func (l *list) PushBack(val int64) {
	nn := newnode()
	nn.val = val
	nn.prev = l.tail
	if l.tail != nil {
		l.tail.next = nn
	}
	l.tail = nn
	if l.head == nil {
		l.head = nn
	}
	l.len++
}

// PushFront adds a new element to the front of the list
func (l *list) PushFront(val int64) {
	nn := newnode()
	nn.val = val
	nn.next = l.head
	if l.head != nil {
		l.head.prev = nn
	}
	l.head = nn
	if l.tail == nil {
		l.tail = nn
	}
	l.len++
}

// Popback removes an element from the end
func (l *list) PopBack() {
	// If the list is empty, do nothing
	if l.Empty() {
		return
	}
	l.len--
	// When the list has only one element
	if l.Len() == 0 {
		l.head = nil
		l.tail = nil
		l.curr = nil
		return
	}
	temp := l.tail
	l.tail = l.tail.prev
	temp.prev = nil
	l.tail.next = nil
	// Reset Current head if it was affected during this operation
	if l.curr == temp {
		l.curr = nil
	}
}

// PopFront removes an element from the front
func (l *list) PopFront() {
	// If the list is empty, do nothing
	if l.Empty() {
		return
	}
	l.len--
	// When the list has only one element
	if l.Len() == 0 {
		l.head = nil
		l.tail = nil
		l.curr = nil
		return
	}
	temp := l.head
	l.head = l.head.next
	temp.next = nil
	l.head.prev = nil
	// Reset Current head if it was affected during this operation
	if l.curr == temp {
		l.curr = nil
	}
}

// Prev sets the current head to the previous element in the list
func (l *list) Prev() List {
	if l.curr == nil {
		return l
	}
	l.curr = l.curr.prev
	return l
}

// Next sets the current pointer to the next element in the list
func (l *list) Next() List {
	if l.curr == nil {
		return l
	}
	l.curr = l.curr.next
	return l
}

// Val returns the element at the the current pointer
func (l *list) Val() int64 {
	return l.curr.val
}

// Begin sets the current pointer to the first element in the list
func (l *list) Begin() List {
	l.curr = l.head
	return l
}

// End sets the current pointer to the last element in the list
func (l *list) End() List {
	l.curr = l.tail
	return l
}

// NotNil checks whether current pointer is not nil
func (l *list) NotNil() bool {
	return l.curr != nil
}

// Len returns the total number of elements in the list
func (l *list) Len() int {
	return l.len
}

// Empty checks whether the list is empty or not
func (l *list) Empty() bool {
	return l.Len() == 0
}

// Reverse reverses the order of elements in the list
func (l *list) Reverse() List {
	if l.Len() <= 1 {
		return l
	}
	for l.Begin(); l.NotNil(); l.Prev() {
		l.curr.prev, l.curr.next = l.curr.next, l.curr.prev
	}
	l.head, l.tail = l.tail, l.head
	return l
}

// ToSlice returns a slice of the elements in the list
func (l *list) ToSlice() []int64 {
	a := make([]int64, 0)
	for l.Begin(); l.NotNil(); l.Next() {
		a = append(a, l.Val())
	}
	return a
}
