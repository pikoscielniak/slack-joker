package joker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddShouldAddUntilMaxSize(t *testing.T) {
	q := newRotateQueue(3)
	q.add("a")
	q.add("b")
	q.add("c")

	assert.Len(t, q.values, 3)

	q.add("d")
	assert.Len(t, q.values, 3)
	assert.Equal(t, cap(q.values), 4)

	assert.Equal(t, []string{"b", "c", "d"}, q.values)

	q.add("e")
	assert.Len(t, q.values, 3)
	assert.Equal(t, cap(q.values), 4)

	assert.Equal(t, []string{"c", "d", "e"}, q.values)
}

func TestPopReturnsProperValue(t *testing.T) {
	q := newRotateQueue(3)
	q.add("a")
	q.add("b")
	q.add("c")

	val := q.pop()

	assert.Equal(t, "c", val)
}

func TestPopReturnsProperValueAfterQueueOverflow(t *testing.T) {
	q := newRotateQueue(3)
	q.add("a")
	q.add("b")
	q.add("c")
	q.add("d")

	val := q.pop()

	assert.Equal(t, "d", val)
}

func TestPopReturnsProperValueWhenQueueIsEmpty(t *testing.T) {
	q := newRotateQueue(3)

	val := q.pop()

	assert.Equal(t, "", val)
}

func TestHasJokeReturnsTrueWhenThereIsValue(t *testing.T) {
	q := newRotateQueue(3)

	q.add("a")
	q.add("b")

	assert.True(t, q.hasJoke("a"))
}

func TestHasJokeReturnsFalseWhenThereIsNoValue(t *testing.T) {
	q := newRotateQueue(3)

	q.add("a")
	q.add("b")
	q.add("c")
	q.add("d")

	assert.False(t, q.hasJoke("a"))
}
