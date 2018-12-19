package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToModelUsingGob(t *testing.T) {
	assert := assert.New(t)

	e := &ProductEntity{
		ID:         1,
		Name:       "焼きそば",
		Price:      500,
		Cost:       100,
		CategoryID: 1,
		Deleted:    false,
	}

	m, err := e.ToModelUsingGob()

	if assert.NoError(err) {
		assert.Equal(e.ID, m.ID)
		assert.Equal(e.Name, m.Name)
		assert.Equal(e.Price, m.Price)
		assert.Equal(e.Cost, m.Cost)
		assert.Equal(0, m.Profit)
		assert.Nil(m.Category)
	}
}

func TestToModelUsingBinary(t *testing.T) {
	assert := assert.New(t)

	e := &ProductEntity{
		ID:         1,
		Name:       "焼きそば",
		Price:      500,
		Cost:       100,
		CategoryID: 1,
		Deleted:    false,
	}

	m, err := e.ToModelUsingBinary()

	if assert.NoError(err) {
		assert.Equal(e.ID, m.ID)
		assert.Equal(e.Name, m.Name)
		assert.Equal(e.Price, m.Price)
		assert.Equal(e.Cost, m.Cost)
		assert.Equal(0, m.Profit)
		assert.Nil(m.Category)
	}
}
