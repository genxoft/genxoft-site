package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	s := "pass"
	v := "1"
	r := "1.0"
	p := NewHealth(s, v, r)
	assert.Equal(t, s, p.Status)
	assert.Equal(t, v, p.Version)
	assert.Equal(t, r, p.ReleaseId)
}
