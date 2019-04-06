package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var l List

func TestMain(m *testing.M) {
	l = NewList([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	m.Run()
}
func TestPrev(t *testing.T) {
	l.End()
	assert.Equal(t, l.Prev().Val(), int64(9))
	assert.Equal(t, l.Prev().Prev().Val(), int64(7))
}

func TestNext(t *testing.T) {
	l.Begin()
	assert.Equal(t, l.Next().Val(), int64(2))
	assert.Equal(t, l.Next().Next().Val(), int64(4))
}

func TestVal(t *testing.T) {
	assert.Equal(t, l.Begin().Val(), int64(1))
	assert.Equal(t, l.End().Val(), int64(10))
}

func TestLen(t *testing.T) {
	assert.Equal(t, l.Len(), 10)
	l.PushBack(11)
	defer l.PopBack()
	assert.Equal(t, l.Len(), 11)
}

func TestPushBack(t *testing.T) {
	l.PushBack(11)
	defer l.PopBack()
	l.PushBack(12)
	defer l.PopBack()
	l.PushBack(13)
	defer l.PopBack()
	assert.Equal(t, l.End().Val(), int64(13))
	assert.Equal(t, l.Prev().Val(), int64(12))
	assert.Equal(t, l.Prev().Val(), int64(11))
	assert.Equal(t, l.Prev().Prev().Val(), int64(9))
}

func TestPopBack(t *testing.T) {
	l.PopBack()
	defer l.PushBack(10)
	l.PopBack()
	defer l.PushBack(9)
	assert.Equal(t, l.Len(), 8)
	assert.Equal(t, l.End().Val(), int64(8))
	assert.Equal(t, l.Prev().Val(), int64(7))
}

func TestPushFront(t *testing.T) {
	l.PushFront(0)
	defer l.PopFront()
	l.PushFront(-1)
	defer l.PopFront()
	l.PushFront(-2)
	defer l.PopFront()
	assert.Equal(t, l.Begin().Val(), int64(-2))
	assert.Equal(t, l.Next().Val(), int64(-1))
	assert.Equal(t, l.Next().Val(), int64(0))
	assert.Equal(t, l.Next().Next().Val(), int64(2))
}

func TestPopFront(t *testing.T) {
	l.PopFront()
	defer l.PushFront(1)
	l.PopFront()
	defer l.PushFront(2)
	assert.Equal(t, l.Len(), 8)
	assert.Equal(t, l.Begin().Val(), int64(3))
	assert.Equal(t, l.Next().Val(), int64(4))
}

func TestNotNil(t *testing.T) {
	l.Begin()
	assert.True(t, l.NotNil())
	l.PopFront()
	defer l.PushFront(1)
	assert.False(t, l.NotNil())
	l.PopBack()
	defer l.PushBack(10)
	assert.False(t, l.NotNil())
}

func TestEmpty(t *testing.T) {
	assert.False(t, l.Empty())
	for i := 1; i <= 10; i++ {
		l.PopFront()
		defer l.PushFront(int64(i))
	}
	assert.True(t, l.Empty())
	l.PopFront()
	l.PopBack()
	assert.True(t, l.Empty())
}

func TestReverse(t *testing.T) {
	l.Reverse()
	defer l.Reverse()
	assert.Equal(t, l.Begin().Val(), int64(10))
	assert.Equal(t, l.End().Val(), int64(1))
}

func TestToSlice(t *testing.T) {
	exp := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := l.ToSlice()
	assert.Equal(t, len(got), 10)
	assert.Equal(t, exp, got)
}
