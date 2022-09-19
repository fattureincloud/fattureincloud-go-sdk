package filter

import "fmt"

type Disjunction struct {
	Left  Expression
	Right Expression
}

func NewDisjunction(left Expression, right Expression) *Disjunction {
	this := Disjunction{}
	this.Left = left
	this.Right = right
	return &this
}

func (c *Disjunction) BuildQuery() string {
	return fmt.Sprintf("(%s or %s)", c.Left.BuildQuery(), c.Right.BuildQuery())
}

func (c *Disjunction) String() string {
	return c.BuildQuery()
}
