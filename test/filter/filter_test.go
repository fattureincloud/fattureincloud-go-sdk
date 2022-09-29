package filter

import (
	"testing"

	"github.com/fattureincloud/fattureincloud-go-sdk/v2/filter"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	e := filter.NewCondition("nation", filter.Operators.EQ, "IT")
	f := filter.NewFilter(e)
	assert.Equal(t, e, f.Expression)

	e1 := filter.NewCondition("company", filter.Operators.EQ, "MadBit Entertainment S.r.l.")
	f.Expression = e1
	assert.Equal(t, e1, f.Expression)
}

func TestWhere(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	e := filter.NewCondition("nation", filter.Operators.EQ, "IT")
	f = filter.NewFilter(e)
	assert.Equal(t, e, f.Expression)

	e1 := filter.NewCondition("company", filter.Operators.EQ, "MadBit Entertainment S.r.l.")
	f.Expression = e1
	assert.Equal(t, e1, f.Expression)
}

func TestWhereExpression(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	e := filter.NewCondition("nation", filter.Operators.EQ, "IT")
	f.WhereExpression(e)
	assert.Equal(t, e, f.Expression)

	e1 := filter.NewCondition("company", filter.Operators.EQ, "MadBit Entertainment S.r.l.")
	f.WhereExpression(e1)
	assert.Equal(t, e1, f.Expression)
}

func TestAnd(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	c1 := filter.NewCondition("city", filter.Operators.EQ, "Warsaw")
	c2 := filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media")
	conj := filter.NewConjunction(c1, c2)

	f.Where("city", filter.Operators.EQ, "Warsaw").And("company", filter.Operators.EQ, "Przewodniczka Social Media")

	assert.Equal(t, conj, f.Expression)
}

func TestAndExpression(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	c1 := filter.NewCondition("city", filter.Operators.EQ, "Warsaw")
	c2 := filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media")
	conj := filter.NewConjunction(c1, c2)

	f.WhereExpression(c1).AndExpression(c2)

	assert.Equal(t, conj, f.Expression)
}

func TestAndFilter(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	c1 := filter.NewCondition("city", filter.Operators.EQ, "Warsaw")
	c2 := filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media")
	conj := filter.NewConjunction(c1, c2)

	f.WhereExpression(c1).AndFilter(*filter.NewFilter(c2))

	assert.Equal(t, conj, f.Expression)
}

func TestOr(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	c1 := filter.NewCondition("city", filter.Operators.EQ, "Warsaw")
	c2 := filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media")
	conj := filter.NewDisjunction(c1, c2)

	f.Where("city", filter.Operators.EQ, "Warsaw").Or("company", filter.Operators.EQ, "Przewodniczka Social Media")

	assert.Equal(t, conj, f.Expression)
}

func TestOrExpression(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	c1 := filter.NewCondition("city", filter.Operators.EQ, "Warsaw")
	c2 := filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media")
	conj := filter.NewDisjunction(c1, c2)

	f.WhereExpression(c1).OrExpression(c2)

	assert.Equal(t, conj, f.Expression)
}

func TestOrFilter(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.Expression, nil)

	c1 := filter.NewCondition("city", filter.Operators.EQ, "Warsaw")
	c2 := filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media")
	conj := filter.NewDisjunction(c1, c2)

	f.WhereExpression(c1).OrFilter(*filter.NewFilter(c2))

	assert.Equal(t, conj, f.Expression)
}

func TestBuildQuery(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.BuildQuery(), "")

	f = filter.NewFilter(filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media"))
	assert.Equal(t, "company = 'Przewodniczka Social Media'", f.BuildQuery())
}

func TestBuildUrlEncodedQuery(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.BuildUrlEncodedQuery(), "")

	f = filter.NewFilter(filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media"))
	assert.Equal(t, "company+%3D+%27Przewodniczka+Social+Media%27", f.BuildUrlEncodedQuery())
}

func TestString(t *testing.T) {
	f := filter.NewEmptyFilter()
	assert.Equal(t, f.String(), "")

	f = filter.NewFilter(filter.NewCondition("company", filter.Operators.EQ, "Przewodniczka Social Media"))
	assert.Equal(t, "company = 'Przewodniczka Social Media'", f.String())
}
