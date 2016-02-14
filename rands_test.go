// Copyright 2016 Nemanja Zbiljic
//

package rands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunes(t *testing.T) {
	assert.Equal(t, 62, len(alphanumeric))
	assert.Equal(t, 52, len(alphabetic))
	assert.Equal(t, 10, len(numeric))
}

func TestMax(t *testing.T) {
	assert.Equal(t, int64(max.Uint64()), maxInt64)
}
