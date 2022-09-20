package filter

// Operator
type Operator string

// List of Operators
var Operators = struct {
	EQ          Operator
	GT          Operator
	GTE         Operator
	LT          Operator
	LTE         Operator
	NEQ         Operator
	IS          Operator
	IS_NOT      Operator
	LIKE        Operator
	CONTAINS    Operator
	STARTS_WITH Operator
	ENDS_WITH   Operator
}{
	EQ:          "=",
	GT:          ">",
	GTE:         ">=",
	LT:          "<",
	LTE:         "<=",
	NEQ:         "<>",
	IS:          "is",
	IS_NOT:      "is not",
	LIKE:        "like",
	CONTAINS:    "contains",
	STARTS_WITH: "starts with",
	ENDS_WITH:   "ends with",
}
