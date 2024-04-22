package hw04lrucache

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestRemoveMiddle(t *testing.T) {
	l := NewList()

	l.PushFront(10) // [10]
	l.PushBack(20)  // [10, 20]
	l.PushBack(30)  // [10, 20, 30]

	element := l.Front().Next // 20
	l.Remove(element)         // [10, 30]
	require.Equal(t, 2, l.Len())
	expected := []int{10, 30}
	require.Equal(t, expected, ListToInt(l))
}

func TestRemoveBack(t *testing.T) {
	l := NewList()

	l.PushFront(10) // [10]
	l.PushBack(20)  // [10, 20]
	l.PushBack(30)  // [10, 20, 30]

	element := l.Back() // 30
	l.Remove(element)   // [10, 20]
	require.Equal(t, 2, l.Len())
	expected := []int{10, 20}
	require.Equal(t, expected, ListToInt(l))
}

func TestRemoveFront(t *testing.T) {
	l := NewList()

	l.PushFront(10) // [10]
	l.PushBack(20)  // [10, 20]
	l.PushBack(30)  // [10, 20, 30]

	element := l.Front() // 30
	l.Remove(element)    // [20, 30]
	require.Equal(t, 2, l.Len())
	expected := []int{20, 30}
	require.Equal(t, expected, ListToInt(l))
}

func TestRemoveAll(t *testing.T) {
	l := NewList()

	l.PushFront(10) // [10]
	l.PushBack(20)  // [10, 20]
	l.PushBack(30)  // [10, 20, 30]

	element := l.Front() // 10
	l.Remove(element)    // [20, 30]
	element = l.Front()  // 20
	l.Remove(element)    // [30]
	element = l.Front()  // 30
	l.Remove(element)    // []

	require.Equal(t, 0, l.Len())
	expected := []int{}
	require.Equal(t, expected, ListToInt(l))
}

func TestRemoveEmpty(t *testing.T) {
	l := NewList()
	l.Remove(l.Front()) // []

	require.Equal(t, 0, l.Len())
	expected := []int{}
	require.Equal(t, expected, ListToInt(l))
}

func TestRemoveOneBack(t *testing.T) {
	l := NewList()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)
	l.PushBack(20)
	l.Remove(l.Back())
	l.Remove(l.Back()) // []
	l.Remove(l.Back())
	l.Remove(l.Front())
	l.Remove(l.Front())
	l.Remove(l.Front())
	l.Remove(l.Back())
	require.Equal(t, 0, l.Len())
	expected := []int{}
	require.Equal(t, expected, ListToInt(l))
}

func TestPushFront(t *testing.T) {
	l := NewList()
	l.PushFront(30)
	l.PushFront(20)
	l.PushFront(10)
	expected := []int{10, 20, 30}
	require.Equal(t, expected, ListToInt(l))
	require.Equal(t, 3, l.Len())
}

func TestPushBack(t *testing.T) {
	l := NewList()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	expected := []int{10, 20, 30}
	require.Equal(t, expected, ListToInt(l))
	require.Equal(t, 3, l.Len())
}

func TestPushBackAndFront(t *testing.T) {
	l := NewList()
	l.PushBack(10)
	l.PushFront(20)
	l.PushBack(30)
	l.PushFront(40)
	l.PushBack(50)
	l.PushFront(60)
	expected := []int{60, 40, 20, 10, 30, 50}
	require.Equal(t, expected, ListToInt(l))
	require.Equal(t, 6, l.Len())
}

func TestMoveToFrontFront(t *testing.T) {
	l := NewList()
	l.PushFront(30)
	l.PushFront(20)
	l.PushFront(10)
	expected := []int{10, 20, 30}
	require.Equal(t, expected, ListToInt(l))
	require.Equal(t, 3, l.Len())

	el := l.Front()
	l.MoveToFront(el)
	expected = []int{10, 20, 30}
	require.Equal(t, expected, ListToInt(l))

	el = l.Back()
	l.MoveToFront(el)
	expected = []int{30, 10, 20}
	require.Equal(t, expected, ListToInt(l))

	el = l.Front().Next
	l.MoveToFront(el)
	expected = []int{10, 30, 20}
	require.Equal(t, expected, ListToInt(l))
}

func TestMoveFrontBack(t *testing.T) {
	l := NewList()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.MoveToFront(l.Back())
	expected := []int{30, 10, 20}
	require.Equal(t, expected, ListToInt(l))
	reverse(expected)
	require.Equal(t, expected, ListToIntReverse(l))

	l.MoveToFront(l.Back())
	expected = []int{20, 30, 10}
	require.Equal(t, expected, ListToInt(l))
	reverse(expected)
	require.Equal(t, expected, ListToIntReverse(l))

	l.MoveToFront(l.Back())
	expected = []int{10, 20, 30}
	require.Equal(t, expected, ListToInt(l))
	reverse(expected)
	require.Equal(t, expected, ListToIntReverse(l))
}

func TestMoveFrontMiddle(t *testing.T) {
	l := NewList()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.MoveToFront(l.Front().Next)
	expected := []int{20, 10, 30}
	require.Equal(t, expected, ListToInt(l))
	reverse(expected)
	require.Equal(t, expected, ListToIntReverse(l))

	l.MoveToFront(l.Front().Next)
	expected = []int{10, 20, 30}
	require.Equal(t, expected, ListToInt(l))
	reverse(expected)
	require.Equal(t, expected, ListToIntReverse(l))

	l.MoveToFront(l.Back().Prev)
	expected = []int{20, 10, 30}
	require.Equal(t, expected, ListToInt(l))
	reverse(expected)
	require.Equal(t, expected, ListToIntReverse(l))
}

func ListToInt(l List) []int {
	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	return elems
}

func ListToIntReverse(l List) []int {
	elems := make([]int, 0, l.Len())
	for i := l.Back(); i != nil; i = i.Prev {
		elems = append(elems, i.Value.(int))
	}
	return elems
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
