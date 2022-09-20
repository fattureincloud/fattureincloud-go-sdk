package filter

import "fmt"

type Conjunction struct {
	Left  Expression
	Right Expression
}

func NewConjunction(left Expression, right Expression) *Conjunction {
	this := Conjunction{}
	this.Left = left
	this.Right = right
	return &this
}

func (c *Conjunction) BuildQuery() string {
	return fmt.Sprintf("(%s and %s)", c.Left.BuildQuery(), c.Right.BuildQuery())
}

func (c *Conjunction) String() string {
	return c.BuildQuery()
}
