package alert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	a := NewSlidingWindowAlert(3)

	assert.False(t, a.AllTrue())
	a.Append(false)

	assert.False(t, a.AllTrue())
	a.Append(false)

	assert.False(t, a.AllTrue())
	a.Append(true)

	assert.False(t, a.AllTrue())
	a.Append(true)

	assert.False(t, a.AllTrue())
	a.Append(true)

	assert.True(t, a.AllTrue())
	a.Append(false)

	assert.False(t, a.AllTrue())
}
