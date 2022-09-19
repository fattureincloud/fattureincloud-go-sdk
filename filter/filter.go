package filter

import "net/url"

type Filter struct {
	Expression Expression
}

func NewEmptyFilter() *Filter {
	this := Filter{}
	this.Expression = nil
	return &this
}

func NewFilter(e Expression) *Filter {
	this := Filter{}
	this.Expression = e
	return &this
}

func (f *Filter) Where(field string, op Operator, value interface{}) *Filter {
	f.Expression = NewCondition(field, op, value)
	return f
}

func (f *Filter) WhereExpression(e Expression) *Filter {
	f.Expression = e
	return f
}

func (f *Filter) And(field string, op Operator, value interface{}) *Filter {
	f.Expression = NewConjunction(
		f.Expression,
		NewCondition(field, op, value),
	)
	return f
}

func (f *Filter) AndExpression(e Expression) *Filter {
	f.Expression = NewConjunction(
		f.Expression,
		e,
	)
	return f
}

func (f *Filter) AndFilter(filter Filter) *Filter {
	f.Expression = NewConjunction(
		f.Expression,
		filter.Expression,
	)
	return f
}

func (f *Filter) Or(field string, op Operator, value interface{}) *Filter {
	f.Expression = NewDisjunction(
		f.Expression,
		NewCondition(field, op, value),
	)
	return f
}

func (f *Filter) OrExpression(e Expression) *Filter {
	f.Expression = NewDisjunction(
		f.Expression,
		e,
	)
	return f
}

func (f *Filter) OrFilter(filter Filter) *Filter {
	f.Expression = NewDisjunction(
		f.Expression,
		filter.Expression,
	)
	return f
}

func (f *Filter) BuildQuery() string {
	if f.Expression != nil {
		return f.Expression.BuildQuery()
	} else {
		return ""
	}
}

func (f *Filter) BuildUrlEncodedQuery() string {
	if f.Expression != nil {
		return url.QueryEscape(f.BuildQuery())
	} else {
		return ""
	}
}

func (f *Filter) String() string {
	return f.BuildQuery()
}
