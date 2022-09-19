package filter

import "fmt"

type Condition struct {
	Field    string
	Operator Operator
	Value    interface{}
}

func NewCondition(field string, op Operator, value interface{}) *Condition {
	this := Condition{}
	this.Field = field
	this.Operator = op
	this.Value = value
	return &this
}

func (c *Condition) BuildQuery() string {
	if c.Value == nil {
		return fmt.Sprintf("%s %s %s", c.Field, string(c.Operator), "null")
	}
	switch c.Value.(type) {
	case bool:
		return fmt.Sprintf("%s %s %t", c.Field, string(c.Operator), c.Value)
	case int:
		return fmt.Sprintf("%s %s %d", c.Field, string(c.Operator), c.Value)
	case float64:
		return fmt.Sprintf("%s %s %f", c.Field, string(c.Operator), c.Value)
	case string:
		return fmt.Sprintf("%s %s '%s'", c.Field, string(c.Operator), c.Value)
	}
	return fmt.Sprintf("%s %s '%s'", c.Field, string(c.Operator), c.Value)
}

func (c *Condition) String() string {
	return c.BuildQuery()
}
