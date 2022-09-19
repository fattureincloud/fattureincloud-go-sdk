package filter

import (
	"testing"

	"github.com/fattureincloud/fattureincloud-go-sdk/filter"
	"github.com/stretchr/testify/assert"
)

func TestCondition(t *testing.T) {
	cBool := filter.NewCondition("qualified", filter.Operators.EQ, true)
	assert.Equal(t, "qualified", cBool.Field)
	assert.Equal(t, filter.Operators.EQ, cBool.Operator)
	assert.Equal(t, "qualified = true", cBool.BuildQuery())
	assert.Equal(t, "qualified = true", cBool.String())

	cBool.Field = "phd"
	cBool.Operator = filter.Operators.NEQ
	assert.Equal(t, "phd", cBool.Field)
	assert.Equal(t, filter.Operators.NEQ, cBool.Operator)

	cBool.Value = false
	assert.Equal(t, false, cBool.Value)
	cBool.Value = true
	assert.Equal(t, true, cBool.Value)

	cInt := filter.NewCondition("ranking", filter.Operators.GT, 10)
	assert.Equal(t, 10, cInt.Value)
	assert.Equal(t, filter.Operators.GT, cInt.Operator)
	assert.Equal(t, "ranking > 10", cInt.BuildQuery())
	assert.Equal(t, "ranking > 10", cInt.String())

	cInt.Value = 6
	assert.Equal(t, 6, cInt.Value)
	cInt.Value = 99
	assert.Equal(t, 99, cInt.Value)

	cString := filter.NewCondition("sweet", filter.Operators.ENDS_WITH, "cioccolato")
	assert.Equal(t, "cioccolato", cString.Value)
	assert.Equal(t, filter.Operators.ENDS_WITH, cString.Operator)
	assert.Equal(t, "sweet ends with 'cioccolato'", cString.BuildQuery())
	assert.Equal(t, "sweet ends with 'cioccolato'", cString.String())

	cString.Value = "fragole"
	assert.Equal(t, "fragole", cString.Value)
	cString.Value = "panna"
	assert.Equal(t, "panna", cString.Value)

	cNull := filter.NewCondition("girlfriend", filter.Operators.IS_NOT, nil)
	assert.Equal(t, nil, cNull.Value)
	assert.Equal(t, "girlfriend is not null", cNull.BuildQuery())
	cNull.Operator = filter.Operators.IS
	assert.Equal(t, "girlfriend is null", cNull.String())
}
