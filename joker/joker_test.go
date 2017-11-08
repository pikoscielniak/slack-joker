package joker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestJokerReturnsThreeJokesAndThenDefaultMsg(t *testing.T) {
	jokes := map[string]bool{}
	for i := 0; i < maxJokes; i++ {
		joke, err := fetch()
		assert.Nil(t, err)
		assert.NotEmpty(t, joke)
		jokes[joke] = true
	}
	msg, err := fetch()
	assert.Nil(t, err)
	assert.Equal(t, defaultMsg, msg)
	assert.Len(t, jokes, maxJokes)
	for _, v := range jokes {
		assert.True(t, v)
	}
}
