package filter

import (
	"testing"

	"github.com/fattureincloud/fattureincloud-go-sdk/v2/filter"
	"github.com/stretchr/testify/assert"
)

func TestOperator(t *testing.T) {
	eq := string(filter.Operators.EQ)
	assert.Equal(t, "=", eq)

	neq := string(filter.Operators.NEQ)
	assert.Equal(t, "<>", neq)

	gt := string(filter.Operators.GT)
	assert.Equal(t, ">", gt)

	gte := string(filter.Operators.GTE)
	assert.Equal(t, ">=", gte)

	lt := string(filter.Operators.LT)
	assert.Equal(t, "<", lt)

	lte := string(filter.Operators.LTE)
	assert.Equal(t, "<=", lte)

	is := string(filter.Operators.IS)
	assert.Equal(t, "is", is)

	isNot := string(filter.Operators.IS_NOT)
	assert.Equal(t, "is not", isNot)

	like := string(filter.Operators.LIKE)
	assert.Equal(t, "like", like)

	contains := string(filter.Operators.CONTAINS)
	assert.Equal(t, "contains", contains)

	notLike := string(filter.Operators.NOT_LIKE)
	assert.Equal(t, "not like", notLike)

	notContains := string(filter.Operators.NOT_CONTAINS)
	assert.Equal(t, "not contains", notContains)

	startsWith := string(filter.Operators.STARTS_WITH)
	assert.Equal(t, "starts with", startsWith)

	endsWith := string(filter.Operators.ENDS_WITH)
	assert.Equal(t, "ends with", endsWith)
}
