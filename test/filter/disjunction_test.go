package filter

import (
	"testing"

	"github.com/fattureincloud/fattureincloud-go-sdk/v2/filter"
	"github.com/stretchr/testify/assert"
)

func TestDisjunction(t *testing.T) {
	left := filter.NewCondition("city", filter.Operators.EQ, "Bergamo")
	right := filter.NewCondition("age", filter.Operators.LT, 35)
	c := filter.NewDisjunction(left, right)
	assert.Equal(t, left, c.Left)
	assert.Equal(t, right, c.Right)
	assert.Equal(t, "(city = 'Bergamo' or age < 35)", c.BuildQuery())
	assert.Equal(t, "(city = 'Bergamo' or age < 35)", c.String())

	left1 := filter.NewCondition("team", filter.Operators.EQ, "Volley Bergamo 1991")
	c.Left = left1
	assert.Equal(t, left1, c.Left)
	assert.Equal(t, "(team = 'Volley Bergamo 1991' or age < 35)", c.BuildQuery())
	assert.Equal(t, "(team = 'Volley Bergamo 1991' or age < 35)", c.String())

	right1 := filter.NewCondition("best_player", filter.Operators.IS_NOT, nil)
	c.Right = right1
	assert.Equal(t, right1, c.Right)
	assert.Equal(t, "(team = 'Volley Bergamo 1991' or best_player is not null)", c.BuildQuery())
	assert.Equal(t, "(team = 'Volley Bergamo 1991' or best_player is not null)", c.String())
}
